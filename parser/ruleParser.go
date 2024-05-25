// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

// Package parser implements a parser for GoAbU ECA rules.
package parser

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/abu-lang/goabu/ecarule"
	antlr_parser "github.com/abu-lang/goabu/parser/internal/antlr"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	grule_parser "github.com/hyperjumptech/grule-rule-engine/antlr"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// localProcessing will contain the events and the local tasks of the rule currently being processed.
type localProcessing struct {
	// types tracks the resource types of the current nodes' resources.
	types map[string]string
	// events contains the event part of the rule currently being processed.
	events []string
	// localTasks contains the local tasks of the rule currently being processed.
	localTasks []ecarule.LocalTask
}

// processing will contain the events and the tasks of the rule currently being processed.
type processing struct {
	// localProcessing contains the events and the local tasks of the rule currently being processed.
	localProcessing
	// remoteTasks contains the remote tasks of the rule currently being processed.
	remoteTasks []ecarule.RemoteTask
	// inAssignLeft reports whether the parser is currently processing an l-value expression.
	inAssignLeft bool
}

// ruleParser is responsible for the parsing step of goabuParser by implementing an ANTLR listener.
// The operations of the ruleParser depends on its current state. Each state implements the common parserState interface.
type ruleParser struct {
	// parserState holds the current state of the ruleParser.
	parserState
	// local holds the state responsible for parsing local tasks.
	local *localParserState
	// remote holds the state responsible for parsing remote tasks and
	// obtaining tasks that can be send to other nodes.
	remote *remoteParserState
	// received holds the state responsible for parsing remote tasks received
	// from other nodes.
	received *receivedParserState
	// rules contains the GoAbU rules parsed by the ruleParser.
	rules []ecarule.Rule
	// processing contains the events and the tasks of the rule currently being processed.
	processing
}

// newRuleParser constructs a ruleParser given the node resource names, along with the resource types, and the node's [*ast.WorkingMemory].
func newRuleParser(types map[string]string, workingMemory *ast.WorkingMemory, errorReporter *pkg.GruleErrorReporter) *ruleParser {
	kb := &ast.KnowledgeBase{
		WorkingMemory: workingMemory,
	}
	gruleParser := grule_parser.NewGruleV3ParserListener(kb, errorReporter)
	res := &ruleParser{
		processing: processing{
			localProcessing: localProcessing{
				types: types,
			},
		},
	}
	res.local = newLocalParserState(gruleParser, &res.localProcessing)
	res.remote = newRemoteParserState(gruleParser, &res.processing)
	res.received = newReceivedParserState(gruleParser, &res.processing)
	res.parserState = res.local
	return res
}

// reset prepares the parser for parsing a different [antlr.TokenStream].
func (l *ruleParser) reset(tokenStream antlr.TokenStream) {
	l.local.reset(tokenStream)
	l.remote.reset(tokenStream)
	l.received.reset(tokenStream)
	l.rules = nil
}

// EnterPrule is called when production prule is entered.
func (l *ruleParser) EnterPrule(ctx *antlr_parser.PruleContext) {
	if l.isParsingHalted() {
		return
	}
	l.localTasks = make([]ecarule.LocalTask, 0)
	l.remoteTasks = make([]ecarule.RemoteTask, 0)
}

// ExitPrule is called when production prule is exited.
func (l *ruleParser) ExitPrule(ctx *antlr_parser.PruleContext) {
	if l.isParsingHalted() {
		return
	}
	l.rules = append(l.rules, ecarule.Rule{
		Name:        ctx.SIMPLENAME().GetText(),
		Events:      l.events,
		LocalTasks:  l.localTasks,
		RemoteTasks: l.remoteTasks,
	})
	l.events = nil
	l.localTasks = nil
	l.remoteTasks = nil
}

// EnterEvents is called when production events is entered.
func (l *ruleParser) EnterEvents(ctx *antlr_parser.EventsContext) {
	if l.isParsingHalted() {
		return
	}
	if len(l.events) > 0 {
		l.parseError(errors.New("syntax error"))
		return
	}
	for i := 0; ctx.SIMPLENAME(i) != nil; i++ {
		l.events = append(l.events, ctx.SIMPLENAME(i).GetText())
	}
}

// EnterDefaultActions is called when production defaultActions is entered.
func (l *ruleParser) EnterDefaultActions(ctx *antlr_parser.DefaultActionsContext) {
	if l.isParsingHalted() {
		return
	}
	cond, err := newBooleanLiteralExpression(l.local.KnowledgeBase.WorkingMemory, true)
	if err != nil {
		l.parseError(errors.New("error during default actions parsing"))
		return
	}
	l.localTasks = append(l.localTasks, ecarule.LocalTask{Condition: cond})

	l.push(&expressionReceiver{
		LocalTask:   &l.localTasks[len(l.localTasks)-1],
		isAccepting: nil,
	})
}

// ExitDefaultActions is called when production defaultActions is exited.
func (l *ruleParser) ExitDefaultActions(ctx *antlr_parser.DefaultActionsContext) {
	if l.isParsingHalted() {
		return
	}
	l.pop()
}

// EnterTask is called when production task is entered.
func (l *ruleParser) EnterTask(ctx *antlr_parser.TaskContext) {
	if l.isParsingHalted() {
		return
	}
	if l.parserState != l.received {
		if ctx.ALL() != nil {
			l.parserState = l.remote
		} else {
			l.parserState = l.local
		}
	}
	l.parserState.EnterTask(ctx)
}

// ExitTask is called when production task is exited.
func (l *ruleParser) ExitTask(ctx *antlr_parser.TaskContext) {
	l.parserState.ExitTask(ctx)
	if l.remote == l.parserState {
		l.parserState = l.local
	}
}

// newAssignVariable constructs a [*ast.Variable] encoding a GoAbU resource.
func newAssignVariable(workingMemory *ast.WorkingMemory, prefix, typ, name string) *ast.Variable {
	pre := ast.NewVariable()
	pre.Name = prefix
	tv := ast.NewVariable()
	tv.Name = typ
	tv.Variable = workingMemory.AddVariable(pre)
	res := ast.NewVariable()
	res.Variable = workingMemory.AddVariable(tv)
	res.ArrayMapSelector = newResourceArrayMapSelector(workingMemory, name)
	res.SetGrlText(fmt.Sprintf(`%s.%s["%s"]`, prefix, typ, name))
	return res
}

// newResourceArrayMapSelector constructs an [*ast.ArrayMapSelector] where r is used as the key.
func newResourceArrayMapSelector(workingMemory *ast.WorkingMemory, r string) *ast.ArrayMapSelector {
	val := reflect.ValueOf(r)
	c := ast.NewConstant()
	c.Value = val
	c.SetGrlText(fmt.Sprintf(`"%s"`, r))
	a := ast.NewExpressionAtom()
	a.Constant = c
	e := ast.NewExpression()
	e.ExpressionAtom = workingMemory.AddExpressionAtom(a)
	res := ast.NewArrayMapSelector()
	res.Expression = workingMemory.AddExpression(e)
	return res
}

// newBooleanLiteralExpression creates an [*ast.Expression] encoding the specified boolean value.
func newBooleanLiteralExpression(workingMemory *ast.WorkingMemory, b bool) (*ast.Expression, error) {
	text := "false"
	if b {
		text = "true"
	}
	cons := ast.NewConstant()
	cons.GrlText = text
	cons.AcceptBooleanLiteral(&ast.BooleanLiteral{Boolean: b})
	atm := ast.NewExpressionAtom()
	atm.GrlText = text
	err := atm.AcceptConstant(cons)
	if err != nil {
		return nil, err
	}
	exp := ast.NewExpression()
	exp.GrlText = text
	err = exp.AcceptExpressionAtom(workingMemory.AddExpressionAtom(atm))
	if err != nil {
		return nil, err
	}
	return workingMemory.AddExpression(exp), nil
}
