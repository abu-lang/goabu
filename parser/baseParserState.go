// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	antlr_parser "github.com/abu-lang/goabu/parser/internal/antlr"

	"github.com/antlr4-go/antlr/v4"
	grule_parser "github.com/hyperjumptech/grule-rule-engine/antlr"
	"github.com/hyperjumptech/grule-rule-engine/antlr/parser/grulev3"
)

// baseParserState is a null object parserState that can be used for defining parserStates through struct embedding.
type baseParserState struct {
	// GruleV3ParserListener is the Grule rule engine's parser which the parser relies upon.
	*grule_parser.GruleV3ParserListener
}

// makeBaseParserState constructs a baseParserState given a [*grule_parser.GruleV3ParserListener].
func makeBaseParserState(parser *grule_parser.GruleV3ParserListener) parserState {
	return baseParserState{
		GruleV3ParserListener: parser,
	}
}

// reset brings back the listener to a clean state so that it can be used on a different stream.
func (l baseParserState) reset(tokenStream antlr.TokenStream) {
	l.PreviousNode = l.PreviousNode[:0]
	for l.Stack.Len() > 0 {
		l.Stack.Pop()
	}
	l.StopParse = false
	l.ErrorCallback.Errors = make([]error, 0)
}

// parseError registers an error, successive calls to isParsingHalted will return true until the next reset.
func (l baseParserState) parseError(err error) {
	l.StopParse = true
	l.ErrorCallback.AddError(err)
}

// isParsingHalted reports whether an error was encountered during the parsing of the current token stream.
func (l baseParserState) isParsingHalted() bool {
	return l.StopParse
}

// push adds an element to the top of the parser's internal AST stack.
func (l baseParserState) push(e any) {
	l.Stack.Push(e)
}

// pop removes and returns the top element from the parser's internal AST stack.
func (l baseParserState) pop() any {
	return l.Stack.Pop()
}

// EnterPrule is called when production prules is entered.
func (l baseParserState) EnterPrules(ctx *antlr_parser.PrulesContext) {}

// ExitPrule is called when production prules is exited.
func (l baseParserState) ExitPrules(ctx *antlr_parser.PrulesContext) {}

// EnterPrule is called when production prule is entered.
func (l baseParserState) EnterPrule(ctx *antlr_parser.PruleContext) {}

// ExitPrule is called when production prule is exited.
func (l baseParserState) ExitPrule(ctx *antlr_parser.PruleContext) {}

// EnterEvents is called when production events is entered.
func (l baseParserState) EnterEvents(ctx *antlr_parser.EventsContext) {}

// ExitEvents is called when production events is exited.
func (l baseParserState) ExitEvents(ctx *antlr_parser.EventsContext) {}

// EnterDefaultActions is called when production defaultActions is entered.
func (l baseParserState) EnterDefaultActions(ctx *antlr_parser.DefaultActionsContext) {}

// ExitDefaultActions is called when production defaultActions is exited.
func (l baseParserState) ExitDefaultActions(ctx *antlr_parser.DefaultActionsContext) {}

// EnterTask is called when production task is entered.
func (l baseParserState) EnterTask(ctx *antlr_parser.TaskContext) {}

// ExitTask is called when production task is exited.
func (l baseParserState) ExitTask(ctx *antlr_parser.TaskContext) {}

// EnterActions is called when production actions is entered.
func (l baseParserState) EnterActions(ctx *antlr_parser.ActionsContext) {}

// ExitActions is called when production actions is exited.
func (l baseParserState) ExitActions(ctx *antlr_parser.ActionsContext) {}

// EnterTailActions is called when production tailActions is entered.
func (l baseParserState) EnterTailActions(ctx *antlr_parser.TailActionsContext) {}

// ExitTailActions is called when production tailActions is exited.
func (l baseParserState) ExitTailActions(ctx *antlr_parser.TailActionsContext) {}

// EnterMaybeActions is called when production maybeActions is entered.
func (l baseParserState) EnterMaybeActions(ctx *antlr_parser.MaybeActionsContext) {}

// ExitMaybeActions is called when production maybeActions is exited.
func (l baseParserState) ExitMaybeActions(ctx *antlr_parser.MaybeActionsContext) {}

// ExitVariable is called when production variable is exited.
func (l baseParserState) ExitVariable(ctx *grulev3.VariableContext) {}
