grammar ecarule;

/* Grammar for expressions. */
import expr;

/* Rule. */
prule : RULE id (IN id)? ON evt (DEFAULT actslist)? task ;

/* Event. */
evt : id SEMICOLON (id SEMICOLON)* ;

/* Task. */
task : FOR (SOME | ALL)? exp DO actslist ;

/* List of actions. */
actslist : act SEMICOLON (act SEMICOLON)* ;

/* Actions. */
act : (THIS)? id ASSIGN exp    // internal action
    | EXT id ASSIGN exp     // external action
    ;

/* Keywords and operators. */
DEFAULT     : 'default' ;
IN          : 'in' ;
ON          : 'on' ;
DO          : 'do' ;
ALL         : 'all' ;
SOME        : 'some' ;
FOR         : 'for' ;
RULE        : 'rule' ;
SEMICOLON   : ';' ;

/* Match (but discard) whitespaces. */
WS: [ \t\r\n]+ -> skip;