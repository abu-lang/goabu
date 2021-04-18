package exprParser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"steel-lang/misc"
)

type expIntp struct {
	exprStr string
	mem    map[string]interface{}
	tree    IExpContext
}

func NewexpIntp(expr string, memory map[string]interface{}) expIntp {
	// Setup the input string Expression
	input := antlr.NewInputStream(expr)
	// Create the Lexer
	lexer := NewexprLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	parser := NewexprParser(stream)
	// Initialize the Parser with the root Expression
	t := parser.Exp()
	return expIntp{exprStr: expr,mem:misc.CopyMap(memory),tree:t}
}

func (v *expIntp) RunexpIntp() interface{} {
	// Create and initialize a Visitor for the Parser
	visitor := NewexprVisitorImpl(v.mem)
	// Visit the Expression
	return visitor.VisitExp(v.tree.(*ExpContext))
}

func (v *expIntp) PevalExp() string {
	// Create and initialize a Visitor for the Parser
	visitor := NewexprVisitorImpl(v.mem)
	// Partially evaluate the Expression
	return visitor.PevalExp(v.tree.(*ExpContext))
}
