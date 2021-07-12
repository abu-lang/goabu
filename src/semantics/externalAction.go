package semantics

import (
	"encoding/json"
	"fmt"
	"reflect"
	"steel-lang/datastructure"
	"steel-lang/misc"
	"strings"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

type externalAction struct {
	Condition      *ast.Expression
	Actions        []datastructure.Action
	CondWorkingSet misc.StringSet
	WorkingSets    []misc.StringSet
	Constants      map[string]interface{}
	IntConstants   map[string]int64
	dataContext    ast.IDataContext
	workingMemory  *ast.WorkingMemory
}

func (a externalAction) String() string {
	return fmt.Sprintf("if %v do:\n  %v", a.Condition.GetGrlText(), datastructure.ActionsToStr(a.Actions))
}

func (a externalAction) cullActions(localResources misc.StringSet) []datastructure.Action {
	var res []datastructure.Action
	for i, action := range a.Actions {
		if localResources.ContainsSet(a.WorkingSets[i]) {
			res = append(res, action)
		}
	}
	return res
}

func (a externalAction) preEvaluatedActions(actions []datastructure.Action) []datastructure.Action {
	if actions == nil {
		return nil
	}
	res := make([]datastructure.Action, 0, len(actions))
	for i, action := range actions {
		res = append(res, datastructure.Action{
			Resource:   action.Resource,
			Expression: a.preEvaluatedAssignment(action.Expression, a.WorkingSets[i]),
		})
	}
	return res
}

func (a externalAction) preEvaluatedAssignment(assign *ast.Assignment, workingSet misc.StringSet) *ast.Assignment {
	res := assign.Clone(pkg.NewCloneTable())
	a.partiallyEvalVariable(res.Variable, misc.MakeStringSet(""), false)
	a.partiallyEvalExpression(res.Expression, workingSet, true)
	return res
}

func (a externalAction) preEvaluatedExpression(exp *ast.Expression, workingSet misc.StringSet) *ast.Expression {
	res := exp.Clone(pkg.NewCloneTable())
	a.partiallyEvalExpression(res, workingSet, true)
	return res
}

func (a externalAction) partiallyEvalExpression(e *ast.Expression, workingSet misc.StringSet, eval bool) {
	if e == nil {
		return
	}
	a.partiallyEvalExpression(e.LeftExpression, workingSet, eval)
	a.partiallyEvalExpression(e.RightExpression, workingSet, eval)
	a.partiallyEvalExpression(e.SingleExpression, workingSet, eval)
	a.partiallyEvalExpressionAtom(e.ExpressionAtom, workingSet, eval)
}

func (a externalAction) partiallyEvalExpressionAtom(e *ast.ExpressionAtom, workingSet misc.StringSet, eval bool) {
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

func (a externalAction) detach(key string, val reflect.Value) {
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		a.IntConstants[key] = val.Int()
	default:
		a.Constants[key] = val.Interface()
	}
}

func (a externalAction) partiallyEvalArgumentList(e *ast.ArgumentList, workingSet misc.StringSet, eval bool) {
	if e == nil {
		return
	}
	for _, arg := range e.Arguments {
		a.partiallyEvalExpression(arg, workingSet, eval)
	}
}

func (a externalAction) partiallyEvalVariable(e *ast.Variable, workingSet misc.StringSet, eval bool) {
	if e == nil {
		return
	}
	if e.ArrayMapSelector != nil {
		a.partiallyEvalExpression(e.ArrayMapSelector.Expression, workingSet, eval)
	}
}

func (a externalAction) attachTypesConsts(types map[string]string) {
	a.attachTypesConstsExpression(a.Condition, types)
	a.attachTypesConstsActions(a.Actions, types)
}

func (a externalAction) attachTypesConstsActions(actions []datastructure.Action, types map[string]string) {
	for _, action := range actions {
		a.attachTypesConstsAssignment(action.Expression, types)
	}
}

func (a externalAction) attachTypesConstsAssignment(e *ast.Assignment, types map[string]string) {
	a.attachTypesConstsVariable(e.Variable, types)
	a.attachTypesConstsExpression(e.Expression, types)
	switch {
	case e.Variable == nil:
		return
	case e.Variable.Variable == nil:
		return
	case e.Variable.Variable.Variable == nil:
		return
	}
	if e.Variable.Variable.Variable.Name == "ext" {
		e.Variable.Variable.Variable.Name = "this"
	}
}

func (a externalAction) attachTypesConstsExpression(e *ast.Expression, types map[string]string) {
	if e == nil {
		return
	}
	a.attachTypesConstsExpression(e.LeftExpression, types)
	a.attachTypesConstsExpression(e.RightExpression, types)
	a.attachTypesConstsExpression(e.SingleExpression, types)
	a.attachTypesConstsExpressionAtom(e.ExpressionAtom, types)
}

func (a externalAction) attachTypesConstsExpressionAtom(e *ast.ExpressionAtom, types map[string]string) {
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
		a.attachTypesConstsArgumentList(e.FunctionCall.ArgumentList, types)
	}
	a.attachTypesConstsExpressionAtom(e.ExpressionAtom, types)
	if e.ArrayMapSelector != nil {
		a.attachTypesConstsExpression(e.ArrayMapSelector.Expression, types)
	}
	a.attachTypesConstsVariable(e.Variable, types)
}

func (a externalAction) attachTypesConstsArgumentList(e *ast.ArgumentList, types map[string]string) {
	if e == nil {
		return
	}
	for _, arg := range e.Arguments {
		a.attachTypesConstsExpression(arg, types)
	}
}

func (a externalAction) attachTypesConstsVariable(e *ast.Variable, types map[string]string) {
	if e == nil {
		return
	}
	if e.ArrayMapSelector != nil {
		a.attachTypesConstsExpression(e.ArrayMapSelector.Expression, types)
		if strings.HasPrefix(e.GetGrlText(), "ext.") {
			switch {
			case e.ArrayMapSelector.Expression == nil:
				return
			case e.ArrayMapSelector.Expression.ExpressionAtom == nil:
				return
			case e.ArrayMapSelector.Expression.ExpressionAtom.Constant == nil:
				return
			}
			text := e.ArrayMapSelector.Expression.ExpressionAtom.Constant.GetGrlText()
			res := strings.Split(text, `"`)[1]
			e.Variable.Name = types[res]
		}
	}
}

func marshalExternalActions(actions []externalAction) ([]byte, error) {
	return json.Marshal(actions)
}

func unmarshalExternalActions(b []byte, types map[string]string) ([]externalAction, error) {
	var eActions []externalAction
	err := json.Unmarshal(b, &eActions)
	if err != nil {
		return nil, err
	}
	for _, action := range eActions {
		action.attachTypesConsts(types)
	}
	return eActions, nil
}
