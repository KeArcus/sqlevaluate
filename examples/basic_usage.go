package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/KeArcus/sqlevaluate/parser"
)

func main() {
	// Basic Usage Example - Simple WHERE clause evaluation
	basicUsageExample()

	fmt.Println("\n" + strings.Repeat("=", 60) + "\n")

	// Comprehensive Examples - Multiple WHERE conditions
	comprehensiveExamples()

	fmt.Println("\n" + strings.Repeat("=", 60) + "\n")

	// WHERE Only Examples - Test without SQL statement prefix
	whereOnlyExamples()
}

func basicUsageExample() {
	fmt.Println("Basic Usage Example")
	fmt.Println("===================")

	// Define environment variables that represent the current context
	envs := map[string]interface{}{
		"user_id":        42,
		"user_role":      "admin",
		"subscription":   "premium",
		"region":         "us-west-2",
		"account_status": "active",
		"login_count":    150,
		"last_login":     "2024-01-15",
	}

	// SQL query with WHERE clause to evaluate
	sql := `SELECT * FROM users 
			WHERE user_id = 42 
			AND user_role = 'admin' 
			AND account_status = 'active'`

	fmt.Printf("SQL: %s\n", sql)
	fmt.Printf("Environment: %+v\n\n", envs)

	// Evaluate the WHERE clause against environment variables
	matches, err := parser.EvaluateSQL(sql, envs)
	if err != nil {
		log.Fatalf("Error evaluating SQL: %v", err)
	}

	fmt.Printf("WHERE clause matches environment: %v\n", matches)

	if matches {
		fmt.Println("✅ Query conditions are satisfied by current environment")
	} else {
		fmt.Println("❌ Query conditions are NOT satisfied by current environment")
	}

	// Example with different environment that doesn't match
	fmt.Println("\nDifferent Environment Example")
	fmt.Println("==============================")

	differentEnvs := map[string]interface{}{
		"user_id":        99,     // Different user
		"user_role":      "user", // Different role
		"subscription":   "free",
		"region":         "eu-west-1",
		"account_status": "active",
		"login_count":    5,
	}

	fmt.Printf("Environment: %+v\n\n", differentEnvs)

	matches2, err := parser.EvaluateSQL(sql, differentEnvs)
	if err != nil {
		log.Fatalf("Error evaluating SQL: %v", err)
	}

	fmt.Printf("WHERE clause matches environment: %v\n", matches2)

	if matches2 {
		fmt.Println("✅ Query conditions are satisfied by current environment")
	} else {
		fmt.Println("❌ Query conditions are NOT satisfied by current environment")
	}
}

func comprehensiveExamples() {
	fmt.Println("Comprehensive WHERE Clause Examples")
	fmt.Println("===================================")

	// Example environment variables
	envs := map[string]interface{}{
		"user_id":    123,
		"status":     "active",
		"region":     "us-west-2",
		"age":        25,
		"name":       "John Doe",
		"is_premium": true,
		"score":      85.5,
	}

	fmt.Printf("Environment variables: %+v\n\n", envs)

	// Example SQL queries with WHERE clauses
	examples := []string{
		"SELECT * FROM users WHERE user_id = 123",
		"SELECT * FROM users WHERE status = 'active'",
		"SELECT * FROM users WHERE age > 20",
		"SELECT * FROM users WHERE user_id = 123 AND status = 'active'",
		"SELECT * FROM users WHERE age >= 25 OR is_premium = TRUE",
		"SELECT * FROM users WHERE name LIKE 'John%'",
		"SELECT * FROM users WHERE region IN ('us-west-2', 'us-east-1')",
		"SELECT * FROM users WHERE user_id != 456",
		"SELECT * FROM users WHERE score > 80.0",
		"SELECT * FROM users WHERE (age > 20 AND status = 'active') OR is_premium = TRUE",
	}

	fmt.Println("Complete SQL Statement Examples:")
	fmt.Println("---------------------------------")
	for i, sql := range examples {
		fmt.Printf("Example %d: %s\n", i+1, sql)

		result, err := parser.EvaluateSQL(sql, envs)
		if err != nil {
			fmt.Printf("  Error: %v\n", err)
		} else {
			fmt.Printf("  Result: %v\n", result)
		}
		fmt.Println()
	}

	// Example of evaluating just WHERE expressions
	fmt.Println("WHERE Expression Examples:")
	fmt.Println("--------------------------")

	whereExamples := []string{
		"user_id = 123",
		"status = 'active' AND age > 20",
		"name LIKE 'John%'",
		"region IN ('us-west-2', 'eu-west-1')",
		"score >= 85.0",
	}

	for _, whereExpr := range whereExamples {
		fmt.Printf("WHERE %s\n", whereExpr)

		result, err := parser.EvaluateWhere(whereExpr, envs)
		if err != nil {
			fmt.Printf("  Error: %v\n", err)
		} else {
			fmt.Printf("  Result: %v\n", result)
		}
		fmt.Println()
	}
}

func whereOnlyExamples() {
	fmt.Println("WHERE Only Examples (Without SQL Statement Prefix)")
	fmt.Println("==================================================")

	// Test environment variables
	envs := map[string]interface{}{
		"user_id":      123,
		"status":       "active",
		"region":       "us-west-2",
		"age":          25,
		"name":         "John Doe",
		"email":        "john@example.com",
		"is_premium":   true,
		"score":        85.5,
		"department":   "Engineering",
		"salary":       75000,
		"created_date": "2022-03-15",
		"bonus":        nil,
		"description":  "  Senior Developer  ",
	}

	fmt.Printf("Environment: %+v\n\n", envs)

	// Test cases for WHERE-only expressions
	testCases := []struct {
		description string
		expression  string
		shouldMatch bool
	}{
		// Basic comparisons
		{"Simple equality", "user_id = 123", true},
		{"String comparison", "status = 'active'", true},
		{"Numeric comparison", "age > 20", true},
		{"Boolean comparison", "is_premium = TRUE", true},
		{"Null comparison", "bonus IS NULL", true},

		// Failed matches
		{"Wrong user ID", "user_id = 456", false},
		{"Wrong status", "status = 'inactive'", false},
		{"Age too high", "age > 30", false},

		// Complex logical operations
		{"AND operation (true)", "user_id = 123 AND status = 'active'", true},
		{"AND operation (false)", "user_id = 123 AND status = 'inactive'", false},
		{"OR operation (true)", "user_id = 456 OR status = 'active'", true},
		{"OR operation (false)", "user_id = 456 OR status = 'inactive'", false},

		// Parentheses grouping
		{"Grouped conditions", "(user_id = 123 OR user_id = 456) AND status = 'active'", true},
		{"Complex grouping", "(age > 20 AND status = 'active') OR (is_premium = TRUE AND region = 'eu-west-1')", true},

		// Pattern matching
		{"LIKE prefix", "name LIKE 'John%'", true},
		{"LIKE suffix", "email LIKE '%example.com'", true},
		{"LIKE contains", "department LIKE '%Engineer%'", true},
		{"NOT LIKE", "status NOT LIKE '%inactive%'", true},

		// IN operations
		{"IN operation (match)", "region IN ('us-west-2', 'us-east-1')", true},
		{"IN operation (no match)", "region IN ('eu-west-1', 'ap-south-1')", false},
		{"NOT IN operation", "status NOT IN ('inactive', 'suspended')", true},

		// Function examples
		{"UPPER function", "UPPER(name) = 'JOHN DOE'", true},
		{"LENGTH function", "LENGTH(name) = 8", true},
		{"TRIM function", "LENGTH(TRIM(description)) > 10", true},
		{"SUBSTR function", "SUBSTR(name, 1, 4) = 'John'", true},
		{"YEAR function", "YEAR(created_date) = 2022", true},
		{"ROUND function", "ROUND(salary) = 75000", true},
		{"COALESCE function", "COALESCE(bonus, 0) = 0", true},

		// Complex function combinations
		{"Nested functions", "LENGTH(UPPER(TRIM(name))) = 8", true},
		{"Function with AND", "YEAR(created_date) = 2022 AND LENGTH(department) > 5", true},
		{"Function in LIKE", "LOWER(email) LIKE '%example%'", true},
	}

	fmt.Println("Testing WHERE-only expressions:")
	fmt.Println("--------------------------------")

	successCount := 0
	totalCount := len(testCases)

	for i, tc := range testCases {
		fmt.Printf("%2d. %s\n", i+1, tc.description)
		fmt.Printf("    Expression: %s\n", tc.expression)

		result, err := parser.EvaluateWhere(tc.expression, envs)
		if err != nil {
			fmt.Printf("    ❌ ERROR: %v\n", err)
		} else {
			if result == tc.shouldMatch {
				fmt.Printf("    ✅ PASS: %v (expected %v)\n", result, tc.shouldMatch)
				successCount++
			} else {
				fmt.Printf("    ❌ FAIL: %v (expected %v)\n", result, tc.shouldMatch)
			}
		}
		fmt.Println()
	}

	fmt.Printf("Summary: %d/%d tests passed (%.1f%%)\n",
		successCount, totalCount, float64(successCount)/float64(totalCount)*100)

	if successCount == totalCount {
		fmt.Println("🎉 All WHERE-only expressions work correctly!")
	} else {
		fmt.Printf("⚠️  %d test(s) failed\n", totalCount-successCount)
	}

	// Additional edge cases
	fmt.Println("\nEdge Cases:")
	fmt.Println("-----------")

	edgeCases := []struct {
		description string
		expression  string
	}{
		{"Empty string handling", "name != ''"},
		{"Zero values", "salary > 0"},
		{"Null checks", "bonus IS NULL"},
		{"Case sensitivity", "UPPER(status) = 'ACTIVE'"},
		{"Unicode strings", "name LIKE '%ö%' OR name LIKE '%ü%'"},
		{"Large numbers", "salary < 1000000"},
		{"Float precision", "score >= 85.5"},
		{"Complex nesting", "((user_id = 123 AND status = 'active') OR (age > 30)) AND (is_premium = TRUE OR salary > 50000)"},
	}

	for i, ec := range edgeCases {
		fmt.Printf("%d. %s\n", i+1, ec.description)
		fmt.Printf("   Expression: %s\n", ec.expression)

		result, err := parser.EvaluateWhere(ec.expression, envs)
		if err != nil {
			fmt.Printf("   Result: ERROR - %v\n", err)
		} else {
			fmt.Printf("   Result: %v\n", result)
		}
		fmt.Println()
	}
}
