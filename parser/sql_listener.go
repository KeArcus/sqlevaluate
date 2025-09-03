// Code generated from grammar/SQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SQL

import "github.com/antlr4-go/antlr/v4"

// SQLListener is a complete listener for a parse tree produced by SQLParser.
type SQLListener interface {
	antlr.ParseTreeListener

	// EnterSql_statement is called when entering the sql_statement production.
	EnterSql_statement(c *Sql_statementContext)

	// EnterSelect_statement is called when entering the select_statement production.
	EnterSelect_statement(c *Select_statementContext)

	// EnterSelect_list is called when entering the select_list production.
	EnterSelect_list(c *Select_listContext)

	// EnterColumn_list is called when entering the column_list production.
	EnterColumn_list(c *Column_listContext)

	// EnterTable_name is called when entering the table_name production.
	EnterTable_name(c *Table_nameContext)

	// EnterColumn_name is called when entering the column_name production.
	EnterColumn_name(c *Column_nameContext)

	// EnterWhere_clause is called when entering the where_clause production.
	EnterWhere_clause(c *Where_clauseContext)

	// EnterInCondition is called when entering the inCondition production.
	EnterInCondition(c *InConditionContext)

	// EnterOrCondition is called when entering the orCondition production.
	EnterOrCondition(c *OrConditionContext)

	// EnterParenthesisCondition is called when entering the parenthesisCondition production.
	EnterParenthesisCondition(c *ParenthesisConditionContext)

	// EnterAndCondition is called when entering the andCondition production.
	EnterAndCondition(c *AndConditionContext)

	// EnterIsNullCondition is called when entering the isNullCondition production.
	EnterIsNullCondition(c *IsNullConditionContext)

	// EnterLikeCondition is called when entering the likeCondition production.
	EnterLikeCondition(c *LikeConditionContext)

	// EnterIsNotNullCondition is called when entering the isNotNullCondition production.
	EnterIsNotNullCondition(c *IsNotNullConditionContext)

	// EnterComparisonCondition is called when entering the comparisonCondition production.
	EnterComparisonCondition(c *ComparisonConditionContext)

	// EnterNotInCondition is called when entering the notInCondition production.
	EnterNotInCondition(c *NotInConditionContext)

	// EnterNotLikeCondition is called when entering the notLikeCondition production.
	EnterNotLikeCondition(c *NotLikeConditionContext)

	// EnterComparison_expression is called when entering the comparison_expression production.
	EnterComparison_expression(c *Comparison_expressionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterComparison_operator is called when entering the comparison_operator production.
	EnterComparison_operator(c *Comparison_operatorContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterValue_list is called when entering the value_list production.
	EnterValue_list(c *Value_listContext)

	// EnterString_literal is called when entering the string_literal production.
	EnterString_literal(c *String_literalContext)

	// EnterNumber_literal is called when entering the number_literal production.
	EnterNumber_literal(c *Number_literalContext)

	// EnterBoolean_literal is called when entering the boolean_literal production.
	EnterBoolean_literal(c *Boolean_literalContext)

	// EnterFunction_call is called when entering the function_call production.
	EnterFunction_call(c *Function_callContext)

	// EnterFunction_name is called when entering the function_name production.
	EnterFunction_name(c *Function_nameContext)

	// EnterFunction_args is called when entering the function_args production.
	EnterFunction_args(c *Function_argsContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitSql_statement is called when exiting the sql_statement production.
	ExitSql_statement(c *Sql_statementContext)

	// ExitSelect_statement is called when exiting the select_statement production.
	ExitSelect_statement(c *Select_statementContext)

	// ExitSelect_list is called when exiting the select_list production.
	ExitSelect_list(c *Select_listContext)

	// ExitColumn_list is called when exiting the column_list production.
	ExitColumn_list(c *Column_listContext)

	// ExitTable_name is called when exiting the table_name production.
	ExitTable_name(c *Table_nameContext)

	// ExitColumn_name is called when exiting the column_name production.
	ExitColumn_name(c *Column_nameContext)

	// ExitWhere_clause is called when exiting the where_clause production.
	ExitWhere_clause(c *Where_clauseContext)

	// ExitInCondition is called when exiting the inCondition production.
	ExitInCondition(c *InConditionContext)

	// ExitOrCondition is called when exiting the orCondition production.
	ExitOrCondition(c *OrConditionContext)

	// ExitParenthesisCondition is called when exiting the parenthesisCondition production.
	ExitParenthesisCondition(c *ParenthesisConditionContext)

	// ExitAndCondition is called when exiting the andCondition production.
	ExitAndCondition(c *AndConditionContext)

	// ExitIsNullCondition is called when exiting the isNullCondition production.
	ExitIsNullCondition(c *IsNullConditionContext)

	// ExitLikeCondition is called when exiting the likeCondition production.
	ExitLikeCondition(c *LikeConditionContext)

	// ExitIsNotNullCondition is called when exiting the isNotNullCondition production.
	ExitIsNotNullCondition(c *IsNotNullConditionContext)

	// ExitComparisonCondition is called when exiting the comparisonCondition production.
	ExitComparisonCondition(c *ComparisonConditionContext)

	// ExitNotInCondition is called when exiting the notInCondition production.
	ExitNotInCondition(c *NotInConditionContext)

	// ExitNotLikeCondition is called when exiting the notLikeCondition production.
	ExitNotLikeCondition(c *NotLikeConditionContext)

	// ExitComparison_expression is called when exiting the comparison_expression production.
	ExitComparison_expression(c *Comparison_expressionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitComparison_operator is called when exiting the comparison_operator production.
	ExitComparison_operator(c *Comparison_operatorContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitValue_list is called when exiting the value_list production.
	ExitValue_list(c *Value_listContext)

	// ExitString_literal is called when exiting the string_literal production.
	ExitString_literal(c *String_literalContext)

	// ExitNumber_literal is called when exiting the number_literal production.
	ExitNumber_literal(c *Number_literalContext)

	// ExitBoolean_literal is called when exiting the boolean_literal production.
	ExitBoolean_literal(c *Boolean_literalContext)

	// ExitFunction_call is called when exiting the function_call production.
	ExitFunction_call(c *Function_callContext)

	// ExitFunction_name is called when exiting the function_name production.
	ExitFunction_name(c *Function_nameContext)

	// ExitFunction_args is called when exiting the function_args production.
	ExitFunction_args(c *Function_argsContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
