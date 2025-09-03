package main

import (
	"fmt"
	"log"

	"github.com/antlr4-sqlparser/parser"
)

func main() {
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
			AND account_status = 'active'
			AND region LIKE 'us-%'`

	fmt.Println("Basic Usage Example")
	fmt.Println("===================")
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
	fmt.Println("\n" + "Different Environment Example")
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
