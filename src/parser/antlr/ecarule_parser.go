// Code generated from EcaruleParser.g4 by ANTLR 4.9.2.
// then MODIFIED by adding grulev3 import and by:
// sed '/type GrlContext struct/,$s/listenerT\.\(Enter\|Exit\)\([A-Za-z]\+\)(s)/c := grulev3.New\2Context(s.parser, s.BaseParserRuleContext, -1)\n\t\tc.BaseParserRuleContext = s.BaseParserRuleContext\n\t\tlistenerT.\1\2(c)/'

package antlr // EcaruleParser
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/hyperjumptech/grule-rule-engine/antlr/parser/grulev3"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 59, 320,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 83, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5,
	2, 89, 10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 97, 10, 3, 12,
	3, 14, 3, 100, 11, 3, 3, 4, 3, 4, 5, 4, 104, 10, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 115, 10, 5, 12, 5, 14, 5, 118, 11,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 7, 7, 125, 10, 7, 12, 7, 14, 7, 128, 11,
	7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 5, 8, 135, 10, 8, 3, 8, 5, 8, 138, 10,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 6,
	14, 161, 10, 14, 13, 14, 14, 14, 162, 3, 15, 3, 15, 5, 15, 167, 10, 15,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 5, 17, 175, 10, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 5, 17, 182, 10, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 204, 10, 17, 12, 17, 14,
	17, 207, 11, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21,
	3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 225, 10,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 7, 23, 233, 10, 23, 12, 23,
	14, 23, 236, 11, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 243, 10,
	24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 7, 25, 252, 10, 25,
	12, 25, 14, 25, 255, 11, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27,
	3, 27, 3, 28, 3, 28, 3, 28, 5, 28, 267, 10, 28, 3, 28, 3, 28, 3, 29, 3,
	29, 3, 29, 3, 30, 3, 30, 3, 30, 7, 30, 277, 10, 30, 12, 30, 14, 30, 280,
	11, 30, 3, 31, 3, 31, 5, 31, 284, 10, 31, 3, 32, 5, 32, 287, 10, 32, 3,
	32, 3, 32, 3, 33, 5, 33, 292, 10, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34,
	5, 34, 299, 10, 34, 3, 35, 5, 35, 302, 10, 35, 3, 35, 3, 35, 3, 36, 5,
	36, 307, 10, 36, 3, 36, 3, 36, 3, 37, 5, 37, 312, 10, 37, 3, 37, 3, 37,
	3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 2, 5, 32, 44, 48, 40, 2, 4, 6, 8, 10,
	12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
	48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 2, 9, 3, 2,
	57, 58, 3, 2, 41, 42, 3, 2, 28, 32, 3, 2, 6, 8, 4, 2, 4, 5, 38, 39, 4,
	2, 27, 27, 33, 37, 3, 2, 22, 23, 2, 320, 2, 78, 3, 2, 2, 2, 4, 92, 3, 2,
	2, 2, 6, 101, 3, 2, 2, 2, 8, 109, 3, 2, 2, 2, 10, 119, 3, 2, 2, 2, 12,
	126, 3, 2, 2, 2, 14, 131, 3, 2, 2, 2, 16, 144, 3, 2, 2, 2, 18, 147, 3,
	2, 2, 2, 20, 149, 3, 2, 2, 2, 22, 151, 3, 2, 2, 2, 24, 154, 3, 2, 2, 2,
	26, 160, 3, 2, 2, 2, 28, 166, 3, 2, 2, 2, 30, 168, 3, 2, 2, 2, 32, 181,
	3, 2, 2, 2, 34, 208, 3, 2, 2, 2, 36, 210, 3, 2, 2, 2, 38, 212, 3, 2, 2,
	2, 40, 214, 3, 2, 2, 2, 42, 216, 3, 2, 2, 2, 44, 224, 3, 2, 2, 2, 46, 242,
	3, 2, 2, 2, 48, 244, 3, 2, 2, 2, 50, 256, 3, 2, 2, 2, 52, 260, 3, 2, 2,
	2, 54, 263, 3, 2, 2, 2, 56, 270, 3, 2, 2, 2, 58, 273, 3, 2, 2, 2, 60, 283,
	3, 2, 2, 2, 62, 286, 3, 2, 2, 2, 64, 291, 3, 2, 2, 2, 66, 298, 3, 2, 2,
	2, 68, 301, 3, 2, 2, 2, 70, 306, 3, 2, 2, 2, 72, 311, 3, 2, 2, 2, 74, 315,
	3, 2, 2, 2, 76, 317, 3, 2, 2, 2, 78, 79, 7, 17, 2, 2, 79, 82, 7, 40, 2,
	2, 80, 81, 7, 54, 2, 2, 81, 83, 7, 40, 2, 2, 82, 80, 3, 2, 2, 2, 82, 83,
	3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 85, 7, 55, 2, 2, 85, 88, 5, 4, 3, 2,
	86, 87, 7, 53, 2, 2, 87, 89, 5, 8, 5, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3,
	2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 91, 5, 6, 4, 2, 91, 3, 3, 2, 2, 2, 92,
	93, 7, 40, 2, 2, 93, 98, 7, 10, 2, 2, 94, 95, 7, 40, 2, 2, 95, 97, 7, 10,
	2, 2, 96, 94, 3, 2, 2, 2, 97, 100, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 98,
	99, 3, 2, 2, 2, 99, 5, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 101, 103, 7, 59,
	2, 2, 102, 104, 9, 2, 2, 2, 103, 102, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2,
	104, 105, 3, 2, 2, 2, 105, 106, 5, 32, 17, 2, 106, 107, 7, 56, 2, 2, 107,
	108, 5, 8, 5, 2, 108, 7, 3, 2, 2, 2, 109, 110, 5, 10, 6, 2, 110, 116, 7,
	10, 2, 2, 111, 112, 5, 10, 6, 2, 112, 113, 7, 10, 2, 2, 113, 115, 3, 2,
	2, 2, 114, 111, 3, 2, 2, 2, 115, 118, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2,
	116, 117, 3, 2, 2, 2, 117, 9, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 119, 120,
	7, 40, 2, 2, 120, 121, 7, 28, 2, 2, 121, 122, 5, 32, 17, 2, 122, 11, 3,
	2, 2, 2, 123, 125, 5, 14, 8, 2, 124, 123, 3, 2, 2, 2, 125, 128, 3, 2, 2,
	2, 126, 124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 129, 3, 2, 2, 2, 128,
	126, 3, 2, 2, 2, 129, 130, 7, 2, 2, 3, 130, 13, 3, 2, 2, 2, 131, 132, 7,
	17, 2, 2, 132, 134, 5, 18, 10, 2, 133, 135, 5, 20, 11, 2, 134, 133, 3,
	2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 137, 3, 2, 2, 2, 136, 138, 5, 16, 9,
	2, 137, 136, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139,
	140, 7, 11, 2, 2, 140, 141, 5, 22, 12, 2, 141, 142, 5, 24, 13, 2, 142,
	143, 7, 12, 2, 2, 143, 15, 3, 2, 2, 2, 144, 145, 7, 26, 2, 2, 145, 146,
	5, 66, 34, 2, 146, 17, 3, 2, 2, 2, 147, 148, 7, 40, 2, 2, 148, 19, 3, 2,
	2, 2, 149, 150, 9, 3, 2, 2, 150, 21, 3, 2, 2, 2, 151, 152, 7, 18, 2, 2,
	152, 153, 5, 32, 17, 2, 153, 23, 3, 2, 2, 2, 154, 155, 7, 19, 2, 2, 155,
	156, 5, 26, 14, 2, 156, 25, 3, 2, 2, 2, 157, 158, 5, 28, 15, 2, 158, 159,
	7, 10, 2, 2, 159, 161, 3, 2, 2, 2, 160, 157, 3, 2, 2, 2, 161, 162, 3, 2,
	2, 2, 162, 160, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 27, 3, 2, 2, 2,
	164, 167, 5, 30, 16, 2, 165, 167, 5, 44, 23, 2, 166, 164, 3, 2, 2, 2, 166,
	165, 3, 2, 2, 2, 167, 29, 3, 2, 2, 2, 168, 169, 5, 48, 25, 2, 169, 170,
	9, 4, 2, 2, 170, 171, 5, 32, 17, 2, 171, 31, 3, 2, 2, 2, 172, 174, 8, 17,
	1, 2, 173, 175, 7, 25, 2, 2, 174, 173, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2,
	175, 176, 3, 2, 2, 2, 176, 177, 7, 13, 2, 2, 177, 178, 5, 32, 17, 2, 178,
	179, 7, 14, 2, 2, 179, 182, 3, 2, 2, 2, 180, 182, 5, 44, 23, 2, 181, 172,
	3, 2, 2, 2, 181, 180, 3, 2, 2, 2, 182, 205, 3, 2, 2, 2, 183, 184, 12, 9,
	2, 2, 184, 185, 5, 34, 18, 2, 185, 186, 5, 32, 17, 10, 186, 204, 3, 2,
	2, 2, 187, 188, 12, 8, 2, 2, 188, 189, 5, 36, 19, 2, 189, 190, 5, 32, 17,
	9, 190, 204, 3, 2, 2, 2, 191, 192, 12, 7, 2, 2, 192, 193, 5, 38, 20, 2,
	193, 194, 5, 32, 17, 8, 194, 204, 3, 2, 2, 2, 195, 196, 12, 6, 2, 2, 196,
	197, 5, 40, 21, 2, 197, 198, 5, 32, 17, 7, 198, 204, 3, 2, 2, 2, 199, 200,
	12, 5, 2, 2, 200, 201, 5, 42, 22, 2, 201, 202, 5, 32, 17, 6, 202, 204,
	3, 2, 2, 2, 203, 183, 3, 2, 2, 2, 203, 187, 3, 2, 2, 2, 203, 191, 3, 2,
	2, 2, 203, 195, 3, 2, 2, 2, 203, 199, 3, 2, 2, 2, 204, 207, 3, 2, 2, 2,
	205, 203, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 33, 3, 2, 2, 2, 207, 205,
	3, 2, 2, 2, 208, 209, 9, 5, 2, 2, 209, 35, 3, 2, 2, 2, 210, 211, 9, 6,
	2, 2, 211, 37, 3, 2, 2, 2, 212, 213, 9, 7, 2, 2, 213, 39, 3, 2, 2, 2, 214,
	215, 7, 20, 2, 2, 215, 41, 3, 2, 2, 2, 216, 217, 7, 21, 2, 2, 217, 43,
	3, 2, 2, 2, 218, 219, 8, 23, 1, 2, 219, 225, 5, 46, 24, 2, 220, 225, 5,
	48, 25, 2, 221, 225, 5, 54, 28, 2, 222, 223, 7, 25, 2, 2, 223, 225, 5,
	44, 23, 3, 224, 218, 3, 2, 2, 2, 224, 220, 3, 2, 2, 2, 224, 221, 3, 2,
	2, 2, 224, 222, 3, 2, 2, 2, 225, 234, 3, 2, 2, 2, 226, 227, 12, 6, 2, 2,
	227, 233, 5, 56, 29, 2, 228, 229, 12, 5, 2, 2, 229, 233, 5, 52, 27, 2,
	230, 231, 12, 4, 2, 2, 231, 233, 5, 50, 26, 2, 232, 226, 3, 2, 2, 2, 232,
	228, 3, 2, 2, 2, 232, 230, 3, 2, 2, 2, 233, 236, 3, 2, 2, 2, 234, 232,
	3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 45, 3, 2, 2, 2, 236, 234, 3, 2,
	2, 2, 237, 243, 5, 74, 38, 2, 238, 243, 5, 66, 34, 2, 239, 243, 5, 60,
	31, 2, 240, 243, 5, 76, 39, 2, 241, 243, 7, 24, 2, 2, 242, 237, 3, 2, 2,
	2, 242, 238, 3, 2, 2, 2, 242, 239, 3, 2, 2, 2, 242, 240, 3, 2, 2, 2, 242,
	241, 3, 2, 2, 2, 243, 47, 3, 2, 2, 2, 244, 245, 8, 25, 1, 2, 245, 246,
	7, 40, 2, 2, 246, 253, 3, 2, 2, 2, 247, 248, 12, 5, 2, 2, 248, 252, 5,
	52, 27, 2, 249, 250, 12, 4, 2, 2, 250, 252, 5, 50, 26, 2, 251, 247, 3,
	2, 2, 2, 251, 249, 3, 2, 2, 2, 252, 255, 3, 2, 2, 2, 253, 251, 3, 2, 2,
	2, 253, 254, 3, 2, 2, 2, 254, 49, 3, 2, 2, 2, 255, 253, 3, 2, 2, 2, 256,
	257, 7, 15, 2, 2, 257, 258, 5, 32, 17, 2, 258, 259, 7, 16, 2, 2, 259, 51,
	3, 2, 2, 2, 260, 261, 7, 9, 2, 2, 261, 262, 7, 40, 2, 2, 262, 53, 3, 2,
	2, 2, 263, 264, 7, 40, 2, 2, 264, 266, 7, 13, 2, 2, 265, 267, 5, 58, 30,
	2, 266, 265, 3, 2, 2, 2, 266, 267, 3, 2, 2, 2, 267, 268, 3, 2, 2, 2, 268,
	269, 7, 14, 2, 2, 269, 55, 3, 2, 2, 2, 270, 271, 7, 9, 2, 2, 271, 272,
	5, 54, 28, 2, 272, 57, 3, 2, 2, 2, 273, 278, 5, 32, 17, 2, 274, 275, 7,
	3, 2, 2, 275, 277, 5, 32, 17, 2, 276, 274, 3, 2, 2, 2, 277, 280, 3, 2,
	2, 2, 278, 276, 3, 2, 2, 2, 278, 279, 3, 2, 2, 2, 279, 59, 3, 2, 2, 2,
	280, 278, 3, 2, 2, 2, 281, 284, 5, 62, 32, 2, 282, 284, 5, 64, 33, 2, 283,
	281, 3, 2, 2, 2, 283, 282, 3, 2, 2, 2, 284, 61, 3, 2, 2, 2, 285, 287, 7,
	5, 2, 2, 286, 285, 3, 2, 2, 2, 286, 287, 3, 2, 2, 2, 287, 288, 3, 2, 2,
	2, 288, 289, 7, 43, 2, 2, 289, 63, 3, 2, 2, 2, 290, 292, 7, 5, 2, 2, 291,
	290, 3, 2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 293, 3, 2, 2, 2, 293, 294,
	7, 45, 2, 2, 294, 65, 3, 2, 2, 2, 295, 299, 5, 68, 35, 2, 296, 299, 5,
	70, 36, 2, 297, 299, 5, 72, 37, 2, 298, 295, 3, 2, 2, 2, 298, 296, 3, 2,
	2, 2, 298, 297, 3, 2, 2, 2, 299, 67, 3, 2, 2, 2, 300, 302, 7, 5, 2, 2,
	301, 300, 3, 2, 2, 2, 301, 302, 3, 2, 2, 2, 302, 303, 3, 2, 2, 2, 303,
	304, 7, 47, 2, 2, 304, 69, 3, 2, 2, 2, 305, 307, 7, 5, 2, 2, 306, 305,
	3, 2, 2, 2, 306, 307, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308, 309, 7, 48,
	2, 2, 309, 71, 3, 2, 2, 2, 310, 312, 7, 5, 2, 2, 311, 310, 3, 2, 2, 2,
	311, 312, 3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313, 314, 7, 49, 2, 2, 314,
	73, 3, 2, 2, 2, 315, 316, 9, 3, 2, 2, 316, 75, 3, 2, 2, 2, 317, 318, 9,
	8, 2, 2, 318, 77, 3, 2, 2, 2, 31, 82, 88, 98, 103, 116, 126, 134, 137,
	162, 166, 174, 181, 203, 205, 224, 232, 234, 242, 251, 253, 266, 278, 283,
	286, 291, 298, 301, 306, 311,
}
var literalNames []string

var symbolicNames = []string{
	"", "", "PLUS", "MINUS", "DIV", "MUL", "MOD", "DOT", "SEMICOLON", "LR_BRACE",
	"RR_BRACE", "LR_BRACKET", "RR_BRACKET", "LS_BRACKET", "RS_BRACKET", "RULE",
	"WHEN", "THEN", "AND", "OR", "TRUE", "FALSE", "NIL_LITERAL", "NEGATION",
	"SALIENCE", "EQUALS", "ASSIGN", "PLUS_ASIGN", "MINUS_ASIGN", "DIV_ASIGN",
	"MUL_ASIGN", "GT", "LT", "GTE", "LTE", "NOTEQUALS", "BITAND", "BITOR",
	"SIMPLENAME", "DQUOTA_STRING", "SQUOTA_STRING", "DECIMAL_FLOAT_LIT", "DECIMAL_EXPONENT",
	"HEX_FLOAT_LIT", "HEX_EXPONENT", "DEC_LIT", "HEX_LIT", "OCT_LIT", "SPACE",
	"COMMENT", "LINE_COMMENT", "DEFAULT", "IN", "ON", "DO", "ALL", "SOME",
	"FOR",
}

var ruleNames = []string{
	"prule", "evt", "task", "actslist", "act", "grl", "ruleEntry", "salience",
	"ruleName", "ruleDescription", "whenScope", "thenScope", "thenExpressionList",
	"thenExpression", "assignment", "expression", "mulDivOperators", "addMinusOperators",
	"comparisonOperator", "andLogicOperator", "orLogicOperator", "expressionAtom",
	"constant", "variable", "arrayMapSelector", "memberVariable", "functionCall",
	"methodCall", "argumentList", "floatLiteral", "decimalFloatLiteral", "hexadecimalFloatLiteral",
	"integerLiteral", "decimalLiteral", "hexadecimalLiteral", "octalLiteral",
	"stringLiteral", "booleanLiteral",
}

type EcaruleParser struct {
	*antlr.BaseParser
}

// NewEcaruleParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *EcaruleParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewEcaruleParser(input antlr.TokenStream) *EcaruleParser {
	this := new(EcaruleParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "EcaruleParser.g4"

	return this
}

// EcaruleParser tokens.
const (
	EcaruleParserEOF               = antlr.TokenEOF
	EcaruleParserT__0              = 1
	EcaruleParserPLUS              = 2
	EcaruleParserMINUS             = 3
	EcaruleParserDIV               = 4
	EcaruleParserMUL               = 5
	EcaruleParserMOD               = 6
	EcaruleParserDOT               = 7
	EcaruleParserSEMICOLON         = 8
	EcaruleParserLR_BRACE          = 9
	EcaruleParserRR_BRACE          = 10
	EcaruleParserLR_BRACKET        = 11
	EcaruleParserRR_BRACKET        = 12
	EcaruleParserLS_BRACKET        = 13
	EcaruleParserRS_BRACKET        = 14
	EcaruleParserRULE              = 15
	EcaruleParserWHEN              = 16
	EcaruleParserTHEN              = 17
	EcaruleParserAND               = 18
	EcaruleParserOR                = 19
	EcaruleParserTRUE              = 20
	EcaruleParserFALSE             = 21
	EcaruleParserNIL_LITERAL       = 22
	EcaruleParserNEGATION          = 23
	EcaruleParserSALIENCE          = 24
	EcaruleParserEQUALS            = 25
	EcaruleParserASSIGN            = 26
	EcaruleParserPLUS_ASIGN        = 27
	EcaruleParserMINUS_ASIGN       = 28
	EcaruleParserDIV_ASIGN         = 29
	EcaruleParserMUL_ASIGN         = 30
	EcaruleParserGT                = 31
	EcaruleParserLT                = 32
	EcaruleParserGTE               = 33
	EcaruleParserLTE               = 34
	EcaruleParserNOTEQUALS         = 35
	EcaruleParserBITAND            = 36
	EcaruleParserBITOR             = 37
	EcaruleParserSIMPLENAME        = 38
	EcaruleParserDQUOTA_STRING     = 39
	EcaruleParserSQUOTA_STRING     = 40
	EcaruleParserDECIMAL_FLOAT_LIT = 41
	EcaruleParserDECIMAL_EXPONENT  = 42
	EcaruleParserHEX_FLOAT_LIT     = 43
	EcaruleParserHEX_EXPONENT      = 44
	EcaruleParserDEC_LIT           = 45
	EcaruleParserHEX_LIT           = 46
	EcaruleParserOCT_LIT           = 47
	EcaruleParserSPACE             = 48
	EcaruleParserCOMMENT           = 49
	EcaruleParserLINE_COMMENT      = 50
	EcaruleParserDEFAULT           = 51
	EcaruleParserIN                = 52
	EcaruleParserON                = 53
	EcaruleParserDO                = 54
	EcaruleParserALL               = 55
	EcaruleParserSOME              = 56
	EcaruleParserFOR               = 57
)

// EcaruleParser rules.
const (
	EcaruleParserRULE_prule                   = 0
	EcaruleParserRULE_evt                     = 1
	EcaruleParserRULE_task                    = 2
	EcaruleParserRULE_actslist                = 3
	EcaruleParserRULE_act                     = 4
	EcaruleParserRULE_grl                     = 5
	EcaruleParserRULE_ruleEntry               = 6
	EcaruleParserRULE_salience                = 7
	EcaruleParserRULE_ruleName                = 8
	EcaruleParserRULE_ruleDescription         = 9
	EcaruleParserRULE_whenScope               = 10
	EcaruleParserRULE_thenScope               = 11
	EcaruleParserRULE_thenExpressionList      = 12
	EcaruleParserRULE_thenExpression          = 13
	EcaruleParserRULE_assignment              = 14
	EcaruleParserRULE_expression              = 15
	EcaruleParserRULE_mulDivOperators         = 16
	EcaruleParserRULE_addMinusOperators       = 17
	EcaruleParserRULE_comparisonOperator      = 18
	EcaruleParserRULE_andLogicOperator        = 19
	EcaruleParserRULE_orLogicOperator         = 20
	EcaruleParserRULE_expressionAtom          = 21
	EcaruleParserRULE_constant                = 22
	EcaruleParserRULE_variable                = 23
	EcaruleParserRULE_arrayMapSelector        = 24
	EcaruleParserRULE_memberVariable          = 25
	EcaruleParserRULE_functionCall            = 26
	EcaruleParserRULE_methodCall              = 27
	EcaruleParserRULE_argumentList            = 28
	EcaruleParserRULE_floatLiteral            = 29
	EcaruleParserRULE_decimalFloatLiteral     = 30
	EcaruleParserRULE_hexadecimalFloatLiteral = 31
	EcaruleParserRULE_integerLiteral          = 32
	EcaruleParserRULE_decimalLiteral          = 33
	EcaruleParserRULE_hexadecimalLiteral      = 34
	EcaruleParserRULE_octalLiteral            = 35
	EcaruleParserRULE_stringLiteral           = 36
	EcaruleParserRULE_booleanLiteral          = 37
)

// IPruleContext is an interface to support dynamic dispatch.
type IPruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPruleContext differentiates from other interfaces.
	IsPruleContext()
}

type PruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPruleContext() *PruleContext {
	var p = new(PruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_prule
	return p
}

func (*PruleContext) IsPruleContext() {}

func NewPruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PruleContext {
	var p = new(PruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_prule

	return p
}

func (s *PruleContext) GetParser() antlr.Parser { return s.parser }

func (s *PruleContext) RULE() antlr.TerminalNode {
	return s.GetToken(EcaruleParserRULE, 0)
}

func (s *PruleContext) AllSIMPLENAME() []antlr.TerminalNode {
	return s.GetTokens(EcaruleParserSIMPLENAME)
}

func (s *PruleContext) SIMPLENAME(i int) antlr.TerminalNode {
	return s.GetToken(EcaruleParserSIMPLENAME, i)
}

func (s *PruleContext) ON() antlr.TerminalNode {
	return s.GetToken(EcaruleParserON, 0)
}

func (s *PruleContext) Evt() IEvtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEvtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEvtContext)
}

func (s *PruleContext) Task() ITaskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITaskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITaskContext)
}

func (s *PruleContext) IN() antlr.TerminalNode {
	return s.GetToken(EcaruleParserIN, 0)
}

func (s *PruleContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(EcaruleParserDEFAULT, 0)
}

func (s *PruleContext) Actslist() IActslistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IActslistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IActslistContext)
}

func (s *PruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		listenerT.EnterPrule(s)
	}
}

func (s *PruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		listenerT.ExitPrule(s)
	}
}

func (p *EcaruleParser) Prule() (localctx IPruleContext) {
	localctx = NewPruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EcaruleParserRULE_prule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(EcaruleParserRULE)
	}
	{
		p.SetState(77)
		p.Match(EcaruleParserSIMPLENAME)
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EcaruleParserIN {
		{
			p.SetState(78)
			p.Match(EcaruleParserIN)
		}
		{
			p.SetState(79)
			p.Match(EcaruleParserSIMPLENAME)
		}

	}
	{
		p.SetState(82)
		p.Match(EcaruleParserON)
	}
	{
		p.SetState(83)
		p.Evt()
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EcaruleParserDEFAULT {
		{
			p.SetState(84)
			p.Match(EcaruleParserDEFAULT)
		}
		{
			p.SetState(85)
			p.Actslist()
		}

	}
	{
		p.SetState(88)
		p.Task()
	}

	return localctx
}

// IEvtContext is an interface to support dynamic dispatch.
type IEvtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEvtContext differentiates from other interfaces.
	IsEvtContext()
}

type EvtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEvtContext() *EvtContext {
	var p = new(EvtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_evt
	return p
}

func (*EvtContext) IsEvtContext() {}

func NewEvtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EvtContext {
	var p = new(EvtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_evt

	return p
}

func (s *EvtContext) GetParser() antlr.Parser { return s.parser }

func (s *EvtContext) AllSIMPLENAME() []antlr.TerminalNode {
	return s.GetTokens(EcaruleParserSIMPLENAME)
}

func (s *EvtContext) SIMPLENAME(i int) antlr.TerminalNode {
	return s.GetToken(EcaruleParserSIMPLENAME, i)
}

func (s *EvtContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(EcaruleParserSEMICOLON)
}

func (s *EvtContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(EcaruleParserSEMICOLON, i)
}

func (s *EvtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EvtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EvtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		listenerT.EnterEvt(s)
	}
}

func (s *EvtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		listenerT.ExitEvt(s)
	}
}

func (p *EcaruleParser) Evt() (localctx IEvtContext) {
	localctx = NewEvtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EcaruleParserRULE_evt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(EcaruleParserSIMPLENAME)
	}
	{
		p.SetState(91)
		p.Match(EcaruleParserSEMICOLON)
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EcaruleParserSIMPLENAME {
		{
			p.SetState(92)
			p.Match(EcaruleParserSIMPLENAME)
		}
		{
			p.SetState(93)
			p.Match(EcaruleParserSEMICOLON)
		}

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITaskContext is an interface to support dynamic dispatch.
type ITaskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTaskContext differentiates from other interfaces.
	IsTaskContext()
}

type TaskContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTaskContext() *TaskContext {
	var p = new(TaskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_task
	return p
}

func (*TaskContext) IsTaskContext() {}

func NewTaskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TaskContext {
	var p = new(TaskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_task

	return p
}

func (s *TaskContext) GetParser() antlr.Parser { return s.parser }

func (s *TaskContext) FOR() antlr.TerminalNode {
	return s.GetToken(EcaruleParserFOR, 0)
}

func (s *TaskContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TaskContext) DO() antlr.TerminalNode {
	return s.GetToken(EcaruleParserDO, 0)
}

func (s *TaskContext) Actslist() IActslistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IActslistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IActslistContext)
}

func (s *TaskContext) SOME() antlr.TerminalNode {
	return s.GetToken(EcaruleParserSOME, 0)
}

func (s *TaskContext) ALL() antlr.TerminalNode {
	return s.GetToken(EcaruleParserALL, 0)
}

func (s *TaskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TaskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TaskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		listenerT.EnterTask(s)
	}
}

func (s *TaskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		listenerT.ExitTask(s)
	}
}

func (p *EcaruleParser) Task() (localctx ITaskContext) {
	localctx = NewTaskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, EcaruleParserRULE_task)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(EcaruleParserFOR)
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EcaruleParserALL || _la == EcaruleParserSOME {
		{
			p.SetState(100)
			_la = p.GetTokenStream().LA(1)

			if !(_la == EcaruleParserALL || _la == EcaruleParserSOME) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(103)
		p.expression(0)
	}
	{
		p.SetState(104)
		p.Match(EcaruleParserDO)
	}
	{
		p.SetState(105)
		p.Actslist()
	}

	return localctx
}

// IActslistContext is an interface to support dynamic dispatch.
type IActslistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActslistContext differentiates from other interfaces.
	IsActslistContext()
}

type ActslistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActslistContext() *ActslistContext {
	var p = new(ActslistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_actslist
	return p
}

func (*ActslistContext) IsActslistContext() {}

func NewActslistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActslistContext {
	var p = new(ActslistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_actslist

	return p
}

func (s *ActslistContext) GetParser() antlr.Parser { return s.parser }

func (s *ActslistContext) AllAct() []IActContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IActContext)(nil)).Elem())
	var tst = make([]IActContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IActContext)
		}
	}

	return tst
}

func (s *ActslistContext) Act(i int) IActContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IActContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IActContext)
}

func (s *ActslistContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(EcaruleParserSEMICOLON)
}

func (s *ActslistContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(EcaruleParserSEMICOLON, i)
}

func (s *ActslistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActslistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActslistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		listenerT.EnterActslist(s)
	}
}

func (s *ActslistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		listenerT.ExitActslist(s)
	}
}

func (p *EcaruleParser) Actslist() (localctx IActslistContext) {
	localctx = NewActslistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EcaruleParserRULE_actslist)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Act()
	}
	{
		p.SetState(108)
		p.Match(EcaruleParserSEMICOLON)
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EcaruleParserSIMPLENAME {
		{
			p.SetState(109)
			p.Act()
		}
		{
			p.SetState(110)
			p.Match(EcaruleParserSEMICOLON)
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IActContext is an interface to support dynamic dispatch.
type IActContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActContext differentiates from other interfaces.
	IsActContext()
}

type ActContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActContext() *ActContext {
	var p = new(ActContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_act
	return p
}

func (*ActContext) IsActContext() {}

func NewActContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActContext {
	var p = new(ActContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_act

	return p
}

func (s *ActContext) GetParser() antlr.Parser { return s.parser }

func (s *ActContext) SIMPLENAME() antlr.TerminalNode {
	return s.GetToken(EcaruleParserSIMPLENAME, 0)
}

func (s *ActContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(EcaruleParserASSIGN, 0)
}

func (s *ActContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ActContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		listenerT.EnterAct(s)
	}
}

func (s *ActContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		listenerT.ExitAct(s)
	}
}

func (p *EcaruleParser) Act() (localctx IActContext) {
	localctx = NewActContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, EcaruleParserRULE_act)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(EcaruleParserSIMPLENAME)
	}
	{
		p.SetState(118)
		p.Match(EcaruleParserASSIGN)
	}
	{
		p.SetState(119)
		p.expression(0)
	}

	return localctx
}

// IGrlContext is an interface to support dynamic dispatch.
type IGrlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGrlContext differentiates from other interfaces.
	IsGrlContext()
}

type GrlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGrlContext() *GrlContext {
	var p = new(GrlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_grl
	return p
}

func (*GrlContext) IsGrlContext() {}

func NewGrlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GrlContext {
	var p = new(GrlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_grl

	return p
}

func (s *GrlContext) GetParser() antlr.Parser { return s.parser }

func (s *GrlContext) EOF() antlr.TerminalNode {
	return s.GetToken(EcaruleParserEOF, 0)
}

func (s *GrlContext) AllRuleEntry() []IRuleEntryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRuleEntryContext)(nil)).Elem())
	var tst = make([]IRuleEntryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRuleEntryContext)
		}
	}

	return tst
}

func (s *GrlContext) RuleEntry(i int) IRuleEntryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleEntryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRuleEntryContext)
}

func (s *GrlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GrlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GrlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewGrlContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterGrl(c)
	}
}

func (s *GrlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewGrlContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitGrl(c)
	}
}

func (p *EcaruleParser) Grl() (localctx IGrlContext) {
	localctx = NewGrlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, EcaruleParserRULE_grl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EcaruleParserRULE {
		{
			p.SetState(121)
			p.RuleEntry()
		}

		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(127)
		p.Match(EcaruleParserEOF)
	}

	return localctx
}

// IRuleEntryContext is an interface to support dynamic dispatch.
type IRuleEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleEntryContext differentiates from other interfaces.
	IsRuleEntryContext()
}

type RuleEntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleEntryContext() *RuleEntryContext {
	var p = new(RuleEntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_ruleEntry
	return p
}

func (*RuleEntryContext) IsRuleEntryContext() {}

func NewRuleEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleEntryContext {
	var p = new(RuleEntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_ruleEntry

	return p
}

func (s *RuleEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleEntryContext) RULE() antlr.TerminalNode {
	return s.GetToken(EcaruleParserRULE, 0)
}

func (s *RuleEntryContext) RuleName() IRuleNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleNameContext)
}

func (s *RuleEntryContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(EcaruleParserLR_BRACE, 0)
}

func (s *RuleEntryContext) WhenScope() IWhenScopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhenScopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhenScopeContext)
}

func (s *RuleEntryContext) ThenScope() IThenScopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThenScopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IThenScopeContext)
}

func (s *RuleEntryContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(EcaruleParserRR_BRACE, 0)
}

func (s *RuleEntryContext) RuleDescription() IRuleDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleDescriptionContext)
}

func (s *RuleEntryContext) Salience() ISalienceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISalienceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISalienceContext)
}

func (s *RuleEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewRuleEntryContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterRuleEntry(c)
	}
}

func (s *RuleEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewRuleEntryContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitRuleEntry(c)
	}
}

func (p *EcaruleParser) RuleEntry() (localctx IRuleEntryContext) {
	localctx = NewRuleEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, EcaruleParserRULE_ruleEntry)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Match(EcaruleParserRULE)
	}
	{
		p.SetState(130)
		p.RuleName()
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EcaruleParserDQUOTA_STRING || _la == EcaruleParserSQUOTA_STRING {
		{
			p.SetState(131)
			p.RuleDescription()
		}

	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EcaruleParserSALIENCE {
		{
			p.SetState(134)
			p.Salience()
		}

	}
	{
		p.SetState(137)
		p.Match(EcaruleParserLR_BRACE)
	}
	{
		p.SetState(138)
		p.WhenScope()
	}
	{
		p.SetState(139)
		p.ThenScope()
	}
	{
		p.SetState(140)
		p.Match(EcaruleParserRR_BRACE)
	}

	return localctx
}

// ISalienceContext is an interface to support dynamic dispatch.
type ISalienceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSalienceContext differentiates from other interfaces.
	IsSalienceContext()
}

type SalienceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySalienceContext() *SalienceContext {
	var p = new(SalienceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_salience
	return p
}

func (*SalienceContext) IsSalienceContext() {}

func NewSalienceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SalienceContext {
	var p = new(SalienceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_salience

	return p
}

func (s *SalienceContext) GetParser() antlr.Parser { return s.parser }

func (s *SalienceContext) SALIENCE() antlr.TerminalNode {
	return s.GetToken(EcaruleParserSALIENCE, 0)
}

func (s *SalienceContext) IntegerLiteral() IIntegerLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerLiteralContext)
}

func (s *SalienceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SalienceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SalienceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewSalienceContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterSalience(c)
	}
}

func (s *SalienceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewSalienceContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitSalience(c)
	}
}

func (p *EcaruleParser) Salience() (localctx ISalienceContext) {
	localctx = NewSalienceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, EcaruleParserRULE_salience)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(EcaruleParserSALIENCE)
	}
	{
		p.SetState(143)
		p.IntegerLiteral()
	}

	return localctx
}

// IRuleNameContext is an interface to support dynamic dispatch.
type IRuleNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleNameContext differentiates from other interfaces.
	IsRuleNameContext()
}

type RuleNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleNameContext() *RuleNameContext {
	var p = new(RuleNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_ruleName
	return p
}

func (*RuleNameContext) IsRuleNameContext() {}

func NewRuleNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleNameContext {
	var p = new(RuleNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_ruleName

	return p
}

func (s *RuleNameContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleNameContext) SIMPLENAME() antlr.TerminalNode {
	return s.GetToken(EcaruleParserSIMPLENAME, 0)
}

func (s *RuleNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewRuleNameContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterRuleName(c)
	}
}

func (s *RuleNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewRuleNameContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitRuleName(c)
	}
}

func (p *EcaruleParser) RuleName() (localctx IRuleNameContext) {
	localctx = NewRuleNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, EcaruleParserRULE_ruleName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(EcaruleParserSIMPLENAME)
	}

	return localctx
}

// IRuleDescriptionContext is an interface to support dynamic dispatch.
type IRuleDescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleDescriptionContext differentiates from other interfaces.
	IsRuleDescriptionContext()
}

type RuleDescriptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleDescriptionContext() *RuleDescriptionContext {
	var p = new(RuleDescriptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_ruleDescription
	return p
}

func (*RuleDescriptionContext) IsRuleDescriptionContext() {}

func NewRuleDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleDescriptionContext {
	var p = new(RuleDescriptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_ruleDescription

	return p
}

func (s *RuleDescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleDescriptionContext) DQUOTA_STRING() antlr.TerminalNode {
	return s.GetToken(EcaruleParserDQUOTA_STRING, 0)
}

func (s *RuleDescriptionContext) SQUOTA_STRING() antlr.TerminalNode {
	return s.GetToken(EcaruleParserSQUOTA_STRING, 0)
}

func (s *RuleDescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleDescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleDescriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewRuleDescriptionContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterRuleDescription(c)
	}
}

func (s *RuleDescriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewRuleDescriptionContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitRuleDescription(c)
	}
}

func (p *EcaruleParser) RuleDescription() (localctx IRuleDescriptionContext) {
	localctx = NewRuleDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, EcaruleParserRULE_ruleDescription)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		_la = p.GetTokenStream().LA(1)

		if !(_la == EcaruleParserDQUOTA_STRING || _la == EcaruleParserSQUOTA_STRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IWhenScopeContext is an interface to support dynamic dispatch.
type IWhenScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhenScopeContext differentiates from other interfaces.
	IsWhenScopeContext()
}

type WhenScopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhenScopeContext() *WhenScopeContext {
	var p = new(WhenScopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_whenScope
	return p
}

func (*WhenScopeContext) IsWhenScopeContext() {}

func NewWhenScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhenScopeContext {
	var p = new(WhenScopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_whenScope

	return p
}

func (s *WhenScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *WhenScopeContext) WHEN() antlr.TerminalNode {
	return s.GetToken(EcaruleParserWHEN, 0)
}

func (s *WhenScopeContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhenScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhenScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhenScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewWhenScopeContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterWhenScope(c)
	}
}

func (s *WhenScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewWhenScopeContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitWhenScope(c)
	}
}

func (p *EcaruleParser) WhenScope() (localctx IWhenScopeContext) {
	localctx = NewWhenScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, EcaruleParserRULE_whenScope)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(EcaruleParserWHEN)
	}
	{
		p.SetState(150)
		p.expression(0)
	}

	return localctx
}

// IThenScopeContext is an interface to support dynamic dispatch.
type IThenScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThenScopeContext differentiates from other interfaces.
	IsThenScopeContext()
}

type ThenScopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThenScopeContext() *ThenScopeContext {
	var p = new(ThenScopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_thenScope
	return p
}

func (*ThenScopeContext) IsThenScopeContext() {}

func NewThenScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThenScopeContext {
	var p = new(ThenScopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_thenScope

	return p
}

func (s *ThenScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *ThenScopeContext) THEN() antlr.TerminalNode {
	return s.GetToken(EcaruleParserTHEN, 0)
}

func (s *ThenScopeContext) ThenExpressionList() IThenExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThenExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IThenExpressionListContext)
}

func (s *ThenScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThenScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThenScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewThenScopeContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterThenScope(c)
	}
}

func (s *ThenScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewThenScopeContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitThenScope(c)
	}
}

func (p *EcaruleParser) ThenScope() (localctx IThenScopeContext) {
	localctx = NewThenScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, EcaruleParserRULE_thenScope)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Match(EcaruleParserTHEN)
	}
	{
		p.SetState(153)
		p.ThenExpressionList()
	}

	return localctx
}

// IThenExpressionListContext is an interface to support dynamic dispatch.
type IThenExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThenExpressionListContext differentiates from other interfaces.
	IsThenExpressionListContext()
}

type ThenExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThenExpressionListContext() *ThenExpressionListContext {
	var p = new(ThenExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_thenExpressionList
	return p
}

func (*ThenExpressionListContext) IsThenExpressionListContext() {}

func NewThenExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThenExpressionListContext {
	var p = new(ThenExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_thenExpressionList

	return p
}

func (s *ThenExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ThenExpressionListContext) AllThenExpression() []IThenExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IThenExpressionContext)(nil)).Elem())
	var tst = make([]IThenExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IThenExpressionContext)
		}
	}

	return tst
}

func (s *ThenExpressionListContext) ThenExpression(i int) IThenExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThenExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IThenExpressionContext)
}

func (s *ThenExpressionListContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(EcaruleParserSEMICOLON)
}

func (s *ThenExpressionListContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(EcaruleParserSEMICOLON, i)
}

func (s *ThenExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThenExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThenExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewThenExpressionListContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterThenExpressionList(c)
	}
}

func (s *ThenExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewThenExpressionListContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitThenExpressionList(c)
	}
}

func (p *EcaruleParser) ThenExpressionList() (localctx IThenExpressionListContext) {
	localctx = NewThenExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, EcaruleParserRULE_thenExpressionList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<EcaruleParserMINUS)|(1<<EcaruleParserTRUE)|(1<<EcaruleParserFALSE)|(1<<EcaruleParserNIL_LITERAL)|(1<<EcaruleParserNEGATION))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(EcaruleParserSIMPLENAME-38))|(1<<(EcaruleParserDQUOTA_STRING-38))|(1<<(EcaruleParserSQUOTA_STRING-38))|(1<<(EcaruleParserDECIMAL_FLOAT_LIT-38))|(1<<(EcaruleParserHEX_FLOAT_LIT-38))|(1<<(EcaruleParserDEC_LIT-38))|(1<<(EcaruleParserHEX_LIT-38))|(1<<(EcaruleParserOCT_LIT-38)))) != 0) {
		{
			p.SetState(155)
			p.ThenExpression()
		}
		{
			p.SetState(156)
			p.Match(EcaruleParserSEMICOLON)
		}

		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IThenExpressionContext is an interface to support dynamic dispatch.
type IThenExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThenExpressionContext differentiates from other interfaces.
	IsThenExpressionContext()
}

type ThenExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThenExpressionContext() *ThenExpressionContext {
	var p = new(ThenExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_thenExpression
	return p
}

func (*ThenExpressionContext) IsThenExpressionContext() {}

func NewThenExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThenExpressionContext {
	var p = new(ThenExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_thenExpression

	return p
}

func (s *ThenExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ThenExpressionContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ThenExpressionContext) ExpressionAtom() IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *ThenExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThenExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThenExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewThenExpressionContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterThenExpression(c)
	}
}

func (s *ThenExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewThenExpressionContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitThenExpression(c)
	}
}

func (p *EcaruleParser) ThenExpression() (localctx IThenExpressionContext) {
	localctx = NewThenExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, EcaruleParserRULE_thenExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(163)
			p.expressionAtom(0)
		}

	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(EcaruleParserASSIGN, 0)
}

func (s *AssignmentContext) PLUS_ASIGN() antlr.TerminalNode {
	return s.GetToken(EcaruleParserPLUS_ASIGN, 0)
}

func (s *AssignmentContext) MINUS_ASIGN() antlr.TerminalNode {
	return s.GetToken(EcaruleParserMINUS_ASIGN, 0)
}

func (s *AssignmentContext) DIV_ASIGN() antlr.TerminalNode {
	return s.GetToken(EcaruleParserDIV_ASIGN, 0)
}

func (s *AssignmentContext) MUL_ASIGN() antlr.TerminalNode {
	return s.GetToken(EcaruleParserMUL_ASIGN, 0)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewAssignmentContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterAssignment(c)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewAssignmentContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitAssignment(c)
	}
}

func (p *EcaruleParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, EcaruleParserRULE_assignment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.variable(0)
	}
	{
		p.SetState(167)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<EcaruleParserASSIGN)|(1<<EcaruleParserPLUS_ASIGN)|(1<<EcaruleParserMINUS_ASIGN)|(1<<EcaruleParserDIV_ASIGN)|(1<<EcaruleParserMUL_ASIGN))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(168)
		p.expression(0)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(EcaruleParserLR_BRACKET, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(EcaruleParserRR_BRACKET, 0)
}

func (s *ExpressionContext) NEGATION() antlr.TerminalNode {
	return s.GetToken(EcaruleParserNEGATION, 0)
}

func (s *ExpressionContext) ExpressionAtom() IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *ExpressionContext) MulDivOperators() IMulDivOperatorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMulDivOperatorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMulDivOperatorsContext)
}

func (s *ExpressionContext) AddMinusOperators() IAddMinusOperatorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddMinusOperatorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAddMinusOperatorsContext)
}

func (s *ExpressionContext) ComparisonOperator() IComparisonOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonOperatorContext)
}

func (s *ExpressionContext) AndLogicOperator() IAndLogicOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndLogicOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAndLogicOperatorContext)
}

func (s *ExpressionContext) OrLogicOperator() IOrLogicOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrLogicOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrLogicOperatorContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewExpressionContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterExpression(c)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewExpressionContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitExpression(c)
	}
}

func (p *EcaruleParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *EcaruleParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, EcaruleParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EcaruleParserNEGATION {
			{
				p.SetState(171)
				p.Match(EcaruleParserNEGATION)
			}

		}
		{
			p.SetState(174)
			p.Match(EcaruleParserLR_BRACKET)
		}
		{
			p.SetState(175)
			p.expression(0)
		}
		{
			p.SetState(176)
			p.Match(EcaruleParserRR_BRACKET)
		}

	case 2:
		{
			p.SetState(178)
			p.expressionAtom(0)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(201)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, EcaruleParserRULE_expression)
				p.SetState(181)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(182)
					p.MulDivOperators()
				}
				{
					p.SetState(183)
					p.expression(8)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, EcaruleParserRULE_expression)
				p.SetState(185)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(186)
					p.AddMinusOperators()
				}
				{
					p.SetState(187)
					p.expression(7)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, EcaruleParserRULE_expression)
				p.SetState(189)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(190)
					p.ComparisonOperator()
				}
				{
					p.SetState(191)
					p.expression(6)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, EcaruleParserRULE_expression)
				p.SetState(193)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(194)
					p.AndLogicOperator()
				}
				{
					p.SetState(195)
					p.expression(5)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, EcaruleParserRULE_expression)
				p.SetState(197)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(198)
					p.OrLogicOperator()
				}
				{
					p.SetState(199)
					p.expression(4)
				}

			}

		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}

	return localctx
}

// IMulDivOperatorsContext is an interface to support dynamic dispatch.
type IMulDivOperatorsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMulDivOperatorsContext differentiates from other interfaces.
	IsMulDivOperatorsContext()
}

type MulDivOperatorsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulDivOperatorsContext() *MulDivOperatorsContext {
	var p = new(MulDivOperatorsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_mulDivOperators
	return p
}

func (*MulDivOperatorsContext) IsMulDivOperatorsContext() {}

func NewMulDivOperatorsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulDivOperatorsContext {
	var p = new(MulDivOperatorsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_mulDivOperators

	return p
}

func (s *MulDivOperatorsContext) GetParser() antlr.Parser { return s.parser }

func (s *MulDivOperatorsContext) MUL() antlr.TerminalNode {
	return s.GetToken(EcaruleParserMUL, 0)
}

func (s *MulDivOperatorsContext) DIV() antlr.TerminalNode {
	return s.GetToken(EcaruleParserDIV, 0)
}

func (s *MulDivOperatorsContext) MOD() antlr.TerminalNode {
	return s.GetToken(EcaruleParserMOD, 0)
}

func (s *MulDivOperatorsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivOperatorsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MulDivOperatorsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewMulDivOperatorsContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterMulDivOperators(c)
	}
}

func (s *MulDivOperatorsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewMulDivOperatorsContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitMulDivOperators(c)
	}
}

func (p *EcaruleParser) MulDivOperators() (localctx IMulDivOperatorsContext) {
	localctx = NewMulDivOperatorsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, EcaruleParserRULE_mulDivOperators)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<EcaruleParserDIV)|(1<<EcaruleParserMUL)|(1<<EcaruleParserMOD))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAddMinusOperatorsContext is an interface to support dynamic dispatch.
type IAddMinusOperatorsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddMinusOperatorsContext differentiates from other interfaces.
	IsAddMinusOperatorsContext()
}

type AddMinusOperatorsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddMinusOperatorsContext() *AddMinusOperatorsContext {
	var p = new(AddMinusOperatorsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_addMinusOperators
	return p
}

func (*AddMinusOperatorsContext) IsAddMinusOperatorsContext() {}

func NewAddMinusOperatorsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddMinusOperatorsContext {
	var p = new(AddMinusOperatorsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_addMinusOperators

	return p
}

func (s *AddMinusOperatorsContext) GetParser() antlr.Parser { return s.parser }

func (s *AddMinusOperatorsContext) PLUS() antlr.TerminalNode {
	return s.GetToken(EcaruleParserPLUS, 0)
}

func (s *AddMinusOperatorsContext) MINUS() antlr.TerminalNode {
	return s.GetToken(EcaruleParserMINUS, 0)
}

func (s *AddMinusOperatorsContext) BITAND() antlr.TerminalNode {
	return s.GetToken(EcaruleParserBITAND, 0)
}

func (s *AddMinusOperatorsContext) BITOR() antlr.TerminalNode {
	return s.GetToken(EcaruleParserBITOR, 0)
}

func (s *AddMinusOperatorsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddMinusOperatorsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddMinusOperatorsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewAddMinusOperatorsContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterAddMinusOperators(c)
	}
}

func (s *AddMinusOperatorsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewAddMinusOperatorsContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitAddMinusOperators(c)
	}
}

func (p *EcaruleParser) AddMinusOperators() (localctx IAddMinusOperatorsContext) {
	localctx = NewAddMinusOperatorsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, EcaruleParserRULE_addMinusOperators)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		_la = p.GetTokenStream().LA(1)

		if !(_la == EcaruleParserPLUS || _la == EcaruleParserMINUS || _la == EcaruleParserBITAND || _la == EcaruleParserBITOR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IComparisonOperatorContext is an interface to support dynamic dispatch.
type IComparisonOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonOperatorContext differentiates from other interfaces.
	IsComparisonOperatorContext()
}

type ComparisonOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOperatorContext() *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_comparisonOperator
	return p
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonOperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(EcaruleParserGT, 0)
}

func (s *ComparisonOperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(EcaruleParserLT, 0)
}

func (s *ComparisonOperatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(EcaruleParserGTE, 0)
}

func (s *ComparisonOperatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(EcaruleParserLTE, 0)
}

func (s *ComparisonOperatorContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(EcaruleParserEQUALS, 0)
}

func (s *ComparisonOperatorContext) NOTEQUALS() antlr.TerminalNode {
	return s.GetToken(EcaruleParserNOTEQUALS, 0)
}

func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewComparisonOperatorContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterComparisonOperator(c)
	}
}

func (s *ComparisonOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewComparisonOperatorContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitComparisonOperator(c)
	}
}

func (p *EcaruleParser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, EcaruleParserRULE_comparisonOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-25)&-(0x1f+1)) == 0 && ((1<<uint((_la-25)))&((1<<(EcaruleParserEQUALS-25))|(1<<(EcaruleParserGT-25))|(1<<(EcaruleParserLT-25))|(1<<(EcaruleParserGTE-25))|(1<<(EcaruleParserLTE-25))|(1<<(EcaruleParserNOTEQUALS-25)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAndLogicOperatorContext is an interface to support dynamic dispatch.
type IAndLogicOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAndLogicOperatorContext differentiates from other interfaces.
	IsAndLogicOperatorContext()
}

type AndLogicOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndLogicOperatorContext() *AndLogicOperatorContext {
	var p = new(AndLogicOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_andLogicOperator
	return p
}

func (*AndLogicOperatorContext) IsAndLogicOperatorContext() {}

func NewAndLogicOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndLogicOperatorContext {
	var p = new(AndLogicOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_andLogicOperator

	return p
}

func (s *AndLogicOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AndLogicOperatorContext) AND() antlr.TerminalNode {
	return s.GetToken(EcaruleParserAND, 0)
}

func (s *AndLogicOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndLogicOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndLogicOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewAndLogicOperatorContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterAndLogicOperator(c)
	}
}

func (s *AndLogicOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewAndLogicOperatorContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitAndLogicOperator(c)
	}
}

func (p *EcaruleParser) AndLogicOperator() (localctx IAndLogicOperatorContext) {
	localctx = NewAndLogicOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, EcaruleParserRULE_andLogicOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(EcaruleParserAND)
	}

	return localctx
}

// IOrLogicOperatorContext is an interface to support dynamic dispatch.
type IOrLogicOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrLogicOperatorContext differentiates from other interfaces.
	IsOrLogicOperatorContext()
}

type OrLogicOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrLogicOperatorContext() *OrLogicOperatorContext {
	var p = new(OrLogicOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_orLogicOperator
	return p
}

func (*OrLogicOperatorContext) IsOrLogicOperatorContext() {}

func NewOrLogicOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrLogicOperatorContext {
	var p = new(OrLogicOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_orLogicOperator

	return p
}

func (s *OrLogicOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OrLogicOperatorContext) OR() antlr.TerminalNode {
	return s.GetToken(EcaruleParserOR, 0)
}

func (s *OrLogicOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrLogicOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrLogicOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewOrLogicOperatorContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterOrLogicOperator(c)
	}
}

func (s *OrLogicOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewOrLogicOperatorContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitOrLogicOperator(c)
	}
}

func (p *EcaruleParser) OrLogicOperator() (localctx IOrLogicOperatorContext) {
	localctx = NewOrLogicOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, EcaruleParserRULE_orLogicOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(EcaruleParserOR)
	}

	return localctx
}

// IExpressionAtomContext is an interface to support dynamic dispatch.
type IExpressionAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionAtomContext differentiates from other interfaces.
	IsExpressionAtomContext()
}

type ExpressionAtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionAtomContext() *ExpressionAtomContext {
	var p = new(ExpressionAtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_expressionAtom
	return p
}

func (*ExpressionAtomContext) IsExpressionAtomContext() {}

func NewExpressionAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionAtomContext {
	var p = new(ExpressionAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_expressionAtom

	return p
}

func (s *ExpressionAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionAtomContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ExpressionAtomContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ExpressionAtomContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ExpressionAtomContext) NEGATION() antlr.TerminalNode {
	return s.GetToken(EcaruleParserNEGATION, 0)
}

func (s *ExpressionAtomContext) ExpressionAtom() IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *ExpressionAtomContext) MethodCall() IMethodCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *ExpressionAtomContext) MemberVariable() IMemberVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberVariableContext)
}

func (s *ExpressionAtomContext) ArrayMapSelector() IArrayMapSelectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayMapSelectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayMapSelectorContext)
}

func (s *ExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewExpressionAtomContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterExpressionAtom(c)
	}
}

func (s *ExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewExpressionAtomContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitExpressionAtom(c)
	}
}

func (p *EcaruleParser) ExpressionAtom() (localctx IExpressionAtomContext) {
	return p.expressionAtom(0)
}

func (p *EcaruleParser) expressionAtom(_p int) (localctx IExpressionAtomContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionAtomContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionAtomContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, EcaruleParserRULE_expressionAtom, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(217)
			p.Constant()
		}

	case 2:
		{
			p.SetState(218)
			p.variable(0)
		}

	case 3:
		{
			p.SetState(219)
			p.FunctionCall()
		}

	case 4:
		{
			p.SetState(220)
			p.Match(EcaruleParserNEGATION)
		}
		{
			p.SetState(221)
			p.expressionAtom(1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(230)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionAtomContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, EcaruleParserRULE_expressionAtom)
				p.SetState(224)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(225)
					p.MethodCall()
				}

			case 2:
				localctx = NewExpressionAtomContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, EcaruleParserRULE_expressionAtom)
				p.SetState(226)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(227)
					p.MemberVariable()
				}

			case 3:
				localctx = NewExpressionAtomContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, EcaruleParserRULE_expressionAtom)
				p.SetState(228)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(229)
					p.ArrayMapSelector()
				}

			}

		}
		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *ConstantContext) IntegerLiteral() IIntegerLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerLiteralContext)
}

func (s *ConstantContext) FloatLiteral() IFloatLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatLiteralContext)
}

func (s *ConstantContext) BooleanLiteral() IBooleanLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *ConstantContext) NIL_LITERAL() antlr.TerminalNode {
	return s.GetToken(EcaruleParserNIL_LITERAL, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewConstantContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterConstant(c)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewConstantContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitConstant(c)
	}
}

func (p *EcaruleParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, EcaruleParserRULE_constant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(235)
			p.StringLiteral()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(236)
			p.IntegerLiteral()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(237)
			p.FloatLiteral()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(238)
			p.BooleanLiteral()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(239)
			p.Match(EcaruleParserNIL_LITERAL)
		}

	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) SIMPLENAME() antlr.TerminalNode {
	return s.GetToken(EcaruleParserSIMPLENAME, 0)
}

func (s *VariableContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VariableContext) MemberVariable() IMemberVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberVariableContext)
}

func (s *VariableContext) ArrayMapSelector() IArrayMapSelectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayMapSelectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayMapSelectorContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewVariableContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterVariable(c)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewVariableContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitVariable(c)
	}
}

func (p *EcaruleParser) Variable() (localctx IVariableContext) {
	return p.variable(0)
}

func (p *EcaruleParser) variable(_p int) (localctx IVariableContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewVariableContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IVariableContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 46
	p.EnterRecursionRule(localctx, 46, EcaruleParserRULE_variable, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(EcaruleParserSIMPLENAME)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(249)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
			case 1:
				localctx = NewVariableContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, EcaruleParserRULE_variable)
				p.SetState(245)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(246)
					p.MemberVariable()
				}

			case 2:
				localctx = NewVariableContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, EcaruleParserRULE_variable)
				p.SetState(247)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(248)
					p.ArrayMapSelector()
				}

			}

		}
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}

	return localctx
}

// IArrayMapSelectorContext is an interface to support dynamic dispatch.
type IArrayMapSelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayMapSelectorContext differentiates from other interfaces.
	IsArrayMapSelectorContext()
}

type ArrayMapSelectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayMapSelectorContext() *ArrayMapSelectorContext {
	var p = new(ArrayMapSelectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_arrayMapSelector
	return p
}

func (*ArrayMapSelectorContext) IsArrayMapSelectorContext() {}

func NewArrayMapSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayMapSelectorContext {
	var p = new(ArrayMapSelectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_arrayMapSelector

	return p
}

func (s *ArrayMapSelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayMapSelectorContext) LS_BRACKET() antlr.TerminalNode {
	return s.GetToken(EcaruleParserLS_BRACKET, 0)
}

func (s *ArrayMapSelectorContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayMapSelectorContext) RS_BRACKET() antlr.TerminalNode {
	return s.GetToken(EcaruleParserRS_BRACKET, 0)
}

func (s *ArrayMapSelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayMapSelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayMapSelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewArrayMapSelectorContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterArrayMapSelector(c)
	}
}

func (s *ArrayMapSelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewArrayMapSelectorContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitArrayMapSelector(c)
	}
}

func (p *EcaruleParser) ArrayMapSelector() (localctx IArrayMapSelectorContext) {
	localctx = NewArrayMapSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, EcaruleParserRULE_arrayMapSelector)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		p.Match(EcaruleParserLS_BRACKET)
	}
	{
		p.SetState(255)
		p.expression(0)
	}
	{
		p.SetState(256)
		p.Match(EcaruleParserRS_BRACKET)
	}

	return localctx
}

// IMemberVariableContext is an interface to support dynamic dispatch.
type IMemberVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberVariableContext differentiates from other interfaces.
	IsMemberVariableContext()
}

type MemberVariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberVariableContext() *MemberVariableContext {
	var p = new(MemberVariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_memberVariable
	return p
}

func (*MemberVariableContext) IsMemberVariableContext() {}

func NewMemberVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberVariableContext {
	var p = new(MemberVariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_memberVariable

	return p
}

func (s *MemberVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberVariableContext) DOT() antlr.TerminalNode {
	return s.GetToken(EcaruleParserDOT, 0)
}

func (s *MemberVariableContext) SIMPLENAME() antlr.TerminalNode {
	return s.GetToken(EcaruleParserSIMPLENAME, 0)
}

func (s *MemberVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewMemberVariableContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterMemberVariable(c)
	}
}

func (s *MemberVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewMemberVariableContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitMemberVariable(c)
	}
}

func (p *EcaruleParser) MemberVariable() (localctx IMemberVariableContext) {
	localctx = NewMemberVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, EcaruleParserRULE_memberVariable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
		p.Match(EcaruleParserDOT)
	}
	{
		p.SetState(259)
		p.Match(EcaruleParserSIMPLENAME)
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) SIMPLENAME() antlr.TerminalNode {
	return s.GetToken(EcaruleParserSIMPLENAME, 0)
}

func (s *FunctionCallContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(EcaruleParserLR_BRACKET, 0)
}

func (s *FunctionCallContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(EcaruleParserRR_BRACKET, 0)
}

func (s *FunctionCallContext) ArgumentList() IArgumentListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewFunctionCallContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterFunctionCall(c)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewFunctionCallContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitFunctionCall(c)
	}
}

func (p *EcaruleParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, EcaruleParserRULE_functionCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.Match(EcaruleParserSIMPLENAME)
	}
	{
		p.SetState(262)
		p.Match(EcaruleParserLR_BRACKET)
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<EcaruleParserMINUS)|(1<<EcaruleParserLR_BRACKET)|(1<<EcaruleParserTRUE)|(1<<EcaruleParserFALSE)|(1<<EcaruleParserNIL_LITERAL)|(1<<EcaruleParserNEGATION))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(EcaruleParserSIMPLENAME-38))|(1<<(EcaruleParserDQUOTA_STRING-38))|(1<<(EcaruleParserSQUOTA_STRING-38))|(1<<(EcaruleParserDECIMAL_FLOAT_LIT-38))|(1<<(EcaruleParserHEX_FLOAT_LIT-38))|(1<<(EcaruleParserDEC_LIT-38))|(1<<(EcaruleParserHEX_LIT-38))|(1<<(EcaruleParserOCT_LIT-38)))) != 0) {
		{
			p.SetState(263)
			p.ArgumentList()
		}

	}
	{
		p.SetState(266)
		p.Match(EcaruleParserRR_BRACKET)
	}

	return localctx
}

// IMethodCallContext is an interface to support dynamic dispatch.
type IMethodCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodCallContext differentiates from other interfaces.
	IsMethodCallContext()
}

type MethodCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodCallContext() *MethodCallContext {
	var p = new(MethodCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_methodCall
	return p
}

func (*MethodCallContext) IsMethodCallContext() {}

func NewMethodCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodCallContext {
	var p = new(MethodCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_methodCall

	return p
}

func (s *MethodCallContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodCallContext) DOT() antlr.TerminalNode {
	return s.GetToken(EcaruleParserDOT, 0)
}

func (s *MethodCallContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *MethodCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewMethodCallContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterMethodCall(c)
	}
}

func (s *MethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewMethodCallContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitMethodCall(c)
	}
}

func (p *EcaruleParser) MethodCall() (localctx IMethodCallContext) {
	localctx = NewMethodCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, EcaruleParserRULE_methodCall)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Match(EcaruleParserDOT)
	}
	{
		p.SetState(269)
		p.FunctionCall()
	}

	return localctx
}

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_argumentList
	return p
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ArgumentListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewArgumentListContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterArgumentList(c)
	}
}

func (s *ArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewArgumentListContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitArgumentList(c)
	}
}

func (p *EcaruleParser) ArgumentList() (localctx IArgumentListContext) {
	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, EcaruleParserRULE_argumentList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(271)
		p.expression(0)
	}
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EcaruleParserT__0 {
		{
			p.SetState(272)
			p.Match(EcaruleParserT__0)
		}
		{
			p.SetState(273)
			p.expression(0)
		}

		p.SetState(278)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFloatLiteralContext is an interface to support dynamic dispatch.
type IFloatLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatLiteralContext differentiates from other interfaces.
	IsFloatLiteralContext()
}

type FloatLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatLiteralContext() *FloatLiteralContext {
	var p = new(FloatLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_floatLiteral
	return p
}

func (*FloatLiteralContext) IsFloatLiteralContext() {}

func NewFloatLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_floatLiteral

	return p
}

func (s *FloatLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatLiteralContext) DecimalFloatLiteral() IDecimalFloatLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalFloatLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecimalFloatLiteralContext)
}

func (s *FloatLiteralContext) HexadecimalFloatLiteral() IHexadecimalFloatLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHexadecimalFloatLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHexadecimalFloatLiteralContext)
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewFloatLiteralContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterFloatLiteral(c)
	}
}

func (s *FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewFloatLiteralContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitFloatLiteral(c)
	}
}

func (p *EcaruleParser) FloatLiteral() (localctx IFloatLiteralContext) {
	localctx = NewFloatLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, EcaruleParserRULE_floatLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(279)
			p.DecimalFloatLiteral()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(280)
			p.HexadecimalFloatLiteral()
		}

	}

	return localctx
}

// IDecimalFloatLiteralContext is an interface to support dynamic dispatch.
type IDecimalFloatLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecimalFloatLiteralContext differentiates from other interfaces.
	IsDecimalFloatLiteralContext()
}

type DecimalFloatLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecimalFloatLiteralContext() *DecimalFloatLiteralContext {
	var p = new(DecimalFloatLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_decimalFloatLiteral
	return p
}

func (*DecimalFloatLiteralContext) IsDecimalFloatLiteralContext() {}

func NewDecimalFloatLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecimalFloatLiteralContext {
	var p = new(DecimalFloatLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_decimalFloatLiteral

	return p
}

func (s *DecimalFloatLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *DecimalFloatLiteralContext) DECIMAL_FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(EcaruleParserDECIMAL_FLOAT_LIT, 0)
}

func (s *DecimalFloatLiteralContext) MINUS() antlr.TerminalNode {
	return s.GetToken(EcaruleParserMINUS, 0)
}

func (s *DecimalFloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalFloatLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecimalFloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewDecimalFloatLiteralContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterDecimalFloatLiteral(c)
	}
}

func (s *DecimalFloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewDecimalFloatLiteralContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitDecimalFloatLiteral(c)
	}
}

func (p *EcaruleParser) DecimalFloatLiteral() (localctx IDecimalFloatLiteralContext) {
	localctx = NewDecimalFloatLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, EcaruleParserRULE_decimalFloatLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EcaruleParserMINUS {
		{
			p.SetState(283)
			p.Match(EcaruleParserMINUS)
		}

	}
	{
		p.SetState(286)
		p.Match(EcaruleParserDECIMAL_FLOAT_LIT)
	}

	return localctx
}

// IHexadecimalFloatLiteralContext is an interface to support dynamic dispatch.
type IHexadecimalFloatLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHexadecimalFloatLiteralContext differentiates from other interfaces.
	IsHexadecimalFloatLiteralContext()
}

type HexadecimalFloatLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHexadecimalFloatLiteralContext() *HexadecimalFloatLiteralContext {
	var p = new(HexadecimalFloatLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_hexadecimalFloatLiteral
	return p
}

func (*HexadecimalFloatLiteralContext) IsHexadecimalFloatLiteralContext() {}

func NewHexadecimalFloatLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HexadecimalFloatLiteralContext {
	var p = new(HexadecimalFloatLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_hexadecimalFloatLiteral

	return p
}

func (s *HexadecimalFloatLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *HexadecimalFloatLiteralContext) HEX_FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(EcaruleParserHEX_FLOAT_LIT, 0)
}

func (s *HexadecimalFloatLiteralContext) MINUS() antlr.TerminalNode {
	return s.GetToken(EcaruleParserMINUS, 0)
}

func (s *HexadecimalFloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HexadecimalFloatLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HexadecimalFloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewHexadecimalFloatLiteralContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterHexadecimalFloatLiteral(c)
	}
}

func (s *HexadecimalFloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewHexadecimalFloatLiteralContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitHexadecimalFloatLiteral(c)
	}
}

func (p *EcaruleParser) HexadecimalFloatLiteral() (localctx IHexadecimalFloatLiteralContext) {
	localctx = NewHexadecimalFloatLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, EcaruleParserRULE_hexadecimalFloatLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EcaruleParserMINUS {
		{
			p.SetState(288)
			p.Match(EcaruleParserMINUS)
		}

	}
	{
		p.SetState(291)
		p.Match(EcaruleParserHEX_FLOAT_LIT)
	}

	return localctx
}

// IIntegerLiteralContext is an interface to support dynamic dispatch.
type IIntegerLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerLiteralContext differentiates from other interfaces.
	IsIntegerLiteralContext()
}

type IntegerLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerLiteralContext() *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_integerLiteral
	return p
}

func (*IntegerLiteralContext) IsIntegerLiteralContext() {}

func NewIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_integerLiteral

	return p
}

func (s *IntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerLiteralContext) DecimalLiteral() IDecimalLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecimalLiteralContext)
}

func (s *IntegerLiteralContext) HexadecimalLiteral() IHexadecimalLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHexadecimalLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHexadecimalLiteralContext)
}

func (s *IntegerLiteralContext) OctalLiteral() IOctalLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOctalLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOctalLiteralContext)
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewIntegerLiteralContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterIntegerLiteral(c)
	}
}

func (s *IntegerLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewIntegerLiteralContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitIntegerLiteral(c)
	}
}

func (p *EcaruleParser) IntegerLiteral() (localctx IIntegerLiteralContext) {
	localctx = NewIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, EcaruleParserRULE_integerLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(293)
			p.DecimalLiteral()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(294)
			p.HexadecimalLiteral()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(295)
			p.OctalLiteral()
		}

	}

	return localctx
}

// IDecimalLiteralContext is an interface to support dynamic dispatch.
type IDecimalLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecimalLiteralContext differentiates from other interfaces.
	IsDecimalLiteralContext()
}

type DecimalLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecimalLiteralContext() *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_decimalLiteral
	return p
}

func (*DecimalLiteralContext) IsDecimalLiteralContext() {}

func NewDecimalLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_decimalLiteral

	return p
}

func (s *DecimalLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *DecimalLiteralContext) DEC_LIT() antlr.TerminalNode {
	return s.GetToken(EcaruleParserDEC_LIT, 0)
}

func (s *DecimalLiteralContext) MINUS() antlr.TerminalNode {
	return s.GetToken(EcaruleParserMINUS, 0)
}

func (s *DecimalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecimalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewDecimalLiteralContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterDecimalLiteral(c)
	}
}

func (s *DecimalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewDecimalLiteralContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitDecimalLiteral(c)
	}
}

func (p *EcaruleParser) DecimalLiteral() (localctx IDecimalLiteralContext) {
	localctx = NewDecimalLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, EcaruleParserRULE_decimalLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EcaruleParserMINUS {
		{
			p.SetState(298)
			p.Match(EcaruleParserMINUS)
		}

	}
	{
		p.SetState(301)
		p.Match(EcaruleParserDEC_LIT)
	}

	return localctx
}

// IHexadecimalLiteralContext is an interface to support dynamic dispatch.
type IHexadecimalLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHexadecimalLiteralContext differentiates from other interfaces.
	IsHexadecimalLiteralContext()
}

type HexadecimalLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHexadecimalLiteralContext() *HexadecimalLiteralContext {
	var p = new(HexadecimalLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_hexadecimalLiteral
	return p
}

func (*HexadecimalLiteralContext) IsHexadecimalLiteralContext() {}

func NewHexadecimalLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HexadecimalLiteralContext {
	var p = new(HexadecimalLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_hexadecimalLiteral

	return p
}

func (s *HexadecimalLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *HexadecimalLiteralContext) HEX_LIT() antlr.TerminalNode {
	return s.GetToken(EcaruleParserHEX_LIT, 0)
}

func (s *HexadecimalLiteralContext) MINUS() antlr.TerminalNode {
	return s.GetToken(EcaruleParserMINUS, 0)
}

func (s *HexadecimalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HexadecimalLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HexadecimalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewHexadecimalLiteralContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterHexadecimalLiteral(c)
	}
}

func (s *HexadecimalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewHexadecimalLiteralContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitHexadecimalLiteral(c)
	}
}

func (p *EcaruleParser) HexadecimalLiteral() (localctx IHexadecimalLiteralContext) {
	localctx = NewHexadecimalLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, EcaruleParserRULE_hexadecimalLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EcaruleParserMINUS {
		{
			p.SetState(303)
			p.Match(EcaruleParserMINUS)
		}

	}
	{
		p.SetState(306)
		p.Match(EcaruleParserHEX_LIT)
	}

	return localctx
}

// IOctalLiteralContext is an interface to support dynamic dispatch.
type IOctalLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOctalLiteralContext differentiates from other interfaces.
	IsOctalLiteralContext()
}

type OctalLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOctalLiteralContext() *OctalLiteralContext {
	var p = new(OctalLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_octalLiteral
	return p
}

func (*OctalLiteralContext) IsOctalLiteralContext() {}

func NewOctalLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OctalLiteralContext {
	var p = new(OctalLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_octalLiteral

	return p
}

func (s *OctalLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *OctalLiteralContext) OCT_LIT() antlr.TerminalNode {
	return s.GetToken(EcaruleParserOCT_LIT, 0)
}

func (s *OctalLiteralContext) MINUS() antlr.TerminalNode {
	return s.GetToken(EcaruleParserMINUS, 0)
}

func (s *OctalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OctalLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OctalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewOctalLiteralContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterOctalLiteral(c)
	}
}

func (s *OctalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewOctalLiteralContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitOctalLiteral(c)
	}
}

func (p *EcaruleParser) OctalLiteral() (localctx IOctalLiteralContext) {
	localctx = NewOctalLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, EcaruleParserRULE_octalLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EcaruleParserMINUS {
		{
			p.SetState(308)
			p.Match(EcaruleParserMINUS)
		}

	}
	{
		p.SetState(311)
		p.Match(EcaruleParserOCT_LIT)
	}

	return localctx
}

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) DQUOTA_STRING() antlr.TerminalNode {
	return s.GetToken(EcaruleParserDQUOTA_STRING, 0)
}

func (s *StringLiteralContext) SQUOTA_STRING() antlr.TerminalNode {
	return s.GetToken(EcaruleParserSQUOTA_STRING, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewStringLiteralContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterStringLiteral(c)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewStringLiteralContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitStringLiteral(c)
	}
}

func (p *EcaruleParser) StringLiteral() (localctx IStringLiteralContext) {
	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, EcaruleParserRULE_stringLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(313)
		_la = p.GetTokenStream().LA(1)

		if !(_la == EcaruleParserDQUOTA_STRING || _la == EcaruleParserSQUOTA_STRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(EcaruleParserTRUE, 0)
}

func (s *BooleanLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(EcaruleParserFALSE, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewBooleanLiteralContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.EnterBooleanLiteral(c)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		c := grulev3.NewBooleanLiteralContext(s.parser, s.BaseParserRuleContext, -1)
		c.BaseParserRuleContext = s.BaseParserRuleContext
		listenerT.ExitBooleanLiteral(c)
	}
}

func (p *EcaruleParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, EcaruleParserRULE_booleanLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(315)
		_la = p.GetTokenStream().LA(1)

		if !(_la == EcaruleParserTRUE || _la == EcaruleParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *EcaruleParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 15:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 21:
		var t *ExpressionAtomContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionAtomContext)
		}
		return p.ExpressionAtom_Sempred(t, predIndex)

	case 23:
		var t *VariableContext = nil
		if localctx != nil {
			t = localctx.(*VariableContext)
		}
		return p.Variable_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *EcaruleParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *EcaruleParser) ExpressionAtom_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *EcaruleParser) Variable_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
