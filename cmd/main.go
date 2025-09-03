package main

import (
	"fmt"

	"github.com/antlr4-sqlparser/parser"
)

func main() {
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

	fmt.Println("ANTLR4 SQL WHERE Clause Evaluator")
	fmt.Println("==================================")
	fmt.Printf("Environment variables: %+v\n\n", envs)

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
	fmt.Println("WHERE Expression Examples")
	fmt.Println("=========================")

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
