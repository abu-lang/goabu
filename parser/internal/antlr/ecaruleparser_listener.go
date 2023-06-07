// Code generated from EcaruleParser.g4 by ANTLR 4.10.1 and MODIFIED by ../Makefile.

package antlr // EcaruleParser
import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/hyperjumptech/grule-rule-engine/antlr/parser/grulev3"
)

// EcaruleParserListener is a complete listener for a parse tree produced by EcaruleParser.
type EcaruleParserListener interface {
	antlr.ParseTreeListener

	// EnterPrules is called when entering the prules production.
	EnterPrules(c *PrulesContext)

	// EnterPrule is called when entering the prule production.
	EnterPrule(c *PruleContext)

	// EnterEvents is called when entering the events production.
	EnterEvents(c *EventsContext)

	// EnterDefaultActions is called when entering the defaultActions production.
	EnterDefaultActions(c *DefaultActionsContext)

	// EnterTask is called when entering the task production.
	EnterTask(c *TaskContext)

	// EnterActions is called when entering the actions production.
	EnterActions(c *ActionsContext)

	// EnterTailActions is called when entering the tailActions production.
	EnterTailActions(c *TailActionsContext)

	// EnterMaybeActions is called when entering the maybeActions production.
	EnterMaybeActions(c *MaybeActionsContext)

	// EnterGrl is called when entering the grl production.
	EnterGrl(c *grulev3.GrlContext)

	// EnterRuleEntry is called when entering the ruleEntry production.
	EnterRuleEntry(c *grulev3.RuleEntryContext)

	// EnterSalience is called when entering the salience production.
	EnterSalience(c *grulev3.SalienceContext)

	// EnterRuleName is called when entering the ruleName production.
	EnterRuleName(c *grulev3.RuleNameContext)

	// EnterRuleDescription is called when entering the ruleDescription production.
	EnterRuleDescription(c *grulev3.RuleDescriptionContext)

	// EnterWhenScope is called when entering the whenScope production.
	EnterWhenScope(c *grulev3.WhenScopeContext)

	// EnterThenScope is called when entering the thenScope production.
	EnterThenScope(c *grulev3.ThenScopeContext)

	// EnterThenExpressionList is called when entering the thenExpressionList production.
	EnterThenExpressionList(c *grulev3.ThenExpressionListContext)

	// EnterThenExpression is called when entering the thenExpression production.
	EnterThenExpression(c *grulev3.ThenExpressionContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *grulev3.AssignmentContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *grulev3.ExpressionContext)

	// EnterMulDivOperators is called when entering the mulDivOperators production.
	EnterMulDivOperators(c *grulev3.MulDivOperatorsContext)

	// EnterAddMinusOperators is called when entering the addMinusOperators production.
	EnterAddMinusOperators(c *grulev3.AddMinusOperatorsContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *grulev3.ComparisonOperatorContext)

	// EnterAndLogicOperator is called when entering the andLogicOperator production.
	EnterAndLogicOperator(c *grulev3.AndLogicOperatorContext)

	// EnterOrLogicOperator is called when entering the orLogicOperator production.
	EnterOrLogicOperator(c *grulev3.OrLogicOperatorContext)

	// EnterExpressionAtom is called when entering the expressionAtom production.
	EnterExpressionAtom(c *grulev3.ExpressionAtomContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *grulev3.ConstantContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *grulev3.VariableContext)

	// EnterArrayMapSelector is called when entering the arrayMapSelector production.
	EnterArrayMapSelector(c *grulev3.ArrayMapSelectorContext)

	// EnterMemberVariable is called when entering the memberVariable production.
	EnterMemberVariable(c *grulev3.MemberVariableContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *grulev3.FunctionCallContext)

	// EnterMethodCall is called when entering the methodCall production.
	EnterMethodCall(c *grulev3.MethodCallContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *grulev3.ArgumentListContext)

	// EnterFloatLiteral is called when entering the floatLiteral production.
	EnterFloatLiteral(c *grulev3.FloatLiteralContext)

	// EnterDecimalFloatLiteral is called when entering the decimalFloatLiteral production.
	EnterDecimalFloatLiteral(c *grulev3.DecimalFloatLiteralContext)

	// EnterHexadecimalFloatLiteral is called when entering the hexadecimalFloatLiteral production.
	EnterHexadecimalFloatLiteral(c *grulev3.HexadecimalFloatLiteralContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *grulev3.IntegerLiteralContext)

	// EnterDecimalLiteral is called when entering the decimalLiteral production.
	EnterDecimalLiteral(c *grulev3.DecimalLiteralContext)

	// EnterHexadecimalLiteral is called when entering the hexadecimalLiteral production.
	EnterHexadecimalLiteral(c *grulev3.HexadecimalLiteralContext)

	// EnterOctalLiteral is called when entering the octalLiteral production.
	EnterOctalLiteral(c *grulev3.OctalLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *grulev3.StringLiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *grulev3.BooleanLiteralContext)

	// ExitPrules is called when exiting the prules production.
	ExitPrules(c *PrulesContext)

	// ExitPrule is called when exiting the prule production.
	ExitPrule(c *PruleContext)

	// ExitEvents is called when exiting the events production.
	ExitEvents(c *EventsContext)

	// ExitDefaultActions is called when exiting the defaultActions production.
	ExitDefaultActions(c *DefaultActionsContext)

	// ExitTask is called when exiting the task production.
	ExitTask(c *TaskContext)

	// ExitActions is called when exiting the actions production.
	ExitActions(c *ActionsContext)

	// ExitTailActions is called when exiting the tailActions production.
	ExitTailActions(c *TailActionsContext)

	// ExitMaybeActions is called when exiting the maybeActions production.
	ExitMaybeActions(c *MaybeActionsContext)

	// ExitGrl is called when exiting the grl production.
	ExitGrl(c *grulev3.GrlContext)

	// ExitRuleEntry is called when exiting the ruleEntry production.
	ExitRuleEntry(c *grulev3.RuleEntryContext)

	// ExitSalience is called when exiting the salience production.
	ExitSalience(c *grulev3.SalienceContext)

	// ExitRuleName is called when exiting the ruleName production.
	ExitRuleName(c *grulev3.RuleNameContext)

	// ExitRuleDescription is called when exiting the ruleDescription production.
	ExitRuleDescription(c *grulev3.RuleDescriptionContext)

	// ExitWhenScope is called when exiting the whenScope production.
	ExitWhenScope(c *grulev3.WhenScopeContext)

	// ExitThenScope is called when exiting the thenScope production.
	ExitThenScope(c *grulev3.ThenScopeContext)

	// ExitThenExpressionList is called when exiting the thenExpressionList production.
	ExitThenExpressionList(c *grulev3.ThenExpressionListContext)

	// ExitThenExpression is called when exiting the thenExpression production.
	ExitThenExpression(c *grulev3.ThenExpressionContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *grulev3.AssignmentContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *grulev3.ExpressionContext)

	// ExitMulDivOperators is called when exiting the mulDivOperators production.
	ExitMulDivOperators(c *grulev3.MulDivOperatorsContext)

	// ExitAddMinusOperators is called when exiting the addMinusOperators production.
	ExitAddMinusOperators(c *grulev3.AddMinusOperatorsContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *grulev3.ComparisonOperatorContext)

	// ExitAndLogicOperator is called when exiting the andLogicOperator production.
	ExitAndLogicOperator(c *grulev3.AndLogicOperatorContext)

	// ExitOrLogicOperator is called when exiting the orLogicOperator production.
	ExitOrLogicOperator(c *grulev3.OrLogicOperatorContext)

	// ExitExpressionAtom is called when exiting the expressionAtom production.
	ExitExpressionAtom(c *grulev3.ExpressionAtomContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *grulev3.ConstantContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *grulev3.VariableContext)

	// ExitArrayMapSelector is called when exiting the arrayMapSelector production.
	ExitArrayMapSelector(c *grulev3.ArrayMapSelectorContext)

	// ExitMemberVariable is called when exiting the memberVariable production.
	ExitMemberVariable(c *grulev3.MemberVariableContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *grulev3.FunctionCallContext)

	// ExitMethodCall is called when exiting the methodCall production.
	ExitMethodCall(c *grulev3.MethodCallContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *grulev3.ArgumentListContext)

	// ExitFloatLiteral is called when exiting the floatLiteral production.
	ExitFloatLiteral(c *grulev3.FloatLiteralContext)

	// ExitDecimalFloatLiteral is called when exiting the decimalFloatLiteral production.
	ExitDecimalFloatLiteral(c *grulev3.DecimalFloatLiteralContext)

	// ExitHexadecimalFloatLiteral is called when exiting the hexadecimalFloatLiteral production.
	ExitHexadecimalFloatLiteral(c *grulev3.HexadecimalFloatLiteralContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *grulev3.IntegerLiteralContext)

	// ExitDecimalLiteral is called when exiting the decimalLiteral production.
	ExitDecimalLiteral(c *grulev3.DecimalLiteralContext)

	// ExitHexadecimalLiteral is called when exiting the hexadecimalLiteral production.
	ExitHexadecimalLiteral(c *grulev3.HexadecimalLiteralContext)

	// ExitOctalLiteral is called when exiting the octalLiteral production.
	ExitOctalLiteral(c *grulev3.OctalLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *grulev3.StringLiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *grulev3.BooleanLiteralContext)
}
