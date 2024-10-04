// Code generated from ./FormulaParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // FormulaParser
import "github.com/antlr4-go/antlr/v4"

type BaseFormulaParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFormulaParserVisitor) VisitCompilationUnit(ctx *CompilationUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitCompareExpression(ctx *CompareExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitLogicExpression(ctx *LogicExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitFieldReference(ctx *FieldReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitExponentiationExpression(ctx *ExponentiationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitArithExpression(ctx *ArithExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitConcatExpression(ctx *ConcatExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitAbs(ctx *AbsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitAddMonths(ctx *AddMonthsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitAnd(ctx *AndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
