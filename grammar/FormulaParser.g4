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
    ;

abs
    : ABS LPAREN expression RPAREN
    ;

addMonths
    : ADDMONTHS LPAREN expression COMMA expression RPAREN
    ;

and
    : AND_FUNC LPAREN expression COMMA expression (COMMA expression)* RPAREN
    ;

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
