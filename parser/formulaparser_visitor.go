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

	// Visit a parse tree produced by FormulaParser#negationExpression.
	VisitNegationExpression(ctx *NegationExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#compareExpression.
	VisitCompareExpression(ctx *CompareExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#positiveExpression.
	VisitPositiveExpression(ctx *PositiveExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#logicExpression.
	VisitLogicExpression(ctx *LogicExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#functionCallExpression.
	VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#fieldReferenceExpression.
	VisitFieldReferenceExpression(ctx *FieldReferenceExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#exponentiationExpression.
	VisitExponentiationExpression(ctx *ExponentiationExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#arithExpression.
	VisitArithExpression(ctx *ArithExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#variableExpression.
	VisitVariableExpression(ctx *VariableExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#concatExpression.
	VisitConcatExpression(ctx *ConcatExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#fieldReference.
	VisitFieldReference(ctx *FieldReferenceContext) interface{}

	// Visit a parse tree produced by FormulaParser#fieldPart.
	VisitFieldPart(ctx *FieldPartContext) interface{}

	// Visit a parse tree produced by FormulaParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by FormulaParser#abs.
	VisitAbs(ctx *AbsContext) interface{}

	// Visit a parse tree produced by FormulaParser#addMonths.
	VisitAddMonths(ctx *AddMonthsContext) interface{}

	// Visit a parse tree produced by FormulaParser#and.
	VisitAnd(ctx *AndContext) interface{}

	// Visit a parse tree produced by FormulaParser#begins.
	VisitBegins(ctx *BeginsContext) interface{}

	// Visit a parse tree produced by FormulaParser#blankvalue.
	VisitBlankvalue(ctx *BlankvalueContext) interface{}

	// Visit a parse tree produced by FormulaParser#br.
	VisitBr(ctx *BrContext) interface{}

	// Visit a parse tree produced by FormulaParser#case.
	VisitCase(ctx *CaseContext) interface{}

	// Visit a parse tree produced by FormulaParser#casesafeid.
	VisitCasesafeid(ctx *CasesafeidContext) interface{}

	// Visit a parse tree produced by FormulaParser#ceiling.
	VisitCeiling(ctx *CeilingContext) interface{}

	// Visit a parse tree produced by FormulaParser#contains.
	VisitContains(ctx *ContainsContext) interface{}

	// Visit a parse tree produced by FormulaParser#currencyrate.
	VisitCurrencyrate(ctx *CurrencyrateContext) interface{}

	// Visit a parse tree produced by FormulaParser#date.
	VisitDate(ctx *DateContext) interface{}

	// Visit a parse tree produced by FormulaParser#datevalue.
	VisitDatevalue(ctx *DatevalueContext) interface{}

	// Visit a parse tree produced by FormulaParser#datetimevalue.
	VisitDatetimevalue(ctx *DatetimevalueContext) interface{}

	// Visit a parse tree produced by FormulaParser#day.
	VisitDay(ctx *DayContext) interface{}

	// Visit a parse tree produced by FormulaParser#distance.
	VisitDistance(ctx *DistanceContext) interface{}

	// Visit a parse tree produced by FormulaParser#exp.
	VisitExp(ctx *ExpContext) interface{}

	// Visit a parse tree produced by FormulaParser#find.
	VisitFind(ctx *FindContext) interface{}

	// Visit a parse tree produced by FormulaParser#floor.
	VisitFloor(ctx *FloorContext) interface{}

	// Visit a parse tree produced by FormulaParser#geolocation.
	VisitGeolocation(ctx *GeolocationContext) interface{}

	// Visit a parse tree produced by FormulaParser#getrecordids.
	VisitGetrecordids(ctx *GetrecordidsContext) interface{}

	// Visit a parse tree produced by FormulaParser#getsessionid.
	VisitGetsessionid(ctx *GetsessionidContext) interface{}

	// Visit a parse tree produced by FormulaParser#hour.
	VisitHour(ctx *HourContext) interface{}

	// Visit a parse tree produced by FormulaParser#htmlencode.
	VisitHtmlencode(ctx *HtmlencodeContext) interface{}

	// Visit a parse tree produced by FormulaParser#hyperlink.
	VisitHyperlink(ctx *HyperlinkContext) interface{}

	// Visit a parse tree produced by FormulaParser#if.
	VisitIf(ctx *IfContext) interface{}

	// Visit a parse tree produced by FormulaParser#image.
	VisitImage(ctx *ImageContext) interface{}

	// Visit a parse tree produced by FormulaParser#imageproxyurl.
	VisitImageproxyurl(ctx *ImageproxyurlContext) interface{}

	// Visit a parse tree produced by FormulaParser#includes.
	VisitIncludes(ctx *IncludesContext) interface{}

	// Visit a parse tree produced by FormulaParser#isblank.
	VisitIsblank(ctx *IsblankContext) interface{}

	// Visit a parse tree produced by FormulaParser#ischanged.
	VisitIschanged(ctx *IschangedContext) interface{}

	// Visit a parse tree produced by FormulaParser#isnew.
	VisitIsnew(ctx *IsnewContext) interface{}

	// Visit a parse tree produced by FormulaParser#isnull.
	VisitIsnull(ctx *IsnullContext) interface{}

	// Visit a parse tree produced by FormulaParser#isnumber.
	VisitIsnumber(ctx *IsnumberContext) interface{}

	// Visit a parse tree produced by FormulaParser#ispickval.
	VisitIspickval(ctx *IspickvalContext) interface{}

	// Visit a parse tree produced by FormulaParser#left.
	VisitLeft(ctx *LeftContext) interface{}

	// Visit a parse tree produced by FormulaParser#len.
	VisitLen(ctx *LenContext) interface{}

	// Visit a parse tree produced by FormulaParser#lower.
	VisitLower(ctx *LowerContext) interface{}

	// Visit a parse tree produced by FormulaParser#lpad.
	VisitLpad(ctx *LpadContext) interface{}

	// Visit a parse tree produced by FormulaParser#max.
	VisitMax(ctx *MaxContext) interface{}

	// Visit a parse tree produced by FormulaParser#mid.
	VisitMid(ctx *MidContext) interface{}

	// Visit a parse tree produced by FormulaParser#min.
	VisitMin(ctx *MinContext) interface{}

	// Visit a parse tree produced by FormulaParser#minute.
	VisitMinute(ctx *MinuteContext) interface{}

	// Visit a parse tree produced by FormulaParser#mod.
	VisitMod(ctx *ModContext) interface{}

	// Visit a parse tree produced by FormulaParser#month.
	VisitMonth(ctx *MonthContext) interface{}

	// Visit a parse tree produced by FormulaParser#not.
	VisitNot(ctx *NotContext) interface{}

	// Visit a parse tree produced by FormulaParser#now.
	VisitNow(ctx *NowContext) interface{}

	// Visit a parse tree produced by FormulaParser#or.
	VisitOr(ctx *OrContext) interface{}

	// Visit a parse tree produced by FormulaParser#priorvalue.
	VisitPriorvalue(ctx *PriorvalueContext) interface{}

	// Visit a parse tree produced by FormulaParser#regex.
	VisitRegex(ctx *RegexContext) interface{}

	// Visit a parse tree produced by FormulaParser#right.
	VisitRight(ctx *RightContext) interface{}

	// Visit a parse tree produced by FormulaParser#round.
	VisitRound(ctx *RoundContext) interface{}

	// Visit a parse tree produced by FormulaParser#substitute.
	VisitSubstitute(ctx *SubstituteContext) interface{}

	// Visit a parse tree produced by FormulaParser#trim.
	VisitTrim(ctx *TrimContext) interface{}

	// Visit a parse tree produced by FormulaParser#text.
	VisitText(ctx *TextContext) interface{}

	// Visit a parse tree produced by FormulaParser#timevalue.
	VisitTimevalue(ctx *TimevalueContext) interface{}

	// Visit a parse tree produced by FormulaParser#today.
	VisitToday(ctx *TodayContext) interface{}

	// Visit a parse tree produced by FormulaParser#upper.
	VisitUpper(ctx *UpperContext) interface{}

	// Visit a parse tree produced by FormulaParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by FormulaParser#year.
	VisitYear(ctx *YearContext) interface{}

	// Visit a parse tree produced by FormulaParser#fieldExpression.
	VisitFieldExpression(ctx *FieldExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#dateExpression.
	VisitDateExpression(ctx *DateExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#valueExpression.
	VisitValueExpression(ctx *ValueExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#resultExpression.
	VisitResultExpression(ctx *ResultExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#substituteValue.
	VisitSubstituteValue(ctx *SubstituteValueContext) interface{}

	// Visit a parse tree produced by FormulaParser#defaultExpression.
	VisitDefaultExpression(ctx *DefaultExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#yearExpression.
	VisitYearExpression(ctx *YearExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#monthExpression.
	VisitMonthExpression(ctx *MonthExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#dayExpression.
	VisitDayExpression(ctx *DayExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#regexExpression.
	VisitRegexExpression(ctx *RegexExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#unitExpression.
	VisitUnitExpression(ctx *UnitExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#searchExpression.
	VisitSearchExpression(ctx *SearchExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#textExpression.
	VisitTextExpression(ctx *TextExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#oldText.
	VisitOldText(ctx *OldTextContext) interface{}

	// Visit a parse tree produced by FormulaParser#replacement.
	VisitReplacement(ctx *ReplacementContext) interface{}

	// Visit a parse tree produced by FormulaParser#startNum.
	VisitStartNum(ctx *StartNumContext) interface{}

	// Visit a parse tree produced by FormulaParser#compareText.
	VisitCompareText(ctx *CompareTextContext) interface{}

	// Visit a parse tree produced by FormulaParser#num.
	VisitNum(ctx *NumContext) interface{}

	// Visit a parse tree produced by FormulaParser#divisor.
	VisitDivisor(ctx *DivisorContext) interface{}

	// Visit a parse tree produced by FormulaParser#digits.
	VisitDigits(ctx *DigitsContext) interface{}

	// Visit a parse tree produced by FormulaParser#latitudeExpression.
	VisitLatitudeExpression(ctx *LatitudeExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#longitudeExpression.
	VisitLongitudeExpression(ctx *LongitudeExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#url.
	VisitUrl(ctx *UrlContext) interface{}

	// Visit a parse tree produced by FormulaParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by FormulaParser#target.
	VisitTarget(ctx *TargetContext) interface{}

	// Visit a parse tree produced by FormulaParser#logicalExpression.
	VisitLogicalExpression(ctx *LogicalExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#ifTrueExpression.
	VisitIfTrueExpression(ctx *IfTrueExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#ifFalseExpression.
	VisitIfFalseExpression(ctx *IfFalseExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#length.
	VisitLength(ctx *LengthContext) interface{}

	// Visit a parse tree produced by FormulaParser#padString.
	VisitPadString(ctx *PadStringContext) interface{}

	// Visit a parse tree produced by FormulaParser#heightExpression.
	VisitHeightExpression(ctx *HeightExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#widthExpression.
	VisitWidthExpression(ctx *WidthExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by FormulaParser#numChars.
	VisitNumChars(ctx *NumCharsContext) interface{}

	// Visit a parse tree produced by FormulaParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by FormulaParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}
}
