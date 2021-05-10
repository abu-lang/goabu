package semantics

import (
	"errors"
	"fmt"
	"reflect"
	"steel-lang/communication"
	"steel-lang/datastructure"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

type State struct {
	Memory datastructure.Resources
	Pool   [][]SemanticAction
}

type MuSteelExecuter struct {
	memory        datastructure.Resources
	types         map[string]string
	pool          [][]SemanticAction
	extPool       []ExternalAction
	parsedActions int
	localLibrary  map[string]datastructure.RuleDict
	globalLibrary map[string]datastructure.RuleDict

	knowledgeLibrary *ast.KnowledgeLibrary
	workingMemory    *ast.WorkingMemory
	dataContext      ast.IDataContext

	agent communication.ISteelAgent
}

func NewMuSteelExecuter(mem datastructure.Resources, agt communication.ISteelAgent) (*MuSteelExecuter, error) {
	res := &MuSteelExecuter{
		memory:           mem.Clone(),
		pool:             make([][]SemanticAction, 0),
		extPool:          make([]ExternalAction, 0),
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
	err := res.dataContext.Add("this", &(res.memory))
	if err != nil {
		return nil, err
	}
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
	err = m.agent.Join()
	if err != nil {
		return err
	}
	return nil
}

func (m *MuSteelExecuter) StopAgent() error {
	return m.agent.Stop()
}

func (m *MuSteelExecuter) SetAgent(agt communication.ISteelAgent) error {
	if m.agent.IsRunning() {
		return errors.New("agent is still running")
	}
	m.agent = agt
	return nil
}

func (m *MuSteelExecuter) GetState() State {
	memCopy := m.memory.Clone()
	poolCopy := make([][]SemanticAction, 0, len(m.pool))
	for _, acts := range m.pool {
		actsCopy := make([]SemanticAction, len(acts))
		copy(actsCopy, acts)
		poolCopy = append(poolCopy, actsCopy)
	}
	return State{Memory: memCopy, Pool: poolCopy}
}

func (m *MuSteelExecuter) IsStable() bool {
	return len(m.pool) == 0
}

func (m *MuSteelExecuter) AddRule(rule *datastructure.Rule) {
	parsed := datastructure.NewParsedRule(rule, m.knowledgeLibrary, m.types)
	m.updateWorkingMemory()
	library := m.localLibrary
	if parsed.Task.Mode != "for" {
		library = m.globalLibrary
	}
	for _, evt := range parsed.Event {
		if library[evt] == nil {
			var dict datastructure.RuleDict = datastructure.MakeRuleDict()
			library[evt] = dict
		}
		library[evt].Insert(parsed)
	}
}

func (m *MuSteelExecuter) AddRules(rules []datastructure.Rule) {
	for _, rule := range rules {
		m.AddRule(&rule)
	}
}

func (m *MuSteelExecuter) AddActions(actions []datastructure.Action) {
	m.pool = append(m.pool, evalActions(m.parseActions(actions), m.dataContext, m.workingMemory))
}

func (m *MuSteelExecuter) AddPool(pl [][]datastructure.Action) {
	for _, actions := range pl {
		m.AddActions(actions)
	}
}

func (m *MuSteelExecuter) Exec() {
	if m.IsStable() || m.workingMemory == nil { // nil workingMemory => m does not have rules nor parsed actions
		return
	}
	actions, index := m.chooseActions()
	fmt.Print("Exec: ")
	m.execActions(actions)
	m.removeActions(index)
}

func (m *MuSteelExecuter) Input(actions []datastructure.Action) {
	sActions := evalActions(m.parseActions(actions), m.dataContext, m.workingMemory)
	fmt.Print("Input: ")
	m.execActions(sActions)
}

func (m *MuSteelExecuter) TestExtPool() {
	if len(m.extPool) == 0 {
		return
	}
	fmt.Print("Ext: ")
	context, workMem := m.NewEmptyGruleStructures("ext")
	extAction := m.extPool[0]
	m.extPool = m.extPool[1:]
	fmt.Println(extAction)
	if extAction.DefaultActions != nil {
		m.pool = append(m.pool, evalActions(extAction.DefaultActions, context, workMem))
	}
	m.pool = appendNonempty(m.pool, condEvalActions(extAction.Condition, extAction.Actions, context, workMem))
}

func (m *MuSteelExecuter) chooseActions() ([]SemanticAction, int) {
	// TODO: implement other strategies
	return m.pool[0], 0
}

func (m *MuSteelExecuter) execActions(actions []SemanticAction) {
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
			Xset = append(Xset, action)
			fmt.Print(action)
		}
	}
	fmt.Println()
	sActions, eActions := m.discovery(Xset)
	m.pool = append(m.pool, sActions...)
	m.extPool = append(m.extPool, eActions...)
}

func (m *MuSteelExecuter) removeActions(index int) {
	m.pool = append(m.pool[:index], m.pool[index+1:len(m.pool)]...)
}

func (m *MuSteelExecuter) discovery(Xset []SemanticAction) ([][]SemanticAction, []ExternalAction) {
	var newpool [][]SemanticAction
	var extActions []ExternalAction
	localRules, globalRules := m.activeRules(Xset)
	for _, rule := range localRules {
		if rule.DefaultActions != nil {
			newpool = append(newpool, evalActions(rule.DefaultActions, m.dataContext, m.workingMemory))
		}
		newpool = appendNonempty(newpool, condEvalActions(rule.Task.Exp, rule.Task.Actions, m.dataContext, m.workingMemory))
	}
	for _, rule := range globalRules {
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
	err := dataContext.Add(name, &(m.memory))
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
	return fmt.Sprintf("Memory: %v\nPool: %v\nExtPool: %v\n", m.memory, m.printPool(), m.printExtPool())
}

func (m *MuSteelExecuter) printPool() string {
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

func (m *MuSteelExecuter) printExtPool() string {
	if len(m.extPool) == 0 {
		return "{}"
	} else {
		str := "{"
		for _, action := range m.extPool {
			str = str + "\n  " + action.String()
		}
		return str + "\n}"
	}
}
