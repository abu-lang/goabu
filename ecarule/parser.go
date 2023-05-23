// Copyright 2023 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package ecarule

import "github.com/hyperjumptech/grule-rule-engine/ast"

// Parser is the interface implemented by parsers of GoAbU rules.
type Parser interface {
	// Parse parses a series of GoAbU rules.
	Parse(...string) ([]Rule, []error)
	// ParseExpressions parses a series of local expressions.
	ParseExpressions(...string) ([]*ast.Expression, []error)
	// ParseActions parses a series of local actions.
	ParseActions(string) ([]Action, []error)
}
