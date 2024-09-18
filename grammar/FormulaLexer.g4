lexer grammar FormulaLexer;

channels {
    WHITESPACE_CHANNEL,
    COMMENT_CHANNEL
}

StringLiteral            : SQSTRING | DQSTRING;

fragment DQSTRING : '"' (DQESC | .)*? '"';
fragment DQESC    : '\\"' | '\\\\';
fragment SQSTRING : '\'' (SQESC | .)*? '\'';
fragment SQESC    : '\\\'' | '\\\\';

IntegerLiteral
    :   Digits
    ;

fragment
Digits
    :   Digit+
    ;

fragment
Digit
    :   '0'
    |   NonZeroDigit
    ;

fragment
NonZeroDigit
    :   [1-9]
    ;

FloatingPointLiteral
    :   Digits '.' Digits?
    ;

BooleanLiteral
    :   T R U E
    |   F A L S E
    ;

NullLiteral
    :   N U L L
    ;

Identifier
    :   Letter LetterOrDigit*
    ;

fragment
Letter
    :   [a-zA-Z$_] // these are the "java letters" below 0xFF
    ;

fragment
LetterOrDigit
    :   [a-zA-Z0-9$_] // these are the "java letters or digits" below 0xFF
    ;

WS  :  [ \t\r\n\u000C]+ -> channel(WHITESPACE_CHANNEL)
    ;

COMMENT
    :   '/*' .*? '*/' -> channel(COMMENT_CHANNEL)
    ;
LPAREN          : '(';
RPAREN          : ')';
DOT             : '.';
COMMA           : ',';
CARET           : '^';
BITAND          : '&';
ADD             : '+';
SUB             : '-';
MUL             : '*';
DIV             : '/';
LT              : '<';
GT              : '>';
ASSIGN          : '=';
LESSANDGREATER  : '<>';
NOTEQUAL        : '!=';
EQUAL           : '==';
AND             : '&&';
OR              : '||';

fragment A : [aA];
fragment B : [bB];
fragment C : [cC];
fragment D : [dD];
fragment E : [eE];
fragment F : [fF];
fragment G : [gG];
fragment H : [hH];
fragment I : [iI];
fragment J : [jJ];
fragment K : [kK];
fragment L : [lL];
fragment M : [mM];
fragment N : [nN];
fragment O : [oO];
fragment P : [pP];
fragment Q : [qQ];
fragment R : [rR];
fragment S : [sS];
fragment T : [tT];
fragment U : [uU];
fragment V : [vV];
fragment W : [wW];
fragment X : [xX];
fragment Y : [yY];
fragment Z : [zZ];
