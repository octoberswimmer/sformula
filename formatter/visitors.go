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

func (v *FormatVisitor) VisitPrimary(ctx *parser.PrimaryContext) interface{} {
	if e := ctx.Expression(); e != nil {
		return fmt.Sprintf("(%s)", v.visitRule(e))
	}
	// Literal
	return ctx.GetText()
}

func (v *FormatVisitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	return v.visitRule(ctx.Primary())
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

func (v *FormatVisitor) VisitOr(ctx *parser.OrContext) interface{} {
	if (len(ctx.GetText()) > 40 && len(ctx.AllExpression()) > 3) || len(ctx.GetText()) > 150 {
		defer restoreWrap(wrap(v))
	}
	expressions := []string{}
	for _, e := range ctx.AllExpression() {
		expressions = append(expressions, v.visitRule(e).(string))
	}
	if v.wrap {
		return fmt.Sprintf("OR(\n%s\n)", v.indent(strings.Join(expressions, ",\n")))
	}
	return fmt.Sprintf("OR(%s)", strings.Join(expressions, ", "))
}

func (v *FormatVisitor) VisitMin(ctx *parser.MinContext) interface{} {
	if (len(ctx.GetText()) > 40 && len(ctx.AllExpression()) > 3) || len(ctx.GetText()) > 150 {
		defer restoreWrap(wrap(v))
	}
	expressions := []string{}
	for _, e := range ctx.AllExpression() {
		expressions = append(expressions, v.visitRule(e).(string))
	}
	if v.wrap {
		return fmt.Sprintf("MIN(\n%s\n)", v.indent(strings.Join(expressions, ",\n")))
	}
	return fmt.Sprintf("MIN(%s)", strings.Join(expressions, ", "))
}

func (v *FormatVisitor) VisitCase(ctx *parser.CaseContext) interface{} {
	defer restoreWrap(wrap(v))
	cases := []string{}
	valueExpressions := ctx.AllValueExpression()
	resultExpressions := ctx.AllResultExpression()
	for i, e := range valueExpressions {
		cases = append(cases, v.indent(fmt.Sprintf("%s, %s", v.visitRule(e), v.visitRule(resultExpressions[i]))))
	}
	return fmt.Sprintf("CASE(\n%s,\n%s,\n%s\n)", v.indent(v.visitRule(ctx.Expression()).(string)),
		v.indent(strings.Join(cases, ",\n")),
		v.indent(v.visitRule(ctx.DefaultExpression()).(string)),
	)
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
	if t := ctx.Target(); t != nil {
		target = v.visitRule(t).(string)
	}
	if v.wrap {
		args := []string{
			v.indent(v.visitRule(ctx.Url()).(string)),
			v.indent(v.visitRule(ctx.Name()).(string)),
		}
		if target != "" {
			args = append(args, v.indent(v.visitRule(ctx.Target()).(string)))
		}
		return fmt.Sprintf("HYPERLINK(\n%s\n)", strings.Join(args, ",\n"))
	}
	args := []string{
		v.visitRule(ctx.Url()).(string),
		v.visitRule(ctx.Name()).(string),
	}
	if target != "" {
		args = append(args, v.visitRule(ctx.Target()).(string))
	}
	return fmt.Sprintf("HYPERLINK(%s)", strings.Join(args, ", "))
}

func (v *FormatVisitor) VisitFind(ctx *parser.FindContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		args := []string{
			v.indent(v.visitRule(ctx.SearchExpression()).(string)),
			v.indent(v.visitRule(ctx.TextExpression()).(string)),
		}
		if s := ctx.StartNum(); s != nil {
			args = append(args, v.indent(v.visitRule(s).(string)))
		}
		return fmt.Sprintf("FIND(\n%s\n)", strings.Join(args, ",\n"))
	}
	args := []string{
		v.visitRule(ctx.SearchExpression()).(string),
		v.visitRule(ctx.TextExpression()).(string),
	}
	if s := ctx.StartNum(); s != nil {
		args = append(args, v.visitRule(s).(string))
	}
	return fmt.Sprintf("FIND(%s)", strings.Join(args, ", "))
}

func (v *FormatVisitor) VisitIf(ctx *parser.IfContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("IF(\n%s,\n%s,\n%s\n)", v.indent(v.visitRule(ctx.LogicalExpression()).(string)),
			v.indent(v.visitRule(ctx.IfTrueExpression()).(string)),
			v.indent(v.visitRule(ctx.IfFalseExpression()).(string)),
		)
	}
	return fmt.Sprintf("IF(%s, %s, %s)", v.visitRule(ctx.LogicalExpression()),
		v.visitRule(ctx.IfTrueExpression()),
		v.visitRule(ctx.IfFalseExpression()),
	)
}

func (v *FormatVisitor) VisitSubstitute(ctx *parser.SubstituteContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("SUBSTITUTE(\n%s,\n%s,\n%s\n)", v.visitRule(ctx.TextExpression()),
			v.indent(v.visitRule(ctx.OldText()).(string)),
			v.indent(v.visitRule(ctx.Replacement()).(string)),
		)
	}
	return fmt.Sprintf("SUBSTITUTE(%s, %s, %s)", v.visitRule(ctx.TextExpression()),
		v.visitRule(ctx.OldText()),
		v.visitRule(ctx.Replacement()),
	)
}

func (v *FormatVisitor) VisitMid(ctx *parser.MidContext) interface{} {
	if len(ctx.GetText()) < 40 {
		defer restoreWrap(unwrap(v))
	}
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

func (v *FormatVisitor) VisitLeft(ctx *parser.LeftContext) interface{} {
	if len(ctx.GetText()) < 40 {
		defer restoreWrap(unwrap(v))
	}
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("LEFT(%s,\n%s\n)", v.visitRule(ctx.TextExpression()), v.indent(v.visitRule(ctx.NumChars()).(string)))
	}
	return fmt.Sprintf("LEFT(%s, %s)", v.visitRule(ctx.TextExpression()), v.visitRule(ctx.NumChars()))
}

func (v *FormatVisitor) VisitRight(ctx *parser.RightContext) interface{} {
	if len(ctx.GetText()) < 40 {
		defer restoreWrap(unwrap(v))
	}
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("RIGHT(%s,\n%s\n)", v.visitRule(ctx.TextExpression()), v.indent(v.visitRule(ctx.NumChars()).(string)))
	}
	return fmt.Sprintf("RIGHT(%s, %s)", v.visitRule(ctx.TextExpression()), v.visitRule(ctx.NumChars()))
}

func (v *FormatVisitor) VisitIncludes(ctx *parser.IncludesContext) interface{} {
	if len(ctx.GetText()) < 40 {
		defer restoreWrap(unwrap(v))
	}
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("INCLUDES(%s,\n%s\n)", v.visitRule(ctx.FieldExpression()), v.indent(v.visitRule(ctx.ValueExpression()).(string)))
	}
	return fmt.Sprintf("INCLUDES(%s, %s)", v.visitRule(ctx.FieldExpression()), v.visitRule(ctx.ValueExpression()))
}

func (v *FormatVisitor) VisitIspickval(ctx *parser.IspickvalContext) interface{} {
	if len(ctx.GetText()) < 40 {
		defer restoreWrap(unwrap(v))
	}
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("ISPICKVAL(\n%s,\n%s\n)", v.indent(v.visitRule(ctx.FieldExpression()).(string)), v.indent(v.visitRule(ctx.ValueExpression()).(string)))
	}
	return fmt.Sprintf("ISPICKVAL(%s, %s)", v.visitRule(ctx.FieldExpression()), v.visitRule(ctx.ValueExpression()))
}

func (v *FormatVisitor) VisitMod(ctx *parser.ModContext) interface{} {
	if len(ctx.GetText()) < 60 {
		defer restoreWrap(unwrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("MOD(\n%s,\n%s\n)", v.indent(v.visitRule(ctx.Num()).(string)), v.indent(v.visitRule(ctx.Divisor()).(string)))
	}
	return fmt.Sprintf("MOD(%s, %s)", v.visitRule(ctx.Num()), v.visitRule(ctx.Divisor()))
}

func (v *FormatVisitor) VisitRound(ctx *parser.RoundContext) interface{} {
	if len(ctx.GetText()) < 60 {
		defer restoreWrap(unwrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("ROUND(\n%s,\n%s\n)", v.indent(v.visitRule(ctx.Num()).(string)), v.indent(v.visitRule(ctx.Digits()).(string)))
	}
	return fmt.Sprintf("ROUND(%s, %s)", v.visitRule(ctx.Num()), v.visitRule(ctx.Digits()))
}

func (v *FormatVisitor) VisitBegins(ctx *parser.BeginsContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("BEGINS(\n%s,\n%s\n)", v.indent(v.visitRule(ctx.TextExpression()).(string)), v.indent(v.visitRule(ctx.CompareText()).(string)))
	}
	return fmt.Sprintf("BEGINS(%s, %s)", v.visitRule(ctx.TextExpression()), v.visitRule(ctx.CompareText()))
}

func (v *FormatVisitor) VisitContains(ctx *parser.ContainsContext) interface{} {
	if len(ctx.GetText()) < 40 {
		defer restoreWrap(unwrap(v))
	} else if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("CONTAINS(\n%s,\n%s\n)", v.indent(v.visitRule(ctx.TextExpression()).(string)), v.indent(v.visitRule(ctx.CompareText()).(string)))
	}
	return fmt.Sprintf("CONTAINS(%s, %s)", v.visitRule(ctx.TextExpression()), v.visitRule(ctx.CompareText()))
}

func (v *FormatVisitor) VisitDate(ctx *parser.DateContext) interface{} {
	if len(ctx.GetText()) < 60 {
		defer restoreWrap(unwrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("DATE(\n%s\n%s\n%s\n)",
			v.indent(v.visitRule(ctx.YearExpression()).(string)),
			v.indent(v.visitRule(ctx.MonthExpression()).(string)),
			v.indent(v.visitRule(ctx.DayExpression()).(string)),
		)
	}
	return fmt.Sprintf("DATE(%s, %s, %s)",
		v.visitRule(ctx.YearExpression()),
		v.visitRule(ctx.MonthExpression()),
		v.visitRule(ctx.DayExpression()),
	)
}

func (v *FormatVisitor) VisitDatevalue(ctx *parser.DatevalueContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("DATEVALUE(\n%s\n)", v.indent(v.visitRule(ctx.Expression()).(string)))
	}
	return fmt.Sprintf("DATEVALUE(%s)", v.visitRule(ctx.Expression()).(string))
}

func (v *FormatVisitor) VisitDatetimevalue(ctx *parser.DatetimevalueContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("DATETIMVALUE(\n%s\n)", v.indent(v.visitRule(ctx.Expression()).(string)))
	}
	return fmt.Sprintf("DATETIMVALUE(%s)", v.visitRule(ctx.Expression()).(string))
}

func (v *FormatVisitor) VisitText(ctx *parser.TextContext) interface{} {
	if len(ctx.GetText()) < 60 {
		defer restoreWrap(unwrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("TEXT(\n%s\n)", v.indent(v.visitRule(ctx.Expression()).(string)))
	}
	return fmt.Sprintf("TEXT(%s)", v.visitRule(ctx.Expression()).(string))
}

func (v *FormatVisitor) VisitFloor(ctx *parser.FloorContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("FLOOR(\n%s\n)", v.indent(v.visitRule(ctx.Expression()).(string)))
	}
	return fmt.Sprintf("FLOOR(%s)", v.visitRule(ctx.Expression()).(string))
}

func (v *FormatVisitor) VisitCasesafeid(ctx *parser.CasesafeidContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("CASESAFEID(\n%s\n)", v.indent(v.visitRule(ctx.Expression()).(string)))
	}
	return fmt.Sprintf("CASESAFEID(%s)", v.visitRule(ctx.Expression()).(string))
}

func (v *FormatVisitor) VisitLen(ctx *parser.LenContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("LEN(\n%s\n)", v.indent(v.visitRule(ctx.Expression()).(string)))
	}
	return fmt.Sprintf("LEN(%s)", v.visitRule(ctx.Expression()).(string))
}

func (v *FormatVisitor) VisitAbs(ctx *parser.AbsContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("ABS(\n%s\n)", v.indent(v.visitRule(ctx.Expression()).(string)))
	}
	return fmt.Sprintf("ABS(%s)", v.visitRule(ctx.Expression()).(string))
}

func (v *FormatVisitor) VisitCeiling(ctx *parser.CeilingContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("CEILING(\n%s\n)", v.indent(v.visitRule(ctx.Expression()).(string)))
	}
	return fmt.Sprintf("CEILING(%s)", v.visitRule(ctx.Expression()).(string))
}

func (v *FormatVisitor) VisitNot(ctx *parser.NotContext) interface{} {
	if len(ctx.GetText()) < 40 {
		defer restoreWrap(unwrap(v))
	}
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("NOT(\n%s\n)", v.indent(v.visitRule(ctx.Expression()).(string)))
	}
	return fmt.Sprintf("NOT(%s)", v.visitRule(ctx.Expression()).(string))
}

func (v *FormatVisitor) VisitValue(ctx *parser.ValueContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("VALUE(\n%s\n)", v.indent(v.visitRule(ctx.Expression()).(string)))
	}
	return fmt.Sprintf("VALUE(%s)", v.visitRule(ctx.Expression()).(string))
}

func (v *FormatVisitor) VisitDay(ctx *parser.DayContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("DAY(\n%s\n)", v.indent(v.visitRule(ctx.Expression()).(string)))
	}
	return fmt.Sprintf("DAY(%s)", v.visitRule(ctx.Expression()).(string))
}

func (v *FormatVisitor) VisitMonth(ctx *parser.MonthContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("MONTH(\n%s\n)", v.indent(v.visitRule(ctx.Expression()).(string)))
	}
	return fmt.Sprintf("MONTH(%s)", v.visitRule(ctx.Expression()).(string))
}

func (v *FormatVisitor) VisitYear(ctx *parser.YearContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("YEAR(\n%s\n)", v.indent(v.visitRule(ctx.Expression()).(string)))
	}
	return fmt.Sprintf("YEAR(%s)", v.visitRule(ctx.Expression()).(string))
}

func (v *FormatVisitor) VisitIsnull(ctx *parser.IsnullContext) interface{} {
	if len(ctx.GetText()) < 40 {
		defer restoreWrap(unwrap(v))
	}
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("ISNULL(\n%s\n)", v.indent(v.visitRule(ctx.Expression()).(string)))
	}
	return fmt.Sprintf("ISNULL(%s)", v.visitRule(ctx.Expression()).(string))
}

func (v *FormatVisitor) VisitTrim(ctx *parser.TrimContext) interface{} {
	if len(ctx.GetText()) > 40 {
		defer restoreWrap(unwrap(v))
	}
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("TRIM(\n%s\n)", v.indent(v.visitRule(ctx.Expression()).(string)))
	}
	return fmt.Sprintf("TRIM(%s)", v.visitRule(ctx.Expression()).(string))
}

func (v *FormatVisitor) VisitIsblank(ctx *parser.IsblankContext) interface{} {
	if len(ctx.GetText()) > 60 {
		defer restoreWrap(wrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("ISBLANK(\n%s\n)", v.indent(v.visitRule(ctx.Expression()).(string)))
	}
	return fmt.Sprintf("ISBLANK(%s)", v.visitRule(ctx.Expression()).(string))
}

func (v *FormatVisitor) VisitBlankvalue(ctx *parser.BlankvalueContext) interface{} {
	if len(ctx.GetText()) < 60 {
		defer restoreWrap(unwrap(v))
	}
	if v.wrap {
		return fmt.Sprintf("BLANKVALUE(\n%s,\n%s\n)", v.indent(v.visitRule(ctx.Expression()).(string)), v.indent(v.visitRule(ctx.SubstituteValue()).(string)))
	}
	return fmt.Sprintf("BLANKVALUE(%s, %s)", v.visitRule(ctx.Expression()), v.visitRule(ctx.SubstituteValue()))
}

func (v *FormatVisitor) VisitToday(ctx *parser.TodayContext) interface{} {
	return "TODAY()"
}

func (v *FormatVisitor) VisitNow(ctx *parser.NowContext) interface{} {
	return "NOW()"
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
			v.indent(v.visitRule(ctx.Url()).(string)),
			v.indent(v.visitRule(ctx.TextExpression()).(string)),
		}
		if h := ctx.HeightExpression(); h != nil {
			args = append(args, v.indent(v.visitRule(h).(string)))
			args = append(args, v.indent(v.visitRule(ctx.WidthExpression()).(string)))
		}
		return fmt.Sprintf("IMAGE(\n%s\n)", strings.Join(args, ",\n"))
	}

	args := []string{
		v.visitRule(ctx.Url()).(string),
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

func (v *FormatVisitor) VisitNum(ctx *parser.NumContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitDivisor(ctx *parser.DivisorContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitDigits(ctx *parser.DigitsContext) interface{} {
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

func (v *FormatVisitor) VisitSearchExpression(ctx *parser.SearchExpressionContext) interface{} {
	return ctx.GetText()
}

func (v *FormatVisitor) VisitUnitExpression(ctx *parser.UnitExpressionContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitUrl(ctx *parser.UrlContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitTextExpression(ctx *parser.TextExpressionContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitOldText(ctx *parser.OldTextContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitReplacement(ctx *parser.ReplacementContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitCompareText(ctx *parser.CompareTextContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitFieldExpression(ctx *parser.FieldExpressionContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitValueExpression(ctx *parser.ValueExpressionContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitTarget(ctx *parser.TargetContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitResultExpression(ctx *parser.ResultExpressionContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitDefaultExpression(ctx *parser.DefaultExpressionContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitName(ctx *parser.NameContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitSubstituteValue(ctx *parser.SubstituteValueContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitYearExpression(ctx *parser.YearExpressionContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitMonthExpression(ctx *parser.MonthExpressionContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *FormatVisitor) VisitDayExpression(ctx *parser.DayExpressionContext) interface{} {
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
	// TODO: Use ChainVisitor to wrap
	return fmt.Sprintf("%s %s %s", v.visitRule(ctx.Expression(0)), ctx.GetChild(1).(antlr.TerminalNode).GetText(), v.visitRule(ctx.Expression(1)))
}

func (v *FormatVisitor) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} {
	defer restoreWrap(unwrap(v))
	cmpToken := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	return fmt.Sprintf("%s %s %s", v.visitRule(ctx.Expression(0)), cmpToken, v.visitRule(ctx.Expression(1)))
}

func (v *FormatVisitor) VisitNegationExpression(ctx *parser.NegationExpressionContext) interface{} {
	return fmt.Sprintf("%s%s", ctx.GetChild(0).(antlr.TerminalNode).GetText(), v.visitRule(ctx.Expression()))
}
