// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

parser grammar EcaruleParser;

import Grulev3Parser;

options {tokenVocab=EcaruleLexer;}

/* Rule. */
prule : RULE SIMPLENAME ON events (DEFAULT actions)? task+ ;

/* Events. */
events : SIMPLENAME+ ;

/* Task. */
task : FOR ALL? expression DO actions ;

/* List of actions. */
actions : assignment tailActions ;
tailActions : SEMICOLON maybeActions | /* epsilon */ ;
maybeActions : actions | /* epsilon */ ;
