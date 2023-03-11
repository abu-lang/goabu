// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

// Package ecarule defines GoAbU ECA rules.
package ecarule

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
)

type Rule struct {
	Name   string
	Events []string
	Tasks  []Task
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

func (a Action) String() string {
	return a.Assignment.GetGrlText()
}
