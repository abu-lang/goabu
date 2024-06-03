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

// receivedParserState is a parserState for parsing remote task received from other nodes.
type receivedParserState struct {
	// baseParserState implements basic parserState behaviour.
	baseParserState
	// processing contains the data related to the rule currently being parsed.
	*processing
	// remoteTypes tracks resource types of the sender node.
	remoteTypes map[string]string
	// isAccepting is false when the node is missing some resource of the received current task, that will be skipped.
	isAccepting bool
}

// newReceivedParserState constructs a receivedParserState given a [*grule_parser.GruleV3ParserListener]
// and a pointer to the parser's processing struct.
func newReceivedParserState(parser *grule_parser.GruleV3ParserListener, proc *processing) *receivedParserState {
	return &receivedParserState{
		baseParserState: baseParserState{
			GruleV3ParserListener: parser,
		},
		processing: proc,
	}
}

// EnterTask is called when production task is entered.
func (l *receivedParserState) EnterTask(ctx *antlr_parser.TaskContext) {
	if l.StopParse {
		return
	}
	l.isAccepting = true
	l.localTasks = append(l.localTasks, ecarule.LocalTask{})
	l.Stack.Push(&expressionReceiver{
		LocalTask:      &l.localTasks[len(l.localTasks)-1],
		isAccepting:    &l.isAccepting,
		isReceivedTask: true,
	})
}

// ExitTask is called when production task is exited.
func (l *receivedParserState) ExitTask(ctx *antlr_parser.TaskContext) {
	if l.StopParse {
		return
	}
	l.Stack.Pop()
}

// EnterAssignment is called when production assignment is entered.
func (l *receivedParserState) EnterAssignment(ctx *grulev3.AssignmentContext) {
	l.inAssignLeft = true
	l.isAccepting = true

	l.GruleV3ParserListener.EnterAssignment(ctx)
}

// ExitVariable is called when production variable is exited.
func (l *receivedParserState) ExitVariable(ctx *grulev3.VariableContext) {
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
	var remote bool = false
	switch {
	case ctx.SIMPLENAME() != nil && len(ctx.SIMPLENAME().GetText()) > 0:
		name = ctx.SIMPLENAME().GetText()
		if name == "this" || name == "ext" {
			return
		}
		if l.inAssignLeft {
			remote = true
		}
	case e.Variable != nil && e.Variable.Name == "ext":
		name = e.Name
		remote = true
	case e.Variable != nil && e.Variable.Name == "this":
		name = e.Name
	}
	presentType := false
	if remote {
		typ, presentType = l.remoteTypes[name]
	} else {
		typ, presentType = l.types[name]
	}
	if !presentType {
		if remote {
			l.parseError(fmt.Errorf("could not determine the type of %s", name))
		} else {
			l.isAccepting = false
		}
	}
	var r *ast.Variable
	if !remote {
		r = newAssignVariable(l.KnowledgeBase.WorkingMemory, "this", typ, name)
	} else {
		r = newAssignVariable(l.KnowledgeBase.WorkingMemory, "ext", typ, name)
	}
	l.Stack.Pop()
	l.Stack.Push(r)
}

// EnterExpression is called when production expression is entered.
func (l *receivedParserState) EnterExpression(ctx *grulev3.ExpressionContext) {
	l.inAssignLeft = false

	l.GruleV3ParserListener.EnterExpression(ctx)
}
