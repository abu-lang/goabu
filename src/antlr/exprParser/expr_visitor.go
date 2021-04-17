// Code generated from /home/michelep/go/src/mSteelProto/antlr/expr.g4 by ANTLR 4.9.1. DO NOT EDIT.

package exprParser // expr
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by exprParser.
type exprVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by exprParser#exp.
	VisitExp(ctx *ExpContext) interface{}

	// Visit a parse tree produced by exprParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by exprParser#val.
	VisitVal(ctx *ValContext) interface{}

	PevalExp(ctx *ExpContext) string
}
