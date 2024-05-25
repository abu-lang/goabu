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

// Action groups an assignment with the name of the involved resource.
type Action struct {
	// Assignment is a simple assignment to Resource's variable. The expression contains only local resources.
	Assignment *ast.Assignment
	// Resource is the name of the resource in the L-value of Assignment.
	Resource string
}

// LocalTask models a local task allowed to modify only local resources.
type LocalTask struct {
	// Condition is an expression with boolean result indicating whether the activated rule is to be evaluated.
	Condition *ast.Expression
	// Actions is a list of assignments where only local resources can appear.
	Actions []Action
}

// RemoteTask models a remote task that can update the resources of the other nodes.
// In remote tasks, remote resources are prefixed with "this." while the local resources are prefixed with "ext.".
type RemoteTask struct {
	// Condition encodes the rule's condition.
	Condition string
	// Actions encodes the actions that are to be performed.
	Actions []string
	// RemoteResources contains all the names of the remote resources of the task.
	RemoteResources []string
	// LocalResources contains all the names of the local resources of the task.
	LocalResources []string
}

// String returns the code of the action's assignment.
func (a Action) String() string {
	return a.Assignment.GetGrlText()
}
