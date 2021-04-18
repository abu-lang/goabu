package ecaruleParser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"steel-lang/datastructure"
)

type ecaruleVisitorImpl struct {
	antlr.ParseTreeVisitor
	rule *datastructure.Rule
}

func NewecaruleVisitorImpl(rulep *datastructure.Rule) ecaruleVisitor {
	return &ecaruleVisitorImpl{ParseTreeVisitor: &BaseecaruleVisitor{}, rule: rulep}
}

func (m *ecaruleVisitorImpl) VisitPrule(ctx *PruleContext) interface{} {
	ruleName := ctx.Id(0).GetText()
	evt := m.VisitEvt(ctx.Evt().(*EvtContext))
	/* TODO: ruleset skipped for now
	var ruleset string
	if ctx.IN() != nil {
		ruleset = ctx.Id(1).GetText()
	} else {
		ruleset = ruleName
	}*/
	var defActs []datastructure.Action
	if ctx.DEFAULT() != nil {
		defActs = m.VisitActslist(ctx.Actslist().(*ActslistContext)).([]datastructure.Action)
	} else {
		defActs = nil
	}
	*m.rule = datastructure.Rule{Name: ruleName, Event: evt.([]string), DefaultActions: defActs, Task: m.VisitTask(ctx.Task().(*TaskContext)).(datastructure.Task)}
	return nil
}

func (m *ecaruleVisitorImpl) VisitEvt(ctx *EvtContext) interface{} {
	var evt []string
	for i := 0; ctx.Id(i) != nil; i++ {
		evt = append(evt,ctx.Id(i).GetText())
	}
	return evt
}

func (m *ecaruleVisitorImpl) VisitTask(ctx *TaskContext) interface{} {
	mode := "for"
	if ctx.ALL() != nil {
		mode = mode + " all"
	}
	if ctx.SOME() != nil {
		mode = mode + " some"
	}
	return  datastructure.Task{Mode: mode, Exp: ctx.Exp().GetText(), Actions: m.VisitActslist(ctx.Actslist().(*ActslistContext)).([]datastructure.Action)}
}

func (m *ecaruleVisitorImpl) VisitActslist(ctx *ActslistContext) interface{} {
	var actslist []datastructure.Action
	for i := 0; ctx.Act(i) != nil; i++ {
		actslist = append(actslist, m.VisitAct(ctx.Act(i).(*ActContext)).(datastructure.Action))
	}
	return actslist
}

func (m *ecaruleVisitorImpl) VisitAct(ctx *ActContext) interface{} {
	var sact datastructure.Action
	if ctx.EXT() != nil {
		sact = datastructure.Action{Resource: ctx.Id().GetText(), External: true, Expression: ctx.Exp().GetText()}
	} else {
		sact = datastructure.Action{Resource: ctx.Id().GetText(), External: false, Expression: ctx.Exp().GetText()}
	}
	return sact
}

// Never used
func (m *ecaruleVisitorImpl) VisitExp(ctx *ExpContext) interface{} {
	return nil
}

// Never used
func (m *ecaruleVisitorImpl) VisitId(ctx *IdContext) interface{} {
	return nil
}

// Never used
func (m *ecaruleVisitorImpl) VisitVal(ctx *ValContext) interface{} {
	return nil
}