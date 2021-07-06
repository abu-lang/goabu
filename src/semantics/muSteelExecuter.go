package semantics

import (
	"errors"
	"fmt"
	"reflect"
	"steel-lang/datastructure"
	"steel-lang/misc"
	"sync"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

type State struct {
	Memory datastructure.Resources
	Pool   [][]SemanticAction
}

type MuSteelExecuter struct {
	memory        datastructure.ResourceController
	types         map[string]string
	pool          [][]SemanticAction
	lockPool      sync.Mutex
	parsedActions int
	localLibrary  map[string]datastructure.RuleDict
	globalLibrary map[string]datastructure.RuleDict

	knowledgeLibrary *ast.KnowledgeLibrary
	workingMemory    *ast.WorkingMemory
	dataContext      ast.IDataContext

	agent ISteelAgent
}

func NewMuSteelExecuter(mem datastructure.ResourceController, rules []datastructure.Rule, agt ISteelAgent) (*MuSteelExecuter, error) {
	res := &MuSteelExecuter{
		memory:           mem.Clone(),
		pool:             make([][]SemanticAction, 0),
		parsedActions:    0,
		localLibrary:     make(map[string]datastructure.RuleDict),
		globalLibrary:    make(map[string]datastructure.RuleDict),
		knowledgeLibrary: ast.NewKnowledgeLibrary(),
		dataContext:      ast.NewDataContext(),
		agent:            agt,
	}
	if !res.memory.IsValid() {
		return nil, errors.New("invalid Resources argument")
	}
	res.types = res.memory.GetTypes()
	resources := res.memory.GetResources()
	err := res.dataContext.Add("this", &(resources))
	if err != nil {
		return nil, err
	}
	res.AddRules(rules)
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
	memCopy := m.memory.Clone().GetResources()
	m.lockPool.Lock()
	poolCopy := make([][]SemanticAction, 0, len(m.pool))
	for _, acts := range m.pool {
		actsCopy := make([]SemanticAction, len(acts))
		copy(actsCopy, acts)
		poolCopy = append(poolCopy, actsCopy)
	}
	m.lockPool.Unlock()
	return State{Memory: memCopy, Pool: poolCopy}
}

func (m *MuSteelExecuter) IsStable() bool {
	m.lockPool.Lock()
	defer m.lockPool.Unlock()
	return len(m.pool) == 0
}

func (m *MuSteelExecuter) AddRule(rule datastructure.Rule) {
	parsed := datastructure.NewParsedRule(rule, m.knowledgeLibrary, m.types)
	m.updateWorkingMemory()
	library := m.localLibrary
	if parsed.Task.Mode != "for" {
		library = m.globalLibrary
	}
	for _, evt := range parsed.Events {
		if library[evt] == nil {
			var dict datastructure.RuleDict = datastructure.MakeRuleDict()
			library[evt] = dict
		}
		library[evt].Insert(parsed)
	}
}

func (m *MuSteelExecuter) AddRules(rules []datastructure.Rule) {
	for _, rule := range rules {
		m.AddRule(rule)
	}
}

func (m *MuSteelExecuter) AddActions(actions []datastructure.Action) {
	m.lockPool.Lock()
	m.pool = append(m.pool, evalActions(m.parseActions(actions), m.dataContext, m.workingMemory))
	m.lockPool.Unlock()
}

func (m *MuSteelExecuter) AddPool(pl [][]datastructure.Action) {
	for _, actions := range pl {
		m.AddActions(actions)
	}
}

func (m *MuSteelExecuter) Exec() {
	m.lockPool.Lock()
	if len(m.pool) == 0 || m.workingMemory == nil { // nil workingMemory => m does not have rules nor parsed actions
		m.lockPool.Unlock()
		return
	}
	actions, index := m.chooseActions()
	m.removeActions(index)
	m.lockPool.Unlock()
	fmt.Print("Exec: ")
	m.execActions(actions)
}

func (m *MuSteelExecuter) Input(actions []datastructure.Action) {
	sActions := evalActions(m.parseActions(actions), m.dataContext, m.workingMemory)
	fmt.Print("Input: ")
	m.execActions(sActions)
}

func (m *MuSteelExecuter) receiveInputs() {
	inputs := m.memory.Inputs()
	for {
		action := <-inputs
		m.Input([]datastructure.Action{action})
	}
}

func (m *MuSteelExecuter) receiveExternalActions() {
	requests, commandRequests := m.agent.ReceivedActions()
	for {
		actionsCh := <-requests
		if actionsCh == nil {
			return
		}
		commandsCh := <-commandRequests
		eActions, err := unmarshalExternalActions(<-actionsCh, m.types)
		if err != nil {
			panic(err)
		}
		var sActions [][]SemanticAction
		m.lockPool.Lock()
		localResources := m.memory.ResourceNames()
		context, workMem := m.NewEmptyGruleStructures("ext")
		for _, eAction := range eActions {
			if localResources.ContainsSet(eAction.CondWorkingSet) {
				actions := eAction.cullActions(localResources)
				if len(actions) == 0 {
					continue
				}
				sActions = appendNonempty(sActions, condEvalActions(eAction.Condition, actions, context, workMem))
			}
		}
		if len(sActions) == 0 {
			commandsCh <- "not_interested"
			m.lockPool.Unlock()
			continue
		}
		commandsCh <- "interested"
		switch <-commandsCh {
		case "do_commit":
			m.pool = append(m.pool, sActions...)
			fallthrough
		case "do_abort":
			commandsCh <- "done"
		}
		m.lockPool.Unlock()
	}
}

func (m *MuSteelExecuter) chooseActions() ([]SemanticAction, int) {
	// TODO: implement other strategies
	return m.pool[0], 0
}

func (m *MuSteelExecuter) execActions(actions []SemanticAction) {
	m.lockPool.Lock()
	var Xset []SemanticAction
	for _, action := range actions {
		variable := action.Variable
		variable = m.workingMemory.AddVariable(variable)
		currentVal, err := variable.Evaluate(m.dataContext, m.workingMemory)
		if err != nil {
			panic(err)
		}
		diff := false
		if currentVal.Kind() == reflect.Interface || action.Value.Kind() == reflect.Interface {
			diff = true
		} else {
			eq, err := pkg.EvaluateEqual(currentVal, action.Value)
			if err != nil {
				panic(err)
			}
			if !eq.Bool() {
				diff = true
				ltype := currentVal.Type()
				rtype := action.Value.Type()
				if !rtype.AssignableTo(ltype) {
					panic(fmt.Errorf("cannot assign a %v to a %v", rtype, ltype))
				}
			}
		}
		if diff {
			err := variable.Assign(action.Value, m.dataContext, m.workingMemory)
			if err != nil {
				panic(err)
			}
			err = m.memory.Modified(action.Resource)
			if err != nil {
				panic(err)
			}
			Xset = append(Xset, action)
			fmt.Print(action)
		}
	}
	fmt.Println()
	sActions, eActions := m.discovery(Xset)
	m.pool = append(m.pool, sActions...)
	m.lockPool.Unlock()
	if len(eActions) > 0 {
		payload, err := marshalExternalActions(eActions)
		if err == nil {
			err = m.agent.ForAll(payload)
		}
		if err != nil {
			fmt.Println(err.Error())
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
		if len(rule.DefaultActions) > 0 {
			newpool = append(newpool, evalActions(rule.DefaultActions, m.dataContext, m.workingMemory))
		}
		newpool = appendNonempty(newpool, condEvalActions(rule.Task.Condition, rule.Task.Actions, m.dataContext, m.workingMemory))
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

func (m *MuSteelExecuter) activeRules(Xset []SemanticAction) (local, global datastructure.RuleDict) {
	local = datastructure.MakeRuleDict()
	global = datastructure.MakeRuleDict()
	for _, act := range Xset {
		local.Add(m.localLibrary[act.Resource])
		global.Add(m.globalLibrary[act.Resource])
	}
	return local, global
}

// Precondition: rule.Task.Mode != "for"
func (m *MuSteelExecuter) preEvaluated(rule *datastructure.ParsedRule) externalAction {
	res := externalAction{
		CondWorkingSet: misc.MakeStringSet(""),
		Constants:      make(map[string]interface{}),
		IntConstants:   make(map[string]int64),
		dataContext:    m.dataContext,
		workingMemory:  m.workingMemory,
	}
	res.WorkingSets = make([]misc.StringSet, 0, len(rule.Task.Actions))
	for _, action := range rule.Task.Actions {
		res.WorkingSets = append(res.WorkingSets, misc.MakeStringSet(action.Resource))
	}
	res.Condition = res.preEvaluatedExpression(rule.Task.Condition, res.CondWorkingSet)
	res.Actions = res.preEvaluatedActions(rule.Task.Actions)
	return res
}

func (m *MuSteelExecuter) updateWorkingMemory() {
	knowledgeBase := m.knowledgeLibrary.NewKnowledgeBaseInstance("dummy", "0.0.0")
	if knowledgeBase == nil {
		return
	}
	defunc := &ast.BuiltInFunctions{
		Knowledge:     knowledgeBase,
		WorkingMemory: knowledgeBase.WorkingMemory,
		DataContext:   m.dataContext,
	}
	m.dataContext.Add("DEFUNC", defunc)
	knowledgeBase.InitializeContext(m.dataContext)
	m.workingMemory = knowledgeBase.WorkingMemory
}

func (m *MuSteelExecuter) NewEmptyGruleStructures(name string) (ast.IDataContext, *ast.WorkingMemory) {
	dataContext := ast.NewDataContext()
	resources := m.memory.GetResources()
	err := dataContext.Add(name, &(resources))
	if err != nil {
		panic(err)
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
	dataContext.Add("DEFUNC", defunc)
	knowledgeBase.InitializeContext(dataContext)
	return dataContext, knowledgeBase.WorkingMemory
}

func (m *MuSteelExecuter) PrintState() string {
	return fmt.Sprintf("Memory: %v\nPool: %v\n", m.memory, m.printPool())
}

func (m *MuSteelExecuter) printPool() string {
	m.lockPool.Lock()
	defer m.lockPool.Unlock()
	if len(m.pool) == 0 {
		return "{}"
	} else {
		str := "{"
		for _, list := range m.pool {
			str = str + "\n  "
			for _, action := range list {
				str = str + action.String()
			}
		}
		return str + "\n}"
	}
}
