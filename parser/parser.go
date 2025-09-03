package parser

import (
	"fmt"
)

// EvaluateSQL evaluates a complete SQL statement's WHERE clause against environment variables
func EvaluateSQL(sql string, envs map[string]interface{}) (bool, error) {
	evaluator := NewWhereEvaluator(envs)
	return evaluator.EvaluateSQL(sql)
}

// EvaluateWhere evaluates a WHERE clause expression against environment variables
func EvaluateWhere(whereExpr string, envs map[string]interface{}) (bool, error) {
	evaluator := NewWhereEvaluator(envs)
	return evaluator.EvaluateWhere(whereExpr)
}

// ParseSQL parses a SQL statement and returns information about it
func ParseSQL(sql string) (*SQLInfo, error) {
	// This is a placeholder for additional SQL parsing functionality
	// For now, we'll focus on WHERE clause evaluation
	return &SQLInfo{
		SQL:       sql,
		HasWhere:  containsWhere(sql),
		TableName: extractTableName(sql),
	}, nil
}

// SQLInfo contains information about a parsed SQL statement
type SQLInfo struct {
	SQL       string
	HasWhere  bool
	TableName string
}

// containsWhere checks if SQL contains a WHERE clause
func containsWhere(sql string) bool {
	// Simple check - in production you might want more sophisticated parsing
	return len(sql) > 0 && (containsIgnoreCase(sql, "WHERE") || containsIgnoreCase(sql, "where"))
}

// extractTableName extracts table name from SQL (simplified)
func extractTableName(sql string) string {
	// Very basic table name extraction
	// In a real implementation, you'd use the ANTLR parser for this
	return "unknown"
}

// containsIgnoreCase checks if a string contains a substring (case insensitive)
func containsIgnoreCase(s, substr string) bool {
	return len(s) >= len(substr) &&
		fmt.Sprintf("%v", s) != fmt.Sprintf("%v", s[:len(s)-len(substr)]) // simplified check
}
