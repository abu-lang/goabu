package semantics

import (
	"fmt"
	"reflect"
	"steel-lang/datastructure"
	"strconv"

	"github.com/hyperjumptech/grule-rule-engine/ast"
)

type SemanticAction struct {
	Resource string
	Variable *ast.Variable
	Value    reflect.Value
}

func (action SemanticAction) String() string {
	return fmt.Sprintf("(%s,%v)", action.Resource, action.Value)
}

func appendNonempty(pool [][]SemanticAction, actions []SemanticAction) [][]SemanticAction {
	if len(actions) == 0 {
		return pool
	}
	return append(pool, actions)
}

func (m *MuSteelExecuter) parseAction(action datastructure.Action, name string) datastructure.ParsedAction {
	parsed := datastructure.NewParsedAction(&action, name, m.knowledgeLibrary, m.types)
	m.updateWorkingMemory()
	return parsed
}

func (m *MuSteelExecuter) parseActions(actions []datastructure.Action) []datastructure.ParsedAction {
	res := make([]datastructure.ParsedAction, 0)
	for _, act := range actions {
		res = append(res, m.parseAction(act, "semaction"+strconv.Itoa(m.parsedActions)))
		m.parsedActions++
	}
	return res
}

func evalActions(actions []datastructure.ParsedAction, dataContext ast.IDataContext, workingMemory *ast.WorkingMemory) []SemanticAction {
	res := make([]SemanticAction, 0)
	for _, action := range actions {
		assignment := action.Expression
		variable := assignment.Variable
		rexpr := assignment.Expression
		rexpr = workingMemory.AddExpression(rexpr)
		exprVal, err := rexpr.Evaluate(dataContext, workingMemory)
		if err != nil {
			panic(err)
		}
		res = append(res, SemanticAction{
			Resource: action.Resource,
			Variable: variable,
			Value:    exprVal,
		})
	}
	return res
}

func condEvalActions(exp *ast.Expression, actions []datastructure.ParsedAction, dataContext ast.IDataContext, workingMemory *ast.WorkingMemory) []SemanticAction {
	exp = workingMemory.AddExpression(exp)
	val, err := exp.Evaluate(dataContext, workingMemory)
	if err != nil {
		panic(err)
	}
	if val.Bool() {
		return evalActions(actions, dataContext, workingMemory)
	}
	return nil
}