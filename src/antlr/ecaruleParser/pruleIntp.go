package ecaruleParser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"steel-lang/datastructure"
)

type pruleIntp struct {
	pruleStr string
	rules   *datastructure.Rule
	tree    IPruleContext
}

func NewpruleIntp(prule string) pruleIntp {
	// Setup the input string Rule
	input := antlr.NewInputStream(prule)
	// Create the Lexer
	lexer := NewecaruleLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	parser := NewecaruleParser(stream)
	// Initialize the Parser with the root Rule
	t := parser.Prule()
	return pruleIntp{pruleStr: prule, tree: t, rules: new(datastructure.Rule)}
}

func (m *pruleIntp) RunpruleIntp() (datastructure.Rule) {
	// Create and initialize a Visitor for the Parser
	visitor := NewecaruleVisitorImpl(m.rules)
	// Visit the Rule
	visitor.VisitPrule(m.tree.(*PruleContext))
	return *m.rules
}
