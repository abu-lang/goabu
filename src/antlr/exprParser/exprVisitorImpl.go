package exprParser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"reflect"
	"strconv"
)

type exprVisitorImpl struct {
	antlr.ParseTreeVisitor
	mem map[string]interface{}
}

func NewexprVisitorImpl(memory map[string]interface{}) exprVisitor {
	return &exprVisitorImpl{ParseTreeVisitor:&BaseexprVisitor{},mem:memory}
}

func (v *exprVisitorImpl) VisitExp(ctx *ExpContext) interface{} {
	// Case 'ROUNDLEFT exp ROUNDRIGHT'
	if ctx.ROUNDLEFT() != nil && ctx.CONCAT() == nil {
		return v.VisitExp(ctx.Exp(0).(*ExpContext))
	}
	// Case 'val'
	if ctx.Val() != nil {
		return v.VisitVal(ctx.Val().(*ValContext))
	}
	// Case 'NOT exp'
	if ctx.NOT() != nil {
		return !v.VisitExp(ctx.Exp(0).(*ExpContext)).(bool)
	}
	// Case 'exp op exp'
	if ctx.op != nil {
		tree1 := v.VisitExp(ctx.Exp(0).(*ExpContext))
		tree2 := v.VisitExp(ctx.Exp(1).(*ExpContext))
		switch ctx.op.GetTokenType() {
		case exprParserMUL:
			if reflect.TypeOf(tree1).String() == "int64" {
				return tree1.(int64) * tree2.(int64)
			} else {
				return tree1.(float64) * tree2.(float64)
			}
		case exprParserDIV:
			if reflect.TypeOf(tree1).String() == "int64" {
				return tree1.(int64) / tree2.(int64)
			} else {
				return tree1.(float64) / tree2.(float64)
			}
		case exprParserPLUS:
			if reflect.TypeOf(tree1).String() == "int64" {
				return tree1.(int64) + tree2.(int64)
			} else {
				return tree1.(float64) + tree2.(float64)
			}
		case exprParserMINUS:
			if reflect.TypeOf(tree1).String() == "int64" {
				return tree1.(int64) - tree2.(int64)
			} else {
				return tree1.(float64) - tree2.(float64)
			}
		case exprParserLT:
			if reflect.TypeOf(tree1).String() == "int64" {
				return tree1.(int64) < tree2.(int64)
			} else {
				return tree1.(float64) < tree2.(float64)
			}
		case exprParserLEQ:
			if reflect.TypeOf(tree1).String() == "int64" {
				return tree1.(int64) <= tree2.(int64)
			} else {
				return tree1.(float64) <= tree2.(float64)
			}
		case exprParserGT:
			if reflect.TypeOf(tree1).String() == "int64" {
				return tree1.(int64) > tree2.(int64)
			} else {
				return tree1.(float64) > tree2.(float64)
			}
		case exprParserGEQ:
			if reflect.TypeOf(tree1).String() == "int64" {
				return tree1.(int64) >= tree2.(int64)
			} else {
				return tree1.(float64) >= tree2.(float64)
			}
		case exprParserEQ:
			return tree1 == tree2
		case exprParserNEQ:
			return tree1 != tree2
		case exprParserAND:
			return tree1.(bool) && tree2.(bool)
		case exprParserOR:
			return tree1.(bool) || tree2.(bool)
		}
	}
	// Case 'CONCAT ROUNDLEFT exp COMMA exp ROUNDRIGHT'
	if ctx.ROUNDLEFT() != nil && ctx.CONCAT() != nil {
		tree1 := v.VisitExp(ctx.Exp(0).(*ExpContext))
		tree2 := v.VisitExp(ctx.Exp(1).(*ExpContext))
		return tree1.(string) + tree2.(string)
	}
	// Case '(THIS)? id' or 'EXT id'
	if ctx.Id() != nil {
		var ret interface{}
		if ctx.EXT() != nil {
			// TODO: implement external lookup
			ret = v.mem[ctx.Id().GetText()]
		} else {
			ret = v.mem[ctx.Id().GetText()]
		}
		var v interface{}
		switch ret.(type) {
		case bool, string:
			v = ret
		case int, int32:
			v = int64(ret.(int))
		case float32:
			v = float64(ret.(float32))
		}
		return v
	}
	return nil // Unreachable
}

// Never used
func (v *exprVisitorImpl) VisitId(ctx *IdContext) interface{} {
	return nil
}

func (v *exprVisitorImpl) VisitVal(ctx *ValContext) interface{} {
	if ctx.BOOL() != nil {
		v , _ := strconv.ParseBool(ctx.GetText())
		return v
	}
	if ctx.INT() != nil {
		v , _ := strconv.ParseInt(ctx.GetText(),10,64)
		return v
	}
	if ctx.DEC() != nil {
		v , _ := strconv.ParseFloat(ctx.GetText(),64)
		return v
	}
	if ctx.STR() != nil {
		return ctx.GetText()[1:len(ctx.GetText())-1]
	}
	return nil // Unreachable
}

func (v *exprVisitorImpl) PevalExp(ctx *ExpContext) string {
	// Case 'ROUNDLEFT exp ROUNDRIGHT'
	if ctx.ROUNDLEFT() != nil && ctx.CONCAT() == nil {
		return ctx.ROUNDLEFT().GetText() + v.PevalExp(ctx.Exp(0).(*ExpContext)) + ctx.ROUNDRIGHT().GetText()
	}
	// Case 'val'
	if ctx.Val() != nil {
		return ctx.Val().GetText()
	}
	// Case 'NOT exp'
	if ctx.NOT() != nil {
		return ctx.NOT().GetText() + v.PevalExp(ctx.Exp(0).(*ExpContext))
	}
	// Case 'exp op exp'
	if ctx.op != nil {
		return v.PevalExp(ctx.Exp(0).(*ExpContext)) + ctx.op.GetText() + v.PevalExp(ctx.Exp(1).(*ExpContext))
	}
	// Case 'CONCAT ROUNDLEFT exp COMMA exp ROUNDRIGHT'
	if ctx.ROUNDLEFT() != nil && ctx.CONCAT() != nil {
		return ctx.CONCAT().GetText() + ctx.ROUNDLEFT().GetText() + v.PevalExp(ctx.Exp(0).(*ExpContext)) + ctx.COMMA().GetText() + v.PevalExp(ctx.Exp(1).(*ExpContext)) + ctx.ROUNDRIGHT().GetText()
	}
	// Case '(THIS)? id' or 'EXT id'
	if ctx.Id() != nil {
		if ctx.EXT() != nil {
			// TODO: implement external lookup
			return fmt.Sprintf("%v", v.mem[ctx.Id().GetText()])
		} else {
			return fmt.Sprintf("%v", v.mem[ctx.Id().GetText()])
		}
	}
	return "" // Unreachable
}
