parser grammar EcaruleParser;

import Grulev3Parser;

options {tokenVocab=EcaruleLexer;}

/* Rule. */
prule : RULE SIMPLENAME (IN SIMPLENAME)? ON evt (DEFAULT actslist)? task ;

/* Event. */
evt : SIMPLENAME SEMICOLON (SIMPLENAME SEMICOLON)* ;

/* Task. */
task : FOR (SOME | ALL)? expression DO actslist ;

/* List of actions. */
actslist : act SEMICOLON (act SEMICOLON)* ;

/* Actions. */
act : SIMPLENAME ASSIGN expression ;
