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
		"", "", "", "", "", "", "", "'('", "')'", "'.'", "','", "'^'", "'!'",
		"'&'", "'+'", "'-'", "'*'", "'/'", "'<'", "'>'", "'='", "'<>'", "'!='",
		"'=='", "'&&'", "'||'",
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
		"COMMA", "CARET", "BANG", "BITAND", "ADD", "SUB", "MUL", "DIV", "LT",
		"GT", "ASSIGN", "LESSANDGREATER", "NOTEQUAL", "EQUAL", "AND", "OR",
		"WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"compilationUnit", "expression", "functionCall", "abs", "addMonths",
		"and", "begins", "blankvalue", "br", "case", "casesafeid", "ceiling",
		"contains", "currencyrate", "date", "datevalue", "datetimevalue", "day",
		"distance", "exp", "find", "floor", "geolocation", "getrecordids", "getsessionid",
		"hour", "htmlencode", "hyperlink", "if", "image", "imageproxyurl", "includes",
		"isblank", "ischanged", "isnew", "isnull", "isnumber", "ispickval",
		"left", "len", "lower", "lpad", "max", "mid", "min", "mod", "month",
		"not", "now", "or", "priorvalue", "regex", "right", "round", "substitute",
		"trim", "text", "today", "upper", "value", "year", "fieldExpression",
		"dateExpression", "valueExpression", "resultExpression", "substituteValue",
		"defaultExpression", "yearExpression", "monthExpression", "dayExpression",
		"regexExpression", "unitExpression", "searchExpression", "textExpression",
		"oldText", "replacement", "startNum", "compareText", "num", "divisor",
		"digits", "latitudeExpression", "longitudeExpression", "url", "name",
		"target", "logicalExpression", "ifTrueExpression", "ifFalseExpression",
		"length", "padString", "heightExpression", "widthExpression", "start",
		"numChars", "primary", "literal",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 112, 774, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
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
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 62, 2,
		63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67, 2, 68,
		7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2, 73, 7,
		73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 2, 77, 7, 77, 2, 78, 7, 78,
		2, 79, 7, 79, 2, 80, 7, 80, 2, 81, 7, 81, 2, 82, 7, 82, 2, 83, 7, 83, 2,
		84, 7, 84, 2, 85, 7, 85, 2, 86, 7, 86, 2, 87, 7, 87, 2, 88, 7, 88, 2, 89,
		7, 89, 2, 90, 7, 90, 2, 91, 7, 91, 2, 92, 7, 92, 2, 93, 7, 93, 2, 94, 7,
		94, 2, 95, 7, 95, 2, 96, 7, 96, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 5, 1, 205, 8, 1, 10, 1, 12, 1, 208, 9, 1, 3, 1, 210,
		8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 3, 1, 224, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1,
		233, 8, 1, 10, 1, 12, 1, 236, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 296, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		5, 5, 315, 8, 5, 10, 5, 12, 5, 318, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 4, 9,
		348, 8, 9, 11, 9, 12, 9, 349, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 3, 20, 423, 8, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 3, 27, 465, 8, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 488, 8, 29, 1, 29, 1, 29,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41,
		1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 559, 8, 41, 1, 41, 1, 41, 1, 42, 1,
		42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 5, 42, 570, 8, 42, 10, 42, 12, 42,
		573, 9, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1,
		43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 5, 44,
		593, 8, 44, 10, 44, 12, 44, 596, 9, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1,
		45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47,
		1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1,
		49, 1, 49, 1, 49, 5, 49, 626, 8, 49, 10, 49, 12, 49, 629, 9, 49, 1, 49,
		1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1,
		51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 53,
		1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 54, 1,
		54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 56,
		1, 56, 1, 56, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 1, 57, 1, 58, 1, 58, 1,
		58, 1, 58, 1, 58, 1, 59, 1, 59, 1, 59, 1, 59, 1, 59, 1, 60, 1, 60, 1, 60,
		1, 60, 1, 60, 1, 61, 1, 61, 1, 62, 1, 62, 1, 63, 1, 63, 1, 64, 1, 64, 1,
		65, 1, 65, 1, 66, 1, 66, 1, 67, 1, 67, 1, 68, 1, 68, 1, 69, 1, 69, 1, 70,
		1, 70, 1, 71, 1, 71, 1, 72, 1, 72, 1, 73, 1, 73, 1, 74, 1, 74, 1, 75, 1,
		75, 1, 76, 1, 76, 1, 77, 1, 77, 1, 78, 1, 78, 1, 79, 1, 79, 1, 80, 1, 80,
		1, 81, 1, 81, 1, 82, 1, 82, 1, 83, 1, 83, 1, 84, 1, 84, 1, 85, 1, 85, 1,
		86, 1, 86, 1, 87, 1, 87, 1, 88, 1, 88, 1, 89, 1, 89, 1, 90, 1, 90, 1, 91,
		1, 91, 1, 92, 1, 92, 1, 93, 1, 93, 1, 94, 1, 94, 1, 95, 1, 95, 1, 95, 1,
		95, 1, 95, 3, 95, 770, 8, 95, 1, 96, 1, 96, 1, 96, 0, 1, 2, 97, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
		80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112,
		114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 136, 138, 140, 142,
		144, 146, 148, 150, 152, 154, 156, 158, 160, 162, 164, 166, 168, 170, 172,
		174, 176, 178, 180, 182, 184, 186, 188, 190, 192, 0, 6, 2, 0, 97, 97, 100,
		100, 1, 0, 99, 102, 1, 0, 103, 104, 1, 0, 105, 108, 1, 0, 109, 110, 1,
		0, 86, 90, 754, 0, 194, 1, 0, 0, 0, 2, 209, 1, 0, 0, 0, 4, 295, 1, 0, 0,
		0, 6, 297, 1, 0, 0, 0, 8, 302, 1, 0, 0, 0, 10, 309, 1, 0, 0, 0, 12, 321,
		1, 0, 0, 0, 14, 328, 1, 0, 0, 0, 16, 335, 1, 0, 0, 0, 18, 339, 1, 0, 0,
		0, 20, 355, 1, 0, 0, 0, 22, 360, 1, 0, 0, 0, 24, 365, 1, 0, 0, 0, 26, 372,
		1, 0, 0, 0, 28, 377, 1, 0, 0, 0, 30, 386, 1, 0, 0, 0, 32, 391, 1, 0, 0,
		0, 34, 396, 1, 0, 0, 0, 36, 401, 1, 0, 0, 0, 38, 410, 1, 0, 0, 0, 40, 415,
		1, 0, 0, 0, 42, 426, 1, 0, 0, 0, 44, 431, 1, 0, 0, 0, 46, 438, 1, 0, 0,
		0, 48, 443, 1, 0, 0, 0, 50, 447, 1, 0, 0, 0, 52, 452, 1, 0, 0, 0, 54, 457,
		1, 0, 0, 0, 56, 468, 1, 0, 0, 0, 58, 477, 1, 0, 0, 0, 60, 491, 1, 0, 0,
		0, 62, 496, 1, 0, 0, 0, 64, 503, 1, 0, 0, 0, 66, 508, 1, 0, 0, 0, 68, 513,
		1, 0, 0, 0, 70, 517, 1, 0, 0, 0, 72, 522, 1, 0, 0, 0, 74, 527, 1, 0, 0,
		0, 76, 534, 1, 0, 0, 0, 78, 541, 1, 0, 0, 0, 80, 546, 1, 0, 0, 0, 82, 551,
		1, 0, 0, 0, 84, 562, 1, 0, 0, 0, 86, 576, 1, 0, 0, 0, 88, 585, 1, 0, 0,
		0, 90, 599, 1, 0, 0, 0, 92, 606, 1, 0, 0, 0, 94, 611, 1, 0, 0, 0, 96, 616,
		1, 0, 0, 0, 98, 620, 1, 0, 0, 0, 100, 632, 1, 0, 0, 0, 102, 637, 1, 0,
		0, 0, 104, 644, 1, 0, 0, 0, 106, 651, 1, 0, 0, 0, 108, 658, 1, 0, 0, 0,
		110, 667, 1, 0, 0, 0, 112, 672, 1, 0, 0, 0, 114, 677, 1, 0, 0, 0, 116,
		681, 1, 0, 0, 0, 118, 686, 1, 0, 0, 0, 120, 691, 1, 0, 0, 0, 122, 696,
		1, 0, 0, 0, 124, 698, 1, 0, 0, 0, 126, 700, 1, 0, 0, 0, 128, 702, 1, 0,
		0, 0, 130, 704, 1, 0, 0, 0, 132, 706, 1, 0, 0, 0, 134, 708, 1, 0, 0, 0,
		136, 710, 1, 0, 0, 0, 138, 712, 1, 0, 0, 0, 140, 714, 1, 0, 0, 0, 142,
		716, 1, 0, 0, 0, 144, 718, 1, 0, 0, 0, 146, 720, 1, 0, 0, 0, 148, 722,
		1, 0, 0, 0, 150, 724, 1, 0, 0, 0, 152, 726, 1, 0, 0, 0, 154, 728, 1, 0,
		0, 0, 156, 730, 1, 0, 0, 0, 158, 732, 1, 0, 0, 0, 160, 734, 1, 0, 0, 0,
		162, 736, 1, 0, 0, 0, 164, 738, 1, 0, 0, 0, 166, 740, 1, 0, 0, 0, 168,
		742, 1, 0, 0, 0, 170, 744, 1, 0, 0, 0, 172, 746, 1, 0, 0, 0, 174, 748,
		1, 0, 0, 0, 176, 750, 1, 0, 0, 0, 178, 752, 1, 0, 0, 0, 180, 754, 1, 0,
		0, 0, 182, 756, 1, 0, 0, 0, 184, 758, 1, 0, 0, 0, 186, 760, 1, 0, 0, 0,
		188, 762, 1, 0, 0, 0, 190, 769, 1, 0, 0, 0, 192, 771, 1, 0, 0, 0, 194,
		195, 3, 2, 1, 0, 195, 1, 1, 0, 0, 0, 196, 197, 6, 1, -1, 0, 197, 210, 3,
		190, 95, 0, 198, 210, 3, 4, 2, 0, 199, 200, 7, 0, 0, 0, 200, 210, 3, 2,
		1, 8, 201, 206, 5, 91, 0, 0, 202, 203, 5, 94, 0, 0, 203, 205, 5, 91, 0,
		0, 204, 202, 1, 0, 0, 0, 205, 208, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 206,
		207, 1, 0, 0, 0, 207, 210, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 209, 196,
		1, 0, 0, 0, 209, 198, 1, 0, 0, 0, 209, 199, 1, 0, 0, 0, 209, 201, 1, 0,
		0, 0, 210, 234, 1, 0, 0, 0, 211, 212, 10, 6, 0, 0, 212, 213, 5, 96, 0,
		0, 213, 233, 3, 2, 1, 7, 214, 215, 10, 5, 0, 0, 215, 216, 5, 98, 0, 0,
		216, 233, 3, 2, 1, 6, 217, 218, 10, 4, 0, 0, 218, 219, 7, 1, 0, 0, 219,
		233, 3, 2, 1, 5, 220, 221, 10, 3, 0, 0, 221, 223, 7, 2, 0, 0, 222, 224,
		5, 105, 0, 0, 223, 222, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 225, 1,
		0, 0, 0, 225, 233, 3, 2, 1, 4, 226, 227, 10, 2, 0, 0, 227, 228, 7, 3, 0,
		0, 228, 233, 3, 2, 1, 3, 229, 230, 10, 1, 0, 0, 230, 231, 7, 4, 0, 0, 231,
		233, 3, 2, 1, 2, 232, 211, 1, 0, 0, 0, 232, 214, 1, 0, 0, 0, 232, 217,
		1, 0, 0, 0, 232, 220, 1, 0, 0, 0, 232, 226, 1, 0, 0, 0, 232, 229, 1, 0,
		0, 0, 233, 236, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0,
		235, 3, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 237, 296, 3, 6, 3, 0, 238, 296,
		3, 8, 4, 0, 239, 296, 3, 10, 5, 0, 240, 296, 3, 12, 6, 0, 241, 296, 3,
		14, 7, 0, 242, 296, 3, 16, 8, 0, 243, 296, 3, 18, 9, 0, 244, 296, 3, 20,
		10, 0, 245, 296, 3, 22, 11, 0, 246, 296, 3, 24, 12, 0, 247, 296, 3, 26,
		13, 0, 248, 296, 3, 28, 14, 0, 249, 296, 3, 30, 15, 0, 250, 296, 3, 32,
		16, 0, 251, 296, 3, 34, 17, 0, 252, 296, 3, 36, 18, 0, 253, 296, 3, 38,
		19, 0, 254, 296, 3, 40, 20, 0, 255, 296, 3, 42, 21, 0, 256, 296, 3, 44,
		22, 0, 257, 296, 3, 46, 23, 0, 258, 296, 3, 48, 24, 0, 259, 296, 3, 50,
		25, 0, 260, 296, 3, 52, 26, 0, 261, 296, 3, 54, 27, 0, 262, 296, 3, 56,
		28, 0, 263, 296, 3, 58, 29, 0, 264, 296, 3, 60, 30, 0, 265, 296, 3, 66,
		33, 0, 266, 296, 3, 68, 34, 0, 267, 296, 3, 74, 37, 0, 268, 296, 3, 62,
		31, 0, 269, 296, 3, 64, 32, 0, 270, 296, 3, 70, 35, 0, 271, 296, 3, 72,
		36, 0, 272, 296, 3, 76, 38, 0, 273, 296, 3, 78, 39, 0, 274, 296, 3, 82,
		41, 0, 275, 296, 3, 80, 40, 0, 276, 296, 3, 84, 42, 0, 277, 296, 3, 86,
		43, 0, 278, 296, 3, 88, 44, 0, 279, 296, 3, 90, 45, 0, 280, 296, 3, 92,
		46, 0, 281, 296, 3, 94, 47, 0, 282, 296, 3, 96, 48, 0, 283, 296, 3, 98,
		49, 0, 284, 296, 3, 100, 50, 0, 285, 296, 3, 102, 51, 0, 286, 296, 3, 104,
		52, 0, 287, 296, 3, 106, 53, 0, 288, 296, 3, 108, 54, 0, 289, 296, 3, 110,
		55, 0, 290, 296, 3, 112, 56, 0, 291, 296, 3, 114, 57, 0, 292, 296, 3, 116,
		58, 0, 293, 296, 3, 118, 59, 0, 294, 296, 3, 120, 60, 0, 295, 237, 1, 0,
		0, 0, 295, 238, 1, 0, 0, 0, 295, 239, 1, 0, 0, 0, 295, 240, 1, 0, 0, 0,
		295, 241, 1, 0, 0, 0, 295, 242, 1, 0, 0, 0, 295, 243, 1, 0, 0, 0, 295,
		244, 1, 0, 0, 0, 295, 245, 1, 0, 0, 0, 295, 246, 1, 0, 0, 0, 295, 247,
		1, 0, 0, 0, 295, 248, 1, 0, 0, 0, 295, 249, 1, 0, 0, 0, 295, 250, 1, 0,
		0, 0, 295, 251, 1, 0, 0, 0, 295, 252, 1, 0, 0, 0, 295, 253, 1, 0, 0, 0,
		295, 254, 1, 0, 0, 0, 295, 255, 1, 0, 0, 0, 295, 256, 1, 0, 0, 0, 295,
		257, 1, 0, 0, 0, 295, 258, 1, 0, 0, 0, 295, 259, 1, 0, 0, 0, 295, 260,
		1, 0, 0, 0, 295, 261, 1, 0, 0, 0, 295, 262, 1, 0, 0, 0, 295, 263, 1, 0,
		0, 0, 295, 264, 1, 0, 0, 0, 295, 265, 1, 0, 0, 0, 295, 266, 1, 0, 0, 0,
		295, 267, 1, 0, 0, 0, 295, 268, 1, 0, 0, 0, 295, 269, 1, 0, 0, 0, 295,
		270, 1, 0, 0, 0, 295, 271, 1, 0, 0, 0, 295, 272, 1, 0, 0, 0, 295, 273,
		1, 0, 0, 0, 295, 274, 1, 0, 0, 0, 295, 275, 1, 0, 0, 0, 295, 276, 1, 0,
		0, 0, 295, 277, 1, 0, 0, 0, 295, 278, 1, 0, 0, 0, 295, 279, 1, 0, 0, 0,
		295, 280, 1, 0, 0, 0, 295, 281, 1, 0, 0, 0, 295, 282, 1, 0, 0, 0, 295,
		283, 1, 0, 0, 0, 295, 284, 1, 0, 0, 0, 295, 285, 1, 0, 0, 0, 295, 286,
		1, 0, 0, 0, 295, 287, 1, 0, 0, 0, 295, 288, 1, 0, 0, 0, 295, 289, 1, 0,
		0, 0, 295, 290, 1, 0, 0, 0, 295, 291, 1, 0, 0, 0, 295, 292, 1, 0, 0, 0,
		295, 293, 1, 0, 0, 0, 295, 294, 1, 0, 0, 0, 296, 5, 1, 0, 0, 0, 297, 298,
		5, 1, 0, 0, 298, 299, 5, 92, 0, 0, 299, 300, 3, 2, 1, 0, 300, 301, 5, 93,
		0, 0, 301, 7, 1, 0, 0, 0, 302, 303, 5, 2, 0, 0, 303, 304, 5, 92, 0, 0,
		304, 305, 3, 124, 62, 0, 305, 306, 5, 95, 0, 0, 306, 307, 3, 156, 78, 0,
		307, 308, 5, 93, 0, 0, 308, 9, 1, 0, 0, 0, 309, 310, 5, 3, 0, 0, 310, 311,
		5, 92, 0, 0, 311, 316, 3, 2, 1, 0, 312, 313, 5, 95, 0, 0, 313, 315, 3,
		2, 1, 0, 314, 312, 1, 0, 0, 0, 315, 318, 1, 0, 0, 0, 316, 314, 1, 0, 0,
		0, 316, 317, 1, 0, 0, 0, 317, 319, 1, 0, 0, 0, 318, 316, 1, 0, 0, 0, 319,
		320, 5, 93, 0, 0, 320, 11, 1, 0, 0, 0, 321, 322, 5, 4, 0, 0, 322, 323,
		5, 92, 0, 0, 323, 324, 3, 146, 73, 0, 324, 325, 5, 95, 0, 0, 325, 326,
		3, 154, 77, 0, 326, 327, 5, 93, 0, 0, 327, 13, 1, 0, 0, 0, 328, 329, 5,
		5, 0, 0, 329, 330, 5, 92, 0, 0, 330, 331, 3, 2, 1, 0, 331, 332, 5, 95,
		0, 0, 332, 333, 3, 130, 65, 0, 333, 334, 5, 93, 0, 0, 334, 15, 1, 0, 0,
		0, 335, 336, 5, 6, 0, 0, 336, 337, 5, 92, 0, 0, 337, 338, 5, 93, 0, 0,
		338, 17, 1, 0, 0, 0, 339, 340, 5, 7, 0, 0, 340, 341, 5, 92, 0, 0, 341,
		347, 3, 2, 1, 0, 342, 343, 5, 95, 0, 0, 343, 344, 3, 126, 63, 0, 344, 345,
		5, 95, 0, 0, 345, 346, 3, 128, 64, 0, 346, 348, 1, 0, 0, 0, 347, 342, 1,
		0, 0, 0, 348, 349, 1, 0, 0, 0, 349, 347, 1, 0, 0, 0, 349, 350, 1, 0, 0,
		0, 350, 351, 1, 0, 0, 0, 351, 352, 5, 95, 0, 0, 352, 353, 3, 132, 66, 0,
		353, 354, 5, 93, 0, 0, 354, 19, 1, 0, 0, 0, 355, 356, 5, 8, 0, 0, 356,
		357, 5, 92, 0, 0, 357, 358, 3, 2, 1, 0, 358, 359, 5, 93, 0, 0, 359, 21,
		1, 0, 0, 0, 360, 361, 5, 9, 0, 0, 361, 362, 5, 92, 0, 0, 362, 363, 3, 2,
		1, 0, 363, 364, 5, 93, 0, 0, 364, 23, 1, 0, 0, 0, 365, 366, 5, 10, 0, 0,
		366, 367, 5, 92, 0, 0, 367, 368, 3, 146, 73, 0, 368, 369, 5, 95, 0, 0,
		369, 370, 3, 154, 77, 0, 370, 371, 5, 93, 0, 0, 371, 25, 1, 0, 0, 0, 372,
		373, 5, 11, 0, 0, 373, 374, 5, 92, 0, 0, 374, 375, 3, 2, 1, 0, 375, 376,
		5, 93, 0, 0, 376, 27, 1, 0, 0, 0, 377, 378, 5, 12, 0, 0, 378, 379, 5, 92,
		0, 0, 379, 380, 3, 134, 67, 0, 380, 381, 5, 95, 0, 0, 381, 382, 3, 136,
		68, 0, 382, 383, 5, 95, 0, 0, 383, 384, 3, 138, 69, 0, 384, 385, 5, 93,
		0, 0, 385, 29, 1, 0, 0, 0, 386, 387, 5, 13, 0, 0, 387, 388, 5, 92, 0, 0,
		388, 389, 3, 2, 1, 0, 389, 390, 5, 93, 0, 0, 390, 31, 1, 0, 0, 0, 391,
		392, 5, 14, 0, 0, 392, 393, 5, 92, 0, 0, 393, 394, 3, 2, 1, 0, 394, 395,
		5, 93, 0, 0, 395, 33, 1, 0, 0, 0, 396, 397, 5, 15, 0, 0, 397, 398, 5, 92,
		0, 0, 398, 399, 3, 2, 1, 0, 399, 400, 5, 93, 0, 0, 400, 35, 1, 0, 0, 0,
		401, 402, 5, 16, 0, 0, 402, 403, 5, 92, 0, 0, 403, 404, 3, 2, 1, 0, 404,
		405, 5, 95, 0, 0, 405, 406, 3, 2, 1, 0, 406, 407, 5, 95, 0, 0, 407, 408,
		3, 142, 71, 0, 408, 409, 5, 93, 0, 0, 409, 37, 1, 0, 0, 0, 410, 411, 5,
		17, 0, 0, 411, 412, 5, 92, 0, 0, 412, 413, 3, 2, 1, 0, 413, 414, 5, 93,
		0, 0, 414, 39, 1, 0, 0, 0, 415, 416, 5, 18, 0, 0, 416, 417, 5, 92, 0, 0,
		417, 418, 3, 144, 72, 0, 418, 419, 5, 95, 0, 0, 419, 422, 3, 146, 73, 0,
		420, 421, 5, 95, 0, 0, 421, 423, 3, 152, 76, 0, 422, 420, 1, 0, 0, 0, 422,
		423, 1, 0, 0, 0, 423, 424, 1, 0, 0, 0, 424, 425, 5, 93, 0, 0, 425, 41,
		1, 0, 0, 0, 426, 427, 5, 19, 0, 0, 427, 428, 5, 92, 0, 0, 428, 429, 3,
		2, 1, 0, 429, 430, 5, 93, 0, 0, 430, 43, 1, 0, 0, 0, 431, 432, 5, 20, 0,
		0, 432, 433, 5, 92, 0, 0, 433, 434, 3, 162, 81, 0, 434, 435, 5, 95, 0,
		0, 435, 436, 3, 164, 82, 0, 436, 437, 5, 93, 0, 0, 437, 45, 1, 0, 0, 0,
		438, 439, 5, 21, 0, 0, 439, 440, 5, 92, 0, 0, 440, 441, 3, 2, 1, 0, 441,
		442, 5, 93, 0, 0, 442, 47, 1, 0, 0, 0, 443, 444, 5, 22, 0, 0, 444, 445,
		5, 92, 0, 0, 445, 446, 5, 93, 0, 0, 446, 49, 1, 0, 0, 0, 447, 448, 5, 23,
		0, 0, 448, 449, 5, 92, 0, 0, 449, 450, 3, 2, 1, 0, 450, 451, 5, 93, 0,
		0, 451, 51, 1, 0, 0, 0, 452, 453, 5, 24, 0, 0, 453, 454, 5, 92, 0, 0, 454,
		455, 3, 2, 1, 0, 455, 456, 5, 93, 0, 0, 456, 53, 1, 0, 0, 0, 457, 458,
		5, 25, 0, 0, 458, 459, 5, 92, 0, 0, 459, 460, 3, 166, 83, 0, 460, 461,
		5, 95, 0, 0, 461, 464, 3, 168, 84, 0, 462, 463, 5, 95, 0, 0, 463, 465,
		3, 170, 85, 0, 464, 462, 1, 0, 0, 0, 464, 465, 1, 0, 0, 0, 465, 466, 1,
		0, 0, 0, 466, 467, 5, 93, 0, 0, 467, 55, 1, 0, 0, 0, 468, 469, 5, 26, 0,
		0, 469, 470, 5, 92, 0, 0, 470, 471, 3, 172, 86, 0, 471, 472, 5, 95, 0,
		0, 472, 473, 3, 174, 87, 0, 473, 474, 5, 95, 0, 0, 474, 475, 3, 176, 88,
		0, 475, 476, 5, 93, 0, 0, 476, 57, 1, 0, 0, 0, 477, 478, 5, 27, 0, 0, 478,
		479, 5, 92, 0, 0, 479, 480, 3, 166, 83, 0, 480, 481, 5, 95, 0, 0, 481,
		487, 3, 146, 73, 0, 482, 483, 5, 95, 0, 0, 483, 484, 3, 182, 91, 0, 484,
		485, 5, 95, 0, 0, 485, 486, 3, 184, 92, 0, 486, 488, 1, 0, 0, 0, 487, 482,
		1, 0, 0, 0, 487, 488, 1, 0, 0, 0, 488, 489, 1, 0, 0, 0, 489, 490, 5, 93,
		0, 0, 490, 59, 1, 0, 0, 0, 491, 492, 5, 28, 0, 0, 492, 493, 5, 92, 0, 0,
		493, 494, 3, 166, 83, 0, 494, 495, 5, 93, 0, 0, 495, 61, 1, 0, 0, 0, 496,
		497, 5, 30, 0, 0, 497, 498, 5, 92, 0, 0, 498, 499, 3, 122, 61, 0, 499,
		500, 5, 95, 0, 0, 500, 501, 3, 126, 63, 0, 501, 502, 5, 93, 0, 0, 502,
		63, 1, 0, 0, 0, 503, 504, 5, 31, 0, 0, 504, 505, 5, 92, 0, 0, 505, 506,
		3, 2, 1, 0, 506, 507, 5, 93, 0, 0, 507, 65, 1, 0, 0, 0, 508, 509, 5, 32,
		0, 0, 509, 510, 5, 92, 0, 0, 510, 511, 3, 122, 61, 0, 511, 512, 5, 93,
		0, 0, 512, 67, 1, 0, 0, 0, 513, 514, 5, 34, 0, 0, 514, 515, 5, 92, 0, 0,
		515, 516, 5, 93, 0, 0, 516, 69, 1, 0, 0, 0, 517, 518, 5, 35, 0, 0, 518,
		519, 5, 92, 0, 0, 519, 520, 3, 2, 1, 0, 520, 521, 5, 93, 0, 0, 521, 71,
		1, 0, 0, 0, 522, 523, 5, 36, 0, 0, 523, 524, 5, 92, 0, 0, 524, 525, 3,
		2, 1, 0, 525, 526, 5, 93, 0, 0, 526, 73, 1, 0, 0, 0, 527, 528, 5, 37, 0,
		0, 528, 529, 5, 92, 0, 0, 529, 530, 3, 122, 61, 0, 530, 531, 5, 95, 0,
		0, 531, 532, 3, 126, 63, 0, 532, 533, 5, 93, 0, 0, 533, 75, 1, 0, 0, 0,
		534, 535, 5, 41, 0, 0, 535, 536, 5, 92, 0, 0, 536, 537, 3, 146, 73, 0,
		537, 538, 5, 95, 0, 0, 538, 539, 3, 188, 94, 0, 539, 540, 5, 93, 0, 0,
		540, 77, 1, 0, 0, 0, 541, 542, 5, 42, 0, 0, 542, 543, 5, 92, 0, 0, 543,
		544, 3, 2, 1, 0, 544, 545, 5, 93, 0, 0, 545, 79, 1, 0, 0, 0, 546, 547,
		5, 46, 0, 0, 547, 548, 5, 92, 0, 0, 548, 549, 3, 2, 1, 0, 549, 550, 5,
		93, 0, 0, 550, 81, 1, 0, 0, 0, 551, 552, 5, 47, 0, 0, 552, 553, 5, 92,
		0, 0, 553, 554, 3, 146, 73, 0, 554, 555, 5, 95, 0, 0, 555, 558, 3, 178,
		89, 0, 556, 557, 5, 95, 0, 0, 557, 559, 3, 180, 90, 0, 558, 556, 1, 0,
		0, 0, 558, 559, 1, 0, 0, 0, 559, 560, 1, 0, 0, 0, 560, 561, 5, 93, 0, 0,
		561, 83, 1, 0, 0, 0, 562, 563, 5, 48, 0, 0, 563, 564, 5, 92, 0, 0, 564,
		565, 3, 2, 1, 0, 565, 566, 5, 95, 0, 0, 566, 571, 3, 2, 1, 0, 567, 568,
		5, 95, 0, 0, 568, 570, 3, 2, 1, 0, 569, 567, 1, 0, 0, 0, 570, 573, 1, 0,
		0, 0, 571, 569, 1, 0, 0, 0, 571, 572, 1, 0, 0, 0, 572, 574, 1, 0, 0, 0,
		573, 571, 1, 0, 0, 0, 574, 575, 5, 93, 0, 0, 575, 85, 1, 0, 0, 0, 576,
		577, 5, 51, 0, 0, 577, 578, 5, 92, 0, 0, 578, 579, 3, 146, 73, 0, 579,
		580, 5, 95, 0, 0, 580, 581, 3, 152, 76, 0, 581, 582, 5, 95, 0, 0, 582,
		583, 3, 188, 94, 0, 583, 584, 5, 93, 0, 0, 584, 87, 1, 0, 0, 0, 585, 586,
		5, 53, 0, 0, 586, 587, 5, 92, 0, 0, 587, 588, 3, 2, 1, 0, 588, 589, 5,
		95, 0, 0, 589, 594, 3, 2, 1, 0, 590, 591, 5, 95, 0, 0, 591, 593, 3, 2,
		1, 0, 592, 590, 1, 0, 0, 0, 593, 596, 1, 0, 0, 0, 594, 592, 1, 0, 0, 0,
		594, 595, 1, 0, 0, 0, 595, 597, 1, 0, 0, 0, 596, 594, 1, 0, 0, 0, 597,
		598, 5, 93, 0, 0, 598, 89, 1, 0, 0, 0, 599, 600, 5, 55, 0, 0, 600, 601,
		5, 92, 0, 0, 601, 602, 3, 156, 78, 0, 602, 603, 5, 95, 0, 0, 603, 604,
		3, 158, 79, 0, 604, 605, 5, 93, 0, 0, 605, 91, 1, 0, 0, 0, 606, 607, 5,
		56, 0, 0, 607, 608, 5, 92, 0, 0, 608, 609, 3, 2, 1, 0, 609, 610, 5, 93,
		0, 0, 610, 93, 1, 0, 0, 0, 611, 612, 5, 57, 0, 0, 612, 613, 5, 92, 0, 0,
		613, 614, 3, 2, 1, 0, 614, 615, 5, 93, 0, 0, 615, 95, 1, 0, 0, 0, 616,
		617, 5, 58, 0, 0, 617, 618, 5, 92, 0, 0, 618, 619, 5, 93, 0, 0, 619, 97,
		1, 0, 0, 0, 620, 621, 5, 60, 0, 0, 621, 622, 5, 92, 0, 0, 622, 627, 3,
		2, 1, 0, 623, 624, 5, 95, 0, 0, 624, 626, 3, 2, 1, 0, 625, 623, 1, 0, 0,
		0, 626, 629, 1, 0, 0, 0, 627, 625, 1, 0, 0, 0, 627, 628, 1, 0, 0, 0, 628,
		630, 1, 0, 0, 0, 629, 627, 1, 0, 0, 0, 630, 631, 5, 93, 0, 0, 631, 99,
		1, 0, 0, 0, 632, 633, 5, 64, 0, 0, 633, 634, 5, 92, 0, 0, 634, 635, 3,
		122, 61, 0, 635, 636, 5, 93, 0, 0, 636, 101, 1, 0, 0, 0, 637, 638, 5, 65,
		0, 0, 638, 639, 5, 92, 0, 0, 639, 640, 3, 146, 73, 0, 640, 641, 5, 95,
		0, 0, 641, 642, 3, 140, 70, 0, 642, 643, 5, 93, 0, 0, 643, 103, 1, 0, 0,
		0, 644, 645, 5, 68, 0, 0, 645, 646, 5, 92, 0, 0, 646, 647, 3, 146, 73,
		0, 647, 648, 5, 95, 0, 0, 648, 649, 3, 188, 94, 0, 649, 650, 5, 93, 0,
		0, 650, 105, 1, 0, 0, 0, 651, 652, 5, 69, 0, 0, 652, 653, 5, 92, 0, 0,
		653, 654, 3, 156, 78, 0, 654, 655, 5, 95, 0, 0, 655, 656, 3, 160, 80, 0,
		656, 657, 5, 93, 0, 0, 657, 107, 1, 0, 0, 0, 658, 659, 5, 73, 0, 0, 659,
		660, 5, 92, 0, 0, 660, 661, 3, 146, 73, 0, 661, 662, 5, 95, 0, 0, 662,
		663, 3, 148, 74, 0, 663, 664, 5, 95, 0, 0, 664, 665, 3, 150, 75, 0, 665,
		666, 5, 93, 0, 0, 666, 109, 1, 0, 0, 0, 667, 668, 5, 78, 0, 0, 668, 669,
		5, 92, 0, 0, 669, 670, 3, 2, 1, 0, 670, 671, 5, 93, 0, 0, 671, 111, 1,
		0, 0, 0, 672, 673, 5, 74, 0, 0, 673, 674, 5, 92, 0, 0, 674, 675, 3, 2,
		1, 0, 675, 676, 5, 93, 0, 0, 676, 113, 1, 0, 0, 0, 677, 678, 5, 77, 0,
		0, 678, 679, 5, 92, 0, 0, 679, 680, 5, 93, 0, 0, 680, 115, 1, 0, 0, 0,
		681, 682, 5, 79, 0, 0, 682, 683, 5, 92, 0, 0, 683, 684, 3, 2, 1, 0, 684,
		685, 5, 93, 0, 0, 685, 117, 1, 0, 0, 0, 686, 687, 5, 82, 0, 0, 687, 688,
		5, 92, 0, 0, 688, 689, 3, 2, 1, 0, 689, 690, 5, 93, 0, 0, 690, 119, 1,
		0, 0, 0, 691, 692, 5, 85, 0, 0, 692, 693, 5, 92, 0, 0, 693, 694, 3, 2,
		1, 0, 694, 695, 5, 93, 0, 0, 695, 121, 1, 0, 0, 0, 696, 697, 3, 2, 1, 0,
		697, 123, 1, 0, 0, 0, 698, 699, 3, 2, 1, 0, 699, 125, 1, 0, 0, 0, 700,
		701, 3, 2, 1, 0, 701, 127, 1, 0, 0, 0, 702, 703, 3, 2, 1, 0, 703, 129,
		1, 0, 0, 0, 704, 705, 3, 2, 1, 0, 705, 131, 1, 0, 0, 0, 706, 707, 3, 2,
		1, 0, 707, 133, 1, 0, 0, 0, 708, 709, 3, 2, 1, 0, 709, 135, 1, 0, 0, 0,
		710, 711, 3, 2, 1, 0, 711, 137, 1, 0, 0, 0, 712, 713, 3, 2, 1, 0, 713,
		139, 1, 0, 0, 0, 714, 715, 3, 2, 1, 0, 715, 141, 1, 0, 0, 0, 716, 717,
		3, 2, 1, 0, 717, 143, 1, 0, 0, 0, 718, 719, 3, 2, 1, 0, 719, 145, 1, 0,
		0, 0, 720, 721, 3, 2, 1, 0, 721, 147, 1, 0, 0, 0, 722, 723, 3, 2, 1, 0,
		723, 149, 1, 0, 0, 0, 724, 725, 3, 2, 1, 0, 725, 151, 1, 0, 0, 0, 726,
		727, 3, 2, 1, 0, 727, 153, 1, 0, 0, 0, 728, 729, 3, 2, 1, 0, 729, 155,
		1, 0, 0, 0, 730, 731, 3, 2, 1, 0, 731, 157, 1, 0, 0, 0, 732, 733, 3, 2,
		1, 0, 733, 159, 1, 0, 0, 0, 734, 735, 3, 2, 1, 0, 735, 161, 1, 0, 0, 0,
		736, 737, 3, 2, 1, 0, 737, 163, 1, 0, 0, 0, 738, 739, 3, 2, 1, 0, 739,
		165, 1, 0, 0, 0, 740, 741, 3, 2, 1, 0, 741, 167, 1, 0, 0, 0, 742, 743,
		3, 2, 1, 0, 743, 169, 1, 0, 0, 0, 744, 745, 3, 2, 1, 0, 745, 171, 1, 0,
		0, 0, 746, 747, 3, 2, 1, 0, 747, 173, 1, 0, 0, 0, 748, 749, 3, 2, 1, 0,
		749, 175, 1, 0, 0, 0, 750, 751, 3, 2, 1, 0, 751, 177, 1, 0, 0, 0, 752,
		753, 3, 2, 1, 0, 753, 179, 1, 0, 0, 0, 754, 755, 3, 2, 1, 0, 755, 181,
		1, 0, 0, 0, 756, 757, 3, 2, 1, 0, 757, 183, 1, 0, 0, 0, 758, 759, 3, 2,
		1, 0, 759, 185, 1, 0, 0, 0, 760, 761, 3, 2, 1, 0, 761, 187, 1, 0, 0, 0,
		762, 763, 3, 2, 1, 0, 763, 189, 1, 0, 0, 0, 764, 770, 3, 192, 96, 0, 765,
		766, 5, 92, 0, 0, 766, 767, 3, 2, 1, 0, 767, 768, 5, 93, 0, 0, 768, 770,
		1, 0, 0, 0, 769, 764, 1, 0, 0, 0, 769, 765, 1, 0, 0, 0, 770, 191, 1, 0,
		0, 0, 771, 772, 7, 5, 0, 0, 772, 193, 1, 0, 0, 0, 16, 206, 209, 223, 232,
		234, 295, 316, 349, 422, 464, 487, 558, 571, 594, 627, 769,
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
	FormulaParserBANG                 = 97
	FormulaParserBITAND               = 98
	FormulaParserADD                  = 99
	FormulaParserSUB                  = 100
	FormulaParserMUL                  = 101
	FormulaParserDIV                  = 102
	FormulaParserLT                   = 103
	FormulaParserGT                   = 104
	FormulaParserASSIGN               = 105
	FormulaParserLESSANDGREATER       = 106
	FormulaParserNOTEQUAL             = 107
	FormulaParserEQUAL                = 108
	FormulaParserAND                  = 109
	FormulaParserOR                   = 110
	FormulaParserWS                   = 111
	FormulaParserCOMMENT              = 112
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
	FormulaParserRULE_includes            = 31
	FormulaParserRULE_isblank             = 32
	FormulaParserRULE_ischanged           = 33
	FormulaParserRULE_isnew               = 34
	FormulaParserRULE_isnull              = 35
	FormulaParserRULE_isnumber            = 36
	FormulaParserRULE_ispickval           = 37
	FormulaParserRULE_left                = 38
	FormulaParserRULE_len                 = 39
	FormulaParserRULE_lower               = 40
	FormulaParserRULE_lpad                = 41
	FormulaParserRULE_max                 = 42
	FormulaParserRULE_mid                 = 43
	FormulaParserRULE_min                 = 44
	FormulaParserRULE_mod                 = 45
	FormulaParserRULE_month               = 46
	FormulaParserRULE_not                 = 47
	FormulaParserRULE_now                 = 48
	FormulaParserRULE_or                  = 49
	FormulaParserRULE_priorvalue          = 50
	FormulaParserRULE_regex               = 51
	FormulaParserRULE_right               = 52
	FormulaParserRULE_round               = 53
	FormulaParserRULE_substitute          = 54
	FormulaParserRULE_trim                = 55
	FormulaParserRULE_text                = 56
	FormulaParserRULE_today               = 57
	FormulaParserRULE_upper               = 58
	FormulaParserRULE_value               = 59
	FormulaParserRULE_year                = 60
	FormulaParserRULE_fieldExpression     = 61
	FormulaParserRULE_dateExpression      = 62
	FormulaParserRULE_valueExpression     = 63
	FormulaParserRULE_resultExpression    = 64
	FormulaParserRULE_substituteValue     = 65
	FormulaParserRULE_defaultExpression   = 66
	FormulaParserRULE_yearExpression      = 67
	FormulaParserRULE_monthExpression     = 68
	FormulaParserRULE_dayExpression       = 69
	FormulaParserRULE_regexExpression     = 70
	FormulaParserRULE_unitExpression      = 71
	FormulaParserRULE_searchExpression    = 72
	FormulaParserRULE_textExpression      = 73
	FormulaParserRULE_oldText             = 74
	FormulaParserRULE_replacement         = 75
	FormulaParserRULE_startNum            = 76
	FormulaParserRULE_compareText         = 77
	FormulaParserRULE_num                 = 78
	FormulaParserRULE_divisor             = 79
	FormulaParserRULE_digits              = 80
	FormulaParserRULE_latitudeExpression  = 81
	FormulaParserRULE_longitudeExpression = 82
	FormulaParserRULE_url                 = 83
	FormulaParserRULE_name                = 84
	FormulaParserRULE_target              = 85
	FormulaParserRULE_logicalExpression   = 86
	FormulaParserRULE_ifTrueExpression    = 87
	FormulaParserRULE_ifFalseExpression   = 88
	FormulaParserRULE_length              = 89
	FormulaParserRULE_padString           = 90
	FormulaParserRULE_heightExpression    = 91
	FormulaParserRULE_widthExpression     = 92
	FormulaParserRULE_start               = 93
	FormulaParserRULE_numChars            = 94
	FormulaParserRULE_primary             = 95
	FormulaParserRULE_literal             = 96
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
		p.SetState(194)
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

func (s *NegationExpressionContext) BANG() antlr.TerminalNode {
	return s.GetToken(FormulaParserBANG, 0)
}

func (s *NegationExpressionContext) SUB() antlr.TerminalNode {
	return s.GetToken(FormulaParserSUB, 0)
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

func (s *EqualityExpressionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(FormulaParserASSIGN, 0)
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
	p.SetState(209)
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
			p.SetState(197)
			p.Primary()
		}

	case FormulaParserABS, FormulaParserADDMONTHS, FormulaParserAND_FUNC, FormulaParserBEGINS, FormulaParserBLANKVALUE, FormulaParserBR, FormulaParserCASE, FormulaParserCASESAFEID, FormulaParserCEILING, FormulaParserCONTAINS, FormulaParserCURRENCYRATE, FormulaParserDATE, FormulaParserDATEVALUE, FormulaParserDATETIMEVALUE, FormulaParserDAY, FormulaParserDISTANCE, FormulaParserEXP, FormulaParserFIND, FormulaParserFLOOR, FormulaParserGEOLOCATION, FormulaParserGETRECORDIDS, FormulaParserGETSESSIONID, FormulaParserHOUR, FormulaParserHTMLENCODE, FormulaParserHYPERLINK, FormulaParserIF, FormulaParserIMAGE, FormulaParserIMAGEPROXYURL, FormulaParserINCLUDES, FormulaParserISBLANK, FormulaParserISCHANGED, FormulaParserISNEW, FormulaParserISNULL, FormulaParserISNUMBER, FormulaParserISPICKVAL, FormulaParserLEFT, FormulaParserLEN, FormulaParserLOWER, FormulaParserLPAD, FormulaParserMAX, FormulaParserMID, FormulaParserMIN, FormulaParserMOD, FormulaParserMONTH, FormulaParserNOT, FormulaParserNOW, FormulaParserOR_FUNC, FormulaParserPRIORVALUE, FormulaParserREGEX, FormulaParserRIGHT, FormulaParserROUND, FormulaParserSUBSTITUTE, FormulaParserTEXT, FormulaParserTODAY, FormulaParserTRIM, FormulaParserUPPER, FormulaParserVALUE, FormulaParserYEAR:
		localctx = NewFunctionCallExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(198)
			p.FunctionCall()
		}

	case FormulaParserBANG, FormulaParserSUB:
		localctx = NewNegationExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(199)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FormulaParserBANG || _la == FormulaParserSUB) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(200)
			p.expression(8)
		}

	case FormulaParserIdentifier:
		localctx = NewFieldReferenceContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(201)
			p.Match(FormulaParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(206)
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
					p.SetState(202)
					p.Match(FormulaParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(203)
					p.Match(FormulaParserIdentifier)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(208)
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
	p.SetState(234)
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
			p.SetState(232)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExponentiationExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expression)
				p.SetState(211)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(212)
					p.Match(FormulaParserCARET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(213)
					p.expression(7)
				}

			case 2:
				localctx = NewConcatExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expression)
				p.SetState(214)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(215)
					p.Match(FormulaParserBITAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(216)
					p.expression(6)
				}

			case 3:
				localctx = NewArithExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expression)
				p.SetState(217)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(218)
					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-99)) & ^0x3f) == 0 && ((int64(1)<<(_la-99))&15) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(219)
					p.expression(5)
				}

			case 4:
				localctx = NewCompareExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expression)
				p.SetState(220)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(221)
					_la = p.GetTokenStream().LA(1)

					if !(_la == FormulaParserLT || _la == FormulaParserGT) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(223)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == FormulaParserASSIGN {
					{
						p.SetState(222)
						p.Match(FormulaParserASSIGN)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(225)
					p.expression(4)
				}

			case 5:
				localctx = NewEqualityExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expression)
				p.SetState(226)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(227)
					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-105)) & ^0x3f) == 0 && ((int64(1)<<(_la-105))&15) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(228)
					p.expression(3)
				}

			case 6:
				localctx = NewLogicExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expression)
				p.SetState(229)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(230)
					_la = p.GetTokenStream().LA(1)

					if !(_la == FormulaParserAND || _la == FormulaParserOR) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(231)
					p.expression(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(236)
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
	Ischanged() IIschangedContext
	Isnew() IIsnewContext
	Ispickval() IIspickvalContext
	Includes() IIncludesContext
	Isblank() IIsblankContext
	Isnull() IIsnullContext
	Isnumber() IIsnumberContext
	Left() ILeftContext
	Len_() ILenContext
	Lpad() ILpadContext
	Lower() ILowerContext
	Max() IMaxContext
	Mid() IMidContext
	Min() IMinContext
	Mod() IModContext
	Month() IMonthContext
	Not() INotContext
	Now() INowContext
	Or() IOrContext
	Priorvalue() IPriorvalueContext
	Regex() IRegexContext
	Right() IRightContext
	Round() IRoundContext
	Substitute() ISubstituteContext
	Trim() ITrimContext
	Text() ITextContext
	Today() ITodayContext
	Upper() IUpperContext
	Value() IValueContext
	Year() IYearContext

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

func (s *FunctionCallContext) Ischanged() IIschangedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIschangedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIschangedContext)
}

func (s *FunctionCallContext) Isnew() IIsnewContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIsnewContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIsnewContext)
}

func (s *FunctionCallContext) Ispickval() IIspickvalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIspickvalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIspickvalContext)
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

func (s *FunctionCallContext) Isblank() IIsblankContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIsblankContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIsblankContext)
}

func (s *FunctionCallContext) Isnull() IIsnullContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIsnullContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIsnullContext)
}

func (s *FunctionCallContext) Isnumber() IIsnumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIsnumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIsnumberContext)
}

func (s *FunctionCallContext) Left() ILeftContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeftContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeftContext)
}

func (s *FunctionCallContext) Len_() ILenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILenContext)
}

func (s *FunctionCallContext) Lpad() ILpadContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILpadContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILpadContext)
}

func (s *FunctionCallContext) Lower() ILowerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILowerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILowerContext)
}

func (s *FunctionCallContext) Max() IMaxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMaxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMaxContext)
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

func (s *FunctionCallContext) Min() IMinContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMinContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMinContext)
}

func (s *FunctionCallContext) Mod() IModContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModContext)
}

func (s *FunctionCallContext) Month() IMonthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMonthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMonthContext)
}

func (s *FunctionCallContext) Not() INotContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotContext)
}

func (s *FunctionCallContext) Now() INowContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INowContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INowContext)
}

func (s *FunctionCallContext) Or() IOrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrContext)
}

func (s *FunctionCallContext) Priorvalue() IPriorvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPriorvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPriorvalueContext)
}

func (s *FunctionCallContext) Regex() IRegexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexContext)
}

func (s *FunctionCallContext) Right() IRightContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRightContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRightContext)
}

func (s *FunctionCallContext) Round() IRoundContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoundContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoundContext)
}

func (s *FunctionCallContext) Substitute() ISubstituteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubstituteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubstituteContext)
}

func (s *FunctionCallContext) Trim() ITrimContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITrimContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITrimContext)
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

func (s *FunctionCallContext) Today() ITodayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITodayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITodayContext)
}

func (s *FunctionCallContext) Upper() IUpperContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpperContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUpperContext)
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

func (s *FunctionCallContext) Year() IYearContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IYearContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IYearContext)
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
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FormulaParserABS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(237)
			p.Abs()
		}

	case FormulaParserADDMONTHS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(238)
			p.AddMonths()
		}

	case FormulaParserAND_FUNC:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(239)
			p.And()
		}

	case FormulaParserBEGINS:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(240)
			p.Begins()
		}

	case FormulaParserBLANKVALUE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(241)
			p.Blankvalue()
		}

	case FormulaParserBR:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(242)
			p.Br()
		}

	case FormulaParserCASE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(243)
			p.Case_()
		}

	case FormulaParserCASESAFEID:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(244)
			p.Casesafeid()
		}

	case FormulaParserCEILING:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(245)
			p.Ceiling()
		}

	case FormulaParserCONTAINS:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(246)
			p.Contains()
		}

	case FormulaParserCURRENCYRATE:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(247)
			p.Currencyrate()
		}

	case FormulaParserDATE:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(248)
			p.Date()
		}

	case FormulaParserDATEVALUE:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(249)
			p.Datevalue()
		}

	case FormulaParserDATETIMEVALUE:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(250)
			p.Datetimevalue()
		}

	case FormulaParserDAY:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(251)
			p.Day()
		}

	case FormulaParserDISTANCE:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(252)
			p.Distance()
		}

	case FormulaParserEXP:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(253)
			p.Exp()
		}

	case FormulaParserFIND:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(254)
			p.Find()
		}

	case FormulaParserFLOOR:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(255)
			p.Floor()
		}

	case FormulaParserGEOLOCATION:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(256)
			p.Geolocation()
		}

	case FormulaParserGETRECORDIDS:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(257)
			p.Getrecordids()
		}

	case FormulaParserGETSESSIONID:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(258)
			p.Getsessionid()
		}

	case FormulaParserHOUR:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(259)
			p.Hour()
		}

	case FormulaParserHTMLENCODE:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(260)
			p.Htmlencode()
		}

	case FormulaParserHYPERLINK:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(261)
			p.Hyperlink()
		}

	case FormulaParserIF:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(262)
			p.If_()
		}

	case FormulaParserIMAGE:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(263)
			p.Image()
		}

	case FormulaParserIMAGEPROXYURL:
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(264)
			p.Imageproxyurl()
		}

	case FormulaParserISCHANGED:
		p.EnterOuterAlt(localctx, 29)
		{
			p.SetState(265)
			p.Ischanged()
		}

	case FormulaParserISNEW:
		p.EnterOuterAlt(localctx, 30)
		{
			p.SetState(266)
			p.Isnew()
		}

	case FormulaParserISPICKVAL:
		p.EnterOuterAlt(localctx, 31)
		{
			p.SetState(267)
			p.Ispickval()
		}

	case FormulaParserINCLUDES:
		p.EnterOuterAlt(localctx, 32)
		{
			p.SetState(268)
			p.Includes()
		}

	case FormulaParserISBLANK:
		p.EnterOuterAlt(localctx, 33)
		{
			p.SetState(269)
			p.Isblank()
		}

	case FormulaParserISNULL:
		p.EnterOuterAlt(localctx, 34)
		{
			p.SetState(270)
			p.Isnull()
		}

	case FormulaParserISNUMBER:
		p.EnterOuterAlt(localctx, 35)
		{
			p.SetState(271)
			p.Isnumber()
		}

	case FormulaParserLEFT:
		p.EnterOuterAlt(localctx, 36)
		{
			p.SetState(272)
			p.Left()
		}

	case FormulaParserLEN:
		p.EnterOuterAlt(localctx, 37)
		{
			p.SetState(273)
			p.Len_()
		}

	case FormulaParserLPAD:
		p.EnterOuterAlt(localctx, 38)
		{
			p.SetState(274)
			p.Lpad()
		}

	case FormulaParserLOWER:
		p.EnterOuterAlt(localctx, 39)
		{
			p.SetState(275)
			p.Lower()
		}

	case FormulaParserMAX:
		p.EnterOuterAlt(localctx, 40)
		{
			p.SetState(276)
			p.Max()
		}

	case FormulaParserMID:
		p.EnterOuterAlt(localctx, 41)
		{
			p.SetState(277)
			p.Mid()
		}

	case FormulaParserMIN:
		p.EnterOuterAlt(localctx, 42)
		{
			p.SetState(278)
			p.Min()
		}

	case FormulaParserMOD:
		p.EnterOuterAlt(localctx, 43)
		{
			p.SetState(279)
			p.Mod()
		}

	case FormulaParserMONTH:
		p.EnterOuterAlt(localctx, 44)
		{
			p.SetState(280)
			p.Month()
		}

	case FormulaParserNOT:
		p.EnterOuterAlt(localctx, 45)
		{
			p.SetState(281)
			p.Not()
		}

	case FormulaParserNOW:
		p.EnterOuterAlt(localctx, 46)
		{
			p.SetState(282)
			p.Now()
		}

	case FormulaParserOR_FUNC:
		p.EnterOuterAlt(localctx, 47)
		{
			p.SetState(283)
			p.Or()
		}

	case FormulaParserPRIORVALUE:
		p.EnterOuterAlt(localctx, 48)
		{
			p.SetState(284)
			p.Priorvalue()
		}

	case FormulaParserREGEX:
		p.EnterOuterAlt(localctx, 49)
		{
			p.SetState(285)
			p.Regex()
		}

	case FormulaParserRIGHT:
		p.EnterOuterAlt(localctx, 50)
		{
			p.SetState(286)
			p.Right()
		}

	case FormulaParserROUND:
		p.EnterOuterAlt(localctx, 51)
		{
			p.SetState(287)
			p.Round()
		}

	case FormulaParserSUBSTITUTE:
		p.EnterOuterAlt(localctx, 52)
		{
			p.SetState(288)
			p.Substitute()
		}

	case FormulaParserTRIM:
		p.EnterOuterAlt(localctx, 53)
		{
			p.SetState(289)
			p.Trim()
		}

	case FormulaParserTEXT:
		p.EnterOuterAlt(localctx, 54)
		{
			p.SetState(290)
			p.Text()
		}

	case FormulaParserTODAY:
		p.EnterOuterAlt(localctx, 55)
		{
			p.SetState(291)
			p.Today()
		}

	case FormulaParserUPPER:
		p.EnterOuterAlt(localctx, 56)
		{
			p.SetState(292)
			p.Upper()
		}

	case FormulaParserVALUE:
		p.EnterOuterAlt(localctx, 57)
		{
			p.SetState(293)
			p.Value()
		}

	case FormulaParserYEAR:
		p.EnterOuterAlt(localctx, 58)
		{
			p.SetState(294)
			p.Year()
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
		p.SetState(297)
		p.Match(FormulaParserABS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(298)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(299)
		p.expression(0)
	}
	{
		p.SetState(300)
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
	DateExpression() IDateExpressionContext
	COMMA() antlr.TerminalNode
	Num() INumContext
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

func (s *AddMonthsContext) DateExpression() IDateExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateExpressionContext)
}

func (s *AddMonthsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, 0)
}

func (s *AddMonthsContext) Num() INumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumContext)
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
		p.SetState(302)
		p.Match(FormulaParserADDMONTHS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(303)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(304)
		p.DateExpression()
	}
	{
		p.SetState(305)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)
		p.Num()
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
	RPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

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

func (s *AndContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *AndContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FormulaParserCOMMA)
}

func (s *AndContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, i)
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
		p.SetState(309)
		p.Match(FormulaParserAND_FUNC)
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
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FormulaParserCOMMA {
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

		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(319)
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
	TextExpression() ITextExpressionContext
	COMMA() antlr.TerminalNode
	CompareText() ICompareTextContext
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

func (s *BeginsContext) TextExpression() ITextExpressionContext {
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

func (s *BeginsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, 0)
}

func (s *BeginsContext) CompareText() ICompareTextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompareTextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompareTextContext)
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
		p.SetState(321)
		p.Match(FormulaParserBEGINS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(322)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(323)
		p.TextExpression()
	}
	{
		p.SetState(324)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(325)
		p.CompareText()
	}
	{
		p.SetState(326)
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
	Expression() IExpressionContext
	COMMA() antlr.TerminalNode
	SubstituteValue() ISubstituteValueContext
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

func (s *BlankvalueContext) Expression() IExpressionContext {
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

func (s *BlankvalueContext) COMMA() antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, 0)
}

func (s *BlankvalueContext) SubstituteValue() ISubstituteValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubstituteValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubstituteValueContext)
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
		p.SetState(328)
		p.Match(FormulaParserBLANKVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(329)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(330)
		p.expression(0)
	}
	{
		p.SetState(331)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(332)
		p.SubstituteValue()
	}
	{
		p.SetState(333)
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
		p.SetState(335)
		p.Match(FormulaParserBR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(336)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
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
	DefaultExpression() IDefaultExpressionContext
	RPAREN() antlr.TerminalNode
	AllValueExpression() []IValueExpressionContext
	ValueExpression(i int) IValueExpressionContext
	AllResultExpression() []IResultExpressionContext
	ResultExpression(i int) IResultExpressionContext

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
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.Match(FormulaParserCASE)
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
		p.expression(0)
	}
	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
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
				p.ValueExpression()
			}
			{
				p.SetState(344)
				p.Match(FormulaParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(345)
				p.ResultExpression()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(349)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(351)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(352)
		p.DefaultExpression()
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
		p.SetState(355)
		p.Match(FormulaParserCASESAFEID)
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
		p.SetState(360)
		p.Match(FormulaParserCEILING)
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

// IContainsContext is an interface to support dynamic dispatch.
type IContainsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTAINS() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	TextExpression() ITextExpressionContext
	COMMA() antlr.TerminalNode
	CompareText() ICompareTextContext
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

func (s *ContainsContext) TextExpression() ITextExpressionContext {
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

func (s *ContainsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, 0)
}

func (s *ContainsContext) CompareText() ICompareTextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompareTextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompareTextContext)
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
		p.SetState(365)
		p.Match(FormulaParserCONTAINS)
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
		p.TextExpression()
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
		p.CompareText()
	}
	{
		p.SetState(370)
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
		p.SetState(372)
		p.Match(FormulaParserCURRENCYRATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(373)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(374)
		p.expression(0)
	}
	{
		p.SetState(375)
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
		p.SetState(377)
		p.Match(FormulaParserDATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(378)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(379)
		p.YearExpression()
	}
	{
		p.SetState(380)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(381)
		p.MonthExpression()
	}
	{
		p.SetState(382)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(383)
		p.DayExpression()
	}
	{
		p.SetState(384)
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
		p.SetState(386)
		p.Match(FormulaParserDATEVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(387)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(388)
		p.expression(0)
	}
	{
		p.SetState(389)
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
		p.SetState(391)
		p.Match(FormulaParserDATETIMEVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(392)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(393)
		p.expression(0)
	}
	{
		p.SetState(394)
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
		p.SetState(396)
		p.Match(FormulaParserDAY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(397)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(398)
		p.expression(0)
	}
	{
		p.SetState(399)
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
		p.SetState(401)
		p.Match(FormulaParserDISTANCE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(402)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(403)
		p.expression(0)
	}
	{
		p.SetState(404)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(405)
		p.expression(0)
	}
	{
		p.SetState(406)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(407)
		p.UnitExpression()
	}
	{
		p.SetState(408)
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
		p.SetState(410)
		p.Match(FormulaParserEXP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(411)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(412)
		p.expression(0)
	}
	{
		p.SetState(413)
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
		p.SetState(415)
		p.Match(FormulaParserFIND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(416)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(417)
		p.SearchExpression()
	}
	{
		p.SetState(418)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(419)
		p.TextExpression()
	}
	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FormulaParserCOMMA {
		{
			p.SetState(420)
			p.Match(FormulaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(421)
			p.StartNum()
		}

	}
	{
		p.SetState(424)
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
		p.SetState(426)
		p.Match(FormulaParserFLOOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(427)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(428)
		p.expression(0)
	}
	{
		p.SetState(429)
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
		p.SetState(431)
		p.Match(FormulaParserGEOLOCATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(432)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(433)
		p.LatitudeExpression()
	}
	{
		p.SetState(434)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(435)
		p.LongitudeExpression()
	}
	{
		p.SetState(436)
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
		p.SetState(438)
		p.Match(FormulaParserGETRECORDIDS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(439)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(440)
		p.expression(0)
	}
	{
		p.SetState(441)
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
		p.SetState(443)
		p.Match(FormulaParserGETSESSIONID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(444)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(445)
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
		p.SetState(447)
		p.Match(FormulaParserHOUR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(448)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(449)
		p.expression(0)
	}
	{
		p.SetState(450)
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
		p.SetState(452)
		p.Match(FormulaParserHTMLENCODE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(453)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(454)
		p.expression(0)
	}
	{
		p.SetState(455)
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
	Url() IUrlContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	Name() INameContext
	RPAREN() antlr.TerminalNode
	Target() ITargetContext

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

func (s *HyperlinkContext) Url() IUrlContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUrlContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUrlContext)
}

func (s *HyperlinkContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FormulaParserCOMMA)
}

func (s *HyperlinkContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, i)
}

func (s *HyperlinkContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *HyperlinkContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *HyperlinkContext) Target() ITargetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
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
		p.SetState(457)
		p.Match(FormulaParserHYPERLINK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(458)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(459)
		p.Url()
	}
	{
		p.SetState(460)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(461)
		p.Name()
	}
	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FormulaParserCOMMA {
		{
			p.SetState(462)
			p.Match(FormulaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(463)
			p.Target()
		}

	}
	{
		p.SetState(466)
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
		p.SetState(468)
		p.Match(FormulaParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(469)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(470)
		p.LogicalExpression()
	}
	{
		p.SetState(471)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(472)
		p.IfTrueExpression()
	}
	{
		p.SetState(473)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(474)
		p.IfFalseExpression()
	}
	{
		p.SetState(475)
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
	Url() IUrlContext
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

func (s *ImageContext) Url() IUrlContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUrlContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUrlContext)
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
		p.SetState(477)
		p.Match(FormulaParserIMAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(478)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(479)
		p.Url()
	}
	{
		p.SetState(480)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(481)
		p.TextExpression()
	}
	p.SetState(487)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FormulaParserCOMMA {
		{
			p.SetState(482)
			p.Match(FormulaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(483)
			p.HeightExpression()
		}
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
			p.WidthExpression()
		}

	}
	{
		p.SetState(489)
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
	Url() IUrlContext
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

func (s *ImageproxyurlContext) Url() IUrlContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUrlContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUrlContext)
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
		p.SetState(491)
		p.Match(FormulaParserIMAGEPROXYURL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(492)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(493)
		p.Url()
	}
	{
		p.SetState(494)
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
	p.EnterRule(localctx, 62, FormulaParserRULE_includes)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(496)
		p.Match(FormulaParserINCLUDES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(497)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(498)
		p.FieldExpression()
	}
	{
		p.SetState(499)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(500)
		p.ValueExpression()
	}
	{
		p.SetState(501)
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

// IIsblankContext is an interface to support dynamic dispatch.
type IIsblankContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ISBLANK() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsIsblankContext differentiates from other interfaces.
	IsIsblankContext()
}

type IsblankContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIsblankContext() *IsblankContext {
	var p = new(IsblankContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_isblank
	return p
}

func InitEmptyIsblankContext(p *IsblankContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_isblank
}

func (*IsblankContext) IsIsblankContext() {}

func NewIsblankContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsblankContext {
	var p = new(IsblankContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_isblank

	return p
}

func (s *IsblankContext) GetParser() antlr.Parser { return s.parser }

func (s *IsblankContext) ISBLANK() antlr.TerminalNode {
	return s.GetToken(FormulaParserISBLANK, 0)
}

func (s *IsblankContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *IsblankContext) Expression() IExpressionContext {
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

func (s *IsblankContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *IsblankContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsblankContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsblankContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterIsblank(s)
	}
}

func (s *IsblankContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitIsblank(s)
	}
}

func (s *IsblankContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitIsblank(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Isblank() (localctx IIsblankContext) {
	localctx = NewIsblankContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, FormulaParserRULE_isblank)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(503)
		p.Match(FormulaParserISBLANK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(504)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(505)
		p.expression(0)
	}
	{
		p.SetState(506)
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

// IIschangedContext is an interface to support dynamic dispatch.
type IIschangedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ISCHANGED() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	FieldExpression() IFieldExpressionContext
	RPAREN() antlr.TerminalNode

	// IsIschangedContext differentiates from other interfaces.
	IsIschangedContext()
}

type IschangedContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIschangedContext() *IschangedContext {
	var p = new(IschangedContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_ischanged
	return p
}

func InitEmptyIschangedContext(p *IschangedContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_ischanged
}

func (*IschangedContext) IsIschangedContext() {}

func NewIschangedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IschangedContext {
	var p = new(IschangedContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_ischanged

	return p
}

func (s *IschangedContext) GetParser() antlr.Parser { return s.parser }

func (s *IschangedContext) ISCHANGED() antlr.TerminalNode {
	return s.GetToken(FormulaParserISCHANGED, 0)
}

func (s *IschangedContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *IschangedContext) FieldExpression() IFieldExpressionContext {
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

func (s *IschangedContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *IschangedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IschangedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IschangedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterIschanged(s)
	}
}

func (s *IschangedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitIschanged(s)
	}
}

func (s *IschangedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitIschanged(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Ischanged() (localctx IIschangedContext) {
	localctx = NewIschangedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, FormulaParserRULE_ischanged)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(508)
		p.Match(FormulaParserISCHANGED)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(509)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(510)
		p.FieldExpression()
	}
	{
		p.SetState(511)
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

// IIsnewContext is an interface to support dynamic dispatch.
type IIsnewContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ISNEW() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsIsnewContext differentiates from other interfaces.
	IsIsnewContext()
}

type IsnewContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIsnewContext() *IsnewContext {
	var p = new(IsnewContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_isnew
	return p
}

func InitEmptyIsnewContext(p *IsnewContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_isnew
}

func (*IsnewContext) IsIsnewContext() {}

func NewIsnewContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsnewContext {
	var p = new(IsnewContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_isnew

	return p
}

func (s *IsnewContext) GetParser() antlr.Parser { return s.parser }

func (s *IsnewContext) ISNEW() antlr.TerminalNode {
	return s.GetToken(FormulaParserISNEW, 0)
}

func (s *IsnewContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *IsnewContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *IsnewContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsnewContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsnewContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterIsnew(s)
	}
}

func (s *IsnewContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitIsnew(s)
	}
}

func (s *IsnewContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitIsnew(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Isnew() (localctx IIsnewContext) {
	localctx = NewIsnewContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, FormulaParserRULE_isnew)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(513)
		p.Match(FormulaParserISNEW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(514)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(515)
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

// IIsnullContext is an interface to support dynamic dispatch.
type IIsnullContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ISNULL() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsIsnullContext differentiates from other interfaces.
	IsIsnullContext()
}

type IsnullContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIsnullContext() *IsnullContext {
	var p = new(IsnullContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_isnull
	return p
}

func InitEmptyIsnullContext(p *IsnullContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_isnull
}

func (*IsnullContext) IsIsnullContext() {}

func NewIsnullContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsnullContext {
	var p = new(IsnullContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_isnull

	return p
}

func (s *IsnullContext) GetParser() antlr.Parser { return s.parser }

func (s *IsnullContext) ISNULL() antlr.TerminalNode {
	return s.GetToken(FormulaParserISNULL, 0)
}

func (s *IsnullContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *IsnullContext) Expression() IExpressionContext {
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

func (s *IsnullContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *IsnullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsnullContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsnullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterIsnull(s)
	}
}

func (s *IsnullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitIsnull(s)
	}
}

func (s *IsnullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitIsnull(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Isnull() (localctx IIsnullContext) {
	localctx = NewIsnullContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, FormulaParserRULE_isnull)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(517)
		p.Match(FormulaParserISNULL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(518)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(519)
		p.expression(0)
	}
	{
		p.SetState(520)
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

// IIsnumberContext is an interface to support dynamic dispatch.
type IIsnumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ISNUMBER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsIsnumberContext differentiates from other interfaces.
	IsIsnumberContext()
}

type IsnumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIsnumberContext() *IsnumberContext {
	var p = new(IsnumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_isnumber
	return p
}

func InitEmptyIsnumberContext(p *IsnumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_isnumber
}

func (*IsnumberContext) IsIsnumberContext() {}

func NewIsnumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsnumberContext {
	var p = new(IsnumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_isnumber

	return p
}

func (s *IsnumberContext) GetParser() antlr.Parser { return s.parser }

func (s *IsnumberContext) ISNUMBER() antlr.TerminalNode {
	return s.GetToken(FormulaParserISNUMBER, 0)
}

func (s *IsnumberContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *IsnumberContext) Expression() IExpressionContext {
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

func (s *IsnumberContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *IsnumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsnumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsnumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterIsnumber(s)
	}
}

func (s *IsnumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitIsnumber(s)
	}
}

func (s *IsnumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitIsnumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Isnumber() (localctx IIsnumberContext) {
	localctx = NewIsnumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, FormulaParserRULE_isnumber)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(522)
		p.Match(FormulaParserISNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(523)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(524)
		p.expression(0)
	}
	{
		p.SetState(525)
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

// IIspickvalContext is an interface to support dynamic dispatch.
type IIspickvalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ISPICKVAL() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	FieldExpression() IFieldExpressionContext
	COMMA() antlr.TerminalNode
	ValueExpression() IValueExpressionContext
	RPAREN() antlr.TerminalNode

	// IsIspickvalContext differentiates from other interfaces.
	IsIspickvalContext()
}

type IspickvalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIspickvalContext() *IspickvalContext {
	var p = new(IspickvalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_ispickval
	return p
}

func InitEmptyIspickvalContext(p *IspickvalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_ispickval
}

func (*IspickvalContext) IsIspickvalContext() {}

func NewIspickvalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IspickvalContext {
	var p = new(IspickvalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_ispickval

	return p
}

func (s *IspickvalContext) GetParser() antlr.Parser { return s.parser }

func (s *IspickvalContext) ISPICKVAL() antlr.TerminalNode {
	return s.GetToken(FormulaParserISPICKVAL, 0)
}

func (s *IspickvalContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *IspickvalContext) FieldExpression() IFieldExpressionContext {
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

func (s *IspickvalContext) COMMA() antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, 0)
}

func (s *IspickvalContext) ValueExpression() IValueExpressionContext {
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

func (s *IspickvalContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *IspickvalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IspickvalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IspickvalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterIspickval(s)
	}
}

func (s *IspickvalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitIspickval(s)
	}
}

func (s *IspickvalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitIspickval(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Ispickval() (localctx IIspickvalContext) {
	localctx = NewIspickvalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, FormulaParserRULE_ispickval)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(527)
		p.Match(FormulaParserISPICKVAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(528)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(529)
		p.FieldExpression()
	}
	{
		p.SetState(530)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(531)
		p.ValueExpression()
	}
	{
		p.SetState(532)
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

// ILeftContext is an interface to support dynamic dispatch.
type ILeftContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	TextExpression() ITextExpressionContext
	COMMA() antlr.TerminalNode
	NumChars() INumCharsContext
	RPAREN() antlr.TerminalNode

	// IsLeftContext differentiates from other interfaces.
	IsLeftContext()
}

type LeftContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeftContext() *LeftContext {
	var p = new(LeftContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_left
	return p
}

func InitEmptyLeftContext(p *LeftContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_left
}

func (*LeftContext) IsLeftContext() {}

func NewLeftContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeftContext {
	var p = new(LeftContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_left

	return p
}

func (s *LeftContext) GetParser() antlr.Parser { return s.parser }

func (s *LeftContext) LEFT() antlr.TerminalNode {
	return s.GetToken(FormulaParserLEFT, 0)
}

func (s *LeftContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *LeftContext) TextExpression() ITextExpressionContext {
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

func (s *LeftContext) COMMA() antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, 0)
}

func (s *LeftContext) NumChars() INumCharsContext {
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

func (s *LeftContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *LeftContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeftContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeftContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterLeft(s)
	}
}

func (s *LeftContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitLeft(s)
	}
}

func (s *LeftContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitLeft(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Left() (localctx ILeftContext) {
	localctx = NewLeftContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, FormulaParserRULE_left)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(534)
		p.Match(FormulaParserLEFT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(535)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(536)
		p.TextExpression()
	}
	{
		p.SetState(537)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(538)
		p.NumChars()
	}
	{
		p.SetState(539)
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

// ILenContext is an interface to support dynamic dispatch.
type ILenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEN() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsLenContext differentiates from other interfaces.
	IsLenContext()
}

type LenContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLenContext() *LenContext {
	var p = new(LenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_len
	return p
}

func InitEmptyLenContext(p *LenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_len
}

func (*LenContext) IsLenContext() {}

func NewLenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LenContext {
	var p = new(LenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_len

	return p
}

func (s *LenContext) GetParser() antlr.Parser { return s.parser }

func (s *LenContext) LEN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLEN, 0)
}

func (s *LenContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *LenContext) Expression() IExpressionContext {
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

func (s *LenContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *LenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterLen(s)
	}
}

func (s *LenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitLen(s)
	}
}

func (s *LenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitLen(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Len_() (localctx ILenContext) {
	localctx = NewLenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, FormulaParserRULE_len)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(541)
		p.Match(FormulaParserLEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(542)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(543)
		p.expression(0)
	}
	{
		p.SetState(544)
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

// ILowerContext is an interface to support dynamic dispatch.
type ILowerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LOWER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsLowerContext differentiates from other interfaces.
	IsLowerContext()
}

type LowerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLowerContext() *LowerContext {
	var p = new(LowerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_lower
	return p
}

func InitEmptyLowerContext(p *LowerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_lower
}

func (*LowerContext) IsLowerContext() {}

func NewLowerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LowerContext {
	var p = new(LowerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_lower

	return p
}

func (s *LowerContext) GetParser() antlr.Parser { return s.parser }

func (s *LowerContext) LOWER() antlr.TerminalNode {
	return s.GetToken(FormulaParserLOWER, 0)
}

func (s *LowerContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *LowerContext) Expression() IExpressionContext {
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

func (s *LowerContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *LowerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LowerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LowerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterLower(s)
	}
}

func (s *LowerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitLower(s)
	}
}

func (s *LowerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitLower(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Lower() (localctx ILowerContext) {
	localctx = NewLowerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, FormulaParserRULE_lower)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(546)
		p.Match(FormulaParserLOWER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(547)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(548)
		p.expression(0)
	}
	{
		p.SetState(549)
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

// ILpadContext is an interface to support dynamic dispatch.
type ILpadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAD() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	TextExpression() ITextExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	Length() ILengthContext
	RPAREN() antlr.TerminalNode
	PadString() IPadStringContext

	// IsLpadContext differentiates from other interfaces.
	IsLpadContext()
}

type LpadContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLpadContext() *LpadContext {
	var p = new(LpadContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_lpad
	return p
}

func InitEmptyLpadContext(p *LpadContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_lpad
}

func (*LpadContext) IsLpadContext() {}

func NewLpadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LpadContext {
	var p = new(LpadContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_lpad

	return p
}

func (s *LpadContext) GetParser() antlr.Parser { return s.parser }

func (s *LpadContext) LPAD() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAD, 0)
}

func (s *LpadContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *LpadContext) TextExpression() ITextExpressionContext {
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

func (s *LpadContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FormulaParserCOMMA)
}

func (s *LpadContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, i)
}

func (s *LpadContext) Length() ILengthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILengthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILengthContext)
}

func (s *LpadContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *LpadContext) PadString() IPadStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPadStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPadStringContext)
}

func (s *LpadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LpadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LpadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterLpad(s)
	}
}

func (s *LpadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitLpad(s)
	}
}

func (s *LpadContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitLpad(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Lpad() (localctx ILpadContext) {
	localctx = NewLpadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, FormulaParserRULE_lpad)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(551)
		p.Match(FormulaParserLPAD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(552)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(553)
		p.TextExpression()
	}
	{
		p.SetState(554)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(555)
		p.Length()
	}
	p.SetState(558)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FormulaParserCOMMA {
		{
			p.SetState(556)
			p.Match(FormulaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(557)
			p.PadString()
		}

	}
	{
		p.SetState(560)
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

// IMaxContext is an interface to support dynamic dispatch.
type IMaxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MAX() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsMaxContext differentiates from other interfaces.
	IsMaxContext()
}

type MaxContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMaxContext() *MaxContext {
	var p = new(MaxContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_max
	return p
}

func InitEmptyMaxContext(p *MaxContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_max
}

func (*MaxContext) IsMaxContext() {}

func NewMaxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MaxContext {
	var p = new(MaxContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_max

	return p
}

func (s *MaxContext) GetParser() antlr.Parser { return s.parser }

func (s *MaxContext) MAX() antlr.TerminalNode {
	return s.GetToken(FormulaParserMAX, 0)
}

func (s *MaxContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *MaxContext) AllExpression() []IExpressionContext {
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

func (s *MaxContext) Expression(i int) IExpressionContext {
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

func (s *MaxContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FormulaParserCOMMA)
}

func (s *MaxContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, i)
}

func (s *MaxContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *MaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MaxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MaxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterMax(s)
	}
}

func (s *MaxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitMax(s)
	}
}

func (s *MaxContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitMax(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Max() (localctx IMaxContext) {
	localctx = NewMaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, FormulaParserRULE_max)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(562)
		p.Match(FormulaParserMAX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(563)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(564)
		p.expression(0)
	}
	{
		p.SetState(565)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(566)
		p.expression(0)
	}
	p.SetState(571)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FormulaParserCOMMA {
		{
			p.SetState(567)
			p.Match(FormulaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(568)
			p.expression(0)
		}

		p.SetState(573)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(574)
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
	p.EnterRule(localctx, 86, FormulaParserRULE_mid)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(576)
		p.Match(FormulaParserMID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(577)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(578)
		p.TextExpression()
	}
	{
		p.SetState(579)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(580)
		p.StartNum()
	}
	{
		p.SetState(581)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(582)
		p.NumChars()
	}
	{
		p.SetState(583)
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

// IMinContext is an interface to support dynamic dispatch.
type IMinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MIN() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsMinContext differentiates from other interfaces.
	IsMinContext()
}

type MinContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMinContext() *MinContext {
	var p = new(MinContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_min
	return p
}

func InitEmptyMinContext(p *MinContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_min
}

func (*MinContext) IsMinContext() {}

func NewMinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MinContext {
	var p = new(MinContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_min

	return p
}

func (s *MinContext) GetParser() antlr.Parser { return s.parser }

func (s *MinContext) MIN() antlr.TerminalNode {
	return s.GetToken(FormulaParserMIN, 0)
}

func (s *MinContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *MinContext) AllExpression() []IExpressionContext {
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

func (s *MinContext) Expression(i int) IExpressionContext {
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

func (s *MinContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FormulaParserCOMMA)
}

func (s *MinContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, i)
}

func (s *MinContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *MinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterMin(s)
	}
}

func (s *MinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitMin(s)
	}
}

func (s *MinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitMin(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Min() (localctx IMinContext) {
	localctx = NewMinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, FormulaParserRULE_min)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(585)
		p.Match(FormulaParserMIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(586)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(587)
		p.expression(0)
	}
	{
		p.SetState(588)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(589)
		p.expression(0)
	}
	p.SetState(594)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FormulaParserCOMMA {
		{
			p.SetState(590)
			p.Match(FormulaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(591)
			p.expression(0)
		}

		p.SetState(596)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(597)
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

// IModContext is an interface to support dynamic dispatch.
type IModContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MOD() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Num() INumContext
	COMMA() antlr.TerminalNode
	Divisor() IDivisorContext
	RPAREN() antlr.TerminalNode

	// IsModContext differentiates from other interfaces.
	IsModContext()
}

type ModContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModContext() *ModContext {
	var p = new(ModContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_mod
	return p
}

func InitEmptyModContext(p *ModContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_mod
}

func (*ModContext) IsModContext() {}

func NewModContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModContext {
	var p = new(ModContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_mod

	return p
}

func (s *ModContext) GetParser() antlr.Parser { return s.parser }

func (s *ModContext) MOD() antlr.TerminalNode {
	return s.GetToken(FormulaParserMOD, 0)
}

func (s *ModContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *ModContext) Num() INumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumContext)
}

func (s *ModContext) COMMA() antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, 0)
}

func (s *ModContext) Divisor() IDivisorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDivisorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDivisorContext)
}

func (s *ModContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *ModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterMod(s)
	}
}

func (s *ModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitMod(s)
	}
}

func (s *ModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitMod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Mod() (localctx IModContext) {
	localctx = NewModContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, FormulaParserRULE_mod)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(599)
		p.Match(FormulaParserMOD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(600)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(601)
		p.Num()
	}
	{
		p.SetState(602)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(603)
		p.Divisor()
	}
	{
		p.SetState(604)
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

// IMonthContext is an interface to support dynamic dispatch.
type IMonthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MONTH() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsMonthContext differentiates from other interfaces.
	IsMonthContext()
}

type MonthContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMonthContext() *MonthContext {
	var p = new(MonthContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_month
	return p
}

func InitEmptyMonthContext(p *MonthContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_month
}

func (*MonthContext) IsMonthContext() {}

func NewMonthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MonthContext {
	var p = new(MonthContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_month

	return p
}

func (s *MonthContext) GetParser() antlr.Parser { return s.parser }

func (s *MonthContext) MONTH() antlr.TerminalNode {
	return s.GetToken(FormulaParserMONTH, 0)
}

func (s *MonthContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *MonthContext) Expression() IExpressionContext {
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

func (s *MonthContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *MonthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MonthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterMonth(s)
	}
}

func (s *MonthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitMonth(s)
	}
}

func (s *MonthContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitMonth(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Month() (localctx IMonthContext) {
	localctx = NewMonthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, FormulaParserRULE_month)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(606)
		p.Match(FormulaParserMONTH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(607)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(608)
		p.expression(0)
	}
	{
		p.SetState(609)
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

// INotContext is an interface to support dynamic dispatch.
type INotContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NOT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsNotContext differentiates from other interfaces.
	IsNotContext()
}

type NotContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotContext() *NotContext {
	var p = new(NotContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_not
	return p
}

func InitEmptyNotContext(p *NotContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_not
}

func (*NotContext) IsNotContext() {}

func NewNotContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotContext {
	var p = new(NotContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_not

	return p
}

func (s *NotContext) GetParser() antlr.Parser { return s.parser }

func (s *NotContext) NOT() antlr.TerminalNode {
	return s.GetToken(FormulaParserNOT, 0)
}

func (s *NotContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *NotContext) Expression() IExpressionContext {
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

func (s *NotContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *NotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterNot(s)
	}
}

func (s *NotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitNot(s)
	}
}

func (s *NotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitNot(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Not() (localctx INotContext) {
	localctx = NewNotContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, FormulaParserRULE_not)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(611)
		p.Match(FormulaParserNOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(612)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(613)
		p.expression(0)
	}
	{
		p.SetState(614)
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

// INowContext is an interface to support dynamic dispatch.
type INowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NOW() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsNowContext differentiates from other interfaces.
	IsNowContext()
}

type NowContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNowContext() *NowContext {
	var p = new(NowContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_now
	return p
}

func InitEmptyNowContext(p *NowContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_now
}

func (*NowContext) IsNowContext() {}

func NewNowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NowContext {
	var p = new(NowContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_now

	return p
}

func (s *NowContext) GetParser() antlr.Parser { return s.parser }

func (s *NowContext) NOW() antlr.TerminalNode {
	return s.GetToken(FormulaParserNOW, 0)
}

func (s *NowContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *NowContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *NowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterNow(s)
	}
}

func (s *NowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitNow(s)
	}
}

func (s *NowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitNow(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Now() (localctx INowContext) {
	localctx = NewNowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, FormulaParserRULE_now)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(616)
		p.Match(FormulaParserNOW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(617)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(618)
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

// IOrContext is an interface to support dynamic dispatch.
type IOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OR_FUNC() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	RPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsOrContext differentiates from other interfaces.
	IsOrContext()
}

type OrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrContext() *OrContext {
	var p = new(OrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_or
	return p
}

func InitEmptyOrContext(p *OrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_or
}

func (*OrContext) IsOrContext() {}

func NewOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrContext {
	var p = new(OrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_or

	return p
}

func (s *OrContext) GetParser() antlr.Parser { return s.parser }

func (s *OrContext) OR_FUNC() antlr.TerminalNode {
	return s.GetToken(FormulaParserOR_FUNC, 0)
}

func (s *OrContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *OrContext) AllExpression() []IExpressionContext {
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

func (s *OrContext) Expression(i int) IExpressionContext {
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

func (s *OrContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *OrContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FormulaParserCOMMA)
}

func (s *OrContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, i)
}

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitOr(s)
	}
}

func (s *OrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitOr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Or() (localctx IOrContext) {
	localctx = NewOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, FormulaParserRULE_or)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(620)
		p.Match(FormulaParserOR_FUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(621)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(622)
		p.expression(0)
	}
	p.SetState(627)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FormulaParserCOMMA {
		{
			p.SetState(623)
			p.Match(FormulaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(624)
			p.expression(0)
		}

		p.SetState(629)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(630)
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

// IPriorvalueContext is an interface to support dynamic dispatch.
type IPriorvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRIORVALUE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	FieldExpression() IFieldExpressionContext
	RPAREN() antlr.TerminalNode

	// IsPriorvalueContext differentiates from other interfaces.
	IsPriorvalueContext()
}

type PriorvalueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPriorvalueContext() *PriorvalueContext {
	var p = new(PriorvalueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_priorvalue
	return p
}

func InitEmptyPriorvalueContext(p *PriorvalueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_priorvalue
}

func (*PriorvalueContext) IsPriorvalueContext() {}

func NewPriorvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PriorvalueContext {
	var p = new(PriorvalueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_priorvalue

	return p
}

func (s *PriorvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *PriorvalueContext) PRIORVALUE() antlr.TerminalNode {
	return s.GetToken(FormulaParserPRIORVALUE, 0)
}

func (s *PriorvalueContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *PriorvalueContext) FieldExpression() IFieldExpressionContext {
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

func (s *PriorvalueContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *PriorvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PriorvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PriorvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterPriorvalue(s)
	}
}

func (s *PriorvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitPriorvalue(s)
	}
}

func (s *PriorvalueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitPriorvalue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Priorvalue() (localctx IPriorvalueContext) {
	localctx = NewPriorvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, FormulaParserRULE_priorvalue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(632)
		p.Match(FormulaParserPRIORVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(633)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(634)
		p.FieldExpression()
	}
	{
		p.SetState(635)
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

// IRegexContext is an interface to support dynamic dispatch.
type IRegexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REGEX() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	TextExpression() ITextExpressionContext
	COMMA() antlr.TerminalNode
	RegexExpression() IRegexExpressionContext
	RPAREN() antlr.TerminalNode

	// IsRegexContext differentiates from other interfaces.
	IsRegexContext()
}

type RegexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegexContext() *RegexContext {
	var p = new(RegexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_regex
	return p
}

func InitEmptyRegexContext(p *RegexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_regex
}

func (*RegexContext) IsRegexContext() {}

func NewRegexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegexContext {
	var p = new(RegexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_regex

	return p
}

func (s *RegexContext) GetParser() antlr.Parser { return s.parser }

func (s *RegexContext) REGEX() antlr.TerminalNode {
	return s.GetToken(FormulaParserREGEX, 0)
}

func (s *RegexContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *RegexContext) TextExpression() ITextExpressionContext {
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

func (s *RegexContext) COMMA() antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, 0)
}

func (s *RegexContext) RegexExpression() IRegexExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexExpressionContext)
}

func (s *RegexContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *RegexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterRegex(s)
	}
}

func (s *RegexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitRegex(s)
	}
}

func (s *RegexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitRegex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Regex() (localctx IRegexContext) {
	localctx = NewRegexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, FormulaParserRULE_regex)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(637)
		p.Match(FormulaParserREGEX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(638)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(639)
		p.TextExpression()
	}
	{
		p.SetState(640)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(641)
		p.RegexExpression()
	}
	{
		p.SetState(642)
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

// IRightContext is an interface to support dynamic dispatch.
type IRightContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RIGHT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	TextExpression() ITextExpressionContext
	COMMA() antlr.TerminalNode
	NumChars() INumCharsContext
	RPAREN() antlr.TerminalNode

	// IsRightContext differentiates from other interfaces.
	IsRightContext()
}

type RightContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRightContext() *RightContext {
	var p = new(RightContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_right
	return p
}

func InitEmptyRightContext(p *RightContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_right
}

func (*RightContext) IsRightContext() {}

func NewRightContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RightContext {
	var p = new(RightContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_right

	return p
}

func (s *RightContext) GetParser() antlr.Parser { return s.parser }

func (s *RightContext) RIGHT() antlr.TerminalNode {
	return s.GetToken(FormulaParserRIGHT, 0)
}

func (s *RightContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *RightContext) TextExpression() ITextExpressionContext {
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

func (s *RightContext) COMMA() antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, 0)
}

func (s *RightContext) NumChars() INumCharsContext {
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

func (s *RightContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *RightContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RightContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RightContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterRight(s)
	}
}

func (s *RightContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitRight(s)
	}
}

func (s *RightContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitRight(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Right() (localctx IRightContext) {
	localctx = NewRightContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, FormulaParserRULE_right)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(644)
		p.Match(FormulaParserRIGHT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(645)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(646)
		p.TextExpression()
	}
	{
		p.SetState(647)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(648)
		p.NumChars()
	}
	{
		p.SetState(649)
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

// IRoundContext is an interface to support dynamic dispatch.
type IRoundContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ROUND() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Num() INumContext
	COMMA() antlr.TerminalNode
	Digits() IDigitsContext
	RPAREN() antlr.TerminalNode

	// IsRoundContext differentiates from other interfaces.
	IsRoundContext()
}

type RoundContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRoundContext() *RoundContext {
	var p = new(RoundContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_round
	return p
}

func InitEmptyRoundContext(p *RoundContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_round
}

func (*RoundContext) IsRoundContext() {}

func NewRoundContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RoundContext {
	var p = new(RoundContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_round

	return p
}

func (s *RoundContext) GetParser() antlr.Parser { return s.parser }

func (s *RoundContext) ROUND() antlr.TerminalNode {
	return s.GetToken(FormulaParserROUND, 0)
}

func (s *RoundContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *RoundContext) Num() INumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumContext)
}

func (s *RoundContext) COMMA() antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, 0)
}

func (s *RoundContext) Digits() IDigitsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDigitsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDigitsContext)
}

func (s *RoundContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *RoundContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoundContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RoundContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterRound(s)
	}
}

func (s *RoundContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitRound(s)
	}
}

func (s *RoundContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitRound(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Round() (localctx IRoundContext) {
	localctx = NewRoundContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, FormulaParserRULE_round)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(651)
		p.Match(FormulaParserROUND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(652)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(653)
		p.Num()
	}
	{
		p.SetState(654)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(655)
		p.Digits()
	}
	{
		p.SetState(656)
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

// ISubstituteContext is an interface to support dynamic dispatch.
type ISubstituteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SUBSTITUTE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	TextExpression() ITextExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	OldText() IOldTextContext
	Replacement() IReplacementContext
	RPAREN() antlr.TerminalNode

	// IsSubstituteContext differentiates from other interfaces.
	IsSubstituteContext()
}

type SubstituteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubstituteContext() *SubstituteContext {
	var p = new(SubstituteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_substitute
	return p
}

func InitEmptySubstituteContext(p *SubstituteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_substitute
}

func (*SubstituteContext) IsSubstituteContext() {}

func NewSubstituteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubstituteContext {
	var p = new(SubstituteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_substitute

	return p
}

func (s *SubstituteContext) GetParser() antlr.Parser { return s.parser }

func (s *SubstituteContext) SUBSTITUTE() antlr.TerminalNode {
	return s.GetToken(FormulaParserSUBSTITUTE, 0)
}

func (s *SubstituteContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *SubstituteContext) TextExpression() ITextExpressionContext {
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

func (s *SubstituteContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FormulaParserCOMMA)
}

func (s *SubstituteContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, i)
}

func (s *SubstituteContext) OldText() IOldTextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOldTextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOldTextContext)
}

func (s *SubstituteContext) Replacement() IReplacementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReplacementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReplacementContext)
}

func (s *SubstituteContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *SubstituteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubstituteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubstituteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterSubstitute(s)
	}
}

func (s *SubstituteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitSubstitute(s)
	}
}

func (s *SubstituteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitSubstitute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Substitute() (localctx ISubstituteContext) {
	localctx = NewSubstituteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, FormulaParserRULE_substitute)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(658)
		p.Match(FormulaParserSUBSTITUTE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(659)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(660)
		p.TextExpression()
	}
	{
		p.SetState(661)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(662)
		p.OldText()
	}
	{
		p.SetState(663)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(664)
		p.Replacement()
	}
	{
		p.SetState(665)
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

// ITrimContext is an interface to support dynamic dispatch.
type ITrimContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRIM() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsTrimContext differentiates from other interfaces.
	IsTrimContext()
}

type TrimContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTrimContext() *TrimContext {
	var p = new(TrimContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_trim
	return p
}

func InitEmptyTrimContext(p *TrimContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_trim
}

func (*TrimContext) IsTrimContext() {}

func NewTrimContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TrimContext {
	var p = new(TrimContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_trim

	return p
}

func (s *TrimContext) GetParser() antlr.Parser { return s.parser }

func (s *TrimContext) TRIM() antlr.TerminalNode {
	return s.GetToken(FormulaParserTRIM, 0)
}

func (s *TrimContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *TrimContext) Expression() IExpressionContext {
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

func (s *TrimContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *TrimContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrimContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TrimContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterTrim(s)
	}
}

func (s *TrimContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitTrim(s)
	}
}

func (s *TrimContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitTrim(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Trim() (localctx ITrimContext) {
	localctx = NewTrimContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, FormulaParserRULE_trim)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(667)
		p.Match(FormulaParserTRIM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(668)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(669)
		p.expression(0)
	}
	{
		p.SetState(670)
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
	p.EnterRule(localctx, 112, FormulaParserRULE_text)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(672)
		p.Match(FormulaParserTEXT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(673)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(674)
		p.expression(0)
	}
	{
		p.SetState(675)
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

// ITodayContext is an interface to support dynamic dispatch.
type ITodayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TODAY() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsTodayContext differentiates from other interfaces.
	IsTodayContext()
}

type TodayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTodayContext() *TodayContext {
	var p = new(TodayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_today
	return p
}

func InitEmptyTodayContext(p *TodayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_today
}

func (*TodayContext) IsTodayContext() {}

func NewTodayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TodayContext {
	var p = new(TodayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_today

	return p
}

func (s *TodayContext) GetParser() antlr.Parser { return s.parser }

func (s *TodayContext) TODAY() antlr.TerminalNode {
	return s.GetToken(FormulaParserTODAY, 0)
}

func (s *TodayContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *TodayContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *TodayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TodayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TodayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterToday(s)
	}
}

func (s *TodayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitToday(s)
	}
}

func (s *TodayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitToday(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Today() (localctx ITodayContext) {
	localctx = NewTodayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, FormulaParserRULE_today)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(677)
		p.Match(FormulaParserTODAY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(678)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(679)
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

// IUpperContext is an interface to support dynamic dispatch.
type IUpperContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UPPER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsUpperContext differentiates from other interfaces.
	IsUpperContext()
}

type UpperContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpperContext() *UpperContext {
	var p = new(UpperContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_upper
	return p
}

func InitEmptyUpperContext(p *UpperContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_upper
}

func (*UpperContext) IsUpperContext() {}

func NewUpperContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpperContext {
	var p = new(UpperContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_upper

	return p
}

func (s *UpperContext) GetParser() antlr.Parser { return s.parser }

func (s *UpperContext) UPPER() antlr.TerminalNode {
	return s.GetToken(FormulaParserUPPER, 0)
}

func (s *UpperContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *UpperContext) Expression() IExpressionContext {
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

func (s *UpperContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *UpperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpperContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UpperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterUpper(s)
	}
}

func (s *UpperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitUpper(s)
	}
}

func (s *UpperContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitUpper(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Upper() (localctx IUpperContext) {
	localctx = NewUpperContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, FormulaParserRULE_upper)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(681)
		p.Match(FormulaParserUPPER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(682)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(683)
		p.expression(0)
	}
	{
		p.SetState(684)
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
	p.EnterRule(localctx, 118, FormulaParserRULE_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(686)
		p.Match(FormulaParserVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(687)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(688)
		p.expression(0)
	}
	{
		p.SetState(689)
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

// IYearContext is an interface to support dynamic dispatch.
type IYearContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	YEAR() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsYearContext differentiates from other interfaces.
	IsYearContext()
}

type YearContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyYearContext() *YearContext {
	var p = new(YearContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_year
	return p
}

func InitEmptyYearContext(p *YearContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_year
}

func (*YearContext) IsYearContext() {}

func NewYearContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *YearContext {
	var p = new(YearContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_year

	return p
}

func (s *YearContext) GetParser() antlr.Parser { return s.parser }

func (s *YearContext) YEAR() antlr.TerminalNode {
	return s.GetToken(FormulaParserYEAR, 0)
}

func (s *YearContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserLPAREN, 0)
}

func (s *YearContext) Expression() IExpressionContext {
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

func (s *YearContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserRPAREN, 0)
}

func (s *YearContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *YearContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *YearContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterYear(s)
	}
}

func (s *YearContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitYear(s)
	}
}

func (s *YearContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitYear(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Year() (localctx IYearContext) {
	localctx = NewYearContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, FormulaParserRULE_year)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(691)
		p.Match(FormulaParserYEAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(692)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(693)
		p.expression(0)
	}
	{
		p.SetState(694)
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
	p.EnterRule(localctx, 122, FormulaParserRULE_fieldExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(696)
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

// IDateExpressionContext is an interface to support dynamic dispatch.
type IDateExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsDateExpressionContext differentiates from other interfaces.
	IsDateExpressionContext()
}

type DateExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateExpressionContext() *DateExpressionContext {
	var p = new(DateExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_dateExpression
	return p
}

func InitEmptyDateExpressionContext(p *DateExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_dateExpression
}

func (*DateExpressionContext) IsDateExpressionContext() {}

func NewDateExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateExpressionContext {
	var p = new(DateExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_dateExpression

	return p
}

func (s *DateExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *DateExpressionContext) Expression() IExpressionContext {
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

func (s *DateExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterDateExpression(s)
	}
}

func (s *DateExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitDateExpression(s)
	}
}

func (s *DateExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitDateExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) DateExpression() (localctx IDateExpressionContext) {
	localctx = NewDateExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, FormulaParserRULE_dateExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(698)
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
	p.EnterRule(localctx, 126, FormulaParserRULE_valueExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(700)
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
	p.EnterRule(localctx, 128, FormulaParserRULE_resultExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(702)
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

// ISubstituteValueContext is an interface to support dynamic dispatch.
type ISubstituteValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsSubstituteValueContext differentiates from other interfaces.
	IsSubstituteValueContext()
}

type SubstituteValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubstituteValueContext() *SubstituteValueContext {
	var p = new(SubstituteValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_substituteValue
	return p
}

func InitEmptySubstituteValueContext(p *SubstituteValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_substituteValue
}

func (*SubstituteValueContext) IsSubstituteValueContext() {}

func NewSubstituteValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubstituteValueContext {
	var p = new(SubstituteValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_substituteValue

	return p
}

func (s *SubstituteValueContext) GetParser() antlr.Parser { return s.parser }

func (s *SubstituteValueContext) Expression() IExpressionContext {
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

func (s *SubstituteValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubstituteValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubstituteValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterSubstituteValue(s)
	}
}

func (s *SubstituteValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitSubstituteValue(s)
	}
}

func (s *SubstituteValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitSubstituteValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) SubstituteValue() (localctx ISubstituteValueContext) {
	localctx = NewSubstituteValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, FormulaParserRULE_substituteValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(704)
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
	p.EnterRule(localctx, 132, FormulaParserRULE_defaultExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(706)
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
	p.EnterRule(localctx, 134, FormulaParserRULE_yearExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(708)
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
	p.EnterRule(localctx, 136, FormulaParserRULE_monthExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(710)
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
	p.EnterRule(localctx, 138, FormulaParserRULE_dayExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(712)
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

// IRegexExpressionContext is an interface to support dynamic dispatch.
type IRegexExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsRegexExpressionContext differentiates from other interfaces.
	IsRegexExpressionContext()
}

type RegexExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegexExpressionContext() *RegexExpressionContext {
	var p = new(RegexExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_regexExpression
	return p
}

func InitEmptyRegexExpressionContext(p *RegexExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_regexExpression
}

func (*RegexExpressionContext) IsRegexExpressionContext() {}

func NewRegexExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegexExpressionContext {
	var p = new(RegexExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_regexExpression

	return p
}

func (s *RegexExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RegexExpressionContext) Expression() IExpressionContext {
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

func (s *RegexExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegexExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterRegexExpression(s)
	}
}

func (s *RegexExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitRegexExpression(s)
	}
}

func (s *RegexExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitRegexExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) RegexExpression() (localctx IRegexExpressionContext) {
	localctx = NewRegexExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, FormulaParserRULE_regexExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(714)
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
	p.EnterRule(localctx, 142, FormulaParserRULE_unitExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(716)
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
	p.EnterRule(localctx, 144, FormulaParserRULE_searchExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(718)
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
	p.EnterRule(localctx, 146, FormulaParserRULE_textExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(720)
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

// IOldTextContext is an interface to support dynamic dispatch.
type IOldTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsOldTextContext differentiates from other interfaces.
	IsOldTextContext()
}

type OldTextContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOldTextContext() *OldTextContext {
	var p = new(OldTextContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_oldText
	return p
}

func InitEmptyOldTextContext(p *OldTextContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_oldText
}

func (*OldTextContext) IsOldTextContext() {}

func NewOldTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OldTextContext {
	var p = new(OldTextContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_oldText

	return p
}

func (s *OldTextContext) GetParser() antlr.Parser { return s.parser }

func (s *OldTextContext) Expression() IExpressionContext {
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

func (s *OldTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OldTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OldTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterOldText(s)
	}
}

func (s *OldTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitOldText(s)
	}
}

func (s *OldTextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitOldText(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) OldText() (localctx IOldTextContext) {
	localctx = NewOldTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 148, FormulaParserRULE_oldText)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(722)
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

// IReplacementContext is an interface to support dynamic dispatch.
type IReplacementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsReplacementContext differentiates from other interfaces.
	IsReplacementContext()
}

type ReplacementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReplacementContext() *ReplacementContext {
	var p = new(ReplacementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_replacement
	return p
}

func InitEmptyReplacementContext(p *ReplacementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_replacement
}

func (*ReplacementContext) IsReplacementContext() {}

func NewReplacementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReplacementContext {
	var p = new(ReplacementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_replacement

	return p
}

func (s *ReplacementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReplacementContext) Expression() IExpressionContext {
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

func (s *ReplacementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReplacementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReplacementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterReplacement(s)
	}
}

func (s *ReplacementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitReplacement(s)
	}
}

func (s *ReplacementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitReplacement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Replacement() (localctx IReplacementContext) {
	localctx = NewReplacementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 150, FormulaParserRULE_replacement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(724)
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
	p.EnterRule(localctx, 152, FormulaParserRULE_startNum)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(726)
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

// ICompareTextContext is an interface to support dynamic dispatch.
type ICompareTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsCompareTextContext differentiates from other interfaces.
	IsCompareTextContext()
}

type CompareTextContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompareTextContext() *CompareTextContext {
	var p = new(CompareTextContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_compareText
	return p
}

func InitEmptyCompareTextContext(p *CompareTextContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_compareText
}

func (*CompareTextContext) IsCompareTextContext() {}

func NewCompareTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareTextContext {
	var p = new(CompareTextContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_compareText

	return p
}

func (s *CompareTextContext) GetParser() antlr.Parser { return s.parser }

func (s *CompareTextContext) Expression() IExpressionContext {
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

func (s *CompareTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompareTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterCompareText(s)
	}
}

func (s *CompareTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitCompareText(s)
	}
}

func (s *CompareTextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitCompareText(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) CompareText() (localctx ICompareTextContext) {
	localctx = NewCompareTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 154, FormulaParserRULE_compareText)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(728)
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

// INumContext is an interface to support dynamic dispatch.
type INumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsNumContext differentiates from other interfaces.
	IsNumContext()
}

type NumContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumContext() *NumContext {
	var p = new(NumContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_num
	return p
}

func InitEmptyNumContext(p *NumContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_num
}

func (*NumContext) IsNumContext() {}

func NewNumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumContext {
	var p = new(NumContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_num

	return p
}

func (s *NumContext) GetParser() antlr.Parser { return s.parser }

func (s *NumContext) Expression() IExpressionContext {
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

func (s *NumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterNum(s)
	}
}

func (s *NumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitNum(s)
	}
}

func (s *NumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitNum(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Num() (localctx INumContext) {
	localctx = NewNumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 156, FormulaParserRULE_num)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(730)
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

// IDivisorContext is an interface to support dynamic dispatch.
type IDivisorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsDivisorContext differentiates from other interfaces.
	IsDivisorContext()
}

type DivisorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDivisorContext() *DivisorContext {
	var p = new(DivisorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_divisor
	return p
}

func InitEmptyDivisorContext(p *DivisorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_divisor
}

func (*DivisorContext) IsDivisorContext() {}

func NewDivisorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DivisorContext {
	var p = new(DivisorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_divisor

	return p
}

func (s *DivisorContext) GetParser() antlr.Parser { return s.parser }

func (s *DivisorContext) Expression() IExpressionContext {
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

func (s *DivisorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivisorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DivisorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterDivisor(s)
	}
}

func (s *DivisorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitDivisor(s)
	}
}

func (s *DivisorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitDivisor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Divisor() (localctx IDivisorContext) {
	localctx = NewDivisorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 158, FormulaParserRULE_divisor)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(732)
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

// IDigitsContext is an interface to support dynamic dispatch.
type IDigitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsDigitsContext differentiates from other interfaces.
	IsDigitsContext()
}

type DigitsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDigitsContext() *DigitsContext {
	var p = new(DigitsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_digits
	return p
}

func InitEmptyDigitsContext(p *DigitsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_digits
}

func (*DigitsContext) IsDigitsContext() {}

func NewDigitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DigitsContext {
	var p = new(DigitsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_digits

	return p
}

func (s *DigitsContext) GetParser() antlr.Parser { return s.parser }

func (s *DigitsContext) Expression() IExpressionContext {
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

func (s *DigitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DigitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DigitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterDigits(s)
	}
}

func (s *DigitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitDigits(s)
	}
}

func (s *DigitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitDigits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Digits() (localctx IDigitsContext) {
	localctx = NewDigitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 160, FormulaParserRULE_digits)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(734)
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
	p.EnterRule(localctx, 162, FormulaParserRULE_latitudeExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(736)
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
	p.EnterRule(localctx, 164, FormulaParserRULE_longitudeExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(738)
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

// IUrlContext is an interface to support dynamic dispatch.
type IUrlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsUrlContext differentiates from other interfaces.
	IsUrlContext()
}

type UrlContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUrlContext() *UrlContext {
	var p = new(UrlContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_url
	return p
}

func InitEmptyUrlContext(p *UrlContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_url
}

func (*UrlContext) IsUrlContext() {}

func NewUrlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UrlContext {
	var p = new(UrlContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_url

	return p
}

func (s *UrlContext) GetParser() antlr.Parser { return s.parser }

func (s *UrlContext) Expression() IExpressionContext {
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

func (s *UrlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UrlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UrlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterUrl(s)
	}
}

func (s *UrlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitUrl(s)
	}
}

func (s *UrlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitUrl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Url() (localctx IUrlContext) {
	localctx = NewUrlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 166, FormulaParserRULE_url)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(740)
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

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_name
	return p
}

func InitEmptyNameContext(p *NameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_name
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) Expression() IExpressionContext {
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

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitName(s)
	}
}

func (s *NameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 168, FormulaParserRULE_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(742)
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

// ITargetContext is an interface to support dynamic dispatch.
type ITargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsTargetContext differentiates from other interfaces.
	IsTargetContext()
}

type TargetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetContext() *TargetContext {
	var p = new(TargetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_target
	return p
}

func InitEmptyTargetContext(p *TargetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_target
}

func (*TargetContext) IsTargetContext() {}

func NewTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetContext {
	var p = new(TargetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_target

	return p
}

func (s *TargetContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetContext) Expression() IExpressionContext {
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

func (s *TargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterTarget(s)
	}
}

func (s *TargetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitTarget(s)
	}
}

func (s *TargetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitTarget(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Target() (localctx ITargetContext) {
	localctx = NewTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 170, FormulaParserRULE_target)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(744)
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
	p.EnterRule(localctx, 172, FormulaParserRULE_logicalExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(746)
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
	p.EnterRule(localctx, 174, FormulaParserRULE_ifTrueExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(748)
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
	p.EnterRule(localctx, 176, FormulaParserRULE_ifFalseExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(750)
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

// ILengthContext is an interface to support dynamic dispatch.
type ILengthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsLengthContext differentiates from other interfaces.
	IsLengthContext()
}

type LengthContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLengthContext() *LengthContext {
	var p = new(LengthContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_length
	return p
}

func InitEmptyLengthContext(p *LengthContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_length
}

func (*LengthContext) IsLengthContext() {}

func NewLengthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthContext {
	var p = new(LengthContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_length

	return p
}

func (s *LengthContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthContext) Expression() IExpressionContext {
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

func (s *LengthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LengthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterLength(s)
	}
}

func (s *LengthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitLength(s)
	}
}

func (s *LengthContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitLength(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) Length() (localctx ILengthContext) {
	localctx = NewLengthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 178, FormulaParserRULE_length)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(752)
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

// IPadStringContext is an interface to support dynamic dispatch.
type IPadStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsPadStringContext differentiates from other interfaces.
	IsPadStringContext()
}

type PadStringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPadStringContext() *PadStringContext {
	var p = new(PadStringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_padString
	return p
}

func InitEmptyPadStringContext(p *PadStringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_padString
}

func (*PadStringContext) IsPadStringContext() {}

func NewPadStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PadStringContext {
	var p = new(PadStringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_padString

	return p
}

func (s *PadStringContext) GetParser() antlr.Parser { return s.parser }

func (s *PadStringContext) Expression() IExpressionContext {
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

func (s *PadStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PadStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PadStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.EnterPadString(s)
	}
}

func (s *PadStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaParserListener); ok {
		listenerT.ExitPadString(s)
	}
}

func (s *PadStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FormulaParserVisitor:
		return t.VisitPadString(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FormulaParser) PadString() (localctx IPadStringContext) {
	localctx = NewPadStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 180, FormulaParserRULE_padString)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(754)
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
	p.EnterRule(localctx, 182, FormulaParserRULE_heightExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(756)
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
	p.EnterRule(localctx, 184, FormulaParserRULE_widthExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(758)
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
	p.EnterRule(localctx, 186, FormulaParserRULE_start)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(760)
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
	p.EnterRule(localctx, 188, FormulaParserRULE_numChars)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(762)
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
	p.EnterRule(localctx, 190, FormulaParserRULE_primary)
	p.SetState(769)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FormulaParserStringLiteral, FormulaParserIntegerLiteral, FormulaParserFloatingPointLiteral, FormulaParserBooleanLiteral, FormulaParserNullLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(764)
			p.Literal()
		}

	case FormulaParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(765)
			p.Match(FormulaParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(766)
			p.expression(0)
		}
		{
			p.SetState(767)
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
	p.EnterRule(localctx, 192, FormulaParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(771)
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
