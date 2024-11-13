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

func (v *BaseFormulaParserVisitor) VisitNegationExpression(ctx *NegationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitCompareExpression(ctx *CompareExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitPositiveExpression(ctx *PositiveExpressionContext) interface{} {
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

func (v *BaseFormulaParserVisitor) VisitFieldReferenceExpression(ctx *FieldReferenceExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitExponentiationExpression(ctx *ExponentiationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitArithExpression(ctx *ArithExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitVariableExpression(ctx *VariableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitConcatExpression(ctx *ConcatExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitFieldReference(ctx *FieldReferenceContext) interface{} {
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

func (v *BaseFormulaParserVisitor) VisitBegins(ctx *BeginsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitBlankvalue(ctx *BlankvalueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitBr(ctx *BrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitCase(ctx *CaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitCasesafeid(ctx *CasesafeidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitCeiling(ctx *CeilingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitContains(ctx *ContainsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitCurrencyrate(ctx *CurrencyrateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitDate(ctx *DateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitDatevalue(ctx *DatevalueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitDatetimevalue(ctx *DatetimevalueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitDay(ctx *DayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitDistance(ctx *DistanceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitExp(ctx *ExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitFind(ctx *FindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitFloor(ctx *FloorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitGeolocation(ctx *GeolocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitGetrecordids(ctx *GetrecordidsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitGetsessionid(ctx *GetsessionidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitHour(ctx *HourContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitHtmlencode(ctx *HtmlencodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitHyperlink(ctx *HyperlinkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitIf(ctx *IfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitImage(ctx *ImageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitImageproxyurl(ctx *ImageproxyurlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitIncludes(ctx *IncludesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitIsblank(ctx *IsblankContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitIschanged(ctx *IschangedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitIsnew(ctx *IsnewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitIsnull(ctx *IsnullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitIsnumber(ctx *IsnumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitIspickval(ctx *IspickvalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitLeft(ctx *LeftContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitLen(ctx *LenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitLower(ctx *LowerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitLpad(ctx *LpadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitMax(ctx *MaxContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitMid(ctx *MidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitMin(ctx *MinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitMod(ctx *ModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitMonth(ctx *MonthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitNot(ctx *NotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitNow(ctx *NowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitOr(ctx *OrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitPriorvalue(ctx *PriorvalueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitRegex(ctx *RegexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitRight(ctx *RightContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitRound(ctx *RoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitSubstitute(ctx *SubstituteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitTrim(ctx *TrimContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitText(ctx *TextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitToday(ctx *TodayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitUpper(ctx *UpperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitYear(ctx *YearContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitFieldExpression(ctx *FieldExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitDateExpression(ctx *DateExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitValueExpression(ctx *ValueExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitResultExpression(ctx *ResultExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitSubstituteValue(ctx *SubstituteValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitDefaultExpression(ctx *DefaultExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitYearExpression(ctx *YearExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitMonthExpression(ctx *MonthExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitDayExpression(ctx *DayExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitRegexExpression(ctx *RegexExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitUnitExpression(ctx *UnitExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitSearchExpression(ctx *SearchExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitTextExpression(ctx *TextExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitOldText(ctx *OldTextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitReplacement(ctx *ReplacementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitStartNum(ctx *StartNumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitCompareText(ctx *CompareTextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitNum(ctx *NumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitDivisor(ctx *DivisorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitDigits(ctx *DigitsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitLatitudeExpression(ctx *LatitudeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitLongitudeExpression(ctx *LongitudeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitUrl(ctx *UrlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitTarget(ctx *TargetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitLogicalExpression(ctx *LogicalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitIfTrueExpression(ctx *IfTrueExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitIfFalseExpression(ctx *IfFalseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitLength(ctx *LengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitPadString(ctx *PadStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitHeightExpression(ctx *HeightExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitWidthExpression(ctx *WidthExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitNumChars(ctx *NumCharsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFormulaParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
