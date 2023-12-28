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

	"github.com/hyperjumptech/grule-rule-engine/antlr"
	"github.com/hyperjumptech/grule-rule-engine/antlr/parser/grulev3"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// processing will contain the events and the tasks of the rule currently being processed.
type processing struct {
	events      []string
	localTasks  []ecarule.LocalTask
	remoteTasks []ecarule.RemoteTask
}

type ecaruleParserListener struct {
	*antlr.GruleV3ParserListener
	Rules        []ecarule.Rule
	types        map[string]string
	allowExt     bool
	inAssignLeft bool
	processing
}

func newEcaruleParserListener(types map[string]string, workingMemory *ast.WorkingMemory, errorReporter *pkg.GruleErrorReporter) *ecaruleParserListener {
	kb := &ast.KnowledgeBase{
		WorkingMemory: workingMemory,
	}
	res := &ecaruleParserListener{
		GruleV3ParserListener: antlr.NewGruleV3ParserListener(kb, errorReporter),
		types:                 types,
	}
	return res
}

// reset brings back the listener to a clean state so that it can be used on a differet tree.
func (l *ecaruleParserListener) reset() {
	l.Rules = nil
	l.PreviousNode = l.PreviousNode[:0]
	for l.Stack.Len() > 0 {
		l.Stack.Pop()
	}
	l.StopParse = false
	l.ErrorCallback.Errors = make([]error, 0)
}

func (l *ecaruleParserListener) parseError(err error) {
	l.StopParse = true
	l.ErrorCallback.AddError(err)
}

// EnterPrule is called when production prules is entered.
func (l *ecaruleParserListener) EnterPrules(ctx *antlr_parser.PrulesContext) {}

// ExitPrule is called when production prules is exited.
func (l *ecaruleParserListener) ExitPrules(ctx *antlr_parser.PrulesContext) {}

// EnterPrule is called when production prule is entered.
func (l *ecaruleParserListener) EnterPrule(ctx *antlr_parser.PruleContext) {
	if l.StopParse {
		return
	}
	l.localTasks = make([]ecarule.LocalTask, 0)
	l.remoteTasks = make([]ecarule.RemoteTask, 0)
}

// ExitPrule is called when production prule is exited.
func (l *ecaruleParserListener) ExitPrule(ctx *antlr_parser.PruleContext) {
	if l.StopParse {
		return
	}
	l.Rules = append(l.Rules, ecarule.Rule{
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
func (l *ecaruleParserListener) EnterEvents(ctx *antlr_parser.EventsContext) {
	if l.StopParse {
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

// ExitEvents is called when production events is exited.
func (l *ecaruleParserListener) ExitEvents(ctx *antlr_parser.EventsContext) {}

// EnterDefaultActions is called when production defaultActions is entered.
func (l *ecaruleParserListener) EnterDefaultActions(ctx *antlr_parser.DefaultActionsContext) {
	if l.StopParse {
		return
	}
	cond, err := l.newBooleanLiteralExpression(true)
	if err != nil {
		l.parseError(errors.New("error during default actions parsing"))
		return
	}
	l.localTasks = append(l.localTasks, ecarule.LocalTask{Condition: cond})
	l.Stack.Push(expressionReceiver{&l.localTasks[len(l.localTasks)-1]})
}

// ExitDefaultActions is called when production defaultActions is exited.
func (l *ecaruleParserListener) ExitDefaultActions(ctx *antlr_parser.DefaultActionsContext) {
	if l.StopParse {
		return
	}
	l.Stack.Pop()
}

// EnterTask is called when production task is entered.
func (l *ecaruleParserListener) EnterTask(ctx *antlr_parser.TaskContext) {
	if l.StopParse {
		return
	}
	if ctx.ALL() != nil {
		l.allowExt = true
		l.remoteTasks = append(l.remoteTasks, ecarule.RemoteTask{})
		l.Stack.Push(tmpExpressionReceiver{&l.remoteTasks[len(l.remoteTasks)-1]})
	} else {
		l.allowExt = false
		l.localTasks = append(l.localTasks, ecarule.LocalTask{})
		l.Stack.Push(expressionReceiver{&l.localTasks[len(l.localTasks)-1]})
	}
}

// ExitTask is called when production task is exited.
func (l *ecaruleParserListener) ExitTask(ctx *antlr_parser.TaskContext) {
	if l.StopParse {
		return
	}
	l.Stack.Pop()
	l.allowExt = false
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

// newBooleanLiteralExpression creates an [*ast.Expression] encoding the specified boolean value.
func (l *ecaruleParserListener) newBooleanLiteralExpression(b bool) (*ast.Expression, error) {
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
	err = exp.AcceptExpressionAtom(l.KnowledgeBase.WorkingMemory.AddExpressionAtom(atm))
	if err != nil {
		return nil, err
	}
	return l.KnowledgeBase.WorkingMemory.AddExpression(exp), nil
}
