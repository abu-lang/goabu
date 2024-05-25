// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	"sync"

	"github.com/abu-lang/goabu/ecarule"
	antlr_parser "github.com/abu-lang/goabu/parser/internal/antlr"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// brokenLock implements the [sync.Locker] interface but does not lock anything.
type brokenLock struct{}

// Lock is a no-op.
func (l brokenLock) Lock() {}

// Unlock is a no-op.
func (l brokenLock) Unlock() {}

// goabuParser is the ANTLR4 based parser for GoAbU rules implementing the [ecarule.Parser]
// interface by relying on the Grule Rule Engine's parser.
type goabuParser struct {
	// lexer is the goabuParser's lexer.
	lexer *antlr_parser.EcaruleLexer
	// lexer is the goabuParser's parser.
	parser *antlr_parser.EcaruleParser
	// errListener is used for registering errors that occur during parsing.
	errListener *pkg.GruleErrorReporter

	// listener builds the rules from the AST using the ANTLR4 listener interface.
	listener *ruleParser
	// lockMemory is the lock that the goabuParser will acquire before operating on the node's working memory.
	lockMemory sync.Locker
}

// New takes as arguments the types of the local resources specified as [github.com/abu-lang/goabu/memory.Resources.Types]
// and an [*ast.WorkingMemory] and creates an [ecarule.Parser]. The parsed expressions will be added to the
// [*ast.WorkingMemory].
//
// If a single [sync.Locker] is passed as argument then the parser will acquire the lock before performing
// actions on the [*ast.WorkingMemory]. Only a single [sync.Locker] can be passed as argument.
func New(types map[string]string, workingMemory *ast.WorkingMemory, args ...any) ecarule.Parser {
	if len(args) > 1 {
		return nil
	}
	res := &goabuParser{errListener: &pkg.GruleErrorReporter{Errors: make([]error, 0)}}
	if len(args) > 0 {
		locker, ok := args[0].(sync.Locker)
		if !ok || locker == nil {
			return nil
		}
		res.lockMemory = locker
	} else {
		res.lockMemory = brokenLock{}
	}
	res.lexer = antlr_parser.NewEcaruleLexer(antlr.NewInputStream(""))
	res.lexer.RemoveErrorListeners()
	res.lexer.AddErrorListener(res.errListener)
	res.parser = antlr_parser.NewEcaruleParser(antlr.NewCommonTokenStream(res.lexer, antlr.TokenDefaultChannel))
	res.parser.BuildParseTrees = true
	res.parser.RemoveErrorListeners()
	res.parser.AddErrorListener(res.errListener)
	res.listener = newRuleParser(types, workingMemory, res.errListener)
	return res
}

// reset prepares the parser for the parsing of a different stream.
func (p *goabuParser) reset(input string) {
	p.lexer.SetInputStream(antlr.NewInputStream(input))
	ts := antlr.NewCommonTokenStream(p.lexer, antlr.TokenDefaultChannel)
	p.listener.reset(ts)
	p.parser.SetInputStream(ts)
}

// errors retrieves the errors encountered during the parsing.
func (lp *goabuParser) errors() []error {
	return lp.errListener.Errors
}

// Parse parses a series of GoAbU rules.
func (p *goabuParser) Parse(rules ...string) ([]ecarule.Rule, []error) {
	res := make([]ecarule.Rule, 0, len(rules))
	p.lockMemory.Lock()
	defer p.lockMemory.Unlock()
	for _, r := range rules {
		p.reset(r)
		tree := p.parser.Prules()
		errs := p.errors()
		if len(errs) > 0 {
			return nil, errs
		}
		antlr.ParseTreeWalkerDefault.Walk(p.listener, tree)
		errs = p.errors()
		if len(errs) > 0 {
			return nil, errs
		}
		res = append(res, p.listener.rules...)
	}
	// update WorkingMemory
	p.listener.local.KnowledgeBase.WorkingMemory.IndexVariables()
	return res, nil
}

// ParseActions parses a series of local actions.
func (p *goabuParser) ParseActions(actions string) ([]ecarule.Action, []error) {
	p.reset(actions)
	tree := p.parser.Actions()
	errs := p.errors()
	if len(errs) > 0 {
		return nil, errs
	}
	task := ecarule.LocalTask{}
	p.listener.push(&expressionReceiver{&task, nil, false})
	p.lockMemory.Lock()
	antlr.ParseTreeWalkerDefault.Walk(p.listener, tree)
	// update WorkingMemory
	p.listener.local.KnowledgeBase.WorkingMemory.IndexVariables()
	p.lockMemory.Unlock()
	errs = p.errors()
	if len(errs) > 0 {
		return nil, errs
	}
	return task.Actions, nil
}

// ParseExpressions parses a series of local expressions.
func (p *goabuParser) ParseExpressions(exps ...string) ([]*ast.Expression, []error) {
	res := make([]*ast.Expression, 0, len(exps))
	p.lockMemory.Lock()
	defer p.lockMemory.Unlock()
	task := ecarule.LocalTask{}
	for _, exp := range exps {
		p.reset(exp)
		p.listener.push(&expressionReceiver{&task, nil, false})
		tree := p.parser.Expression()
		errs := p.errors()
		if len(errs) > 0 {
			return nil, errs
		}
		antlr.ParseTreeWalkerDefault.Walk(p.listener, tree)
		errs = p.errors()
		if len(errs) > 0 {
			return nil, errs
		}
		res = append(res, task.Condition)
		task.Condition = nil
	}
	// update WorkingMemory
	p.listener.local.KnowledgeBase.WorkingMemory.IndexVariables()
	return res, nil
}

// ParseRemoteTasks parses a series of received tasks into local tasks that can be executed.
func (p *goabuParser) ParseRemoteTasks(remoteTypes map[string]string, tasks ...ecarule.RemoteTask) ([]ecarule.LocalTask, []error) {
	res := make([]ecarule.LocalTask, 0)
	p.lockMemory.Lock()
	defer p.lockMemory.Unlock()
	for _, rTask := range tasks {
		str := "for " + rTask.Condition + " do "
		for _, act := range rTask.Actions {
			str += act
			str += ", "
		}
		p.reset(str)
		tree := p.parser.Task()
		errs := p.errors()
		if len(errs) > 0 {
			return nil, errs
		}
		p.listener.received.remoteTypes = remoteTypes
		p.listener.parserState = p.listener.received
		antlr.ParseTreeWalkerDefault.Walk(p.listener, tree)
		errs = p.errors()
		if len(errs) > 0 {
			return nil, errs
		}

		lTask := p.listener.localTasks[0]
		if lTask.Condition != nil {
			res = append(res, lTask)
		}
		p.listener.parserState = p.listener.local
		p.listener.received.remoteTypes = nil
	}
	// update WorkingMemory
	p.listener.local.KnowledgeBase.WorkingMemory.IndexVariables()
	return res, nil
}
