// Code generated from /home/michelep/Gitlab/steel-lang/src/antlr/ecarule.g4 by ANTLR 4.9.1. DO NOT EDIT.

package ecaruleParser // ecarule
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ecaruleParser.
type ecaruleVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ecaruleParser#prule.
	VisitPrule(ctx *PruleContext) interface{}

	// Visit a parse tree produced by ecaruleParser#evt.
	VisitEvt(ctx *EvtContext) interface{}

	// Visit a parse tree produced by ecaruleParser#task.
	VisitTask(ctx *TaskContext) interface{}

	// Visit a parse tree produced by ecaruleParser#actslist.
	VisitActslist(ctx *ActslistContext) interface{}

	// Visit a parse tree produced by ecaruleParser#act.
	VisitAct(ctx *ActContext) interface{}

	// Visit a parse tree produced by ecaruleParser#exp.
	VisitExp(ctx *ExpContext) interface{}

	// Visit a parse tree produced by ecaruleParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by ecaruleParser#val.
	VisitVal(ctx *ValContext) interface{}
}
