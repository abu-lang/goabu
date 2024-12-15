// Copyright 2023 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	antlr_parser "github.com/abu-lang/goabu/parser/internal/antlr"
	"github.com/antlr4-go/antlr/v4"
)

// ValidateIdentifiers returns a boolean value for each argument. If the value is true
// then the argument can be used as an identifier in GoAbU rules.
func ValidateIdentifiers(ids ...string) []bool {
	res := make([]bool, 0, len(ids))
	lexer := antlr_parser.NewEcaruleLexer(antlr.NewInputStream(""))
	lexer.RemoveErrorListeners()
	for _, n := range ids {
		if n != "this" && n != "ext" {
			lexer.SetInputStream(antlr.NewInputStream(n))
			token := lexer.NextToken()
			if token.GetLine() == 1 && token.GetColumn() == 0 &&
				lexer.GetCharIndex() == len(n) &&
				antlr_parser.EcaruleLexerSIMPLENAME == token.GetTokenType() {
				res = append(res, true)
				continue
			}
		}
		res = append(res, false)
	}
	return res
}
