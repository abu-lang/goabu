// Copyright 2024 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	"fmt"

	"github.com/abu-lang/goabu/ecarule"
	antlr_parser "github.com/abu-lang/goabu/parser/internal/antlr"
	"github.com/abu-lang/goabu/stringset"

	"github.com/antlr4-go/antlr/v4"
	grule_parser "github.com/hyperjumptech/grule-rule-engine/antlr"
	"github.com/hyperjumptech/grule-rule-engine/antlr/parser/grulev3"
	"github.com/hyperjumptech/grule-rule-engine/ast"
)

// remoteParserState is a parserState for parsing remote tasks of a GoAbU rule for allowing transfer to other nodes.
type remoteParserState struct {
	// baseParserState implements basic parserState behaviour.
	baseParserState
	// processing contains the data related to the rule currently being parsed.
	*processing
	// rewriter is used to swap local and remote resources before sending the task to the other nodes.
	rewriter *antlr.TokenStreamRewriter
	// localResources contains the names of the current task's local resources (before the swap performed by rewriter).
	localResources stringset.Set
	// remoteResources contains the names of the current task's remote resources (before the swap performed by rewriter).
	remoteResources stringset.Set
}

// newRemoteParserState constructs a remoteParserState given a [*grule_parser.GruleV3ParserListener]
// and a pointer to the parser's processing struct.
func newRemoteParserState(parser *grule_parser.GruleV3ParserListener, proc *processing) *remoteParserState {
	return &remoteParserState{
		baseParserState: baseParserState{
			GruleV3ParserListener: parser,
		},
		processing:      proc,
		localResources:  stringset.Make(),
		remoteResources: stringset.Make(),
	}
}

// reset prepares the parser for parsing a different [antlr_parser.TokenStream].
func (l *remoteParserState) reset(tokenStream antlr.TokenStream) {
	l.baseParserState.reset(tokenStream)
	l.rewriter = antlr.NewTokenStreamRewriter(tokenStream)
}

// EnterTask is called when production task is entered.
func (l *remoteParserState) EnterTask(ctx *antlr_parser.TaskContext) {
	if l.StopParse {
		return
	}
	l.remoteTasks = append(l.remoteTasks, ecarule.RemoteTask{})
	l.Stack.Push(nullExpressionReceiver{&l.remoteTasks[len(l.remoteTasks)-1]})
	l.localResources = stringset.Make()
	l.remoteResources = stringset.Make()
}

// ExitTask is called when production task is exited.
func (l *remoteParserState) ExitTask(ctx *antlr_parser.TaskContext) {
	exprRec, ok := l.Stack.Peek().(nullExpressionReceiver)
	if !ok {
		l.StopParse = true
		return
	}
	task := exprRec.RemoteTask
	modifiedExp := l.rewriter.GetText("default", antlr.Interval{
		Start: ctx.Expression().GetStart().GetTokenIndex(),
		Stop:  ctx.Expression().GetStop().GetTokenIndex(),
	})
	task.Condition = modifiedExp
	task.LocalResources = l.localResources.Slice()
	task.RemoteResources = l.remoteResources.Slice()
	l.Stack.Pop()
}

// EnterAssignment is called when production assignment is entered.
func (l *remoteParserState) EnterAssignment(ctx *grulev3.AssignmentContext) {
	l.inAssignLeft = true

	l.GruleV3ParserListener.EnterAssignment(ctx)
}

// ExitAssignment is called when production assignment is exited.
func (l *remoteParserState) ExitAssignment(ctx *grulev3.AssignmentContext) {
	l.Stack.Pop()
	exprRec, ok := l.Stack.Peek().(nullExpressionReceiver)
	if !ok {
		l.StopParse = true
		return
	}
	modifiedExp := l.rewriter.GetText("default", antlr.Interval{Start: ctx.GetStart().GetTokenIndex(), Stop: ctx.GetStop().GetTokenIndex()})
	exprRec.Actions = append(exprRec.Actions, modifiedExp)
}

// ExitVariable is called when production variable is exited.
func (l *remoteParserState) ExitVariable(ctx *grulev3.VariableContext) {
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
	remote := false
	switch {
	case ctx.SIMPLENAME() != nil && len(ctx.SIMPLENAME().GetText()) > 0:
		name = ctx.SIMPLENAME().GetText()
		if name == "this" {
			l.rewriter.ReplaceTokenDefault(ctx.SIMPLENAME().GetSymbol(), ctx.SIMPLENAME().GetSymbol(), "ext")
			return
		}
		if name == "ext" {
			l.rewriter.ReplaceTokenDefault(ctx.SIMPLENAME().GetSymbol(), ctx.SIMPLENAME().GetSymbol(), "this")
			return
		}
		if l.inAssignLeft {
			remote = true
			l.rewriter.InsertBeforeToken(antlr.DefaultProgramName, ctx.SIMPLENAME().GetSymbol(), "this.")
		} else {
			l.rewriter.InsertBeforeToken(antlr.DefaultProgramName, ctx.SIMPLENAME().GetSymbol(), "ext.")
		}
	case e.Variable != nil && e.Variable.Name == "ext":
		name = e.Name
		remote = true
	case e.Variable != nil && e.Variable.Name == "this":
		if l.inAssignLeft {
			l.parseError(fmt.Errorf("local actions are not allowed in 'for all' tasks"))
			return
		}
		name = e.Name
	}
	presentType := false
	if remote {
		typ = "Void"
		presentType = true
	} else {
		typ, presentType = l.types[name]
	}
	if !presentType {
		l.parseError(fmt.Errorf("could not determine the type of %s", name))
	}
	var r *ast.Variable
	if !remote {
		r = newAssignVariable(l.KnowledgeBase.WorkingMemory, "this", typ, name)
		l.localResources.Insert(name)
	} else {
		r = newAssignVariable(l.KnowledgeBase.WorkingMemory, "ext", typ, name)
		l.remoteResources.Insert(name)
	}
	l.Stack.Pop()
	l.Stack.Push(r)
}

// EnterExpression is called when production expression is entered.
func (l *remoteParserState) EnterExpression(ctx *grulev3.ExpressionContext) {
	l.inAssignLeft = false

	l.GruleV3ParserListener.EnterExpression(ctx)
}
