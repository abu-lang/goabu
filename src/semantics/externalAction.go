package semantics

import (
	"fmt"
	"steel-lang/datastructure"
	"strings"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

type ExternalAction struct {
	DefaultActions []datastructure.ParsedAction
	Condition      *ast.Expression
	Actions        []datastructure.ParsedAction
}

func (action ExternalAction) String() string {
	tail := fmt.Sprintf("if %v do:\n  %v", action.Condition.GetGrlText(), datastructure.ActionsToStr(action.Actions))
	if len(action.DefaultActions) > 0 {
		return fmt.Sprintf("default %v\n  ", datastructure.ActionsToStr(action.DefaultActions)) + tail
	}
	return tail
}

// Precondition: rule.Task.Mode != "for"
func (m *MuSteelExecuter) preEvaluated(rule *datastructure.ParsedRule) ExternalAction {
	return ExternalAction{
		DefaultActions: m.preEvaluatedActions(rule.DefaultActions),
		Condition:      m.preEvaluatedExpression(rule.Task.Exp),
		Actions:        m.preEvaluatedActions(rule.Task.Actions),
	}
}

func (m *MuSteelExecuter) preEvaluatedActions(actions []datastructure.ParsedAction) []datastructure.ParsedAction {
	if actions == nil {
		return nil
	}
	res := make([]datastructure.ParsedAction, 0, len(actions))
	for _, action := range actions {
		res = append(res, datastructure.ParsedAction{
			Resource:   action.Resource,
			Expression: m.preEvaluatedAssignment(action.Expression),
		})
	}
	return res
}

func (m *MuSteelExecuter) preEvaluatedAssignment(assign *ast.Assignment) *ast.Assignment {
	res := assign.Clone(pkg.NewCloneTable())
	m.partiallyEvalExpression(res.Expression)
	return res
}

func (m *MuSteelExecuter) preEvaluatedExpression(exp *ast.Expression) *ast.Expression {
	res := exp.Clone(pkg.NewCloneTable())
	m.partiallyEvalExpression(res)
	return res
}

func (m *MuSteelExecuter) partiallyEvalExpression(e *ast.Expression) {
	if e == nil {
		return
	}
	m.partiallyEvalExpression(e.LeftExpression)
	m.partiallyEvalExpression(e.RightExpression)
	m.partiallyEvalExpression(e.SingleExpression)
	m.partiallyEvalExpressionAtom(e.ExpressionAtom)
}

func (m *MuSteelExecuter) partiallyEvalExpressionAtom(e *ast.ExpressionAtom) {
	if e == nil {
		return
	}
	if e.FunctionCall != nil {
		m.partiallyEvalArgumentList(e.FunctionCall.ArgumentList)
	}
	m.partiallyEvalExpressionAtom(e.ExpressionAtom)
	if e.ArrayMapSelector != nil {
		m.partiallyEvalExpression(e.ArrayMapSelector.Expression)
	}
	if e.Variable == nil {
		return
	}
	if strings.HasPrefix(e.Variable.GetGrlText(), "this.") {
		variable := m.workingMemory.AddVariable(e.Variable)
		val, err := variable.Evaluate(m.dataContext, m.workingMemory)
		if err != nil {
			panic(err)
		}
		e.Variable = nil
		constant := ast.NewConstant()
		constant.Value = val
		e.Constant = constant
	} else {
		m.partiallyEvalVariable(e.Variable)
	}
}

func (m *MuSteelExecuter) partiallyEvalArgumentList(e *ast.ArgumentList) {
	if e == nil {
		return
	}
	for _, arg := range e.Arguments {
		m.partiallyEvalExpression(arg)
	}
}

func (m *MuSteelExecuter) partiallyEvalVariable(e *ast.Variable) {
	if e == nil {
		return
	}
	if e.ArrayMapSelector != nil {
		m.partiallyEvalExpression(e.ArrayMapSelector.Expression)
	}
}
