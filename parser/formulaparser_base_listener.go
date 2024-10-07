// Code generated from ./FormulaParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // FormulaParser
import "github.com/antlr4-go/antlr/v4"

// BaseFormulaParserListener is a complete listener for a parse tree produced by FormulaParser.
type BaseFormulaParserListener struct{}

var _ FormulaParserListener = &BaseFormulaParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFormulaParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFormulaParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFormulaParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFormulaParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseFormulaParserListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseFormulaParserListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseFormulaParserListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseFormulaParserListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterNegationExpression is called when production negationExpression is entered.
func (s *BaseFormulaParserListener) EnterNegationExpression(ctx *NegationExpressionContext) {}

// ExitNegationExpression is called when production negationExpression is exited.
func (s *BaseFormulaParserListener) ExitNegationExpression(ctx *NegationExpressionContext) {}

// EnterCompareExpression is called when production compareExpression is entered.
func (s *BaseFormulaParserListener) EnterCompareExpression(ctx *CompareExpressionContext) {}

// ExitCompareExpression is called when production compareExpression is exited.
func (s *BaseFormulaParserListener) ExitCompareExpression(ctx *CompareExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseFormulaParserListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseFormulaParserListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterLogicExpression is called when production logicExpression is entered.
func (s *BaseFormulaParserListener) EnterLogicExpression(ctx *LogicExpressionContext) {}

// ExitLogicExpression is called when production logicExpression is exited.
func (s *BaseFormulaParserListener) ExitLogicExpression(ctx *LogicExpressionContext) {}

// EnterFunctionCallExpression is called when production functionCallExpression is entered.
func (s *BaseFormulaParserListener) EnterFunctionCallExpression(ctx *FunctionCallExpressionContext) {}

// ExitFunctionCallExpression is called when production functionCallExpression is exited.
func (s *BaseFormulaParserListener) ExitFunctionCallExpression(ctx *FunctionCallExpressionContext) {}

// EnterFieldReference is called when production fieldReference is entered.
func (s *BaseFormulaParserListener) EnterFieldReference(ctx *FieldReferenceContext) {}

// ExitFieldReference is called when production fieldReference is exited.
func (s *BaseFormulaParserListener) ExitFieldReference(ctx *FieldReferenceContext) {}

// EnterExponentiationExpression is called when production exponentiationExpression is entered.
func (s *BaseFormulaParserListener) EnterExponentiationExpression(ctx *ExponentiationExpressionContext) {
}

// ExitExponentiationExpression is called when production exponentiationExpression is exited.
func (s *BaseFormulaParserListener) ExitExponentiationExpression(ctx *ExponentiationExpressionContext) {
}

// EnterArithExpression is called when production arithExpression is entered.
func (s *BaseFormulaParserListener) EnterArithExpression(ctx *ArithExpressionContext) {}

// ExitArithExpression is called when production arithExpression is exited.
func (s *BaseFormulaParserListener) ExitArithExpression(ctx *ArithExpressionContext) {}

// EnterConcatExpression is called when production concatExpression is entered.
func (s *BaseFormulaParserListener) EnterConcatExpression(ctx *ConcatExpressionContext) {}

// ExitConcatExpression is called when production concatExpression is exited.
func (s *BaseFormulaParserListener) ExitConcatExpression(ctx *ConcatExpressionContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseFormulaParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseFormulaParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterAbs is called when production abs is entered.
func (s *BaseFormulaParserListener) EnterAbs(ctx *AbsContext) {}

// ExitAbs is called when production abs is exited.
func (s *BaseFormulaParserListener) ExitAbs(ctx *AbsContext) {}

// EnterAddMonths is called when production addMonths is entered.
func (s *BaseFormulaParserListener) EnterAddMonths(ctx *AddMonthsContext) {}

// ExitAddMonths is called when production addMonths is exited.
func (s *BaseFormulaParserListener) ExitAddMonths(ctx *AddMonthsContext) {}

// EnterAnd is called when production and is entered.
func (s *BaseFormulaParserListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production and is exited.
func (s *BaseFormulaParserListener) ExitAnd(ctx *AndContext) {}

// EnterBegins is called when production begins is entered.
func (s *BaseFormulaParserListener) EnterBegins(ctx *BeginsContext) {}

// ExitBegins is called when production begins is exited.
func (s *BaseFormulaParserListener) ExitBegins(ctx *BeginsContext) {}

// EnterBlankvalue is called when production blankvalue is entered.
func (s *BaseFormulaParserListener) EnterBlankvalue(ctx *BlankvalueContext) {}

// ExitBlankvalue is called when production blankvalue is exited.
func (s *BaseFormulaParserListener) ExitBlankvalue(ctx *BlankvalueContext) {}

// EnterBr is called when production br is entered.
func (s *BaseFormulaParserListener) EnterBr(ctx *BrContext) {}

// ExitBr is called when production br is exited.
func (s *BaseFormulaParserListener) ExitBr(ctx *BrContext) {}

// EnterCase is called when production case is entered.
func (s *BaseFormulaParserListener) EnterCase(ctx *CaseContext) {}

// ExitCase is called when production case is exited.
func (s *BaseFormulaParserListener) ExitCase(ctx *CaseContext) {}

// EnterCasesafeid is called when production casesafeid is entered.
func (s *BaseFormulaParserListener) EnterCasesafeid(ctx *CasesafeidContext) {}

// ExitCasesafeid is called when production casesafeid is exited.
func (s *BaseFormulaParserListener) ExitCasesafeid(ctx *CasesafeidContext) {}

// EnterCeiling is called when production ceiling is entered.
func (s *BaseFormulaParserListener) EnterCeiling(ctx *CeilingContext) {}

// ExitCeiling is called when production ceiling is exited.
func (s *BaseFormulaParserListener) ExitCeiling(ctx *CeilingContext) {}

// EnterContains is called when production contains is entered.
func (s *BaseFormulaParserListener) EnterContains(ctx *ContainsContext) {}

// ExitContains is called when production contains is exited.
func (s *BaseFormulaParserListener) ExitContains(ctx *ContainsContext) {}

// EnterCurrencyrate is called when production currencyrate is entered.
func (s *BaseFormulaParserListener) EnterCurrencyrate(ctx *CurrencyrateContext) {}

// ExitCurrencyrate is called when production currencyrate is exited.
func (s *BaseFormulaParserListener) ExitCurrencyrate(ctx *CurrencyrateContext) {}

// EnterDate is called when production date is entered.
func (s *BaseFormulaParserListener) EnterDate(ctx *DateContext) {}

// ExitDate is called when production date is exited.
func (s *BaseFormulaParserListener) ExitDate(ctx *DateContext) {}

// EnterDatevalue is called when production datevalue is entered.
func (s *BaseFormulaParserListener) EnterDatevalue(ctx *DatevalueContext) {}

// ExitDatevalue is called when production datevalue is exited.
func (s *BaseFormulaParserListener) ExitDatevalue(ctx *DatevalueContext) {}

// EnterDatetimevalue is called when production datetimevalue is entered.
func (s *BaseFormulaParserListener) EnterDatetimevalue(ctx *DatetimevalueContext) {}

// ExitDatetimevalue is called when production datetimevalue is exited.
func (s *BaseFormulaParserListener) ExitDatetimevalue(ctx *DatetimevalueContext) {}

// EnterDay is called when production day is entered.
func (s *BaseFormulaParserListener) EnterDay(ctx *DayContext) {}

// ExitDay is called when production day is exited.
func (s *BaseFormulaParserListener) ExitDay(ctx *DayContext) {}

// EnterDistance is called when production distance is entered.
func (s *BaseFormulaParserListener) EnterDistance(ctx *DistanceContext) {}

// ExitDistance is called when production distance is exited.
func (s *BaseFormulaParserListener) ExitDistance(ctx *DistanceContext) {}

// EnterExp is called when production exp is entered.
func (s *BaseFormulaParserListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BaseFormulaParserListener) ExitExp(ctx *ExpContext) {}

// EnterFind is called when production find is entered.
func (s *BaseFormulaParserListener) EnterFind(ctx *FindContext) {}

// ExitFind is called when production find is exited.
func (s *BaseFormulaParserListener) ExitFind(ctx *FindContext) {}

// EnterFloor is called when production floor is entered.
func (s *BaseFormulaParserListener) EnterFloor(ctx *FloorContext) {}

// ExitFloor is called when production floor is exited.
func (s *BaseFormulaParserListener) ExitFloor(ctx *FloorContext) {}

// EnterGeolocation is called when production geolocation is entered.
func (s *BaseFormulaParserListener) EnterGeolocation(ctx *GeolocationContext) {}

// ExitGeolocation is called when production geolocation is exited.
func (s *BaseFormulaParserListener) ExitGeolocation(ctx *GeolocationContext) {}

// EnterGetrecordids is called when production getrecordids is entered.
func (s *BaseFormulaParserListener) EnterGetrecordids(ctx *GetrecordidsContext) {}

// ExitGetrecordids is called when production getrecordids is exited.
func (s *BaseFormulaParserListener) ExitGetrecordids(ctx *GetrecordidsContext) {}

// EnterGetsessionid is called when production getsessionid is entered.
func (s *BaseFormulaParserListener) EnterGetsessionid(ctx *GetsessionidContext) {}

// ExitGetsessionid is called when production getsessionid is exited.
func (s *BaseFormulaParserListener) ExitGetsessionid(ctx *GetsessionidContext) {}

// EnterHour is called when production hour is entered.
func (s *BaseFormulaParserListener) EnterHour(ctx *HourContext) {}

// ExitHour is called when production hour is exited.
func (s *BaseFormulaParserListener) ExitHour(ctx *HourContext) {}

// EnterHtmlencode is called when production htmlencode is entered.
func (s *BaseFormulaParserListener) EnterHtmlencode(ctx *HtmlencodeContext) {}

// ExitHtmlencode is called when production htmlencode is exited.
func (s *BaseFormulaParserListener) ExitHtmlencode(ctx *HtmlencodeContext) {}

// EnterHyperlink is called when production hyperlink is entered.
func (s *BaseFormulaParserListener) EnterHyperlink(ctx *HyperlinkContext) {}

// ExitHyperlink is called when production hyperlink is exited.
func (s *BaseFormulaParserListener) ExitHyperlink(ctx *HyperlinkContext) {}

// EnterIf is called when production if is entered.
func (s *BaseFormulaParserListener) EnterIf(ctx *IfContext) {}

// ExitIf is called when production if is exited.
func (s *BaseFormulaParserListener) ExitIf(ctx *IfContext) {}

// EnterImage is called when production image is entered.
func (s *BaseFormulaParserListener) EnterImage(ctx *ImageContext) {}

// ExitImage is called when production image is exited.
func (s *BaseFormulaParserListener) ExitImage(ctx *ImageContext) {}

// EnterImageproxyurl is called when production imageproxyurl is entered.
func (s *BaseFormulaParserListener) EnterImageproxyurl(ctx *ImageproxyurlContext) {}

// ExitImageproxyurl is called when production imageproxyurl is exited.
func (s *BaseFormulaParserListener) ExitImageproxyurl(ctx *ImageproxyurlContext) {}

// EnterMid is called when production mid is entered.
func (s *BaseFormulaParserListener) EnterMid(ctx *MidContext) {}

// ExitMid is called when production mid is exited.
func (s *BaseFormulaParserListener) ExitMid(ctx *MidContext) {}

// EnterIncludes is called when production includes is entered.
func (s *BaseFormulaParserListener) EnterIncludes(ctx *IncludesContext) {}

// ExitIncludes is called when production includes is exited.
func (s *BaseFormulaParserListener) ExitIncludes(ctx *IncludesContext) {}

// EnterText is called when production text is entered.
func (s *BaseFormulaParserListener) EnterText(ctx *TextContext) {}

// ExitText is called when production text is exited.
func (s *BaseFormulaParserListener) ExitText(ctx *TextContext) {}

// EnterValue is called when production value is entered.
func (s *BaseFormulaParserListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseFormulaParserListener) ExitValue(ctx *ValueContext) {}

// EnterFieldExpression is called when production fieldExpression is entered.
func (s *BaseFormulaParserListener) EnterFieldExpression(ctx *FieldExpressionContext) {}

// ExitFieldExpression is called when production fieldExpression is exited.
func (s *BaseFormulaParserListener) ExitFieldExpression(ctx *FieldExpressionContext) {}

// EnterValueExpression is called when production valueExpression is entered.
func (s *BaseFormulaParserListener) EnterValueExpression(ctx *ValueExpressionContext) {}

// ExitValueExpression is called when production valueExpression is exited.
func (s *BaseFormulaParserListener) ExitValueExpression(ctx *ValueExpressionContext) {}

// EnterResultExpression is called when production resultExpression is entered.
func (s *BaseFormulaParserListener) EnterResultExpression(ctx *ResultExpressionContext) {}

// ExitResultExpression is called when production resultExpression is exited.
func (s *BaseFormulaParserListener) ExitResultExpression(ctx *ResultExpressionContext) {}

// EnterDefaultExpression is called when production defaultExpression is entered.
func (s *BaseFormulaParserListener) EnterDefaultExpression(ctx *DefaultExpressionContext) {}

// ExitDefaultExpression is called when production defaultExpression is exited.
func (s *BaseFormulaParserListener) ExitDefaultExpression(ctx *DefaultExpressionContext) {}

// EnterYearExpression is called when production yearExpression is entered.
func (s *BaseFormulaParserListener) EnterYearExpression(ctx *YearExpressionContext) {}

// ExitYearExpression is called when production yearExpression is exited.
func (s *BaseFormulaParserListener) ExitYearExpression(ctx *YearExpressionContext) {}

// EnterMonthExpression is called when production monthExpression is entered.
func (s *BaseFormulaParserListener) EnterMonthExpression(ctx *MonthExpressionContext) {}

// ExitMonthExpression is called when production monthExpression is exited.
func (s *BaseFormulaParserListener) ExitMonthExpression(ctx *MonthExpressionContext) {}

// EnterDayExpression is called when production dayExpression is entered.
func (s *BaseFormulaParserListener) EnterDayExpression(ctx *DayExpressionContext) {}

// ExitDayExpression is called when production dayExpression is exited.
func (s *BaseFormulaParserListener) ExitDayExpression(ctx *DayExpressionContext) {}

// EnterUnitExpression is called when production unitExpression is entered.
func (s *BaseFormulaParserListener) EnterUnitExpression(ctx *UnitExpressionContext) {}

// ExitUnitExpression is called when production unitExpression is exited.
func (s *BaseFormulaParserListener) ExitUnitExpression(ctx *UnitExpressionContext) {}

// EnterSearchExpression is called when production searchExpression is entered.
func (s *BaseFormulaParserListener) EnterSearchExpression(ctx *SearchExpressionContext) {}

// ExitSearchExpression is called when production searchExpression is exited.
func (s *BaseFormulaParserListener) ExitSearchExpression(ctx *SearchExpressionContext) {}

// EnterTextExpression is called when production textExpression is entered.
func (s *BaseFormulaParserListener) EnterTextExpression(ctx *TextExpressionContext) {}

// ExitTextExpression is called when production textExpression is exited.
func (s *BaseFormulaParserListener) ExitTextExpression(ctx *TextExpressionContext) {}

// EnterStartNum is called when production startNum is entered.
func (s *BaseFormulaParserListener) EnterStartNum(ctx *StartNumContext) {}

// ExitStartNum is called when production startNum is exited.
func (s *BaseFormulaParserListener) ExitStartNum(ctx *StartNumContext) {}

// EnterLatitudeExpression is called when production latitudeExpression is entered.
func (s *BaseFormulaParserListener) EnterLatitudeExpression(ctx *LatitudeExpressionContext) {}

// ExitLatitudeExpression is called when production latitudeExpression is exited.
func (s *BaseFormulaParserListener) ExitLatitudeExpression(ctx *LatitudeExpressionContext) {}

// EnterLongitudeExpression is called when production longitudeExpression is entered.
func (s *BaseFormulaParserListener) EnterLongitudeExpression(ctx *LongitudeExpressionContext) {}

// ExitLongitudeExpression is called when production longitudeExpression is exited.
func (s *BaseFormulaParserListener) ExitLongitudeExpression(ctx *LongitudeExpressionContext) {}

// EnterUrlExpression is called when production urlExpression is entered.
func (s *BaseFormulaParserListener) EnterUrlExpression(ctx *UrlExpressionContext) {}

// ExitUrlExpression is called when production urlExpression is exited.
func (s *BaseFormulaParserListener) ExitUrlExpression(ctx *UrlExpressionContext) {}

// EnterNameExpression is called when production nameExpression is entered.
func (s *BaseFormulaParserListener) EnterNameExpression(ctx *NameExpressionContext) {}

// ExitNameExpression is called when production nameExpression is exited.
func (s *BaseFormulaParserListener) ExitNameExpression(ctx *NameExpressionContext) {}

// EnterTargetExpression is called when production targetExpression is entered.
func (s *BaseFormulaParserListener) EnterTargetExpression(ctx *TargetExpressionContext) {}

// ExitTargetExpression is called when production targetExpression is exited.
func (s *BaseFormulaParserListener) ExitTargetExpression(ctx *TargetExpressionContext) {}

// EnterLogicalExpression is called when production logicalExpression is entered.
func (s *BaseFormulaParserListener) EnterLogicalExpression(ctx *LogicalExpressionContext) {}

// ExitLogicalExpression is called when production logicalExpression is exited.
func (s *BaseFormulaParserListener) ExitLogicalExpression(ctx *LogicalExpressionContext) {}

// EnterIfTrueExpression is called when production ifTrueExpression is entered.
func (s *BaseFormulaParserListener) EnterIfTrueExpression(ctx *IfTrueExpressionContext) {}

// ExitIfTrueExpression is called when production ifTrueExpression is exited.
func (s *BaseFormulaParserListener) ExitIfTrueExpression(ctx *IfTrueExpressionContext) {}

// EnterIfFalseExpression is called when production ifFalseExpression is entered.
func (s *BaseFormulaParserListener) EnterIfFalseExpression(ctx *IfFalseExpressionContext) {}

// ExitIfFalseExpression is called when production ifFalseExpression is exited.
func (s *BaseFormulaParserListener) ExitIfFalseExpression(ctx *IfFalseExpressionContext) {}

// EnterHeightExpression is called when production heightExpression is entered.
func (s *BaseFormulaParserListener) EnterHeightExpression(ctx *HeightExpressionContext) {}

// ExitHeightExpression is called when production heightExpression is exited.
func (s *BaseFormulaParserListener) ExitHeightExpression(ctx *HeightExpressionContext) {}

// EnterWidthExpression is called when production widthExpression is entered.
func (s *BaseFormulaParserListener) EnterWidthExpression(ctx *WidthExpressionContext) {}

// ExitWidthExpression is called when production widthExpression is exited.
func (s *BaseFormulaParserListener) ExitWidthExpression(ctx *WidthExpressionContext) {}

// EnterStart is called when production start is entered.
func (s *BaseFormulaParserListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseFormulaParserListener) ExitStart(ctx *StartContext) {}

// EnterNumChars is called when production numChars is entered.
func (s *BaseFormulaParserListener) EnterNumChars(ctx *NumCharsContext) {}

// ExitNumChars is called when production numChars is exited.
func (s *BaseFormulaParserListener) ExitNumChars(ctx *NumCharsContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseFormulaParserListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseFormulaParserListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseFormulaParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseFormulaParserListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseFormulaParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseFormulaParserListener) ExitLiteral(ctx *LiteralContext) {}
