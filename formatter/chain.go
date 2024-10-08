package formatter

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/sformula/parser"
	log "github.com/sirupsen/logrus"
)

type ChainVisitor struct {
	parser.BaseFormulaParserVisitor
}

func NewChainVisitor() *ChainVisitor {
	return &ChainVisitor{}
}

func (v *ChainVisitor) visitRule(node antlr.RuleNode) interface{} {
	result := node.Accept(v)
	if r, ok := result.(int); ok {
		return r
	}
	if result == nil {
		log.Debug(fmt.Sprintf("missing ChainVisitor function for %T", node))
	}
	return 0
}

func (v *ChainVisitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	return v.visitRule(ctx.Primary())
}

func (v *ChainVisitor) VisitPrimary(ctx *parser.PrimaryContext) interface{} {
	if e := ctx.Expression(); e != nil {
		return 1 + v.visitRule(e).(int)
	}
	// Literal
	return 1
}

func (v *ChainVisitor) VisitFunctionCallExpression(ctx *parser.FunctionCallExpressionContext) interface{} {
	return 1
}

func (v *ChainVisitor) VisitLogicExpression(ctx *parser.LogicExpressionContext) interface{} {
	return 1 + v.visitRule(ctx.Expression(0)).(int) + v.visitRule(ctx.Expression(1)).(int)
}

func (v *ChainVisitor) VisitExponentiationExpression(ctx *parser.ExponentiationExpressionContext) interface{} {
	return 1 + v.visitRule(ctx.Expression(0)).(int) + v.visitRule(ctx.Expression(1)).(int)
}

func (v *ChainVisitor) VisitConcatExpression(ctx *parser.ConcatExpressionContext) interface{} {
	return 1 + v.visitRule(ctx.Expression(0)).(int) + v.visitRule(ctx.Expression(1)).(int)
}

func (v *ChainVisitor) VisitArithExpression(ctx *parser.ArithExpressionContext) interface{} {
	return 1 + v.visitRule(ctx.Expression(0)).(int) + v.visitRule(ctx.Expression(1)).(int)
}

func (v *ChainVisitor) VisitCompareExpression(ctx *parser.CompareExpressionContext) interface{} {
	return 1 + v.visitRule(ctx.Expression(0)).(int) + v.visitRule(ctx.Expression(1)).(int)
}

func (v *ChainVisitor) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} {
	return 1 + v.visitRule(ctx.Expression(0)).(int) + v.visitRule(ctx.Expression(1)).(int)
}

func (v *ChainVisitor) VisitFieldReference(ctx *parser.FieldReferenceContext) interface{} {
	return 0
}

func (v *ChainVisitor) VisitNegationExpression(ctx *parser.NegationExpressionContext) interface{} {
	return 0
}
