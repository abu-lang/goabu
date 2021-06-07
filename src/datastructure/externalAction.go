package datastructure

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

type ExternalAction struct {
	Condition      *ast.Expression
	Actions        []ParsedAction
	CondWorkingSet StringSet
	WorkingSets    []StringSet
	Constants      map[string]interface{}
	IntConstants   map[string]int64
	dataContext    ast.IDataContext
	workingMemory  *ast.WorkingMemory
}

func (a ExternalAction) String() string {
	return fmt.Sprintf("if %v do:\n  %v", a.Condition.GetGrlText(), ActionsToStr(a.Actions))
}

func (a ExternalAction) CullActions(localResources StringSet) []ParsedAction {
	var res []ParsedAction
	for i, action := range a.Actions {
		if localResources.ContainsSet(a.WorkingSets[i]) {
			res = append(res, action)
		}
	}
	return res
}

func (a ExternalAction) preEvaluatedActions(actions []ParsedAction) []ParsedAction {
	if actions == nil {
		return nil
	}
	res := make([]ParsedAction, 0, len(actions))
	for i, action := range actions {
		res = append(res, ParsedAction{
			Resource:   action.Resource,
			Expression: a.preEvaluatedAssignment(action.Expression, a.WorkingSets[i]),
		})
	}
	return res
}

func (a ExternalAction) preEvaluatedAssignment(assign *ast.Assignment, workingSet StringSet) *ast.Assignment {
	res := assign.Clone(pkg.NewCloneTable())
	a.partiallyEvalVariable(res.Variable, MakeStringSet(""), false)
	a.partiallyEvalExpression(res.Expression, workingSet, true)
	return res
}

func (a ExternalAction) preEvaluatedExpression(exp *ast.Expression, workingSet StringSet) *ast.Expression {
	res := exp.Clone(pkg.NewCloneTable())
	a.partiallyEvalExpression(res, workingSet, true)
	return res
}

func (a ExternalAction) partiallyEvalExpression(e *ast.Expression, workingSet StringSet, eval bool) {
	if e == nil {
		return
	}
	a.partiallyEvalExpression(e.LeftExpression, workingSet, eval)
	a.partiallyEvalExpression(e.RightExpression, workingSet, eval)
	a.partiallyEvalExpression(e.SingleExpression, workingSet, eval)
	a.partiallyEvalExpressionAtom(e.ExpressionAtom, workingSet, eval)
}

func (a ExternalAction) partiallyEvalExpressionAtom(e *ast.ExpressionAtom, workingSet StringSet, eval bool) {
	if e == nil {
		return
	}
	if e.Constant != nil {
		a.detach(e.Constant.GetAstID(), e.Constant.Value)
	}
	if e.FunctionCall != nil {
		a.partiallyEvalArgumentList(e.FunctionCall.ArgumentList, workingSet, eval)
	}
	a.partiallyEvalExpressionAtom(e.ExpressionAtom, workingSet, eval)
	if e.ArrayMapSelector != nil {
		a.partiallyEvalExpression(e.ArrayMapSelector.Expression, workingSet, eval)
	}
	if e.Variable == nil {
		return
	}
	if eval && strings.HasPrefix(e.Variable.GetGrlText(), "this.") {
		variable := a.workingMemory.AddVariable(e.Variable)
		val, err := variable.Evaluate(a.dataContext, a.workingMemory)
		if err != nil {
			panic(err)
		}
		e.Variable = nil
		constant := ast.NewConstant()
		if val.Kind() == reflect.String {
			constant.SetGrlText(fmt.Sprintf(`"%s"`, val.String()))
		}
		constant.Value = val
		e.Constant = constant
		a.detach(constant.GetAstID(), val)
	} else if eval && strings.HasPrefix(e.Variable.GetGrlText(), "ext.") {
		a.partiallyEvalVariable(e.Variable, workingSet, eval)
		switch {
		case e.Variable.ArrayMapSelector == nil:
			return
		case e.Variable.ArrayMapSelector.Expression == nil:
			return
		case e.Variable.ArrayMapSelector.Expression.ExpressionAtom == nil:
			return
		case e.Variable.ArrayMapSelector.Expression.ExpressionAtom.Constant == nil:
			return
		}
		text := e.Variable.ArrayMapSelector.Expression.ExpressionAtom.Constant.GetGrlText()
		res := strings.Split(text, `"`)[1]
		workingSet.Insert(res)
	} else {
		a.partiallyEvalVariable(e.Variable, workingSet, eval)
	}
}

func (a ExternalAction) detach(key string, val reflect.Value) {
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		a.IntConstants[key] = val.Int()
	default:
		a.Constants[key] = val.Interface()
	}
}

func (a ExternalAction) partiallyEvalArgumentList(e *ast.ArgumentList, workingSet StringSet, eval bool) {
	if e == nil {
		return
	}
	for _, arg := range e.Arguments {
		a.partiallyEvalExpression(arg, workingSet, eval)
	}
}

func (a ExternalAction) partiallyEvalVariable(e *ast.Variable, workingSet StringSet, eval bool) {
	if e == nil {
		return
	}
	if e.ArrayMapSelector != nil {
		a.partiallyEvalExpression(e.ArrayMapSelector.Expression, workingSet, eval)
	}
}

func (a ExternalAction) attachConstants() {
	a.attachConstantsExpression(a.Condition)
	a.attachConstantsActions(a.Actions)
}

func (a ExternalAction) attachConstantsActions(actions []ParsedAction) {
	for _, action := range actions {
		a.attachConstantsAssignment(action.Expression)
	}
}

func (a ExternalAction) attachConstantsAssignment(e *ast.Assignment) {
	a.attachConstantsVariable(e.Variable)
	a.attachConstantsExpression(e.Expression)
}

func (a ExternalAction) attachConstantsExpression(e *ast.Expression) {
	if e == nil {
		return
	}
	a.attachConstantsExpression(e.LeftExpression)
	a.attachConstantsExpression(e.RightExpression)
	a.attachConstantsExpression(e.SingleExpression)
	a.attachConstantsExpressionAtom(e.ExpressionAtom)
}

func (a ExternalAction) attachConstantsExpressionAtom(e *ast.ExpressionAtom) {
	if e == nil {
		return
	}
	if e.Constant != nil {
		val, present := a.Constants[e.Constant.GetAstID()]
		if present {
			e.Constant.Value = reflect.ValueOf(val)
		}
		integer, present := a.IntConstants[e.Constant.GetAstID()]
		if present {
			e.Constant.Value = reflect.ValueOf(integer)
		}
	}
	if e.FunctionCall != nil {
		a.attachConstantsArgumentList(e.FunctionCall.ArgumentList)
	}
	a.attachConstantsExpressionAtom(e.ExpressionAtom)
	if e.ArrayMapSelector != nil {
		a.attachConstantsExpression(e.ArrayMapSelector.Expression)
	}
	a.attachConstantsVariable(e.Variable)
}

func (a ExternalAction) attachConstantsArgumentList(e *ast.ArgumentList) {
	if e == nil {
		return
	}
	for _, arg := range e.Arguments {
		a.attachConstantsExpression(arg)
	}
}

func (a ExternalAction) attachConstantsVariable(e *ast.Variable) {
	if e == nil {
		return
	}
	if e.ArrayMapSelector != nil {
		a.attachConstantsExpression(e.ArrayMapSelector.Expression)
	}
}

func MarshalExternalActions(actions []ExternalAction) ([]byte, error) {
	return json.Marshal(actions)
}

func UnmarshalExternalActions(b []byte) ([]ExternalAction, error) {
	var eActions []ExternalAction
	err := json.Unmarshal(b, &eActions)
	if err != nil {
		return nil, err
	}
	for _, action := range eActions {
		action.attachConstants()
	}
	return eActions, nil
}
