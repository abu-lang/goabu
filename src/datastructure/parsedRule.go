package datastructure

import (
	"errors"

	"github.com/hyperjumptech/grule-rule-engine/ast"
)

type ParsedRule struct {
	Name           string
	Events         []string
	DefaultActions []ParsedAction
	Task           ParsedTask
}

type ParsedAction struct {
	Resource   string
	Expression *ast.Assignment
}

type ParsedTask struct {
	Mode      string
	Condition *ast.Expression
	Actions   []ParsedAction
}

func (t *ParsedTask) AcceptExpression(exp *ast.Expression) error {
	if t.Condition != nil {
		return errors.New("task condition already assigned")
	}
	t.Condition = exp
	return nil
}

func (a ParsedAction) String() string {
	return a.Expression.GetGrlText()
}

func ActionsToStr(actions []ParsedAction) string {
	res := ""
	for _, action := range actions {
		res += action.String() + "; "
	}
	return res
}
