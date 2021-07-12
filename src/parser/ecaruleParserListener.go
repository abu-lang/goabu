package parser

import (
	"fmt"
	"reflect"
	"steel-lang/datastructure"
	antlr_parser "steel-lang/parser/antlr"
	"strings"

	"github.com/hyperjumptech/grule-rule-engine/antlr"
	"github.com/hyperjumptech/grule-rule-engine/antlr/parser/grulev3"
	"github.com/hyperjumptech/grule-rule-engine/ast"
)

type ecaruleParserListener struct {
	*antlr.GruleV3ParserListener
	types map[string]string
	Rule  *datastructure.Rule
}

func NewEcaruleParserListener(types map[string]string, workingMemory *ast.WorkingMemory, ecb func(e error)) *ecaruleParserListener {
	kb := &ast.KnowledgeBase{
		Name:          "dummy1",
		Version:       "0.0.1",
		RuleEntries:   make(map[string]*ast.RuleEntry),
		WorkingMemory: workingMemory,
	}
	return &ecaruleParserListener{
		GruleV3ParserListener: antlr.NewGruleV3ParserListener(kb, ecb),
		types:                 types,
		Rule:                  &datastructure.Rule{},
	}
}

// EnterPrule is called when production prule is entered.
func (l *ecaruleParserListener) EnterPrule(ctx *antlr_parser.PruleContext) {
	l.Rule.Name = ctx.SIMPLENAME(0).GetText()
}

// ExitPrule is called when production prule is exited.
func (l *ecaruleParserListener) ExitPrule(ctx *antlr_parser.PruleContext) {}

// EnterEvt is called when production evt is entered.
func (l *ecaruleParserListener) EnterEvt(ctx *antlr_parser.EvtContext) {
	for i := 0; ctx.SIMPLENAME(i) != nil; i++ {
		l.Rule.Events = append(l.Rule.Events, ctx.SIMPLENAME(i).GetText())
	}
}

// ExitEvt is called when production evt is exited.
func (l *ecaruleParserListener) ExitEvt(ctx *antlr_parser.EvtContext) {}

// EnterTask is called when production task is entered.
func (l *ecaruleParserListener) EnterTask(ctx *antlr_parser.TaskContext) {
	mode := "for"
	if ctx.ALL() != nil {
		mode = mode + " all"
	}
	if ctx.SOME() != nil {
		mode = mode + " some"
	}
	l.Rule.Task.Mode = mode
	l.Stack.Push(&l.Rule.Task)
}

// ExitTask is called when production task is exited.
func (l *ecaruleParserListener) ExitTask(ctx *antlr_parser.TaskContext) {
	l.Stack.Pop()
}

// EnterActslist is called when production actslist is entered.
func (l *ecaruleParserListener) EnterActslist(ctx *antlr_parser.ActslistContext) {}

// ExitActslist is called when production actslist is exited.
func (l *ecaruleParserListener) ExitActslist(ctx *antlr_parser.ActslistContext) {}

// EnterAct is called when production act is entered.
func (l *ecaruleParserListener) EnterAct(ctx *antlr_parser.ActContext) {
	assign := ast.NewAssignment()
	assign.IsAssign = true
	assign.SetGrlText(ctx.GetText())
	l.Stack.Push(assign)
}

// ExitAct is called when production act is exited.
func (l *ecaruleParserListener) ExitAct(ctx *antlr_parser.ActContext) {
	inDefault := l.Rule.Task.Mode == ""
	local := inDefault || l.Rule.Task.Mode == "for"
	assign := l.Stack.Pop().(*ast.Assignment)
	dest := ctx.SIMPLENAME().GetText()
	if !local {
		assign.Variable = l.GruleV3ParserListener.KnowledgeBase.WorkingMemory.AddVariable(l.NewExtAssignVariable(dest))
	} else {
		t, present := l.types[dest]
		if !present {
			l.GruleV3ParserListener.StopParse = true
			l.GruleV3ParserListener.ErrorCallback(fmt.Errorf("could not determine the type of %s", dest))
		}
		assign.Variable = l.GruleV3ParserListener.KnowledgeBase.WorkingMemory.AddVariable(l.NewThisAssignVariable(dest, t))
	}
	action := datastructure.Action{
		Resource:   dest,
		Expression: assign,
	}
	if inDefault {
		l.Rule.DefaultActions = append(l.Rule.DefaultActions, action)
	} else {
		l.Rule.Task.Actions = append(l.Rule.Task.Actions, action)
	}
}

func (l *ecaruleParserListener) ExitExpressionAtom(ctx *grulev3.ExpressionAtomContext) {
	e := l.Stack.Peek().(*ast.ExpressionAtom)
	errFunc := func(err error) {
		l.GruleV3ParserListener.StopParse = true
		l.GruleV3ParserListener.ErrorCallback(err)
	}
	switch {
	case strings.HasPrefix(e.GetGrlText(), "ext."):
		if e.Variable == nil {
			defer errFunc(fmt.Errorf("expected external resource got %s", e.GetGrlText()))
			break
		}
		r := e.Variable.Name
		e.Variable = l.GruleV3ParserListener.KnowledgeBase.WorkingMemory.AddVariable(l.NewExtAssignVariable(r))
	case strings.HasPrefix(e.GetGrlText(), "this."):
		if e.Variable == nil {
			defer errFunc(fmt.Errorf("expected local resource got %s", e.GetGrlText()))
			break
		}
		r := e.Variable.Name
		t, present := l.types[r]
		if !present {
			defer errFunc(fmt.Errorf("could not determine the type of %s", r))
			break
		}
		e.Variable = l.GruleV3ParserListener.KnowledgeBase.WorkingMemory.AddVariable(l.NewThisAssignVariable(r, t))
	}
	l.GruleV3ParserListener.ExitExpressionAtom(ctx)
}

func (l *ecaruleParserListener) NewThisAssignVariable(r string, t string) *ast.Variable {
	this := ast.NewVariable()
	this.Name = "this"
	tv := ast.NewVariable()
	tv.Name = t
	tv.Variable = l.GruleV3ParserListener.KnowledgeBase.WorkingMemory.AddVariable(this)
	res := ast.NewVariable()
	res.Variable = l.GruleV3ParserListener.KnowledgeBase.WorkingMemory.AddVariable(tv)
	res.ArrayMapSelector = l.NewResourceArrayMapSelector(r)
	res.SetGrlText(fmt.Sprintf(`this.%s["%s"]`, t, r))
	return res
}

func (l *ecaruleParserListener) NewExtAssignVariable(r string) *ast.Variable {
	ext := ast.NewVariable()
	ext.Name = "ext"
	t := ast.NewVariable()
	t.Name = "Void"
	t.Variable = l.GruleV3ParserListener.KnowledgeBase.WorkingMemory.AddVariable(ext)
	res := ast.NewVariable()
	res.Variable = l.GruleV3ParserListener.KnowledgeBase.WorkingMemory.AddVariable(t)
	res.ArrayMapSelector = l.NewResourceArrayMapSelector(r)
	res.SetGrlText(fmt.Sprintf(`ext.Void["%s"]`, r))
	return res
}

func (l *ecaruleParserListener) NewResourceArrayMapSelector(r string) *ast.ArrayMapSelector {
	val := reflect.ValueOf(r)
	c := ast.NewConstant()
	c.Value = val
	c.SetGrlText(fmt.Sprintf(`"%s"`, r))
	a := ast.NewExpressionAtom()
	a.Constant = c
	e := ast.NewExpression()
	e.ExpressionAtom = l.GruleV3ParserListener.KnowledgeBase.WorkingMemory.AddExpressionAtom(a)
	res := ast.NewArrayMapSelector()
	res.Expression = l.GruleV3ParserListener.KnowledgeBase.WorkingMemory.AddExpression(e)
	return res
}
