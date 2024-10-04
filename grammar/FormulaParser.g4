parser grammar FormulaParser;
options {tokenVocab=FormulaLexer;}

compilationUnit
    :   expression
    ;

expression
    :   primary                                                   # primaryExpression
    |   functionCall                                              # functionCallExpression
    |   SUB expression                                            # negationExpression
    |   Identifier (DOT Identifier)*                              # fieldReference
    |   expression CARET expression                               # exponentiationExpression
    |   expression BITAND expression                              # concatExpression
    |   expression (ADD|SUB|DIV|MUL) expression                   # arithExpression
    |   expression (GT | LT) ASSIGN? expression                   # compareExpression
    |   expression (EQUAL | NOTEQUAL | LESSANDGREATER) expression # equalityExpression
    |   expression (AND | OR) expression                          # logicExpression
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

    ;

abs : ABS LPAREN expression RPAREN ;
addMonths : ADDMONTHS LPAREN expression COMMA expression RPAREN ;
and : AND_FUNC LPAREN expression COMMA expression (COMMA expression)* RPAREN ;
begins : BEGINS LPAREN expression COMMA expression RPAREN ;
blankvalue : BLANKVALUE LPAREN expression COMMA expression RPAREN ;
// TODO: only allow BR() without without whitespace within parentheses and
// surrounded by concatenation operators?
br : BR LPAREN RPAREN ;
case : CASE LPAREN expression COMMA valueExpression COMMA resultExpression (COMMA valueExpression COMMA resultExpression)* defaultExpression RPAREN ;
casesafeid : CASESAFEID LPAREN expression RPAREN ;
ceiling : CEILING LPAREN expression RPAREN ;
contains : CONTAINS LPAREN expression COMMA expression RPAREN ;
currencyrate : CURRENCYRATE LPAREN expression RPAREN ;
date : DATE LPAREN yearExpression COMMA monthExpression COMMA dayExpression RPAREN ;
datevalue : DATEVALUE LPAREN expression RPAREN ;
datetimevalue : DATETIMEVALUE LPAREN expression RPAREN ;
day : DAY LPAREN expression RPAREN ;
distance : DISTANCE LPAREN expression COMMA expression COMMA unitExpression RPAREN ;
exp : EXP LPAREN expression RPAREN ;
find : FIND LPAREN searchExpression COMMA textExpression (COMMA startExpression)? RPAREN ;
floor : FLOOR LPAREN expression RPAREN ;
geolocation : GEOLOCATION LPAREN latitudeExpression COMMA longitudeExpression RPAREN ;
// TOOD: Only allow literals, $SObjectType.<object>?
getrecordids : GETRECORDIDS LPAREN expression RPAREN ;
getsessionid : GETSESSIONID LPAREN RPAREN ;
hour : HOUR LPAREN expression RPAREN ;
htmlencode : HTMLENCODE LPAREN expression RPAREN ;
hyperlink : HYPERLINK LPAREN urlExpression COMMA nameExpression (COMMA targetExpression)? RPAREN ;
if : IF LPAREN logicalExpression COMMA ifTrueExpression COMMA ifFalseExpression RPAREN ;
image : IMAGE LPAREN urlExpression COMMA textExpression (COMMA heightExpression COMMA widthExpression)? RPAREN ;


valueExpression : expression ;
resultExpression : expression ;
defaultExpression : expression ;

yearExpression : expression ;
monthExpression : expression ;
dayExpression : expression ;

unitExpression : expression ;

searchExpression : expression ;
textExpression : expression ;
startExpression : expression ;

latitudeExpression : expression ;
longitudeExpression : expression ;

urlExpression : expression ;
nameExpression : expression ;
targetExpression : expression ;

logicalExpression : expression ;
ifTrueExpression : expression ;
ifFalseExpression : expression ;

heightExpression : expression ;
widthExpression : expression ;

primary
    :   literal
    |   LPAREN expression RPAREN
    ;

arguments
    :   expression (COMMA expression)*
    ;

literal
    :   StringLiteral
    |   IntegerLiteral
    |   FloatingPointLiteral
    |   BooleanLiteral
    |   NullLiteral
    ;
