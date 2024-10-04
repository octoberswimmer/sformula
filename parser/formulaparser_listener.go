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

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)
}
