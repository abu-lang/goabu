// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

// Package ecarule defines GoAbU ECA rules.
package ecarule

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
	Tasks          []Task
}

type Action struct {
	Resource   string
	Assignment *ast.Assignment
}

type Task struct {
	External  bool
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
	n := ""
	if validAssignment(a) {
		n = strings.Trim(a.Variable.ArrayMapSelector.Expression.ExpressionAtom.Constant.GetGrlText(), `"`)
	}
	if n == "" {
		return errors.New("invalid assignment: " + a.GetGrlText())
	}
	r.DefaultActions = append(r.DefaultActions, Action{Resource: n, Assignment: a})
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
	n := ""
	if validAssignment(a) {
		n = strings.Trim(a.Variable.ArrayMapSelector.Expression.ExpressionAtom.Constant.GetGrlText(), `"`)
	}
	if n == "" {
		return errors.New("invalid assignment: " + a.GetGrlText())
	}
	t.Actions = append(t.Actions, Action{Resource: n, Assignment: a})
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
	return a.Assignment.GetGrlText()
}

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
