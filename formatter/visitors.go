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
	if (len(ctx.GetText()) > 40 && len(ctx.AllExpression()) > 3) || len(ctx.GetText()) > 150 {
		defer restoreWrap(wrap(v))
	}
	expressions := []string{}
	for _, e := range ctx.AllExpression() {
		expressions = append(expressions, v.visitRule(e).(string))
	}
	if v.wrap {
		return fmt.Sprintf("AND(\n%s\n)", v.indent(strings.Join(expressions, ",\n")))
	}
	return fmt.Sprintf("AND(%s)", strings.Join(expressions, ", "))
}

func (v *FormatVisitor) VisitDistance(ctx *parser.DistanceContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("DISTANCE(\n%s,\n%s,\n%s\n)",
			v.indent(v.visitRule(ctx.Expression(0)).(string)),
			v.indent(v.visitRule(ctx.Expression(1)).(string)),
			v.indent(v.visitRule(ctx.UnitExpression()).(string)),
		)
	}
	return fmt.Sprintf("DISTANCE(%s, %s, %s)", v.visitRule(ctx.Expression(0)), v.visitRule(ctx.Expression(1)), v.visitRule(ctx.UnitExpression()))
}

func (v *FormatVisitor) VisitGeolocation(ctx *parser.GeolocationContext) interface{} {
	if len(ctx.GetText()) < 60 {
		defer restoreWrap(unwrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("GEOLOCATION(\n%s,\n%s\n)",
			v.indent(v.visitRule(ctx.LatitudeExpression()).(string)),
			v.indent(v.visitRule(ctx.LongitudeExpression()).(string)),
		)
	}
	return fmt.Sprintf("GEOLOCATION(%s, %s)", v.visitRule(ctx.LatitudeExpression()), v.visitRule(ctx.LongitudeExpression()))
}

func (v *FormatVisitor) VisitLatitudeExpression(ctx *parser.LatitudeExpressionContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitLongitudeExpression(ctx *parser.LongitudeExpressionContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitFieldReference(ctx *parser.FieldReferenceContext) interface{} {
	return ctx.GetText()
}

func (v *FormatVisitor) VisitUnitExpression(ctx *parser.UnitExpressionContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitArithExpression(ctx *parser.ArithExpressionContext) interface{} {
	return fmt.Sprintf("%s %s %s", v.visitRule(ctx.Expression(0)), ctx.GetChild(1).(antlr.TerminalNode).GetText(), v.visitRule(ctx.Expression(1)))
}

func (v *FormatVisitor) VisitNegationExpression(ctx *parser.NegationExpressionContext) interface{} {
	return fmt.Sprintf("%s%s", ctx.GetChild(0).(antlr.TerminalNode).GetText(), v.visitRule(ctx.Expression()))
}
