parser grammar FormulaParser;
options {tokenVocab=FormulaLexer;}

compilationUnit
    :   expression
    ;

expression
    :   primary                                                   # primaryExpression
    |   functionCall                                              # functionCallExpression
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

valueExpression : expression ;
resultExpression : expression ;
defaultExpression : expression ;

yearExpression : expression ;
monthExpression : expression ;
dayExpression : expression ;

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
