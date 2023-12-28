// Copyright 2023 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/abu-lang/goabu/ecarule"
	"github.com/hyperjumptech/grule-rule-engine/ast"
)

// tmpExpressionReceiver is a struct wrapping an [*ecarule.LocalTask] for allowing tasks to implement the
// receiver interfaces required by Grule rule engine's parser (without exporting the methods on
// the [ecarule.RemoteTask] type).
type tmpExpressionReceiver struct {
	*ecarule.RemoteTask
}

// AcceptAssignment will accept an Assignment into this remote task by creating a new corresponding Action.
// It is used by EcaruleParserListener.
// One who wants to use this function in a different context should make sure that the Assignment
// satisfies the constraints implied by this remote task and by the Rule owner.
func (t tmpExpressionReceiver) AcceptAssignment(a *ast.Assignment) error {
	if !a.IsAssign {
		return fmt.Errorf("assigment %s only assigment operator '=' is supported", a.GetGrlText())
	}
	n := ""
	if validAssignment(a) {
		n = strings.Trim(a.Variable.ArrayMapSelector.Expression.ExpressionAtom.Constant.GetGrlText(), `"`)
	}
	if n == "" {
		return errors.New("invalid assignment: " + a.GetGrlText())
	}
	t.Actions = append(t.Actions, ecarule.Action{Resource: n, Assignment: a})
	return nil
}

// AcceptExpression will accept an Expression into the Condition of this remote task.
// It is used by EcaruleParserListener.
// One who wants to use this function in a different context should make sure that the Expression
// satisfies the constraints implied by this remote task and by the Rule owner.
func (t tmpExpressionReceiver) AcceptExpression(exp *ast.Expression) error {
	if t.Condition != nil {
		return errors.New("task condition already assigned")
	}
	t.Condition = exp
	return nil
}
