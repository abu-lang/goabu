package datastructure

import (
	"errors"

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
