package formatter

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/sformula/parser"

	log "github.com/sirupsen/logrus"
)

type testErrorListener struct {
	*antlr.DefaultErrorListener
	t *testing.T
}

func TestCompilationUnit(t *testing.T) {
	if testing.Verbose() {
		log.SetLevel(log.DebugLevel)

	}
	tests :=
		[]struct {
			input  string
			output string
		}{
			{
				// Single arguments are allowed to AND
				`IF(AND(true), true, false)`,
				`IF(AND(true), true, false)`,
			},
			{
				`IF(OR(true), true, false)`,
				`IF(OR(true), true, false)`,
			},
			{
				`AND(true,false)`,
				`AND(true, false)`,
			},
			{
				`AND(!true,false)`,
				`AND(!true, false)`,
			},
			{
				`AND(true,true,false)`,
				`AND(true, true, false)`,
			},
			{
				`AND(AND(true,true),false)`,
				`AND(AND(true, true), false)`,
			},
			{
				`AND(true,true,false,true,false,true,false,true,false,true)`,
				`AND(
	true,
	true,
	false,
	true,
	false,
	true,
	false,
	true,
	false,
	true
)`,
			},
			{
				`GEOLOCATION(37.775,-122.418)`,
				`GEOLOCATION(37.775, -122.418)`,
			},
			{
				`DISTANCE(warehouse_location__c, GEOLOCATION(37.775,-122.418), 'km')`,
				`DISTANCE(
	warehouse_location__c,
	GEOLOCATION(37.775, -122.418),
	'km'
)`,
			},
			{
				`$Api.Session_ID`,
				`$Api.Session_ID`,
			},
			{
				`HYPERLINK("ymsgr:sendIM?" & Yahoo_Name__c, IMAGE("http://opi.yahoo.com/online?u=" & Yahoo_Name__c & "&m;=g&t;=0", "Yahoo"))`,
				`HYPERLINK(
	"ymsgr:sendIM?" & Yahoo_Name__c,
	IMAGE("http://opi.yahoo.com/online?u=" & Yahoo_Name__c & "&m;=g&t;=0", "Yahoo")
)`,
			},
			{
				`Owner:User.Id = $User.Id`,
				`Owner:User.Id = $User.Id`,
			},
			{
				`OR((AND(NOT(ISBLANK(Next_Step_Start__c)), ISBLANK(Next_Step__c))), (AND(NOT(ISBLANK(Next_Step_End__c)), ISBLANK(Next_Step__c))))`,
				`OR(
	(AND(
		NOT(ISBLANK(Next_Step_Start__c)),
		ISBLANK(Next_Step__c)
	)),
	(AND(
		NOT(ISBLANK(Next_Step_End__c)),
		ISBLANK(Next_Step__c)
	))
)`,
			},
			{
				`1++++3--------5-+-+-++--2`,
				`1 + +++3 - -------5 - +-+-++--2`,
			},
			{
				`NOW() + (1/24)`,
				`NOW() + (1/24)`,
			},
			{
				`BLANKVALUE(Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Id,
				BLANKVALUE(Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Id,
				BLANKVALUE(Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Id,
				BLANKVALUE(Parent.Parent.Parent.Parent.Parent.Parent.Parent.Id,
				BLANKVALUE(Parent.Parent.Parent.Parent.Parent.Parent.Id,
				BLANKVALUE(Parent.Parent.Parent.Parent.Parent.Id,
				BLANKVALUE(Parent.Parent.Parent.Parent.Id,
				BLANKVALUE(Parent.Parent.Parent.Id,
				BLANKVALUE(Parent.Parent.Id,
				BLANKVALUE(Parent.Id,
				Id))))))))))`,
				`BLANKVALUE(
	Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Id,
	BLANKVALUE(
		Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Id,
		BLANKVALUE(
			Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Id,
			BLANKVALUE(
				Parent.Parent.Parent.Parent.Parent.Parent.Parent.Id,
				BLANKVALUE(
					Parent.Parent.Parent.Parent.Parent.Parent.Id,
					BLANKVALUE(
						Parent.Parent.Parent.Parent.Parent.Id,
						BLANKVALUE(
							Parent.Parent.Parent.Parent.Id,
							BLANKVALUE(
								Parent.Parent.Parent.Id,
								BLANKVALUE(
									Parent.Parent.Id,
									BLANKVALUE(Parent.Id, Id)
								)
							)
						)
					)
				)
			)
		)
	)
)`,
			},
		}
	for _, tt := range tests {
		input := antlr.NewInputStream(tt.input)
		lexer := parser.NewFormulaLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		p := parser.NewFormulaParser(stream)
		p.RemoveErrorListeners()
		p.AddErrorListener(&testErrorListener{t: t})

		v := NewFormatVisitor(stream)
		out, ok := v.visitRule(p.CompilationUnit()).(string)
		if !ok {
			t.Errorf("Unexpected result parsing formula")
		}
		if out != tt.output {
			t.Errorf("unexpected format.  expected:\n%s\ngot:\n%s\n", tt.output, out)
		}
	}
}

// Flow formulas allow variable expressions
func TestFlowFormulas(t *testing.T) {
	if testing.Verbose() {
		log.SetLevel(log.DebugLevel)

	}
	tests :=
		[]struct {
			input  string
			output string
		}{
			{
				`{!$Api.Session_ID}`,
				`{!$Api.Session_ID}`,
			},
		}
	for _, tt := range tests {
		out, err := FormatFlowFormula(tt.input)
		if err != nil {
			t.Errorf("error parsing formula: %s", err.Error())
		}
		if out != tt.output {
			t.Errorf("unexpected format.  expected:\n%s\ngot:\n%s\n", tt.output, out)
		}
	}
}
