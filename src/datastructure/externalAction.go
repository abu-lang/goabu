package datastructure

import (
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
	a.partiallyEvalExpression(res.Expression, workingSet)
	return res
}

func (a ExternalAction) preEvaluatedExpression(exp *ast.Expression, workingSet StringSet) *ast.Expression {
	res := exp.Clone(pkg.NewCloneTable())
	a.partiallyEvalExpression(res, workingSet)
	return res
}

func (a ExternalAction) partiallyEvalExpression(e *ast.Expression, workingSet StringSet) {
	if e == nil {
		return
	}
	a.partiallyEvalExpression(e.LeftExpression, workingSet)
	a.partiallyEvalExpression(e.RightExpression, workingSet)
	a.partiallyEvalExpression(e.SingleExpression, workingSet)
	a.partiallyEvalExpressionAtom(e.ExpressionAtom, workingSet)
}

func (a ExternalAction) partiallyEvalExpressionAtom(e *ast.ExpressionAtom, workingSet StringSet) {
	if e == nil {
		return
	}
	if e.Constant != nil {
		a.Constants[e.Constant.GetAstID()] = e.Constant.Value.Interface()
	}
	if e.FunctionCall != nil {
		a.partiallyEvalArgumentList(e.FunctionCall.ArgumentList, workingSet)
	}
	a.partiallyEvalExpressionAtom(e.ExpressionAtom, workingSet)
	if e.ArrayMapSelector != nil {
		a.partiallyEvalExpression(e.ArrayMapSelector.Expression, workingSet)
	}
	if e.Variable == nil {
		return
	}
	if strings.HasPrefix(e.Variable.GetGrlText(), "this.") {
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
		a.Constants[constant.GetAstID()] = val.Interface()
	} else if strings.HasPrefix(e.Variable.GetGrlText(), "ext.") {
		a.partiallyEvalVariable(e.Variable, workingSet)
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
		a.partiallyEvalVariable(e.Variable, workingSet)
	}
}

func (a ExternalAction) partiallyEvalArgumentList(e *ast.ArgumentList, workingSet StringSet) {
	if e == nil {
		return
	}
	for _, arg := range e.Arguments {
		a.partiallyEvalExpression(arg, workingSet)
	}
}

func (a ExternalAction) partiallyEvalVariable(e *ast.Variable, workingSet StringSet) {
	if e == nil {
		return
	}
	if e.ArrayMapSelector != nil {
		a.partiallyEvalExpression(e.ArrayMapSelector.Expression, workingSet)
	}
}

func (a ExternalAction) AttachConstants() {
	a.attachConstantsExpression(a.Condition)
	a.attachConstantsActions(a.Actions)
}

func (a ExternalAction) attachConstantsActions(actions []ParsedAction) {
	for _, action := range actions {
		a.attachConstantsExpression(action.Expression.Expression)
	}
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
