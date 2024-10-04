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
		"valueExpression", "resultExpression", "defaultExpression", "yearExpression",
		"monthExpression", "dayExpression", "primary", "arguments", "literal",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 111, 252, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 63, 8, 1,
		10, 1, 12, 1, 66, 9, 1, 3, 1, 68, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 82, 8, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 91, 8, 1, 10, 1, 12, 1, 94, 9, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 3, 2, 111, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5,
		5, 132, 8, 5, 10, 5, 12, 5, 135, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 5, 9, 169, 8, 9, 10, 9, 12, 9, 172, 9, 9, 1, 9, 1, 9, 1,
		9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 3, 24, 240, 8, 24, 1, 25, 1, 25, 1, 25, 5, 25, 245, 8, 25, 10, 25,
		12, 25, 248, 9, 25, 1, 26, 1, 26, 1, 26, 0, 1, 2, 27, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
		48, 50, 52, 0, 5, 1, 0, 98, 101, 1, 0, 102, 103, 1, 0, 105, 107, 1, 0,
		108, 109, 1, 0, 86, 90, 252, 0, 54, 1, 0, 0, 0, 2, 67, 1, 0, 0, 0, 4, 110,
		1, 0, 0, 0, 6, 112, 1, 0, 0, 0, 8, 117, 1, 0, 0, 0, 10, 124, 1, 0, 0, 0,
		12, 138, 1, 0, 0, 0, 14, 145, 1, 0, 0, 0, 16, 152, 1, 0, 0, 0, 18, 156,
		1, 0, 0, 0, 20, 176, 1, 0, 0, 0, 22, 181, 1, 0, 0, 0, 24, 186, 1, 0, 0,
		0, 26, 193, 1, 0, 0, 0, 28, 198, 1, 0, 0, 0, 30, 207, 1, 0, 0, 0, 32, 212,
		1, 0, 0, 0, 34, 217, 1, 0, 0, 0, 36, 222, 1, 0, 0, 0, 38, 224, 1, 0, 0,
		0, 40, 226, 1, 0, 0, 0, 42, 228, 1, 0, 0, 0, 44, 230, 1, 0, 0, 0, 46, 232,
		1, 0, 0, 0, 48, 239, 1, 0, 0, 0, 50, 241, 1, 0, 0, 0, 52, 249, 1, 0, 0,
		0, 54, 55, 3, 2, 1, 0, 55, 1, 1, 0, 0, 0, 56, 57, 6, 1, -1, 0, 57, 68,
		3, 48, 24, 0, 58, 68, 3, 4, 2, 0, 59, 64, 5, 91, 0, 0, 60, 61, 5, 94, 0,
		0, 61, 63, 5, 91, 0, 0, 62, 60, 1, 0, 0, 0, 63, 66, 1, 0, 0, 0, 64, 62,
		1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 68, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0,
		67, 56, 1, 0, 0, 0, 67, 58, 1, 0, 0, 0, 67, 59, 1, 0, 0, 0, 68, 92, 1,
		0, 0, 0, 69, 70, 10, 6, 0, 0, 70, 71, 5, 96, 0, 0, 71, 91, 3, 2, 1, 7,
		72, 73, 10, 5, 0, 0, 73, 74, 5, 97, 0, 0, 74, 91, 3, 2, 1, 6, 75, 76, 10,
		4, 0, 0, 76, 77, 7, 0, 0, 0, 77, 91, 3, 2, 1, 5, 78, 79, 10, 3, 0, 0, 79,
		81, 7, 1, 0, 0, 80, 82, 5, 104, 0, 0, 81, 80, 1, 0, 0, 0, 81, 82, 1, 0,
		0, 0, 82, 83, 1, 0, 0, 0, 83, 91, 3, 2, 1, 4, 84, 85, 10, 2, 0, 0, 85,
		86, 7, 2, 0, 0, 86, 91, 3, 2, 1, 3, 87, 88, 10, 1, 0, 0, 88, 89, 7, 3,
		0, 0, 89, 91, 3, 2, 1, 2, 90, 69, 1, 0, 0, 0, 90, 72, 1, 0, 0, 0, 90, 75,
		1, 0, 0, 0, 90, 78, 1, 0, 0, 0, 90, 84, 1, 0, 0, 0, 90, 87, 1, 0, 0, 0,
		91, 94, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 3, 1, 0,
		0, 0, 94, 92, 1, 0, 0, 0, 95, 111, 3, 6, 3, 0, 96, 111, 3, 8, 4, 0, 97,
		111, 3, 10, 5, 0, 98, 111, 3, 12, 6, 0, 99, 111, 3, 14, 7, 0, 100, 111,
		3, 16, 8, 0, 101, 111, 3, 18, 9, 0, 102, 111, 3, 20, 10, 0, 103, 111, 3,
		22, 11, 0, 104, 111, 3, 24, 12, 0, 105, 111, 3, 26, 13, 0, 106, 111, 3,
		28, 14, 0, 107, 111, 3, 30, 15, 0, 108, 111, 3, 32, 16, 0, 109, 111, 3,
		34, 17, 0, 110, 95, 1, 0, 0, 0, 110, 96, 1, 0, 0, 0, 110, 97, 1, 0, 0,
		0, 110, 98, 1, 0, 0, 0, 110, 99, 1, 0, 0, 0, 110, 100, 1, 0, 0, 0, 110,
		101, 1, 0, 0, 0, 110, 102, 1, 0, 0, 0, 110, 103, 1, 0, 0, 0, 110, 104,
		1, 0, 0, 0, 110, 105, 1, 0, 0, 0, 110, 106, 1, 0, 0, 0, 110, 107, 1, 0,
		0, 0, 110, 108, 1, 0, 0, 0, 110, 109, 1, 0, 0, 0, 111, 5, 1, 0, 0, 0, 112,
		113, 5, 1, 0, 0, 113, 114, 5, 92, 0, 0, 114, 115, 3, 2, 1, 0, 115, 116,
		5, 93, 0, 0, 116, 7, 1, 0, 0, 0, 117, 118, 5, 2, 0, 0, 118, 119, 5, 92,
		0, 0, 119, 120, 3, 2, 1, 0, 120, 121, 5, 95, 0, 0, 121, 122, 3, 2, 1, 0,
		122, 123, 5, 93, 0, 0, 123, 9, 1, 0, 0, 0, 124, 125, 5, 3, 0, 0, 125, 126,
		5, 92, 0, 0, 126, 127, 3, 2, 1, 0, 127, 128, 5, 95, 0, 0, 128, 133, 3,
		2, 1, 0, 129, 130, 5, 95, 0, 0, 130, 132, 3, 2, 1, 0, 131, 129, 1, 0, 0,
		0, 132, 135, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134,
		136, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 136, 137, 5, 93, 0, 0, 137, 11,
		1, 0, 0, 0, 138, 139, 5, 4, 0, 0, 139, 140, 5, 92, 0, 0, 140, 141, 3, 2,
		1, 0, 141, 142, 5, 95, 0, 0, 142, 143, 3, 2, 1, 0, 143, 144, 5, 93, 0,
		0, 144, 13, 1, 0, 0, 0, 145, 146, 5, 5, 0, 0, 146, 147, 5, 92, 0, 0, 147,
		148, 3, 2, 1, 0, 148, 149, 5, 95, 0, 0, 149, 150, 3, 2, 1, 0, 150, 151,
		5, 93, 0, 0, 151, 15, 1, 0, 0, 0, 152, 153, 5, 6, 0, 0, 153, 154, 5, 92,
		0, 0, 154, 155, 5, 93, 0, 0, 155, 17, 1, 0, 0, 0, 156, 157, 5, 7, 0, 0,
		157, 158, 5, 92, 0, 0, 158, 159, 3, 2, 1, 0, 159, 160, 5, 95, 0, 0, 160,
		161, 3, 36, 18, 0, 161, 162, 5, 95, 0, 0, 162, 170, 3, 38, 19, 0, 163,
		164, 5, 95, 0, 0, 164, 165, 3, 36, 18, 0, 165, 166, 5, 95, 0, 0, 166, 167,
		3, 38, 19, 0, 167, 169, 1, 0, 0, 0, 168, 163, 1, 0, 0, 0, 169, 172, 1,
		0, 0, 0, 170, 168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 173, 1, 0, 0,
		0, 172, 170, 1, 0, 0, 0, 173, 174, 3, 40, 20, 0, 174, 175, 5, 93, 0, 0,
		175, 19, 1, 0, 0, 0, 176, 177, 5, 8, 0, 0, 177, 178, 5, 92, 0, 0, 178,
		179, 3, 2, 1, 0, 179, 180, 5, 93, 0, 0, 180, 21, 1, 0, 0, 0, 181, 182,
		5, 9, 0, 0, 182, 183, 5, 92, 0, 0, 183, 184, 3, 2, 1, 0, 184, 185, 5, 93,
		0, 0, 185, 23, 1, 0, 0, 0, 186, 187, 5, 10, 0, 0, 187, 188, 5, 92, 0, 0,
		188, 189, 3, 2, 1, 0, 189, 190, 5, 95, 0, 0, 190, 191, 3, 2, 1, 0, 191,
		192, 5, 93, 0, 0, 192, 25, 1, 0, 0, 0, 193, 194, 5, 11, 0, 0, 194, 195,
		5, 92, 0, 0, 195, 196, 3, 2, 1, 0, 196, 197, 5, 93, 0, 0, 197, 27, 1, 0,
		0, 0, 198, 199, 5, 12, 0, 0, 199, 200, 5, 92, 0, 0, 200, 201, 3, 42, 21,
		0, 201, 202, 5, 95, 0, 0, 202, 203, 3, 44, 22, 0, 203, 204, 5, 95, 0, 0,
		204, 205, 3, 46, 23, 0, 205, 206, 5, 93, 0, 0, 206, 29, 1, 0, 0, 0, 207,
		208, 5, 13, 0, 0, 208, 209, 5, 92, 0, 0, 209, 210, 3, 2, 1, 0, 210, 211,
		5, 93, 0, 0, 211, 31, 1, 0, 0, 0, 212, 213, 5, 14, 0, 0, 213, 214, 5, 92,
		0, 0, 214, 215, 3, 2, 1, 0, 215, 216, 5, 93, 0, 0, 216, 33, 1, 0, 0, 0,
		217, 218, 5, 15, 0, 0, 218, 219, 5, 92, 0, 0, 219, 220, 3, 2, 1, 0, 220,
		221, 5, 93, 0, 0, 221, 35, 1, 0, 0, 0, 222, 223, 3, 2, 1, 0, 223, 37, 1,
		0, 0, 0, 224, 225, 3, 2, 1, 0, 225, 39, 1, 0, 0, 0, 226, 227, 3, 2, 1,
		0, 227, 41, 1, 0, 0, 0, 228, 229, 3, 2, 1, 0, 229, 43, 1, 0, 0, 0, 230,
		231, 3, 2, 1, 0, 231, 45, 1, 0, 0, 0, 232, 233, 3, 2, 1, 0, 233, 47, 1,
		0, 0, 0, 234, 240, 3, 52, 26, 0, 235, 236, 5, 92, 0, 0, 236, 237, 3, 2,
		1, 0, 237, 238, 5, 93, 0, 0, 238, 240, 1, 0, 0, 0, 239, 234, 1, 0, 0, 0,
		239, 235, 1, 0, 0, 0, 240, 49, 1, 0, 0, 0, 241, 246, 3, 2, 1, 0, 242, 243,
		5, 95, 0, 0, 243, 245, 3, 2, 1, 0, 244, 242, 1, 0, 0, 0, 245, 248, 1, 0,
		0, 0, 246, 244, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 51, 1, 0, 0, 0,
		248, 246, 1, 0, 0, 0, 249, 250, 7, 4, 0, 0, 250, 53, 1, 0, 0, 0, 10, 64,
		67, 81, 90, 92, 110, 133, 170, 239, 246,
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
	FormulaParserRULE_compilationUnit   = 0
	FormulaParserRULE_expression        = 1
	FormulaParserRULE_functionCall      = 2
	FormulaParserRULE_abs               = 3
	FormulaParserRULE_addMonths         = 4
	FormulaParserRULE_and               = 5
	FormulaParserRULE_begins            = 6
	FormulaParserRULE_blankvalue        = 7
	FormulaParserRULE_br                = 8
	FormulaParserRULE_case              = 9
	FormulaParserRULE_casesafeid        = 10
	FormulaParserRULE_ceiling           = 11
	FormulaParserRULE_contains          = 12
	FormulaParserRULE_currencyrate      = 13
	FormulaParserRULE_date              = 14
	FormulaParserRULE_datevalue         = 15
	FormulaParserRULE_datetimevalue     = 16
	FormulaParserRULE_day               = 17
	FormulaParserRULE_valueExpression   = 18
	FormulaParserRULE_resultExpression  = 19
	FormulaParserRULE_defaultExpression = 20
	FormulaParserRULE_yearExpression    = 21
	FormulaParserRULE_monthExpression   = 22
	FormulaParserRULE_dayExpression     = 23
	FormulaParserRULE_primary           = 24
	FormulaParserRULE_arguments         = 25
	FormulaParserRULE_literal           = 26
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
		p.SetState(54)
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
	p.SetState(67)
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
			p.SetState(57)
			p.Primary()
		}

	case FormulaParserABS, FormulaParserADDMONTHS, FormulaParserAND_FUNC, FormulaParserBEGINS, FormulaParserBLANKVALUE, FormulaParserBR, FormulaParserCASE, FormulaParserCASESAFEID, FormulaParserCEILING, FormulaParserCONTAINS, FormulaParserCURRENCYRATE, FormulaParserDATE, FormulaParserDATEVALUE, FormulaParserDATETIMEVALUE, FormulaParserDAY:
		localctx = NewFunctionCallExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(58)
			p.FunctionCall()
		}

	case FormulaParserIdentifier:
		localctx = NewFieldReferenceContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(59)
			p.Match(FormulaParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(64)
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
					p.SetState(60)
					p.Match(FormulaParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(61)
					p.Match(FormulaParserIdentifier)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(66)
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
	p.SetState(92)
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
			p.SetState(90)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExponentiationExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expression)
				p.SetState(69)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(70)
					p.Match(FormulaParserCARET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(71)
					p.expression(7)
				}

			case 2:
				localctx = NewConcatExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expression)
				p.SetState(72)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(73)
					p.Match(FormulaParserBITAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(74)
					p.expression(6)
				}

			case 3:
				localctx = NewArithExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expression)
				p.SetState(75)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(76)
					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-98)) & ^0x3f) == 0 && ((int64(1)<<(_la-98))&15) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(77)
					p.expression(5)
				}

			case 4:
				localctx = NewCompareExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expression)
				p.SetState(78)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(79)
					_la = p.GetTokenStream().LA(1)

					if !(_la == FormulaParserLT || _la == FormulaParserGT) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(81)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == FormulaParserASSIGN {
					{
						p.SetState(80)
						p.Match(FormulaParserASSIGN)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(83)
					p.expression(4)
				}

			case 5:
				localctx = NewEqualityExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expression)
				p.SetState(84)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(85)
					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-105)) & ^0x3f) == 0 && ((int64(1)<<(_la-105))&7) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(86)
					p.expression(3)
				}

			case 6:
				localctx = NewLogicExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expression)
				p.SetState(87)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(88)
					_la = p.GetTokenStream().LA(1)

					if !(_la == FormulaParserAND || _la == FormulaParserOR) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(89)
					p.expression(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(94)
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
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FormulaParserABS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(95)
			p.Abs()
		}

	case FormulaParserADDMONTHS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(96)
			p.AddMonths()
		}

	case FormulaParserAND_FUNC:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(97)
			p.And()
		}

	case FormulaParserBEGINS:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(98)
			p.Begins()
		}

	case FormulaParserBLANKVALUE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(99)
			p.Blankvalue()
		}

	case FormulaParserBR:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(100)
			p.Br()
		}

	case FormulaParserCASE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(101)
			p.Case_()
		}

	case FormulaParserCASESAFEID:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(102)
			p.Casesafeid()
		}

	case FormulaParserCEILING:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(103)
			p.Ceiling()
		}

	case FormulaParserCONTAINS:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(104)
			p.Contains()
		}

	case FormulaParserCURRENCYRATE:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(105)
			p.Currencyrate()
		}

	case FormulaParserDATE:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(106)
			p.Date()
		}

	case FormulaParserDATEVALUE:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(107)
			p.Datevalue()
		}

	case FormulaParserDATETIMEVALUE:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(108)
			p.Datetimevalue()
		}

	case FormulaParserDAY:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(109)
			p.Day()
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
		p.SetState(112)
		p.Match(FormulaParserABS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)
		p.expression(0)
	}
	{
		p.SetState(115)
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
		p.SetState(117)
		p.Match(FormulaParserADDMONTHS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.expression(0)
	}
	{
		p.SetState(120)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.expression(0)
	}
	{
		p.SetState(122)
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
		p.SetState(124)
		p.Match(FormulaParserAND_FUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
		p.expression(0)
	}
	{
		p.SetState(127)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.expression(0)
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FormulaParserCOMMA {
		{
			p.SetState(129)
			p.Match(FormulaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.expression(0)
		}

		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(136)
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
		p.SetState(138)
		p.Match(FormulaParserBEGINS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.expression(0)
	}
	{
		p.SetState(141)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		p.expression(0)
	}
	{
		p.SetState(143)
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
		p.SetState(145)
		p.Match(FormulaParserBLANKVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
		p.expression(0)
	}
	{
		p.SetState(148)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.expression(0)
	}
	{
		p.SetState(150)
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
		p.SetState(152)
		p.Match(FormulaParserBR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
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
		p.SetState(156)
		p.Match(FormulaParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.expression(0)
	}
	{
		p.SetState(159)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)
		p.ValueExpression()
	}
	{
		p.SetState(161)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.ResultExpression()
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FormulaParserCOMMA {
		{
			p.SetState(163)
			p.Match(FormulaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)
			p.ValueExpression()
		}
		{
			p.SetState(165)
			p.Match(FormulaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.ResultExpression()
		}

		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(173)
		p.DefaultExpression()
	}
	{
		p.SetState(174)
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
		p.SetState(176)
		p.Match(FormulaParserCASESAFEID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(178)
		p.expression(0)
	}
	{
		p.SetState(179)
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
		p.SetState(181)
		p.Match(FormulaParserCEILING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(183)
		p.expression(0)
	}
	{
		p.SetState(184)
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
		p.SetState(186)
		p.Match(FormulaParserCONTAINS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(188)
		p.expression(0)
	}
	{
		p.SetState(189)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(190)
		p.expression(0)
	}
	{
		p.SetState(191)
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
		p.SetState(193)
		p.Match(FormulaParserCURRENCYRATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(194)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(195)
		p.expression(0)
	}
	{
		p.SetState(196)
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
		p.SetState(198)
		p.Match(FormulaParserDATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)
		p.YearExpression()
	}
	{
		p.SetState(201)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.MonthExpression()
	}
	{
		p.SetState(203)
		p.Match(FormulaParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
		p.DayExpression()
	}
	{
		p.SetState(205)
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
		p.SetState(207)
		p.Match(FormulaParserDATEVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.expression(0)
	}
	{
		p.SetState(210)
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
		p.SetState(212)
		p.Match(FormulaParserDATETIMEVALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
		p.expression(0)
	}
	{
		p.SetState(215)
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
		p.SetState(217)
		p.Match(FormulaParserDAY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(218)
		p.Match(FormulaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(219)
		p.expression(0)
	}
	{
		p.SetState(220)
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
	p.EnterRule(localctx, 36, FormulaParserRULE_valueExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
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
	p.EnterRule(localctx, 38, FormulaParserRULE_resultExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
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
	p.EnterRule(localctx, 40, FormulaParserRULE_defaultExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
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
	p.EnterRule(localctx, 42, FormulaParserRULE_yearExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
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
	p.EnterRule(localctx, 44, FormulaParserRULE_monthExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
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
	p.EnterRule(localctx, 46, FormulaParserRULE_dayExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
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
	p.EnterRule(localctx, 48, FormulaParserRULE_primary)
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FormulaParserStringLiteral, FormulaParserIntegerLiteral, FormulaParserFloatingPointLiteral, FormulaParserBooleanLiteral, FormulaParserNullLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(234)
			p.Literal()
		}

	case FormulaParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(235)
			p.Match(FormulaParserLPAREN)
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
	p.EnterRule(localctx, 50, FormulaParserRULE_arguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.expression(0)
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FormulaParserCOMMA {
		{
			p.SetState(242)
			p.Match(FormulaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.expression(0)
		}

		p.SetState(248)
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
	p.EnterRule(localctx, 52, FormulaParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
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
