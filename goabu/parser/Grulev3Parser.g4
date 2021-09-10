// this parser grammar was obtained by slightly MODIFYING the grulev3 grammar
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

parser grammar Grulev3Parser;

grl
    : ruleEntry* EOF
    ;

ruleEntry
    : RULE ruleName ruleDescription? salience? LR_BRACE whenScope thenScope RR_BRACE
    ;

salience
    : SALIENCE integerLiteral
    ;

ruleName
    : SIMPLENAME
    ;

ruleDescription
    : DQUOTA_STRING | SQUOTA_STRING
    ;

whenScope
    : WHEN  expression
    ;

thenScope
    : THEN  thenExpressionList
    ;

thenExpressionList
    : (thenExpression SEMICOLON)+
    ;

thenExpression
    : assignment
    | expressionAtom
    ;

assignment
    : variable (ASSIGN | PLUS_ASIGN | MINUS_ASIGN | DIV_ASIGN | MUL_ASIGN) expression
    ;

expression
    : expression mulDivOperators expression
    | expression addMinusOperators expression
    | expression comparisonOperator expression
    | expression andLogicOperator expression
    | expression orLogicOperator expression
    | NEGATION? LR_BRACKET expression RR_BRACKET
    | expressionAtom
    ;

mulDivOperators
    : MUL | DIV | MOD
    ;

addMinusOperators
    : PLUS | MINUS | BITAND | BITOR
    ;

comparisonOperator
    : GT | LT | GTE | LTE | EQUALS | NOTEQUALS
    ;

andLogicOperator
    : AND
    ;

orLogicOperator
    : OR
    ;

expressionAtom
    : constant
    | variable
    | functionCall
    | expressionAtom methodCall
    | expressionAtom memberVariable
    | expressionAtom arrayMapSelector
    | NEGATION expressionAtom
    ;

constant
    : stringLiteral
    | integerLiteral
    | floatLiteral
    | booleanLiteral
    | NIL_LITERAL
    ;

variable
    : variable memberVariable
    | variable arrayMapSelector
    | SIMPLENAME
    ;

arrayMapSelector
    : LS_BRACKET expression RS_BRACKET
    ;

memberVariable
    : DOT SIMPLENAME
    ;

functionCall
    : SIMPLENAME LR_BRACKET argumentList? RR_BRACKET
    ;

methodCall
    : DOT functionCall
    ;

argumentList
    :  expression ( T__0 expression )*
    ;

floatLiteral
    : decimalFloatLiteral
    | hexadecimalFloatLiteral
    ;

decimalFloatLiteral
    : MINUS? DECIMAL_FLOAT_LIT
    ;

hexadecimalFloatLiteral
    : MINUS? HEX_FLOAT_LIT
    ;

integerLiteral
    : decimalLiteral
    | hexadecimalLiteral
    | octalLiteral
    ;

decimalLiteral
    : MINUS? DEC_LIT
    ;

hexadecimalLiteral
    : MINUS? HEX_LIT
    ;

octalLiteral
    : MINUS? OCT_LIT
    ;

stringLiteral
    : DQUOTA_STRING | SQUOTA_STRING
    ;

booleanLiteral
    : TRUE | FALSE
    ;
