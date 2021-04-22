package semantics

import (
	"fmt"
	"steel-lang/antlr/exprParser"
	"steel-lang/datastructure"
	"steel-lang/misc"

	"github.com/hyperjumptech/grule-rule-engine/ast"
)

type State struct {
	Memory map[string]interface{}
	Pool   [][]SemanticAction
}

type stateMemory struct {
	Resources map[string]interface{}
}

type SemanticAction struct {
	Resource string
	Value    interface{}
}

func (action SemanticAction) String() string {
	return fmt.Sprintf("(%s,%v)", action.Resource, action.Value)
}

func joinPool(pool1, pool2 [][]SemanticAction) [][]SemanticAction {
	return append(pool1, pool2...)
}

type MuSteelExecuter struct {
	memory  stateMemory
	pool    [][]SemanticAction
	rules   []datastructure.Rule
	library map[string]*datastructure.RuleDict

	knowledgeLibrary *ast.KnowledgeLibrary
	dataContext      ast.IDataContext
}

func NewMuSteelExecuter(mem map[string]interface{}, pl [][]SemanticAction, rulesp []datastructure.Rule) *MuSteelExecuter {
	res := &MuSteelExecuter{
		memory: stateMemory{
			Resources: misc.CopyMap(mem),
		},
		pool:             pl,
		rules:            rulesp,
		library:          make(map[string]*datastructure.RuleDict),
		knowledgeLibrary: ast.NewKnowledgeLibrary(),
		dataContext:      ast.NewDataContext(),
	}
	err := res.dataContext.Add("this", &(res.memory))
	if err != nil {
		panic(err)
	}
	return res
}

func EvalCondition(cnd string, mem map[string]interface{}) bool {
	expintp := exprParser.NewexpIntp(cnd, mem)
	v := expintp.RunexpIntp()
	return v.(bool)
}

func EvalExpression(expr string, mem map[string]interface{}) interface{} {
	expintp := exprParser.NewexpIntp(expr, mem)
	return expintp.RunexpIntp()
}

func (m *MuSteelExecuter) GetState() State {
	memCopy := misc.CopyMap(m.memory.Resources)
	poolCopy := make([][]SemanticAction, len(m.pool))
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

func (m *MuSteelExecuter) Exec() {
	list, index := m.chooseActslist()
	var Xset []SemanticAction
	str := ""
	for _, act := range list {
		if m.memory.Resources[act.Resource] != act.Value {
			m.memory.Resources[act.Resource] = act.Value
			Xset = append(Xset, act)
			str = str + act.String()
		}
	}
	m.removeActslist(index)
	m.pool = joinPool(m.pool, m.discovery(Xset))
	fmt.Println("Exec: " + str)
}

func (m *MuSteelExecuter) Input(Xset []SemanticAction) {
	str := ""
	for _, sact := range Xset {
		m.memory.Resources[sact.Resource] = sact.Value
		str = str + sact.String()
	}
	m.pool = joinPool(m.pool, m.discovery(Xset))
	fmt.Println("Input: " + str)
}

func (m *MuSteelExecuter) chooseActslist() ([]SemanticAction, int) {
	// TODO: implement other strategies
	return m.pool[0], 0
}

func (m *MuSteelExecuter) removeActslist(index int) {
	m.pool = append(m.pool[:index], m.pool[index+1:len(m.pool)]...)
}

func (m *MuSteelExecuter) discovery(Xset []SemanticAction) [][]SemanticAction {
	var newpool [][]SemanticAction
	for _, rule := range m.rules {
		evt := rule.Event
		var triggered = false
		for _, act := range Xset {
			for _, id := range evt {
				if id == act.Resource {
					triggered = true
					break
				}
			}
		}
		if triggered {
			if rule.DefaultActions != nil {
				newpool = joinPool(newpool, [][]SemanticAction{m.discoveryActslist(rule.DefaultActions)})
			}
			newpool = joinPool(newpool, m.discoveryTask(rule.Task))
		}
	}
	return newpool
}

func (m *MuSteelExecuter) discoveryActslist(actslist []datastructure.Action) []SemanticAction {
	var sacts []SemanticAction
	var val interface{}
	for _, act := range actslist {
		expintp := exprParser.NewexpIntp(act.Expression, m.memory.Resources)
		val = expintp.RunexpIntp()
		sacts = append(sacts, SemanticAction{Resource: act.Resource, Value: val})
	}
	return sacts
}

func (m *MuSteelExecuter) discoveryTask(task datastructure.Task) [][]SemanticAction {
	switch task.Mode {
	case "for":
		cndintp := exprParser.NewexpIntp(task.Exp, m.memory.Resources)
		condition := cndintp.PevalExp()
		if EvalCondition(condition, m.memory.Resources) {
			var sacts []SemanticAction
			for _, act := range task.Actions {
				val := EvalExpression(act.Expression, m.memory.Resources)
				sacts = append(sacts, SemanticAction{Resource: act.Resource, Value: val})
			}
			return [][]SemanticAction{sacts}
		}
	case "for some":
		// TODO: implement
	case "for all":
		// TODO: implement
	}
	return nil
}

func (m *MuSteelExecuter) PrintState() string {
	return fmt.Sprintf("Memory: %v\nPool: %v\n", m.printMemory(), m.printPool())
}

func (m *MuSteelExecuter) printMemory() string {
	var str string = "[ "
	for key, value := range m.memory.Resources {
		str = str + fmt.Sprintf("(%T)%s->%v ", value, key, value)
	}
	return str + "]"
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
