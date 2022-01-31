package goabu

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"sync"

	"github.com/abu-lang/goabu/config"
	"github.com/abu-lang/goabu/ecarule"
	"github.com/abu-lang/goabu/memory"
	"github.com/abu-lang/goabu/parser"
	antlr_parser "github.com/abu-lang/goabu/parser/antlr"
	"github.com/abu-lang/goabu/stringset"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type State struct {
	Memory memory.Resources
	Pool   []Update
}

type Executer struct {
	memory        memory.ResourceController
	lockMemory    sync.RWMutex
	types         map[string]string
	pool          []Update
	coordinator   execCoordinator
	lockPool      sync.Mutex
	localLibrary  map[string]ecarule.RuleDict
	globalLibrary map[string]ecarule.RuleDict
	lockRules     sync.Mutex

	workingMemory *ast.WorkingMemory
	dataContext   ast.IDataContext

	lexerParserPool sync.Pool

	agent     Agent
	lockAgent sync.Mutex

	logLevel zap.AtomicLevel
	logger   *zap.Logger

	optimistExec   bool
	optimistInput  bool
	lockOptimistic sync.Mutex
}

func NewExecuter(mem memory.ResourceController, rules []string, agt Agent, lc config.LogConfig) (*Executer, error) {
	res := &Executer{
		memory:        mem.Copy(),
		pool:          make([]Update, 0),
		coordinator:   newCoordinator(),
		localLibrary:  make(map[string]ecarule.RuleDict),
		globalLibrary: make(map[string]ecarule.RuleDict),
		agent:         agt,
		lexerParserPool: sync.Pool{
			New: func() interface{} {
				return parser.NewEcaruleLexerParser()
			},
		},
	}
	if res.memory.HasDuplicates() {
		return nil, errors.New("multiple resources have the same name")
	}
	lp := res.lexerParserPool.Get().(*parser.EcaruleLexerParser)
	defer res.lexerParserPool.Put(lp)
	lp.Lexer.RemoveErrorListeners()
	err := validNames(res.memory.ResourceNames(), lp.Lexer)
	if err != nil {
		return nil, err
	}
	res.types = res.memory.Types()
	res.dataContext, res.workingMemory, err = res.newEmptyGruleStructures("this")
	if err != nil {
		return nil, err
	}
	if lc.Encoding == "" {
		lc.Encoding = "console"
	}
	if lc.Encoding != "console" && lc.Encoding != "json" {
		return nil, fmt.Errorf("unsupported log encoding: %s", lc.Encoding)
	}
	zapCfg, ok := config.LogPreset(lc.Encoding).(zap.Config)
	if !ok {
		return nil, errors.New("invalid logger preset")
	}
	res.logLevel = zapCfg.Level
	res.logger, err = zapCfg.Build()
	if err != nil {
		return nil, err
	}
	res.SetLogLevel(lc.Level)
	err = res.AddRules(rules...)
	if err != nil {
		return nil, err
	}
	err = mem.Start()
	if err != nil {
		return nil, err
	}
	go res.receiveInputs()
	err = res.StartAgent()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *Executer) StartAgent() error {
	m.lockAgent.Lock()
	defer m.lockAgent.Unlock()
	err := m.agent.Start()
	if err != nil {
		return err
	}
	go m.receiveExternalActions()
	err = m.agent.Join()
	if err != nil {
		return err
	}
	return nil
}

func (m *Executer) StopAgent() error {
	m.lockAgent.Lock()
	defer m.lockAgent.Unlock()
	return m.agent.Stop()
}

func (m *Executer) SetAgent(agt Agent) error {
	m.lockAgent.Lock()
	defer m.lockAgent.Unlock()
	if m.agent.IsRunning() {
		return errors.New("agent is still running")
	}
	m.agent = agt
	return nil
}

func (m *Executer) TakeState() State {
	m.coordinator.requestWrite(false)
	m.coordinator.fixWorkingSetWrite(stringset.Make())
	m.lockMemory.RLock()
	memCopy := m.memory.Copy().GetResources()
	m.lockMemory.RUnlock()
	m.lockPool.Lock()
	poolCopy := make([]Update, 0, len(m.pool))
	for _, update := range m.pool {
		var updateCopy Update = make([]Assignment, len(update))
		copy(updateCopy, update)
		poolCopy = append(poolCopy, updateCopy)
	}
	m.lockPool.Unlock()
	m.coordinator.confirmWrite()
	m.coordinator.closeWrite()
	return State{Memory: memCopy, Pool: poolCopy}
}

func (m *Executer) DoIfStable(f func()) bool {
	m.coordinator.requestWrite(false)
	m.coordinator.fixWorkingSetWrite(stringset.Make())
	m.lockPool.Lock()
	stable := len(m.pool) == 0
	if stable {
		f()
	}
	m.lockPool.Unlock()
	m.coordinator.confirmWrite()
	m.coordinator.closeWrite()
	return stable
}

func (m *Executer) HasRule(name string) bool {
	m.lockRules.Lock()
	defer m.lockRules.Unlock()
	return m.hasRuleAux(name)
}

func (m *Executer) AddRules(rules ...string) error {
	if len(rules) == 0 {
		return nil
	}
	m.lockRules.Lock()
	defer m.lockRules.Unlock()
	if len(rules) == 1 {
		return m.addRuleAux(rules[0])
	}
	return addList(rules, m.addRuleAux)
}

func (m *Executer) Exec() {
	m.coordinator.requestWrite(m.HasOptimisticExec())
	defer m.coordinator.closeWrite()
	m.lockPool.Lock()
	if len(m.pool) == 0 {
		m.lockPool.Unlock()
		return
	}
	update, index := m.chooseUpdate()
	m.lockPool.Unlock()
	m.logger.Info(fmt.Sprintf("Exec: %v", update), zap.String("act", "exec"), zap.Array("update", updateLogger(update)))
	workingSet := stringset.Make()
	for _, action := range update {
		workingSet.Insert(action.Resource)
	}
	m.coordinator.fixWorkingSetWrite(workingSet)
	m.lockPool.Lock()
	m.removeUpdate(index)
	m.lockPool.Unlock()
	m.execUpdate(update)
	m.logger.Debug("Terminated Exec", zap.String("act", "exec"))
	m.logger.Sync()
}

func (m *Executer) Input(actions string) error {
	parsed, err := m.parseActions(actions)
	if err != nil {
		return err
	}
	workingSet := stringset.Make()
	for _, p := range parsed {
		workingSet.Insert(p.Resource)
	}
	m.coordinator.requestWrite(m.HasOptimisticInput())
	defer m.coordinator.closeWrite()
	m.coordinator.fixWorkingSetWrite(workingSet)
	m.lockMemory.RLock()
	update, err := evalActions(parsed, m.dataContext, m.workingMemory)
	if err != nil {
		m.logger.Panic("Error during input actions evaluation: "+err.Error(),
			zap.String("act", "eval"),
			zap.String("obj", actions))
	}
	m.lockMemory.RUnlock()
	m.logger.Info("Input: "+actions, zap.String("act", "input"), zap.Array("update", updateLogger(update)))
	m.execUpdate(update)
	m.logger.Debug("Processed input", zap.String("act", "input"))
	m.logger.Sync()
	return nil
}

func (m *Executer) LogLevel() int {
	switch m.logLevel.Level() {
	case zapcore.DebugLevel:
		return config.LogDebug
	case zapcore.InfoLevel:
		return config.LogInfo
	case zapcore.WarnLevel:
		return config.LogWarning
	case zapcore.ErrorLevel:
		return config.LogError
	case zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		return config.LogFatal
	}
	// should not be reached
	return -2
}

func (m *Executer) SetLogLevel(l int) {
	m.agent.SetLogLevel(l)
	if l < config.LogDebug {
		l = config.LogDebug
	} else if l > config.LogFatal {
		l = config.LogFatal
	}
	zapLevel := zapcore.InfoLevel
	switch l {
	case config.LogDebug:
		zapLevel = zapcore.DebugLevel
	case config.LogWarning:
		zapLevel = zapcore.WarnLevel
	case config.LogError:
		zapLevel = zapcore.ErrorLevel
	case config.LogFatal:
		zapLevel = zapcore.PanicLevel
	}
	m.logLevel.SetLevel(zapLevel)
}

func (m *Executer) SetOptimisticExec(b bool) {
	m.lockOptimistic.Lock()
	m.optimistExec = b
	m.lockOptimistic.Unlock()
}

func (m *Executer) SetOptimisticInput(b bool) {
	m.lockOptimistic.Lock()
	m.optimistInput = b
	m.lockOptimistic.Unlock()
}

func (m *Executer) HasOptimisticExec() bool {
	m.lockOptimistic.Lock()
	defer m.lockOptimistic.Unlock()
	return m.optimistExec
}

func (m *Executer) HasOptimisticInput() bool {
	m.lockOptimistic.Lock()
	defer m.lockOptimistic.Unlock()
	return m.optimistInput
}

func (m *Executer) chooseUpdate() (Update, int) {
	// TODO: implement other strategies
	return m.pool[0], 0
}

func (m *Executer) execUpdate(update Update) {
	var executed Update
	m.lockMemory.Lock()
	for _, action := range update {
		variable := action.variable
		variable = m.workingMemory.AddVariable(variable)
		currentVal, err := variable.Evaluate(m.dataContext, m.workingMemory)
		if err != nil {
			m.logger.Panic(fmt.Sprintf("Could not evaluate resource %s: %s", action.Resource, err.Error()),
				zap.String("act", "eval_var"),
				zap.String("obj", action.Resource))
		}
		if reflect.DeepEqual(currentVal, action.Value) {
			m.logger.Debug(fmt.Sprintf("Skipping action %v: resource value would not change", action),
				zap.String("act", "assign"),
				zap.Object("action", assignmentLogger(action)))
			continue
		}
		ltype := currentVal.Type()
		rtype := action.Value.Type()
		if !rtype.AssignableTo(ltype) {
			m.logger.DPanic(fmt.Sprintf("Skipping action %v: cannot assign a %v to a %v", action, rtype, ltype),
				zap.String("act", "assign"),
				zap.Object("action", assignmentLogger(action)))
		} else {
			err := variable.Assign(action.Value, m.dataContext, m.workingMemory)
			if err != nil {
				m.logger.Panic("Could not perform assignment: "+err.Error(),
					zap.String("act", "assign"),
					zap.Object("action", assignmentLogger(action)))
			}
			m.memory.Modified(action.Resource)
			executed = append(executed, action)
			m.logger.Debug("Executed action: "+action.String(),
				zap.String("act", "assign"),
				zap.Object("action", assignmentLogger(action)),
				zap.String("evt", action.Resource))
		}
	}
	updates, eActions := m.discovery(executed)
	m.lockMemory.Unlock()
	m.lockPool.Lock()
	m.pool = append(m.pool, updates...)
	m.lockPool.Unlock()
	m.coordinator.confirmWrite()
	m.logger.Info(fmt.Sprintf("Discovery found %d updates", len(updates)),
		zap.String("act", "discovery"),
		zap.Array("updates", poolLogger(updates)))
	if len(eActions) > 0 {
		payload, err := marshalExternalActions(eActions)
		if err != nil {
			m.logger.Panic("Error during external actions marshalling: "+err.Error(),
				zap.String("act", "marshalling"),
				zap.String("obj", "external actions"))
		}
		tentatives := 0
		for {
			err = m.agent.ForAll(payload)
			if err == nil {
				break
			}
			tentatives++
			if tentatives%10 == 0 {
				m.logger.Error(fmt.Sprintf("Failed %d transactions", tentatives),
					zap.String("act", "for_all"),
					zap.Int("transactions", tentatives))
			}
		}
	}
}

func (m *Executer) removeUpdate(index int) {
	m.pool = append(m.pool[:index], m.pool[index+1:len(m.pool)]...)
}

func (m *Executer) discovery(u Update) ([]Update, []externalAction) {
	var newpool []Update
	var extActions []externalAction
	localRules, globalRules := m.activeRules(u)
	for _, rule := range localRules {
		var defaults, tActions Update
		var err error
		if len(rule.DefaultActions) > 0 {
			defaults, err = evalActions(rule.DefaultActions, m.dataContext, m.workingMemory)
			if err != nil {
				m.logger.Panic("Error during default actions evaluation: "+err.Error(),
					zap.String("act", "eval"),
					zap.String("obj", "default actions"))
			}
		}
		tActions, err = condEvalActions(rule.Task.Condition, rule.Task.Actions, m.dataContext, m.workingMemory)
		if err != nil {
			m.logger.Panic("Error during actions evaluation: "+err.Error(),
				zap.String("act", "eval"),
				zap.String("obj", "actions"))
		}
		newpool = appendNonempty(newpool, append(defaults, tActions...))
	}
	for _, rule := range globalRules {
		if len(rule.DefaultActions) > 0 {
			defaults, err := evalActions(rule.DefaultActions, m.dataContext, m.workingMemory)
			if err != nil {
				m.logger.Panic("Error during default actions evaluation: "+err.Error(),
					zap.String("act", "eval"),
					zap.String("obj", "default actions"))
			}
			newpool = append(newpool, defaults)
		}
		ext := m.preEvaluated(rule)
		extActions = append(extActions, ext)
	}
	return newpool, extActions
}

func (m *Executer) activeRules(u Update) (local, global ecarule.RuleDict) {
	local = ecarule.MakeRuleDict()
	global = ecarule.MakeRuleDict()
	m.lockRules.Lock()
	for _, act := range u {
		local.Add(m.localLibrary[act.Resource])
		global.Add(m.globalLibrary[act.Resource])
	}
	m.lockRules.Unlock()
	return local, global
}

// Precondition: rule.Task.External
func (m *Executer) preEvaluated(rule *ecarule.Rule) externalAction {
	res := externalAction{
		CondWorkingSet: stringset.Make(),
		Constants:      make(map[string]interface{}),
		IntConstants:   make(map[string]int64),
		dataContext:    m.dataContext,
		workingMemory:  m.workingMemory,
	}
	res.WorkingSets = make([]stringset.Set, 0, len(rule.Task.Actions))
	for _, action := range rule.Task.Actions {
		res.WorkingSets = append(res.WorkingSets, stringset.Make(action.Resource))
	}
	res.Condition = res.preEvaluatedExpression(rule.Task.Condition, res.CondWorkingSet)
	res.Actions = res.preEvaluatedActions(rule.Task.Actions)
	return res
}

func (m *Executer) hasRuleAux(name string) bool {
	for _, d := range m.localLibrary {
		if d.Has(name) {
			return true
		}
	}
	for _, d := range m.globalLibrary {
		if d.Has(name) {
			return true
		}
	}
	return false
}

func (m *Executer) addRuleAux(r string) error {
	rule, err := m.parseRule(r)
	if err != nil {
		return err
	}
	if m.hasRuleAux(rule.Name) {
		return fmt.Errorf("there is already a rule named %s", rule.Name)
	}

	library := m.localLibrary
	if rule.Task.External {
		library = m.globalLibrary
	}
	for _, evt := range rule.Events {
		if library[evt] == nil {
			var dict ecarule.RuleDict = ecarule.MakeRuleDict()
			library[evt] = dict
		}
		library[evt].Insert(rule)
	}
	m.logger.Debug("Introduced new rule", zap.String("act", "add_rule"), zap.String("obj", r))
	return nil
}

func (m *Executer) addActions(actions string) error {
	parsed, err := m.parseActions(actions)
	if err != nil {
		return err
	}
	update, err := evalActions(parsed, m.dataContext, m.workingMemory)
	if err != nil {
		return err
	}
	m.pool = append(m.pool, update)
	return nil
}

func (m *Executer) addPool(pl []string) error {
	return addList(pl, m.addActions)
}

func (m *Executer) parseRule(r string) (*ecarule.Rule, error) {
	lp := m.lexerParserPool.Get().(*parser.EcaruleLexerParser)
	defer m.lexerParserPool.Put(lp)
	lp.Reset(r)
	tree := lp.Parser.Prule()
	errs := lp.Errors()
	if len(errs) > 0 {
		for _, err := range errs {
			m.logger.Error("error during parsing: "+err.Error(),
				zap.String("act", "parse"),
				zap.String("obj", r))
		}
		m.logger.Sync()
		return nil, errs[0]
	}

	listener := parser.NewEcaruleParserListener(m.types, m.workingMemory)

	m.lockMemory.Lock()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	// update WorkingMemory
	m.workingMemory.IndexVariables()
	m.lockMemory.Unlock()

	errs = listener.Errors()
	if len(errs) > 0 {
		for _, err := range errs {
			m.logger.Error("error during parsing: "+err.Error(),
				zap.String("act", "parse"),
				zap.String("obj", r))
		}
		m.logger.Sync()
		return nil, errs[0]
	}
	return listener.Rule, nil
}

func (m *Executer) parseActions(actions string) ([]ecarule.Action, error) {
	lp := m.lexerParserPool.Get().(*parser.EcaruleLexerParser)
	defer m.lexerParserPool.Put(lp)
	lp.Reset(actions)
	tree := lp.Parser.Actions()
	errs := lp.Errors()
	if len(errs) > 0 {
		for _, err := range errs {
			m.logger.Error("error during parsing: "+err.Error(),
				zap.String("act", "parse"),
				zap.String("obj", actions))
		}
		m.logger.Sync()
		return nil, errs[0]
	}

	listener := parser.NewEcaruleParserListener(m.types, m.workingMemory)

	m.lockMemory.Lock()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	// update WorkingMemory
	m.workingMemory.IndexVariables()
	m.lockMemory.Unlock()

	errs = listener.Errors()
	if len(errs) > 0 {
		for _, err := range errs {
			m.logger.Error("error during parsing: "+err.Error(),
				zap.String("act", "parse"),
				zap.String("obj", actions))
		}
		m.logger.Sync()
		return nil, errs[0]
	}
	return listener.Rule.DefaultActions, nil
}

func (m *Executer) newEmptyGruleStructures(name string) (ast.IDataContext, *ast.WorkingMemory, error) {
	dataContext := ast.NewDataContext()
	resources := m.memory.GetResources()
	err := dataContext.Add(name, &(resources))
	if err != nil {
		return dataContext, nil, err
	}
	kbName := "dummy_" + name
	version := "0.0.0"
	knowledgeBase := &ast.KnowledgeBase{
		Name:          kbName,
		Version:       version,
		RuleEntries:   make(map[string]*ast.RuleEntry),
		WorkingMemory: ast.NewWorkingMemory(kbName, version),
	}
	defunc := &ast.BuiltInFunctions{
		Knowledge:     knowledgeBase,
		WorkingMemory: knowledgeBase.WorkingMemory,
		DataContext:   dataContext,
	}
	err = dataContext.Add("DEFUNC", defunc)
	if err != nil {
		return dataContext, nil, err
	}
	knowledgeBase.InitializeContext(dataContext)
	return dataContext, knowledgeBase.WorkingMemory, nil
}

func validNames(names []string, lexer *antlr_parser.EcaruleLexer) error {
	if len(names) == 0 {
		return errors.New("no resource specified")
	}
	for _, n := range names {
		if n != "this" && n != "ext" {
			lexer.SetInputStream(antlr.NewInputStream(n))
			token := lexer.NextToken()
			if token.GetLine() == 1 && token.GetColumn() == 0 &&
				lexer.GetCharIndex() == len(n) &&
				antlr_parser.EcaruleLexerSIMPLENAME == token.GetTokenType() {
				continue
			}
		}
		return fmt.Errorf(`invalid resource name: "%s"`, n)
	}
	return nil
}

func addList(strs []string, add func(string) error) error {
	var fstErr error
	failed := ""
	for i, s := range strs {
		err := add(s)
		if err != nil {
			failed += strconv.Itoa(i) + ", "
			if fstErr == nil {
				fstErr = err
			}
		}
	}
	if fstErr != nil {
		return fmt.Errorf("could not add elements with indexes %s as %s", failed[:len(failed)-2], fstErr.Error())
	}
	return nil
}
