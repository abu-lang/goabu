// Code generated from EcaruleParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package antlr // EcaruleParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEcaruleParserListener is a complete listener for a parse tree produced by EcaruleParser.
type BaseEcaruleParserListener struct{}

var _ EcaruleParserListener = &BaseEcaruleParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEcaruleParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEcaruleParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEcaruleParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEcaruleParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrule is called when production prule is entered.
func (s *BaseEcaruleParserListener) EnterPrule(ctx *PruleContext) {}

// ExitPrule is called when production prule is exited.
func (s *BaseEcaruleParserListener) ExitPrule(ctx *PruleContext) {}

// EnterEvt is called when production evt is entered.
func (s *BaseEcaruleParserListener) EnterEvt(ctx *EvtContext) {}

// ExitEvt is called when production evt is exited.
func (s *BaseEcaruleParserListener) ExitEvt(ctx *EvtContext) {}

// EnterTask is called when production task is entered.
func (s *BaseEcaruleParserListener) EnterTask(ctx *TaskContext) {}

// ExitTask is called when production task is exited.
func (s *BaseEcaruleParserListener) ExitTask(ctx *TaskContext) {}

// EnterActslist is called when production actslist is entered.
func (s *BaseEcaruleParserListener) EnterActslist(ctx *ActslistContext) {}

// ExitActslist is called when production actslist is exited.
func (s *BaseEcaruleParserListener) ExitActslist(ctx *ActslistContext) {}

// EnterAct is called when production act is entered.
func (s *BaseEcaruleParserListener) EnterAct(ctx *ActContext) {}

// ExitAct is called when production act is exited.
func (s *BaseEcaruleParserListener) ExitAct(ctx *ActContext) {}

// EnterGrl is called when production grl is entered.
func (s *BaseEcaruleParserListener) EnterGrl(ctx *GrlContext) {}

// ExitGrl is called when production grl is exited.
func (s *BaseEcaruleParserListener) ExitGrl(ctx *GrlContext) {}

// EnterRuleEntry is called when production ruleEntry is entered.
func (s *BaseEcaruleParserListener) EnterRuleEntry(ctx *RuleEntryContext) {}

// ExitRuleEntry is called when production ruleEntry is exited.
func (s *BaseEcaruleParserListener) ExitRuleEntry(ctx *RuleEntryContext) {}

// EnterSalience is called when production salience is entered.
func (s *BaseEcaruleParserListener) EnterSalience(ctx *SalienceContext) {}

// ExitSalience is called when production salience is exited.
func (s *BaseEcaruleParserListener) ExitSalience(ctx *SalienceContext) {}

// EnterRuleName is called when production ruleName is entered.
func (s *BaseEcaruleParserListener) EnterRuleName(ctx *RuleNameContext) {}

// ExitRuleName is called when production ruleName is exited.
func (s *BaseEcaruleParserListener) ExitRuleName(ctx *RuleNameContext) {}

// EnterRuleDescription is called when production ruleDescription is entered.
func (s *BaseEcaruleParserListener) EnterRuleDescription(ctx *RuleDescriptionContext) {}

// ExitRuleDescription is called when production ruleDescription is exited.
func (s *BaseEcaruleParserListener) ExitRuleDescription(ctx *RuleDescriptionContext) {}

// EnterWhenScope is called when production whenScope is entered.
func (s *BaseEcaruleParserListener) EnterWhenScope(ctx *WhenScopeContext) {}

// ExitWhenScope is called when production whenScope is exited.
func (s *BaseEcaruleParserListener) ExitWhenScope(ctx *WhenScopeContext) {}

// EnterThenScope is called when production thenScope is entered.
func (s *BaseEcaruleParserListener) EnterThenScope(ctx *ThenScopeContext) {}

// ExitThenScope is called when production thenScope is exited.
func (s *BaseEcaruleParserListener) ExitThenScope(ctx *ThenScopeContext) {}

// EnterThenExpressionList is called when production thenExpressionList is entered.
func (s *BaseEcaruleParserListener) EnterThenExpressionList(ctx *ThenExpressionListContext) {}

// ExitThenExpressionList is called when production thenExpressionList is exited.
func (s *BaseEcaruleParserListener) ExitThenExpressionList(ctx *ThenExpressionListContext) {}

// EnterThenExpression is called when production thenExpression is entered.
func (s *BaseEcaruleParserListener) EnterThenExpression(ctx *ThenExpressionContext) {}

// ExitThenExpression is called when production thenExpression is exited.
func (s *BaseEcaruleParserListener) ExitThenExpression(ctx *ThenExpressionContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseEcaruleParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseEcaruleParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEcaruleParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEcaruleParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMulDivOperators is called when production mulDivOperators is entered.
func (s *BaseEcaruleParserListener) EnterMulDivOperators(ctx *MulDivOperatorsContext) {}

// ExitMulDivOperators is called when production mulDivOperators is exited.
func (s *BaseEcaruleParserListener) ExitMulDivOperators(ctx *MulDivOperatorsContext) {}

// EnterAddMinusOperators is called when production addMinusOperators is entered.
func (s *BaseEcaruleParserListener) EnterAddMinusOperators(ctx *AddMinusOperatorsContext) {}

// ExitAddMinusOperators is called when production addMinusOperators is exited.
func (s *BaseEcaruleParserListener) ExitAddMinusOperators(ctx *AddMinusOperatorsContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseEcaruleParserListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseEcaruleParserListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterAndLogicOperator is called when production andLogicOperator is entered.
func (s *BaseEcaruleParserListener) EnterAndLogicOperator(ctx *AndLogicOperatorContext) {}

// ExitAndLogicOperator is called when production andLogicOperator is exited.
func (s *BaseEcaruleParserListener) ExitAndLogicOperator(ctx *AndLogicOperatorContext) {}

// EnterOrLogicOperator is called when production orLogicOperator is entered.
func (s *BaseEcaruleParserListener) EnterOrLogicOperator(ctx *OrLogicOperatorContext) {}

// ExitOrLogicOperator is called when production orLogicOperator is exited.
func (s *BaseEcaruleParserListener) ExitOrLogicOperator(ctx *OrLogicOperatorContext) {}

// EnterExpressionAtom is called when production expressionAtom is entered.
func (s *BaseEcaruleParserListener) EnterExpressionAtom(ctx *ExpressionAtomContext) {}

// ExitExpressionAtom is called when production expressionAtom is exited.
func (s *BaseEcaruleParserListener) ExitExpressionAtom(ctx *ExpressionAtomContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseEcaruleParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseEcaruleParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseEcaruleParserListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseEcaruleParserListener) ExitVariable(ctx *VariableContext) {}

// EnterArrayMapSelector is called when production arrayMapSelector is entered.
func (s *BaseEcaruleParserListener) EnterArrayMapSelector(ctx *ArrayMapSelectorContext) {}

// ExitArrayMapSelector is called when production arrayMapSelector is exited.
func (s *BaseEcaruleParserListener) ExitArrayMapSelector(ctx *ArrayMapSelectorContext) {}

// EnterMemberVariable is called when production memberVariable is entered.
func (s *BaseEcaruleParserListener) EnterMemberVariable(ctx *MemberVariableContext) {}

// ExitMemberVariable is called when production memberVariable is exited.
func (s *BaseEcaruleParserListener) ExitMemberVariable(ctx *MemberVariableContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseEcaruleParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseEcaruleParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterMethodCall is called when production methodCall is entered.
func (s *BaseEcaruleParserListener) EnterMethodCall(ctx *MethodCallContext) {}

// ExitMethodCall is called when production methodCall is exited.
func (s *BaseEcaruleParserListener) ExitMethodCall(ctx *MethodCallContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseEcaruleParserListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseEcaruleParserListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *BaseEcaruleParserListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *BaseEcaruleParserListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterDecimalFloatLiteral is called when production decimalFloatLiteral is entered.
func (s *BaseEcaruleParserListener) EnterDecimalFloatLiteral(ctx *DecimalFloatLiteralContext) {}

// ExitDecimalFloatLiteral is called when production decimalFloatLiteral is exited.
func (s *BaseEcaruleParserListener) ExitDecimalFloatLiteral(ctx *DecimalFloatLiteralContext) {}

// EnterHexadecimalFloatLiteral is called when production hexadecimalFloatLiteral is entered.
func (s *BaseEcaruleParserListener) EnterHexadecimalFloatLiteral(ctx *HexadecimalFloatLiteralContext) {
}

// ExitHexadecimalFloatLiteral is called when production hexadecimalFloatLiteral is exited.
func (s *BaseEcaruleParserListener) ExitHexadecimalFloatLiteral(ctx *HexadecimalFloatLiteralContext) {
}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BaseEcaruleParserListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BaseEcaruleParserListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *BaseEcaruleParserListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *BaseEcaruleParserListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterHexadecimalLiteral is called when production hexadecimalLiteral is entered.
func (s *BaseEcaruleParserListener) EnterHexadecimalLiteral(ctx *HexadecimalLiteralContext) {}

// ExitHexadecimalLiteral is called when production hexadecimalLiteral is exited.
func (s *BaseEcaruleParserListener) ExitHexadecimalLiteral(ctx *HexadecimalLiteralContext) {}

// EnterOctalLiteral is called when production octalLiteral is entered.
func (s *BaseEcaruleParserListener) EnterOctalLiteral(ctx *OctalLiteralContext) {}

// ExitOctalLiteral is called when production octalLiteral is exited.
func (s *BaseEcaruleParserListener) ExitOctalLiteral(ctx *OctalLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseEcaruleParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseEcaruleParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseEcaruleParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseEcaruleParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}
