// Code generated from ./FormulaParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // FormulaParser
import "github.com/antlr4-go/antlr/v4"

// FormulaParserListener is a complete listener for a parse tree produced by FormulaParser.
type FormulaParserListener interface {
	antlr.ParseTreeListener

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterCompareExpression is called when entering the compareExpression production.
	EnterCompareExpression(c *CompareExpressionContext)

	// EnterEqualityExpression is called when entering the equalityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterLogicExpression is called when entering the logicExpression production.
	EnterLogicExpression(c *LogicExpressionContext)

	// EnterFunctionCallExpression is called when entering the functionCallExpression production.
	EnterFunctionCallExpression(c *FunctionCallExpressionContext)

	// EnterFieldReference is called when entering the fieldReference production.
	EnterFieldReference(c *FieldReferenceContext)

	// EnterExponentiationExpression is called when entering the exponentiationExpression production.
	EnterExponentiationExpression(c *ExponentiationExpressionContext)

	// EnterArithExpression is called when entering the arithExpression production.
	EnterArithExpression(c *ArithExpressionContext)

	// EnterConcatExpression is called when entering the concatExpression production.
	EnterConcatExpression(c *ConcatExpressionContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterAbs is called when entering the abs production.
	EnterAbs(c *AbsContext)

	// EnterAddMonths is called when entering the addMonths production.
	EnterAddMonths(c *AddMonthsContext)

	// EnterAnd is called when entering the and production.
	EnterAnd(c *AndContext)

	// EnterBegins is called when entering the begins production.
	EnterBegins(c *BeginsContext)

	// EnterBlankvalue is called when entering the blankvalue production.
	EnterBlankvalue(c *BlankvalueContext)

	// EnterBr is called when entering the br production.
	EnterBr(c *BrContext)

	// EnterCase is called when entering the case production.
	EnterCase(c *CaseContext)

	// EnterCasesafeid is called when entering the casesafeid production.
	EnterCasesafeid(c *CasesafeidContext)

	// EnterCeiling is called when entering the ceiling production.
	EnterCeiling(c *CeilingContext)

	// EnterContains is called when entering the contains production.
	EnterContains(c *ContainsContext)

	// EnterCurrencyrate is called when entering the currencyrate production.
	EnterCurrencyrate(c *CurrencyrateContext)

	// EnterDate is called when entering the date production.
	EnterDate(c *DateContext)

	// EnterDatevalue is called when entering the datevalue production.
	EnterDatevalue(c *DatevalueContext)

	// EnterDatetimevalue is called when entering the datetimevalue production.
	EnterDatetimevalue(c *DatetimevalueContext)

	// EnterDay is called when entering the day production.
	EnterDay(c *DayContext)

	// EnterValueExpression is called when entering the valueExpression production.
	EnterValueExpression(c *ValueExpressionContext)

	// EnterResultExpression is called when entering the resultExpression production.
	EnterResultExpression(c *ResultExpressionContext)

	// EnterDefaultExpression is called when entering the defaultExpression production.
	EnterDefaultExpression(c *DefaultExpressionContext)

	// EnterYearExpression is called when entering the yearExpression production.
	EnterYearExpression(c *YearExpressionContext)

	// EnterMonthExpression is called when entering the monthExpression production.
	EnterMonthExpression(c *MonthExpressionContext)

	// EnterDayExpression is called when entering the dayExpression production.
	EnterDayExpression(c *DayExpressionContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitCompareExpression is called when exiting the compareExpression production.
	ExitCompareExpression(c *CompareExpressionContext)

	// ExitEqualityExpression is called when exiting the equalityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitLogicExpression is called when exiting the logicExpression production.
	ExitLogicExpression(c *LogicExpressionContext)

	// ExitFunctionCallExpression is called when exiting the functionCallExpression production.
	ExitFunctionCallExpression(c *FunctionCallExpressionContext)

	// ExitFieldReference is called when exiting the fieldReference production.
	ExitFieldReference(c *FieldReferenceContext)

	// ExitExponentiationExpression is called when exiting the exponentiationExpression production.
	ExitExponentiationExpression(c *ExponentiationExpressionContext)

	// ExitArithExpression is called when exiting the arithExpression production.
	ExitArithExpression(c *ArithExpressionContext)

	// ExitConcatExpression is called when exiting the concatExpression production.
	ExitConcatExpression(c *ConcatExpressionContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitAbs is called when exiting the abs production.
	ExitAbs(c *AbsContext)

	// ExitAddMonths is called when exiting the addMonths production.
	ExitAddMonths(c *AddMonthsContext)

	// ExitAnd is called when exiting the and production.
	ExitAnd(c *AndContext)

	// ExitBegins is called when exiting the begins production.
	ExitBegins(c *BeginsContext)

	// ExitBlankvalue is called when exiting the blankvalue production.
	ExitBlankvalue(c *BlankvalueContext)

	// ExitBr is called when exiting the br production.
	ExitBr(c *BrContext)

	// ExitCase is called when exiting the case production.
	ExitCase(c *CaseContext)

	// ExitCasesafeid is called when exiting the casesafeid production.
	ExitCasesafeid(c *CasesafeidContext)

	// ExitCeiling is called when exiting the ceiling production.
	ExitCeiling(c *CeilingContext)

	// ExitContains is called when exiting the contains production.
	ExitContains(c *ContainsContext)

	// ExitCurrencyrate is called when exiting the currencyrate production.
	ExitCurrencyrate(c *CurrencyrateContext)

	// ExitDate is called when exiting the date production.
	ExitDate(c *DateContext)

	// ExitDatevalue is called when exiting the datevalue production.
	ExitDatevalue(c *DatevalueContext)

	// ExitDatetimevalue is called when exiting the datetimevalue production.
	ExitDatetimevalue(c *DatetimevalueContext)

	// ExitDay is called when exiting the day production.
	ExitDay(c *DayContext)

	// ExitValueExpression is called when exiting the valueExpression production.
	ExitValueExpression(c *ValueExpressionContext)

	// ExitResultExpression is called when exiting the resultExpression production.
	ExitResultExpression(c *ResultExpressionContext)

	// ExitDefaultExpression is called when exiting the defaultExpression production.
	ExitDefaultExpression(c *DefaultExpressionContext)

	// ExitYearExpression is called when exiting the yearExpression production.
	ExitYearExpression(c *YearExpressionContext)

	// ExitMonthExpression is called when exiting the monthExpression production.
	ExitMonthExpression(c *MonthExpressionContext)

	// ExitDayExpression is called when exiting the dayExpression production.
	ExitDayExpression(c *DayExpressionContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)
}
