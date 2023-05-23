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

// expressionReceiver is a struct wrapping an [*ecarule.Task] for allowing tasks to implement the
// receiver interfaces required by Grule rule engine's parser (without exporting the methods on
// the [ecarule.Task] type).
type expressionReceiver struct {
	*ecarule.Task
}

// AcceptAssignment will accept an Assignment into this Task by creating a new corresponding Action.
// It is used by EcaruleParserListener.
// One who wants to use this function in a different context should make sure that the Assignment
// satisfies the constraints implied by this Task and by the Rule owner.
func (t expressionReceiver) AcceptAssignment(a *ast.Assignment) error {
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

// AcceptExpression will accept an Expression into the Condition of this Task.
// It is used by EcaruleParserListener.
// One who wants to use this function in a different context should make sure that the Expression
// satisfies the constraints implied by this Task and by the Rule owner.
func (t expressionReceiver) AcceptExpression(exp *ast.Expression) error {
	if t.Condition != nil {
		return errors.New("task condition already assigned")
	}
	t.Condition = exp
	return nil
}

// validAssignment performs some checks on the passed [*ast.Assignment] for verifying if the l-value denotes
// a plausible parsed GoAbU rule's resource.
func validAssignment(a *ast.Assignment) bool {
	if a.Variable == nil ||
		a.Variable.ArrayMapSelector == nil ||
		a.Variable.ArrayMapSelector.Expression == nil ||
		a.Variable.ArrayMapSelector.Expression.ExpressionAtom == nil ||
		a.Variable.ArrayMapSelector.Expression.ExpressionAtom.Constant == nil {
		return false
	}
	ok := false
	v := a.Variable
	if v.Variable != nil && v.Variable.Variable != nil && v.Variable.Variable.Variable == nil {
		switch v.Variable.Name {
		case "Bool", "Integer", "Float", "Text", "Time", "Other":
			if v.Variable.Variable.Name == "this" {
				ok = true
			}
		case "Void":
			if v.Variable.Variable.Name == "ext" {
				ok = true
			}
		}
	}
	return ok
}
