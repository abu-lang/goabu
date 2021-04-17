grammar expr;

/* Expressions. */
exp : ROUNDLEFT exp ROUNDRIGHT                      // parenthesis
    | val                                           // value
    | NOT exp                                       // unary logical opeartion
    | exp op=(MUL | DIV) exp                        // arithmetic operations
    | exp op=(PLUS | MINUS) exp                     // arithmetic operations
    | exp op=(LT | LEQ | GT | GEQ) exp              // arithmetic comparison
    | exp op=(EQ | NEQ) exp                         // expressions equality
    | exp op=(AND | OR) exp                         // binary logical opeartions
    | CONCAT ROUNDLEFT exp COMMA exp ROUNDRIGHT     // string concatenation
    | (THIS)? id                                    // internal resource access
    | EXT id                                        // external resource access
    ;

/* Identifier. */
id : ID ;

/* Values. */
val : UNDEF   // undefined value
      | BOOL    // boolean value
      | INT     // integer value
      | DEC     // decimal value
      | STR     // string value
      ;

/* Keywords and operators. */
THIS        : 'this.' ;
EXT         : 'ext.' ;
UNDEF       : 'null' ;
AND         : 'and' ;
OR          : 'or' ;
NOT         : 'not' ;
EQ          : '==' ;
NEQ         : '=/=' ;
LT          : '<' ;
LEQ         : '<=' ;
GT          : '>' ;
GEQ         : '>=' ;
ASSIGN      : '=' ;
PLUS        : '+' ;
MINUS       : '-' ;
DIV         : '/' ;
MUL         : '*' ;
CONCAT      : 'concat' ;
ROUNDLEFT   : '(' ;
ROUNDRIGHT  : ')' ;
COMMA       : ',' ;

/* Boolean token definition. */
BOOL : 'true' | 'false' ;

/* Numeral tokens definition. */
INT : '0' | POS | '-' POS;
DEC : (INT | '-' '0') '.' DIGIT+;
fragment POS: POSDIGIT DIGIT*;
fragment DIGIT: '0' | POSDIGIT;
fragment POSDIGIT: [1-9] ;

/* String token definition. */
STR : '"' STRCHR* '"' ;
fragment STRCHR : ~["\\] | ESC ;
fragment ESC : '\\' [btnfr"'\\] ;

/* Identifier token definition. */
ID : [a-zA-Z]+[a-zA-Z0-9_]* ;

/* Match (but discard) whitespaces. */
WS: [ \t\r\n]+ -> skip;