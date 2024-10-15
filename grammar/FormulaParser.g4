parser grammar FormulaParser;
options {tokenVocab=FormulaLexer;}

compilationUnit
    :   expression
    ;

expression
    :   primary                                                            # primaryExpression
    |   functionCall                                                       # functionCallExpression
    |   (BANG | SUB) expression                                            # negationExpression
    |   ADD expression                                                     # positiveExpression
    |   Identifier (DOT Identifier)*                                       # fieldReference
    |   expression CARET expression                                        # exponentiationExpression
    |   expression BITAND expression                                       # concatExpression
    |   expression (ADD | SUB | DIV | MUL) expression                      # arithExpression
    |   expression (GT | LT) ASSIGN? expression                            # compareExpression
    |   expression (ASSIGN | EQUAL | NOTEQUAL | LESSANDGREATER) expression # equalityExpression
    |   expression (AND | OR) expression                                   # logicExpression
    ;


functionCall
    : abs
    | addMonths
    | and
    | begins
    | blankvalue
    | br
    | case
    | casesafeid
    | ceiling
    | contains
    | currencyrate
    | date
    | datevalue
    | datetimevalue
    | day
    | distance
    | exp
    | find
    | floor
    | geolocation
    | getrecordids
    | getsessionid
    | hour
    | htmlencode
    | hyperlink
    | if
    | image
    | imageproxyurl
    | ischanged
    | isnew
    | ispickval
    | includes
    | isblank
    | isnull
    | isnumber
    | left
    | len
    | lpad
    | lower
    | max
    | mid
    | min
    | mod
    | month
    | not
    | now
    | or
    | priorvalue
    | regex
    | right
    | round
    | substitute
    | trim
    | text
    | today
    | upper
    | value
    | year

    ;

abs : ABS LPAREN expression RPAREN ;
addMonths : ADDMONTHS LPAREN dateExpression COMMA num RPAREN ;
and : AND_FUNC LPAREN expression (COMMA expression)* RPAREN ;
begins : BEGINS LPAREN textExpression COMMA compareText RPAREN ;
blankvalue : BLANKVALUE LPAREN expression COMMA substituteValue RPAREN ;
// TODO: only allow BR() without without whitespace within parentheses and
// surrounded by concatenation operators?
br : BR LPAREN RPAREN ;
case : CASE LPAREN expression (COMMA valueExpression COMMA resultExpression)+ COMMA defaultExpression RPAREN ;
casesafeid : CASESAFEID LPAREN expression RPAREN ;
ceiling : CEILING LPAREN expression RPAREN ;
contains : CONTAINS LPAREN textExpression COMMA compareText RPAREN ;
currencyrate : CURRENCYRATE LPAREN expression RPAREN ;
date : DATE LPAREN yearExpression COMMA monthExpression COMMA dayExpression RPAREN ;
datevalue : DATEVALUE LPAREN expression RPAREN ;
datetimevalue : DATETIMEVALUE LPAREN expression RPAREN ;
day : DAY LPAREN expression RPAREN ;
distance : DISTANCE LPAREN expression COMMA expression COMMA unitExpression RPAREN ;
exp : EXP LPAREN expression RPAREN ;
find : FIND LPAREN searchExpression COMMA textExpression (COMMA startNum)? RPAREN ;
floor : FLOOR LPAREN expression RPAREN ;
geolocation : GEOLOCATION LPAREN latitudeExpression COMMA longitudeExpression RPAREN ;
// TOOD: Only allow literals, $SObjectType.<object>?
getrecordids : GETRECORDIDS LPAREN expression RPAREN ;
getsessionid : GETSESSIONID LPAREN RPAREN ;
hour : HOUR LPAREN expression RPAREN ;
htmlencode : HTMLENCODE LPAREN expression RPAREN ;
hyperlink : HYPERLINK LPAREN url COMMA name (COMMA target)? RPAREN ;
if : IF LPAREN logicalExpression COMMA ifTrueExpression COMMA ifFalseExpression RPAREN ;
image : IMAGE LPAREN url COMMA textExpression (COMMA heightExpression COMMA widthExpression)? RPAREN ;
imageproxyurl : IMAGEPROXYURL LPAREN url RPAREN ;
includes: INCLUDES LPAREN fieldExpression COMMA valueExpression RPAREN ;
isblank : ISBLANK LPAREN expression RPAREN ;
ischanged : ISCHANGED LPAREN fieldExpression RPAREN ;
isnew : ISNEW LPAREN RPAREN ;
isnull : ISNULL LPAREN expression RPAREN ;
isnumber : ISNUMBER LPAREN expression RPAREN ;
ispickval: ISPICKVAL LPAREN fieldExpression COMMA valueExpression RPAREN ;
left : LEFT LPAREN textExpression COMMA numChars RPAREN ;
len : LEN LPAREN expression RPAREN ;
lower : LOWER LPAREN expression RPAREN ;
lpad : LPAD LPAREN textExpression COMMA length (COMMA padString)? RPAREN ;
max : MAX LPAREN expression COMMA expression (COMMA expression)* RPAREN ;
mid : MID LPAREN textExpression COMMA startNum COMMA numChars RPAREN ;
min : MIN LPAREN expression COMMA expression (COMMA expression)* RPAREN ;
mod : MOD LPAREN num COMMA divisor RPAREN ;
month : MONTH LPAREN expression RPAREN ;
not : NOT LPAREN expression RPAREN ;
now : NOW LPAREN RPAREN ;
or : OR_FUNC LPAREN expression (COMMA expression)* RPAREN ;
priorvalue : PRIORVALUE LPAREN fieldExpression RPAREN ;
regex : REGEX LPAREN textExpression COMMA regexExpression RPAREN ;
right : RIGHT LPAREN textExpression COMMA numChars RPAREN ;
round : ROUND LPAREN num COMMA digits RPAREN ;
substitute : SUBSTITUTE LPAREN textExpression COMMA oldText COMMA replacement RPAREN ;
trim : TRIM LPAREN expression RPAREN ;
text : TEXT LPAREN expression RPAREN ;
today : TODAY LPAREN RPAREN ;
upper : UPPER LPAREN expression RPAREN ;
value : VALUE LPAREN expression RPAREN ;
year : YEAR LPAREN expression RPAREN ;


fieldExpression : expression ;

dateExpression : expression ;

valueExpression : expression ;
resultExpression : expression ;
substituteValue : expression ;
defaultExpression : expression ;

yearExpression : expression ;
monthExpression : expression ;
dayExpression : expression ;

regexExpression : expression ;

unitExpression : expression ;

searchExpression : expression ;
textExpression : expression ;

oldText : expression ;
replacement : expression ;

startNum : expression ;

compareText: expression ;

num : expression ;
divisor : expression ;
digits : expression ;

latitudeExpression : expression ;
longitudeExpression : expression ;

url : expression ;
name : expression ;
target : expression ;

logicalExpression : expression ;
ifTrueExpression : expression ;
ifFalseExpression : expression ;

length : expression ;
padString : expression ;

heightExpression : expression ;
widthExpression : expression ;

start : expression ;
numChars : expression ;

primary
    :   literal
    |   LPAREN expression RPAREN
    ;

literal
    :   StringLiteral
    |   IntegerLiteral
    |   FloatingPointLiteral
    |   BooleanLiteral
    |   NullLiteral
    ;
