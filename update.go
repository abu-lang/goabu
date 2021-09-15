package goabu

import (
	"fmt"
	"reflect"

	"github.com/abu-lang/goabu/ecarule"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"go.uber.org/zap/zapcore"
)

type Update []Assignment

type Assignment struct {
	Resource string
	variable *ast.Variable
	Value    reflect.Value
}

func (a Assignment) String() string {
	return fmt.Sprintf("(%s,%v)", a.Resource, a.Value)
}

func appendNonempty(pool []Update, u Update) []Update {
	if len(u) == 0 {
		return pool
	}
	return append(pool, u)
}

func evalActions(actions []ecarule.Action, dataContext ast.IDataContext, workingMemory *ast.WorkingMemory) (Update, error) {
	res := make([]Assignment, 0)
	for _, action := range actions {
		assignment := action.Assignment
		variable := assignment.Variable
		rexpr := assignment.Expression
		rexpr = workingMemory.AddExpression(rexpr)
		exprVal, err := rexpr.Evaluate(dataContext, workingMemory)
		if err != nil {
			return nil, err
		}
		res = append(res, Assignment{
			Resource: action.Resource,
			variable: variable,
			Value:    exprVal,
		})
	}
	return res, nil
}

func condEvalActions(exp *ast.Expression, actions []ecarule.Action, dataContext ast.IDataContext, workingMemory *ast.WorkingMemory) (Update, error) {
	exp = workingMemory.AddExpression(exp)
	val, err := exp.Evaluate(dataContext, workingMemory)
	if err != nil {
		return nil, err
	}
	if val.Bool() {
		return evalActions(actions, dataContext, workingMemory)
	}
	return nil, nil
}

//----------------------------------LOGGER------------------------------------

type assignmentLogger Assignment

func (l assignmentLogger) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("typ", l.variable.Variable.Name)
	enc.AddString("res", l.Resource)
	enc.AddString("val", fmt.Sprint(l.Value.Interface()))
	return nil
}

type updateLogger Update

func (l updateLogger) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	for _, a := range l {
		err := enc.AppendObject(assignmentLogger(a))
		if err != nil {
			return err
		}
	}
	return nil
}

type poolLogger []Update

func (l poolLogger) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	for _, u := range l {
		err := enc.AppendArray(updateLogger(u))
		if err != nil {
			return err
		}
	}
	return nil
}
