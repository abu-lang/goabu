package semantics

import (
	"errors"
	"fmt"
	"reflect"
	"steel-lang/datastructure"
	"strconv"

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
	parsedActions int
	library       map[string]*datastructure.RuleDict

	knowledgeLibrary *ast.KnowledgeLibrary
	workingMemory    *ast.WorkingMemory
	dataContext      ast.IDataContext
}

func NewMuSteelExecuter(mem datastructure.Resources) (*MuSteelExecuter, error) {
	res := &MuSteelExecuter{
		memory:           mem.Clone(),
		pool:             make([][]SemanticAction, 0),
		parsedActions:    0,
		library:          make(map[string]*datastructure.RuleDict),
		knowledgeLibrary: ast.NewKnowledgeLibrary(),
		dataContext:      ast.NewDataContext(),
	}
	if !res.memory.IsValid() {
		return nil, errors.New("invalid Resources argument")
	}
	res.types = res.memory.GetTypes()
	err := res.dataContext.Add("this", &(res.memory))
	if err != nil {
		return nil, err
	}
	return res, nil
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
	for _, evt := range parsed.Event {
		if m.library[evt] == nil {
			var dict datastructure.RuleDict = make(map[string]*datastructure.ParsedRule)
			m.library[evt] = &dict
		}
		m.library[evt].Insert(parsed)
	}
}

func (m *MuSteelExecuter) AddRules(rules []datastructure.Rule) {
	for _, rule := range rules {
		m.AddRule(&rule)
	}
}

func (m *MuSteelExecuter) AddActions(actions []datastructure.Action) {
	m.pool = append(m.pool, m.parseActions(actions))
}

func (m *MuSteelExecuter) AddPool(pl [][]datastructure.Action) {
	for _, actions := range pl {
		m.AddActions(actions)
	}
}

func (m *MuSteelExecuter) Exec() {
	if m.workingMemory == nil { // => m does not have rules nor actions
		return
	}
	actions, index := m.chooseActions()
	var Xset []SemanticAction
	str := ""
	for _, action := range actions {
		variable := action.Variable
		variable = m.workingMemory.AddVariable(variable)
		currentVal, err := variable.Evaluate(m.dataContext, m.workingMemory)
		if err != nil {
			panic(err)
		}
		diff := false
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
		if diff {
			err := variable.Assign(action.Value, m.dataContext, m.workingMemory)
			if err != nil {
				panic(err)
			}
			Xset = append(Xset, action)
			str = str + action.String()
		}
	}
	m.removeActions(index)
	m.pool = joinPool(m.pool, m.discovery(Xset))
	fmt.Println("Exec: " + str)
}

func (m *MuSteelExecuter) Input(actions []datastructure.Action) {
	sactions := m.parseActions(actions)
	str := ""
	for _, action := range sactions {
		variable := action.Variable
		variable = m.workingMemory.AddVariable(variable)
		currentVal, err := variable.Evaluate(m.dataContext, m.workingMemory)
		if err != nil {
			panic(err)
		}
		ltype := currentVal.Type()
		rtype := action.Value.Type()
		if !rtype.AssignableTo(ltype) {
			panic(fmt.Errorf("cannot assign a %v to a %v", rtype, ltype))
		}
		err = variable.Assign(action.Value, m.dataContext, m.workingMemory)
		if err != nil {
			panic(err)
		}
		str = str + action.String()
	}
	m.pool = joinPool(m.pool, m.discovery(sactions))
	fmt.Println("Input: " + str)
}

func (m *MuSteelExecuter) chooseActions() ([]SemanticAction, int) {
	// TODO: implement other strategies
	return m.pool[0], 0
}

func (m *MuSteelExecuter) removeActions(index int) {
	m.pool = append(m.pool[:index], m.pool[index+1:len(m.pool)]...)
}

func (m *MuSteelExecuter) discovery(Xset []SemanticAction) [][]SemanticAction {
	var newpool [][]SemanticAction
	rules := m.activeRules(Xset)
	if rules.Empty() {
		return newpool
	}
	for _, rule := range *rules {
		if rule.DefaultActions != nil {
			newpool = joinPool(newpool, [][]SemanticAction{m.discoveryActions(rule.DefaultActions)})
		}
		newpool = joinPool(newpool, m.discoveryTask(rule.Task))
	}
	return newpool
}

func (m *MuSteelExecuter) discoveryActions(acts []datastructure.ParsedAction) []SemanticAction {
	var sacts []SemanticAction
	for _, action := range acts {
		assignment := action.Expression
		variable := assignment.Variable
		rexpr := assignment.Expression
		rexpr = m.workingMemory.AddExpression(rexpr)
		exprVal, err := rexpr.Evaluate(m.dataContext, m.workingMemory)
		if err != nil {
			panic(err)
		}
		sacts = append(sacts, SemanticAction{
			Resource: action.Resource,
			Variable: variable,
			Value:    exprVal,
		})
	}
	return sacts
}

func (m *MuSteelExecuter) discoveryTask(task datastructure.ParsedTask) [][]SemanticAction {
	switch task.Mode {
	case "for":
		exp := task.Exp
		exp = m.workingMemory.AddExpression(exp)
		val, err := exp.Evaluate(m.dataContext, m.workingMemory)
		if err != nil {
			panic(err)
		}
		if val.Bool() {
			return [][]SemanticAction{m.discoveryActions(task.Actions)}
		}
	case "for some":
		// TODO: implement
	case "for all":
		// TODO: implement
	}
	return nil
}

func (m *MuSteelExecuter) activeRules(Xset []SemanticAction) *datastructure.RuleDict {
	var dict datastructure.RuleDict = make(map[string]*datastructure.ParsedRule)
	var res *datastructure.RuleDict = &dict
	for _, act := range Xset {
		res.Add(m.library[act.Resource])
	}
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
	m.workingMemory = knowledgeBase.WorkingMemory
}

func (m *MuSteelExecuter) PrintState() string {
	return fmt.Sprintf("Memory: %v\nPool: %v\n", m.memory, m.printPool())
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

//----------------------------SEMANTIC ACTION---------------------------------

type SemanticAction struct {
	Resource string
	Variable *ast.Variable
	Value    reflect.Value
}

func joinPool(pool1, pool2 [][]SemanticAction) [][]SemanticAction {
	return append(pool1, pool2...)
}

func (m *MuSteelExecuter) parseActions(actions []datastructure.Action) []SemanticAction {
	res := make([]SemanticAction, 0)
	for _, act := range actions {
		res = append(res, m.parseAction(act, "semaction"+strconv.Itoa(m.parsedActions)))
		m.parsedActions++
	}
	return res
}

func (m *MuSteelExecuter) parseAction(action datastructure.Action, name string) SemanticAction {
	parsed := datastructure.NewParsedAction(&action, name, m.knowledgeLibrary, m.types)
	m.updateWorkingMemory()
	assignment := parsed.Expression
	variable := assignment.Variable
	rexpr := assignment.Expression
	rexpr = m.workingMemory.AddExpression(rexpr)
	exprVal, err := rexpr.Evaluate(m.dataContext, m.workingMemory)
	if err != nil {
		panic(err)
	}
	return SemanticAction{
		Resource: parsed.Resource,
		Variable: variable,
		Value:    exprVal,
	}
}

func (action SemanticAction) String() string {
	return fmt.Sprintf("(%s,%v)", action.Resource, action.Value)
}
