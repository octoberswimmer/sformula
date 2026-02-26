package parser

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

type countingErrorListener struct {
	*antlr.DefaultErrorListener
	count int
}

func (l *countingErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.count++
}

func parseFormula(t *testing.T, formula string) {
	t.Helper()
	input := antlr.NewInputStream(formula)
	lexer := NewFormulaLexer(input)
	lexerErrors := &countingErrorListener{DefaultErrorListener: &antlr.DefaultErrorListener{}}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewFormulaParser(stream)
	parserErrors := &countingErrorListener{DefaultErrorListener: &antlr.DefaultErrorListener{}}
	p.RemoveErrorListeners()
	p.AddErrorListener(parserErrors)

	p.CompilationUnit()

	if lexerErrors.count != 0 || parserErrors.count != 0 {
		t.Fatalf("formula parse errors (lexer=%d, parser=%d) for %q", lexerErrors.count, parserErrors.count, formula)
	}
}

func TestStringLiteralsInFormulas(t *testing.T) {
	tests := []string{
		`"Simple" & "String"`,
		`'Simple' & 'String'`,
		`'first' & "second"`,
	}

	for _, formula := range tests {
		t.Run(formula, func(t *testing.T) {
			parseFormula(t, formula)
		})
	}
}

func TestClearFormulaParserDFA(t *testing.T) {
	// Parse a formula to populate the DFA cache
	parseFormula(t, `IF(1 > 0, "yes", "no")`)

	// Clear should not panic
	ClearFormulaParserDFA()

	// Parsing should still work after clearing
	parseFormula(t, `1 + 2 * 3`)

	// Multiple clears should be safe
	ClearFormulaParserDFA()
	ClearFormulaParserDFA()

	parseFormula(t, `TEXT(123)`)
}

func TestClearFormulaLexerDFA(t *testing.T) {
	// Parse a formula to populate the DFA cache
	parseFormula(t, `"string" & 'another'`)

	// Clear should not panic
	ClearFormulaLexerDFA()

	// Parsing should still work after clearing
	parseFormula(t, `123.456 + 789`)

	// Multiple clears should be safe
	ClearFormulaLexerDFA()
	ClearFormulaLexerDFA()

	parseFormula(t, `Account.Name`)
}

func TestTrigFunctions(t *testing.T) {
	tests := []string{
		`ACOS(0.5)`,
		`ASIN(0.5)`,
		`ATAN(1)`,
		`ATAN2(1, 2)`,
		`COS(0)`,
		`SIN(0)`,
		`TAN(0)`,
		`PI()`,
	}

	for _, formula := range tests {
		t.Run(formula, func(t *testing.T) {
			parseFormula(t, formula)
		})
	}
}

func TestStringFunctions(t *testing.T) {
	tests := []string{
		`ASCII("A")`,
		`CHR(65)`,
		`INITCAP("hello world")`,
	}

	for _, formula := range tests {
		t.Run(formula, func(t *testing.T) {
			parseFormula(t, formula)
		})
	}
}

func TestDateTimeFunctions(t *testing.T) {
	tests := []string{
		`DAYOFYEAR(TODAY())`,
		`ISOWEEK(TODAY())`,
		`ISOYEAR(TODAY())`,
		`UNIXTIMESTAMP(NOW())`,
		`FROMUNIXTIME(1234567890)`,
	}

	for _, formula := range tests {
		t.Run(formula, func(t *testing.T) {
			parseFormula(t, formula)
		})
	}
}

func TestNumericFunctions(t *testing.T) {
	tests := []string{
		`TRUNC(3.14159, 2)`,
		`TRUNC(123.456, 0)`,
	}

	for _, formula := range tests {
		t.Run(formula, func(t *testing.T) {
			parseFormula(t, formula)
		})
	}
}

func TestPicklistFunctions(t *testing.T) {
	tests := []string{
		`PICKLISTCOUNT(Field__c)`,
	}

	for _, formula := range tests {
		t.Run(formula, func(t *testing.T) {
			parseFormula(t, formula)
		})
	}
}

func TestFieldReferenceWithValueKeyword(t *testing.T) {
	tests := []string{
		`value.Field__c`,
		`value.Payment_Method__c`,
		`Object.value`,
		`value.Nested.Field`,
	}

	for _, formula := range tests {
		t.Run(formula, func(t *testing.T) {
			parseFormula(t, formula)
		})
	}
}
