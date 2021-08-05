package datastructure

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hyperjumptech/grule-rule-engine/ast"
)

type Rule struct {
	Name           string
	Events         []string
	DefaultActions []Action
	Task           Task
}

type Action struct {
	Resource   string
	Expression *ast.Assignment
}

type Task struct {
	Mode      string
	Condition *ast.Expression
	Actions   []Action
}

// AcceptAssignment will accept an Assignment into this Rule by creating a new corresponding DefaultAction.
// It is used by EcaruleParserListener.
// One who wants to use this function in a different context should make sure that the Assignment
// satisfies the constraints implied by this Rule and by the Rule owner.
func (r *Rule) AcceptAssignment(a *ast.Assignment) error {
	if !a.IsAssign {
		return fmt.Errorf("assigment %s only assigment operator '=' is supported", a.GetGrlText())
	}
	n := strings.Trim(a.Variable.ArrayMapSelector.Expression.ExpressionAtom.Constant.GetGrlText(), `"`)
	r.DefaultActions = append(r.DefaultActions, Action{Resource: n, Expression: a})
	return nil
}

// AcceptAssignment will accept an Assignment into this Task by creating a new corresponding Action.
// It is used by EcaruleParserListener.
// One who wants to use this function in a different context should make sure that the Assignment
// satisfies the constraints implied by this Task and by the Rule owner.
func (t *Task) AcceptAssignment(a *ast.Assignment) error {
	if !a.IsAssign {
		return fmt.Errorf("assigment %s only assigment operator '=' is supported", a.GetGrlText())
	}
	n := strings.Trim(a.Variable.ArrayMapSelector.Expression.ExpressionAtom.Constant.GetGrlText(), `"`)
	t.Actions = append(t.Actions, Action{Resource: n, Expression: a})
	return nil
}

// AcceptExpression will accept an Expression into the Condition of this Task.
// It is used by EcaruleParserListener.
// One who wants to use this function in a different context should make sure that the Expression
// satisfies the constraints implied by this Task and by the Rule owner.
func (t *Task) AcceptExpression(exp *ast.Expression) error {
	if t.Condition != nil {
		return errors.New("task condition already assigned")
	}
	t.Condition = exp
	return nil
}

func (a Action) String() string {
	return a.Expression.GetGrlText()
}

func ActionsToStr(actions []Action) string {
	res := ""
	for _, action := range actions {
		res += action.String() + "; "
	}
	return res
}
