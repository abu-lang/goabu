// Code generated from /home/michelep/go/src/mSteelProto/antlr/expr.g4 by ANTLR 4.9.1. DO NOT EDIT.

package exprParser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 29, 209,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 144,
	10, 23, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 150, 10, 24, 3, 25, 3, 25, 3,
	25, 5, 25, 155, 10, 25, 3, 25, 3, 25, 6, 25, 159, 10, 25, 13, 25, 14, 25,
	160, 3, 26, 3, 26, 7, 26, 165, 10, 26, 12, 26, 14, 26, 168, 11, 26, 3,
	27, 3, 27, 5, 27, 172, 10, 27, 3, 28, 3, 28, 3, 29, 3, 29, 7, 29, 178,
	10, 29, 12, 29, 14, 29, 181, 11, 29, 3, 29, 3, 29, 3, 30, 3, 30, 5, 30,
	187, 10, 30, 3, 31, 3, 31, 3, 31, 3, 32, 6, 32, 193, 10, 32, 13, 32, 14,
	32, 194, 3, 32, 7, 32, 198, 10, 32, 12, 32, 14, 32, 201, 11, 32, 3, 33,
	6, 33, 204, 10, 33, 13, 33, 14, 33, 205, 3, 33, 3, 33, 2, 2, 34, 3, 3,
	5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22,
	43, 23, 45, 24, 47, 25, 49, 26, 51, 2, 53, 2, 55, 2, 57, 27, 59, 2, 61,
	2, 63, 28, 65, 29, 3, 2, 8, 3, 2, 51, 59, 4, 2, 36, 36, 94, 94, 10, 2,
	36, 36, 41, 41, 94, 94, 100, 100, 104, 104, 112, 112, 116, 116, 118, 118,
	4, 2, 67, 92, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 5, 2, 11,
	12, 15, 15, 34, 34, 2, 215, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15,
	3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2,
	23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2,
	2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2,
	2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2,
	2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 63, 3,
	2, 2, 2, 2, 65, 3, 2, 2, 2, 3, 67, 3, 2, 2, 2, 5, 73, 3, 2, 2, 2, 7, 78,
	3, 2, 2, 2, 9, 83, 3, 2, 2, 2, 11, 87, 3, 2, 2, 2, 13, 90, 3, 2, 2, 2,
	15, 94, 3, 2, 2, 2, 17, 97, 3, 2, 2, 2, 19, 101, 3, 2, 2, 2, 21, 103, 3,
	2, 2, 2, 23, 106, 3, 2, 2, 2, 25, 108, 3, 2, 2, 2, 27, 111, 3, 2, 2, 2,
	29, 113, 3, 2, 2, 2, 31, 115, 3, 2, 2, 2, 33, 117, 3, 2, 2, 2, 35, 119,
	3, 2, 2, 2, 37, 121, 3, 2, 2, 2, 39, 128, 3, 2, 2, 2, 41, 130, 3, 2, 2,
	2, 43, 132, 3, 2, 2, 2, 45, 143, 3, 2, 2, 2, 47, 149, 3, 2, 2, 2, 49, 154,
	3, 2, 2, 2, 51, 162, 3, 2, 2, 2, 53, 171, 3, 2, 2, 2, 55, 173, 3, 2, 2,
	2, 57, 175, 3, 2, 2, 2, 59, 186, 3, 2, 2, 2, 61, 188, 3, 2, 2, 2, 63, 192,
	3, 2, 2, 2, 65, 203, 3, 2, 2, 2, 67, 68, 7, 118, 2, 2, 68, 69, 7, 106,
	2, 2, 69, 70, 7, 107, 2, 2, 70, 71, 7, 117, 2, 2, 71, 72, 7, 48, 2, 2,
	72, 4, 3, 2, 2, 2, 73, 74, 7, 103, 2, 2, 74, 75, 7, 122, 2, 2, 75, 76,
	7, 118, 2, 2, 76, 77, 7, 48, 2, 2, 77, 6, 3, 2, 2, 2, 78, 79, 7, 112, 2,
	2, 79, 80, 7, 119, 2, 2, 80, 81, 7, 110, 2, 2, 81, 82, 7, 110, 2, 2, 82,
	8, 3, 2, 2, 2, 83, 84, 7, 99, 2, 2, 84, 85, 7, 112, 2, 2, 85, 86, 7, 102,
	2, 2, 86, 10, 3, 2, 2, 2, 87, 88, 7, 113, 2, 2, 88, 89, 7, 116, 2, 2, 89,
	12, 3, 2, 2, 2, 90, 91, 7, 112, 2, 2, 91, 92, 7, 113, 2, 2, 92, 93, 7,
	118, 2, 2, 93, 14, 3, 2, 2, 2, 94, 95, 7, 63, 2, 2, 95, 96, 7, 63, 2, 2,
	96, 16, 3, 2, 2, 2, 97, 98, 7, 63, 2, 2, 98, 99, 7, 49, 2, 2, 99, 100,
	7, 63, 2, 2, 100, 18, 3, 2, 2, 2, 101, 102, 7, 62, 2, 2, 102, 20, 3, 2,
	2, 2, 103, 104, 7, 62, 2, 2, 104, 105, 7, 63, 2, 2, 105, 22, 3, 2, 2, 2,
	106, 107, 7, 64, 2, 2, 107, 24, 3, 2, 2, 2, 108, 109, 7, 64, 2, 2, 109,
	110, 7, 63, 2, 2, 110, 26, 3, 2, 2, 2, 111, 112, 7, 63, 2, 2, 112, 28,
	3, 2, 2, 2, 113, 114, 7, 45, 2, 2, 114, 30, 3, 2, 2, 2, 115, 116, 7, 47,
	2, 2, 116, 32, 3, 2, 2, 2, 117, 118, 7, 49, 2, 2, 118, 34, 3, 2, 2, 2,
	119, 120, 7, 44, 2, 2, 120, 36, 3, 2, 2, 2, 121, 122, 7, 101, 2, 2, 122,
	123, 7, 113, 2, 2, 123, 124, 7, 112, 2, 2, 124, 125, 7, 101, 2, 2, 125,
	126, 7, 99, 2, 2, 126, 127, 7, 118, 2, 2, 127, 38, 3, 2, 2, 2, 128, 129,
	7, 42, 2, 2, 129, 40, 3, 2, 2, 2, 130, 131, 7, 43, 2, 2, 131, 42, 3, 2,
	2, 2, 132, 133, 7, 46, 2, 2, 133, 44, 3, 2, 2, 2, 134, 135, 7, 118, 2,
	2, 135, 136, 7, 116, 2, 2, 136, 137, 7, 119, 2, 2, 137, 144, 7, 103, 2,
	2, 138, 139, 7, 104, 2, 2, 139, 140, 7, 99, 2, 2, 140, 141, 7, 110, 2,
	2, 141, 142, 7, 117, 2, 2, 142, 144, 7, 103, 2, 2, 143, 134, 3, 2, 2, 2,
	143, 138, 3, 2, 2, 2, 144, 46, 3, 2, 2, 2, 145, 150, 7, 50, 2, 2, 146,
	150, 5, 51, 26, 2, 147, 148, 7, 47, 2, 2, 148, 150, 5, 51, 26, 2, 149,
	145, 3, 2, 2, 2, 149, 146, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 150, 48, 3,
	2, 2, 2, 151, 155, 5, 47, 24, 2, 152, 153, 7, 47, 2, 2, 153, 155, 7, 50,
	2, 2, 154, 151, 3, 2, 2, 2, 154, 152, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2,
	156, 158, 7, 48, 2, 2, 157, 159, 5, 53, 27, 2, 158, 157, 3, 2, 2, 2, 159,
	160, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 50, 3,
	2, 2, 2, 162, 166, 5, 55, 28, 2, 163, 165, 5, 53, 27, 2, 164, 163, 3, 2,
	2, 2, 165, 168, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2,
	167, 52, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 169, 172, 7, 50, 2, 2, 170,
	172, 5, 55, 28, 2, 171, 169, 3, 2, 2, 2, 171, 170, 3, 2, 2, 2, 172, 54,
	3, 2, 2, 2, 173, 174, 9, 2, 2, 2, 174, 56, 3, 2, 2, 2, 175, 179, 7, 36,
	2, 2, 176, 178, 5, 59, 30, 2, 177, 176, 3, 2, 2, 2, 178, 181, 3, 2, 2,
	2, 179, 177, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 182, 3, 2, 2, 2, 181,
	179, 3, 2, 2, 2, 182, 183, 7, 36, 2, 2, 183, 58, 3, 2, 2, 2, 184, 187,
	10, 3, 2, 2, 185, 187, 5, 61, 31, 2, 186, 184, 3, 2, 2, 2, 186, 185, 3,
	2, 2, 2, 187, 60, 3, 2, 2, 2, 188, 189, 7, 94, 2, 2, 189, 190, 9, 4, 2,
	2, 190, 62, 3, 2, 2, 2, 191, 193, 9, 5, 2, 2, 192, 191, 3, 2, 2, 2, 193,
	194, 3, 2, 2, 2, 194, 192, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 199,
	3, 2, 2, 2, 196, 198, 9, 6, 2, 2, 197, 196, 3, 2, 2, 2, 198, 201, 3, 2,
	2, 2, 199, 197, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200, 64, 3, 2, 2, 2,
	201, 199, 3, 2, 2, 2, 202, 204, 9, 7, 2, 2, 203, 202, 3, 2, 2, 2, 204,
	205, 3, 2, 2, 2, 205, 203, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 207,
	3, 2, 2, 2, 207, 208, 8, 33, 2, 2, 208, 66, 3, 2, 2, 2, 14, 2, 143, 149,
	154, 160, 166, 171, 179, 186, 194, 199, 205, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'this.'", "'ext.'", "'null'", "'and'", "'or'", "'not'", "'=='", "'=/='",
	"'<'", "'<='", "'>'", "'>='", "'='", "'+'", "'-'", "'/'", "'*'", "'concat'",
	"'('", "')'", "','",
}

var lexerSymbolicNames = []string{
	"", "THIS", "EXT", "UNDEF", "AND", "OR", "NOT", "EQ", "NEQ", "LT", "LEQ",
	"GT", "GEQ", "ASSIGN", "PLUS", "MINUS", "DIV", "MUL", "CONCAT", "ROUNDLEFT",
	"ROUNDRIGHT", "COMMA", "BOOL", "INT", "DEC", "STR", "ID", "WS",
}

var lexerRuleNames = []string{
	"THIS", "EXT", "UNDEF", "AND", "OR", "NOT", "EQ", "NEQ", "LT", "LEQ", "GT",
	"GEQ", "ASSIGN", "PLUS", "MINUS", "DIV", "MUL", "CONCAT", "ROUNDLEFT",
	"ROUNDRIGHT", "COMMA", "BOOL", "INT", "DEC", "POS", "DIGIT", "POSDIGIT",
	"STR", "STRCHR", "ESC", "ID", "WS",
}

type exprLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewexprLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *exprLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewexprLexer(input antlr.CharStream) *exprLexer {
	l := new(exprLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "expr.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// exprLexer tokens.
const (
	exprLexerTHIS       = 1
	exprLexerEXT        = 2
	exprLexerUNDEF      = 3
	exprLexerAND        = 4
	exprLexerOR         = 5
	exprLexerNOT        = 6
	exprLexerEQ         = 7
	exprLexerNEQ        = 8
	exprLexerLT         = 9
	exprLexerLEQ        = 10
	exprLexerGT         = 11
	exprLexerGEQ        = 12
	exprLexerASSIGN     = 13
	exprLexerPLUS       = 14
	exprLexerMINUS      = 15
	exprLexerDIV        = 16
	exprLexerMUL        = 17
	exprLexerCONCAT     = 18
	exprLexerROUNDLEFT  = 19
	exprLexerROUNDRIGHT = 20
	exprLexerCOMMA      = 21
	exprLexerBOOL       = 22
	exprLexerINT        = 23
	exprLexerDEC        = 24
	exprLexerSTR        = 25
	exprLexerID         = 26
	exprLexerWS         = 27
)
