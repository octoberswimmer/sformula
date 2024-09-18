parser grammar FormulaParser;
options {tokenVocab=FormulaLexer;}

compilationUnit
    :   expression
    ;

expression
    :   primary                                                   # primaryExpression
    |   Identifier LPAREN arguments? RPAREN                       # functionCall
    |   Identifier (DOT Identifier)*                              # fieldReference
    |   expression CARET expression                               # binaryExpression
    |   expression BITAND expression                              # binaryExpression
    |   expression ADD expression                                 # binaryExpression
    |   expression SUB expression                                 # binaryExpression
    |   expression MUL expression                                 # binaryExpression
    |   expression DIV expression                                 # binaryExpression
    |   expression (GT | LT) ASSIGN? expression                   # binaryExpression
    |   expression (EQUAL | NOTEQUAL | LESSANDGREATER) expression # binaryExpression
    |   expression (AND | OR) expression                          # binaryExpression
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
