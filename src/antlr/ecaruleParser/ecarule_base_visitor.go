// Code generated from /home/michelep/Gitlab/steel-lang/src/antlr/ecarule.g4 by ANTLR 4.9.1. DO NOT EDIT.

package ecaruleParser // ecarule
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseecaruleVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseecaruleVisitor) VisitPrule(ctx *PruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseecaruleVisitor) VisitEvt(ctx *EvtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseecaruleVisitor) VisitTask(ctx *TaskContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseecaruleVisitor) VisitActslist(ctx *ActslistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseecaruleVisitor) VisitAct(ctx *ActContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseecaruleVisitor) VisitExp(ctx *ExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseecaruleVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseecaruleVisitor) VisitVal(ctx *ValContext) interface{} {
	return v.VisitChildren(ctx)
}
