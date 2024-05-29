// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	"fmt"

	"github.com/abu-lang/goabu/ecarule"
	antlr_parser "github.com/abu-lang/goabu/parser/internal/antlr"

	grule_parser "github.com/hyperjumptech/grule-rule-engine/antlr"
	"github.com/hyperjumptech/grule-rule-engine/antlr/parser/grulev3"
	"github.com/hyperjumptech/grule-rule-engine/ast"
)

// localParserState is a parserState for parsing local tasks.
type localParserState struct {
	// baseParserState implements basic parserState behaviour.
	baseParserState
	// localProcessing contains the data related to the rule currently being parsed.
	*localProcessing
}

// newLocalParserState constructs a localParserState given a [*grule_parser.GruleV3ParserListener]
// and a pointer to the parser's localProcessing struct.
func newLocalParserState(parser *grule_parser.GruleV3ParserListener, proc *localProcessing) *localParserState {
	return &localParserState{
		baseParserState: baseParserState{
			GruleV3ParserListener: parser,
		},
		localProcessing: proc,
	}
}

// EnterTask is called when production task is entered.
func (l *localParserState) EnterTask(ctx *antlr_parser.TaskContext) {
	if l.StopParse {
		return
	}
	l.localTasks = append(l.localTasks, ecarule.LocalTask{})
	l.Stack.Push(&expressionReceiver{
		LocalTask:   &l.localTasks[len(l.localTasks)-1],
		isAccepting: nil,
	})
}

// ExitTask is called when production task is exited.
func (l *localParserState) ExitTask(ctx *antlr_parser.TaskContext) {
	if l.StopParse {
		return
	}
	l.Stack.Pop()
}

// ExitVariable is called when production variable is exited.
func (l *localParserState) ExitVariable(ctx *grulev3.VariableContext) {
	defer l.GruleV3ParserListener.ExitVariable(ctx)
	if l.StopParse {
		return
	}
	e, ok := l.Stack.Peek().(*ast.Variable)
	if !ok {
		return
	}
	var name string
	var typ string
	switch {
	case ctx.SIMPLENAME() != nil && len(ctx.SIMPLENAME().GetText()) > 0:
		name = ctx.SIMPLENAME().GetText()
		if name == "this" || name == "ext" {
			return
		}
	case e.Variable != nil && e.Variable.Name == "ext":
		l.parseError(fmt.Errorf("external variable %s is not allowed in this context", e.GetGrlText()))
		return
	case e.Variable != nil && e.Variable.Name == "this":
		name = e.Name
	}
	presentType := false
	typ, presentType = l.types[name]
	if !presentType {
		l.parseError(fmt.Errorf("could not determine the type of %s", name))
	}
	r := newAssignVariable(l.KnowledgeBase.WorkingMemory, "this", typ, name)
	l.Stack.Pop()
	l.Stack.Push(r)
}
