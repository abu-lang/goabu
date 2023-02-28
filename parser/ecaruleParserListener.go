// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

// Package parser implements a parser for GoAbU ECA rules.
package parser

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/abu-lang/goabu/ecarule"
	antlr_parser "github.com/abu-lang/goabu/parser/antlr"

	"github.com/hyperjumptech/grule-rule-engine/antlr"
	"github.com/hyperjumptech/grule-rule-engine/antlr/parser/grulev3"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

type ecaruleParserListener struct {
	*antlr.GruleV3ParserListener
	Rule         *ecarule.Rule
	types        map[string]string
	allowExt     bool
	inAssignLeft bool
}

func newEcaruleParserListener(types map[string]string, workingMemory *ast.WorkingMemory, errorReporter *pkg.GruleErrorReporter) *ecaruleParserListener {
	kb := &ast.KnowledgeBase{
		WorkingMemory: workingMemory,
	}
	res := &ecaruleParserListener{
		GruleV3ParserListener: antlr.NewGruleV3ParserListener(kb, errorReporter),
		types:                 types,
		Rule:                  &ecarule.Rule{},
	}
	res.Stack.Push(res.Rule)
	return res
}

// reset brings back the listener to a clean state so that it can be used on a differet tree.
func (l *ecaruleParserListener) reset() {
	l.Rule = &ecarule.Rule{}
	l.PreviousNode = l.PreviousNode[:0]
	for l.Stack.Len() > 0 {
		l.Stack.Pop()
	}
	l.Stack.Push(l.Rule)
	l.StopParse = false
	l.ErrorCallback.Errors = make([]error, 0)

}

func (l *ecaruleParserListener) parseError(err error) {
	l.StopParse = true
	l.ErrorCallback.AddError(err)
}

// EnterPrule is called when production prule is entered.
func (l *ecaruleParserListener) EnterPrule(ctx *antlr_parser.PruleContext) {
	if l.StopParse {
		return
	}
	l.Rule.Name = ctx.SIMPLENAME().GetText()
	t := 0
	for ; ctx.Task(t) != nil; t++ {
	}
	l.Rule.Tasks = make([]ecarule.Task, 0, t)
}

// ExitPrule is called when production prule is exited.
func (l *ecaruleParserListener) ExitPrule(ctx *antlr_parser.PruleContext) {
	if l.StopParse {
		return
	}
	l.Stack.Pop()
}

// EnterEvents is called when production events is entered.
func (l *ecaruleParserListener) EnterEvents(ctx *antlr_parser.EventsContext) {
	if l.StopParse {
		return
	}
	rule, ok := l.Stack.Peek().(*ecarule.Rule)
	if !ok {
		l.parseError(errors.New("syntax error"))
		return
	}
	for i := 0; ctx.SIMPLENAME(i) != nil; i++ {
		rule.Events = append(rule.Events, ctx.SIMPLENAME(i).GetText())
	}
}

// ExitEvents is called when production events is exited.
func (l *ecaruleParserListener) ExitEvents(ctx *antlr_parser.EventsContext) {}

// EnterTask is called when production task is entered.
func (l *ecaruleParserListener) EnterTask(ctx *antlr_parser.TaskContext) {
	if l.StopParse {
		return
	}
	rule, ok := l.Stack.Peek().(*ecarule.Rule)
	if !ok {
		l.parseError(errors.New("syntax error"))
		return
	}
	all := ctx.ALL() != nil
	rule.Tasks = append(rule.Tasks, ecarule.Task{External: all})
	l.allowExt = all
	l.Stack.Push(&rule.Tasks[len(rule.Tasks)-1])
}

// ExitTask is called when production task is exited.
func (l *ecaruleParserListener) ExitTask(ctx *antlr_parser.TaskContext) {
	if l.StopParse {
		return
	}
	l.Stack.Pop()
	t, ok := l.Stack.Peek().(*ecarule.Task)
	if ok {
		l.allowExt = t.External
	} else {
		l.allowExt = false
	}
}

// EnterActions is called when production actions is entered.
func (l *ecaruleParserListener) EnterActions(ctx *antlr_parser.ActionsContext) {}

// ExitActions is called when production actions is exited.
func (l *ecaruleParserListener) ExitActions(ctx *antlr_parser.ActionsContext) {}

// EnterTailActions is called when production tailActions is entered.
func (l *ecaruleParserListener) EnterTailActions(ctx *antlr_parser.TailActionsContext) {}

// ExitTailActions is called when production tailActions is exited.
func (l *ecaruleParserListener) ExitTailActions(ctx *antlr_parser.TailActionsContext) {}

// EnterMaybeActions is called when production maybeActions is entered.
func (l *ecaruleParserListener) EnterMaybeActions(ctx *antlr_parser.MaybeActionsContext) {}

// ExitMaybeActions is called when production maybeActions is exited.
func (l *ecaruleParserListener) ExitMaybeActions(ctx *antlr_parser.MaybeActionsContext) {}

func (l *ecaruleParserListener) EnterAssignment(ctx *grulev3.AssignmentContext) {
	l.inAssignLeft = true

	l.GruleV3ParserListener.EnterAssignment(ctx)
}

func (l *ecaruleParserListener) ExitVariable(ctx *grulev3.VariableContext) {
	if !l.StopParse {
		var r *ast.Variable = nil
		e, ok := l.Stack.Peek().(*ast.Variable)
		if ok {
			switch {
			case ctx.SIMPLENAME() != nil && len(ctx.SIMPLENAME().GetText()) > 0:
				n := ctx.SIMPLENAME().GetText()
				if n == "this" || n == "ext" {
					break
				}
				if l.inAssignLeft && l.allowExt {
					r = l.newExtAssignVariable(n)
				} else {
					t, present := l.types[n]
					if !present {
						l.parseError(fmt.Errorf("could not determine the type of %s", n))
						break
					}
					r = l.newThisAssignVariable(n, t)
				}
			case e.Variable != nil && e.Variable.Name == "ext":
				if !l.allowExt {
					l.parseError(fmt.Errorf("external variable %s is not allowed in this context", e.GetGrlText()))
					break
				}
				r = l.newExtAssignVariable(e.Name)
			case e.Variable != nil && e.Variable.Name == "this":
				if l.inAssignLeft && l.allowExt {
					l.parseError(fmt.Errorf("local actions are not allowed in 'for all' tasks"))
					break
				}
				n := e.Name
				t, present := l.types[n]
				if !present {
					l.parseError(fmt.Errorf("could not determine the type of %s", n))
					break
				}
				r = l.newThisAssignVariable(n, t)
			}
		}
		if r != nil {
			l.Stack.Pop()
			l.Stack.Push(r)
		}
	}

	l.GruleV3ParserListener.ExitVariable(ctx)
}

func (l *ecaruleParserListener) EnterExpression(ctx *grulev3.ExpressionContext) {
	l.inAssignLeft = false

	l.GruleV3ParserListener.EnterExpression(ctx)
}

func (l *ecaruleParserListener) newThisAssignVariable(r string, t string) *ast.Variable {
	this := ast.NewVariable()
	this.Name = "this"
	tv := ast.NewVariable()
	tv.Name = t
	tv.Variable = l.KnowledgeBase.WorkingMemory.AddVariable(this)
	res := ast.NewVariable()
	res.Variable = l.KnowledgeBase.WorkingMemory.AddVariable(tv)
	res.ArrayMapSelector = l.newResourceArrayMapSelector(r)
	res.SetGrlText(fmt.Sprintf(`this.%s["%s"]`, t, r))
	return res
}

func (l *ecaruleParserListener) newExtAssignVariable(r string) *ast.Variable {
	ext := ast.NewVariable()
	ext.Name = "ext"
	t := ast.NewVariable()
	t.Name = "Void"
	t.Variable = l.KnowledgeBase.WorkingMemory.AddVariable(ext)
	res := ast.NewVariable()
	res.Variable = l.KnowledgeBase.WorkingMemory.AddVariable(t)
	res.ArrayMapSelector = l.newResourceArrayMapSelector(r)
	res.SetGrlText(fmt.Sprintf(`ext.Void["%s"]`, r))
	return res
}

func (l *ecaruleParserListener) newResourceArrayMapSelector(r string) *ast.ArrayMapSelector {
	val := reflect.ValueOf(r)
	c := ast.NewConstant()
	c.Value = val
	c.SetGrlText(fmt.Sprintf(`"%s"`, r))
	a := ast.NewExpressionAtom()
	a.Constant = c
	e := ast.NewExpression()
	e.ExpressionAtom = l.KnowledgeBase.WorkingMemory.AddExpressionAtom(a)
	res := ast.NewArrayMapSelector()
	res.Expression = l.KnowledgeBase.WorkingMemory.AddExpression(e)
	return res
}
