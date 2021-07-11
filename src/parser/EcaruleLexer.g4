// this lexer grammar was obtained by slightly MODIFYING the grulev3 grammar
// of the Grule Rule Engine which is released under the following license:

//  Copyright hyperjumptech/grule-rule-engine Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// source: "https://github.com/hyperjumptech/grule-rule-engine/blob/e63c3f6444865c7b76ed5c19e97dc2a4ed62810d/antlr/grulev3.g4"

lexer grammar EcaruleLexer;

options {tokenVocab=EcaruleLexer;}

fragment A                  : [aA] ;
fragment B                  : [bB] ;
fragment C                  : [cC] ;
fragment D                  : [dD] ;
fragment E                  : [eE] ;
fragment F                  : [fF] ;
fragment G                  : [gG] ;
fragment H                  : [hH] ;
fragment I                  : [iI] ;
fragment J                  : [jJ] ;
fragment K                  : [kK] ;
fragment L                  : [lL] ;
fragment M                  : [mM] ;
fragment N                  : [nN] ;
fragment O                  : [oO] ;
fragment P                  : [pP] ;
fragment Q                  : [qQ] ;
fragment R                  : [rR] ;
fragment S                  : [sS] ;
fragment T                  : [tT] ;
fragment U                  : [uU] ;
fragment V                  : [vV] ;
fragment W                  : [wW] ;
fragment X                  : [xX] ;
fragment Y                  : [yY] ;
fragment Z                  : [zZ] ;

fragment ISC                : 'A' .. 'Z'
                            | 'a' .. 'z'
                            | '\u00C0' .. '\u00D6'
                            | '\u00D8' .. '\u00F6'
                            | '\u00F8' .. '\u02FF'
                            | '\u0370' .. '\u037D'
                            | '\u037F' .. '\u1FFF'
                            | '\u200C' .. '\u200D'
                            | '\u2070' .. '\u218F'
                            | '\u2C00' .. '\u2FEF'
                            | '\u3001' .. '\uD7FF'
                            | '\uF900' .. '\uFDCF'
                            | '\uFDF0' .. '\uFFFD'
                            ;

fragment IC                 : ISC
                            | '0' .. '9'
                            | '_'
                            | '\u00B7'
                            | '\u0300' .. '\u036F'
                            | '\u203F' .. '\u2040'
                            ;


T__0                        : ',' ;
PLUS                        : '+' ;
MINUS                       : '-' ;
DIV                         : '/' ;
MUL                         : '*' ;
MOD                         : '%' ;
DOT                         : '.' ;
SEMICOLON                   : ';' ;

LR_BRACE                    : '{';
RR_BRACE                    : '}';
LR_BRACKET                  : '(';
RR_BRACKET                  : ')';
LS_BRACKET                  : '[';
RS_BRACKET                  : ']';

RULE                        : R U L E  ;
WHEN                        : W H E N ;
THEN                        : T H E N ;
AND                         : '&&' ;
OR                          : '||' ;
TRUE                        : T R U E ;
FALSE                       : F A L S E ;
NIL_LITERAL                 : N I L ;
NEGATION                    : '!' ;
SALIENCE                    : S A L I E N C E ;

EQUALS                      : '==' ;
ASSIGN                      : '=' ;
PLUS_ASIGN                  : '+=' ;
MINUS_ASIGN                 : '-=' ;
DIV_ASIGN                   : '/=' ;
MUL_ASIGN                   : '*=' ;
GT                          : '>' ;
LT                          : '<' ;
GTE                         : '>=' ;
LTE                         : '<=' ;
NOTEQUALS                   : '!=' ;

BITAND                      : '&';
BITOR                       : '|';

// START EcaruleParser UNSHARED TOKENS
DEFAULT     : 'default' ;
IN          : 'in' ;
ON          : 'on' ;
DO          : 'do' ;
ALL         : 'all' ;
SOME        : 'some' ;
FOR         : 'for' ;
// END   EcaruleParser UNSHARED TOKENS

SIMPLENAME                  : ISC IC*;

DQUOTA_STRING               : '"' ( '\\'. | '""' | ~('"'| '\\') )* '"';
SQUOTA_STRING               : '\'' ('\\'. | '\'\'' | ~('\'' | '\\'))* '\'';


DECIMAL_FLOAT_LIT           : DEC_LIT DOT DEC_DIGITS DECIMAL_EXPONENT?
                            | DEC_LIT DECIMAL_EXPONENT
                            | DOT DEC_DIGITS DECIMAL_EXPONENT?
                            ;

DECIMAL_EXPONENT            : E (PLUS|MINUS)? DEC_DIGITS;

HEX_FLOAT_LIT               : '0' X HEX_MANTISA HEX_EXPONENT
                            ;

fragment HEX_MANTISA        : HEX_DIGITS DOT HEX_DIGITS?
                            | HEX_DIGITS
                            | DOT HEX_DIGITS
                            ;

HEX_EXPONENT                : P (PLUS|MINUS)? DEC_DIGITS
                            ;

DEC_LIT                     : '0'
                            | [1-9] DEC_DIGITS?
                            ;

HEX_LIT                     : '0' X HEX_DIGITS;
OCT_LIT                     : '0' OCT_DIGITS;

fragment HEX_DIGITS         : HEX_DIGIT+;
fragment DEC_DIGITS         : DEC_DIGIT+;
fragment OCT_DIGITS         : OCT_DIGIT+;
fragment DEC_DIGIT          : [0-9];
fragment OCT_DIGIT          : [0-7];
fragment HEX_DIGIT          : [0-9a-fA-F];

// IGNORED TOKENS
SPACE                       : [ \t\r\n]+    -> skip;
COMMENT                     : '/*' .*? '*/' -> skip;
LINE_COMMENT                : '//' ~[\r\n]* -> skip;
