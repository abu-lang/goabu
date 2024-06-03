// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

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
	"github.com/abu-lang/goabu/stringset"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Executer struct {
	memory         memory.ResourceController
	lockMemory     sync.RWMutex
	types          map[string]string
	pool           []Update
	coordinator    execCoordinator
	updateReceiver chan<- preparedUpdates
	lockPool       sync.Mutex
	ruleLibrary    map[string]ecarule.RuleDict
	lockRules      sync.Mutex
	invariants     []*ast.Expression

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

func NewExecuter(
	mem memory.ResourceController,
	rules []string,
	agt Agent,
	lc config.LogConfig,
	invariants ...string,
) (*Executer, error) {
	res := &Executer{
		memory:      mem.Copy(),
		pool:        make([]Update, 0),
		coordinator: newCoordinator(),
		ruleLibrary: make(map[string]ecarule.RuleDict),
		invariants:  make([]*ast.Expression, 0, len(invariants)),
		agent:       agt,
	}
	if res.memory.HasDuplicates() {
		return nil, errors.New("multiple resources have the same name")
	}
	err := validNames(res.memory.ResourceNames())
	if err != nil {
		return nil, err
	}
	res.types = res.memory.Types()
	res.dataContext, res.workingMemory, err = newEmptyGruleStructures(map[string]memory.Resources{"this": res.memory.GetResources()})
	if err != nil {
		return nil, err
	}
	res.lexerParserPool = sync.Pool{
		New: func() interface{} {
			return parser.New(res.types, res.workingMemory, &res.lockMemory)
		},
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
	err = res.addInvariants(invariants...)
	if err != nil {
		return nil, err
	}
	err = res.AddRules(rules...)
	if err != nil {
		return nil, err
	}
	res.updateReceiver = res.startUpdateReceiver()
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

func (m *Executer) TakeState() (memory.Resources, []Update) {
	m.coordinator.requestWrite(false)
	m.lockMemory.RLock()
	memCopy := m.memory.Copy().GetResources()
	m.lockMemory.RUnlock()
	lock := make(chan bool)
	m.updateReceiver <- preparedUpdates{confirm: lock}
	lock <- false // no updates are added
	poolCopy := make([]Update, 0, len(m.pool))
	for _, update := range m.pool {
		var updateCopy Update = make([]Assignment, len(update))
		copy(updateCopy, update)
		poolCopy = append(poolCopy, updateCopy)
	}
	<-lock
	m.coordinator.closeWrite()
	return memCopy, poolCopy
}

func (m *Executer) DoIfStable(f func()) bool {
	m.coordinator.requestWrite(false)
	lock := make(chan bool)
	m.updateReceiver <- preparedUpdates{confirm: lock}
	lock <- false // no updates are added
	stable := len(m.pool) == 0
	if stable {
		f()
	}
	<-lock
	m.coordinator.closeWrite()
	return stable
}

func (m *Executer) HasRule(name string) bool {
	m.lockRules.Lock()
	defer m.lockRules.Unlock()
	return m.hasRuleAux(name)
}

// AddRules adds a list of GoAbU rules to the node's knowledge base.
func (m *Executer) AddRules(rules ...string) error {
	if len(rules) == 0 {
		return nil
	}
	m.lockRules.Lock()
	defer m.lockRules.Unlock()
	parser := m.lexerParserPool.Get().(ecarule.Parser)
	defer m.lexerParserPool.Put(parser)
	parsedRules, errs := parser.Parse(rules...)
	if len(errs) > 0 {
		for _, err := range errs {
			m.logger.Error("error during parsing: "+err.Error(),
				zap.String("act", "parse"),
				zap.Strings("obj", rules))
		}
		m.logger.Sync()
		return errs[0]
	}
	if len(parsedRules) == 1 {
		return m.addRuleAux(parsedRules[0])
	}
	return addList(parsedRules, m.addRuleAux)
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
	m.logger.Info(fmt.Sprintf("Exec: %v", update), zap.String("act", "exec"), zapUpdate("update", update))
	workingSet := stringset.Make()
	for _, action := range update {
		workingSet.Insert(action.Resource)
	}
	m.coordinator.fixWorkingSetWrite(workingSet)
	m.lockPool.Lock()
	m.removeUpdate(index)
	m.lockPool.Unlock()
	m.lockMemory.Lock()
	var modified stringset.Set
	if len(m.invariants) > 0 {
		copy := m.memory.Extract(workingSet.Slice())
		modified = m.applyUpdate(update, false)
		if !m.invariantsOk() {
			m.memory.Enclose(copy)
			for _, action := range update {
				if modified.Has(action.Resource) {
					modified.Remove(action.Resource)
					m.workingMemory.ResetVariable(action.variable)
				}
			}
			m.lockMemory.Unlock()
			m.coordinator.confirmWrite()
			m.logger.Info(fmt.Sprintf("Exec-Fail: %v would violate the invariants", update),
				zap.String("act", "exec-fail"),
				zapUpdate("update", update))
			return
		}
	} else {
		modified = m.applyUpdate(update, false)
	}
	m.signalModified(modified)
	m.discovery(modified)
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
	m.logger.Info("Input: "+actions, zap.String("act", "input"), zapUpdate("update", update))
	m.lockMemory.Lock()
	m.discovery(m.applyUpdate(update, true))
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

func (m *Executer) removeUpdate(index int) {
	m.pool = append(m.pool[:index], m.pool[index+1:len(m.pool)]...)
}

func (m *Executer) applyUpdate(update Update, input bool) stringset.Set {
	modified := stringset.Make()
	for _, action := range update {
		variable := action.variable
		variable = m.workingMemory.AddVariable(variable)
		currentVal, err := variable.Evaluate(m.dataContext, m.workingMemory)
		if err != nil {
			m.logger.Panic(fmt.Sprintf("Could not evaluate resource %s: %s", action.Resource, err.Error()),
				zap.String("act", "eval_var"),
				zapUpdate("action", action))
		}
		if reflect.DeepEqual(currentVal, action.Value) {
			m.logger.Debug(fmt.Sprintf("Skipping action %v: resource value would not change", action),
				zap.String("act", "assign"),
				zapUpdate("action", action))
			continue
		}
		ltype := currentVal.Type()
		rtype := action.Value.Type()
		if !rtype.AssignableTo(ltype) {
			m.logger.DPanic(fmt.Sprintf("Skipping action %v: cannot assign a %v to a %v", action, rtype, ltype),
				zap.String("act", "assign"),
				zapUpdate("action", action))
		} else {
			err := variable.Assign(action.Value, m.dataContext, m.workingMemory)
			if err != nil {
				m.logger.Panic("Could not perform assignment: "+err.Error(),
					zap.String("act", "assign"),
					zapUpdate("action", action))
			}
			modified.Insert(action.Resource)
			if input {
				m.memory.Modified(action.Resource)
				m.logger.Debug(fmt.Sprintf("Modified resource \"%s\"", action.Resource),
					zap.String("act", "assign"),
					zapUpdate("action", action))
			}
		}
	}
	return modified
}

func (m *Executer) invariantsOk() bool {
	for _, inv := range m.invariants {
		val, err := inv.Evaluate(m.dataContext, m.workingMemory)
		if err != nil {
			m.logger.Panic("Could not evaluate invariant: "+err.Error(),
				zap.String("act", "eval_inv"),
				zap.String("obj", inv.GetGrlText()))
		}
		if !val.Bool() {
			return false
		}
	}
	return true
}

func (m *Executer) signalModified(modified stringset.Set) {
	for r := range modified {
		m.memory.Modified(r)
		m.logger.Debug(fmt.Sprintf("Modified resource \"%s\"", r),
			zap.String("act", "assign"),
			zap.Any(r, m.resourceValue(r).Interface()),
		)
	}
}

func (m *Executer) resourceValue(resource string) reflect.Value {
	// TODO: Simplify resource access
	return reflect.ValueOf(m.memory.GetResources()).FieldByName(m.types[resource]).MapIndex(reflect.ValueOf(resource))
}

// discovery given a set of modified resource names adds to the pool the updates coming from the
// triggered local rules and adds the updates from the global rules in the pools of the other nodes.
func (m *Executer) discovery(modified stringset.Set) {
	updates, wire := m.triggeredActions(modified)
	m.lockMemory.Unlock()
	ok := make(chan bool)
	m.updateReceiver <- preparedUpdates{updates: updates, confirm: ok}
	ok <- true
	<-ok
	m.coordinator.confirmWrite()
	m.logger.Info(fmt.Sprintf("Discovery found %d updates", len(updates)),
		zap.String("act", "discovery"),
		zapUpdates("updates", updates))
	if len(wire.Tasks) > 0 {
		payload, err := marshalWireTasks(wire)
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

// triggeredActions, given a set of modified resources, calculates the local updates and the partially evaluated tasks
// that are to be sent to the other nodes.
func (m *Executer) triggeredActions(modified stringset.Set) ([]Update, wireTasks) {
	var newpool []Update
	var wTask wireTasks
	rules := m.activeRules(modified)
	localResources := stringset.Make()
	for _, rule := range rules {
		for _, task := range rule.LocalTasks {
			tActions, err := condEvalActions(task.Condition, task.Actions, m.dataContext, m.workingMemory)
			if err != nil {
				m.logger.Panic("Error during actions evaluation: "+err.Error(),
					zap.String("act", "eval"),
					zap.String("obj", "actions"))
			}
			newpool = appendNonempty(newpool, tActions)
		}
		for _, task := range rule.RemoteTasks {
			wTask.Tasks = append(wTask.Tasks, task)
			localResources.Add(stringset.Make(task.LocalResources...))
		}
	}
	wTask.Resources = m.memory.Extract(localResources.Slice())
	return newpool, wTask
}

func (m *Executer) activeRules(modified stringset.Set) ecarule.RuleDict {
	res := ecarule.MakeRuleDict()
	m.lockRules.Lock()
	for resource := range modified {
		res.Add(m.ruleLibrary[resource])
	}
	m.lockRules.Unlock()
	return res
}

func (m *Executer) hasRuleAux(name string) bool {
	for _, d := range m.ruleLibrary {
		if d.Has(name) {
			return true
		}
	}
	return false
}

// addRuleAux adds a [ecarule.Rule] to the node's knowledge base.
func (m *Executer) addRuleAux(rule ecarule.Rule) error {
	if m.hasRuleAux(rule.Name) {
		return fmt.Errorf("there is already a rule named %s", rule.Name)
	}
	for _, evt := range rule.Events {
		if m.ruleLibrary[evt] == nil {
			m.ruleLibrary[evt] = ecarule.MakeRuleDict()
		}
		m.ruleLibrary[evt].Insert(&rule)
	}
	m.logger.Debug("Introduced new rule", zap.String("act", "add_rule"), zap.String("obj", rule.Name))
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

// addInvariants constructs the node's invariant by forming a logical conjunction from
// a list of local boolean expressions.
func (m *Executer) addInvariants(invs ...string) error {
	if len(invs) == 0 {
		return nil
	}
	parser := m.lexerParserPool.Get().(ecarule.Parser)
	defer m.lexerParserPool.Put(parser)
	exps, errs := parser.ParseExpressions(invs...)
	if len(errs) > 0 {
		for _, err := range errs {
			m.logger.Error("error during parsing: "+err.Error(),
				zap.String("act", "parse"),
				zap.Strings("obj", invs))
		}
		m.logger.Sync()
		return errs[0]
	}
	for i, exp := range exps {
		val, err := exp.Evaluate(m.dataContext, m.workingMemory)
		if err != nil {
			m.logger.Error("Could not evaluate invariant: "+err.Error(),
				zap.String("act", "eval_inv"),
				zap.String("obj", invs[i]))
			return err
		}
		if val.Kind() != reflect.Bool {
			m.logger.Error("Invariant with non-boolean type",
				zap.String("act", "add_inv"),
				zap.String("obj", invs[i]))
			return fmt.Errorf("type of invariant #%d is not boolean", i)
		}
		m.invariants = append(m.invariants, exp)
	}
	return nil
}

// parseActions parses a series of local actions.
func (m *Executer) parseActions(actions string) ([]ecarule.Action, error) {
	parser := m.lexerParserPool.Get().(ecarule.Parser)
	defer m.lexerParserPool.Put(parser)
	res, errs := parser.ParseActions(actions)
	if len(errs) > 0 {
		for _, err := range errs {
			m.logger.Error("error during parsing: "+err.Error(),
				zap.String("act", "parse"),
				zap.String("obj", actions))
		}
		m.logger.Sync()
		return nil, errs[0]
	}

	return res, nil
}

// newEmptyGruleStructures creates a clean working memory and data context containing
// the [memory.Resources] from resources as struct instances referenced by the map keys.
func newEmptyGruleStructures(resources map[string]memory.Resources) (ast.IDataContext, *ast.WorkingMemory, error) {
	dataContext := ast.NewDataContext()
	kbName := "dummy"
	for name, rs := range resources {
		rs := rs
		kbName += "_" + name
		err := dataContext.Add(name, &(rs))
		if err != nil {
			return dataContext, nil, err
		}
	}
	version := "0.0.0"
	knowledgeBase := &ast.KnowledgeBase{
		Name:          kbName,
		Version:       version,
		RuleEntries:   make(map[string]*ast.RuleEntry),
		WorkingMemory: ast.NewWorkingMemory(kbName, version),
	}
	err := dataContext.Add("DEFUNC",
		makeBuiltinFunctions(
			knowledgeBase,
			knowledgeBase.WorkingMemory,
			dataContext,
		))
	if err != nil {
		return dataContext, nil, err
	}
	knowledgeBase.InitializeContext(dataContext)
	return dataContext, knowledgeBase.WorkingMemory, nil
}

// validNames verifies if the arguments can be used as valid identifiers in GoAbU rules.
func validNames(names []string) error {
	if len(names) == 0 {
		return errors.New("no resource specified")
	}
	results := parser.ValidateIdentifiers(names...)
	for i, n := range names {
		if !results[i] {
			return fmt.Errorf(`invalid resource name: "%s"`, n)
		}
	}
	return nil
}

func addList[T any](objs []T, add func(T) error) error {
	var fstErr error
	failed := ""
	for i, obj := range objs {
		err := add(obj)
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
