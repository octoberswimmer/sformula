// Code generated from ./FormulaParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // FormulaParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type FormulaParser struct {
	*antlr.BaseParser
}

var FormulaParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func formulaparserParserInit() {
	staticData := &FormulaParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "'('", "')'", "'.'", "','", "'^'", "'&'",
		"'+'", "'-'", "'*'", "'/'", "'<'", "'>'", "'='", "'<>'", "'!='", "'=='",
		"'&&'", "'||'",
	}
	staticData.SymbolicNames = []string{
		"", "ABS", "ADDMONTHS", "AND_FUNC", "BEGINS", "BLANKVALUE", "BR", "CASE",
		"CASESAFEID", "CEILING", "CONTAINS", "CURRENCYRATE", "DATE", "DATEVALUE",
		"DATETIMEVALUE", "DAY", "DISTANCE", "EXP", "FIND", "FLOOR", "GEOLOCATION",
		"GETRECORDIDS", "GETSESSIONID", "HOUR", "HTMLENCODE", "HYPERLINK", "IF",
		"IMAGE", "IMAGEPROXYURL", "INCLUDE", "INCLUDES", "ISBLANK", "ISCHANGED",
		"ISCLONE", "ISNEW", "ISNULL", "ISNUMBER", "ISPICKVAL", "JSENCODE", "JSINHTMLENCODE",
		"JUNCTIONIDLIST", "LEFT", "LEN", "LINKTO", "LN", "LOG", "LOWER", "LPAD",
		"MAX", "MCEILING", "MFLOOR", "MID", "MILLISECOND", "MIN", "MINUTE",
		"MOD", "MONTH", "NOT", "NOW", "NULLVALUE", "OR_FUNC", "PARENTGROUPVAL",
		"PREDICT", "PREVGROUPVAL", "PRIORVALUE", "REGEX", "REQUIRESCRIPT", "REVERSE",
		"RIGHT", "ROUND", "RPAD", "SECOND", "SQRT", "SUBSTITUTE", "TEXT", "TIMENOW",
		"TIMEVALUE", "TODAY", "TRIM", "UPPER", "URLENCODE", "URLFOR", "VALUE",
		"VLOOKUP", "WEEKDAY", "YEAR", "StringLiteral", "IntegerLiteral", "FloatingPointLiteral",
		"BooleanLiteral", "NullLiteral", "Identifier", "LPAREN", "RPAREN", "DOT",
		"COMMA", "CARET", "BITAND", "ADD", "SUB", "MUL", "DIV", "LT", "GT",
		"ASSIGN", "LESSANDGREATER", "NOTEQUAL", "EQUAL", "AND", "OR", "WS",
		"COMMENT",
	}
	staticData.RuleNames = []string{
		"compilationUnit", "expression", "functionCall", "abs", "addMonths",
		"and", "begins", "blankvalue", "br", "case", "casesafeid", "ceiling",
		"contains", "currencyrate", "date", "datevalue", "datetimevalue", "day",
		"distance", "exp", "find", "floor", "geolocation", "getrecordids", "getsessionid",
		"hour", "htmlencode", "hyperlink", "if", "image", "imageproxyurl", "mid",
		"includes", "text", "value", "fieldExpression", "valueExpression", "resultExpression",
		"defaultExpression", "yearExpression", "monthExpression", "dayExpression",
		"unitExpression", "searchExpression", "textExpression", "startNum",
		"latitudeExpression", "longitudeExpression", "urlExpression", "nameExpression",
		"targetExpression", "logicalExpression", "ifTrueExpression", "ifFalseExpression",
		"heightExpression", "widthExpression", "start", "numChars", "primary",
		"arguments", "literal",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 111, 494, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 133, 8, 1, 10, 1, 12, 1, 136, 9, 1,
		3, 1, 138, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 3, 1, 152, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 5, 1, 161, 8, 1, 10, 1, 12, 1, 164, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 198, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 5, 5, 219, 8, 5, 10, 5, 12, 5, 222, 9, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 256, 8, 9, 10, 9, 12, 9, 259, 9, 9,
		1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 331, 8, 20, 1, 20,
		1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 373, 8, 27,
		1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		3, 29, 396, 8, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37,
		1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1,
		43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48,
		1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 1,
		53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56, 1, 56, 1, 57, 1, 57, 1, 58, 1, 58,
		1, 58, 1, 58, 1, 58, 3, 58, 482, 8, 58, 1, 59, 1, 59, 1, 59, 5, 59, 487,
		8, 59, 10, 59, 12, 59, 490, 9, 59, 1, 60, 1, 60, 1, 60, 0, 1, 2, 61, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
		40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74,
		76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108,
		110, 112, 114, 116, 118, 120, 0, 5, 1, 0, 98, 101, 1, 0, 102, 103, 1, 0,
		105, 107, 1, 0, 108, 109, 1, 0, 86, 90, 481, 0, 122, 1, 0, 0, 0, 2, 137,
		1, 0, 0, 0, 4, 197, 1, 0, 0, 0, 6, 199, 1, 0, 0, 0, 8, 204, 1, 0, 0, 0,
		10, 211, 1, 0, 0, 0, 12, 225, 1, 0, 0, 0, 14, 232, 1, 0, 0, 0, 16, 239,
		1, 0, 0, 0, 18, 243, 1, 0, 0, 0, 20, 263, 1, 0, 0, 0, 22, 268, 1, 0, 0,
		0, 24, 273, 1, 0, 0, 0, 26, 280, 1, 0, 0, 0, 28, 285, 1, 0, 0, 0, 30, 294,
		1, 0, 0, 0, 32, 299, 1, 0, 0, 0, 34, 304, 1, 0, 0, 0, 36, 309, 1, 0, 0,
		0, 38, 318, 1, 0, 0, 0, 40, 323, 1, 0, 0, 0, 42, 334, 1, 0, 0, 0, 44, 339,
		1, 0, 0, 0, 46, 346, 1, 0, 0, 0, 48, 351, 1, 0, 0, 0, 50, 355, 1, 0, 0,
		0, 52, 360, 1, 0, 0, 0, 54, 365, 1, 0, 0, 0, 56, 376, 1, 0, 0, 0, 58, 385,
		1, 0, 0, 0, 60, 399, 1, 0, 0, 0, 62, 404, 1, 0, 0, 0, 64, 413, 1, 0, 0,
		0, 66, 420, 1, 0, 0, 0, 68, 425, 1, 0, 0, 0, 70, 430, 1, 0, 0, 0, 72, 432,
		1, 0, 0, 0, 74, 434, 1, 0, 0, 0, 76, 436, 1, 0, 0, 0, 78, 438, 1, 0, 0,
		0, 80, 440, 1, 0, 0, 0, 82, 442, 1, 0, 0, 0, 84, 444, 1, 0, 0, 0, 86, 446,
		1, 0, 0, 0, 88, 448, 1, 0, 0, 0, 90, 450, 1, 0, 0, 0, 92, 452, 1, 0, 0,
		0, 94, 454, 1, 0, 0, 0, 96, 456, 1, 0, 0, 0, 98, 458, 1, 0, 0, 0, 100,
		460, 1, 0, 0, 0, 102, 462, 1, 0, 0, 0, 104, 464, 1, 0, 0, 0, 106, 466,
		1, 0, 0, 0, 108, 468, 1, 0, 0, 0, 110, 470, 1, 0, 0, 0, 112, 472, 1, 0,
		0, 0, 114, 474, 1, 0, 0, 0, 116, 481, 1, 0, 0, 0, 118, 483, 1, 0, 0, 0,
		120, 491, 1, 0, 0, 0, 122, 123, 3, 2, 1, 0, 123, 1, 1, 0, 0, 0, 124, 125,
		6, 1, -1, 0, 125, 138, 3, 116, 58, 0, 126, 138, 3, 4, 2, 0, 127, 128, 5,
		99, 0, 0, 128, 138, 3, 2, 1, 8, 129, 134, 5, 91, 0, 0, 130, 131, 5, 94,
		0, 0, 131, 133, 5, 91, 0, 0, 132, 130, 1, 0, 0, 0, 133, 136, 1, 0, 0, 0,
		134, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 138, 1, 0, 0, 0, 136,
		134, 1, 0, 0, 0, 137, 124, 1, 0, 0, 0, 137, 126, 1, 0, 0, 0, 137, 127,
		1, 0, 0, 0, 137, 129, 1, 0, 0, 0, 138, 162, 1, 0, 0, 0, 139, 140, 10, 6,
		0, 0, 140, 141, 5, 96, 0, 0, 141, 161, 3, 2, 1, 7, 142, 143, 10, 5, 0,
		0, 143, 144, 5, 97, 0, 0, 144, 161, 3, 2, 1, 6, 145, 146, 10, 4, 0, 0,
		146, 147, 7, 0, 0, 0, 147, 161, 3, 2, 1, 5, 148, 149, 10, 3, 0, 0, 149,
		151, 7, 1, 0, 0, 150, 152, 5, 104, 0, 0, 151, 150, 1, 0, 0, 0, 151, 152,
		1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 161, 3, 2, 1, 4, 154, 155, 10, 2,
		0, 0, 155, 156, 7, 2, 0, 0, 156, 161, 3, 2, 1, 3, 157, 158, 10, 1, 0, 0,
		158, 159, 7, 3, 0, 0, 159, 161, 3, 2, 1, 2, 160, 139, 1, 0, 0, 0, 160,
		142, 1, 0, 0, 0, 160, 145, 1, 0, 0, 0, 160, 148, 1, 0, 0, 0, 160, 154,
		1, 0, 0, 0, 160, 157, 1, 0, 0, 0, 161, 164, 1, 0, 0, 0, 162, 160, 1, 0,
		0, 0, 162, 163, 1, 0, 0, 0, 163, 3, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 165,
		198, 3, 6, 3, 0, 166, 198, 3, 8, 4, 0, 167, 198, 3, 10, 5, 0, 168, 198,
		3, 12, 6, 0, 169, 198, 3, 14, 7, 0, 170, 198, 3, 16, 8, 0, 171, 198, 3,
		18, 9, 0, 172, 198, 3, 20, 10, 0, 173, 198, 3, 22, 11, 0, 174, 198, 3,
		24, 12, 0, 175, 198, 3, 26, 13, 0, 176, 198, 3, 28, 14, 0, 177, 198, 3,
		30, 15, 0, 178, 198, 3, 32, 16, 0, 179, 198, 3, 34, 17, 0, 180, 198, 3,
		36, 18, 0, 181, 198, 3, 38, 19, 0, 182, 198, 3, 40, 20, 0, 183, 198, 3,
		42, 21, 0, 184, 198, 3, 44, 22, 0, 185, 198, 3, 46, 23, 0, 186, 198, 3,
		48, 24, 0, 187, 198, 3, 50, 25, 0, 188, 198, 3, 52, 26, 0, 189, 198, 3,
		54, 27, 0, 190, 198, 3, 56, 28, 0, 191, 198, 3, 58, 29, 0, 192, 198, 3,
		60, 30, 0, 193, 198, 3, 64, 32, 0, 194, 198, 3, 62, 31, 0, 195, 198, 3,
		66, 33, 0, 196, 198, 3, 68, 34, 0, 197, 165, 1, 0, 0, 0, 197, 166, 1, 0,
		0, 0, 197, 167, 1, 0, 0, 0, 197, 168, 1, 0, 0, 0, 197, 169, 1, 0, 0, 0,
		197, 170, 1, 0, 0, 0, 197, 171, 1, 0, 0, 0, 197, 172, 1, 0, 0, 0, 197,
		173, 1, 0, 0, 0, 197, 174, 1, 0, 0, 0, 197, 175, 1, 0, 0, 0, 197, 176,
		1, 0, 0, 0, 197, 177, 1, 0, 0, 0, 197, 178, 1, 0, 0, 0, 197, 179, 1, 0,
		0, 0, 197, 180, 1, 0, 0, 0, 197, 181, 1, 0, 0, 0, 197, 182, 1, 0, 0, 0,
		197, 183, 1, 0, 0, 0, 197, 184, 1, 0, 0, 0, 197, 185, 1, 0, 0, 0, 197,
		186, 1, 0, 0, 0, 197, 187, 1, 0, 0, 0, 197, 188, 1, 0, 0, 0, 197, 189,
		1, 0, 0, 0, 197, 190, 1, 0, 0, 0, 197, 191, 1, 0, 0, 0, 197, 192, 1, 0,
		0, 0, 197, 193, 1, 0, 0, 0, 197, 194, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0,
		197, 196, 1, 0, 0, 0, 198, 5, 1, 0, 0, 0, 199, 200, 5, 1, 0, 0, 200, 201,
		5, 92, 0, 0, 201, 202, 3, 2, 1, 0, 202, 203, 5, 93, 0, 0, 203, 7, 1, 0,
		0, 0, 204, 205, 5, 2, 0, 0, 205, 206, 5, 92, 0, 0, 206, 207, 3, 2, 1, 0,
		207, 208, 5, 95, 0, 0, 208, 209, 3, 2, 1, 0, 209, 210, 5, 93, 0, 0, 210,
		9, 1, 0, 0, 0, 211, 212, 5, 3, 0, 0, 212, 213, 5, 92, 0, 0, 213, 214, 3,
		2, 1, 0, 214, 215, 5, 95, 0, 0, 215, 220, 3, 2, 1, 0, 216, 217, 5, 95,
		0, 0, 217, 219, 3, 2, 1, 0, 218, 216, 1, 0, 0, 0, 219, 222, 1, 0, 0, 0,
		220, 218, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 223, 1, 0, 0, 0, 222,
		220, 1, 0, 0, 0, 223, 224, 5, 93, 0, 0, 224, 11, 1, 0, 0, 0, 225, 226,
		5, 4, 0, 0, 226, 227, 5, 92, 0, 0, 227, 228, 3, 2, 1, 0, 228, 229, 5, 95,
		0, 0, 229, 230, 3, 2, 1, 0, 230, 231, 5, 93, 0, 0, 231, 13, 1, 0, 0, 0,
		232, 233, 5, 5, 0, 0, 233, 234, 5, 92, 0, 0, 234, 235, 3, 2, 1, 0, 235,
		236, 5, 95, 0, 0, 236, 237, 3, 2, 1, 0, 237, 238, 5, 93, 0, 0, 238, 15,
		1, 0, 0, 0, 239, 240, 5, 6, 0, 0, 240, 241, 5, 92, 0, 0, 241, 242, 5, 93,
		0, 0, 242, 17, 1, 0, 0, 0, 243, 244, 5, 7, 0, 0, 244, 245, 5, 92, 0, 0,
		245, 246, 3, 2, 1, 0, 246, 247, 5, 95, 0, 0, 247, 248, 3, 72, 36, 0, 248,
		249, 5, 95, 0, 0, 249, 257, 3, 74, 37, 0, 250, 251, 5, 95, 0, 0, 251, 252,
		3, 72, 36, 0, 252, 253, 5, 95, 0, 0, 253, 254, 3, 74, 37, 0, 254, 256,
		1, 0, 0, 0, 255, 250, 1, 0, 0, 0, 256, 259, 1, 0, 0, 0, 257, 255, 1, 0,
		0, 0, 257, 258, 1, 0, 0, 0, 258, 260, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0,
		260, 261, 3, 76, 38, 0, 261, 262, 5, 93, 0, 0, 262, 19, 1, 0, 0, 0, 263,
		264, 5, 8, 0, 0, 264, 265, 5, 92, 0, 0, 265, 266, 3, 2, 1, 0, 266, 267,
		5, 93, 0, 0, 267, 21, 1, 0, 0, 0, 268, 269, 5, 9, 0, 0, 269, 270, 5, 92,
		0, 0, 270, 271, 3, 2, 1, 0, 271, 272, 5, 93, 0, 0, 272, 23, 1, 0, 0, 0,
		273, 274, 5, 10, 0, 0, 274, 275, 5, 92, 0, 0, 275, 276, 3, 2, 1, 0, 276,
		277, 5, 95, 0, 0, 277, 278, 3, 2, 1, 0, 278, 279, 5, 93, 0, 0, 279, 25,
		1, 0, 0, 0, 280, 281, 5, 11, 0, 0, 281, 282, 5, 92, 0, 0, 282, 283, 3,
		2, 1, 0, 283, 284, 5, 93, 0, 0, 284, 27, 1, 0, 0, 0, 285, 286, 5, 12, 0,
		0, 286, 287, 5, 92, 0, 0, 287, 288, 3, 78, 39, 0, 288, 289, 5, 95, 0, 0,
		289, 290, 3, 80, 40, 0, 290, 291, 5, 95, 0, 0, 291, 292, 3, 82, 41, 0,
		292, 293, 5, 93, 0, 0, 293, 29, 1, 0, 0, 0, 294, 295, 5, 13, 0, 0, 295,
		296, 5, 92, 0, 0, 296, 297, 3, 2, 1, 0, 297, 298, 5, 93, 0, 0, 298, 31,
		1, 0, 0, 0, 299, 300, 5, 14, 0, 0, 300, 301, 5, 92, 0, 0, 301, 302, 3,
		2, 1, 0, 302, 303, 5, 93, 0, 0, 303, 33, 1, 0, 0, 0, 304, 305, 5, 15, 0,
		0, 305, 306, 5, 92, 0, 0, 306, 307, 3, 2, 1, 0, 307, 308, 5, 93, 0, 0,
		308, 35, 1, 0, 0, 0, 309, 310, 5, 16, 0, 0, 310, 311, 5, 92, 0, 0, 311,
		312, 3, 2, 1, 0, 312, 313, 5, 95, 0, 0, 313, 314, 3, 2, 1, 0, 314, 315,
		5, 95, 0, 0, 315, 316, 3, 84, 42, 0, 316, 317, 5, 93, 0, 0, 317, 37, 1,
		0, 0, 0, 318, 319, 5, 17, 0, 0, 319, 320, 5, 92, 0, 0, 320, 321, 3, 2,
		1, 0, 321, 322, 5, 93, 0, 0, 322, 39, 1, 0, 0, 0, 323, 324, 5, 18, 0, 0,
		324, 325, 5, 92, 0, 0, 325, 326, 3, 86, 43, 0, 326, 327, 5, 95, 0, 0, 327,
		330, 3, 88, 44, 0, 328, 329, 5, 95, 0, 0, 329, 331, 3, 90, 45, 0, 330,
		328, 1, 0, 0, 0, 330, 331, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332, 333,
		5, 93, 0, 0, 333, 41, 1, 0, 0, 0, 334, 335, 5, 19, 0, 0, 335, 336, 5, 92,
		0, 0, 336, 337, 3, 2, 1, 0, 337, 338, 5, 93, 0, 0, 338, 43, 1, 0, 0, 0,
		339, 340, 5, 20, 0, 0, 340, 341, 5, 92, 0, 0, 341, 342, 3, 92, 46, 0, 342,
		343, 5, 95, 0, 0, 343, 344, 3, 94, 47, 0, 344, 345, 5, 93, 0, 0, 345, 45,
		1, 0, 0, 0, 346, 347, 5, 21, 0, 0, 347, 348, 5, 92, 0, 0, 348, 349, 3,
		2, 1, 0, 349, 350, 5, 93, 0, 0, 350, 47, 1, 0, 0, 0, 351, 352, 5, 22, 0,
		0, 352, 353, 5, 92, 0, 0, 353, 354, 5, 93, 0, 0, 354, 49, 1, 0, 0, 0, 355,
		356, 5, 23, 0, 0, 356, 357, 5, 92, 0, 0, 357, 358, 3, 2, 1, 0, 358, 359,
		5, 93, 0, 0, 359, 51, 1, 0, 0, 0, 360, 361, 5, 24, 0, 0, 361, 362, 5, 92,
		0, 0, 362, 363, 3, 2, 1, 0, 363, 364, 5, 93, 0, 0, 364, 53, 1, 0, 0, 0,
		365, 366, 5, 25, 0, 0, 366, 367, 5, 92, 0, 0, 367, 368, 3, 96, 48, 0, 368,
		369, 5, 95, 0, 0, 369, 372, 3, 98, 49, 0, 370, 371, 5, 95, 0, 0, 371, 373,
		3, 100, 50, 0, 372, 370, 1, 0, 0, 0, 372, 373, 1, 0, 0, 0, 373, 374, 1,
		0, 0, 0, 374, 375, 5, 93, 0, 0, 375, 55, 1, 0, 0, 0, 376, 377, 5, 26, 0,
		0, 377, 378, 5, 92, 0, 0, 378, 379, 3, 102, 51, 0, 379, 380, 5, 95, 0,
		0, 380, 381, 3, 104, 52, 0, 381, 382, 5, 95, 0, 0, 382, 383, 3, 106, 53,
		0, 383, 384, 5, 93, 0, 0, 384, 57, 1, 0, 0, 0, 385, 386, 5, 27, 0, 0, 386,
		387, 5, 92, 0, 0, 387, 388, 3, 96, 48, 0, 388, 389, 5, 95, 0, 0, 389, 395,
		3, 88, 44, 0, 390, 391, 5, 95, 0, 0, 391, 392, 3, 108, 54, 0, 392, 393,
		5, 95, 0, 0, 393, 394, 3, 110, 55, 0, 394, 396, 1, 0, 0, 0, 395, 390, 1,
		0, 0, 0, 395, 396, 1, 0, 0, 0, 396, 397, 1, 0, 0, 0, 397, 398, 5, 93, 0,
		0, 398, 59, 1, 0, 0, 0, 399, 400, 5, 28, 0, 0, 400, 401, 5, 92, 0, 0, 401,
		402, 3, 96, 48, 0, 402, 403, 5, 93, 0, 0, 403, 61, 1, 0, 0, 0, 404, 405,
		5, 51, 0, 0, 405, 406, 5, 92, 0, 0, 406, 407, 3, 88, 44, 0, 407, 408, 5,
		95, 0, 0, 408, 409, 3, 90, 45, 0, 409, 410, 5, 95, 0, 0, 410, 411, 3, 114,
		57, 0, 411, 412, 5, 93, 0, 0, 412, 63, 1, 0, 0, 0, 413, 414, 5, 30, 0,
		0, 414, 415, 5, 92, 0, 0, 415, 416, 3, 70, 35, 0, 416, 417, 5, 95, 0, 0,
		417, 418, 3, 72, 36, 0, 418, 419, 5, 93, 0, 0, 419, 65, 1, 0, 0, 0, 420,
		421, 5, 74, 0, 0, 421, 422, 5, 92, 0, 0, 422, 423, 3, 2, 1, 0, 423, 424,
		5, 93, 0, 0, 424, 67, 1, 0, 0, 0, 425, 426, 5, 82, 0, 0, 426, 427, 5, 92,
		0, 0, 427, 428, 3, 2, 1, 0, 428, 429, 5, 93, 0, 0, 429, 69, 1, 0, 0, 0,
		430, 431, 3, 2, 1, 0, 431, 71, 1, 0, 0, 0, 432, 433, 3, 2, 1, 0, 433, 73,
		1, 0, 0, 0, 434, 435, 3, 2, 1, 0, 435, 75, 1, 0, 0, 0, 436, 437, 3, 2,
		1, 0, 437, 77, 1, 0, 0, 0, 438, 439, 3, 2, 1, 0, 439, 79, 1, 0, 0, 0, 440,
		441, 3, 2, 1, 0, 441, 81, 1, 0, 0, 0, 442, 443, 3, 2, 1, 0, 443, 83, 1,
		0, 0, 0, 444, 445, 3, 2, 1, 0, 445, 85, 1, 0, 0, 0, 446, 447, 3, 2, 1,
		0, 447, 87, 1, 0, 0, 0, 448, 449, 3, 2, 1, 0, 449, 89, 1, 0, 0, 0, 450,
		451, 3, 2, 1, 0, 451, 91, 1, 0, 0, 0, 452, 453, 3, 2, 1, 0, 453, 93, 1,
		0, 0, 0, 454, 455, 3, 2, 1, 0, 455, 95, 1, 0, 0, 0, 456, 457, 3, 2, 1,
		0, 457, 97, 1, 0, 0, 0, 458, 459, 3, 2, 1, 0, 459, 99, 1, 0, 0, 0, 460,
		461, 3, 2, 1, 0, 461, 101, 1, 0, 0, 0, 462, 463, 3, 2, 1, 0, 463, 103,
		1, 0, 0, 0, 464, 465, 3, 2, 1, 0, 465, 105, 1, 0, 0, 0, 466, 467, 3, 2,
		1, 0, 467, 107, 1, 0, 0, 0, 468, 469, 3, 2, 1, 0, 469, 109, 1, 0, 0, 0,
		470, 471, 3, 2, 1, 0, 471, 111, 1, 0, 0, 0, 472, 473, 3, 2, 1, 0, 473,
		113, 1, 0, 0, 0, 474, 475, 3, 2, 1, 0, 475, 115, 1, 0, 0, 0, 476, 482,
		3, 120, 60, 0, 477, 478, 5, 92, 0, 0, 478, 479, 3, 2, 1, 0, 479, 480, 5,
		93, 0, 0, 480, 482, 1, 0, 0, 0, 481, 476, 1, 0, 0, 0, 481, 477, 1, 0, 0,
		0, 482, 117, 1, 0, 0, 0, 483, 488, 3, 2, 1, 0, 484, 485, 5, 95, 0, 0, 485,
		487, 3, 2, 1, 0, 486, 484, 1, 0, 0, 0, 487, 490, 1, 0, 0, 0, 488, 486,
		1, 0, 0, 0, 488, 489, 1, 0, 0, 0, 489, 119, 1, 0, 0, 0, 490, 488, 1, 0,
		0, 0, 491, 492, 7, 4, 0, 0, 492, 121, 1, 0, 0, 0, 13, 134, 137, 151, 160,
		162, 197, 220, 257, 330, 372, 395, 481, 488,
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

// FormulaParserInit initializes any static state used to implement FormulaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFormulaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FormulaParserInit() {
	staticData := &FormulaParserParserStaticData
	staticData.once.Do(formulaparserParserInit)
}

// NewFormulaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFormulaParser(input antlr.TokenStream) *FormulaParser {
	FormulaParserInit()
	this := new(FormulaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &FormulaParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "FormulaParser.g4"

	return this
}

// FormulaParser tokens.
const (
	FormulaParserEOF                  = antlr.TokenEOF
	FormulaParserABS                  = 1
	FormulaParserADDMONTHS            = 2
	FormulaParserAND_FUNC             = 3
	FormulaParserBEGINS               = 4
	FormulaParserBLANKVALUE           = 5
	FormulaParserBR                   = 6
	FormulaParserCASE                 = 7
	FormulaParserCASESAFEID           = 8
	FormulaParserCEILING              = 9
	FormulaParserCONTAINS             = 10
	FormulaParserCURRENCYRATE         = 11
	FormulaParserDATE                 = 12
	FormulaParserDATEVALUE            = 13
	FormulaParserDATETIMEVALUE        = 14
	FormulaParserDAY                  = 15
	FormulaParserDISTANCE             = 16
	FormulaParserEXP                  = 17
	FormulaParserFIND                 = 18
	FormulaParserFLOOR                = 19
	FormulaParserGEOLOCATION          = 20
	FormulaParserGETRECORDIDS         = 21
	FormulaParserGETSESSIONID         = 22
	FormulaParserHOUR                 = 23
	FormulaParserHTMLENCODE           = 24
	FormulaParserHYPERLINK            = 25
	FormulaParserIF                   = 26
	FormulaParserIMAGE                = 27
	FormulaParserIMAGEPROXYURL        = 28
	FormulaParserINCLUDE              = 29
	FormulaParserINCLUDES             = 30
	FormulaParserISBLANK              = 31
	FormulaParserISCHANGED            = 32
	FormulaParserISCLONE              = 33
	FormulaParserISNEW                = 34
	FormulaParserISNULL               = 35
	FormulaParserISNUMBER             = 36
	FormulaParserISPICKVAL            = 37
	FormulaParserJSENCODE             = 38
	FormulaParserJSINHTMLENCODE       = 39
	FormulaParserJUNCTIONIDLIST       = 40
	FormulaParserLEFT                 = 41
	FormulaParserLEN                  = 42
	FormulaParserLINKTO               = 43
	FormulaParserLN                   = 44
	FormulaParserLOG                  = 45
	FormulaParserLOWER                = 46
	FormulaParserLPAD                 = 47
	FormulaParserMAX                  = 48
	FormulaParserMCEILING             = 49
	FormulaParserMFLOOR               = 50
	FormulaParserMID                  = 51
	FormulaParserMILLISECOND          = 52
	FormulaParserMIN                  = 53
	FormulaParserMINUTE               = 54
	FormulaParserMOD                  = 55
	FormulaParserMONTH                = 56
	FormulaParserNOT                  = 57
	FormulaParserNOW                  = 58
	FormulaParserNULLVALUE            = 59
	FormulaParserOR_FUNC              = 60
	FormulaParserPARENTGROUPVAL       = 61
	FormulaParserPREDICT              = 62
	FormulaParserPREVGROUPVAL         = 63
	FormulaParserPRIORVALUE           = 64
	FormulaParserREGEX                = 65
	FormulaParserREQUIRESCRIPT        = 66
	FormulaParserREVERSE              = 67
	FormulaParserRIGHT                = 68
	FormulaParserROUND                = 69
	FormulaParserRPAD                 = 70
	FormulaParserSECOND               = 71
	FormulaParserSQRT                 = 72
	FormulaParserSUBSTITUTE           = 73
	FormulaParserTEXT                 = 74
	FormulaParserTIMENOW              = 75
	FormulaParserTIMEVALUE            = 76
	FormulaParserTODAY                = 77
	FormulaParserTRIM                 = 78
	FormulaParserUPPER                = 79
	FormulaParserURLENCODE            = 80
	FormulaParserURLFOR               = 81
	FormulaParserVALUE                = 82
	FormulaParserVLOOKUP              = 83
	FormulaParserWEEKDAY              = 84
	FormulaParserYEAR                 = 85
	FormulaParserStringLiteral        = 86
	FormulaParserIntegerLiteral       = 87
	FormulaParserFloatingPointLiteral = 88
	FormulaParserBooleanLiteral       = 89
	FormulaParserNullLiteral          = 90
	FormulaParserIdentifier           = 91
	FormulaParserLPAREN               = 92
	FormulaParserRPAREN               = 93
	FormulaParserDOT                  = 94
	FormulaParserCOMMA                = 95
	FormulaParserCARET                = 96
	FormulaParserBITAND               = 97
	FormulaParserADD                  = 98
	FormulaParserSUB                  = 99
	FormulaParserMUL                  = 100
	FormulaParserDIV                  = 101
	FormulaParserLT                   = 102
	FormulaParserGT                   = 103
	FormulaParserASSIGN               = 104
	FormulaParserLESSANDGREATER       = 105
	FormulaParserNOTEQUAL             = 106
	FormulaParserEQUAL                = 107
	FormulaParserAND                  = 108
	FormulaParserOR                   = 109
	FormulaParserWS                   = 110
	FormulaParserCOMMENT              = 111
)

// FormulaParser rules.
const (
	FormulaParserRULE_compilationUnit     = 0
	FormulaParserRULE_expression          = 1
	FormulaParserRULE_functionCall        = 2
	FormulaParserRULE_abs                 = 3
	FormulaParserRULE_addMonths           = 4
	FormulaParserRULE_and                 = 5
	FormulaParserRULE_begins              = 6
	FormulaParserRULE_blankvalue          = 7
	FormulaParserRULE_br                  = 8
	FormulaParserRULE_case                = 9
	FormulaParserRULE_casesafeid          = 10
	FormulaParserRULE_ceiling             = 11
	FormulaParserRULE_contains            = 12
	FormulaParserRULE_currencyrate        = 13
	FormulaParserRULE_date                = 14
	FormulaParserRULE_datevalue           = 15
	FormulaParserRULE_datetimevalue       = 16
	FormulaParserRULE_day                 = 17
	FormulaParserRULE_distance            = 18
	FormulaParserRULE_exp                 = 19
	FormulaParserRULE_find                = 20
	FormulaParserRULE_floor               = 21
	FormulaParserRULE_geolocation         = 22
	FormulaParserRULE_getrecordids        = 23
	FormulaParserRULE_getsessionid        = 24
	FormulaParserRULE_hour                = 25
	FormulaParserRULE_htmlencode          = 26
	FormulaParserRULE_hyperlink           = 27
	FormulaParserRULE_if                  = 28
	FormulaParserRULE_image               = 29
	FormulaParserRULE_imageproxyurl       = 30
	FormulaParserRULE_mid                 = 31
	FormulaParserRULE_includes            = 32
	FormulaParserRULE_text                = 33
	FormulaParserRULE_value               = 34
	FormulaParserRULE_fieldExpression     = 35
	FormulaParserRULE_valueExpression     = 36
	FormulaParserRULE_resultExpression    = 37
	FormulaParserRULE_defaultExpression   = 38
	FormulaParserRULE_yearExpression      = 39
	FormulaParserRULE_monthExpression     = 40
	FormulaParserRULE_dayExpression       = 41
	FormulaParserRULE_unitExpression      = 42
	FormulaParserRULE_searchExpression    = 43
	FormulaParserRULE_textExpression      = 44
	FormulaParserRULE_startNum            = 45
	FormulaParserRULE_latitudeExpression  = 46
	FormulaParserRULE_longitudeExpression = 47
	FormulaParserRULE_urlExpression       = 48
	FormulaParserRULE_nameExpression      = 49
	FormulaParserRULE_targetExpression    = 50
	FormulaParserRULE_logicalExpression   = 51
	FormulaParserRULE_ifTrueExpression    = 52
	FormulaParserRULE_ifFalseExpression   = 53
	FormulaParserRULE_heightExpression    = 54
	FormulaParserRULE_widthExpression     = 55
	FormulaParserRULE_start               = 56
	FormulaParserRULE_numChars            = 57
	FormulaParserRULE_primary             = 58
	FormulaParserRULE_arguments           = 59
	FormulaParserRULE_literal             = 60
)

// ICompilationUnitContext is an interface to support dynamic dispatch.
type ICompilationUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsCompilationUnitContext differentiates from other interfaces.
	IsCompilationUnitContext()
}

type CompilationUnitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompilationUnitContext() *CompilationUnitContext {
	var p = new(CompilationUnitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_compilationUnit
	return p
}

func InitEmptyCompilationUnitContext(p *CompilationUnitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_compilationUnit
}

func (*CompilationUnitContext) IsCompilationUnitContext() {}

func NewCompilationUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompilationUnitContext {
	var p = new(CompilationUnitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_compilationUnit

	return p
}

func (s *CompilationUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *CompilationUnitContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CompilationUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompilationUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompilationUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitCompilationUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) CompilationUnit() (localctx ICompilationUnitContext) {
	localctx = NewCompilationUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FormulaParserRULE_compilationUnit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrimaryExpressionContext struct {
	ExpressionContext
}

func NewPrimaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitPrimaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegationExpressionContext struct {
	ExpressionContext
}

func NewNegationExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegationExpressionContext {
	var p = new(NegationExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NegationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegationExpressionContext) SUB() antlr.TerminalNode {
	return s.GetToken(FormulaParserSUB, 0)
}

func (s *NegationExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NegationExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterNegationExpression(s)
	}
}

func (s *NegationExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitNegationExpression(s)
	}
}

func (s *NegationExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitNegationExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompareExpressionContext struct {
	ExpressionContext
}

func NewCompareExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareExpressionContext {
	var p = new(CompareExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CompareExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *CompareExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CompareExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(FormulaParserGT, 0)
}

func (s *CompareExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(FormulaParserLT, 0)
}

func (s *CompareExpressionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(FormulaParserASSIGN, 0)
}

func (s *CompareExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterCompareExpression(s)
	}
}

func (s *CompareExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitCompareExpression(s)
	}
}

func (s *CompareExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitCompareExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualityExpressionContext struct {
	ExpressionContext
}

func NewEqualityExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *EqualityExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqualityExpressionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(FormulaParserEQUAL, 0)
}

func (s *EqualityExpressionContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(FormulaParserNOTEQUAL, 0)
}

func (s *EqualityExpressionContext) LESSANDGREATER() antlr.TerminalNode {
	return s.GetToken(FormulaParserLESSANDGREATER, 0)
}

func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitEqualityExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicExpressionContext struct {
	ExpressionContext
}

func NewLogicExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicExpressionContext {
	var p = new(LogicExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LogicExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogicExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(FormulaParserAND, 0)
}

func (s *LogicExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(FormulaParserOR, 0)
}

func (s *LogicExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterLogicExpression(s)
	}
}

func (s *LogicExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitLogicExpression(s)
	}
}

func (s *LogicExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitLogicExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallExpressionContext struct {
	ExpressionContext
}

func NewFunctionCallExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallExpressionContext {
	var p = new(FunctionCallExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FunctionCallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallExpressionContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FunctionCallExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterFunctionCallExpression(s)
	}
}

func (s *FunctionCallExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitFunctionCallExpression(s)
	}
}

func (s *FunctionCallExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitFunctionCallExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type FieldReferenceContext struct {
	ExpressionContext
}

func NewFieldReferenceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldReferenceContext {
	var p = new(FieldReferenceContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FieldReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldReferenceContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(FormulaParserIdentifier)
}

func (s *FieldReferenceContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(FormulaParserIdentifier, i)
}

func (s *FieldReferenceContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(FormulaParserDOT)
}

func (s *FieldReferenceContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(FormulaParserDOT, i)
}

func (s *FieldReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterFieldReference(s)
	}
}

func (s *FieldReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitFieldReference(s)
	}
}

func (s *FieldReferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitFieldReference(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExponentiationExpressionContext struct {
	ExpressionContext
}

func NewExponentiationExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExponentiationExpressionContext {
	var p = new(ExponentiationExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExponentiationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExponentiationExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExponentiationExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExponentiationExpressionContext) CARET() antlr.TerminalNode {
	return s.GetToken(FormulaParserCARET, 0)
}

func (s *ExponentiationExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterExponentiationExpression(s)
	}
}

func (s *ExponentiationExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitExponentiationExpression(s)
	}
}

func (s *ExponentiationExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitExponentiationExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArithExpressionContext struct {
	ExpressionContext
}

func NewArithExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArithExpressionContext {
	var p = new(ArithExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ArithExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArithExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArithExpressionContext) ADD() antlr.TerminalNode {
	return s.GetToken(FormulaParserADD, 0)
}

func (s *ArithExpressionContext) SUB() antlr.TerminalNode {
	return s.GetToken(FormulaParserSUB, 0)
}

func (s *ArithExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(FormulaParserDIV, 0)
}

func (s *ArithExpressionContext) MUL() antlr.TerminalNode {
	return s.GetToken(FormulaParserMUL, 0)
}

func (s *ArithExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterArithExpression(s)
	}
}

func (s *ArithExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitArithExpression(s)
	}
}

func (s *ArithExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitArithExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConcatExpressionContext struct {
	ExpressionContext
}

func NewConcatExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConcatExpressionContext {
	var p = new(ConcatExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ConcatExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConcatExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ConcatExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConcatExpressionContext) BITAND() antlr.TerminalNode {
	return s.GetToken(FormulaParserBITAND, 0)
}

func (s *ConcatExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterConcatExpression(s)
	}
}

func (s *ConcatExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitConcatExpression(s)
	}
}

func (s *ConcatExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitConcatExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *FormulaParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, FormulaParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FormulaParserStringLiteral, FormulaParserIntegerLiteral, FormulaParserFloatingPointLiteral, FormulaParserBooleanLiteral, FormulaParserNullLiteral, FormulaParserLPAREN:
		localctx = NewPrimaryExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(125)
			p.Primary()
		}

	case FormulaParserABS, FormulaParserADDMONTHS, FormulaParserAND_FUNC, FormulaParserBEGINS, FormulaParserBLANKVALUE, FormulaParserBR, FormulaParserCASE, FormulaParserCASESAFEID, FormulaParserCEILING, FormulaParserCONTAINS, FormulaParserCURRENCYRATE, FormulaParserDATE, FormulaParserDATEVALUE, FormulaParserDATETIMEVALUE, FormulaParserDAY, FormulaParserDISTANCE, FormulaParserEXP, FormulaParserFIND, FormulaParserFLOOR, FormulaParserGEOLOCATION, FormulaParserGETRECORDIDS, FormulaParserGETSESSIONID, FormulaParserHOUR, FormulaParserHTMLENCODE, FormulaParserHYPERLINK, FormulaParserIF, FormulaParserIMAGE, FormulaParserIMAGEPROXYURL, FormulaParserINCLUDES, FormulaParserMID, FormulaParserTEXT, FormulaParserVALUE:
		localctx = NewFunctionCallExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(126)
			p.FunctionCall()
		}

	case FormulaParserSUB:
		localctx = NewNegationExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(127)
			p.Match(FormulaParserSUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.expression(8)
		}

	case FormulaParserIdentifier:
		localctx = NewFieldReferenceContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(129)
			p.Match(FormulaParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(130)
					p.Match(FormulaParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(131)
					p.Match(FormulaParserIdentifier)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(136)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(160)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExponentiationExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expression)
				p.SetState(139)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(140)
					p.Match(FormulaParserCARET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(141)
					p.expression(7)
				}

			case 2:
				localctx = NewConcatExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expression)
				p.SetState(142)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(143)
					p.Match(FormulaParserBITAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(144)
					p.expression(6)
				}

			case 3:
				localctx = NewArithExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expression)
				p.SetState(145)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(146)
					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-98)) & ^0x3f) == 0 && ((int64(1)<<(_la-98))&15) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(147)
					p.expression(5)
				}

			case 4:
				localctx = NewCompareExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expression)
				p.SetState(148)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(149)
					_la = p.GetTokenStream().LA(1)

					if !(_la == FormulaParserLT || _la == FormulaParserGT) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(151)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == FormulaParserASSIGN {
					{
						p.SetState(150)
						p.Match(FormulaParserASSIGN)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(153)
					p.expression(4)
				}

			case 5:
				localctx = NewEqualityExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expression)
				p.SetState(154)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(155)
					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-105)) & ^0x3f) == 0 && ((int64(1)<<(_la-105))&7) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(156)
					p.expression(3)
				}

			case 6:
				localctx = NewLogicExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expression)
				p.SetState(157)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(158)
					_la = p.GetTokenStream().LA(1)

					if !(_la == FormulaParserAND || _la == FormulaParserOR) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(159)
					p.expression(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Abs() IAbsContext
	AddMonths() IAddMonthsContext
	And() IAndContext
	Begins() IBeginsContext
	Blankvalue() IBlankvalueContext
	Br() IBrContext
	Case_() ICaseContext
	Casesafeid() ICasesafeidContext
	Ceiling() ICeilingContext
	Contains() IContainsContext
	Currencyrate() ICurrencyrateContext
	Date() IDateContext
	Datevalue() IDatevalueContext
	Datetimevalue() IDatetimevalueContext
	Day() IDayContext
	Distance() IDistanceContext
	Exp() IExpContext
	Find() IFindContext
	Floor() IFloorContext
	Geolocation() IGeolocationContext
	Getrecordids() IGetrecordidsContext
	Getsessionid() IGetsessionidContext
	Hour() IHourContext
	Htmlencode() IHtmlencodeContext
	Hyperlink() IHyperlinkContext
	If_() IIfContext
	Image() IImageContext
	Imageproxyurl() IImageproxyurlContext
	Includes() IIncludesContext
	Mid() IMidContext
	Text() ITextContext
	Value() IValueContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) Abs() IAbsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAbsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAbsContext)
}

func (s *FunctionCallContext) AddMonths() IAddMonthsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddMonthsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddMonthsContext)
}

func (s *FunctionCallContext) And() IAndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndContext)
}

func (s *FunctionCallContext) Begins() IBeginsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBeginsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBeginsContext)
}

func (s *FunctionCallContext) Blankvalue() IBlankvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlankvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlankvalueContext)
}

func (s *FunctionCallContext) Br() IBrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBrContext)
}

func (s *FunctionCallContext) Case_() ICaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseContext)
}

func (s *FunctionCallContext) Casesafeid() ICasesafeidContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICasesafeidContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICasesafeidContext)
}

func (s *FunctionCallContext) Ceiling() ICeilingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICeilingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICeilingContext)
}

func (s *FunctionCallContext) Contains() IContainsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContainsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContainsContext)
}

func (s *FunctionCallContext) Currencyrate() ICurrencyrateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICurrencyrateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICurrencyrateContext)
}

func (s *FunctionCallContext) Date() IDateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateContext)
}

func (s *FunctionCallContext) Datevalue() IDatevalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatevalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatevalueContext)
}

func (s *FunctionCallContext) Datetimevalue() IDatetimevalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatetimevalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatetimevalueContext)
}

func (s *FunctionCallContext) Day() IDayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDayContext)
}

func (s *FunctionCallContext) Distance() IDistanceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDistanceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDistanceContext)
}

func (s *FunctionCallContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *FunctionCallContext) Find() IFindContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFindContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFindContext)
}

func (s *FunctionCallContext) Floor() IFloorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloorContext)
}

func (s *FunctionCallContext) Geolocation() IGeolocationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGeolocationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGeolocationContext)
}

func (s *FunctionCallContext) Getrecordids() IGetrecordidsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGetrecordidsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGetrecordidsContext)
}

func (s *FunctionCallContext) Getsessionid() IGetsessionidContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGetsessionidContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGetsessionidContext)
}

func (s *FunctionCallContext) Hour() IHourContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHourContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHourContext)
}

func (s *FunctionCallContext) Htmlencode() IHtmlencodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHtmlencodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHtmlencodeContext)
}

func (s *FunctionCallContext) Hyperlink() IHyperlinkContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHyperlinkContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHyperlinkContext)
}

func (s *FunctionCallContext) If_() IIfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfContext)
}

func (s *FunctionCallContext) Image() IImageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImageContext)
}

func (s *FunctionCallContext) Imageproxyurl() IImageproxyurlContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImageproxyurlContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImageproxyurlContext)
}

func (s *FunctionCallContext) Includes() IIncludesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncludesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncludesContext)
}

func (s *FunctionCallContext) Mid() IMidContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMidContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMidContext)
}

func (s *FunctionCallContext) Text() ITextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *FunctionCallContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FormulaParserRULE_functionCall)
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FormulaParserABS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(165)
			p.Abs()
		}

	case FormulaParserADDMONTHS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(166)
			p.AddMonths()
		}

	case FormulaParserAND_FUNC:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(167)
			p.And()
		}

	case FormulaParserBEGINS:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(168)
			p.Begins()
		}

	case FormulaParserBLANKVALUE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(169)
			p.Blankvalue()
		}

	case FormulaParserBR:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(170)
			p.Br()
		}

	case FormulaParserCASE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(171)
			p.Case_()
		}

	case FormulaParserCASESAFEID:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(172)
			p.Casesafeid()
		}

	case FormulaParserCEILING:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(173)
			p.Ceiling()
		}

	case FormulaParserCONTAINS:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(174)
			p.Contains()
		}

	case FormulaParserCURRENCYRATE:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(175)
			p.Currencyrate()
		}

	case FormulaParserDATE:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(176)
			p.Date()
		}

	case FormulaParserDATEVALUE:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(177)
			p.Datevalue()
		}

	case FormulaParserDATETIMEVALUE:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(178)
			p.Datetimevalue()
		}

	case FormulaParserDAY:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(179)
			p.Day()
		}

	case FormulaParserDISTANCE:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(180)
			p.Distance()
		}

	case FormulaParserEXP:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(181)
			p.Exp()
		}

	case FormulaParserFIND:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(182)
			p.Find()
		}

	case FormulaParserFLOOR:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(183)
			p.Floor()
		}

	case FormulaParserGEOLOCATION:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(184)
			p.Geolocation()
		}

	case FormulaParserGETRECORDIDS:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(185)
			p.Getrecordids()
		}

	case FormulaParserGETSESSIONID:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(186)
			p.Getsessionid()
		}

	case FormulaParserHOUR:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(187)
			p.Hour()
		}

	case FormulaParserHTMLENCODE:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(188)
			p.Htmlencode()
		}

	case FormulaParserHYPERLINK:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(189)
			p.Hyperlink()
		}

	case FormulaParserIF:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(190)
			p.If_()
		}

	case FormulaParserIMAGE:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(191)
			p.Image()
		}

	case FormulaParserIMAGEPROXYURL:
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(192)
			p.Imageproxyurl()
		}

	case FormulaParserINCLUDES:
		p.EnterOuterAlt(localctx, 29)
		{
			p.SetState(193)
			p.Includes()
		}

	case FormulaParserMID:
		p.EnterOuterAlt(localctx, 30)
		{
			p.SetState(194)
			p.Mid()
		}

	case FormulaParserTEXT:
		p.EnterOuterAlt(localctx, 31)
		{
			p.SetState(195)
			p.Text()
		}

	case FormulaParserVALUE:
		p.EnterOuterAlt(localctx, 32)
		{
			p.SetState(196)
			p.Value()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAbsContext is an interface to support dynamic dispatch.
type IAbsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ABS() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsAbsContext differentiates from other interfaces.
	IsAbsContext()
}

type AbsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAbsContext() *AbsContext {
	var p = new(AbsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_abs
	return p
}

func InitEmptyAbsContext(p *AbsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_abs
}

func (*AbsContext) IsAbsContext() {}

func NewAbsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AbsContext {
	var p = new(AbsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_abs

	return p
}

func (s *AbsContext) GetParser() antlr.Parser { return s.parser }

func (s *AbsContext) ABS() antlr.TerminalNode {
	return s.GetToken(FormulaParserABS, 0)
}

func (s *AbsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *AbsContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AbsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *AbsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AbsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterAbs(s)
	}
}

func (s *AbsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitAbs(s)
	}
}

func (s *AbsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitAbs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Abs() (localctx IAbsContext) {
	localctx = NewAbsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FormulaParserRULE_abs)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(FormulaParserABS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.expression(0)
	}
	{
		p.SetState(202)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddMonthsContext is an interface to support dynamic dispatch.
type IAddMonthsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ADDMONTHS() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	COMMA() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsAddMonthsContext differentiates from other interfaces.
	IsAddMonthsContext()
}

type AddMonthsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddMonthsContext() *AddMonthsContext {
	var p = new(AddMonthsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_addMonths
	return p
}

func InitEmptyAddMonthsContext(p *AddMonthsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_addMonths
}

func (*AddMonthsContext) IsAddMonthsContext() {}

func NewAddMonthsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddMonthsContext {
	var p = new(AddMonthsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_addMonths

	return p
}

func (s *AddMonthsContext) GetParser() antlr.Parser { return s.parser }

func (s *AddMonthsContext) ADDMONTHS() antlr.TerminalNode {
	return s.GetToken(FormulaParserADDMONTHS, 0)
}

func (s *AddMonthsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *AddMonthsContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AddMonthsContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddMonthsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, 0)
}

func (s *AddMonthsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *AddMonthsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddMonthsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddMonthsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterAddMonths(s)
	}
}

func (s *AddMonthsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitAddMonths(s)
	}
}

func (s *AddMonthsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitAddMonths(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) AddMonths() (localctx IAddMonthsContext) {
	localctx = NewAddMonthsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FormulaParserRULE_addMonths)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Match(FormulaParserADDMONTHS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.expression(0)
	}
	{
		p.SetState(207)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.expression(0)
	}
	{
		p.SetState(209)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAndContext is an interface to support dynamic dispatch.
type IAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AND_FUNC() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsAndContext differentiates from other interfaces.
	IsAndContext()
}

type AndContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndContext() *AndContext {
	var p = new(AndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_and
	return p
}

func InitEmptyAndContext(p *AndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_and
}

func (*AndContext) IsAndContext() {}

func NewAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndContext {
	var p = new(AndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_and

	return p
}

func (s *AndContext) GetParser() antlr.Parser { return s.parser }

func (s *AndContext) AND_FUNC() antlr.TerminalNode {
	return s.GetToken(FormulaParserAND_FUNC, 0)
}

func (s *AndContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *AndContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AndContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AndContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FormulaParserCOMMA)
}

func (s *AndContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, i)
}

func (s *AndContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitAnd(s)
	}
}

func (s *AndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) And() (localctx IAndContext) {
	localctx = NewAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FormulaParserRULE_and)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.Match(FormulaParserAND_FUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(212)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.expression(0)
	}
	{
		p.SetState(214)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.expression(0)
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FormulaParserCOMMA {
		{
			p.SetState(216)
			p.Match(FormulaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(217)
			p.expression(0)
		}

		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(223)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBeginsContext is an interface to support dynamic dispatch.
type IBeginsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BEGINS() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	COMMA() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsBeginsContext differentiates from other interfaces.
	IsBeginsContext()
}

type BeginsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBeginsContext() *BeginsContext {
	var p = new(BeginsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_begins
	return p
}

func InitEmptyBeginsContext(p *BeginsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_begins
}

func (*BeginsContext) IsBeginsContext() {}

func NewBeginsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BeginsContext {
	var p = new(BeginsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_begins

	return p
}

func (s *BeginsContext) GetParser() antlr.Parser { return s.parser }

func (s *BeginsContext) BEGINS() antlr.TerminalNode {
	return s.GetToken(FormulaParserBEGINS, 0)
}

func (s *BeginsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *BeginsContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BeginsContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BeginsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, 0)
}

func (s *BeginsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *BeginsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BeginsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BeginsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterBegins(s)
	}
}

func (s *BeginsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitBegins(s)
	}
}

func (s *BeginsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitBegins(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Begins() (localctx IBeginsContext) {
	localctx = NewBeginsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FormulaParserRULE_begins)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)
		p.Match(FormulaParserBEGINS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(226)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(227)
		p.expression(0)
	}
	{
		p.SetState(228)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.expression(0)
	}
	{
		p.SetState(230)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlankvalueContext is an interface to support dynamic dispatch.
type IBlankvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BLANKVALUE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	COMMA() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsBlankvalueContext differentiates from other interfaces.
	IsBlankvalueContext()
}

type BlankvalueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlankvalueContext() *BlankvalueContext {
	var p = new(BlankvalueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_blankvalue
	return p
}

func InitEmptyBlankvalueContext(p *BlankvalueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_blankvalue
}

func (*BlankvalueContext) IsBlankvalueContext() {}

func NewBlankvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlankvalueContext {
	var p = new(BlankvalueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_blankvalue

	return p
}

func (s *BlankvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *BlankvalueContext) BLANKVALUE() antlr.TerminalNode {
	return s.GetToken(FormulaParserBLANKVALUE, 0)
}

func (s *BlankvalueContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *BlankvalueContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BlankvalueContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BlankvalueContext) COMMA() antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, 0)
}

func (s *BlankvalueContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *BlankvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlankvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlankvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterBlankvalue(s)
	}
}

func (s *BlankvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitBlankvalue(s)
	}
}

func (s *BlankvalueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitBlankvalue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Blankvalue() (localctx IBlankvalueContext) {
	localctx = NewBlankvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FormulaParserRULE_blankvalue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.Match(FormulaParserBLANKVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(233)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(234)
		p.expression(0)
	}
	{
		p.SetState(235)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(236)
		p.expression(0)
	}
	{
		p.SetState(237)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBrContext is an interface to support dynamic dispatch.
type IBrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BR() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsBrContext differentiates from other interfaces.
	IsBrContext()
}

type BrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBrContext() *BrContext {
	var p = new(BrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_br
	return p
}

func InitEmptyBrContext(p *BrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_br
}

func (*BrContext) IsBrContext() {}

func NewBrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BrContext {
	var p = new(BrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_br

	return p
}

func (s *BrContext) GetParser() antlr.Parser { return s.parser }

func (s *BrContext) BR() antlr.TerminalNode {
	return s.GetToken(FormulaParserBR, 0)
}

func (s *BrContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *BrContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *BrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterBr(s)
	}
}

func (s *BrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitBr(s)
	}
}

func (s *BrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitBr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Br() (localctx IBrContext) {
	localctx = NewBrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FormulaParserRULE_br)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(FormulaParserBR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(240)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(241)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICaseContext is an interface to support dynamic dispatch.
type ICaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllValueExpression() []IValueExpressionContext
	ValueExpression(i int) IValueExpressionContext
	AllResultExpression() []IResultExpressionContext
	ResultExpression(i int) IResultExpressionContext
	DefaultExpression() IDefaultExpressionContext
	RPAREN() antlr.TerminalNode

	// IsCaseContext differentiates from other interfaces.
	IsCaseContext()
}

type CaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseContext() *CaseContext {
	var p = new(CaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_case
	return p
}

func InitEmptyCaseContext(p *CaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_case
}

func (*CaseContext) IsCaseContext() {}

func NewCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseContext {
	var p = new(CaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_case

	return p
}

func (s *CaseContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseContext) CASE() antlr.TerminalNode {
	return s.GetToken(FormulaParserCASE, 0)
}

func (s *CaseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *CaseContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CaseContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FormulaParserCOMMA)
}

func (s *CaseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, i)
}

func (s *CaseContext) AllValueExpression() []IValueExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueExpressionContext); ok {
			len++
		}
	}

	tst := make([]IValueExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueExpressionContext); ok {
			tst[i] = t.(IValueExpressionContext)
			i++
		}
	}

	return tst
}

func (s *CaseContext) ValueExpression(i int) IValueExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExpressionContext)
}

func (s *CaseContext) AllResultExpression() []IResultExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IResultExpressionContext); ok {
			len++
		}
	}

	tst := make([]IResultExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IResultExpressionContext); ok {
			tst[i] = t.(IResultExpressionContext)
			i++
		}
	}

	return tst
}

func (s *CaseContext) ResultExpression(i int) IResultExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResultExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResultExpressionContext)
}

func (s *CaseContext) DefaultExpression() IDefaultExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultExpressionContext)
}

func (s *CaseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *CaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterCase(s)
	}
}

func (s *CaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitCase(s)
	}
}

func (s *CaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Case_() (localctx ICaseContext) {
	localctx = NewCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FormulaParserRULE_case)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(FormulaParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
		p.expression(0)
	}
	{
		p.SetState(246)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(247)
		p.ValueExpression()
	}
	{
		p.SetState(248)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.ResultExpression()
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FormulaParserCOMMA {
		{
			p.SetState(250)
			p.Match(FormulaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(251)
			p.ValueExpression()
		}
		{
			p.SetState(252)
			p.Match(FormulaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(253)
			p.ResultExpression()
		}

		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(260)
		p.DefaultExpression()
	}
	{
		p.SetState(261)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICasesafeidContext is an interface to support dynamic dispatch.
type ICasesafeidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASESAFEID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsCasesafeidContext differentiates from other interfaces.
	IsCasesafeidContext()
}

type CasesafeidContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCasesafeidContext() *CasesafeidContext {
	var p = new(CasesafeidContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_casesafeid
	return p
}

func InitEmptyCasesafeidContext(p *CasesafeidContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_casesafeid
}

func (*CasesafeidContext) IsCasesafeidContext() {}

func NewCasesafeidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasesafeidContext {
	var p = new(CasesafeidContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_casesafeid

	return p
}

func (s *CasesafeidContext) GetParser() antlr.Parser { return s.parser }

func (s *CasesafeidContext) CASESAFEID() antlr.TerminalNode {
	return s.GetToken(FormulaParserCASESAFEID, 0)
}

func (s *CasesafeidContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *CasesafeidContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CasesafeidContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *CasesafeidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasesafeidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CasesafeidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterCasesafeid(s)
	}
}

func (s *CasesafeidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitCasesafeid(s)
	}
}

func (s *CasesafeidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitCasesafeid(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Casesafeid() (localctx ICasesafeidContext) {
	localctx = NewCasesafeidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FormulaParserRULE_casesafeid)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		p.Match(FormulaParserCASESAFEID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(264)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(265)
		p.expression(0)
	}
	{
		p.SetState(266)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICeilingContext is an interface to support dynamic dispatch.
type ICeilingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CEILING() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsCeilingContext differentiates from other interfaces.
	IsCeilingContext()
}

type CeilingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCeilingContext() *CeilingContext {
	var p = new(CeilingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_ceiling
	return p
}

func InitEmptyCeilingContext(p *CeilingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_ceiling
}

func (*CeilingContext) IsCeilingContext() {}

func NewCeilingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CeilingContext {
	var p = new(CeilingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_ceiling

	return p
}

func (s *CeilingContext) GetParser() antlr.Parser { return s.parser }

func (s *CeilingContext) CEILING() antlr.TerminalNode {
	return s.GetToken(FormulaParserCEILING, 0)
}

func (s *CeilingContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *CeilingContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CeilingContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *CeilingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CeilingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CeilingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterCeiling(s)
	}
}

func (s *CeilingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitCeiling(s)
	}
}

func (s *CeilingContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitCeiling(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Ceiling() (localctx ICeilingContext) {
	localctx = NewCeilingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FormulaParserRULE_ceiling)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Match(FormulaParserCEILING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(270)
		p.expression(0)
	}
	{
		p.SetState(271)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContainsContext is an interface to support dynamic dispatch.
type IContainsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTAINS() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	COMMA() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsContainsContext differentiates from other interfaces.
	IsContainsContext()
}

type ContainsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContainsContext() *ContainsContext {
	var p = new(ContainsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_contains
	return p
}

func InitEmptyContainsContext(p *ContainsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_contains
}

func (*ContainsContext) IsContainsContext() {}

func NewContainsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContainsContext {
	var p = new(ContainsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_contains

	return p
}

func (s *ContainsContext) GetParser() antlr.Parser { return s.parser }

func (s *ContainsContext) CONTAINS() antlr.TerminalNode {
	return s.GetToken(FormulaParserCONTAINS, 0)
}

func (s *ContainsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *ContainsContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ContainsContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ContainsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, 0)
}

func (s *ContainsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *ContainsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContainsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContainsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterContains(s)
	}
}

func (s *ContainsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitContains(s)
	}
}

func (s *ContainsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitContains(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Contains() (localctx IContainsContext) {
	localctx = NewContainsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FormulaParserRULE_contains)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		p.Match(FormulaParserCONTAINS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(274)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(275)
		p.expression(0)
	}
	{
		p.SetState(276)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(277)
		p.expression(0)
	}
	{
		p.SetState(278)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICurrencyrateContext is an interface to support dynamic dispatch.
type ICurrencyrateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CURRENCYRATE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsCurrencyrateContext differentiates from other interfaces.
	IsCurrencyrateContext()
}

type CurrencyrateContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCurrencyrateContext() *CurrencyrateContext {
	var p = new(CurrencyrateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_currencyrate
	return p
}

func InitEmptyCurrencyrateContext(p *CurrencyrateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_currencyrate
}

func (*CurrencyrateContext) IsCurrencyrateContext() {}

func NewCurrencyrateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CurrencyrateContext {
	var p = new(CurrencyrateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_currencyrate

	return p
}

func (s *CurrencyrateContext) GetParser() antlr.Parser { return s.parser }

func (s *CurrencyrateContext) CURRENCYRATE() antlr.TerminalNode {
	return s.GetToken(FormulaParserCURRENCYRATE, 0)
}

func (s *CurrencyrateContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *CurrencyrateContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CurrencyrateContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *CurrencyrateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CurrencyrateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CurrencyrateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterCurrencyrate(s)
	}
}

func (s *CurrencyrateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitCurrencyrate(s)
	}
}

func (s *CurrencyrateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitCurrencyrate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Currencyrate() (localctx ICurrencyrateContext) {
	localctx = NewCurrencyrateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FormulaParserRULE_currencyrate)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(280)
		p.Match(FormulaParserCURRENCYRATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(281)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(282)
		p.expression(0)
	}
	{
		p.SetState(283)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDateContext is an interface to support dynamic dispatch.
type IDateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DATE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	YearExpression() IYearExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	MonthExpression() IMonthExpressionContext
	DayExpression() IDayExpressionContext
	RPAREN() antlr.TerminalNode

	// IsDateContext differentiates from other interfaces.
	IsDateContext()
}

type DateContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateContext() *DateContext {
	var p = new(DateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_date
	return p
}

func InitEmptyDateContext(p *DateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_date
}

func (*DateContext) IsDateContext() {}

func NewDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateContext {
	var p = new(DateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_date

	return p
}

func (s *DateContext) GetParser() antlr.Parser { return s.parser }

func (s *DateContext) DATE() antlr.TerminalNode {
	return s.GetToken(FormulaParserDATE, 0)
}

func (s *DateContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *DateContext) YearExpression() IYearExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IYearExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IYearExpressionContext)
}

func (s *DateContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FormulaParserCOMMA)
}

func (s *DateContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, i)
}

func (s *DateContext) MonthExpression() IMonthExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMonthExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMonthExpressionContext)
}

func (s *DateContext) DayExpression() IDayExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDayExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDayExpressionContext)
}

func (s *DateContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *DateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterDate(s)
	}
}

func (s *DateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitDate(s)
	}
}

func (s *DateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitDate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Date() (localctx IDateContext) {
	localctx = NewDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FormulaParserRULE_date)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		p.Match(FormulaParserDATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(286)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(287)
		p.YearExpression()
	}
	{
		p.SetState(288)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(289)
		p.MonthExpression()
	}
	{
		p.SetState(290)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(291)
		p.DayExpression()
	}
	{
		p.SetState(292)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDatevalueContext is an interface to support dynamic dispatch.
type IDatevalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DATEVALUE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsDatevalueContext differentiates from other interfaces.
	IsDatevalueContext()
}

type DatevalueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatevalueContext() *DatevalueContext {
	var p = new(DatevalueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_datevalue
	return p
}

func InitEmptyDatevalueContext(p *DatevalueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_datevalue
}

func (*DatevalueContext) IsDatevalueContext() {}

func NewDatevalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatevalueContext {
	var p = new(DatevalueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_datevalue

	return p
}

func (s *DatevalueContext) GetParser() antlr.Parser { return s.parser }

func (s *DatevalueContext) DATEVALUE() antlr.TerminalNode {
	return s.GetToken(FormulaParserDATEVALUE, 0)
}

func (s *DatevalueContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *DatevalueContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DatevalueContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *DatevalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatevalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatevalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterDatevalue(s)
	}
}

func (s *DatevalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitDatevalue(s)
	}
}

func (s *DatevalueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitDatevalue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Datevalue() (localctx IDatevalueContext) {
	localctx = NewDatevalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FormulaParserRULE_datevalue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(294)
		p.Match(FormulaParserDATEVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(295)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(296)
		p.expression(0)
	}
	{
		p.SetState(297)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDatetimevalueContext is an interface to support dynamic dispatch.
type IDatetimevalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DATETIMEVALUE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsDatetimevalueContext differentiates from other interfaces.
	IsDatetimevalueContext()
}

type DatetimevalueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatetimevalueContext() *DatetimevalueContext {
	var p = new(DatetimevalueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_datetimevalue
	return p
}

func InitEmptyDatetimevalueContext(p *DatetimevalueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_datetimevalue
}

func (*DatetimevalueContext) IsDatetimevalueContext() {}

func NewDatetimevalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatetimevalueContext {
	var p = new(DatetimevalueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_datetimevalue

	return p
}

func (s *DatetimevalueContext) GetParser() antlr.Parser { return s.parser }

func (s *DatetimevalueContext) DATETIMEVALUE() antlr.TerminalNode {
	return s.GetToken(FormulaParserDATETIMEVALUE, 0)
}

func (s *DatetimevalueContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *DatetimevalueContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DatetimevalueContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *DatetimevalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatetimevalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatetimevalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterDatetimevalue(s)
	}
}

func (s *DatetimevalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitDatetimevalue(s)
	}
}

func (s *DatetimevalueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitDatetimevalue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Datetimevalue() (localctx IDatetimevalueContext) {
	localctx = NewDatetimevalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FormulaParserRULE_datetimevalue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.Match(FormulaParserDATETIMEVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(300)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(301)
		p.expression(0)
	}
	{
		p.SetState(302)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDayContext is an interface to support dynamic dispatch.
type IDayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DAY() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsDayContext differentiates from other interfaces.
	IsDayContext()
}

type DayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDayContext() *DayContext {
	var p = new(DayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_day
	return p
}

func InitEmptyDayContext(p *DayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_day
}

func (*DayContext) IsDayContext() {}

func NewDayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DayContext {
	var p = new(DayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_day

	return p
}

func (s *DayContext) GetParser() antlr.Parser { return s.parser }

func (s *DayContext) DAY() antlr.TerminalNode {
	return s.GetToken(FormulaParserDAY, 0)
}

func (s *DayContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *DayContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DayContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *DayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterDay(s)
	}
}

func (s *DayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitDay(s)
	}
}

func (s *DayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitDay(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Day() (localctx IDayContext) {
	localctx = NewDayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, FormulaParserRULE_day)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(304)
		p.Match(FormulaParserDAY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(305)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)
		p.expression(0)
	}
	{
		p.SetState(307)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDistanceContext is an interface to support dynamic dispatch.
type IDistanceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DISTANCE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	UnitExpression() IUnitExpressionContext
	RPAREN() antlr.TerminalNode

	// IsDistanceContext differentiates from other interfaces.
	IsDistanceContext()
}

type DistanceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDistanceContext() *DistanceContext {
	var p = new(DistanceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_distance
	return p
}

func InitEmptyDistanceContext(p *DistanceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_distance
}

func (*DistanceContext) IsDistanceContext() {}

func NewDistanceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DistanceContext {
	var p = new(DistanceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_distance

	return p
}

func (s *DistanceContext) GetParser() antlr.Parser { return s.parser }

func (s *DistanceContext) DISTANCE() antlr.TerminalNode {
	return s.GetToken(FormulaParserDISTANCE, 0)
}

func (s *DistanceContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *DistanceContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *DistanceContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DistanceContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FormulaParserCOMMA)
}

func (s *DistanceContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, i)
}

func (s *DistanceContext) UnitExpression() IUnitExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitExpressionContext)
}

func (s *DistanceContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *DistanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DistanceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DistanceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterDistance(s)
	}
}

func (s *DistanceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitDistance(s)
	}
}

func (s *DistanceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitDistance(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Distance() (localctx IDistanceContext) {
	localctx = NewDistanceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, FormulaParserRULE_distance)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		p.Match(FormulaParserDISTANCE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(310)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(311)
		p.expression(0)
	}
	{
		p.SetState(312)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(313)
		p.expression(0)
	}
	{
		p.SetState(314)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(315)
		p.UnitExpression()
	}
	{
		p.SetState(316)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXP() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_exp
	return p
}

func InitEmptyExpContext(p *ExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_exp
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) EXP() antlr.TerminalNode {
	return s.GetToken(FormulaParserEXP, 0)
}

func (s *ExpContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *ExpContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitExp(s)
	}
}

func (s *ExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Exp() (localctx IExpContext) {
	localctx = NewExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, FormulaParserRULE_exp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(318)
		p.Match(FormulaParserEXP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(319)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(320)
		p.expression(0)
	}
	{
		p.SetState(321)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFindContext is an interface to support dynamic dispatch.
type IFindContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FIND() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	SearchExpression() ISearchExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	TextExpression() ITextExpressionContext
	RPAREN() antlr.TerminalNode
	StartNum() IStartNumContext

	// IsFindContext differentiates from other interfaces.
	IsFindContext()
}

type FindContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFindContext() *FindContext {
	var p = new(FindContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_find
	return p
}

func InitEmptyFindContext(p *FindContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_find
}

func (*FindContext) IsFindContext() {}

func NewFindContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FindContext {
	var p = new(FindContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_find

	return p
}

func (s *FindContext) GetParser() antlr.Parser { return s.parser }

func (s *FindContext) FIND() antlr.TerminalNode {
	return s.GetToken(FormulaParserFIND, 0)
}

func (s *FindContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *FindContext) SearchExpression() ISearchExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchExpressionContext)
}

func (s *FindContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FormulaParserCOMMA)
}

func (s *FindContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, i)
}

func (s *FindContext) TextExpression() ITextExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextExpressionContext)
}

func (s *FindContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *FindContext) StartNum() IStartNumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStartNumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStartNumContext)
}

func (s *FindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FindContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterFind(s)
	}
}

func (s *FindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitFind(s)
	}
}

func (s *FindContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitFind(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Find() (localctx IFindContext) {
	localctx = NewFindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, FormulaParserRULE_find)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(323)
		p.Match(FormulaParserFIND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(324)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(325)
		p.SearchExpression()
	}
	{
		p.SetState(326)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(327)
		p.TextExpression()
	}
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FormulaParserCOMMA {
		{
			p.SetState(328)
			p.Match(FormulaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)
			p.StartNum()
		}

	}
	{
		p.SetState(332)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFloorContext is an interface to support dynamic dispatch.
type IFloorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FLOOR() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsFloorContext differentiates from other interfaces.
	IsFloorContext()
}

type FloorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloorContext() *FloorContext {
	var p = new(FloorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_floor
	return p
}

func InitEmptyFloorContext(p *FloorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_floor
}

func (*FloorContext) IsFloorContext() {}

func NewFloorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloorContext {
	var p = new(FloorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_floor

	return p
}

func (s *FloorContext) GetParser() antlr.Parser { return s.parser }

func (s *FloorContext) FLOOR() antlr.TerminalNode {
	return s.GetToken(FormulaParserFLOOR, 0)
}

func (s *FloorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *FloorContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FloorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *FloorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterFloor(s)
	}
}

func (s *FloorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitFloor(s)
	}
}

func (s *FloorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitFloor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Floor() (localctx IFloorContext) {
	localctx = NewFloorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, FormulaParserRULE_floor)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)
		p.Match(FormulaParserFLOOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(335)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(336)
		p.expression(0)
	}
	{
		p.SetState(337)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGeolocationContext is an interface to support dynamic dispatch.
type IGeolocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GEOLOCATION() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	LatitudeExpression() ILatitudeExpressionContext
	COMMA() antlr.TerminalNode
	LongitudeExpression() ILongitudeExpressionContext
	RPAREN() antlr.TerminalNode

	// IsGeolocationContext differentiates from other interfaces.
	IsGeolocationContext()
}

type GeolocationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeolocationContext() *GeolocationContext {
	var p = new(GeolocationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_geolocation
	return p
}

func InitEmptyGeolocationContext(p *GeolocationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_geolocation
}

func (*GeolocationContext) IsGeolocationContext() {}

func NewGeolocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeolocationContext {
	var p = new(GeolocationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_geolocation

	return p
}

func (s *GeolocationContext) GetParser() antlr.Parser { return s.parser }

func (s *GeolocationContext) GEOLOCATION() antlr.TerminalNode {
	return s.GetToken(FormulaParserGEOLOCATION, 0)
}

func (s *GeolocationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *GeolocationContext) LatitudeExpression() ILatitudeExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILatitudeExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILatitudeExpressionContext)
}

func (s *GeolocationContext) COMMA() antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, 0)
}

func (s *GeolocationContext) LongitudeExpression() ILongitudeExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILongitudeExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILongitudeExpressionContext)
}

func (s *GeolocationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *GeolocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeolocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeolocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterGeolocation(s)
	}
}

func (s *GeolocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitGeolocation(s)
	}
}

func (s *GeolocationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitGeolocation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Geolocation() (localctx IGeolocationContext) {
	localctx = NewGeolocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, FormulaParserRULE_geolocation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.Match(FormulaParserGEOLOCATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(340)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(341)
		p.LatitudeExpression()
	}
	{
		p.SetState(342)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(343)
		p.LongitudeExpression()
	}
	{
		p.SetState(344)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGetrecordidsContext is an interface to support dynamic dispatch.
type IGetrecordidsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GETRECORDIDS() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsGetrecordidsContext differentiates from other interfaces.
	IsGetrecordidsContext()
}

type GetrecordidsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGetrecordidsContext() *GetrecordidsContext {
	var p = new(GetrecordidsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_getrecordids
	return p
}

func InitEmptyGetrecordidsContext(p *GetrecordidsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_getrecordids
}

func (*GetrecordidsContext) IsGetrecordidsContext() {}

func NewGetrecordidsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GetrecordidsContext {
	var p = new(GetrecordidsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_getrecordids

	return p
}

func (s *GetrecordidsContext) GetParser() antlr.Parser { return s.parser }

func (s *GetrecordidsContext) GETRECORDIDS() antlr.TerminalNode {
	return s.GetToken(FormulaParserGETRECORDIDS, 0)
}

func (s *GetrecordidsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *GetrecordidsContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GetrecordidsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *GetrecordidsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetrecordidsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GetrecordidsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterGetrecordids(s)
	}
}

func (s *GetrecordidsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitGetrecordids(s)
	}
}

func (s *GetrecordidsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitGetrecordids(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Getrecordids() (localctx IGetrecordidsContext) {
	localctx = NewGetrecordidsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, FormulaParserRULE_getrecordids)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(346)
		p.Match(FormulaParserGETRECORDIDS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(347)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(348)
		p.expression(0)
	}
	{
		p.SetState(349)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGetsessionidContext is an interface to support dynamic dispatch.
type IGetsessionidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GETSESSIONID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsGetsessionidContext differentiates from other interfaces.
	IsGetsessionidContext()
}

type GetsessionidContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGetsessionidContext() *GetsessionidContext {
	var p = new(GetsessionidContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_getsessionid
	return p
}

func InitEmptyGetsessionidContext(p *GetsessionidContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_getsessionid
}

func (*GetsessionidContext) IsGetsessionidContext() {}

func NewGetsessionidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GetsessionidContext {
	var p = new(GetsessionidContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_getsessionid

	return p
}

func (s *GetsessionidContext) GetParser() antlr.Parser { return s.parser }

func (s *GetsessionidContext) GETSESSIONID() antlr.TerminalNode {
	return s.GetToken(FormulaParserGETSESSIONID, 0)
}

func (s *GetsessionidContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *GetsessionidContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *GetsessionidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetsessionidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GetsessionidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterGetsessionid(s)
	}
}

func (s *GetsessionidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitGetsessionid(s)
	}
}

func (s *GetsessionidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitGetsessionid(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Getsessionid() (localctx IGetsessionidContext) {
	localctx = NewGetsessionidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, FormulaParserRULE_getsessionid)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(351)
		p.Match(FormulaParserGETSESSIONID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(352)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(353)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHourContext is an interface to support dynamic dispatch.
type IHourContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HOUR() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsHourContext differentiates from other interfaces.
	IsHourContext()
}

type HourContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHourContext() *HourContext {
	var p = new(HourContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_hour
	return p
}

func InitEmptyHourContext(p *HourContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_hour
}

func (*HourContext) IsHourContext() {}

func NewHourContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HourContext {
	var p = new(HourContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_hour

	return p
}

func (s *HourContext) GetParser() antlr.Parser { return s.parser }

func (s *HourContext) HOUR() antlr.TerminalNode {
	return s.GetToken(FormulaParserHOUR, 0)
}

func (s *HourContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *HourContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *HourContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *HourContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HourContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HourContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterHour(s)
	}
}

func (s *HourContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitHour(s)
	}
}

func (s *HourContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitHour(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Hour() (localctx IHourContext) {
	localctx = NewHourContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, FormulaParserRULE_hour)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.Match(FormulaParserHOUR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(356)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(357)
		p.expression(0)
	}
	{
		p.SetState(358)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHtmlencodeContext is an interface to support dynamic dispatch.
type IHtmlencodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HTMLENCODE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsHtmlencodeContext differentiates from other interfaces.
	IsHtmlencodeContext()
}

type HtmlencodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHtmlencodeContext() *HtmlencodeContext {
	var p = new(HtmlencodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_htmlencode
	return p
}

func InitEmptyHtmlencodeContext(p *HtmlencodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_htmlencode
}

func (*HtmlencodeContext) IsHtmlencodeContext() {}

func NewHtmlencodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HtmlencodeContext {
	var p = new(HtmlencodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_htmlencode

	return p
}

func (s *HtmlencodeContext) GetParser() antlr.Parser { return s.parser }

func (s *HtmlencodeContext) HTMLENCODE() antlr.TerminalNode {
	return s.GetToken(FormulaParserHTMLENCODE, 0)
}

func (s *HtmlencodeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *HtmlencodeContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *HtmlencodeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *HtmlencodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HtmlencodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HtmlencodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterHtmlencode(s)
	}
}

func (s *HtmlencodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitHtmlencode(s)
	}
}

func (s *HtmlencodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitHtmlencode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Htmlencode() (localctx IHtmlencodeContext) {
	localctx = NewHtmlencodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, FormulaParserRULE_htmlencode)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(360)
		p.Match(FormulaParserHTMLENCODE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(361)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(362)
		p.expression(0)
	}
	{
		p.SetState(363)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHyperlinkContext is an interface to support dynamic dispatch.
type IHyperlinkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HYPERLINK() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	UrlExpression() IUrlExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	NameExpression() INameExpressionContext
	RPAREN() antlr.TerminalNode
	TargetExpression() ITargetExpressionContext

	// IsHyperlinkContext differentiates from other interfaces.
	IsHyperlinkContext()
}

type HyperlinkContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHyperlinkContext() *HyperlinkContext {
	var p = new(HyperlinkContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_hyperlink
	return p
}

func InitEmptyHyperlinkContext(p *HyperlinkContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_hyperlink
}

func (*HyperlinkContext) IsHyperlinkContext() {}

func NewHyperlinkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HyperlinkContext {
	var p = new(HyperlinkContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_hyperlink

	return p
}

func (s *HyperlinkContext) GetParser() antlr.Parser { return s.parser }

func (s *HyperlinkContext) HYPERLINK() antlr.TerminalNode {
	return s.GetToken(FormulaParserHYPERLINK, 0)
}

func (s *HyperlinkContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *HyperlinkContext) UrlExpression() IUrlExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUrlExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUrlExpressionContext)
}

func (s *HyperlinkContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FormulaParserCOMMA)
}

func (s *HyperlinkContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, i)
}

func (s *HyperlinkContext) NameExpression() INameExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameExpressionContext)
}

func (s *HyperlinkContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *HyperlinkContext) TargetExpression() ITargetExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetExpressionContext)
}

func (s *HyperlinkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HyperlinkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HyperlinkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterHyperlink(s)
	}
}

func (s *HyperlinkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitHyperlink(s)
	}
}

func (s *HyperlinkContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitHyperlink(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Hyperlink() (localctx IHyperlinkContext) {
	localctx = NewHyperlinkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, FormulaParserRULE_hyperlink)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		p.Match(FormulaParserHYPERLINK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(366)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(367)
		p.UrlExpression()
	}
	{
		p.SetState(368)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(369)
		p.NameExpression()
	}
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FormulaParserCOMMA {
		{
			p.SetState(370)
			p.Match(FormulaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(371)
			p.TargetExpression()
		}

	}
	{
		p.SetState(374)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfContext is an interface to support dynamic dispatch.
type IIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	LogicalExpression() ILogicalExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	IfTrueExpression() IIfTrueExpressionContext
	IfFalseExpression() IIfFalseExpressionContext
	RPAREN() antlr.TerminalNode

	// IsIfContext differentiates from other interfaces.
	IsIfContext()
}

type IfContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfContext() *IfContext {
	var p = new(IfContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_if
	return p
}

func InitEmptyIfContext(p *IfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_if
}

func (*IfContext) IsIfContext() {}

func NewIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfContext {
	var p = new(IfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_if

	return p
}

func (s *IfContext) GetParser() antlr.Parser { return s.parser }

func (s *IfContext) IF() antlr.TerminalNode {
	return s.GetToken(FormulaParserIF, 0)
}

func (s *IfContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *IfContext) LogicalExpression() ILogicalExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalExpressionContext)
}

func (s *IfContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FormulaParserCOMMA)
}

func (s *IfContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, i)
}

func (s *IfContext) IfTrueExpression() IIfTrueExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfTrueExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfTrueExpressionContext)
}

func (s *IfContext) IfFalseExpression() IIfFalseExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfFalseExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfFalseExpressionContext)
}

func (s *IfContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterIf(s)
	}
}

func (s *IfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitIf(s)
	}
}

func (s *IfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitIf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) If_() (localctx IIfContext) {
	localctx = NewIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, FormulaParserRULE_if)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.Match(FormulaParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(377)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(378)
		p.LogicalExpression()
	}
	{
		p.SetState(379)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(380)
		p.IfTrueExpression()
	}
	{
		p.SetState(381)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(382)
		p.IfFalseExpression()
	}
	{
		p.SetState(383)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImageContext is an interface to support dynamic dispatch.
type IImageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMAGE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	UrlExpression() IUrlExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	TextExpression() ITextExpressionContext
	RPAREN() antlr.TerminalNode
	HeightExpression() IHeightExpressionContext
	WidthExpression() IWidthExpressionContext

	// IsImageContext differentiates from other interfaces.
	IsImageContext()
}

type ImageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImageContext() *ImageContext {
	var p = new(ImageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_image
	return p
}

func InitEmptyImageContext(p *ImageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_image
}

func (*ImageContext) IsImageContext() {}

func NewImageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImageContext {
	var p = new(ImageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_image

	return p
}

func (s *ImageContext) GetParser() antlr.Parser { return s.parser }

func (s *ImageContext) IMAGE() antlr.TerminalNode {
	return s.GetToken(FormulaParserIMAGE, 0)
}

func (s *ImageContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *ImageContext) UrlExpression() IUrlExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUrlExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUrlExpressionContext)
}

func (s *ImageContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FormulaParserCOMMA)
}

func (s *ImageContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, i)
}

func (s *ImageContext) TextExpression() ITextExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextExpressionContext)
}

func (s *ImageContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *ImageContext) HeightExpression() IHeightExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHeightExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHeightExpressionContext)
}

func (s *ImageContext) WidthExpression() IWidthExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWidthExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWidthExpressionContext)
}

func (s *ImageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterImage(s)
	}
}

func (s *ImageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitImage(s)
	}
}

func (s *ImageContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitImage(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Image() (localctx IImageContext) {
	localctx = NewImageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, FormulaParserRULE_image)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(385)
		p.Match(FormulaParserIMAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(386)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(387)
		p.UrlExpression()
	}
	{
		p.SetState(388)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(389)
		p.TextExpression()
	}
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FormulaParserCOMMA {
		{
			p.SetState(390)
			p.Match(FormulaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(391)
			p.HeightExpression()
		}
		{
			p.SetState(392)
			p.Match(FormulaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(393)
			p.WidthExpression()
		}

	}
	{
		p.SetState(397)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImageproxyurlContext is an interface to support dynamic dispatch.
type IImageproxyurlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMAGEPROXYURL() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	UrlExpression() IUrlExpressionContext
	RPAREN() antlr.TerminalNode

	// IsImageproxyurlContext differentiates from other interfaces.
	IsImageproxyurlContext()
}

type ImageproxyurlContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImageproxyurlContext() *ImageproxyurlContext {
	var p = new(ImageproxyurlContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_imageproxyurl
	return p
}

func InitEmptyImageproxyurlContext(p *ImageproxyurlContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_imageproxyurl
}

func (*ImageproxyurlContext) IsImageproxyurlContext() {}

func NewImageproxyurlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImageproxyurlContext {
	var p = new(ImageproxyurlContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_imageproxyurl

	return p
}

func (s *ImageproxyurlContext) GetParser() antlr.Parser { return s.parser }

func (s *ImageproxyurlContext) IMAGEPROXYURL() antlr.TerminalNode {
	return s.GetToken(FormulaParserIMAGEPROXYURL, 0)
}

func (s *ImageproxyurlContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *ImageproxyurlContext) UrlExpression() IUrlExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUrlExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUrlExpressionContext)
}

func (s *ImageproxyurlContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *ImageproxyurlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImageproxyurlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImageproxyurlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterImageproxyurl(s)
	}
}

func (s *ImageproxyurlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitImageproxyurl(s)
	}
}

func (s *ImageproxyurlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitImageproxyurl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Imageproxyurl() (localctx IImageproxyurlContext) {
	localctx = NewImageproxyurlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, FormulaParserRULE_imageproxyurl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(399)
		p.Match(FormulaParserIMAGEPROXYURL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(400)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(401)
		p.UrlExpression()
	}
	{
		p.SetState(402)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMidContext is an interface to support dynamic dispatch.
type IMidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	TextExpression() ITextExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	StartNum() IStartNumContext
	NumChars() INumCharsContext
	RPAREN() antlr.TerminalNode

	// IsMidContext differentiates from other interfaces.
	IsMidContext()
}

type MidContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMidContext() *MidContext {
	var p = new(MidContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_mid
	return p
}

func InitEmptyMidContext(p *MidContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_mid
}

func (*MidContext) IsMidContext() {}

func NewMidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MidContext {
	var p = new(MidContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_mid

	return p
}

func (s *MidContext) GetParser() antlr.Parser { return s.parser }

func (s *MidContext) MID() antlr.TerminalNode {
	return s.GetToken(FormulaParserMID, 0)
}

func (s *MidContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *MidContext) TextExpression() ITextExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextExpressionContext)
}

func (s *MidContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FormulaParserCOMMA)
}

func (s *MidContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, i)
}

func (s *MidContext) StartNum() IStartNumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStartNumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStartNumContext)
}

func (s *MidContext) NumChars() INumCharsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumCharsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumCharsContext)
}

func (s *MidContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *MidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterMid(s)
	}
}

func (s *MidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitMid(s)
	}
}

func (s *MidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitMid(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Mid() (localctx IMidContext) {
	localctx = NewMidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, FormulaParserRULE_mid)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(404)
		p.Match(FormulaParserMID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(405)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(406)
		p.TextExpression()
	}
	{
		p.SetState(407)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(408)
		p.StartNum()
	}
	{
		p.SetState(409)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(410)
		p.NumChars()
	}
	{
		p.SetState(411)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIncludesContext is an interface to support dynamic dispatch.
type IIncludesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INCLUDES() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	FieldExpression() IFieldExpressionContext
	COMMA() antlr.TerminalNode
	ValueExpression() IValueExpressionContext
	RPAREN() antlr.TerminalNode

	// IsIncludesContext differentiates from other interfaces.
	IsIncludesContext()
}

type IncludesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncludesContext() *IncludesContext {
	var p = new(IncludesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_includes
	return p
}

func InitEmptyIncludesContext(p *IncludesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_includes
}

func (*IncludesContext) IsIncludesContext() {}

func NewIncludesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncludesContext {
	var p = new(IncludesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_includes

	return p
}

func (s *IncludesContext) GetParser() antlr.Parser { return s.parser }

func (s *IncludesContext) INCLUDES() antlr.TerminalNode {
	return s.GetToken(FormulaParserINCLUDES, 0)
}

func (s *IncludesContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *IncludesContext) FieldExpression() IFieldExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldExpressionContext)
}

func (s *IncludesContext) COMMA() antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, 0)
}

func (s *IncludesContext) ValueExpression() IValueExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExpressionContext)
}

func (s *IncludesContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *IncludesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncludesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncludesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterIncludes(s)
	}
}

func (s *IncludesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitIncludes(s)
	}
}

func (s *IncludesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitIncludes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Includes() (localctx IIncludesContext) {
	localctx = NewIncludesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, FormulaParserRULE_includes)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(413)
		p.Match(FormulaParserINCLUDES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(414)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(415)
		p.FieldExpression()
	}
	{
		p.SetState(416)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(417)
		p.ValueExpression()
	}
	{
		p.SetState(418)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITextContext is an interface to support dynamic dispatch.
type ITextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsTextContext differentiates from other interfaces.
	IsTextContext()
}

type TextContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextContext() *TextContext {
	var p = new(TextContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_text
	return p
}

func InitEmptyTextContext(p *TextContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_text
}

func (*TextContext) IsTextContext() {}

func NewTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextContext {
	var p = new(TextContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_text

	return p
}

func (s *TextContext) GetParser() antlr.Parser { return s.parser }

func (s *TextContext) TEXT() antlr.TerminalNode {
	return s.GetToken(FormulaParserTEXT, 0)
}

func (s *TextContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *TextContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TextContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *TextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterText(s)
	}
}

func (s *TextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitText(s)
	}
}

func (s *TextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitText(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Text() (localctx ITextContext) {
	localctx = NewTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, FormulaParserRULE_text)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(420)
		p.Match(FormulaParserTEXT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(421)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(422)
		p.expression(0)
	}
	{
		p.SetState(423)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VALUE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) VALUE() antlr.TerminalNode {
	return s.GetToken(FormulaParserVALUE, 0)
}

func (s *ValueContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *ValueContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ValueContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, FormulaParserRULE_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(425)
		p.Match(FormulaParserVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(426)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(427)
		p.expression(0)
	}
	{
		p.SetState(428)
		p.Match(FormulaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldExpressionContext is an interface to support dynamic dispatch.
type IFieldExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsFieldExpressionContext differentiates from other interfaces.
	IsFieldExpressionContext()
}

type FieldExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldExpressionContext() *FieldExpressionContext {
	var p = new(FieldExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_fieldExpression
	return p
}

func InitEmptyFieldExpressionContext(p *FieldExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_fieldExpression
}

func (*FieldExpressionContext) IsFieldExpressionContext() {}

func NewFieldExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldExpressionContext {
	var p = new(FieldExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_fieldExpression

	return p
}

func (s *FieldExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FieldExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterFieldExpression(s)
	}
}

func (s *FieldExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitFieldExpression(s)
	}
}

func (s *FieldExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitFieldExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) FieldExpression() (localctx IFieldExpressionContext) {
	localctx = NewFieldExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, FormulaParserRULE_fieldExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(430)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueExpressionContext is an interface to support dynamic dispatch.
type IValueExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsValueExpressionContext differentiates from other interfaces.
	IsValueExpressionContext()
}

type ValueExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueExpressionContext() *ValueExpressionContext {
	var p = new(ValueExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_valueExpression
	return p
}

func InitEmptyValueExpressionContext(p *ValueExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_valueExpression
}

func (*ValueExpressionContext) IsValueExpressionContext() {}

func NewValueExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueExpressionContext {
	var p = new(ValueExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_valueExpression

	return p
}

func (s *ValueExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ValueExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterValueExpression(s)
	}
}

func (s *ValueExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitValueExpression(s)
	}
}

func (s *ValueExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitValueExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) ValueExpression() (localctx IValueExpressionContext) {
	localctx = NewValueExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, FormulaParserRULE_valueExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(432)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IResultExpressionContext is an interface to support dynamic dispatch.
type IResultExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsResultExpressionContext differentiates from other interfaces.
	IsResultExpressionContext()
}

type ResultExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResultExpressionContext() *ResultExpressionContext {
	var p = new(ResultExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_resultExpression
	return p
}

func InitEmptyResultExpressionContext(p *ResultExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_resultExpression
}

func (*ResultExpressionContext) IsResultExpressionContext() {}

func NewResultExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResultExpressionContext {
	var p = new(ResultExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_resultExpression

	return p
}

func (s *ResultExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ResultExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ResultExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResultExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResultExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterResultExpression(s)
	}
}

func (s *ResultExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitResultExpression(s)
	}
}

func (s *ResultExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitResultExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) ResultExpression() (localctx IResultExpressionContext) {
	localctx = NewResultExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, FormulaParserRULE_resultExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(434)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefaultExpressionContext is an interface to support dynamic dispatch.
type IDefaultExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsDefaultExpressionContext differentiates from other interfaces.
	IsDefaultExpressionContext()
}

type DefaultExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultExpressionContext() *DefaultExpressionContext {
	var p = new(DefaultExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_defaultExpression
	return p
}

func InitEmptyDefaultExpressionContext(p *DefaultExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_defaultExpression
}

func (*DefaultExpressionContext) IsDefaultExpressionContext() {}

func NewDefaultExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultExpressionContext {
	var p = new(DefaultExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_defaultExpression

	return p
}

func (s *DefaultExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DefaultExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterDefaultExpression(s)
	}
}

func (s *DefaultExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitDefaultExpression(s)
	}
}

func (s *DefaultExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitDefaultExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) DefaultExpression() (localctx IDefaultExpressionContext) {
	localctx = NewDefaultExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, FormulaParserRULE_defaultExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(436)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IYearExpressionContext is an interface to support dynamic dispatch.
type IYearExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsYearExpressionContext differentiates from other interfaces.
	IsYearExpressionContext()
}

type YearExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyYearExpressionContext() *YearExpressionContext {
	var p = new(YearExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_yearExpression
	return p
}

func InitEmptyYearExpressionContext(p *YearExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_yearExpression
}

func (*YearExpressionContext) IsYearExpressionContext() {}

func NewYearExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *YearExpressionContext {
	var p = new(YearExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_yearExpression

	return p
}

func (s *YearExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *YearExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *YearExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *YearExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *YearExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterYearExpression(s)
	}
}

func (s *YearExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitYearExpression(s)
	}
}

func (s *YearExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitYearExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) YearExpression() (localctx IYearExpressionContext) {
	localctx = NewYearExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, FormulaParserRULE_yearExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(438)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMonthExpressionContext is an interface to support dynamic dispatch.
type IMonthExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsMonthExpressionContext differentiates from other interfaces.
	IsMonthExpressionContext()
}

type MonthExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMonthExpressionContext() *MonthExpressionContext {
	var p = new(MonthExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_monthExpression
	return p
}

func InitEmptyMonthExpressionContext(p *MonthExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_monthExpression
}

func (*MonthExpressionContext) IsMonthExpressionContext() {}

func NewMonthExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MonthExpressionContext {
	var p = new(MonthExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_monthExpression

	return p
}

func (s *MonthExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MonthExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MonthExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonthExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MonthExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterMonthExpression(s)
	}
}

func (s *MonthExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitMonthExpression(s)
	}
}

func (s *MonthExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitMonthExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) MonthExpression() (localctx IMonthExpressionContext) {
	localctx = NewMonthExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, FormulaParserRULE_monthExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(440)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDayExpressionContext is an interface to support dynamic dispatch.
type IDayExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsDayExpressionContext differentiates from other interfaces.
	IsDayExpressionContext()
}

type DayExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDayExpressionContext() *DayExpressionContext {
	var p = new(DayExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_dayExpression
	return p
}

func InitEmptyDayExpressionContext(p *DayExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_dayExpression
}

func (*DayExpressionContext) IsDayExpressionContext() {}

func NewDayExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DayExpressionContext {
	var p = new(DayExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_dayExpression

	return p
}

func (s *DayExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *DayExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DayExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DayExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DayExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterDayExpression(s)
	}
}

func (s *DayExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitDayExpression(s)
	}
}

func (s *DayExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitDayExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) DayExpression() (localctx IDayExpressionContext) {
	localctx = NewDayExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, FormulaParserRULE_dayExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(442)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnitExpressionContext is an interface to support dynamic dispatch.
type IUnitExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsUnitExpressionContext differentiates from other interfaces.
	IsUnitExpressionContext()
}

type UnitExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnitExpressionContext() *UnitExpressionContext {
	var p = new(UnitExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_unitExpression
	return p
}

func InitEmptyUnitExpressionContext(p *UnitExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_unitExpression
}

func (*UnitExpressionContext) IsUnitExpressionContext() {}

func NewUnitExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitExpressionContext {
	var p = new(UnitExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_unitExpression

	return p
}

func (s *UnitExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnitExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnitExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterUnitExpression(s)
	}
}

func (s *UnitExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitUnitExpression(s)
	}
}

func (s *UnitExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitUnitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) UnitExpression() (localctx IUnitExpressionContext) {
	localctx = NewUnitExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, FormulaParserRULE_unitExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(444)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISearchExpressionContext is an interface to support dynamic dispatch.
type ISearchExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsSearchExpressionContext differentiates from other interfaces.
	IsSearchExpressionContext()
}

type SearchExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchExpressionContext() *SearchExpressionContext {
	var p = new(SearchExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_searchExpression
	return p
}

func InitEmptySearchExpressionContext(p *SearchExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_searchExpression
}

func (*SearchExpressionContext) IsSearchExpressionContext() {}

func NewSearchExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchExpressionContext {
	var p = new(SearchExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_searchExpression

	return p
}

func (s *SearchExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SearchExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SearchExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterSearchExpression(s)
	}
}

func (s *SearchExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitSearchExpression(s)
	}
}

func (s *SearchExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitSearchExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) SearchExpression() (localctx ISearchExpressionContext) {
	localctx = NewSearchExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, FormulaParserRULE_searchExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(446)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITextExpressionContext is an interface to support dynamic dispatch.
type ITextExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsTextExpressionContext differentiates from other interfaces.
	IsTextExpressionContext()
}

type TextExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextExpressionContext() *TextExpressionContext {
	var p = new(TextExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_textExpression
	return p
}

func InitEmptyTextExpressionContext(p *TextExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_textExpression
}

func (*TextExpressionContext) IsTextExpressionContext() {}

func NewTextExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextExpressionContext {
	var p = new(TextExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_textExpression

	return p
}

func (s *TextExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *TextExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TextExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterTextExpression(s)
	}
}

func (s *TextExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitTextExpression(s)
	}
}

func (s *TextExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitTextExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) TextExpression() (localctx ITextExpressionContext) {
	localctx = NewTextExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, FormulaParserRULE_textExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(448)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStartNumContext is an interface to support dynamic dispatch.
type IStartNumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsStartNumContext differentiates from other interfaces.
	IsStartNumContext()
}

type StartNumContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartNumContext() *StartNumContext {
	var p = new(StartNumContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_startNum
	return p
}

func InitEmptyStartNumContext(p *StartNumContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_startNum
}

func (*StartNumContext) IsStartNumContext() {}

func NewStartNumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartNumContext {
	var p = new(StartNumContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_startNum

	return p
}

func (s *StartNumContext) GetParser() antlr.Parser { return s.parser }

func (s *StartNumContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StartNumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartNumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartNumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterStartNum(s)
	}
}

func (s *StartNumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitStartNum(s)
	}
}

func (s *StartNumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitStartNum(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) StartNum() (localctx IStartNumContext) {
	localctx = NewStartNumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, FormulaParserRULE_startNum)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(450)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILatitudeExpressionContext is an interface to support dynamic dispatch.
type ILatitudeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsLatitudeExpressionContext differentiates from other interfaces.
	IsLatitudeExpressionContext()
}

type LatitudeExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLatitudeExpressionContext() *LatitudeExpressionContext {
	var p = new(LatitudeExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_latitudeExpression
	return p
}

func InitEmptyLatitudeExpressionContext(p *LatitudeExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_latitudeExpression
}

func (*LatitudeExpressionContext) IsLatitudeExpressionContext() {}

func NewLatitudeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LatitudeExpressionContext {
	var p = new(LatitudeExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_latitudeExpression

	return p
}

func (s *LatitudeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LatitudeExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LatitudeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LatitudeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LatitudeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterLatitudeExpression(s)
	}
}

func (s *LatitudeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitLatitudeExpression(s)
	}
}

func (s *LatitudeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitLatitudeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) LatitudeExpression() (localctx ILatitudeExpressionContext) {
	localctx = NewLatitudeExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, FormulaParserRULE_latitudeExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(452)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILongitudeExpressionContext is an interface to support dynamic dispatch.
type ILongitudeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsLongitudeExpressionContext differentiates from other interfaces.
	IsLongitudeExpressionContext()
}

type LongitudeExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLongitudeExpressionContext() *LongitudeExpressionContext {
	var p = new(LongitudeExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_longitudeExpression
	return p
}

func InitEmptyLongitudeExpressionContext(p *LongitudeExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_longitudeExpression
}

func (*LongitudeExpressionContext) IsLongitudeExpressionContext() {}

func NewLongitudeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LongitudeExpressionContext {
	var p = new(LongitudeExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_longitudeExpression

	return p
}

func (s *LongitudeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LongitudeExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LongitudeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LongitudeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LongitudeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterLongitudeExpression(s)
	}
}

func (s *LongitudeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitLongitudeExpression(s)
	}
}

func (s *LongitudeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitLongitudeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) LongitudeExpression() (localctx ILongitudeExpressionContext) {
	localctx = NewLongitudeExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, FormulaParserRULE_longitudeExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(454)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUrlExpressionContext is an interface to support dynamic dispatch.
type IUrlExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsUrlExpressionContext differentiates from other interfaces.
	IsUrlExpressionContext()
}

type UrlExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUrlExpressionContext() *UrlExpressionContext {
	var p = new(UrlExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_urlExpression
	return p
}

func InitEmptyUrlExpressionContext(p *UrlExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_urlExpression
}

func (*UrlExpressionContext) IsUrlExpressionContext() {}

func NewUrlExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UrlExpressionContext {
	var p = new(UrlExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_urlExpression

	return p
}

func (s *UrlExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UrlExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UrlExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UrlExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UrlExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterUrlExpression(s)
	}
}

func (s *UrlExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitUrlExpression(s)
	}
}

func (s *UrlExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitUrlExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) UrlExpression() (localctx IUrlExpressionContext) {
	localctx = NewUrlExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, FormulaParserRULE_urlExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(456)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INameExpressionContext is an interface to support dynamic dispatch.
type INameExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsNameExpressionContext differentiates from other interfaces.
	IsNameExpressionContext()
}

type NameExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameExpressionContext() *NameExpressionContext {
	var p = new(NameExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_nameExpression
	return p
}

func InitEmptyNameExpressionContext(p *NameExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_nameExpression
}

func (*NameExpressionContext) IsNameExpressionContext() {}

func NewNameExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameExpressionContext {
	var p = new(NameExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_nameExpression

	return p
}

func (s *NameExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *NameExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NameExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterNameExpression(s)
	}
}

func (s *NameExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitNameExpression(s)
	}
}

func (s *NameExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitNameExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) NameExpression() (localctx INameExpressionContext) {
	localctx = NewNameExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, FormulaParserRULE_nameExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(458)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITargetExpressionContext is an interface to support dynamic dispatch.
type ITargetExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsTargetExpressionContext differentiates from other interfaces.
	IsTargetExpressionContext()
}

type TargetExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetExpressionContext() *TargetExpressionContext {
	var p = new(TargetExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_targetExpression
	return p
}

func InitEmptyTargetExpressionContext(p *TargetExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_targetExpression
}

func (*TargetExpressionContext) IsTargetExpressionContext() {}

func NewTargetExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetExpressionContext {
	var p = new(TargetExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_targetExpression

	return p
}

func (s *TargetExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TargetExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterTargetExpression(s)
	}
}

func (s *TargetExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitTargetExpression(s)
	}
}

func (s *TargetExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitTargetExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) TargetExpression() (localctx ITargetExpressionContext) {
	localctx = NewTargetExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, FormulaParserRULE_targetExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(460)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogicalExpressionContext is an interface to support dynamic dispatch.
type ILogicalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsLogicalExpressionContext differentiates from other interfaces.
	IsLogicalExpressionContext()
}

type LogicalExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalExpressionContext() *LogicalExpressionContext {
	var p = new(LogicalExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_logicalExpression
	return p
}

func InitEmptyLogicalExpressionContext(p *LogicalExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_logicalExpression
}

func (*LogicalExpressionContext) IsLogicalExpressionContext() {}

func NewLogicalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalExpressionContext {
	var p = new(LogicalExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_logicalExpression

	return p
}

func (s *LogicalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogicalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterLogicalExpression(s)
	}
}

func (s *LogicalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitLogicalExpression(s)
	}
}

func (s *LogicalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitLogicalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) LogicalExpression() (localctx ILogicalExpressionContext) {
	localctx = NewLogicalExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, FormulaParserRULE_logicalExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(462)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfTrueExpressionContext is an interface to support dynamic dispatch.
type IIfTrueExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsIfTrueExpressionContext differentiates from other interfaces.
	IsIfTrueExpressionContext()
}

type IfTrueExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfTrueExpressionContext() *IfTrueExpressionContext {
	var p = new(IfTrueExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_ifTrueExpression
	return p
}

func InitEmptyIfTrueExpressionContext(p *IfTrueExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_ifTrueExpression
}

func (*IfTrueExpressionContext) IsIfTrueExpressionContext() {}

func NewIfTrueExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfTrueExpressionContext {
	var p = new(IfTrueExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_ifTrueExpression

	return p
}

func (s *IfTrueExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *IfTrueExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfTrueExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfTrueExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfTrueExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterIfTrueExpression(s)
	}
}

func (s *IfTrueExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitIfTrueExpression(s)
	}
}

func (s *IfTrueExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitIfTrueExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) IfTrueExpression() (localctx IIfTrueExpressionContext) {
	localctx = NewIfTrueExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, FormulaParserRULE_ifTrueExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(464)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfFalseExpressionContext is an interface to support dynamic dispatch.
type IIfFalseExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsIfFalseExpressionContext differentiates from other interfaces.
	IsIfFalseExpressionContext()
}

type IfFalseExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfFalseExpressionContext() *IfFalseExpressionContext {
	var p = new(IfFalseExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_ifFalseExpression
	return p
}

func InitEmptyIfFalseExpressionContext(p *IfFalseExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_ifFalseExpression
}

func (*IfFalseExpressionContext) IsIfFalseExpressionContext() {}

func NewIfFalseExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfFalseExpressionContext {
	var p = new(IfFalseExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_ifFalseExpression

	return p
}

func (s *IfFalseExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *IfFalseExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfFalseExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfFalseExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfFalseExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterIfFalseExpression(s)
	}
}

func (s *IfFalseExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitIfFalseExpression(s)
	}
}

func (s *IfFalseExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitIfFalseExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) IfFalseExpression() (localctx IIfFalseExpressionContext) {
	localctx = NewIfFalseExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, FormulaParserRULE_ifFalseExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(466)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHeightExpressionContext is an interface to support dynamic dispatch.
type IHeightExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsHeightExpressionContext differentiates from other interfaces.
	IsHeightExpressionContext()
}

type HeightExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeightExpressionContext() *HeightExpressionContext {
	var p = new(HeightExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_heightExpression
	return p
}

func InitEmptyHeightExpressionContext(p *HeightExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_heightExpression
}

func (*HeightExpressionContext) IsHeightExpressionContext() {}

func NewHeightExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeightExpressionContext {
	var p = new(HeightExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_heightExpression

	return p
}

func (s *HeightExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *HeightExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *HeightExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeightExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeightExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterHeightExpression(s)
	}
}

func (s *HeightExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitHeightExpression(s)
	}
}

func (s *HeightExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitHeightExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) HeightExpression() (localctx IHeightExpressionContext) {
	localctx = NewHeightExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, FormulaParserRULE_heightExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(468)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWidthExpressionContext is an interface to support dynamic dispatch.
type IWidthExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsWidthExpressionContext differentiates from other interfaces.
	IsWidthExpressionContext()
}

type WidthExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWidthExpressionContext() *WidthExpressionContext {
	var p = new(WidthExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_widthExpression
	return p
}

func InitEmptyWidthExpressionContext(p *WidthExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_widthExpression
}

func (*WidthExpressionContext) IsWidthExpressionContext() {}

func NewWidthExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WidthExpressionContext {
	var p = new(WidthExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_widthExpression

	return p
}

func (s *WidthExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *WidthExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WidthExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WidthExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WidthExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterWidthExpression(s)
	}
}

func (s *WidthExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitWidthExpression(s)
	}
}

func (s *WidthExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitWidthExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) WidthExpression() (localctx IWidthExpressionContext) {
	localctx = NewWidthExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, FormulaParserRULE_widthExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(470)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_start
	return p
}

func InitEmptyStartContext(p *StartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_start
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Start_() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, FormulaParserRULE_start)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(472)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumCharsContext is an interface to support dynamic dispatch.
type INumCharsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsNumCharsContext differentiates from other interfaces.
	IsNumCharsContext()
}

type NumCharsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumCharsContext() *NumCharsContext {
	var p = new(NumCharsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_numChars
	return p
}

func InitEmptyNumCharsContext(p *NumCharsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_numChars
}

func (*NumCharsContext) IsNumCharsContext() {}

func NewNumCharsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumCharsContext {
	var p = new(NumCharsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_numChars

	return p
}

func (s *NumCharsContext) GetParser() antlr.Parser { return s.parser }

func (s *NumCharsContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NumCharsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumCharsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumCharsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterNumChars(s)
	}
}

func (s *NumCharsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitNumChars(s)
	}
}

func (s *NumCharsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitNumChars(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) NumChars() (localctx INumCharsContext) {
	localctx = NewNumCharsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, FormulaParserRULE_numChars)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(474)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal() ILiteralContext
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PrimaryContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *PrimaryContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, FormulaParserRULE_primary)
	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FormulaParserStringLiteral, FormulaParserIntegerLiteral, FormulaParserFloatingPointLiteral, FormulaParserBooleanLiteral, FormulaParserNullLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(476)
			p.Literal()
		}

	case FormulaParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(477)
			p.Match(FormulaParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(478)
			p.expression(0)
		}
		{
			p.SetState(479)
			p.Match(FormulaParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FormulaParserCOMMA)
}

func (s *ArgumentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, i)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, FormulaParserRULE_arguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(483)
		p.expression(0)
	}
	p.SetState(488)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FormulaParserCOMMA {
		{
			p.SetState(484)
			p.Match(FormulaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(485)
			p.expression(0)
		}

		p.SetState(490)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StringLiteral() antlr.TerminalNode
	IntegerLiteral() antlr.TerminalNode
	FloatingPointLiteral() antlr.TerminalNode
	BooleanLiteral() antlr.TerminalNode
	NullLiteral() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(FormulaParserStringLiteral, 0)
}

func (s *LiteralContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(FormulaParserIntegerLiteral, 0)
}

func (s *LiteralContext) FloatingPointLiteral() antlr.TerminalNode {
	return s.GetToken(FormulaParserFloatingPointLiteral, 0)
}

func (s *LiteralContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(FormulaParserBooleanLiteral, 0)
}

func (s *LiteralContext) NullLiteral() antlr.TerminalNode {
	return s.GetToken(FormulaParserNullLiteral, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, FormulaParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(491)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-86)) & ^0x3f) == 0 && ((int64(1)<<(_la-86))&31) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *FormulaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *FormulaParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
