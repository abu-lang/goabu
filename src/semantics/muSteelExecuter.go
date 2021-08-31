package semantics

import (
	"errors"
	"fmt"
	"reflect"
	"steel-lang/config"
	"steel-lang/ecarule"
	"steel-lang/memory"
	"steel-lang/parser"
	"steel-lang/stringset"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type State struct {
	Memory memory.Resources
	Pool   [][]SemanticAction
}

type MuSteelExecuter struct {
	memory        memory.ResourceController
	lockMemory    sync.RWMutex
	types         map[string]string
	pool          [][]SemanticAction
	coordinator   execCoordinator
	lockPool      sync.Mutex
	localLibrary  map[string]ecarule.RuleDict
	globalLibrary map[string]ecarule.RuleDict
	lockRules     sync.Mutex

	workingMemory *ast.WorkingMemory
	dataContext   ast.IDataContext

	lexerParserPool sync.Pool

	agent ISteelAgent

	logLevel zap.AtomicLevel
	logger   *zap.Logger

	optimistExec   bool
	optimistInput  bool
	lockOptimistic sync.Mutex
}

func NewMuSteelExecuter(mem memory.ResourceController, rules []string, agt ISteelAgent, lc config.LogConfig) (*MuSteelExecuter, error) {
	res := &MuSteelExecuter{
		memory:        mem.Copy(),
		pool:          make([][]SemanticAction, 0),
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
	if !res.memory.IsValid() {
		return nil, errors.New("invalid Resources argument")
	}
	res.types = res.memory.GetTypes()
	var err error
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
	err = res.AddRules(rules)
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

func (m *MuSteelExecuter) StartAgent() error {
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

func (m *MuSteelExecuter) StopAgent() error {
	return m.agent.Stop()
}

func (m *MuSteelExecuter) SetAgent(agt ISteelAgent) error {
	if m.agent.IsRunning() {
		return errors.New("agent is still running")
	}
	m.agent = agt
	return nil
}

func (m *MuSteelExecuter) GetState() State {
	m.coordinator.requestWrite(false)
	m.coordinator.fixWorkingSetWrite(stringset.Make(""))
	m.lockMemory.RLock()
	memCopy := m.memory.Copy().GetResources()
	m.lockMemory.RUnlock()
	m.lockPool.Lock()
	poolCopy := make([][]SemanticAction, 0, len(m.pool))
	for _, acts := range m.pool {
		actsCopy := make([]SemanticAction, len(acts))
		copy(actsCopy, acts)
		poolCopy = append(poolCopy, actsCopy)
	}
	m.lockPool.Unlock()
	m.coordinator.confirmWrite()
	m.coordinator.closeWrite()
	return State{Memory: memCopy, Pool: poolCopy}
}

func (m *MuSteelExecuter) DoIfStable(f func()) bool {
	m.coordinator.requestWrite(false)
	m.coordinator.fixWorkingSetWrite(stringset.Make(""))
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

func (m *MuSteelExecuter) HasRule(name string) bool {
	m.lockRules.Lock()
	defer m.lockRules.Unlock()
	return m.hasRuleAux(name)
}

func (m *MuSteelExecuter) AddRule(r string) error {
	m.lockRules.Lock()
	defer m.lockRules.Unlock()
	return m.addRuleAux(r)
}

func (m *MuSteelExecuter) AddRules(rules []string) error {
	m.lockRules.Lock()
	defer m.lockRules.Unlock()
	return addList(rules, m.addRuleAux)
}

func (m *MuSteelExecuter) Exec() {
	m.coordinator.requestWrite(m.HasOptimisticExec())
	defer m.coordinator.closeWrite()
	m.lockPool.Lock()
	if len(m.pool) == 0 {
		m.lockPool.Unlock()
		return
	}
	actions, index := m.chooseActions()
	m.lockPool.Unlock()
	m.logger.Info(fmt.Sprintf("Exec: %v", actions), zap.String("act", "exec"), zap.Array("actions", updateLogger(actions)))
	workingSet := stringset.Make("")
	for _, action := range actions {
		workingSet.Insert(action.Resource)
	}
	m.coordinator.fixWorkingSetWrite(workingSet)
	m.lockPool.Lock()
	m.removeActions(index)
	m.lockPool.Unlock()
	m.execActions(actions)
	m.logger.Debug("Terminated Exec", zap.String("act", "exec"))
	m.logger.Sync()
}

func (m *MuSteelExecuter) Input(actions string) error {
	parsed, err := m.parseActions(actions)
	if err != nil {
		return err
	}
	workingSet := stringset.Make("")
	for _, p := range parsed {
		workingSet.Insert(p.Resource)
	}
	m.coordinator.requestWrite(m.HasOptimisticInput())
	defer m.coordinator.closeWrite()
	m.coordinator.fixWorkingSetWrite(workingSet)
	m.lockMemory.RLock()
	sActions := evalActions(parsed, m.dataContext, m.workingMemory)
	m.lockMemory.RUnlock()
	m.logger.Info("Input: "+actions, zap.String("act", "input"), zap.Array("actions", updateLogger(sActions)))
	m.execActions(sActions)
	m.logger.Debug("Processed input", zap.String("act", "input"))
	m.logger.Sync()
	return nil
}

func (m *MuSteelExecuter) LogLevel() int {
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

func (m *MuSteelExecuter) SetLogLevel(l int) {
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

func (m *MuSteelExecuter) SetOptimisticExec(b bool) {
	m.lockOptimistic.Lock()
	m.optimistExec = b
	m.lockOptimistic.Unlock()
}

func (m *MuSteelExecuter) SetOptimisticInput(b bool) {
	m.lockOptimistic.Lock()
	m.optimistInput = b
	m.lockOptimistic.Unlock()
}

func (m *MuSteelExecuter) HasOptimisticExec() bool {
	m.lockOptimistic.Lock()
	defer m.lockOptimistic.Unlock()
	return m.optimistExec
}

func (m *MuSteelExecuter) HasOptimisticInput() bool {
	m.lockOptimistic.Lock()
	defer m.lockOptimistic.Unlock()
	return m.optimistInput
}

func (m *MuSteelExecuter) chooseActions() ([]SemanticAction, int) {
	// TODO: implement other strategies
	return m.pool[0], 0
}

func (m *MuSteelExecuter) execActions(actions []SemanticAction) {
	var Xset []SemanticAction
	m.lockMemory.Lock()
	for _, action := range actions {
		variable := action.Variable
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
				zap.Object("action", actionLogger(action)))
			continue
		}
		ltype := currentVal.Type()
		rtype := action.Value.Type()
		if !rtype.AssignableTo(ltype) {
			m.logger.DPanic(fmt.Sprintf("Skipping action %v: cannot assign a %v to a %v", action, rtype, ltype),
				zap.String("act", "assign"),
				zap.Object("action", actionLogger(action)))
		} else {
			err := variable.Assign(action.Value, m.dataContext, m.workingMemory)
			if err != nil {
				m.logger.Panic("Could not perform assingment: "+err.Error(),
					zap.String("act", "assign"),
					zap.Object("action", actionLogger(action)))
			}
			m.memory.Modified(action.Resource)
			Xset = append(Xset, action)
			m.logger.Debug("Executed action: "+action.String(),
				zap.String("act", "assign"),
				zap.Object("action", actionLogger(action)),
				zap.String("evt", action.Resource))
		}
	}
	sActions, eActions := m.discovery(Xset)
	m.lockMemory.Unlock()
	m.lockPool.Lock()
	m.pool = append(m.pool, sActions...)
	m.lockPool.Unlock()
	m.coordinator.confirmWrite()
	m.logger.Info(fmt.Sprintf("Discovery found %d updates", len(sActions)),
		zap.String("act", "discovery"),
		zap.Array("updates", poolLogger(sActions)))
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

func (m *MuSteelExecuter) removeActions(index int) {
	m.pool = append(m.pool[:index], m.pool[index+1:len(m.pool)]...)
}

func (m *MuSteelExecuter) discovery(Xset []SemanticAction) ([][]SemanticAction, []externalAction) {
	var newpool [][]SemanticAction
	var extActions []externalAction
	localRules, globalRules := m.activeRules(Xset)
	for _, rule := range localRules {
		var defaults []SemanticAction
		if len(rule.DefaultActions) > 0 {
			defaults = evalActions(rule.DefaultActions, m.dataContext, m.workingMemory)
		}
		newpool = appendNonempty(newpool,
			append(defaults, condEvalActions(rule.Task.Condition, rule.Task.Actions, m.dataContext, m.workingMemory)...))
	}
	for _, rule := range globalRules {
		if len(rule.DefaultActions) > 0 {
			newpool = append(newpool, evalActions(rule.DefaultActions, m.dataContext, m.workingMemory))
		}
		ext := m.preEvaluated(rule)
		extActions = append(extActions, ext)
	}
	return newpool, extActions
}

func (m *MuSteelExecuter) activeRules(Xset []SemanticAction) (local, global ecarule.RuleDict) {
	local = ecarule.MakeRuleDict()
	global = ecarule.MakeRuleDict()
	m.lockRules.Lock()
	for _, act := range Xset {
		local.Add(m.localLibrary[act.Resource])
		global.Add(m.globalLibrary[act.Resource])
	}
	m.lockRules.Unlock()
	return local, global
}

// Precondition: rule.Task.Mode != "for"
func (m *MuSteelExecuter) preEvaluated(rule *ecarule.Rule) externalAction {
	res := externalAction{
		CondWorkingSet: stringset.Make(""),
		Constants:      make(map[string]interface{}),
		IntConstants:   make(map[string]int64),
		dataContext:    m.dataContext,
		workingMemory:  m.workingMemory,
	}
	res.WorkingSets = make([]stringset.StringSet, 0, len(rule.Task.Actions))
	for _, action := range rule.Task.Actions {
		res.WorkingSets = append(res.WorkingSets, stringset.Make(action.Resource))
	}
	res.Condition = res.preEvaluatedExpression(rule.Task.Condition, res.CondWorkingSet)
	res.Actions = res.preEvaluatedActions(rule.Task.Actions)
	return res
}

func (m *MuSteelExecuter) hasRuleAux(name string) bool {
	for _, d := range m.localLibrary {
		if d.Contains(name) {
			return true
		}
	}
	for _, d := range m.globalLibrary {
		if d.Contains(name) {
			return true
		}
	}
	return false
}

func (m *MuSteelExecuter) addRuleAux(r string) error {
	rule, err := m.parseRule(r)
	if err != nil {
		return err
	}
	if m.hasRuleAux(rule.Name) {
		return fmt.Errorf("there is already a rule named %s", rule.Name)
	}

	library := m.localLibrary
	if rule.Task.Mode != "for" {
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

func (m *MuSteelExecuter) addActions(actions string) error {
	parsed, err := m.parseActions(actions)
	if err != nil {
		return err
	}
	m.pool = append(m.pool, evalActions(parsed, m.dataContext, m.workingMemory))
	return nil
}

func (m *MuSteelExecuter) addPool(pl []string) error {
	return addList(pl, m.addActions)
}

func (m *MuSteelExecuter) parseRule(r string) (*ecarule.Rule, error) {
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

func (m *MuSteelExecuter) parseActions(actions string) ([]ecarule.Action, error) {
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

func (m *MuSteelExecuter) newEmptyGruleStructures(name string) (ast.IDataContext, *ast.WorkingMemory, error) {
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
