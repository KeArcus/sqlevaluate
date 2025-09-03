// Code generated from grammar/SQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SQL

import "github.com/antlr4-go/antlr/v4"

// BaseSQLListener is a complete listener for a parse tree produced by SQLParser.
type BaseSQLListener struct{}

var _ SQLListener = &BaseSQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSql_statement is called when production sql_statement is entered.
func (s *BaseSQLListener) EnterSql_statement(ctx *Sql_statementContext) {}

// ExitSql_statement is called when production sql_statement is exited.
func (s *BaseSQLListener) ExitSql_statement(ctx *Sql_statementContext) {}

// EnterSelect_statement is called when production select_statement is entered.
func (s *BaseSQLListener) EnterSelect_statement(ctx *Select_statementContext) {}

// ExitSelect_statement is called when production select_statement is exited.
func (s *BaseSQLListener) ExitSelect_statement(ctx *Select_statementContext) {}

// EnterSelect_list is called when production select_list is entered.
func (s *BaseSQLListener) EnterSelect_list(ctx *Select_listContext) {}

// ExitSelect_list is called when production select_list is exited.
func (s *BaseSQLListener) ExitSelect_list(ctx *Select_listContext) {}

// EnterColumn_list is called when production column_list is entered.
func (s *BaseSQLListener) EnterColumn_list(ctx *Column_listContext) {}

// ExitColumn_list is called when production column_list is exited.
func (s *BaseSQLListener) ExitColumn_list(ctx *Column_listContext) {}

// EnterTable_name is called when production table_name is entered.
func (s *BaseSQLListener) EnterTable_name(ctx *Table_nameContext) {}

// ExitTable_name is called when production table_name is exited.
func (s *BaseSQLListener) ExitTable_name(ctx *Table_nameContext) {}

// EnterColumn_name is called when production column_name is entered.
func (s *BaseSQLListener) EnterColumn_name(ctx *Column_nameContext) {}

// ExitColumn_name is called when production column_name is exited.
func (s *BaseSQLListener) ExitColumn_name(ctx *Column_nameContext) {}

// EnterWhere_clause is called when production where_clause is entered.
func (s *BaseSQLListener) EnterWhere_clause(ctx *Where_clauseContext) {}

// ExitWhere_clause is called when production where_clause is exited.
func (s *BaseSQLListener) ExitWhere_clause(ctx *Where_clauseContext) {}

// EnterInCondition is called when production inCondition is entered.
func (s *BaseSQLListener) EnterInCondition(ctx *InConditionContext) {}

// ExitInCondition is called when production inCondition is exited.
func (s *BaseSQLListener) ExitInCondition(ctx *InConditionContext) {}

// EnterOrCondition is called when production orCondition is entered.
func (s *BaseSQLListener) EnterOrCondition(ctx *OrConditionContext) {}

// ExitOrCondition is called when production orCondition is exited.
func (s *BaseSQLListener) ExitOrCondition(ctx *OrConditionContext) {}

// EnterParenthesisCondition is called when production parenthesisCondition is entered.
func (s *BaseSQLListener) EnterParenthesisCondition(ctx *ParenthesisConditionContext) {}

// ExitParenthesisCondition is called when production parenthesisCondition is exited.
func (s *BaseSQLListener) ExitParenthesisCondition(ctx *ParenthesisConditionContext) {}

// EnterAndCondition is called when production andCondition is entered.
func (s *BaseSQLListener) EnterAndCondition(ctx *AndConditionContext) {}

// ExitAndCondition is called when production andCondition is exited.
func (s *BaseSQLListener) ExitAndCondition(ctx *AndConditionContext) {}

// EnterIsNullCondition is called when production isNullCondition is entered.
func (s *BaseSQLListener) EnterIsNullCondition(ctx *IsNullConditionContext) {}

// ExitIsNullCondition is called when production isNullCondition is exited.
func (s *BaseSQLListener) ExitIsNullCondition(ctx *IsNullConditionContext) {}

// EnterLikeCondition is called when production likeCondition is entered.
func (s *BaseSQLListener) EnterLikeCondition(ctx *LikeConditionContext) {}

// ExitLikeCondition is called when production likeCondition is exited.
func (s *BaseSQLListener) ExitLikeCondition(ctx *LikeConditionContext) {}

// EnterIsNotNullCondition is called when production isNotNullCondition is entered.
func (s *BaseSQLListener) EnterIsNotNullCondition(ctx *IsNotNullConditionContext) {}

// ExitIsNotNullCondition is called when production isNotNullCondition is exited.
func (s *BaseSQLListener) ExitIsNotNullCondition(ctx *IsNotNullConditionContext) {}

// EnterComparisonCondition is called when production comparisonCondition is entered.
func (s *BaseSQLListener) EnterComparisonCondition(ctx *ComparisonConditionContext) {}

// ExitComparisonCondition is called when production comparisonCondition is exited.
func (s *BaseSQLListener) ExitComparisonCondition(ctx *ComparisonConditionContext) {}

// EnterNotInCondition is called when production notInCondition is entered.
func (s *BaseSQLListener) EnterNotInCondition(ctx *NotInConditionContext) {}

// ExitNotInCondition is called when production notInCondition is exited.
func (s *BaseSQLListener) ExitNotInCondition(ctx *NotInConditionContext) {}

// EnterNotLikeCondition is called when production notLikeCondition is entered.
func (s *BaseSQLListener) EnterNotLikeCondition(ctx *NotLikeConditionContext) {}

// ExitNotLikeCondition is called when production notLikeCondition is exited.
func (s *BaseSQLListener) ExitNotLikeCondition(ctx *NotLikeConditionContext) {}

// EnterComparison_expression is called when production comparison_expression is entered.
func (s *BaseSQLListener) EnterComparison_expression(ctx *Comparison_expressionContext) {}

// ExitComparison_expression is called when production comparison_expression is exited.
func (s *BaseSQLListener) ExitComparison_expression(ctx *Comparison_expressionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSQLListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSQLListener) ExitExpression(ctx *ExpressionContext) {}

// EnterComparison_operator is called when production comparison_operator is entered.
func (s *BaseSQLListener) EnterComparison_operator(ctx *Comparison_operatorContext) {}

// ExitComparison_operator is called when production comparison_operator is exited.
func (s *BaseSQLListener) ExitComparison_operator(ctx *Comparison_operatorContext) {}

// EnterValue is called when production value is entered.
func (s *BaseSQLListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseSQLListener) ExitValue(ctx *ValueContext) {}

// EnterValue_list is called when production value_list is entered.
func (s *BaseSQLListener) EnterValue_list(ctx *Value_listContext) {}

// ExitValue_list is called when production value_list is exited.
func (s *BaseSQLListener) ExitValue_list(ctx *Value_listContext) {}

// EnterString_literal is called when production string_literal is entered.
func (s *BaseSQLListener) EnterString_literal(ctx *String_literalContext) {}

// ExitString_literal is called when production string_literal is exited.
func (s *BaseSQLListener) ExitString_literal(ctx *String_literalContext) {}

// EnterNumber_literal is called when production number_literal is entered.
func (s *BaseSQLListener) EnterNumber_literal(ctx *Number_literalContext) {}

// ExitNumber_literal is called when production number_literal is exited.
func (s *BaseSQLListener) ExitNumber_literal(ctx *Number_literalContext) {}

// EnterBoolean_literal is called when production boolean_literal is entered.
func (s *BaseSQLListener) EnterBoolean_literal(ctx *Boolean_literalContext) {}

// ExitBoolean_literal is called when production boolean_literal is exited.
func (s *BaseSQLListener) ExitBoolean_literal(ctx *Boolean_literalContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BaseSQLListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BaseSQLListener) ExitFunction_call(ctx *Function_callContext) {}

// EnterFunction_name is called when production function_name is entered.
func (s *BaseSQLListener) EnterFunction_name(ctx *Function_nameContext) {}

// ExitFunction_name is called when production function_name is exited.
func (s *BaseSQLListener) ExitFunction_name(ctx *Function_nameContext) {}

// EnterFunction_args is called when production function_args is entered.
func (s *BaseSQLListener) EnterFunction_args(ctx *Function_argsContext) {}

// ExitFunction_args is called when production function_args is exited.
func (s *BaseSQLListener) ExitFunction_args(ctx *Function_argsContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseSQLListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseSQLListener) ExitIdentifier(ctx *IdentifierContext) {}
