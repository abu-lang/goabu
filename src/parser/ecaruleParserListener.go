package parser

import (
	"errors"
	"fmt"
	"reflect"
	"steel-lang/datastructure"
	antlr_parser "steel-lang/parser/antlr"
	"strings"

	"github.com/hyperjumptech/grule-rule-engine/antlr"
	"github.com/hyperjumptech/grule-rule-engine/antlr/parser/grulev3"
	"github.com/hyperjumptech/grule-rule-engine/ast"
)

const (
	condUnknown = iota
	condLocal
	condExternal
)

type EcaruleParserListener struct {
	*antlr.GruleV3ParserListener
	Rule              *datastructure.Rule
	types             map[string]string
	allowExt          bool
	allowLocalActions bool
	condition         int
}

func NewEcaruleParserListener(types map[string]string, workingMemory *ast.WorkingMemory, ecb func(e error)) *EcaruleParserListener {
	kb := &ast.KnowledgeBase{
		Name:          "dummy1",
		Version:       "0.0.1",
		RuleEntries:   make(map[string]*ast.RuleEntry),
		WorkingMemory: workingMemory,
	}
	res := &EcaruleParserListener{
		GruleV3ParserListener: antlr.NewGruleV3ParserListener(kb, ecb),
		types:                 types,
		allowLocalActions:     true,
		Rule:                  &datastructure.Rule{},
	}
	res.Stack.Push(res.Rule)
	return res
}

func (l *EcaruleParserListener) parseError(err error) {
	l.StopParse = true
	l.ErrorCallback(err)
}

// EnterPrule is called when production prule is entered.
func (l *EcaruleParserListener) EnterPrule(ctx *antlr_parser.PruleContext) {
	if l.StopParse {
		return
	}
	l.Rule.Name = ctx.SIMPLENAME(0).GetText()
}

// ExitPrule is called when production prule is exited.
func (l *EcaruleParserListener) ExitPrule(ctx *antlr_parser.PruleContext) {
	if l.StopParse {
		return
	}
	l.Stack.Pop()
}

// EnterEvents is called when production events is entered.
func (l *EcaruleParserListener) EnterEvents(ctx *antlr_parser.EventsContext) {
	if l.StopParse {
		return
	}
	rule, ok := l.Stack.Peek().(*datastructure.Rule)
	if !ok {
		l.parseError(errors.New("syntax error"))
		return
	}
	for i := 0; ctx.SIMPLENAME(i) != nil; i++ {
		rule.Events = append(rule.Events, ctx.SIMPLENAME(i).GetText())
	}
}

// ExitEvents is called when production events is exited.
func (l *EcaruleParserListener) ExitEvents(ctx *antlr_parser.EventsContext) {}

// EnterTask is called when production task is entered.
func (l *EcaruleParserListener) EnterTask(ctx *antlr_parser.TaskContext) {
	if l.StopParse {
		return
	}
	rule, ok := l.Stack.Peek().(*datastructure.Rule)
	if !ok {
		l.parseError(errors.New("syntax error"))
		return
	}
	mode := "for"
	if ctx.ALL() != nil {
		mode = mode + " all"
	}
	if ctx.SOME() != nil {
		mode = mode + " some"
	}
	l.allowExt = mode != "for"
	rule.Task.Mode = mode
	l.Stack.Push(&rule.Task)
}

// ExitTask is called when production task is exited.
func (l *EcaruleParserListener) ExitTask(ctx *antlr_parser.TaskContext) {
	if l.StopParse {
		return
	}
	l.condition = condUnknown
	l.allowLocalActions = true
	l.Stack.Pop()
	t, ok := l.Stack.Peek().(*datastructure.Task)
	if ok {
		l.allowExt = t.Mode != "for"
	} else {
		l.allowExt = false
	}
}

// EnterActions is called when production actions is entered.
func (l *EcaruleParserListener) EnterActions(ctx *antlr_parser.ActionsContext) {}

// ExitActions is called when production actions is exited.
func (l *EcaruleParserListener) ExitActions(ctx *antlr_parser.ActionsContext) {}

// EnterTailActions is called when production tailActions is entered.
func (l *EcaruleParserListener) EnterTailActions(ctx *antlr_parser.TailActionsContext) {}

// ExitTailActions is called when production tailActions is exited.
func (l *EcaruleParserListener) ExitTailActions(ctx *antlr_parser.TailActionsContext) {}

// EnterMaybeActions is called when production maybeActions is entered.
func (l *EcaruleParserListener) EnterMaybeActions(ctx *antlr_parser.MaybeActionsContext) {}

// ExitMaybeActions is called when production maybeActions is exited.
func (l *EcaruleParserListener) ExitMaybeActions(ctx *antlr_parser.MaybeActionsContext) {}

func (l *EcaruleParserListener) EnterAssignment(ctx *grulev3.AssignmentContext) {
	switch l.condition {
	case condUnknown:
		l.condition = condLocal
		l.allowLocalActions = true
	case condExternal:
		l.allowLocalActions = false
	}

	l.GruleV3ParserListener.EnterAssignment(ctx)
}

func (l *EcaruleParserListener) ExitAssignment(ctx *grulev3.AssignmentContext) {
	e, ok := l.Stack.Peek().(*ast.Assignment)
	if ok {
		text := e.GetGrlText()
		if !strings.HasPrefix(text, "ext.") {
			if !l.allowLocalActions {
				l.parseError(fmt.Errorf("local action %s is not allowed because the condition contains external resources", text))
			}
		}
		e.Variable = l.newTypedVariable(text, e.Variable)
	}

	l.GruleV3ParserListener.ExitAssignment(ctx)
}

func (l *EcaruleParserListener) ExitExpressionAtom(ctx *grulev3.ExpressionAtomContext) {
	e, ok := l.Stack.Peek().(*ast.ExpressionAtom)
	if ok {
		text := e.GetGrlText()
		switch {
		case strings.HasPrefix(text, "ext."):
			if l.condition == condUnknown {
				l.condition = condExternal
			}
			fallthrough
		case strings.HasPrefix(text, "this."):
			e.Variable = l.newTypedVariable(text, e.Variable)
		}
	}

	l.GruleV3ParserListener.ExitExpressionAtom(ctx)
}

func (l *EcaruleParserListener) newTypedVariable(text string, v *ast.Variable) *ast.Variable {
	switch {
	case strings.HasPrefix(text, "ext."):
		if v == nil {
			l.parseError(fmt.Errorf("expected external resource got %s", text))
			break
		}
		if !l.allowExt {
			l.parseError(fmt.Errorf("external variable %s is not allowed in this context", v.GetGrlText()))
			break
		}
		r := v.Name
		return l.KnowledgeBase.WorkingMemory.AddVariable(l.newExtAssignVariable(r))
	default:
		if v == nil {
			l.parseError(fmt.Errorf("expected local resource got %s", text))
			break
		}
		r := v.Name
		t, present := l.types[r]
		if !present {
			l.parseError(fmt.Errorf("could not determine the type of %s", r))
			break
		}
		return l.KnowledgeBase.WorkingMemory.AddVariable(l.newThisAssignVariable(r, t))
	}
	return v
}

func (l *EcaruleParserListener) newThisAssignVariable(r string, t string) *ast.Variable {
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

func (l *EcaruleParserListener) newExtAssignVariable(r string) *ast.Variable {
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

func (l *EcaruleParserListener) newResourceArrayMapSelector(r string) *ast.ArrayMapSelector {
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
