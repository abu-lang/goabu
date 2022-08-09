// Code generated from EcaruleParser.g4 by ANTLR 4.9.2 and MODIFIED by ../Makefile.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 57, 312,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 87, 10, 2,
	3, 2, 3, 2, 3, 3, 6, 3, 92, 10, 3, 13, 3, 14, 3, 93, 3, 4, 3, 4, 5, 4,
	98, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6,
	5, 6, 110, 10, 6, 3, 7, 3, 7, 5, 7, 114, 10, 7, 3, 8, 7, 8, 117, 10, 8,
	12, 8, 14, 8, 120, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 5, 9, 127, 10,
	9, 3, 9, 5, 9, 130, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 15, 6, 15, 153, 10, 15, 13, 15, 14, 15, 154, 3, 16,
	3, 16, 5, 16, 159, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 5,
	18, 167, 10, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 174, 10, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18,
	196, 10, 18, 12, 18, 14, 18, 199, 11, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 5, 24, 217, 10, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 7,
	24, 225, 10, 24, 12, 24, 14, 24, 228, 11, 24, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 5, 25, 235, 10, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 7, 26, 244, 10, 26, 12, 26, 14, 26, 247, 11, 26, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 5, 29, 259, 10, 29, 3,
	29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 7, 31, 269, 10, 31,
	12, 31, 14, 31, 272, 11, 31, 3, 32, 3, 32, 5, 32, 276, 10, 32, 3, 33, 5,
	33, 279, 10, 33, 3, 33, 3, 33, 3, 34, 5, 34, 284, 10, 34, 3, 34, 3, 34,
	3, 35, 3, 35, 3, 35, 5, 35, 291, 10, 35, 3, 36, 5, 36, 294, 10, 36, 3,
	36, 3, 36, 3, 37, 5, 37, 299, 10, 37, 3, 37, 3, 37, 3, 38, 5, 38, 304,
	10, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 2, 5, 34, 46,
	50, 41, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
	72, 74, 76, 78, 2, 8, 3, 2, 41, 42, 3, 2, 28, 32, 3, 2, 6, 8, 4, 2, 4,
	5, 38, 39, 4, 2, 27, 27, 33, 37, 3, 2, 22, 23, 2, 311, 2, 80, 3, 2, 2,
	2, 4, 91, 3, 2, 2, 2, 6, 95, 3, 2, 2, 2, 8, 103, 3, 2, 2, 2, 10, 109, 3,
	2, 2, 2, 12, 113, 3, 2, 2, 2, 14, 118, 3, 2, 2, 2, 16, 123, 3, 2, 2, 2,
	18, 136, 3, 2, 2, 2, 20, 139, 3, 2, 2, 2, 22, 141, 3, 2, 2, 2, 24, 143,
	3, 2, 2, 2, 26, 146, 3, 2, 2, 2, 28, 152, 3, 2, 2, 2, 30, 158, 3, 2, 2,
	2, 32, 160, 3, 2, 2, 2, 34, 173, 3, 2, 2, 2, 36, 200, 3, 2, 2, 2, 38, 202,
	3, 2, 2, 2, 40, 204, 3, 2, 2, 2, 42, 206, 3, 2, 2, 2, 44, 208, 3, 2, 2,
	2, 46, 216, 3, 2, 2, 2, 48, 234, 3, 2, 2, 2, 50, 236, 3, 2, 2, 2, 52, 248,
	3, 2, 2, 2, 54, 252, 3, 2, 2, 2, 56, 255, 3, 2, 2, 2, 58, 262, 3, 2, 2,
	2, 60, 265, 3, 2, 2, 2, 62, 275, 3, 2, 2, 2, 64, 278, 3, 2, 2, 2, 66, 283,
	3, 2, 2, 2, 68, 290, 3, 2, 2, 2, 70, 293, 3, 2, 2, 2, 72, 298, 3, 2, 2,
	2, 74, 303, 3, 2, 2, 2, 76, 307, 3, 2, 2, 2, 78, 309, 3, 2, 2, 2, 80, 81,
	7, 17, 2, 2, 81, 82, 7, 40, 2, 2, 82, 83, 7, 53, 2, 2, 83, 86, 5, 4, 3,
	2, 84, 85, 7, 54, 2, 2, 85, 87, 5, 8, 5, 2, 86, 84, 3, 2, 2, 2, 86, 87,
	3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 89, 5, 6, 4, 2, 89, 3, 3, 2, 2, 2,
	90, 92, 7, 40, 2, 2, 91, 90, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 91, 3,
	2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 5, 3, 2, 2, 2, 95, 97, 7, 55, 2, 2, 96,
	98, 7, 56, 2, 2, 97, 96, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 99, 3, 2,
	2, 2, 99, 100, 5, 34, 18, 2, 100, 101, 7, 57, 2, 2, 101, 102, 5, 8, 5,
	2, 102, 7, 3, 2, 2, 2, 103, 104, 5, 32, 17, 2, 104, 105, 5, 10, 6, 2, 105,
	9, 3, 2, 2, 2, 106, 107, 7, 10, 2, 2, 107, 110, 5, 12, 7, 2, 108, 110,
	3, 2, 2, 2, 109, 106, 3, 2, 2, 2, 109, 108, 3, 2, 2, 2, 110, 11, 3, 2,
	2, 2, 111, 114, 5, 8, 5, 2, 112, 114, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2,
	113, 112, 3, 2, 2, 2, 114, 13, 3, 2, 2, 2, 115, 117, 5, 16, 9, 2, 116,
	115, 3, 2, 2, 2, 117, 120, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119,
	3, 2, 2, 2, 119, 121, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 121, 122, 7, 2,
	2, 3, 122, 15, 3, 2, 2, 2, 123, 124, 7, 17, 2, 2, 124, 126, 5, 20, 11,
	2, 125, 127, 5, 22, 12, 2, 126, 125, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2,
	127, 129, 3, 2, 2, 2, 128, 130, 5, 18, 10, 2, 129, 128, 3, 2, 2, 2, 129,
	130, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 132, 7, 11, 2, 2, 132, 133,
	5, 24, 13, 2, 133, 134, 5, 26, 14, 2, 134, 135, 7, 12, 2, 2, 135, 17, 3,
	2, 2, 2, 136, 137, 7, 26, 2, 2, 137, 138, 5, 68, 35, 2, 138, 19, 3, 2,
	2, 2, 139, 140, 7, 40, 2, 2, 140, 21, 3, 2, 2, 2, 141, 142, 9, 2, 2, 2,
	142, 23, 3, 2, 2, 2, 143, 144, 7, 18, 2, 2, 144, 145, 5, 34, 18, 2, 145,
	25, 3, 2, 2, 2, 146, 147, 7, 19, 2, 2, 147, 148, 5, 28, 15, 2, 148, 27,
	3, 2, 2, 2, 149, 150, 5, 30, 16, 2, 150, 151, 7, 10, 2, 2, 151, 153, 3,
	2, 2, 2, 152, 149, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 152, 3, 2, 2,
	2, 154, 155, 3, 2, 2, 2, 155, 29, 3, 2, 2, 2, 156, 159, 5, 32, 17, 2, 157,
	159, 5, 46, 24, 2, 158, 156, 3, 2, 2, 2, 158, 157, 3, 2, 2, 2, 159, 31,
	3, 2, 2, 2, 160, 161, 5, 50, 26, 2, 161, 162, 9, 3, 2, 2, 162, 163, 5,
	34, 18, 2, 163, 33, 3, 2, 2, 2, 164, 166, 8, 18, 1, 2, 165, 167, 7, 25,
	2, 2, 166, 165, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2,
	168, 169, 7, 13, 2, 2, 169, 170, 5, 34, 18, 2, 170, 171, 7, 14, 2, 2, 171,
	174, 3, 2, 2, 2, 172, 174, 5, 46, 24, 2, 173, 164, 3, 2, 2, 2, 173, 172,
	3, 2, 2, 2, 174, 197, 3, 2, 2, 2, 175, 176, 12, 9, 2, 2, 176, 177, 5, 36,
	19, 2, 177, 178, 5, 34, 18, 10, 178, 196, 3, 2, 2, 2, 179, 180, 12, 8,
	2, 2, 180, 181, 5, 38, 20, 2, 181, 182, 5, 34, 18, 9, 182, 196, 3, 2, 2,
	2, 183, 184, 12, 7, 2, 2, 184, 185, 5, 40, 21, 2, 185, 186, 5, 34, 18,
	8, 186, 196, 3, 2, 2, 2, 187, 188, 12, 6, 2, 2, 188, 189, 5, 42, 22, 2,
	189, 190, 5, 34, 18, 7, 190, 196, 3, 2, 2, 2, 191, 192, 12, 5, 2, 2, 192,
	193, 5, 44, 23, 2, 193, 194, 5, 34, 18, 6, 194, 196, 3, 2, 2, 2, 195, 175,
	3, 2, 2, 2, 195, 179, 3, 2, 2, 2, 195, 183, 3, 2, 2, 2, 195, 187, 3, 2,
	2, 2, 195, 191, 3, 2, 2, 2, 196, 199, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2,
	197, 198, 3, 2, 2, 2, 198, 35, 3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 200, 201,
	9, 4, 2, 2, 201, 37, 3, 2, 2, 2, 202, 203, 9, 5, 2, 2, 203, 39, 3, 2, 2,
	2, 204, 205, 9, 6, 2, 2, 205, 41, 3, 2, 2, 2, 206, 207, 7, 20, 2, 2, 207,
	43, 3, 2, 2, 2, 208, 209, 7, 21, 2, 2, 209, 45, 3, 2, 2, 2, 210, 211, 8,
	24, 1, 2, 211, 217, 5, 48, 25, 2, 212, 217, 5, 50, 26, 2, 213, 217, 5,
	56, 29, 2, 214, 215, 7, 25, 2, 2, 215, 217, 5, 46, 24, 3, 216, 210, 3,
	2, 2, 2, 216, 212, 3, 2, 2, 2, 216, 213, 3, 2, 2, 2, 216, 214, 3, 2, 2,
	2, 217, 226, 3, 2, 2, 2, 218, 219, 12, 6, 2, 2, 219, 225, 5, 58, 30, 2,
	220, 221, 12, 5, 2, 2, 221, 225, 5, 54, 28, 2, 222, 223, 12, 4, 2, 2, 223,
	225, 5, 52, 27, 2, 224, 218, 3, 2, 2, 2, 224, 220, 3, 2, 2, 2, 224, 222,
	3, 2, 2, 2, 225, 228, 3, 2, 2, 2, 226, 224, 3, 2, 2, 2, 226, 227, 3, 2,
	2, 2, 227, 47, 3, 2, 2, 2, 228, 226, 3, 2, 2, 2, 229, 235, 5, 76, 39, 2,
	230, 235, 5, 68, 35, 2, 231, 235, 5, 62, 32, 2, 232, 235, 5, 78, 40, 2,
	233, 235, 7, 24, 2, 2, 234, 229, 3, 2, 2, 2, 234, 230, 3, 2, 2, 2, 234,
	231, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2, 234, 233, 3, 2, 2, 2, 235, 49, 3,
	2, 2, 2, 236, 237, 8, 26, 1, 2, 237, 238, 7, 40, 2, 2, 238, 245, 3, 2,
	2, 2, 239, 240, 12, 5, 2, 2, 240, 244, 5, 54, 28, 2, 241, 242, 12, 4, 2,
	2, 242, 244, 5, 52, 27, 2, 243, 239, 3, 2, 2, 2, 243, 241, 3, 2, 2, 2,
	244, 247, 3, 2, 2, 2, 245, 243, 3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246,
	51, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 248, 249, 7, 15, 2, 2, 249, 250,
	5, 34, 18, 2, 250, 251, 7, 16, 2, 2, 251, 53, 3, 2, 2, 2, 252, 253, 7,
	9, 2, 2, 253, 254, 7, 40, 2, 2, 254, 55, 3, 2, 2, 2, 255, 256, 7, 40, 2,
	2, 256, 258, 7, 13, 2, 2, 257, 259, 5, 60, 31, 2, 258, 257, 3, 2, 2, 2,
	258, 259, 3, 2, 2, 2, 259, 260, 3, 2, 2, 2, 260, 261, 7, 14, 2, 2, 261,
	57, 3, 2, 2, 2, 262, 263, 7, 9, 2, 2, 263, 264, 5, 56, 29, 2, 264, 59,
	3, 2, 2, 2, 265, 270, 5, 34, 18, 2, 266, 267, 7, 3, 2, 2, 267, 269, 5,
	34, 18, 2, 268, 266, 3, 2, 2, 2, 269, 272, 3, 2, 2, 2, 270, 268, 3, 2,
	2, 2, 270, 271, 3, 2, 2, 2, 271, 61, 3, 2, 2, 2, 272, 270, 3, 2, 2, 2,
	273, 276, 5, 64, 33, 2, 274, 276, 5, 66, 34, 2, 275, 273, 3, 2, 2, 2, 275,
	274, 3, 2, 2, 2, 276, 63, 3, 2, 2, 2, 277, 279, 7, 5, 2, 2, 278, 277, 3,
	2, 2, 2, 278, 279, 3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 280, 281, 7, 43, 2,
	2, 281, 65, 3, 2, 2, 2, 282, 284, 7, 5, 2, 2, 283, 282, 3, 2, 2, 2, 283,
	284, 3, 2, 2, 2, 284, 285, 3, 2, 2, 2, 285, 286, 7, 45, 2, 2, 286, 67,
	3, 2, 2, 2, 287, 291, 5, 70, 36, 2, 288, 291, 5, 72, 37, 2, 289, 291, 5,
	74, 38, 2, 290, 287, 3, 2, 2, 2, 290, 288, 3, 2, 2, 2, 290, 289, 3, 2,
	2, 2, 291, 69, 3, 2, 2, 2, 292, 294, 7, 5, 2, 2, 293, 292, 3, 2, 2, 2,
	293, 294, 3, 2, 2, 2, 294, 295, 3, 2, 2, 2, 295, 296, 7, 47, 2, 2, 296,
	71, 3, 2, 2, 2, 297, 299, 7, 5, 2, 2, 298, 297, 3, 2, 2, 2, 298, 299, 3,
	2, 2, 2, 299, 300, 3, 2, 2, 2, 300, 301, 7, 48, 2, 2, 301, 73, 3, 2, 2,
	2, 302, 304, 7, 5, 2, 2, 303, 302, 3, 2, 2, 2, 303, 304, 3, 2, 2, 2, 304,
	305, 3, 2, 2, 2, 305, 306, 7, 49, 2, 2, 306, 75, 3, 2, 2, 2, 307, 308,
	9, 2, 2, 2, 308, 77, 3, 2, 2, 2, 309, 310, 9, 7, 2, 2, 310, 79, 3, 2, 2,
	2, 31, 86, 93, 97, 109, 113, 118, 126, 129, 154, 158, 166, 173, 195, 197,
	216, 224, 226, 234, 243, 245, 258, 270, 275, 278, 283, 290, 293, 298, 303,
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
	"COMMENT", "LINE_COMMENT", "ON", "DEFAULT", "FOR", "ALL", "DO",
}

var ruleNames = []string{
	"prule", "events", "task", "actions", "tailActions", "maybeActions", "grl",
	"ruleEntry", "salience", "ruleName", "ruleDescription", "whenScope", "thenScope",
	"thenExpressionList", "thenExpression", "assignment", "expression", "mulDivOperators",
	"addMinusOperators", "comparisonOperator", "andLogicOperator", "orLogicOperator",
	"expressionAtom", "constant", "variable", "arrayMapSelector", "memberVariable",
	"functionCall", "methodCall", "argumentList", "floatLiteral", "decimalFloatLiteral",
	"hexadecimalFloatLiteral", "integerLiteral", "decimalLiteral", "hexadecimalLiteral",
	"octalLiteral", "stringLiteral", "booleanLiteral",
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
	EcaruleParserON                = 51
	EcaruleParserDEFAULT           = 52
	EcaruleParserFOR               = 53
	EcaruleParserALL               = 54
	EcaruleParserDO                = 55
)

// EcaruleParser rules.
const (
	EcaruleParserRULE_prule                   = 0
	EcaruleParserRULE_events                  = 1
	EcaruleParserRULE_task                    = 2
	EcaruleParserRULE_actions                 = 3
	EcaruleParserRULE_tailActions             = 4
	EcaruleParserRULE_maybeActions            = 5
	EcaruleParserRULE_grl                     = 6
	EcaruleParserRULE_ruleEntry               = 7
	EcaruleParserRULE_salience                = 8
	EcaruleParserRULE_ruleName                = 9
	EcaruleParserRULE_ruleDescription         = 10
	EcaruleParserRULE_whenScope               = 11
	EcaruleParserRULE_thenScope               = 12
	EcaruleParserRULE_thenExpressionList      = 13
	EcaruleParserRULE_thenExpression          = 14
	EcaruleParserRULE_assignment              = 15
	EcaruleParserRULE_expression              = 16
	EcaruleParserRULE_mulDivOperators         = 17
	EcaruleParserRULE_addMinusOperators       = 18
	EcaruleParserRULE_comparisonOperator      = 19
	EcaruleParserRULE_andLogicOperator        = 20
	EcaruleParserRULE_orLogicOperator         = 21
	EcaruleParserRULE_expressionAtom          = 22
	EcaruleParserRULE_constant                = 23
	EcaruleParserRULE_variable                = 24
	EcaruleParserRULE_arrayMapSelector        = 25
	EcaruleParserRULE_memberVariable          = 26
	EcaruleParserRULE_functionCall            = 27
	EcaruleParserRULE_methodCall              = 28
	EcaruleParserRULE_argumentList            = 29
	EcaruleParserRULE_floatLiteral            = 30
	EcaruleParserRULE_decimalFloatLiteral     = 31
	EcaruleParserRULE_hexadecimalFloatLiteral = 32
	EcaruleParserRULE_integerLiteral          = 33
	EcaruleParserRULE_decimalLiteral          = 34
	EcaruleParserRULE_hexadecimalLiteral      = 35
	EcaruleParserRULE_octalLiteral            = 36
	EcaruleParserRULE_stringLiteral           = 37
	EcaruleParserRULE_booleanLiteral          = 38
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

func (s *PruleContext) SIMPLENAME() antlr.TerminalNode {
	return s.GetToken(EcaruleParserSIMPLENAME, 0)
}

func (s *PruleContext) ON() antlr.TerminalNode {
	return s.GetToken(EcaruleParserON, 0)
}

func (s *PruleContext) Events() IEventsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEventsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEventsContext)
}

func (s *PruleContext) Task() ITaskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITaskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITaskContext)
}

func (s *PruleContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(EcaruleParserDEFAULT, 0)
}

func (s *PruleContext) Actions() IActionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IActionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IActionsContext)
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
		p.SetState(78)
		p.Match(EcaruleParserRULE)
	}
	{
		p.SetState(79)
		p.Match(EcaruleParserSIMPLENAME)
	}
	{
		p.SetState(80)
		p.Match(EcaruleParserON)
	}
	{
		p.SetState(81)
		p.Events()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EcaruleParserDEFAULT {
		{
			p.SetState(82)
			p.Match(EcaruleParserDEFAULT)
		}
		{
			p.SetState(83)
			p.Actions()
		}

	}
	{
		p.SetState(86)
		p.Task()
	}

	return localctx
}

// IEventsContext is an interface to support dynamic dispatch.
type IEventsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEventsContext differentiates from other interfaces.
	IsEventsContext()
}

type EventsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEventsContext() *EventsContext {
	var p = new(EventsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_events
	return p
}

func (*EventsContext) IsEventsContext() {}

func NewEventsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EventsContext {
	var p = new(EventsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_events

	return p
}

func (s *EventsContext) GetParser() antlr.Parser { return s.parser }

func (s *EventsContext) AllSIMPLENAME() []antlr.TerminalNode {
	return s.GetTokens(EcaruleParserSIMPLENAME)
}

func (s *EventsContext) SIMPLENAME(i int) antlr.TerminalNode {
	return s.GetToken(EcaruleParserSIMPLENAME, i)
}

func (s *EventsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EventsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EventsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		listenerT.EnterEvents(s)
	}
}

func (s *EventsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		listenerT.ExitEvents(s)
	}
}

func (p *EcaruleParser) Events() (localctx IEventsContext) {
	localctx = NewEventsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EcaruleParserRULE_events)
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
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EcaruleParserSIMPLENAME {
		{
			p.SetState(88)
			p.Match(EcaruleParserSIMPLENAME)
		}

		p.SetState(91)
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

func (s *TaskContext) Actions() IActionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IActionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IActionsContext)
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
		p.SetState(93)
		p.Match(EcaruleParserFOR)
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EcaruleParserALL {
		{
			p.SetState(94)
			p.Match(EcaruleParserALL)
		}

	}
	{
		p.SetState(97)
		p.expression(0)
	}
	{
		p.SetState(98)
		p.Match(EcaruleParserDO)
	}
	{
		p.SetState(99)
		p.Actions()
	}

	return localctx
}

// IActionsContext is an interface to support dynamic dispatch.
type IActionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActionsContext differentiates from other interfaces.
	IsActionsContext()
}

type ActionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionsContext() *ActionsContext {
	var p = new(ActionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_actions
	return p
}

func (*ActionsContext) IsActionsContext() {}

func NewActionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionsContext {
	var p = new(ActionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_actions

	return p
}

func (s *ActionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionsContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ActionsContext) TailActions() ITailActionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITailActionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITailActionsContext)
}

func (s *ActionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		listenerT.EnterActions(s)
	}
}

func (s *ActionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		listenerT.ExitActions(s)
	}
}

func (p *EcaruleParser) Actions() (localctx IActionsContext) {
	localctx = NewActionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EcaruleParserRULE_actions)

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
		p.SetState(101)
		p.Assignment()
	}
	{
		p.SetState(102)
		p.TailActions()
	}

	return localctx
}

// ITailActionsContext is an interface to support dynamic dispatch.
type ITailActionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTailActionsContext differentiates from other interfaces.
	IsTailActionsContext()
}

type TailActionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTailActionsContext() *TailActionsContext {
	var p = new(TailActionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_tailActions
	return p
}

func (*TailActionsContext) IsTailActionsContext() {}

func NewTailActionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TailActionsContext {
	var p = new(TailActionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_tailActions

	return p
}

func (s *TailActionsContext) GetParser() antlr.Parser { return s.parser }

func (s *TailActionsContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(EcaruleParserSEMICOLON, 0)
}

func (s *TailActionsContext) MaybeActions() IMaybeActionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMaybeActionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMaybeActionsContext)
}

func (s *TailActionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TailActionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TailActionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		listenerT.EnterTailActions(s)
	}
}

func (s *TailActionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		listenerT.ExitTailActions(s)
	}
}

func (p *EcaruleParser) TailActions() (localctx ITailActionsContext) {
	localctx = NewTailActionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, EcaruleParserRULE_tailActions)

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

	p.SetState(107)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case EcaruleParserSEMICOLON:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)
			p.Match(EcaruleParserSEMICOLON)
		}
		{
			p.SetState(105)
			p.MaybeActions()
		}

	case EcaruleParserEOF, EcaruleParserFOR:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMaybeActionsContext is an interface to support dynamic dispatch.
type IMaybeActionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMaybeActionsContext differentiates from other interfaces.
	IsMaybeActionsContext()
}

type MaybeActionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMaybeActionsContext() *MaybeActionsContext {
	var p = new(MaybeActionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EcaruleParserRULE_maybeActions
	return p
}

func (*MaybeActionsContext) IsMaybeActionsContext() {}

func NewMaybeActionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MaybeActionsContext {
	var p = new(MaybeActionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EcaruleParserRULE_maybeActions

	return p
}

func (s *MaybeActionsContext) GetParser() antlr.Parser { return s.parser }

func (s *MaybeActionsContext) Actions() IActionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IActionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IActionsContext)
}

func (s *MaybeActionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MaybeActionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MaybeActionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		listenerT.EnterMaybeActions(s)
	}
}

func (s *MaybeActionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EcaruleParserListener); ok {
		listenerT.ExitMaybeActions(s)
	}
}

func (p *EcaruleParser) MaybeActions() (localctx IMaybeActionsContext) {
	localctx = NewMaybeActionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, EcaruleParserRULE_maybeActions)

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

	p.SetState(111)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case EcaruleParserSIMPLENAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.Actions()
		}

	case EcaruleParserEOF, EcaruleParserFOR:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 12, EcaruleParserRULE_grl)
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
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EcaruleParserRULE {
		{
			p.SetState(113)
			p.RuleEntry()
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(119)
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
	p.EnterRule(localctx, 14, EcaruleParserRULE_ruleEntry)
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
		p.SetState(121)
		p.Match(EcaruleParserRULE)
	}
	{
		p.SetState(122)
		p.RuleName()
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EcaruleParserDQUOTA_STRING || _la == EcaruleParserSQUOTA_STRING {
		{
			p.SetState(123)
			p.RuleDescription()
		}

	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EcaruleParserSALIENCE {
		{
			p.SetState(126)
			p.Salience()
		}

	}
	{
		p.SetState(129)
		p.Match(EcaruleParserLR_BRACE)
	}
	{
		p.SetState(130)
		p.WhenScope()
	}
	{
		p.SetState(131)
		p.ThenScope()
	}
	{
		p.SetState(132)
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
	p.EnterRule(localctx, 16, EcaruleParserRULE_salience)

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
		p.SetState(134)
		p.Match(EcaruleParserSALIENCE)
	}
	{
		p.SetState(135)
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
	p.EnterRule(localctx, 18, EcaruleParserRULE_ruleName)

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
		p.SetState(137)
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
	p.EnterRule(localctx, 20, EcaruleParserRULE_ruleDescription)
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
		p.SetState(139)
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
	p.EnterRule(localctx, 22, EcaruleParserRULE_whenScope)

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
		p.SetState(141)
		p.Match(EcaruleParserWHEN)
	}
	{
		p.SetState(142)
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
	p.EnterRule(localctx, 24, EcaruleParserRULE_thenScope)

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
		p.SetState(144)
		p.Match(EcaruleParserTHEN)
	}
	{
		p.SetState(145)
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
	p.EnterRule(localctx, 26, EcaruleParserRULE_thenExpressionList)
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
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<EcaruleParserMINUS)|(1<<EcaruleParserTRUE)|(1<<EcaruleParserFALSE)|(1<<EcaruleParserNIL_LITERAL)|(1<<EcaruleParserNEGATION))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(EcaruleParserSIMPLENAME-38))|(1<<(EcaruleParserDQUOTA_STRING-38))|(1<<(EcaruleParserSQUOTA_STRING-38))|(1<<(EcaruleParserDECIMAL_FLOAT_LIT-38))|(1<<(EcaruleParserHEX_FLOAT_LIT-38))|(1<<(EcaruleParserDEC_LIT-38))|(1<<(EcaruleParserHEX_LIT-38))|(1<<(EcaruleParserOCT_LIT-38)))) != 0) {
		{
			p.SetState(147)
			p.ThenExpression()
		}
		{
			p.SetState(148)
			p.Match(EcaruleParserSEMICOLON)
		}

		p.SetState(152)
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
	p.EnterRule(localctx, 28, EcaruleParserRULE_thenExpression)

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

	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(154)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(155)
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
	p.EnterRule(localctx, 30, EcaruleParserRULE_assignment)
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
		p.SetState(158)
		p.variable(0)
	}
	{
		p.SetState(159)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<EcaruleParserASSIGN)|(1<<EcaruleParserPLUS_ASIGN)|(1<<EcaruleParserMINUS_ASIGN)|(1<<EcaruleParserDIV_ASIGN)|(1<<EcaruleParserMUL_ASIGN))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(160)
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
	_startState := 32
	p.EnterRecursionRule(localctx, 32, EcaruleParserRULE_expression, _p)
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
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EcaruleParserNEGATION {
			{
				p.SetState(163)
				p.Match(EcaruleParserNEGATION)
			}

		}
		{
			p.SetState(166)
			p.Match(EcaruleParserLR_BRACKET)
		}
		{
			p.SetState(167)
			p.expression(0)
		}
		{
			p.SetState(168)
			p.Match(EcaruleParserRR_BRACKET)
		}

	case 2:
		{
			p.SetState(170)
			p.expressionAtom(0)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(193)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, EcaruleParserRULE_expression)
				p.SetState(173)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(174)
					p.MulDivOperators()
				}
				{
					p.SetState(175)
					p.expression(8)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, EcaruleParserRULE_expression)
				p.SetState(177)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(178)
					p.AddMinusOperators()
				}
				{
					p.SetState(179)
					p.expression(7)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, EcaruleParserRULE_expression)
				p.SetState(181)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(182)
					p.ComparisonOperator()
				}
				{
					p.SetState(183)
					p.expression(6)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, EcaruleParserRULE_expression)
				p.SetState(185)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(186)
					p.AndLogicOperator()
				}
				{
					p.SetState(187)
					p.expression(5)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, EcaruleParserRULE_expression)
				p.SetState(189)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(190)
					p.OrLogicOperator()
				}
				{
					p.SetState(191)
					p.expression(4)
				}

			}

		}
		p.SetState(197)
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
	p.EnterRule(localctx, 34, EcaruleParserRULE_mulDivOperators)
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
		p.SetState(198)
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
	p.EnterRule(localctx, 36, EcaruleParserRULE_addMinusOperators)
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
		p.SetState(200)
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
	p.EnterRule(localctx, 38, EcaruleParserRULE_comparisonOperator)
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
		p.SetState(202)
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
	p.EnterRule(localctx, 40, EcaruleParserRULE_andLogicOperator)

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
		p.SetState(204)
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
	p.EnterRule(localctx, 42, EcaruleParserRULE_orLogicOperator)

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
	_startState := 44
	p.EnterRecursionRule(localctx, 44, EcaruleParserRULE_expressionAtom, _p)

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
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(209)
			p.Constant()
		}

	case 2:
		{
			p.SetState(210)
			p.variable(0)
		}

	case 3:
		{
			p.SetState(211)
			p.FunctionCall()
		}

	case 4:
		{
			p.SetState(212)
			p.Match(EcaruleParserNEGATION)
		}
		{
			p.SetState(213)
			p.expressionAtom(1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(222)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionAtomContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, EcaruleParserRULE_expressionAtom)
				p.SetState(216)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(217)
					p.MethodCall()
				}

			case 2:
				localctx = NewExpressionAtomContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, EcaruleParserRULE_expressionAtom)
				p.SetState(218)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(219)
					p.MemberVariable()
				}

			case 3:
				localctx = NewExpressionAtomContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, EcaruleParserRULE_expressionAtom)
				p.SetState(220)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(221)
					p.ArrayMapSelector()
				}

			}

		}
		p.SetState(226)
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
	p.EnterRule(localctx, 46, EcaruleParserRULE_constant)

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

	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(227)
			p.StringLiteral()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(228)
			p.IntegerLiteral()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(229)
			p.FloatLiteral()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(230)
			p.BooleanLiteral()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(231)
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
	_startState := 48
	p.EnterRecursionRule(localctx, 48, EcaruleParserRULE_variable, _p)

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
		p.SetState(235)
		p.Match(EcaruleParserSIMPLENAME)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(241)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
			case 1:
				localctx = NewVariableContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, EcaruleParserRULE_variable)
				p.SetState(237)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(238)
					p.MemberVariable()
				}

			case 2:
				localctx = NewVariableContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, EcaruleParserRULE_variable)
				p.SetState(239)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(240)
					p.ArrayMapSelector()
				}

			}

		}
		p.SetState(245)
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
	p.EnterRule(localctx, 50, EcaruleParserRULE_arrayMapSelector)

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
		p.SetState(246)
		p.Match(EcaruleParserLS_BRACKET)
	}
	{
		p.SetState(247)
		p.expression(0)
	}
	{
		p.SetState(248)
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
	p.EnterRule(localctx, 52, EcaruleParserRULE_memberVariable)

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
		p.SetState(250)
		p.Match(EcaruleParserDOT)
	}
	{
		p.SetState(251)
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
	p.EnterRule(localctx, 54, EcaruleParserRULE_functionCall)
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
		p.SetState(253)
		p.Match(EcaruleParserSIMPLENAME)
	}
	{
		p.SetState(254)
		p.Match(EcaruleParserLR_BRACKET)
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<EcaruleParserMINUS)|(1<<EcaruleParserLR_BRACKET)|(1<<EcaruleParserTRUE)|(1<<EcaruleParserFALSE)|(1<<EcaruleParserNIL_LITERAL)|(1<<EcaruleParserNEGATION))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(EcaruleParserSIMPLENAME-38))|(1<<(EcaruleParserDQUOTA_STRING-38))|(1<<(EcaruleParserSQUOTA_STRING-38))|(1<<(EcaruleParserDECIMAL_FLOAT_LIT-38))|(1<<(EcaruleParserHEX_FLOAT_LIT-38))|(1<<(EcaruleParserDEC_LIT-38))|(1<<(EcaruleParserHEX_LIT-38))|(1<<(EcaruleParserOCT_LIT-38)))) != 0) {
		{
			p.SetState(255)
			p.ArgumentList()
		}

	}
	{
		p.SetState(258)
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
	p.EnterRule(localctx, 56, EcaruleParserRULE_methodCall)

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
		p.SetState(260)
		p.Match(EcaruleParserDOT)
	}
	{
		p.SetState(261)
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
	p.EnterRule(localctx, 58, EcaruleParserRULE_argumentList)
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
		p.SetState(263)
		p.expression(0)
	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EcaruleParserT__0 {
		{
			p.SetState(264)
			p.Match(EcaruleParserT__0)
		}
		{
			p.SetState(265)
			p.expression(0)
		}

		p.SetState(270)
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
	p.EnterRule(localctx, 60, EcaruleParserRULE_floatLiteral)

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

	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(271)
			p.DecimalFloatLiteral()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(272)
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
	p.EnterRule(localctx, 62, EcaruleParserRULE_decimalFloatLiteral)
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
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EcaruleParserMINUS {
		{
			p.SetState(275)
			p.Match(EcaruleParserMINUS)
		}

	}
	{
		p.SetState(278)
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
	p.EnterRule(localctx, 64, EcaruleParserRULE_hexadecimalFloatLiteral)
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
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EcaruleParserMINUS {
		{
			p.SetState(280)
			p.Match(EcaruleParserMINUS)
		}

	}
	{
		p.SetState(283)
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
	p.EnterRule(localctx, 66, EcaruleParserRULE_integerLiteral)

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

	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(285)
			p.DecimalLiteral()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(286)
			p.HexadecimalLiteral()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(287)
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
	p.EnterRule(localctx, 68, EcaruleParserRULE_decimalLiteral)
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
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EcaruleParserMINUS {
		{
			p.SetState(290)
			p.Match(EcaruleParserMINUS)
		}

	}
	{
		p.SetState(293)
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
	p.EnterRule(localctx, 70, EcaruleParserRULE_hexadecimalLiteral)
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
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EcaruleParserMINUS {
		{
			p.SetState(295)
			p.Match(EcaruleParserMINUS)
		}

	}
	{
		p.SetState(298)
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
	p.EnterRule(localctx, 72, EcaruleParserRULE_octalLiteral)
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
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EcaruleParserMINUS {
		{
			p.SetState(300)
			p.Match(EcaruleParserMINUS)
		}

	}
	{
		p.SetState(303)
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
	p.EnterRule(localctx, 74, EcaruleParserRULE_stringLiteral)
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
		p.SetState(305)
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
	p.EnterRule(localctx, 76, EcaruleParserRULE_booleanLiteral)
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
		p.SetState(307)
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
	case 16:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 22:
		var t *ExpressionAtomContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionAtomContext)
		}
		return p.ExpressionAtom_Sempred(t, predIndex)

	case 24:
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
