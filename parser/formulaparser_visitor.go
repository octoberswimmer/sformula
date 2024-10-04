// Code generated from ./FormulaParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // FormulaParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by FormulaParser.
type FormulaParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FormulaParser#compilationUnit.
	VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

	// Visit a parse tree produced by FormulaParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#compareExpression.
	VisitCompareExpression(ctx *CompareExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#logicExpression.
	VisitLogicExpression(ctx *LogicExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#functionCallExpression.
	VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#fieldReference.
	VisitFieldReference(ctx *FieldReferenceContext) interface{}

	// Visit a parse tree produced by FormulaParser#exponentiationExpression.
	VisitExponentiationExpression(ctx *ExponentiationExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#arithExpression.
	VisitArithExpression(ctx *ArithExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#concatExpression.
	VisitConcatExpression(ctx *ConcatExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by FormulaParser#abs.
	VisitAbs(ctx *AbsContext) interface{}

	// Visit a parse tree produced by FormulaParser#addMonths.
	VisitAddMonths(ctx *AddMonthsContext) interface{}

	// Visit a parse tree produced by FormulaParser#and.
	VisitAnd(ctx *AndContext) interface{}

	// Visit a parse tree produced by FormulaParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by FormulaParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by FormulaParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}
}
