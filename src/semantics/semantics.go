package semantics

import (
	"fmt"
	"steel-lang/antlr/exprParser"
	"steel-lang/datastructure"
	"steel-lang/misc"
)

type SemanticAction struct {
	Resource string
	Value interface{}
}

func printSemanticAction(action SemanticAction) string {
	return fmt.Sprintf("(%s,%v)", action.Resource, action.Value)
}

type State struct {
	Memory map[string]interface{}
	Pool   [][]SemanticAction
}

func PrintState(state State) string {
	return fmt.Sprintf("Memory: %v\nPool: %v\n", PrintMemory(state.Memory), PrintPool(state.Pool))
}

func PrintMemory(memory map[string]interface{}) string {
	var str string = "[ "
	for key , value := range memory {
		str = str + fmt.Sprintf("(%T)%s->%v ", value, key, value)
	}
	return str + "]"
}

func PrintPool(pool [][]SemanticAction) string {
	if len(pool) == 0 {
		return "{}"
	} else {
		str := "{"
		for _, list := range pool {
			str = str + "\n  "
			for _, action := range list {
				str = str + printSemanticAction(action)
			}
		}
		return str + "\n}"
	}
}

func EvalCondition(cnd string, mem map[string]interface{}) bool {
	expintp := exprParser.NewexpIntp(cnd,mem)
	v := expintp.RunexpIntp()
	return v.(bool)
}

func EvalExpression(expr string, mem map[string]interface{}) interface{} {
	expintp := exprParser.NewexpIntp(expr,mem)
	return expintp.RunexpIntp()
}

type mSteelExecuter struct {
	state *State
	rules []datastructure.Rule
}

func NewmSteelExecuter(statep *State, rulesp []datastructure.Rule) mSteelExecuter {
	return mSteelExecuter{state: statep, rules: rulesp}
}

func (m mSteelExecuter) GetState() State {
	mem := misc.CopyMap(m.state.Memory)
	var pool [][]SemanticAction
	for _, a := range m.state.Pool {
		var list []SemanticAction
		for _, e := range a {
			list = append(list,e)
		}
		pool = append(pool,list)
	}
	return State{Memory: mem, Pool: pool}
}

func (m mSteelExecuter) IsStable() bool {
	return len(m.state.Pool) == 0
}

func (m mSteelExecuter) Exec() {
	list, index := m.chooseActslist()
	var Xset []SemanticAction
	str := ""
	for _, act := range list {
		if m.state.Memory[act.Resource] != act.Value {
			m.state.Memory[act.Resource] = act.Value
			Xset = append(Xset,act)
			str = str + printSemanticAction(act)
		}
	}
	m.removeActslist(index)
	m.state.Pool = joinPool(m.state.Pool,m.discovery(Xset))
	fmt.Println("Exec: " + str)
}

func (m mSteelExecuter) chooseActslist() ([]SemanticAction, int) {
	// TODO: implement other strategies
	return m.state.Pool[0], 0
}

func (m mSteelExecuter) removeActslist(index int) {
	m.state.Pool = append(m.state.Pool[:index], m.state.Pool[index+1:len(m.state.Pool)] ...)
}

func joinPool(pool1, pool2 [][]SemanticAction) [][]SemanticAction {
	return append(pool1, pool2 ...)
}

func (m mSteelExecuter) discovery(Xset []SemanticAction) [][]SemanticAction {
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

func (m mSteelExecuter) discoveryActslist(actslist []datastructure.Action) []SemanticAction {
	var sacts []SemanticAction
	var val interface{}
	for _, act := range actslist {
		expintp := exprParser.NewexpIntp(act.Expression, m.state.Memory)
		val = expintp.RunexpIntp()
		sacts = append(sacts,SemanticAction{Resource: act.Resource, Value: val})
	}
	return sacts
}

func (m mSteelExecuter) discoveryTask(task datastructure.Task) [][]SemanticAction {
	switch task.Mode {
	case "for":
		cndintp := exprParser.NewexpIntp(task.Exp,m.state.Memory)
		condition := cndintp.PevalExp()
		if EvalCondition(condition,m.state.Memory) {
			var sacts []SemanticAction
			for _, act := range task.Actions {
				val := EvalExpression(act.Expression,m.state.Memory)
				sacts = append(sacts, SemanticAction{Resource: act.Resource, Value: val})
			}
			return [][]SemanticAction{sacts}
		}
		break
	case "for some":
		// TODO: implement
		break
	case "for all":
		// TODO: implement
		break
	}
	return nil
}

