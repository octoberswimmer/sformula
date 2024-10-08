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

	// EnterIncludes is called when entering the includes production.
	EnterIncludes(c *IncludesContext)

	// EnterIsblank is called when entering the isblank production.
	EnterIsblank(c *IsblankContext)

	// EnterIsnull is called when entering the isnull production.
	EnterIsnull(c *IsnullContext)

	// EnterIsnumber is called when entering the isnumber production.
	EnterIsnumber(c *IsnumberContext)

	// EnterIspickval is called when entering the ispickval production.
	EnterIspickval(c *IspickvalContext)

	// EnterLeft is called when entering the left production.
	EnterLeft(c *LeftContext)

	// EnterLen is called when entering the len production.
	EnterLen(c *LenContext)

	// EnterLpad is called when entering the lpad production.
	EnterLpad(c *LpadContext)

	// EnterMid is called when entering the mid production.
	EnterMid(c *MidContext)

	// EnterMin is called when entering the min production.
	EnterMin(c *MinContext)

	// EnterMod is called when entering the mod production.
	EnterMod(c *ModContext)

	// EnterMonth is called when entering the month production.
	EnterMonth(c *MonthContext)

	// EnterNot is called when entering the not production.
	EnterNot(c *NotContext)

	// EnterNow is called when entering the now production.
	EnterNow(c *NowContext)

	// EnterOr is called when entering the or production.
	EnterOr(c *OrContext)

	// EnterRight is called when entering the right production.
	EnterRight(c *RightContext)

	// EnterRound is called when entering the round production.
	EnterRound(c *RoundContext)

	// EnterSubstitute is called when entering the substitute production.
	EnterSubstitute(c *SubstituteContext)

	// EnterTrim is called when entering the trim production.
	EnterTrim(c *TrimContext)

	// EnterText is called when entering the text production.
	EnterText(c *TextContext)

	// EnterToday is called when entering the today production.
	EnterToday(c *TodayContext)

	// EnterUpper is called when entering the upper production.
	EnterUpper(c *UpperContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterYear is called when entering the year production.
	EnterYear(c *YearContext)

	// EnterFieldExpression is called when entering the fieldExpression production.
	EnterFieldExpression(c *FieldExpressionContext)

	// EnterValueExpression is called when entering the valueExpression production.
	EnterValueExpression(c *ValueExpressionContext)

	// EnterResultExpression is called when entering the resultExpression production.
	EnterResultExpression(c *ResultExpressionContext)

	// EnterSubstituteValue is called when entering the substituteValue production.
	EnterSubstituteValue(c *SubstituteValueContext)

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

	// EnterOldText is called when entering the oldText production.
	EnterOldText(c *OldTextContext)

	// EnterReplacement is called when entering the replacement production.
	EnterReplacement(c *ReplacementContext)

	// EnterStartNum is called when entering the startNum production.
	EnterStartNum(c *StartNumContext)

	// EnterCompareText is called when entering the compareText production.
	EnterCompareText(c *CompareTextContext)

	// EnterNum is called when entering the num production.
	EnterNum(c *NumContext)

	// EnterDivisor is called when entering the divisor production.
	EnterDivisor(c *DivisorContext)

	// EnterDigits is called when entering the digits production.
	EnterDigits(c *DigitsContext)

	// EnterLatitudeExpression is called when entering the latitudeExpression production.
	EnterLatitudeExpression(c *LatitudeExpressionContext)

	// EnterLongitudeExpression is called when entering the longitudeExpression production.
	EnterLongitudeExpression(c *LongitudeExpressionContext)

	// EnterUrl is called when entering the url production.
	EnterUrl(c *UrlContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterTarget is called when entering the target production.
	EnterTarget(c *TargetContext)

	// EnterLogicalExpression is called when entering the logicalExpression production.
	EnterLogicalExpression(c *LogicalExpressionContext)

	// EnterIfTrueExpression is called when entering the ifTrueExpression production.
	EnterIfTrueExpression(c *IfTrueExpressionContext)

	// EnterIfFalseExpression is called when entering the ifFalseExpression production.
	EnterIfFalseExpression(c *IfFalseExpressionContext)

	// EnterLength is called when entering the length production.
	EnterLength(c *LengthContext)

	// EnterPadString is called when entering the padString production.
	EnterPadString(c *PadStringContext)

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

	// ExitIncludes is called when exiting the includes production.
	ExitIncludes(c *IncludesContext)

	// ExitIsblank is called when exiting the isblank production.
	ExitIsblank(c *IsblankContext)

	// ExitIsnull is called when exiting the isnull production.
	ExitIsnull(c *IsnullContext)

	// ExitIsnumber is called when exiting the isnumber production.
	ExitIsnumber(c *IsnumberContext)

	// ExitIspickval is called when exiting the ispickval production.
	ExitIspickval(c *IspickvalContext)

	// ExitLeft is called when exiting the left production.
	ExitLeft(c *LeftContext)

	// ExitLen is called when exiting the len production.
	ExitLen(c *LenContext)

	// ExitLpad is called when exiting the lpad production.
	ExitLpad(c *LpadContext)

	// ExitMid is called when exiting the mid production.
	ExitMid(c *MidContext)

	// ExitMin is called when exiting the min production.
	ExitMin(c *MinContext)

	// ExitMod is called when exiting the mod production.
	ExitMod(c *ModContext)

	// ExitMonth is called when exiting the month production.
	ExitMonth(c *MonthContext)

	// ExitNot is called when exiting the not production.
	ExitNot(c *NotContext)

	// ExitNow is called when exiting the now production.
	ExitNow(c *NowContext)

	// ExitOr is called when exiting the or production.
	ExitOr(c *OrContext)

	// ExitRight is called when exiting the right production.
	ExitRight(c *RightContext)

	// ExitRound is called when exiting the round production.
	ExitRound(c *RoundContext)

	// ExitSubstitute is called when exiting the substitute production.
	ExitSubstitute(c *SubstituteContext)

	// ExitTrim is called when exiting the trim production.
	ExitTrim(c *TrimContext)

	// ExitText is called when exiting the text production.
	ExitText(c *TextContext)

	// ExitToday is called when exiting the today production.
	ExitToday(c *TodayContext)

	// ExitUpper is called when exiting the upper production.
	ExitUpper(c *UpperContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitYear is called when exiting the year production.
	ExitYear(c *YearContext)

	// ExitFieldExpression is called when exiting the fieldExpression production.
	ExitFieldExpression(c *FieldExpressionContext)

	// ExitValueExpression is called when exiting the valueExpression production.
	ExitValueExpression(c *ValueExpressionContext)

	// ExitResultExpression is called when exiting the resultExpression production.
	ExitResultExpression(c *ResultExpressionContext)

	// ExitSubstituteValue is called when exiting the substituteValue production.
	ExitSubstituteValue(c *SubstituteValueContext)

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

	// ExitOldText is called when exiting the oldText production.
	ExitOldText(c *OldTextContext)

	// ExitReplacement is called when exiting the replacement production.
	ExitReplacement(c *ReplacementContext)

	// ExitStartNum is called when exiting the startNum production.
	ExitStartNum(c *StartNumContext)

	// ExitCompareText is called when exiting the compareText production.
	ExitCompareText(c *CompareTextContext)

	// ExitNum is called when exiting the num production.
	ExitNum(c *NumContext)

	// ExitDivisor is called when exiting the divisor production.
	ExitDivisor(c *DivisorContext)

	// ExitDigits is called when exiting the digits production.
	ExitDigits(c *DigitsContext)

	// ExitLatitudeExpression is called when exiting the latitudeExpression production.
	ExitLatitudeExpression(c *LatitudeExpressionContext)

	// ExitLongitudeExpression is called when exiting the longitudeExpression production.
	ExitLongitudeExpression(c *LongitudeExpressionContext)

	// ExitUrl is called when exiting the url production.
	ExitUrl(c *UrlContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitTarget is called when exiting the target production.
	ExitTarget(c *TargetContext)

	// ExitLogicalExpression is called when exiting the logicalExpression production.
	ExitLogicalExpression(c *LogicalExpressionContext)

	// ExitIfTrueExpression is called when exiting the ifTrueExpression production.
	ExitIfTrueExpression(c *IfTrueExpressionContext)

	// ExitIfFalseExpression is called when exiting the ifFalseExpression production.
	ExitIfFalseExpression(c *IfFalseExpressionContext)

	// ExitLength is called when exiting the length production.
	ExitLength(c *LengthContext)

	// ExitPadString is called when exiting the padString production.
	ExitPadString(c *PadStringContext)

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
