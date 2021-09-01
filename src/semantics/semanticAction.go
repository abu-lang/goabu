package semantics

import (
	"fmt"
	"reflect"
	"steel-lang/ecarule"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"go.uber.org/zap/zapcore"
)

type SemanticAction struct {
	Resource string
	Variable *ast.Variable
	Value    reflect.Value
}

func (a SemanticAction) String() string {
	return fmt.Sprintf("(%s,%v)", a.Resource, a.Value)
}

func appendNonempty(pool [][]SemanticAction, actions []SemanticAction) [][]SemanticAction {
	if len(actions) == 0 {
		return pool
	}
	return append(pool, actions)
}

func evalActions(actions []ecarule.Action, dataContext ast.IDataContext, workingMemory *ast.WorkingMemory) []SemanticAction {
	res := make([]SemanticAction, 0)
	for _, action := range actions {
		assignment := action.Assignment
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

func condEvalActions(exp *ast.Expression, actions []ecarule.Action, dataContext ast.IDataContext, workingMemory *ast.WorkingMemory) []SemanticAction {
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

//----------------------------------LOGGER------------------------------------

type actionLogger SemanticAction

func (l actionLogger) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("typ", l.Variable.Variable.Name)
	enc.AddString("res", l.Resource)
	enc.AddString("val", fmt.Sprint(l.Value.Interface()))
	return nil
}

type updateLogger []SemanticAction

func (l updateLogger) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	for _, a := range l {
		err := enc.AppendObject(actionLogger(a))
		if err != nil {
			return err
		}
	}
	return nil
}

type poolLogger [][]SemanticAction

func (l poolLogger) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	for _, u := range l {
		err := enc.AppendArray(updateLogger(u))
		if err != nil {
			return err
		}
	}
	return nil
}
