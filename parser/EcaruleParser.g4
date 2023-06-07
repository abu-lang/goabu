// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

parser grammar EcaruleParser;

import Grulev3Parser;

options {tokenVocab=EcaruleLexer;}

/* Rules. */
prules : prule+ ;

/* Rule. */
prule : RULE SIMPLENAME ON events defaultActions? task+ ;

/* Events. */
events : SIMPLENAME+ ;

/* Default actions. */
defaultActions : DEFAULT actions ;

/* Task. */
task : FOR ALL? expression DO actions ;

/* List of actions. */
actions : assignment tailActions ;
tailActions : T__0 maybeActions | /* epsilon */ ;
maybeActions : actions | /* epsilon */ ;
