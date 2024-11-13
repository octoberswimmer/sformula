lexer grammar FormulaLexer;

channels {
    WHITESPACE_CHANNEL,
    COMMENT_CHANNEL
}

// Functions
ABS            : A B S;
ADDMONTHS      : A D D M O N T H S;
AND_FUNC       : A N D;
BEGINS         : B E G I N S;
BLANKVALUE     : B L A N K V A L U E;
BR             : B R;
CASE           : C A S E;
CASESAFEID     : C A S E S A F E I D;
CEILING        : C E I L I N G;
CONTAINS       : C O N T A I N S;
CURRENCYRATE   : C U R R E N C Y R A T E;
DATE           : D A T E;
DATEVALUE      : D A T E V A L U E;
DATETIMEVALUE  : D A T E T I M E V A L U E;
DAY            : D A Y;
DISTANCE       : D I S T A N C E;
EXP            : E X P;
FIND           : F I N D;
FLOOR          : F L O O R;
GEOLOCATION    : G E O L O C A T I O N;
GETRECORDIDS   : G E T R E C O R D I D S;
GETSESSIONID   : G E T S E S S I O N I D;
HOUR           : H O U R;
HTMLENCODE     : H T M L E N C O D E;
HYPERLINK      : H Y P E R L I N K;
IF             : I F;
IMAGE          : I M A G E;
IMAGEPROXYURL  : I M A G E P R O X Y U R L;
INCLUDE        : I N C L U D E;
INCLUDES       : I N C L U D E S;
ISBLANK        : I S B L A N K;
ISCHANGED      : I S C H A N G E D;
ISCLONE        : I S C L O N E;
ISNEW          : I S N E W;
ISNULL         : I S N U L L;
ISNUMBER       : I S N U M B E R;
ISPICKVAL      : I S P I C K V A L;
JSENCODE       : J S E N C O D E;
JSINHTMLENCODE : J S I N H T M L E N C O D E;
JUNCTIONIDLIST : J U N C T I O N I D L I S T;
LEFT           : L E F T;
LEN            : L E N;
LINKTO         : L I N K T O;
LN             : L N;
LOG            : L O G;
LOWER          : L O W E R;
LPAD           : L P A D;
MAX            : M A X;
MCEILING       : M C E I L I N G;
MFLOOR         : M F L O O R;
MID            : M I D;
MILLISECOND    : M I L L I S E C O N D;
MIN            : M I N;
MINUTE         : M I N U T E;
MOD            : M O D;
MONTH          : M O N T H;
NOT            : N O T;
NOW            : N O W;
NULLVALUE      : N U L L V A L U E;
OR_FUNC        : O R;
PARENTGROUPVAL : P A R E N T G R O U P V A L;
PREDICT        : P R E D I C T;
PREVGROUPVAL   : P R E V G R O U P V A L;
PRIORVALUE     : P R I O R V A L U E;
REGEX          : R E G E X;
REQUIRESCRIPT  : R E Q U I R E S C R I P T;
REVERSE        : R E V E R S E;
RIGHT          : R I G H T;
ROUND          : R O U N D;
RPAD           : R P A D;
SECOND         : S E C O N D;
SQRT           : S Q R T;
SUBSTITUTE     : S U B S T I T U T E;
TEXT           : T E X T;
TIMENOW        : T I M E N O W;
TIMEVALUE      : T I M E V A L U E;
TODAY          : T O D A Y;
TRIM           : T R I M;
UPPER          : U P P E R;
URLENCODE      : U R L E N C O D E;
URLFOR         : U R L F O R;
VALUE          : V A L U E;
VLOOKUP        : V L O O K U P;
WEEKDAY        : W E E K D A Y;
YEAR           : Y E A R;

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
    :   [a-zA-Z0-9$_:] // these are the "java letters or digits" below 0xFF
    ;

LPAREN          : '(';
RPAREN          : ')';
DOT             : '.';
COMMA           : ',';

CARET           : '^';
BANG            : '!';
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

// Flows support variable expressions
BRACEBANG  : '{!';
RBRACE  : '}';

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

WS  :  [ \t\r\n\u000C]+ -> channel(WHITESPACE_CHANNEL)
    ;

COMMENT
    :   '/*' .*? '*/' -> channel(COMMENT_CHANNEL)
    ;
