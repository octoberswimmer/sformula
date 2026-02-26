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
				`NOT TRUE`,
				`NOT TRUE`,
			},
			{
				`NOT ISBLANK(Field__c)`,
				`NOT ISBLANK(Field__c)`,
			},
			{
				`NOT(ISBLANK(Field__c))`,
				`NOT(ISBLANK(Field__c))`,
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
			{
				// NULLVALUE function - simple
				`NULLVALUE(Name__c, "Unknown")`,
				`NULLVALUE(Name__c, "Unknown")`,
			},
			{
				// NULLVALUE function - with expression
				`NULLVALUE(Account.Name, "No Account")`,
				`NULLVALUE(Account.Name, "No Account")`,
			},
			{
				// NULLVALUE function - nested
				`NULLVALUE(Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Id,
			NULLVALUE(Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Id,
			NULLVALUE(Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Id,
			NULLVALUE(Parent.Parent.Parent.Parent.Parent.Parent.Parent.Id,
			NULLVALUE(Parent.Parent.Parent.Parent.Parent.Parent.Id,
			NULLVALUE(Parent.Parent.Parent.Parent.Parent.Id,
			NULLVALUE(Parent.Parent.Parent.Parent.Id,
			NULLVALUE(Parent.Parent.Parent.Id,
			NULLVALUE(Parent.Parent.Id,
			NULLVALUE(Parent.Id,
			Id))))))))))`,
				`NULLVALUE(
	Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Id,
	NULLVALUE(
		Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Id,
		NULLVALUE(
			Parent.Parent.Parent.Parent.Parent.Parent.Parent.Parent.Id,
			NULLVALUE(
				Parent.Parent.Parent.Parent.Parent.Parent.Parent.Id,
				NULLVALUE(
					Parent.Parent.Parent.Parent.Parent.Parent.Id,
					NULLVALUE(
						Parent.Parent.Parent.Parent.Parent.Id,
						NULLVALUE(
							Parent.Parent.Parent.Parent.Id,
							NULLVALUE(
								Parent.Parent.Parent.Id,
								NULLVALUE(
									Parent.Parent.Id,
									NULLVALUE(Parent.Id, Id)
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
			{
				// Array indexing support
				`items[0]`,
				`items[0]`,
			},
			{
				`items[0].Name`,
				`items[0].Name`,
			},
			{
				`clinics[0].Location__r.Name`,
				`clinics[0].Location__r.Name`,
			},
			{
				// Array indexing with complex expressions
				`items[index + 1].field`,
				`items[index + 1].field`,
			},
			{
				// Array indexing in functions
				`IF(items[0].Active__c, items[0].Name, "Inactive")`,
				`IF(items[0].Active__c, items[0].Name, "Inactive")`,
			},
			{
				// Array indexing with string concatenation
				`"Item: " & items[0].Name & " - " & items[0].Status__c`,
				`"Item: " & items[0].Name & " - " & items[0].Status__c`,
			},
			{
				// Complex Visualforce expression with array indexing
				`clinics[0].Location__r.SchedulingOwner__r.User__r.FirstName & ' ' & clinics[0].Location__r.SchedulingOwner__r.User__r.LastName`,
				`clinics[0].Location__r.SchedulingOwner__r.User__r.FirstName & ' ' & clinics[0].Location__r.SchedulingOwner__r.User__r.LastName`,
			},
			{
				// Array size check with IF
				`IF(clinics.size > 1, "clinics", "a clinic")`,
				`IF(clinics.size > 1, "clinics", "a clinic")`,
			},
			{
				// HOUR function test - simple
				`HOUR(NOW())`,
				`HOUR(NOW())`,
			},
			{
				// WEEKDAY function test - simple
				`WEEKDAY(TODAY())`,
				`WEEKDAY(TODAY())`,
			},
			{
				// HOUR function test - with field reference
				`HOUR(CreatedDate)`,
				`HOUR(CreatedDate)`,
			},
			{
				// HOUR function test - with expression
				`HOUR(NOW() + 1)`,
				`HOUR(NOW() + 1)`,
			},
			{
				// HOUR function test - wrapped due to length
				`HOUR(DATETIMEVALUE("2024-01-01 14:30:00") + (1/24))`,
				`HOUR(DATETIMEVALUE("2024-01-01 14:30:00") + (1/24))`,
			},
			{
				// HOUR function test - in IF condition
				`IF(HOUR(NOW()) > 12, "Afternoon", "Morning")`,
				`IF(HOUR(NOW()) > 12, "Afternoon", "Morning")`,
			},
			{
				// HOUR function test - comparison
				`HOUR(CreatedDate) >= 9 && HOUR(CreatedDate) < 17`,
				`HOUR(CreatedDate) >= 9 &&
	HOUR(CreatedDate) < 17`,
			},
			{
				// TIMEVALUE function test - simple
				`TIMEVALUE("14:30:00")`,
				`TIMEVALUE("14:30:00")`,
			},
			{
				// TIMEVALUE function test - with field reference
				`TIMEVALUE(TextField__c)`,
				`TIMEVALUE(TextField__c)`,
			},
			{
				// TIMEVALUE function test - with concatenation
				`TIMEVALUE(HOUR(NOW()) & ":00:00")`,
				`TIMEVALUE(HOUR(NOW()) & ":00:00")`,
			},
			{
				// TIMEVALUE function test - wrapped due to length
				`TIMEVALUE("12:00:00") + (1/24)`,
				`TIMEVALUE("12:00:00") + (1/24)`,
			},
			{
				// TIMEVALUE function test - in IF condition
				`IF(TIMEVALUE("09:00:00") < NOW(), "Past 9 AM", "Before 9 AM")`,
				`IF(TIMEVALUE("09:00:00") < NOW(), "Past 9 AM", "Before 9 AM")`,
			},
			{
				// MINUTE function test - simple
				`MINUTE(NOW())`,
				`MINUTE(NOW())`,
			},
			{
				// MINUTE function test - with field reference
				`MINUTE(CreatedDate)`,
				`MINUTE(CreatedDate)`,
			},
			{
				// MINUTE function test - with expression
				`MINUTE(NOW() + (30/1440))`,
				`MINUTE(NOW() + (30 / 1440))`,
			},
			{
				// MINUTE function test - wrapped due to length
				`MINUTE(DATETIMEVALUE("2024-01-01 14:30:45"))`,
				`MINUTE(DATETIMEVALUE("2024-01-01 14:30:45"))`,
			},
			{
				// MINUTE function test - in IF condition
				`IF(MINUTE(NOW()) >= 30, "Past half hour", "Before half hour")`,
				`IF(MINUTE(NOW()) >= 30, "Past half hour", "Before half hour")`,
			},
			{
				// MINUTE function test - comparison
				`MINUTE(StartTime__c) == 0 || MINUTE(StartTime__c) == 30`,
				`MINUTE(StartTime__c) == 0 ||
	MINUTE(StartTime__c) == 30`,
			},
			{
				// negativeExpression - simple negative number
				`-5`,
				`-5`,
			},
			{
				// negativeExpression - negative decimal
				`-3.14159`,
				`-3.14159`,
			},
			{
				// negativeExpression - negative field reference
				`-Amount__c`,
				`-Amount__c`,
			},
			{
				// negativeExpression - negative function result
				`-ABS(Amount__c)`,
				`-ABS(Amount__c)`,
			},
			{
				// negativeExpression - negative expression in parentheses
				`-(5 + 3)`,
				`-(5 + 3)`,
			},
			{
				// negativeExpression - double negative
				`--5`,
				`--5`,
			},
			{
				// negativeExpression - triple negative
				`---Amount__c`,
				`---Amount__c`,
			},
			{
				// negativeExpression - negative in arithmetic
				`10 + -5`,
				`10 + -5`,
			},
			{
				// negativeExpression - negative in multiplication
				`-3 * 4`,
				`-3 * 4`,
			},
			{
				// negativeExpression - negative in division
				`-10 / 2`,
				`-10 / 2`,
			},
			{
				// negativeExpression - negative in comparison
				`Amount__c > -100`,
				`Amount__c > -100`,
			},
			{
				// negativeExpression - negative in IF condition
				`IF(Amount__c < -50, "Negative", "Positive")`,
				`IF(Amount__c < -50, "Negative", "Positive")`,
			},
			{
				// negativeExpression - negative in function arguments
				`MAX(-10, -20, -5)`,
				`MAX(-10, -20, -5)`,
			},
			{
				// negativeExpression - complex negative expression
				`-(Amount__c * 0.1) + -TaxRate__c`,
				`-(Amount__c * 0.1) + -TaxRate__c`,
			},
			{
				// negativeExpression - negative with nested functions
				`-ROUND(Amount__c * 0.15, 2)`,
				`-ROUND(Amount__c * 0.15, 2)`,
			},
			{
				// negativeExpression - negative in CASE expression
				`CASE(Status__c, "Refund", -Amount__c, "Credit", -Amount__c/2, 0)`,
				`CASE(
	Status__c,
		"Refund", -Amount__c,
		"Credit", -Amount__c / 2,
	0
)`,
			},
			{
				// negativeExpression - negative with array indexing
				`-items[0].Amount__c`,
				`-items[0].Amount__c`,
			},
			{
				// negativeExpression - negative in complex arithmetic
				`(Price__c * -Discount__c) + -ShippingCost__c`,
				`(Price__c * -Discount__c) + -ShippingCost__c`,
			},
			{
				// LPAD function - simple
				`LPAD("ab", 5, "x")`,
				`LPAD("ab", 5, "x")`,
			},
			{
				// LPAD function - with field reference
				`LPAD(Name__c, 10, "0")`,
				`LPAD(Name__c, 10, "0")`,
			},
			{
				// LPAD function - without pad string (uses space)
				`LPAD(Name__c, 15)`,
				`LPAD(Name__c, 15)`,
			},
			{
				// LPAD function - with expression for length
				`LPAD(Code__c, MaxLength__c, "0")`,
				`LPAD(Code__c, MaxLength__c, "0")`,
			},
			{
				// RPAD function - simple
				`RPAD("ab", 5, "x")`,
				`RPAD("ab", 5, "x")`,
			},
			{
				// RPAD function - with field reference
				`RPAD(Name__c, 10, "-")`,
				`RPAD(Name__c, 10, "-")`,
			},
			{
				// RPAD function - without pad string (uses space)
				`RPAD(Name__c, 20)`,
				`RPAD(Name__c, 20)`,
			},
			{
				// RPAD function - with expression for length
				`RPAD(Description__c, MaxLength__c, ".")`,
				`RPAD(Description__c, MaxLength__c, ".")`,
			},
			{
				// LPAD function - in concatenation
				`LPAD(AccountNumber__c, 8, "0") & "-" & RPAD(Name__c, 20)`,
				`LPAD(AccountNumber__c, 8, "0") & "-" & RPAD(Name__c, 20)`,
			},
			// Trigonometric functions
			{
				`ACOS(0.5)`,
				`ACOS(0.5)`,
			},
			{
				`ASIN(0.5)`,
				`ASIN(0.5)`,
			},
			{
				`ATAN(1)`,
				`ATAN(1)`,
			},
			{
				`ATAN2(1, 2)`,
				`ATAN2(1, 2)`,
			},
			{
				`COS(0)`,
				`COS(0)`,
			},
			{
				`SIN(0)`,
				`SIN(0)`,
			},
			{
				`TAN(0)`,
				`TAN(0)`,
			},
			{
				`PI()`,
				`PI()`,
			},
			{
				// Trig functions with field references
				`COS(Angle__c) * Radius__c`,
				`COS(Angle__c) * Radius__c`,
			},
			{
				// ATAN2 with field references
				`ATAN2(Y__c, X__c)`,
				`ATAN2(Y__c, X__c)`,
			},
			// String functions
			{
				`ASCII("A")`,
				`ASCII("A")`,
			},
			{
				`CHR(65)`,
				`CHR(65)`,
			},
			{
				`INITCAP("hello world")`,
				`INITCAP("hello world")`,
			},
			{
				// INITCAP with field reference
				`INITCAP(Name__c)`,
				`INITCAP(Name__c)`,
			},
			// Date/Time functions
			{
				`DAYOFYEAR(TODAY())`,
				`DAYOFYEAR(TODAY())`,
			},
			{
				`ISOWEEK(TODAY())`,
				`ISOWEEK(TODAY())`,
			},
			{
				`ISOYEAR(TODAY())`,
				`ISOYEAR(TODAY())`,
			},
			{
				`UNIXTIMESTAMP(NOW())`,
				`UNIXTIMESTAMP(NOW())`,
			},
			{
				`FROMUNIXTIME(1234567890)`,
				`FROMUNIXTIME(1234567890)`,
			},
			{
				// DAYOFYEAR with field reference
				`DAYOFYEAR(CreatedDate)`,
				`DAYOFYEAR(CreatedDate)`,
			},
			{
				// ISOWEEK with field reference
				`ISOWEEK(CloseDate__c)`,
				`ISOWEEK(CloseDate__c)`,
			},
			// Numeric functions
			{
				`TRUNC(3.14159, 2)`,
				`TRUNC(3.14159, 2)`,
			},
			{
				`TRUNC(Amount__c, 0)`,
				`TRUNC(Amount__c, 0)`,
			},
			// Picklist functions
			{
				`PICKLISTCOUNT(MultiSelect__c)`,
				`PICKLISTCOUNT(MultiSelect__c)`,
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
			{
				// Visualforce expressions with array indexing
				`{!items[0]}`,
				`{!items[0]}`,
			},
			{
				`{!clinics[0].Location__r.Name}`,
				`{!clinics[0].Location__r.Name}`,
			},
			{
				`{!items[index].Name}`,
				`{!items[index].Name}`,
			},
			{
				// HOUR function in flow formula
				`HOUR({!$Flow.CurrentDateTime})`,
				`HOUR({!$Flow.CurrentDateTime})`,
			},
			{
				// HOUR function with flow variable in comparison
				`HOUR({!StartTime}) < 12`,
				`HOUR({!StartTime}) < 12`,
			},
			{
				// TIMEVALUE function in flow formula
				`TIMEVALUE({!$Flow.CurrentTime})`,
				`TIMEVALUE({!$Flow.CurrentTime})`,
			},
			{
				// TIMEVALUE function with flow variable
				`TIMEVALUE({!TimeString}) > TIMEVALUE("12:00:00")`,
				`TIMEVALUE({!TimeString}) > TIMEVALUE("12:00:00")`,
			},
			{
				// MINUTE function in flow formula
				`MINUTE({!$Flow.CurrentDateTime})`,
				`MINUTE({!$Flow.CurrentDateTime})`,
			},
			{
				// MINUTE function with flow variable in calculation
				`MINUTE({!AppointmentTime}) + 15`,
				`MINUTE({!AppointmentTime}) + 15`,
			},
			{
				// Test case sensitivity for .value (should preserve original case)
				`ISBLANK(TRIM({!Phone.value}))`,
				`ISBLANK(TRIM({!Phone.value}))`,
			},
			{
				// Test uppercase .VALUE (should preserve original case)
				`ISBLANK(TRIM({!Phone.VALUE}))`,
				`ISBLANK(TRIM({!Phone.VALUE}))`,
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
