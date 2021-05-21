package semantics

import (
	"fmt"
	"reflect"
	"steel-lang/datastructure"
	"strings"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

type ExternalAction struct {
	Condition     *ast.Expression
	Actions       []datastructure.ParsedAction
	WorkingSet    datastructure.StringSet
	WriteSet      datastructure.StringSet
	Constants     map[string]interface{}
	dataContext   ast.IDataContext
	workingMemory *ast.WorkingMemory
}

func (action ExternalAction) String() string {
	return fmt.Sprintf("if %v do:\n  %v", action.Condition.GetGrlText(), datastructure.ActionsToStr(action.Actions))
}

// Precondition: rule.Task.Mode != "for"
func (m *MuSteelExecuter) preEvaluated(rule *datastructure.ParsedRule) ExternalAction {
	res := ExternalAction{
		WorkingSet:    datastructure.MakeStringSet(""),
		WriteSet:      datastructure.MakeStringSet(""),
		Constants:     make(map[string]interface{}),
		dataContext:   m.dataContext,
		workingMemory: m.workingMemory,
	}
	res.Condition = res.preEvaluatedExpression(rule.Task.Condition)
	res.Actions = res.preEvaluatedActions(rule.Task.Actions)
	return res
}

func (a ExternalAction) preEvaluatedActions(actions []datastructure.ParsedAction) []datastructure.ParsedAction {
	if actions == nil {
		return nil
	}
	res := make([]datastructure.ParsedAction, 0, len(actions))
	for _, action := range actions {
		res = append(res, datastructure.ParsedAction{
			Resource:   action.Resource,
			Expression: a.preEvaluatedAssignment(action.Expression),
		})
		a.WorkingSet.Insert(action.Resource)
		a.WriteSet.Insert(action.Resource)
	}
	return res
}

func (a ExternalAction) preEvaluatedAssignment(assign *ast.Assignment) *ast.Assignment {
	res := assign.Clone(pkg.NewCloneTable())
	a.partiallyEvalExpression(res.Expression)
	return res
}

func (a ExternalAction) preEvaluatedExpression(exp *ast.Expression) *ast.Expression {
	res := exp.Clone(pkg.NewCloneTable())
	a.partiallyEvalExpression(res)
	return res
}

func (a ExternalAction) partiallyEvalExpression(e *ast.Expression) {
	if e == nil {
		return
	}
	a.partiallyEvalExpression(e.LeftExpression)
	a.partiallyEvalExpression(e.RightExpression)
	a.partiallyEvalExpression(e.SingleExpression)
	a.partiallyEvalExpressionAtom(e.ExpressionAtom)
}

func (a ExternalAction) partiallyEvalExpressionAtom(e *ast.ExpressionAtom) {
	if e == nil {
		return
	}
	if e.Constant != nil {
		a.Constants[e.Constant.GetAstID()] = e.Constant.Value.Interface()
	}
	if e.FunctionCall != nil {
		a.partiallyEvalArgumentList(e.FunctionCall.ArgumentList)
	}
	a.partiallyEvalExpressionAtom(e.ExpressionAtom)
	if e.ArrayMapSelector != nil {
		a.partiallyEvalExpression(e.ArrayMapSelector.Expression)
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
		a.partiallyEvalVariable(e.Variable)
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
		a.WorkingSet.Insert(res)
	} else {
		a.partiallyEvalVariable(e.Variable)
	}
}

func (a ExternalAction) partiallyEvalArgumentList(e *ast.ArgumentList) {
	if e == nil {
		return
	}
	for _, arg := range e.Arguments {
		a.partiallyEvalExpression(arg)
	}
}

func (a ExternalAction) partiallyEvalVariable(e *ast.Variable) {
	if e == nil {
		return
	}
	if e.ArrayMapSelector != nil {
		a.partiallyEvalExpression(e.ArrayMapSelector.Expression)
	}
}

func (a ExternalAction) attachConstants() {
	a.attachConstantsExpression(a.Condition)
	a.attachConstantsActions(a.Actions)
}

func (a ExternalAction) attachConstantsActions(actions []datastructure.ParsedAction) {
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
