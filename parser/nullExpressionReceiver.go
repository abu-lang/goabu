// Copyright 2024 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	"github.com/abu-lang/goabu/ecarule"
	"github.com/hyperjumptech/grule-rule-engine/ast"
)

// nullExpressionReceiver is a null object expression receiver.
type nullExpressionReceiver struct {
	*ecarule.RemoteTask
}

// AcceptAssignment is a no-op.
func (n nullExpressionReceiver) AcceptAssignment(a *ast.Assignment) error {
	return nil
}

// AcceptExpression is a no-op.
func (n nullExpressionReceiver) AcceptExpression(exp *ast.Expression) error {
	return nil
}
