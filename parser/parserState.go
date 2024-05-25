// Copyright 2024 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	antlr_parser "github.com/abu-lang/goabu/parser/internal/antlr"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// parserState defines the behavior shared by every state of ruleParser.
type parserState interface {
	antlr_parser.EcaruleParserListener
	// reset prepares the parser for parsing a different antlr_parser.TokenStream.
	reset(antlr.TokenStream)
	// parseError register an error, successive call to isParsingHalted will return true until the next reset.
	parseError(error)
	// isParsingHalted reports whether an error was encountered during the parsing of the current token stream.
	isParsingHalted() bool
	// pop removes and returns the top element from the parser's internal AST stack.
	pop() any
	// push adds an element to the top of the parser's internal AST stack.
	push(any)
}
