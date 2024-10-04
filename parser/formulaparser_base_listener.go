// Code generated from ./FormulaParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // FormulaParser
import "github.com/antlr4-go/antlr/v4"

// BaseFormulaParserListener is a complete listener for a parse tree produced by FormulaParser.
type BaseFormulaParserListener struct{}

var _ FormulaParserListener = &BaseFormulaParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFormulaParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFormulaParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFormulaParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFormulaParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseFormulaParserListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseFormulaParserListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseFormulaParserListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseFormulaParserListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterCompareExpression is called when production compareExpression is entered.
func (s *BaseFormulaParserListener) EnterCompareExpression(ctx *CompareExpressionContext) {}

// ExitCompareExpression is called when production compareExpression is exited.
func (s *BaseFormulaParserListener) ExitCompareExpression(ctx *CompareExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseFormulaParserListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseFormulaParserListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterLogicExpression is called when production logicExpression is entered.
func (s *BaseFormulaParserListener) EnterLogicExpression(ctx *LogicExpressionContext) {}

// ExitLogicExpression is called when production logicExpression is exited.
func (s *BaseFormulaParserListener) ExitLogicExpression(ctx *LogicExpressionContext) {}

// EnterFunctionCallExpression is called when production functionCallExpression is entered.
func (s *BaseFormulaParserListener) EnterFunctionCallExpression(ctx *FunctionCallExpressionContext) {}

// ExitFunctionCallExpression is called when production functionCallExpression is exited.
func (s *BaseFormulaParserListener) ExitFunctionCallExpression(ctx *FunctionCallExpressionContext) {}

// EnterFieldReference is called when production fieldReference is entered.
func (s *BaseFormulaParserListener) EnterFieldReference(ctx *FieldReferenceContext) {}

// ExitFieldReference is called when production fieldReference is exited.
func (s *BaseFormulaParserListener) ExitFieldReference(ctx *FieldReferenceContext) {}

// EnterExponentiationExpression is called when production exponentiationExpression is entered.
func (s *BaseFormulaParserListener) EnterExponentiationExpression(ctx *ExponentiationExpressionContext) {
}

// ExitExponentiationExpression is called when production exponentiationExpression is exited.
func (s *BaseFormulaParserListener) ExitExponentiationExpression(ctx *ExponentiationExpressionContext) {
}

// EnterArithExpression is called when production arithExpression is entered.
func (s *BaseFormulaParserListener) EnterArithExpression(ctx *ArithExpressionContext) {}

// ExitArithExpression is called when production arithExpression is exited.
func (s *BaseFormulaParserListener) ExitArithExpression(ctx *ArithExpressionContext) {}

// EnterConcatExpression is called when production concatExpression is entered.
func (s *BaseFormulaParserListener) EnterConcatExpression(ctx *ConcatExpressionContext) {}

// ExitConcatExpression is called when production concatExpression is exited.
func (s *BaseFormulaParserListener) ExitConcatExpression(ctx *ConcatExpressionContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseFormulaParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseFormulaParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterAbs is called when production abs is entered.
func (s *BaseFormulaParserListener) EnterAbs(ctx *AbsContext) {}

// ExitAbs is called when production abs is exited.
func (s *BaseFormulaParserListener) ExitAbs(ctx *AbsContext) {}

// EnterAddMonths is called when production addMonths is entered.
func (s *BaseFormulaParserListener) EnterAddMonths(ctx *AddMonthsContext) {}

// ExitAddMonths is called when production addMonths is exited.
func (s *BaseFormulaParserListener) ExitAddMonths(ctx *AddMonthsContext) {}

// EnterAnd is called when production and is entered.
func (s *BaseFormulaParserListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production and is exited.
func (s *BaseFormulaParserListener) ExitAnd(ctx *AndContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseFormulaParserListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseFormulaParserListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseFormulaParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseFormulaParserListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseFormulaParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseFormulaParserListener) ExitLiteral(ctx *LiteralContext) {}
