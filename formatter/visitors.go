package formatter

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/sformula/parser"
)

func (v *FormatVisitor) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	return ctx.GetText()
}

func (v *FormatVisitor) VisitFunctionCallExpression(ctx *parser.FunctionCallExpressionContext) interface{} {
	return v.visitRule(ctx.FunctionCall())
}

func (v *FormatVisitor) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
	child := ctx.GetChild(0).(antlr.RuleNode)
	return v.visitRule(child)
}

func (v *FormatVisitor) VisitAnd(ctx *parser.AndContext) interface{} {
	wrap := v.wrap || (len(ctx.GetText()) > 40 && len(ctx.AllExpression()) > 3) || len(ctx.GetText()) > 150
	expressions := []string{}
	for _, e := range ctx.AllExpression() {
		expressions = append(expressions, v.visitRule(e).(string))
	}
	if wrap {
		return fmt.Sprintf("AND(\n%s\n)", v.indent(strings.Join(expressions, ",\n")))
	}
	return fmt.Sprintf("AND(%s)", strings.Join(expressions, ", "))
}
