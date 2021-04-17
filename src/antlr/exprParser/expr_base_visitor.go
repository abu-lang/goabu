// Code generated from /home/michelep/go/src/mSteelProto/antlr/expr.g4 by ANTLR 4.9.1. DO NOT EDIT.

package exprParser // expr
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseexprVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseexprVisitor) VisitExp(ctx *ExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseexprVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseexprVisitor) VisitVal(ctx *ValContext) interface{} {
	return v.VisitChildren(ctx)
}
