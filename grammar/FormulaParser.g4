parser grammar FormulaParser;
options {tokenVocab=FormulaLexer;}

compilationUnit
    :   expression
    ;

expression
    :   primary                                                            # primaryExpression
    |   functionCall                                                       # functionCallExpression
    |   NOT expression                                                     # negationExpression
    |   BANG expression                                                    # negationExpression
    |   SUB expression                                                     # negativeExpression
    |   ADD expression                                                     # positiveExpression
    |   fieldReference                                                     # fieldReferenceExpression
    |   expression CARET expression                                        # exponentiationExpression
    |   expression (MUL | DIV) expression                                  # multiplicativeExpression
    |   expression (ADD | SUB) expression                                  # additiveExpression
    |   expression BITAND expression                                       # concatExpression
    |   expression (GT | LT) ASSIGN? expression                            # compareExpression
    |   expression (ASSIGN | EQUAL | NOTEQUAL | LESSANDGREATER) expression # equalityExpression
    |   expression (AND | OR) expression                                   # logicExpression
    |   BRACEBANG fieldReference RBRACE                                    # variableExpression
    ;

fieldReference
    : fieldPart (DOT fieldPart)*
    | fieldReference DOT VALUE
    ;

fieldPart
    : (Identifier | VALUE) (LBRACK expression RBRACK)?
    ;

functionCall
    : abs
    | acos
    | addMonths
    | and
    | ascii
    | asin
    | atan
    | atan2
    | begins
    | blankvalue
    | nullvalue
    | br
    | case
    | casesafeid
    | ceiling
    | chr
    | contains
    | cos
    | currencyrate
    | date
    | datevalue
    | datetimevalue
    | day
    | dayofyear
    | weekday
    | distance
    | exp
    | find
    | floor
    | formatduration
    | fromunixtime
    | geolocation
    | getrecordids
    | getsessionid
    | hour
    | htmlencode
    | hyperlink
    | if
    | image
    | imageproxyurl
    | initcap
    | ischanged
    | isnew
    | isoweek
    | isoyear
    | ispickval
    | includes
    | isblank
    | isnull
    | isnumber
    | left
    | len
    | ln
    | log
    | lpad
    | lower
    | max
    | mid
    | min
    | minute
    | mod
    | month
    | not
    | now
    | or
    | pi
    | picklistcount
    | priorvalue
    | regex
    | right
    | round
    | rpad
    | sin
    | sqrt
    | substitute
    | tan
    | trim
    | trunc
    | text
    | timevalue
    | today
    | unixtimestamp
    | upper
    | value
    | year

    ;

abs : ABS LPAREN expression RPAREN ;
acos : ACOS LPAREN expression RPAREN ;
addMonths : ADDMONTHS LPAREN dateExpression COMMA num RPAREN ;
ascii : ASCII LPAREN expression RPAREN ;
asin : ASIN LPAREN expression RPAREN ;
atan : ATAN LPAREN expression RPAREN ;
atan2 : ATAN2 LPAREN yExpression COMMA xExpression RPAREN ;
and : AND_FUNC LPAREN expression (COMMA expression)* RPAREN ;
begins : BEGINS LPAREN textExpression COMMA compareText RPAREN ;
blankvalue : BLANKVALUE LPAREN expression COMMA substituteValue RPAREN ;
nullvalue : NULLVALUE LPAREN expression COMMA substituteValue RPAREN ;
// TODO: only allow BR() without without whitespace within parentheses and
// surrounded by concatenation operators?
br : BR LPAREN RPAREN ;
case : CASE LPAREN expression (COMMA valueExpression COMMA resultExpression)+ COMMA defaultExpression RPAREN ;
casesafeid : CASESAFEID LPAREN expression RPAREN ;
ceiling : CEILING LPAREN expression RPAREN ;
chr : CHR LPAREN expression RPAREN ;
contains : CONTAINS LPAREN textExpression COMMA compareText RPAREN ;
cos : COS LPAREN expression RPAREN ;
currencyrate : CURRENCYRATE LPAREN expression RPAREN ;
date : DATE LPAREN yearExpression COMMA monthExpression COMMA dayExpression RPAREN ;
datevalue : DATEVALUE LPAREN expression RPAREN ;
datetimevalue : DATETIMEVALUE LPAREN expression RPAREN ;
day : DAY LPAREN expression RPAREN ;
dayofyear : DAYOFYEAR LPAREN expression RPAREN ;
weekday : WEEKDAY LPAREN expression RPAREN ;
distance : DISTANCE LPAREN expression COMMA expression COMMA unitExpression RPAREN ;
exp : EXP LPAREN expression RPAREN ;
find : FIND LPAREN searchExpression COMMA textExpression (COMMA startNum)? RPAREN ;
floor : FLOOR LPAREN expression RPAREN ;
formatduration : FORMATDURATION LPAREN expression (COMMA expression)? RPAREN ;
fromunixtime : FROMUNIXTIME LPAREN expression RPAREN ;
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
initcap : INITCAP LPAREN expression RPAREN ;
includes: INCLUDES LPAREN fieldExpression COMMA valueExpression RPAREN ;
isblank : ISBLANK LPAREN expression RPAREN ;
ischanged : ISCHANGED LPAREN fieldExpression RPAREN ;
isnew : ISNEW LPAREN RPAREN ;
isoweek : ISOWEEK LPAREN expression RPAREN ;
isoyear : ISOYEAR LPAREN expression RPAREN ;
isnull : ISNULL LPAREN expression RPAREN ;
isnumber : ISNUMBER LPAREN expression RPAREN ;
ispickval: ISPICKVAL LPAREN fieldExpression COMMA valueExpression RPAREN ;
left : LEFT LPAREN textExpression COMMA numChars RPAREN ;
len : LEN LPAREN expression RPAREN ;
ln : LN LPAREN expression RPAREN ;
log : LOG LPAREN expression RPAREN ;
lower : LOWER LPAREN expression RPAREN ;
lpad : LPAD LPAREN textExpression COMMA length (COMMA padString)? RPAREN ;
max : MAX LPAREN expression COMMA expression (COMMA expression)* RPAREN ;
mid : MID LPAREN textExpression COMMA startNum COMMA numChars RPAREN ;
min : MIN LPAREN expression COMMA expression (COMMA expression)* RPAREN ;
minute : MINUTE LPAREN expression RPAREN ;
mod : MOD LPAREN num COMMA divisor RPAREN ;
month : MONTH LPAREN expression RPAREN ;
not : NOT LPAREN expression RPAREN ;
now : NOW LPAREN RPAREN ;
or : OR_FUNC LPAREN expression (COMMA expression)* RPAREN ;
pi : PI LPAREN RPAREN ;
picklistcount : PICKLISTCOUNT LPAREN fieldExpression RPAREN ;
priorvalue : PRIORVALUE LPAREN fieldExpression RPAREN ;
regex : REGEX LPAREN textExpression COMMA regexExpression RPAREN ;
right : RIGHT LPAREN textExpression COMMA numChars RPAREN ;
round : ROUND LPAREN num COMMA digits RPAREN ;
rpad : RPAD LPAREN textExpression COMMA length (COMMA padString)? RPAREN ;
sin : SIN LPAREN expression RPAREN ;
sqrt : SQRT LPAREN expression RPAREN ;
substitute : SUBSTITUTE LPAREN textExpression COMMA oldText COMMA replacement RPAREN ;
tan : TAN LPAREN expression RPAREN ;
trim : TRIM LPAREN expression RPAREN ;
trunc : TRUNC LPAREN num COMMA digits RPAREN ;
text : TEXT LPAREN expression RPAREN ;
timevalue : TIMEVALUE LPAREN expression RPAREN ;
today : TODAY LPAREN RPAREN ;
unixtimestamp : UNIXTIMESTAMP LPAREN expression RPAREN ;
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

yExpression : expression ;
xExpression : expression ;

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
