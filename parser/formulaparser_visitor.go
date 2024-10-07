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

	// Visit a parse tree produced by FormulaParser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#logicExpression.
	VisitLogicExpression(ctx *LogicExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#functionCallExpression.
	VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#fieldReference.
	VisitFieldReference(ctx *FieldReferenceContext) interface{}

	// Visit a parse tree produced by FormulaParser#exponentiationExpression.
	VisitExponentiationExpression(ctx *ExponentiationExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#arithExpression.
	VisitArithExpression(ctx *ArithExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#concatExpression.
	VisitConcatExpression(ctx *ConcatExpressionContext) interface{}

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

	// Visit a parse tree produced by FormulaParser#mid.
	VisitMid(ctx *MidContext) interface{}

	// Visit a parse tree produced by FormulaParser#includes.
	VisitIncludes(ctx *IncludesContext) interface{}

	// Visit a parse tree produced by FormulaParser#text.
	VisitText(ctx *TextContext) interface{}

	// Visit a parse tree produced by FormulaParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by FormulaParser#fieldExpression.
	VisitFieldExpression(ctx *FieldExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#valueExpression.
	VisitValueExpression(ctx *ValueExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#resultExpression.
	VisitResultExpression(ctx *ResultExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#defaultExpression.
	VisitDefaultExpression(ctx *DefaultExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#yearExpression.
	VisitYearExpression(ctx *YearExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#monthExpression.
	VisitMonthExpression(ctx *MonthExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#dayExpression.
	VisitDayExpression(ctx *DayExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#unitExpression.
	VisitUnitExpression(ctx *UnitExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#searchExpression.
	VisitSearchExpression(ctx *SearchExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#textExpression.
	VisitTextExpression(ctx *TextExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#startNum.
	VisitStartNum(ctx *StartNumContext) interface{}

	// Visit a parse tree produced by FormulaParser#latitudeExpression.
	VisitLatitudeExpression(ctx *LatitudeExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#longitudeExpression.
	VisitLongitudeExpression(ctx *LongitudeExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#urlExpression.
	VisitUrlExpression(ctx *UrlExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#nameExpression.
	VisitNameExpression(ctx *NameExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#targetExpression.
	VisitTargetExpression(ctx *TargetExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#logicalExpression.
	VisitLogicalExpression(ctx *LogicalExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#ifTrueExpression.
	VisitIfTrueExpression(ctx *IfTrueExpressionContext) interface{}

	// Visit a parse tree produced by FormulaParser#ifFalseExpression.
	VisitIfFalseExpression(ctx *IfFalseExpressionContext) interface{}

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

	// Visit a parse tree produced by FormulaParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by FormulaParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}
}
