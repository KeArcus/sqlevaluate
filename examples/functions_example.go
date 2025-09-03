package main

import (
	"fmt"

	"github.com/antlr4-sqlparser/parser"
)

func main() {
	// Example environment variables with various data types
	envs := map[string]interface{}{
		"user_name":    "john doe",
		"email":        "john.doe@company.com",
		"description":  "  Senior Software Engineer  ",
		"salary":       85750.50,
		"bonus":        nil,
		"created_date": "2022-03-15",
		"age":          28,
		"status":       "active",
		"department":   "Engineering",
		"rating":       4.7,
	}

	fmt.Println("SQL Functions Example")
	fmt.Println("=====================")
	fmt.Printf("Environment: %+v\n\n", envs)

	// String function examples
	stringExamples := []struct {
		description string
		sql         string
	}{
		{
			"Convert name to uppercase",
			"SELECT * FROM users WHERE UPPER(user_name) = 'JOHN DOE'",
		},
		{
			"Check email domain",
			"SELECT * FROM users WHERE LOWER(email) LIKE '%company.com'",
		},
		{
			"Get trimmed description length",
			"SELECT * FROM users WHERE LENGTH(TRIM(description)) > 20",
		},
		{
			"Extract first name",
			"SELECT * FROM users WHERE SUBSTR(user_name, 1, 4) = 'john'",
		},
		{
			"Build full email with concatenation",
			"SELECT * FROM users WHERE CONCAT(UPPER(user_name), '@COMPANY.COM') LIKE '%JOHN%'",
		},
	}

	fmt.Println("String Functions:")
	fmt.Println("-----------------")
	for _, example := range stringExamples {
		result, err := parser.EvaluateSQL(example.sql, envs)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("✓ %s: %v\n", example.description, result)
		fmt.Printf("  SQL: %s\n\n", example.sql)
	}

	// Math function examples
	mathExamples := []struct {
		description string
		sql         string
	}{
		{
			"Round salary to nearest thousand",
			"SELECT * FROM users WHERE ROUND(salary, -3) = 86000",
		},
		{
			"Check age is positive (using ABS)",
			"SELECT * FROM users WHERE ABS(age) = age",
		},
		{
			"Rating ceiling check",
			"SELECT * FROM users WHERE CEIL(rating) = 5",
		},
		{
			"Rating floor check",
			"SELECT * FROM users WHERE FLOOR(rating) >= 4",
		},
		{
			"Check if age is even (using MOD)",
			"SELECT * FROM users WHERE MOD(age, 2) = 0",
		},
	}

	fmt.Println("Math Functions:")
	fmt.Println("---------------")
	for _, example := range mathExamples {
		result, err := parser.EvaluateSQL(example.sql, envs)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("✓ %s: %v\n", example.description, result)
		fmt.Printf("  SQL: %s\n\n", example.sql)
	}

	// Date function examples
	dateExamples := []struct {
		description string
		sql         string
	}{
		{
			"Check creation year",
			"SELECT * FROM users WHERE YEAR(created_date) = 2022",
		},
		{
			"Check creation month",
			"SELECT * FROM users WHERE MONTH(created_date) >= 3",
		},
		{
			"Extract day from date",
			"SELECT * FROM users WHERE DAY(created_date) = 15",
		},
		{
			"Date extraction to string",
			"SELECT * FROM users WHERE DATE(created_date) = '2022-03-15'",
		},
	}

	fmt.Println("Date Functions:")
	fmt.Println("---------------")
	for _, example := range dateExamples {
		result, err := parser.EvaluateSQL(example.sql, envs)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("✓ %s: %v\n", example.description, result)
		fmt.Printf("  SQL: %s\n\n", example.sql)
	}

	// NULL handling function examples
	nullExamples := []struct {
		description string
		sql         string
	}{
		{
			"Use COALESCE for null bonus",
			"SELECT * FROM users WHERE COALESCE(bonus, 0) = 0",
		},
		{
			"Check if status is not equal to inactive (NULLIF)",
			"SELECT * FROM users WHERE NULLIF(status, 'inactive') = 'active'",
		},
		{
			"Use ISNULL to check bonus",
			"SELECT * FROM users WHERE ISNULL(bonus) = TRUE",
		},
		{
			"Use ISNULL with replacement value",
			"SELECT * FROM users WHERE ISNULL(bonus, 5000) = 5000",
		},
	}

	fmt.Println("NULL Handling Functions:")
	fmt.Println("------------------------")
	for _, example := range nullExamples {
		result, err := parser.EvaluateSQL(example.sql, envs)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("✓ %s: %v\n", example.description, result)
		fmt.Printf("  SQL: %s\n\n", example.sql)
	}

	// Complex nested function examples
	complexExamples := []struct {
		description string
		sql         string
	}{
		{
			"Nested string functions",
			"SELECT * FROM users WHERE LENGTH(UPPER(TRIM(user_name))) = 8",
		},
		{
			"Function in LIKE pattern",
			"SELECT * FROM users WHERE UPPER(email) LIKE CONCAT('%', UPPER('company'), '%')",
		},
		{
			"Multiple function conditions",
			"SELECT * FROM users WHERE YEAR(created_date) = 2022 AND LENGTH(department) > 5 AND ROUND(salary) > 85000",
		},
		{
			"Complex boolean expression with functions",
			"SELECT * FROM users WHERE CONTAINS(email, '@company.com') = TRUE AND CEIL(rating) >= 5",
		},
	}

	fmt.Println("Complex Function Combinations:")
	fmt.Println("------------------------------")
	for _, example := range complexExamples {
		result, err := parser.EvaluateSQL(example.sql, envs)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("✓ %s: %v\n", example.description, result)
		fmt.Printf("  SQL: %s\n\n", example.sql)
	}
}
