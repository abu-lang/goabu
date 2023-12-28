// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

// Package ecarule defines GoAbU ECA rules.
package ecarule

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
)

// Rule models a GoAbU (Event-Condition-Action) rule.
type Rule struct {
	// Name specifies the rule's name.
	Name string
	// Events is a list of resource names. The rule is activated when any of the listed resources changes its value.
	Events []string
	// LocalTasks contains the rule's local tasks that can modify only local resources when the condition matches.
	LocalTasks []LocalTask
	// RemoteTasks contains the rule's remote tasks that modify the resources of the other nodes matching the condition.
	RemoteTasks []RemoteTask
}

type Action struct {
	Resource   string
	Assignment *ast.Assignment
}

// LocalTask models a local task allowed to modify only local resources.
type LocalTask struct {
	// Condition is an expression with boolean result indicating whether the activated rule is to be evaluated
	Condition *ast.Expression
	// Actions is a list of assignment where only local resources can appear.
	Actions []Action
}

// RemoteTask models a remote task that can update the resources of the other nodes.
type RemoteTask struct {
	Condition *ast.Expression
	Actions   []Action
}

func (a Action) String() string {
	return a.Assignment.GetGrlText()
}
