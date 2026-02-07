package parser

import "github.com/antlr4-go/antlr/v4"

// ClearFormulaParserDFA clears the DFA cache for the Formula parser.
// This can free significant memory after formula parsing is complete.
// Note: This should only be called when no parsing is in progress.
func ClearFormulaParserDFA() {
	staticData := &FormulaParserParserStaticData
	if staticData.atn == nil || staticData.decisionToDFA == nil {
		return
	}
	// Clear each DFA in the array by replacing with fresh empty DFA
	atn := staticData.atn
	for i := range staticData.decisionToDFA {
		staticData.decisionToDFA[i] = antlr.NewDFA(atn.DecisionToState[i], i)
	}
	// Clear the prediction context cache which can also grow during parsing
	if staticData.PredictionContextCache != nil {
		staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	}
}

// ClearFormulaLexerDFA clears the DFA cache for the Formula lexer.
func ClearFormulaLexerDFA() {
	staticData := &FormulaLexerLexerStaticData
	if staticData.atn == nil || staticData.decisionToDFA == nil {
		return
	}
	// Clear each DFA in the array by replacing with fresh empty DFA
	atn := staticData.atn
	for i := range staticData.decisionToDFA {
		staticData.decisionToDFA[i] = antlr.NewDFA(atn.DecisionToState[i], i)
	}
	// Clear the prediction context cache which can also grow during parsing
	if staticData.PredictionContextCache != nil {
		staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	}
}
