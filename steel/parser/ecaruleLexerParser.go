package parser

import (
	antlr_parser "steel/parser/antlr"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

type EcaruleLexerParser struct {
	Lexer       *antlr_parser.EcaruleLexer
	Parser      *antlr_parser.EcaruleParser
	errListener *pkg.GruleErrorReporter
}

func NewEcaruleLexerParser() *EcaruleLexerParser {
	res := &EcaruleLexerParser{errListener: &pkg.GruleErrorReporter{Errors: make([]error, 0)}}
	res.Lexer = antlr_parser.NewEcaruleLexer(antlr.NewInputStream(""))
	res.Lexer.RemoveErrorListeners()
	res.Lexer.AddErrorListener(res.errListener)
	res.Parser = antlr_parser.NewEcaruleParser(antlr.NewCommonTokenStream(res.Lexer, antlr.TokenDefaultChannel))
	res.Parser.BuildParseTrees = true
	res.Parser.RemoveErrorListeners()
	res.Parser.AddErrorListener(res.errListener)
	return res
}

func (lp *EcaruleLexerParser) Reset(input string) {
	lp.errListener.Errors = lp.errListener.Errors[:0]
	lp.Lexer.SetInputStream(antlr.NewInputStream(input))
	lp.Parser.SetInputStream(antlr.NewCommonTokenStream(lp.Lexer, antlr.TokenDefaultChannel))
}

func (lp *EcaruleLexerParser) Errors() []error {
	return lp.errListener.Errors
}
