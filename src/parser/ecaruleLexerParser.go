package parser

import (
	antlr_parser "steel-lang/parser/antlr"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type EcaruleLexerParser struct {
	Lexer  *antlr_parser.EcaruleLexer
	Parser *antlr_parser.EcaruleParser
}

func NewEcaruleLexerParser() *EcaruleLexerParser {
	res := &EcaruleLexerParser{}
	res.Lexer = antlr_parser.NewEcaruleLexer(antlr.NewInputStream(""))
	res.Parser = antlr_parser.NewEcaruleParser(antlr.NewCommonTokenStream(res.Lexer, antlr.TokenDefaultChannel))
	res.Parser.BuildParseTrees = true
	return res
}

func (lp *EcaruleLexerParser) Reset(input string) {
	lp.Lexer.SetInputStream(antlr.NewInputStream(input))
	lp.Parser.SetInputStream(antlr.NewCommonTokenStream(lp.Lexer, antlr.TokenDefaultChannel))
}
