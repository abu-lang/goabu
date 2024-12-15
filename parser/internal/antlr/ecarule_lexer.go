// Code generated from EcaruleLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package antlr

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type EcaruleLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var EcaruleLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ecarulelexerLexerInit() {
	staticData := &EcaruleLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.SymbolicNames = []string{
		"", "", "PLUS", "MINUS", "DIV", "MUL", "MOD", "DOT", "SEMICOLON", "LR_BRACE",
		"RR_BRACE", "LR_BRACKET", "RR_BRACKET", "LS_BRACKET", "RS_BRACKET",
		"RULE", "WHEN", "THEN", "AND", "OR", "TRUE", "FALSE", "NIL_LITERAL",
		"NEGATION", "SALIENCE", "EQUALS", "ASSIGN", "PLUS_ASIGN", "MINUS_ASIGN",
		"DIV_ASIGN", "MUL_ASIGN", "GT", "LT", "GTE", "LTE", "NOTEQUALS", "BITAND",
		"BITOR", "SIMPLENAME", "DQUOTA_STRING", "SQUOTA_STRING", "DECIMAL_FLOAT_LIT",
		"DECIMAL_EXPONENT", "HEX_FLOAT_LIT", "HEX_EXPONENT", "DEC_LIT", "HEX_LIT",
		"OCT_LIT", "SPACE", "COMMENT", "LINE_COMMENT", "ON", "DEFAULT", "FOR",
		"ALL", "DO",
	}
	staticData.RuleNames = []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "ISC", "IC",
		"T__0", "PLUS", "MINUS", "DIV", "MUL", "MOD", "DOT", "SEMICOLON", "LR_BRACE",
		"RR_BRACE", "LR_BRACKET", "RR_BRACKET", "LS_BRACKET", "RS_BRACKET",
		"RULE", "WHEN", "THEN", "AND", "OR", "TRUE", "FALSE", "NIL_LITERAL",
		"NEGATION", "SALIENCE", "EQUALS", "ASSIGN", "PLUS_ASIGN", "MINUS_ASIGN",
		"DIV_ASIGN", "MUL_ASIGN", "GT", "LT", "GTE", "LTE", "NOTEQUALS", "BITAND",
		"BITOR", "ON", "DEFAULT", "FOR", "ALL", "DO", "SIMPLENAME", "DQUOTA_STRING",
		"SQUOTA_STRING", "DECIMAL_FLOAT_LIT", "DECIMAL_EXPONENT", "HEX_FLOAT_LIT",
		"HEX_MANTISA", "HEX_EXPONENT", "DEC_LIT", "HEX_LIT", "OCT_LIT", "HEX_DIGITS",
		"DEC_DIGITS", "OCT_DIGITS", "DEC_DIGIT", "OCT_DIGIT", "HEX_DIGIT", "SPACE",
		"COMMENT", "LINE_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 55, 516, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2,
		73, 7, 73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 2, 77, 7, 77, 2, 78,
		7, 78, 2, 79, 7, 79, 2, 80, 7, 80, 2, 81, 7, 81, 2, 82, 7, 82, 2, 83, 7,
		83, 2, 84, 7, 84, 2, 85, 7, 85, 2, 86, 7, 86, 2, 87, 7, 87, 2, 88, 7, 88,
		2, 89, 7, 89, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26,
		1, 26, 1, 27, 1, 27, 3, 27, 238, 8, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1,
		30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35,
		1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1,
		40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43,
		1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1,
		46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48,
		1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1,
		51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52,
		1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 1, 56, 1, 56, 1,
		56, 1, 57, 1, 57, 1, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1, 60,
		1, 61, 1, 61, 1, 61, 1, 62, 1, 62, 1, 62, 1, 63, 1, 63, 1, 64, 1, 64, 1,
		65, 1, 65, 1, 65, 1, 66, 1, 66, 1, 66, 1, 66, 1, 66, 1, 66, 1, 66, 1, 66,
		1, 67, 1, 67, 1, 67, 1, 67, 1, 68, 1, 68, 1, 68, 1, 68, 1, 69, 1, 69, 1,
		69, 1, 70, 1, 70, 5, 70, 373, 8, 70, 10, 70, 12, 70, 376, 9, 70, 1, 71,
		1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 5, 71, 384, 8, 71, 10, 71, 12, 71, 387,
		9, 71, 1, 71, 1, 71, 1, 72, 1, 72, 1, 72, 1, 72, 1, 72, 1, 72, 5, 72, 397,
		8, 72, 10, 72, 12, 72, 400, 9, 72, 1, 72, 1, 72, 1, 73, 1, 73, 1, 73, 1,
		73, 3, 73, 408, 8, 73, 1, 73, 1, 73, 1, 73, 1, 73, 1, 73, 1, 73, 3, 73,
		416, 8, 73, 3, 73, 418, 8, 73, 1, 74, 1, 74, 1, 74, 3, 74, 423, 8, 74,
		1, 74, 1, 74, 1, 75, 1, 75, 1, 75, 1, 75, 1, 75, 1, 76, 1, 76, 1, 76, 3,
		76, 435, 8, 76, 1, 76, 1, 76, 1, 76, 1, 76, 3, 76, 441, 8, 76, 1, 77, 1,
		77, 1, 77, 3, 77, 446, 8, 77, 1, 77, 1, 77, 1, 78, 1, 78, 1, 78, 3, 78,
		453, 8, 78, 3, 78, 455, 8, 78, 1, 79, 1, 79, 1, 79, 1, 79, 1, 80, 1, 80,
		1, 80, 1, 81, 4, 81, 465, 8, 81, 11, 81, 12, 81, 466, 1, 82, 4, 82, 470,
		8, 82, 11, 82, 12, 82, 471, 1, 83, 4, 83, 475, 8, 83, 11, 83, 12, 83, 476,
		1, 84, 1, 84, 1, 85, 1, 85, 1, 86, 1, 86, 1, 87, 4, 87, 486, 8, 87, 11,
		87, 12, 87, 487, 1, 87, 1, 87, 1, 88, 1, 88, 1, 88, 1, 88, 5, 88, 496,
		8, 88, 10, 88, 12, 88, 499, 9, 88, 1, 88, 1, 88, 1, 88, 1, 88, 1, 88, 1,
		89, 1, 89, 1, 89, 1, 89, 5, 89, 510, 8, 89, 10, 89, 12, 89, 513, 9, 89,
		1, 89, 1, 89, 1, 497, 0, 90, 1, 0, 3, 0, 5, 0, 7, 0, 9, 0, 11, 0, 13, 0,
		15, 0, 17, 0, 19, 0, 21, 0, 23, 0, 25, 0, 27, 0, 29, 0, 31, 0, 33, 0, 35,
		0, 37, 0, 39, 0, 41, 0, 43, 0, 45, 0, 47, 0, 49, 0, 51, 0, 53, 0, 55, 0,
		57, 1, 59, 2, 61, 3, 63, 4, 65, 5, 67, 6, 69, 7, 71, 8, 73, 9, 75, 10,
		77, 11, 79, 12, 81, 13, 83, 14, 85, 15, 87, 16, 89, 17, 91, 18, 93, 19,
		95, 20, 97, 21, 99, 22, 101, 23, 103, 24, 105, 25, 107, 26, 109, 27, 111,
		28, 113, 29, 115, 30, 117, 31, 119, 32, 121, 33, 123, 34, 125, 35, 127,
		36, 129, 37, 131, 51, 133, 52, 135, 53, 137, 54, 139, 55, 141, 38, 143,
		39, 145, 40, 147, 41, 149, 42, 151, 43, 153, 0, 155, 44, 157, 45, 159,
		46, 161, 47, 163, 0, 165, 0, 167, 0, 169, 0, 171, 0, 173, 0, 175, 48, 177,
		49, 179, 50, 1, 0, 36, 2, 0, 65, 65, 97, 97, 2, 0, 66, 66, 98, 98, 2, 0,
		67, 67, 99, 99, 2, 0, 68, 68, 100, 100, 2, 0, 69, 69, 101, 101, 2, 0, 70,
		70, 102, 102, 2, 0, 71, 71, 103, 103, 2, 0, 72, 72, 104, 104, 2, 0, 73,
		73, 105, 105, 2, 0, 74, 74, 106, 106, 2, 0, 75, 75, 107, 107, 2, 0, 76,
		76, 108, 108, 2, 0, 77, 77, 109, 109, 2, 0, 78, 78, 110, 110, 2, 0, 79,
		79, 111, 111, 2, 0, 80, 80, 112, 112, 2, 0, 81, 81, 113, 113, 2, 0, 82,
		82, 114, 114, 2, 0, 83, 83, 115, 115, 2, 0, 84, 84, 116, 116, 2, 0, 85,
		85, 117, 117, 2, 0, 86, 86, 118, 118, 2, 0, 87, 87, 119, 119, 2, 0, 88,
		88, 120, 120, 2, 0, 89, 89, 121, 121, 2, 0, 90, 90, 122, 122, 13, 0, 65,
		90, 97, 122, 192, 214, 216, 246, 248, 767, 880, 893, 895, 8191, 8204, 8205,
		8304, 8591, 11264, 12271, 12289, 55295, 63744, 64975, 65008, 65533, 5,
		0, 48, 57, 95, 95, 183, 183, 768, 879, 8255, 8256, 2, 0, 34, 34, 92, 92,
		2, 0, 39, 39, 92, 92, 1, 0, 49, 57, 1, 0, 48, 57, 1, 0, 48, 55, 3, 0, 48,
		57, 65, 70, 97, 102, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13,
		507, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1,
		0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71,
		1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0,
		79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0,
		0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0,
		0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1,
		0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0,
		109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0,
		0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1, 0, 0, 0, 0, 123,
		1, 0, 0, 0, 0, 125, 1, 0, 0, 0, 0, 127, 1, 0, 0, 0, 0, 129, 1, 0, 0, 0,
		0, 131, 1, 0, 0, 0, 0, 133, 1, 0, 0, 0, 0, 135, 1, 0, 0, 0, 0, 137, 1,
		0, 0, 0, 0, 139, 1, 0, 0, 0, 0, 141, 1, 0, 0, 0, 0, 143, 1, 0, 0, 0, 0,
		145, 1, 0, 0, 0, 0, 147, 1, 0, 0, 0, 0, 149, 1, 0, 0, 0, 0, 151, 1, 0,
		0, 0, 0, 155, 1, 0, 0, 0, 0, 157, 1, 0, 0, 0, 0, 159, 1, 0, 0, 0, 0, 161,
		1, 0, 0, 0, 0, 175, 1, 0, 0, 0, 0, 177, 1, 0, 0, 0, 0, 179, 1, 0, 0, 0,
		1, 181, 1, 0, 0, 0, 3, 183, 1, 0, 0, 0, 5, 185, 1, 0, 0, 0, 7, 187, 1,
		0, 0, 0, 9, 189, 1, 0, 0, 0, 11, 191, 1, 0, 0, 0, 13, 193, 1, 0, 0, 0,
		15, 195, 1, 0, 0, 0, 17, 197, 1, 0, 0, 0, 19, 199, 1, 0, 0, 0, 21, 201,
		1, 0, 0, 0, 23, 203, 1, 0, 0, 0, 25, 205, 1, 0, 0, 0, 27, 207, 1, 0, 0,
		0, 29, 209, 1, 0, 0, 0, 31, 211, 1, 0, 0, 0, 33, 213, 1, 0, 0, 0, 35, 215,
		1, 0, 0, 0, 37, 217, 1, 0, 0, 0, 39, 219, 1, 0, 0, 0, 41, 221, 1, 0, 0,
		0, 43, 223, 1, 0, 0, 0, 45, 225, 1, 0, 0, 0, 47, 227, 1, 0, 0, 0, 49, 229,
		1, 0, 0, 0, 51, 231, 1, 0, 0, 0, 53, 233, 1, 0, 0, 0, 55, 237, 1, 0, 0,
		0, 57, 239, 1, 0, 0, 0, 59, 241, 1, 0, 0, 0, 61, 243, 1, 0, 0, 0, 63, 245,
		1, 0, 0, 0, 65, 247, 1, 0, 0, 0, 67, 249, 1, 0, 0, 0, 69, 251, 1, 0, 0,
		0, 71, 253, 1, 0, 0, 0, 73, 255, 1, 0, 0, 0, 75, 257, 1, 0, 0, 0, 77, 259,
		1, 0, 0, 0, 79, 261, 1, 0, 0, 0, 81, 263, 1, 0, 0, 0, 83, 265, 1, 0, 0,
		0, 85, 267, 1, 0, 0, 0, 87, 272, 1, 0, 0, 0, 89, 277, 1, 0, 0, 0, 91, 282,
		1, 0, 0, 0, 93, 285, 1, 0, 0, 0, 95, 288, 1, 0, 0, 0, 97, 293, 1, 0, 0,
		0, 99, 299, 1, 0, 0, 0, 101, 303, 1, 0, 0, 0, 103, 305, 1, 0, 0, 0, 105,
		314, 1, 0, 0, 0, 107, 317, 1, 0, 0, 0, 109, 319, 1, 0, 0, 0, 111, 322,
		1, 0, 0, 0, 113, 325, 1, 0, 0, 0, 115, 328, 1, 0, 0, 0, 117, 331, 1, 0,
		0, 0, 119, 333, 1, 0, 0, 0, 121, 335, 1, 0, 0, 0, 123, 338, 1, 0, 0, 0,
		125, 341, 1, 0, 0, 0, 127, 344, 1, 0, 0, 0, 129, 346, 1, 0, 0, 0, 131,
		348, 1, 0, 0, 0, 133, 351, 1, 0, 0, 0, 135, 359, 1, 0, 0, 0, 137, 363,
		1, 0, 0, 0, 139, 367, 1, 0, 0, 0, 141, 370, 1, 0, 0, 0, 143, 377, 1, 0,
		0, 0, 145, 390, 1, 0, 0, 0, 147, 417, 1, 0, 0, 0, 149, 419, 1, 0, 0, 0,
		151, 426, 1, 0, 0, 0, 153, 440, 1, 0, 0, 0, 155, 442, 1, 0, 0, 0, 157,
		454, 1, 0, 0, 0, 159, 456, 1, 0, 0, 0, 161, 460, 1, 0, 0, 0, 163, 464,
		1, 0, 0, 0, 165, 469, 1, 0, 0, 0, 167, 474, 1, 0, 0, 0, 169, 478, 1, 0,
		0, 0, 171, 480, 1, 0, 0, 0, 173, 482, 1, 0, 0, 0, 175, 485, 1, 0, 0, 0,
		177, 491, 1, 0, 0, 0, 179, 505, 1, 0, 0, 0, 181, 182, 7, 0, 0, 0, 182,
		2, 1, 0, 0, 0, 183, 184, 7, 1, 0, 0, 184, 4, 1, 0, 0, 0, 185, 186, 7, 2,
		0, 0, 186, 6, 1, 0, 0, 0, 187, 188, 7, 3, 0, 0, 188, 8, 1, 0, 0, 0, 189,
		190, 7, 4, 0, 0, 190, 10, 1, 0, 0, 0, 191, 192, 7, 5, 0, 0, 192, 12, 1,
		0, 0, 0, 193, 194, 7, 6, 0, 0, 194, 14, 1, 0, 0, 0, 195, 196, 7, 7, 0,
		0, 196, 16, 1, 0, 0, 0, 197, 198, 7, 8, 0, 0, 198, 18, 1, 0, 0, 0, 199,
		200, 7, 9, 0, 0, 200, 20, 1, 0, 0, 0, 201, 202, 7, 10, 0, 0, 202, 22, 1,
		0, 0, 0, 203, 204, 7, 11, 0, 0, 204, 24, 1, 0, 0, 0, 205, 206, 7, 12, 0,
		0, 206, 26, 1, 0, 0, 0, 207, 208, 7, 13, 0, 0, 208, 28, 1, 0, 0, 0, 209,
		210, 7, 14, 0, 0, 210, 30, 1, 0, 0, 0, 211, 212, 7, 15, 0, 0, 212, 32,
		1, 0, 0, 0, 213, 214, 7, 16, 0, 0, 214, 34, 1, 0, 0, 0, 215, 216, 7, 17,
		0, 0, 216, 36, 1, 0, 0, 0, 217, 218, 7, 18, 0, 0, 218, 38, 1, 0, 0, 0,
		219, 220, 7, 19, 0, 0, 220, 40, 1, 0, 0, 0, 221, 222, 7, 20, 0, 0, 222,
		42, 1, 0, 0, 0, 223, 224, 7, 21, 0, 0, 224, 44, 1, 0, 0, 0, 225, 226, 7,
		22, 0, 0, 226, 46, 1, 0, 0, 0, 227, 228, 7, 23, 0, 0, 228, 48, 1, 0, 0,
		0, 229, 230, 7, 24, 0, 0, 230, 50, 1, 0, 0, 0, 231, 232, 7, 25, 0, 0, 232,
		52, 1, 0, 0, 0, 233, 234, 7, 26, 0, 0, 234, 54, 1, 0, 0, 0, 235, 238, 3,
		53, 26, 0, 236, 238, 7, 27, 0, 0, 237, 235, 1, 0, 0, 0, 237, 236, 1, 0,
		0, 0, 238, 56, 1, 0, 0, 0, 239, 240, 5, 44, 0, 0, 240, 58, 1, 0, 0, 0,
		241, 242, 5, 43, 0, 0, 242, 60, 1, 0, 0, 0, 243, 244, 5, 45, 0, 0, 244,
		62, 1, 0, 0, 0, 245, 246, 5, 47, 0, 0, 246, 64, 1, 0, 0, 0, 247, 248, 5,
		42, 0, 0, 248, 66, 1, 0, 0, 0, 249, 250, 5, 37, 0, 0, 250, 68, 1, 0, 0,
		0, 251, 252, 5, 46, 0, 0, 252, 70, 1, 0, 0, 0, 253, 254, 5, 59, 0, 0, 254,
		72, 1, 0, 0, 0, 255, 256, 5, 123, 0, 0, 256, 74, 1, 0, 0, 0, 257, 258,
		5, 125, 0, 0, 258, 76, 1, 0, 0, 0, 259, 260, 5, 40, 0, 0, 260, 78, 1, 0,
		0, 0, 261, 262, 5, 41, 0, 0, 262, 80, 1, 0, 0, 0, 263, 264, 5, 91, 0, 0,
		264, 82, 1, 0, 0, 0, 265, 266, 5, 93, 0, 0, 266, 84, 1, 0, 0, 0, 267, 268,
		3, 35, 17, 0, 268, 269, 3, 41, 20, 0, 269, 270, 3, 23, 11, 0, 270, 271,
		3, 9, 4, 0, 271, 86, 1, 0, 0, 0, 272, 273, 3, 45, 22, 0, 273, 274, 3, 15,
		7, 0, 274, 275, 3, 9, 4, 0, 275, 276, 3, 27, 13, 0, 276, 88, 1, 0, 0, 0,
		277, 278, 3, 39, 19, 0, 278, 279, 3, 15, 7, 0, 279, 280, 3, 9, 4, 0, 280,
		281, 3, 27, 13, 0, 281, 90, 1, 0, 0, 0, 282, 283, 5, 38, 0, 0, 283, 284,
		5, 38, 0, 0, 284, 92, 1, 0, 0, 0, 285, 286, 5, 124, 0, 0, 286, 287, 5,
		124, 0, 0, 287, 94, 1, 0, 0, 0, 288, 289, 3, 39, 19, 0, 289, 290, 3, 35,
		17, 0, 290, 291, 3, 41, 20, 0, 291, 292, 3, 9, 4, 0, 292, 96, 1, 0, 0,
		0, 293, 294, 3, 11, 5, 0, 294, 295, 3, 1, 0, 0, 295, 296, 3, 23, 11, 0,
		296, 297, 3, 37, 18, 0, 297, 298, 3, 9, 4, 0, 298, 98, 1, 0, 0, 0, 299,
		300, 3, 27, 13, 0, 300, 301, 3, 17, 8, 0, 301, 302, 3, 23, 11, 0, 302,
		100, 1, 0, 0, 0, 303, 304, 5, 33, 0, 0, 304, 102, 1, 0, 0, 0, 305, 306,
		3, 37, 18, 0, 306, 307, 3, 1, 0, 0, 307, 308, 3, 23, 11, 0, 308, 309, 3,
		17, 8, 0, 309, 310, 3, 9, 4, 0, 310, 311, 3, 27, 13, 0, 311, 312, 3, 5,
		2, 0, 312, 313, 3, 9, 4, 0, 313, 104, 1, 0, 0, 0, 314, 315, 5, 61, 0, 0,
		315, 316, 5, 61, 0, 0, 316, 106, 1, 0, 0, 0, 317, 318, 5, 61, 0, 0, 318,
		108, 1, 0, 0, 0, 319, 320, 5, 43, 0, 0, 320, 321, 5, 61, 0, 0, 321, 110,
		1, 0, 0, 0, 322, 323, 5, 45, 0, 0, 323, 324, 5, 61, 0, 0, 324, 112, 1,
		0, 0, 0, 325, 326, 5, 47, 0, 0, 326, 327, 5, 61, 0, 0, 327, 114, 1, 0,
		0, 0, 328, 329, 5, 42, 0, 0, 329, 330, 5, 61, 0, 0, 330, 116, 1, 0, 0,
		0, 331, 332, 5, 62, 0, 0, 332, 118, 1, 0, 0, 0, 333, 334, 5, 60, 0, 0,
		334, 120, 1, 0, 0, 0, 335, 336, 5, 62, 0, 0, 336, 337, 5, 61, 0, 0, 337,
		122, 1, 0, 0, 0, 338, 339, 5, 60, 0, 0, 339, 340, 5, 61, 0, 0, 340, 124,
		1, 0, 0, 0, 341, 342, 5, 33, 0, 0, 342, 343, 5, 61, 0, 0, 343, 126, 1,
		0, 0, 0, 344, 345, 5, 38, 0, 0, 345, 128, 1, 0, 0, 0, 346, 347, 5, 124,
		0, 0, 347, 130, 1, 0, 0, 0, 348, 349, 3, 29, 14, 0, 349, 350, 3, 27, 13,
		0, 350, 132, 1, 0, 0, 0, 351, 352, 3, 7, 3, 0, 352, 353, 3, 9, 4, 0, 353,
		354, 3, 11, 5, 0, 354, 355, 3, 1, 0, 0, 355, 356, 3, 41, 20, 0, 356, 357,
		3, 23, 11, 0, 357, 358, 3, 39, 19, 0, 358, 134, 1, 0, 0, 0, 359, 360, 3,
		11, 5, 0, 360, 361, 3, 29, 14, 0, 361, 362, 3, 35, 17, 0, 362, 136, 1,
		0, 0, 0, 363, 364, 3, 1, 0, 0, 364, 365, 3, 23, 11, 0, 365, 366, 3, 23,
		11, 0, 366, 138, 1, 0, 0, 0, 367, 368, 3, 7, 3, 0, 368, 369, 3, 29, 14,
		0, 369, 140, 1, 0, 0, 0, 370, 374, 3, 53, 26, 0, 371, 373, 3, 55, 27, 0,
		372, 371, 1, 0, 0, 0, 373, 376, 1, 0, 0, 0, 374, 372, 1, 0, 0, 0, 374,
		375, 1, 0, 0, 0, 375, 142, 1, 0, 0, 0, 376, 374, 1, 0, 0, 0, 377, 385,
		5, 34, 0, 0, 378, 379, 5, 92, 0, 0, 379, 384, 9, 0, 0, 0, 380, 381, 5,
		34, 0, 0, 381, 384, 5, 34, 0, 0, 382, 384, 8, 28, 0, 0, 383, 378, 1, 0,
		0, 0, 383, 380, 1, 0, 0, 0, 383, 382, 1, 0, 0, 0, 384, 387, 1, 0, 0, 0,
		385, 383, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 388, 1, 0, 0, 0, 387,
		385, 1, 0, 0, 0, 388, 389, 5, 34, 0, 0, 389, 144, 1, 0, 0, 0, 390, 398,
		5, 39, 0, 0, 391, 392, 5, 92, 0, 0, 392, 397, 9, 0, 0, 0, 393, 394, 5,
		39, 0, 0, 394, 397, 5, 39, 0, 0, 395, 397, 8, 29, 0, 0, 396, 391, 1, 0,
		0, 0, 396, 393, 1, 0, 0, 0, 396, 395, 1, 0, 0, 0, 397, 400, 1, 0, 0, 0,
		398, 396, 1, 0, 0, 0, 398, 399, 1, 0, 0, 0, 399, 401, 1, 0, 0, 0, 400,
		398, 1, 0, 0, 0, 401, 402, 5, 39, 0, 0, 402, 146, 1, 0, 0, 0, 403, 404,
		3, 157, 78, 0, 404, 405, 3, 69, 34, 0, 405, 407, 3, 165, 82, 0, 406, 408,
		3, 149, 74, 0, 407, 406, 1, 0, 0, 0, 407, 408, 1, 0, 0, 0, 408, 418, 1,
		0, 0, 0, 409, 410, 3, 157, 78, 0, 410, 411, 3, 149, 74, 0, 411, 418, 1,
		0, 0, 0, 412, 413, 3, 69, 34, 0, 413, 415, 3, 165, 82, 0, 414, 416, 3,
		149, 74, 0, 415, 414, 1, 0, 0, 0, 415, 416, 1, 0, 0, 0, 416, 418, 1, 0,
		0, 0, 417, 403, 1, 0, 0, 0, 417, 409, 1, 0, 0, 0, 417, 412, 1, 0, 0, 0,
		418, 148, 1, 0, 0, 0, 419, 422, 3, 9, 4, 0, 420, 423, 3, 59, 29, 0, 421,
		423, 3, 61, 30, 0, 422, 420, 1, 0, 0, 0, 422, 421, 1, 0, 0, 0, 422, 423,
		1, 0, 0, 0, 423, 424, 1, 0, 0, 0, 424, 425, 3, 165, 82, 0, 425, 150, 1,
		0, 0, 0, 426, 427, 5, 48, 0, 0, 427, 428, 3, 47, 23, 0, 428, 429, 3, 153,
		76, 0, 429, 430, 3, 155, 77, 0, 430, 152, 1, 0, 0, 0, 431, 432, 3, 163,
		81, 0, 432, 434, 3, 69, 34, 0, 433, 435, 3, 163, 81, 0, 434, 433, 1, 0,
		0, 0, 434, 435, 1, 0, 0, 0, 435, 441, 1, 0, 0, 0, 436, 441, 3, 163, 81,
		0, 437, 438, 3, 69, 34, 0, 438, 439, 3, 163, 81, 0, 439, 441, 1, 0, 0,
		0, 440, 431, 1, 0, 0, 0, 440, 436, 1, 0, 0, 0, 440, 437, 1, 0, 0, 0, 441,
		154, 1, 0, 0, 0, 442, 445, 3, 31, 15, 0, 443, 446, 3, 59, 29, 0, 444, 446,
		3, 61, 30, 0, 445, 443, 1, 0, 0, 0, 445, 444, 1, 0, 0, 0, 445, 446, 1,
		0, 0, 0, 446, 447, 1, 0, 0, 0, 447, 448, 3, 165, 82, 0, 448, 156, 1, 0,
		0, 0, 449, 455, 5, 48, 0, 0, 450, 452, 7, 30, 0, 0, 451, 453, 3, 165, 82,
		0, 452, 451, 1, 0, 0, 0, 452, 453, 1, 0, 0, 0, 453, 455, 1, 0, 0, 0, 454,
		449, 1, 0, 0, 0, 454, 450, 1, 0, 0, 0, 455, 158, 1, 0, 0, 0, 456, 457,
		5, 48, 0, 0, 457, 458, 3, 47, 23, 0, 458, 459, 3, 163, 81, 0, 459, 160,
		1, 0, 0, 0, 460, 461, 5, 48, 0, 0, 461, 462, 3, 167, 83, 0, 462, 162, 1,
		0, 0, 0, 463, 465, 3, 173, 86, 0, 464, 463, 1, 0, 0, 0, 465, 466, 1, 0,
		0, 0, 466, 464, 1, 0, 0, 0, 466, 467, 1, 0, 0, 0, 467, 164, 1, 0, 0, 0,
		468, 470, 3, 169, 84, 0, 469, 468, 1, 0, 0, 0, 470, 471, 1, 0, 0, 0, 471,
		469, 1, 0, 0, 0, 471, 472, 1, 0, 0, 0, 472, 166, 1, 0, 0, 0, 473, 475,
		3, 171, 85, 0, 474, 473, 1, 0, 0, 0, 475, 476, 1, 0, 0, 0, 476, 474, 1,
		0, 0, 0, 476, 477, 1, 0, 0, 0, 477, 168, 1, 0, 0, 0, 478, 479, 7, 31, 0,
		0, 479, 170, 1, 0, 0, 0, 480, 481, 7, 32, 0, 0, 481, 172, 1, 0, 0, 0, 482,
		483, 7, 33, 0, 0, 483, 174, 1, 0, 0, 0, 484, 486, 7, 34, 0, 0, 485, 484,
		1, 0, 0, 0, 486, 487, 1, 0, 0, 0, 487, 485, 1, 0, 0, 0, 487, 488, 1, 0,
		0, 0, 488, 489, 1, 0, 0, 0, 489, 490, 6, 87, 0, 0, 490, 176, 1, 0, 0, 0,
		491, 492, 5, 47, 0, 0, 492, 493, 5, 42, 0, 0, 493, 497, 1, 0, 0, 0, 494,
		496, 9, 0, 0, 0, 495, 494, 1, 0, 0, 0, 496, 499, 1, 0, 0, 0, 497, 498,
		1, 0, 0, 0, 497, 495, 1, 0, 0, 0, 498, 500, 1, 0, 0, 0, 499, 497, 1, 0,
		0, 0, 500, 501, 5, 42, 0, 0, 501, 502, 5, 47, 0, 0, 502, 503, 1, 0, 0,
		0, 503, 504, 6, 88, 0, 0, 504, 178, 1, 0, 0, 0, 505, 506, 5, 47, 0, 0,
		506, 507, 5, 47, 0, 0, 507, 511, 1, 0, 0, 0, 508, 510, 8, 35, 0, 0, 509,
		508, 1, 0, 0, 0, 510, 513, 1, 0, 0, 0, 511, 509, 1, 0, 0, 0, 511, 512,
		1, 0, 0, 0, 512, 514, 1, 0, 0, 0, 513, 511, 1, 0, 0, 0, 514, 515, 6, 89,
		0, 0, 515, 180, 1, 0, 0, 0, 22, 0, 237, 374, 383, 385, 396, 398, 407, 415,
		417, 422, 434, 440, 445, 452, 454, 466, 471, 476, 487, 497, 511, 1, 6,
		0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// EcaruleLexerInit initializes any static state used to implement EcaruleLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewEcaruleLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func EcaruleLexerInit() {
	staticData := &EcaruleLexerLexerStaticData
	staticData.once.Do(ecarulelexerLexerInit)
}

// NewEcaruleLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewEcaruleLexer(input antlr.CharStream) *EcaruleLexer {
	EcaruleLexerInit()
	l := new(EcaruleLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &EcaruleLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "EcaruleLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// EcaruleLexer tokens.
const (
	EcaruleLexerT__0              = 1
	EcaruleLexerPLUS              = 2
	EcaruleLexerMINUS             = 3
	EcaruleLexerDIV               = 4
	EcaruleLexerMUL               = 5
	EcaruleLexerMOD               = 6
	EcaruleLexerDOT               = 7
	EcaruleLexerSEMICOLON         = 8
	EcaruleLexerLR_BRACE          = 9
	EcaruleLexerRR_BRACE          = 10
	EcaruleLexerLR_BRACKET        = 11
	EcaruleLexerRR_BRACKET        = 12
	EcaruleLexerLS_BRACKET        = 13
	EcaruleLexerRS_BRACKET        = 14
	EcaruleLexerRULE              = 15
	EcaruleLexerWHEN              = 16
	EcaruleLexerTHEN              = 17
	EcaruleLexerAND               = 18
	EcaruleLexerOR                = 19
	EcaruleLexerTRUE              = 20
	EcaruleLexerFALSE             = 21
	EcaruleLexerNIL_LITERAL       = 22
	EcaruleLexerNEGATION          = 23
	EcaruleLexerSALIENCE          = 24
	EcaruleLexerEQUALS            = 25
	EcaruleLexerASSIGN            = 26
	EcaruleLexerPLUS_ASIGN        = 27
	EcaruleLexerMINUS_ASIGN       = 28
	EcaruleLexerDIV_ASIGN         = 29
	EcaruleLexerMUL_ASIGN         = 30
	EcaruleLexerGT                = 31
	EcaruleLexerLT                = 32
	EcaruleLexerGTE               = 33
	EcaruleLexerLTE               = 34
	EcaruleLexerNOTEQUALS         = 35
	EcaruleLexerBITAND            = 36
	EcaruleLexerBITOR             = 37
	EcaruleLexerSIMPLENAME        = 38
	EcaruleLexerDQUOTA_STRING     = 39
	EcaruleLexerSQUOTA_STRING     = 40
	EcaruleLexerDECIMAL_FLOAT_LIT = 41
	EcaruleLexerDECIMAL_EXPONENT  = 42
	EcaruleLexerHEX_FLOAT_LIT     = 43
	EcaruleLexerHEX_EXPONENT      = 44
	EcaruleLexerDEC_LIT           = 45
	EcaruleLexerHEX_LIT           = 46
	EcaruleLexerOCT_LIT           = 47
	EcaruleLexerSPACE             = 48
	EcaruleLexerCOMMENT           = 49
	EcaruleLexerLINE_COMMENT      = 50
	EcaruleLexerON                = 51
	EcaruleLexerDEFAULT           = 52
	EcaruleLexerFOR               = 53
	EcaruleLexerALL               = 54
	EcaruleLexerDO                = 55
)
