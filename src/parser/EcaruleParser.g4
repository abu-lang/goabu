parser grammar EcaruleParser;

import Grulev3Parser;

options {tokenVocab=EcaruleLexer;}

/* Rule. */
prule : RULE SIMPLENAME (IN SIMPLENAME)? ON events (DEFAULT actions)? task ;

/* Events. */
events : SIMPLENAME+ ;

/* Task. */
task : FOR (SOME | ALL)? expression DO actions ;

/* List of actions. */
actions : assignment tailActions ;
tailActions : SEMICOLON maybeActions | /* epsilon */ ;
maybeActions : actions | /* epsilon */ ;
