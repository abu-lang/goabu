package parser

import (
	antlr_parser "steel-lang/parser/antlr"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func TokenStream(s string) *antlr.CommonTokenStream {
	lexer := antlr_parser.NewEcaruleLexer(antlr.NewInputStream(s))
	return antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
}
