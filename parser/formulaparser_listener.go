// Code generated from ./FormulaParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // FormulaParser
import "github.com/antlr4-go/antlr/v4"

// FormulaParserListener is a complete listener for a parse tree produced by FormulaParser.
type FormulaParserListener interface {
	antlr.ParseTreeListener

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterNegationExpression is called when entering the negationExpression production.
	EnterNegationExpression(c *NegationExpressionContext)

	// EnterCompareExpression is called when entering the compareExpression production.
	EnterCompareExpression(c *CompareExpressionContext)

	// EnterEqualityExpression is called when entering the equalityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterLogicExpression is called when entering the logicExpression production.
	EnterLogicExpression(c *LogicExpressionContext)

	// EnterFunctionCallExpression is called when entering the functionCallExpression production.
	EnterFunctionCallExpression(c *FunctionCallExpressionContext)

	// EnterFieldReference is called when entering the fieldReference production.
	EnterFieldReference(c *FieldReferenceContext)

	// EnterExponentiationExpression is called when entering the exponentiationExpression production.
	EnterExponentiationExpression(c *ExponentiationExpressionContext)

	// EnterArithExpression is called when entering the arithExpression production.
	EnterArithExpression(c *ArithExpressionContext)

	// EnterConcatExpression is called when entering the concatExpression production.
	EnterConcatExpression(c *ConcatExpressionContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterAbs is called when entering the abs production.
	EnterAbs(c *AbsContext)

	// EnterAddMonths is called when entering the addMonths production.
	EnterAddMonths(c *AddMonthsContext)

	// EnterAnd is called when entering the and production.
	EnterAnd(c *AndContext)

	// EnterBegins is called when entering the begins production.
	EnterBegins(c *BeginsContext)

	// EnterBlankvalue is called when entering the blankvalue production.
	EnterBlankvalue(c *BlankvalueContext)

	// EnterBr is called when entering the br production.
	EnterBr(c *BrContext)

	// EnterCase is called when entering the case production.
	EnterCase(c *CaseContext)

	// EnterCasesafeid is called when entering the casesafeid production.
	EnterCasesafeid(c *CasesafeidContext)

	// EnterCeiling is called when entering the ceiling production.
	EnterCeiling(c *CeilingContext)

	// EnterContains is called when entering the contains production.
	EnterContains(c *ContainsContext)

	// EnterCurrencyrate is called when entering the currencyrate production.
	EnterCurrencyrate(c *CurrencyrateContext)

	// EnterDate is called when entering the date production.
	EnterDate(c *DateContext)

	// EnterDatevalue is called when entering the datevalue production.
	EnterDatevalue(c *DatevalueContext)

	// EnterDatetimevalue is called when entering the datetimevalue production.
	EnterDatetimevalue(c *DatetimevalueContext)

	// EnterDay is called when entering the day production.
	EnterDay(c *DayContext)

	// EnterDistance is called when entering the distance production.
	EnterDistance(c *DistanceContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// EnterFind is called when entering the find production.
	EnterFind(c *FindContext)

	// EnterFloor is called when entering the floor production.
	EnterFloor(c *FloorContext)

	// EnterGeolocation is called when entering the geolocation production.
	EnterGeolocation(c *GeolocationContext)

	// EnterGetrecordids is called when entering the getrecordids production.
	EnterGetrecordids(c *GetrecordidsContext)

	// EnterGetsessionid is called when entering the getsessionid production.
	EnterGetsessionid(c *GetsessionidContext)

	// EnterHour is called when entering the hour production.
	EnterHour(c *HourContext)

	// EnterHtmlencode is called when entering the htmlencode production.
	EnterHtmlencode(c *HtmlencodeContext)

	// EnterHyperlink is called when entering the hyperlink production.
	EnterHyperlink(c *HyperlinkContext)

	// EnterIf is called when entering the if production.
	EnterIf(c *IfContext)

	// EnterImage is called when entering the image production.
	EnterImage(c *ImageContext)

	// EnterImageproxyurl is called when entering the imageproxyurl production.
	EnterImageproxyurl(c *ImageproxyurlContext)

	// EnterMid is called when entering the mid production.
	EnterMid(c *MidContext)

	// EnterIncludes is called when entering the includes production.
	EnterIncludes(c *IncludesContext)

	// EnterText is called when entering the text production.
	EnterText(c *TextContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterFieldExpression is called when entering the fieldExpression production.
	EnterFieldExpression(c *FieldExpressionContext)

	// EnterValueExpression is called when entering the valueExpression production.
	EnterValueExpression(c *ValueExpressionContext)

	// EnterResultExpression is called when entering the resultExpression production.
	EnterResultExpression(c *ResultExpressionContext)

	// EnterDefaultExpression is called when entering the defaultExpression production.
	EnterDefaultExpression(c *DefaultExpressionContext)

	// EnterYearExpression is called when entering the yearExpression production.
	EnterYearExpression(c *YearExpressionContext)

	// EnterMonthExpression is called when entering the monthExpression production.
	EnterMonthExpression(c *MonthExpressionContext)

	// EnterDayExpression is called when entering the dayExpression production.
	EnterDayExpression(c *DayExpressionContext)

	// EnterUnitExpression is called when entering the unitExpression production.
	EnterUnitExpression(c *UnitExpressionContext)

	// EnterSearchExpression is called when entering the searchExpression production.
	EnterSearchExpression(c *SearchExpressionContext)

	// EnterTextExpression is called when entering the textExpression production.
	EnterTextExpression(c *TextExpressionContext)

	// EnterStartNum is called when entering the startNum production.
	EnterStartNum(c *StartNumContext)

	// EnterLatitudeExpression is called when entering the latitudeExpression production.
	EnterLatitudeExpression(c *LatitudeExpressionContext)

	// EnterLongitudeExpression is called when entering the longitudeExpression production.
	EnterLongitudeExpression(c *LongitudeExpressionContext)

	// EnterUrlExpression is called when entering the urlExpression production.
	EnterUrlExpression(c *UrlExpressionContext)

	// EnterNameExpression is called when entering the nameExpression production.
	EnterNameExpression(c *NameExpressionContext)

	// EnterTargetExpression is called when entering the targetExpression production.
	EnterTargetExpression(c *TargetExpressionContext)

	// EnterLogicalExpression is called when entering the logicalExpression production.
	EnterLogicalExpression(c *LogicalExpressionContext)

	// EnterIfTrueExpression is called when entering the ifTrueExpression production.
	EnterIfTrueExpression(c *IfTrueExpressionContext)

	// EnterIfFalseExpression is called when entering the ifFalseExpression production.
	EnterIfFalseExpression(c *IfFalseExpressionContext)

	// EnterHeightExpression is called when entering the heightExpression production.
	EnterHeightExpression(c *HeightExpressionContext)

	// EnterWidthExpression is called when entering the widthExpression production.
	EnterWidthExpression(c *WidthExpressionContext)

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterNumChars is called when entering the numChars production.
	EnterNumChars(c *NumCharsContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitNegationExpression is called when exiting the negationExpression production.
	ExitNegationExpression(c *NegationExpressionContext)

	// ExitCompareExpression is called when exiting the compareExpression production.
	ExitCompareExpression(c *CompareExpressionContext)

	// ExitEqualityExpression is called when exiting the equalityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitLogicExpression is called when exiting the logicExpression production.
	ExitLogicExpression(c *LogicExpressionContext)

	// ExitFunctionCallExpression is called when exiting the functionCallExpression production.
	ExitFunctionCallExpression(c *FunctionCallExpressionContext)

	// ExitFieldReference is called when exiting the fieldReference production.
	ExitFieldReference(c *FieldReferenceContext)

	// ExitExponentiationExpression is called when exiting the exponentiationExpression production.
	ExitExponentiationExpression(c *ExponentiationExpressionContext)

	// ExitArithExpression is called when exiting the arithExpression production.
	ExitArithExpression(c *ArithExpressionContext)

	// ExitConcatExpression is called when exiting the concatExpression production.
	ExitConcatExpression(c *ConcatExpressionContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitAbs is called when exiting the abs production.
	ExitAbs(c *AbsContext)

	// ExitAddMonths is called when exiting the addMonths production.
	ExitAddMonths(c *AddMonthsContext)

	// ExitAnd is called when exiting the and production.
	ExitAnd(c *AndContext)

	// ExitBegins is called when exiting the begins production.
	ExitBegins(c *BeginsContext)

	// ExitBlankvalue is called when exiting the blankvalue production.
	ExitBlankvalue(c *BlankvalueContext)

	// ExitBr is called when exiting the br production.
	ExitBr(c *BrContext)

	// ExitCase is called when exiting the case production.
	ExitCase(c *CaseContext)

	// ExitCasesafeid is called when exiting the casesafeid production.
	ExitCasesafeid(c *CasesafeidContext)

	// ExitCeiling is called when exiting the ceiling production.
	ExitCeiling(c *CeilingContext)

	// ExitContains is called when exiting the contains production.
	ExitContains(c *ContainsContext)

	// ExitCurrencyrate is called when exiting the currencyrate production.
	ExitCurrencyrate(c *CurrencyrateContext)

	// ExitDate is called when exiting the date production.
	ExitDate(c *DateContext)

	// ExitDatevalue is called when exiting the datevalue production.
	ExitDatevalue(c *DatevalueContext)

	// ExitDatetimevalue is called when exiting the datetimevalue production.
	ExitDatetimevalue(c *DatetimevalueContext)

	// ExitDay is called when exiting the day production.
	ExitDay(c *DayContext)

	// ExitDistance is called when exiting the distance production.
	ExitDistance(c *DistanceContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)

	// ExitFind is called when exiting the find production.
	ExitFind(c *FindContext)

	// ExitFloor is called when exiting the floor production.
	ExitFloor(c *FloorContext)

	// ExitGeolocation is called when exiting the geolocation production.
	ExitGeolocation(c *GeolocationContext)

	// ExitGetrecordids is called when exiting the getrecordids production.
	ExitGetrecordids(c *GetrecordidsContext)

	// ExitGetsessionid is called when exiting the getsessionid production.
	ExitGetsessionid(c *GetsessionidContext)

	// ExitHour is called when exiting the hour production.
	ExitHour(c *HourContext)

	// ExitHtmlencode is called when exiting the htmlencode production.
	ExitHtmlencode(c *HtmlencodeContext)

	// ExitHyperlink is called when exiting the hyperlink production.
	ExitHyperlink(c *HyperlinkContext)

	// ExitIf is called when exiting the if production.
	ExitIf(c *IfContext)

	// ExitImage is called when exiting the image production.
	ExitImage(c *ImageContext)

	// ExitImageproxyurl is called when exiting the imageproxyurl production.
	ExitImageproxyurl(c *ImageproxyurlContext)

	// ExitMid is called when exiting the mid production.
	ExitMid(c *MidContext)

	// ExitIncludes is called when exiting the includes production.
	ExitIncludes(c *IncludesContext)

	// ExitText is called when exiting the text production.
	ExitText(c *TextContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitFieldExpression is called when exiting the fieldExpression production.
	ExitFieldExpression(c *FieldExpressionContext)

	// ExitValueExpression is called when exiting the valueExpression production.
	ExitValueExpression(c *ValueExpressionContext)

	// ExitResultExpression is called when exiting the resultExpression production.
	ExitResultExpression(c *ResultExpressionContext)

	// ExitDefaultExpression is called when exiting the defaultExpression production.
	ExitDefaultExpression(c *DefaultExpressionContext)

	// ExitYearExpression is called when exiting the yearExpression production.
	ExitYearExpression(c *YearExpressionContext)

	// ExitMonthExpression is called when exiting the monthExpression production.
	ExitMonthExpression(c *MonthExpressionContext)

	// ExitDayExpression is called when exiting the dayExpression production.
	ExitDayExpression(c *DayExpressionContext)

	// ExitUnitExpression is called when exiting the unitExpression production.
	ExitUnitExpression(c *UnitExpressionContext)

	// ExitSearchExpression is called when exiting the searchExpression production.
	ExitSearchExpression(c *SearchExpressionContext)

	// ExitTextExpression is called when exiting the textExpression production.
	ExitTextExpression(c *TextExpressionContext)

	// ExitStartNum is called when exiting the startNum production.
	ExitStartNum(c *StartNumContext)

	// ExitLatitudeExpression is called when exiting the latitudeExpression production.
	ExitLatitudeExpression(c *LatitudeExpressionContext)

	// ExitLongitudeExpression is called when exiting the longitudeExpression production.
	ExitLongitudeExpression(c *LongitudeExpressionContext)

	// ExitUrlExpression is called when exiting the urlExpression production.
	ExitUrlExpression(c *UrlExpressionContext)

	// ExitNameExpression is called when exiting the nameExpression production.
	ExitNameExpression(c *NameExpressionContext)

	// ExitTargetExpression is called when exiting the targetExpression production.
	ExitTargetExpression(c *TargetExpressionContext)

	// ExitLogicalExpression is called when exiting the logicalExpression production.
	ExitLogicalExpression(c *LogicalExpressionContext)

	// ExitIfTrueExpression is called when exiting the ifTrueExpression production.
	ExitIfTrueExpression(c *IfTrueExpressionContext)

	// ExitIfFalseExpression is called when exiting the ifFalseExpression production.
	ExitIfFalseExpression(c *IfFalseExpressionContext)

	// ExitHeightExpression is called when exiting the heightExpression production.
	ExitHeightExpression(c *HeightExpressionContext)

	// ExitWidthExpression is called when exiting the widthExpression production.
	ExitWidthExpression(c *WidthExpressionContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitNumChars is called when exiting the numChars production.
	ExitNumChars(c *NumCharsContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)
}
