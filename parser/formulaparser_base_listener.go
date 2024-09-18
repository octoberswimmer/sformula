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

// EnterBinaryExpression is called when production binaryExpression is entered.
func (s *BaseFormulaParserListener) EnterBinaryExpression(ctx *BinaryExpressionContext) {}

// ExitBinaryExpression is called when production binaryExpression is exited.
func (s *BaseFormulaParserListener) ExitBinaryExpression(ctx *BinaryExpressionContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseFormulaParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseFormulaParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterFieldReference is called when production fieldReference is entered.
func (s *BaseFormulaParserListener) EnterFieldReference(ctx *FieldReferenceContext) {}

// ExitFieldReference is called when production fieldReference is exited.
func (s *BaseFormulaParserListener) ExitFieldReference(ctx *FieldReferenceContext) {}

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
