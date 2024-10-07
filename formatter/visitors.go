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

func (v *FormatVisitor) VisitHyperlink(ctx *parser.HyperlinkContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	target := ""
	if t := ctx.TargetExpression(); t != nil {
		target = v.visitRule(t).(string)
	}
	if v.wrap {
		args := []string{
			v.indent(v.visitRule(ctx.UrlExpression()).(string)),
			v.indent(v.visitRule(ctx.NameExpression()).(string)),
		}
		if target != "" {
			args = append(args, v.indent(v.visitRule(ctx.TargetExpression()).(string)))
		}
		return fmt.Sprintf("HYPERLINK(\n%s\n)", strings.Join(args, ",\n"))
	}
	args := []string{
		v.visitRule(ctx.UrlExpression()).(string),
		v.visitRule(ctx.NameExpression()).(string),
	}
	if target != "" {
		args = append(args, v.visitRule(ctx.TargetExpression()).(string))
	}
	return fmt.Sprintf("HYPERLINK(%s)", strings.Join(args, ", "))
}

func (v *FormatVisitor) VisitIf(ctx *parser.IfContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("IF(%s,\n%s,\n%s\n)", v.visitRule(ctx.LogicalExpression()),
			v.indent(v.visitRule(ctx.IfTrueExpression()).(string)),
			v.indent(v.visitRule(ctx.IfFalseExpression()).(string)),
		)
	}
	return fmt.Sprintf("IF(%s, %s, %s)", v.visitRule(ctx.LogicalExpression()), v.visitRule(ctx.IfTrueExpression()), v.visitRule(ctx.IfFalseExpression()))
}

func (v *FormatVisitor) VisitMid(ctx *parser.MidContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("MID(\n%s,\n%s,\n%s\n)", v.indent(v.visitRule(ctx.TextExpression()).(string)),
			v.indent(v.visitRule(ctx.StartNum()).(string)),
			v.indent(v.visitRule(ctx.NumChars()).(string)),
		)
	}
	return fmt.Sprintf("MID(%s, %s, %s)", v.visitRule(ctx.TextExpression()), v.visitRule(ctx.StartNum()), v.visitRule(ctx.NumChars()))
}

func (v *FormatVisitor) VisitIncludes(ctx *parser.IncludesContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("INCLUDES(%s,\n%s\n)", v.visitRule(ctx.FieldExpression()), v.indent(v.visitRule(ctx.ValueExpression()).(string)))
	}
	return fmt.Sprintf("INCLUDES(%s, %s)", v.visitRule(ctx.FieldExpression()), v.visitRule(ctx.ValueExpression()))
}

func (v *FormatVisitor) VisitDatevalue(ctx *parser.DatevalueContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("DATEVALUE(\n%s\n)", v.indent(v.visitRule(ctx.Expression()).(string)))
	}
	return fmt.Sprintf("DATEVALUE(\n%s\n)", v.visitRule(ctx.Expression()).(string))
}

func (v *FormatVisitor) VisitText(ctx *parser.TextContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("TEXT(\n%s\n)", v.indent(v.visitRule(ctx.Expression()).(string)))
	}
	return fmt.Sprintf("TEXT(\n%s\n)", v.visitRule(ctx.Expression()).(string))
}

func (v *FormatVisitor) VisitValue(ctx *parser.ValueContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("VALUE(\n%s\n)", v.indent(v.visitRule(ctx.Expression()).(string)))
	}
	return fmt.Sprintf("VALUE(\n%s\n)", v.visitRule(ctx.Expression()).(string))
}

func (v *FormatVisitor) VisitBr(ctx *parser.BrContext) interface{} {
	return "BR()"
}

func (v *FormatVisitor) VisitImage(ctx *parser.ImageContext) interface{} {
	if len(ctx.GetText()) < 80 {
		defer restoreWrap(unwrap(v))
	}
	if v.wrap {
		args := []string{
			v.indent(v.visitRule(ctx.UrlExpression()).(string)),
			v.indent(v.visitRule(ctx.TextExpression()).(string)),
		}
		if h := ctx.HeightExpression(); h != nil {
			args = append(args, v.indent(v.visitRule(h).(string)))
			args = append(args, v.indent(v.visitRule(ctx.WidthExpression()).(string)))
		}
		return fmt.Sprintf("IMAGE(\n%s\n)", strings.Join(args, ",\n"))
	}

	args := []string{
		v.visitRule(ctx.UrlExpression()).(string),
		v.visitRule(ctx.TextExpression()).(string),
	}
	if h := ctx.HeightExpression(); h != nil {
		args = append(args, v.visitRule(h).(string))
		args = append(args, v.visitRule(ctx.WidthExpression()).(string))
	}
	return fmt.Sprintf("IMAGE(%s)", strings.Join(args, ", "))
}

func (v *FormatVisitor) VisitLatitudeExpression(ctx *parser.LatitudeExpressionContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitLongitudeExpression(ctx *parser.LongitudeExpressionContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitLogicalExpression(ctx *parser.LogicalExpressionContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitIfTrueExpression(ctx *parser.IfTrueExpressionContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitIfFalseExpression(ctx *parser.IfFalseExpressionContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitStartNum(ctx *parser.StartNumContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitNumChars(ctx *parser.NumCharsContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitFieldReference(ctx *parser.FieldReferenceContext) interface{} {
	return ctx.GetText()
}

func (v *FormatVisitor) VisitUnitExpression(ctx *parser.UnitExpressionContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitUrlExpression(ctx *parser.UrlExpressionContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitTextExpression(ctx *parser.TextExpressionContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitFieldExpression(ctx *parser.FieldExpressionContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitValueExpression(ctx *parser.ValueExpressionContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitNameExpression(ctx *parser.NameExpressionContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitCompareExpression(ctx *parser.CompareExpressionContext) interface{} {
	cmpToken := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	if ctx.ASSIGN() != nil {
		cmpToken += "="
	}
	return fmt.Sprintf("%s %s %s", v.visitRule(ctx.Expression(0)), cmpToken, v.visitRule(ctx.Expression(1)))
}

func (v *FormatVisitor) VisitConcatExpression(ctx *parser.ConcatExpressionContext) interface{} {
	return fmt.Sprintf("%s & %s", v.visitRule(ctx.Expression(0)), v.visitRule(ctx.Expression(1)))
}

func (v *FormatVisitor) VisitArithExpression(ctx *parser.ArithExpressionContext) interface{} {
	return fmt.Sprintf("%s %s %s", v.visitRule(ctx.Expression(0)), ctx.GetChild(1).(antlr.TerminalNode).GetText(), v.visitRule(ctx.Expression(1)))
}

func (v *FormatVisitor) VisitLogicExpression(ctx *parser.LogicExpressionContext) interface{} {
	return fmt.Sprintf("%s %s %s", v.visitRule(ctx.Expression(0)), ctx.GetChild(1).(antlr.TerminalNode).GetText(), v.visitRule(ctx.Expression(1)))
}

func (v *FormatVisitor) VisitNegationExpression(ctx *parser.NegationExpressionContext) interface{} {
	return fmt.Sprintf("%s%s", ctx.GetChild(0).(antlr.TerminalNode).GetText(), v.visitRule(ctx.Expression()))
}
