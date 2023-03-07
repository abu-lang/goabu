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

// goabuParser is an ANTLR4 based parser for GoAbU rules implementing the [ecarule.Parser]
// interface relying on the Grule Rule Engine's parser.
type goabuParser struct {
	lexer       *antlr_parser.EcaruleLexer
	parser      *antlr_parser.EcaruleParser
	errListener *pkg.GruleErrorReporter

	listener   *ecaruleParserListener
	lockMemory sync.Locker
}

// brokenLock implements the [sync.Locker] interface but does not lock anything.
type brokenLock struct{}

// Lock is a no-op.
func (l brokenLock) Lock() {}

// Unlock is a no-op.
func (l brokenLock) Unlock() {}

// New takes as arguments the types of the local resources specified as [github.com/abu-lang/goabu/memory.Resources.Types]
// and an [*ast.WorkingMemory] and creates an [ecarule.Parser]. The parsed expressions will be added to the
// [*ast.WorkingMemory].
//
// If a single [sync.Locker] is passed as argument then the parser will acquire the lock before performing
// actions on the [*ast.WorkingMemory]. Only a single [sync.Locker] can be passed as argument.
func New(types map[string]string, workingMemory *ast.WorkingMemory, memoryLocker ...sync.Locker) ecarule.Parser {
	if len(memoryLocker) > 1 {
		return nil
	}
	res := &goabuParser{errListener: &pkg.GruleErrorReporter{Errors: make([]error, 0)}}
	if len(memoryLocker) == 1 {
		res.lockMemory = memoryLocker[0]
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
	res.listener = newEcaruleParserListener(types, workingMemory, res.errListener)
	return res
}

// reset prepares the parser for the parsing of a different string.
func (p *goabuParser) reset(input string) {
	p.listener.reset()
	p.lexer.SetInputStream(antlr.NewInputStream(input))
	p.parser.SetInputStream(antlr.NewCommonTokenStream(p.lexer, antlr.TokenDefaultChannel))
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
		tree := p.parser.Prule()
		errs := p.errors()
		if len(errs) > 0 {
			return nil, errs
		}
		antlr.ParseTreeWalkerDefault.Walk(p.listener, tree)
		errs = p.errors()
		if len(errs) > 0 {
			return nil, errs
		}
		res = append(res, *p.listener.Rule)
	}
	// update WorkingMemory
	p.listener.KnowledgeBase.WorkingMemory.IndexVariables()
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
	p.lockMemory.Lock()
	antlr.ParseTreeWalkerDefault.Walk(p.listener, tree)
	// update WorkingMemory
	p.listener.KnowledgeBase.WorkingMemory.IndexVariables()
	p.lockMemory.Unlock()
	errs = p.errors()
	if len(errs) > 0 {
		return nil, errs
	}
	return p.listener.Rule.DefaultActions, nil
}

// ParseExpressions parses a series of local expressions.
func (p *goabuParser) ParseExpressions(exps ...string) ([]*ast.Expression, []error) {
	res := make([]*ast.Expression, 0, len(exps))
	p.lockMemory.Lock()
	defer p.lockMemory.Unlock()
	task := ecarule.Task{}
	for _, exp := range exps {
		p.reset(exp)
		p.listener.Stack.Push(&task)
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
	p.listener.KnowledgeBase.WorkingMemory.IndexVariables()
	return res, nil
}
