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

// expressionReceiver is a struct wrapping an [*ecarule.LocalTask] for allowing tasks to implement the
// receiver interfaces required by Grule rule engine's parser (without exporting the methods on
// the [ecarule.LocalTask] type).
type expressionReceiver struct {
	*ecarule.LocalTask
	isAccepting    *bool
	isReceivedTask bool
}

// AcceptAssignment will accept an [*ast.Assignment] into this local task by creating a new corresponding [ecarule.Action].
// It is used by ruleParser.
// One who wants to use this function in a different context should make sure that the [*ast.Assignment]
// satisfies the constraints implied by this local task and by the rule owner.
func (t *expressionReceiver) AcceptAssignment(a *ast.Assignment) error {
	if t.isAccepting != nil && !*t.isAccepting {
		return nil
	}
	if !a.IsAssign {
		return fmt.Errorf("assignment %s only assignment operator '=' is supported", a.GetGrlText())
	}
	n := ""
	if t.validAssignment(a) {
		n = strings.Trim(a.Variable.ArrayMapSelector.Expression.ExpressionAtom.Constant.GetGrlText(), `"`)
	}
	if n == "" {
		return errors.New("invalid assignment: " + a.GetGrlText())
	}
	t.Actions = append(t.Actions, ecarule.Action{Resource: n, Assignment: a})
	return nil
}

// AcceptExpression will accept an [*ast.Expression] into the [ecarule.Condition] of this local task.
// It is used by ruleParser.
// One who wants to use this function in a different context should make sure that the [*ast.Expression]
// satisfies the constraints implied by this local task and by the rule owner.
func (t *expressionReceiver) AcceptExpression(exp *ast.Expression) error {
	if t.isAccepting != nil && !*t.isAccepting {
		return nil
	}
	if t.Condition != nil {
		return errors.New("task condition already assigned")
	}
	t.Condition = exp
	return nil
}

// validAssignment performs some checks on the passed [*ast.Assignment] for verifying whether the l-value denotes
// a plausible parsed GoAbU rule's resource.
func (t *expressionReceiver) validAssignment(a *ast.Assignment) bool {
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
		ok = t.isTypeOk(v.Variable.Variable.Name, v.Variable.Name)
	}
	return ok
}

// isTypeOk reports whether the resource type is coherent with the current task.
func (t *expressionReceiver) isTypeOk(prefix, typ string) bool {
	switch typ {
	case "Bool", "Integer", "Float", "Text", "Time", "Other":
		if prefix == "this" || (t.isReceivedTask && prefix == "ext") {
			return true
		}
	case "Void":
		if prefix == "ext" && !t.isReceivedTask {
			return true
		}
	}
	return false
}
