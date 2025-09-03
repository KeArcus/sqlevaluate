package parser

import (
	"testing"
)

func TestWhereEvaluator_StringFunctions(t *testing.T) {
	envs := map[string]interface{}{
		"name":        "John Doe",
		"email":       "john.doe@company.com",
		"description": "  Senior Software Engineer  ",
		"status":      "active",
		"department":  "Engineering",
	}

	tests := []struct {
		name        string
		whereExpr   string
		expected    bool
		description string
	}{
		{
			"UPPER function",
			"UPPER(name) = 'JOHN DOE'",
			true,
			"Convert string to uppercase",
		},
		{
			"LOWER function",
			"LOWER(department) = 'engineering'",
			true,
			"Convert string to lowercase",
		},
		{
			"LENGTH function",
			"LENGTH(name) = 8",
			true,
			"Get string length",
		},
		{
			"SUBSTR function - with length",
			"SUBSTR(name, 1, 4) = 'John'",
			true,
			"Extract substring with start and length",
		},
		{
			"SUBSTR function - from position",
			"SUBSTR(email, 10) = 'company.com'",
			true,
			"Extract substring from position to end",
		},
		{
			"TRIM function",
			"TRIM(description) = 'Senior Software Engineer'",
			true,
			"Trim whitespace from string",
		},
		{
			"CONCAT function",
			"CONCAT(name, ' - ', department) = 'John Doe - Engineering'",
			true,
			"Concatenate multiple strings",
		},
		{
			"REPLACE function",
			"REPLACE(email, 'company.com', 'example.com') = 'john.doe@example.com'",
			true,
			"Replace substring in string",
		},
		{
			"CONTAINS function",
			"CONTAINS(email, '@company.com') = TRUE",
			true,
			"Check if string contains substring",
		},
		{
			"Complex string function combination",
			"LENGTH(TRIM(description)) > 20 AND UPPER(status) = 'ACTIVE'",
			true,
			"Combine multiple string functions",
		},
	}

	evaluator := NewWhereEvaluator(envs)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := evaluator.EvaluateWhere(tt.whereExpr)
			if err != nil {
				t.Errorf("EvaluateWhere() error = %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("EvaluateWhere() = %v, want %v\nExpression: %s\nDescription: %s",
					result, tt.expected, tt.whereExpr, tt.description)
			}
		})
	}
}

func TestWhereEvaluator_MathFunctions(t *testing.T) {
	envs := map[string]interface{}{
		"score":       85.7,
		"negative":    -42.5,
		"temperature": 23.456,
		"count":       157,
		"ratio":       0.875,
	}

	tests := []struct {
		name        string
		whereExpr   string
		expected    bool
		description string
	}{
		{
			"ABS function",
			"ABS(negative) = 42.5",
			true,
			"Get absolute value",
		},
		{
			"ROUND function - no precision",
			"ROUND(score) = 86",
			true,
			"Round to nearest integer",
		},
		{
			"ROUND function - with precision",
			"ROUND(temperature, 2) = 23.46",
			true,
			"Round to specified decimal places",
		},
		{
			"CEIL function",
			"CEIL(score) = 86",
			true,
			"Round up to nearest integer",
		},
		{
			"FLOOR function",
			"FLOOR(score) = 85",
			true,
			"Round down to nearest integer",
		},
		{
			"MOD function",
			"MOD(count, 10) = 7",
			true,
			"Get remainder of division",
		},
		{
			"Complex math combination",
			"ROUND(ABS(negative), 1) = 42.5",
			true,
			"Combine multiple math functions",
		},
		{
			"Math function in comparison",
			"CEIL(ratio) >= 1",
			true,
			"Use function result in comparison",
		},
	}

	evaluator := NewWhereEvaluator(envs)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := evaluator.EvaluateWhere(tt.whereExpr)
			if err != nil {
				t.Errorf("EvaluateWhere() error = %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("EvaluateWhere() = %v, want %v\nExpression: %s\nDescription: %s",
					result, tt.expected, tt.whereExpr, tt.description)
			}
		})
	}
}

func TestWhereEvaluator_DateTimeFunctions(t *testing.T) {
	envs := map[string]interface{}{
		"created_at": "2023-06-15",
		"updated_at": "2023-06-15 14:30:25",
		"birth_date": "1990-03-25",
		"login_time": "2024-01-15 09:15:30",
	}

	tests := []struct {
		name        string
		whereExpr   string
		expected    bool
		description string
	}{
		{
			"DATE function",
			"DATE(updated_at) = '2023-06-15'",
			true,
			"Extract date from datetime",
		},
		{
			"YEAR function",
			"YEAR(birth_date) = 1990",
			true,
			"Extract year from date",
		},
		{
			"MONTH function",
			"MONTH(created_at) = 6",
			true,
			"Extract month from date",
		},
		{
			"DAY function",
			"DAY(birth_date) = 25",
			true,
			"Extract day from date",
		},
		{
			"Complex date comparison",
			"YEAR(created_at) = 2023 AND MONTH(created_at) >= 6",
			true,
			"Combine multiple date functions",
		},
		{
			"Date range check",
			"YEAR(birth_date) >= 1990 AND YEAR(birth_date) <= 2000",
			true,
			"Check if date is in range",
		},
	}

	evaluator := NewWhereEvaluator(envs)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := evaluator.EvaluateWhere(tt.whereExpr)
			if err != nil {
				t.Errorf("EvaluateWhere() error = %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("EvaluateWhere() = %v, want %v\nExpression: %s\nDescription: %s",
					result, tt.expected, tt.whereExpr, tt.description)
			}
		})
	}
}

func TestWhereEvaluator_NullHandlingFunctions(t *testing.T) {
	envs := map[string]interface{}{
		"name":        "John",
		"description": nil,
		"email":       "john@example.com",
		"phone":       nil,
		"status":      "active",
		"score":       85,
	}

	tests := []struct {
		name        string
		whereExpr   string
		expected    bool
		description string
	}{
		{
			"COALESCE function - first non-null",
			"COALESCE(description, 'No description') = 'No description'",
			true,
			"Return first non-null value",
		},
		{
			"COALESCE function - existing value",
			"COALESCE(name, 'Unknown') = 'John'",
			true,
			"Return existing value when not null",
		},
		{
			"NULLIF function - equal values",
			"NULLIF(status, 'active') = NULL",
			true,
			"Return NULL when values are equal",
		},
		{
			"NULLIF function - different values",
			"NULLIF(status, 'inactive') = 'active'",
			true,
			"Return first value when values are different",
		},
		{
			"ISNULL function - check null",
			"ISNULL(description) = TRUE",
			true,
			"Check if value is null",
		},
		{
			"ISNULL function - with replacement",
			"ISNULL(phone, 'No phone') = 'No phone'",
			true,
			"Replace null with default value",
		},
	}

	evaluator := NewWhereEvaluator(envs)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := evaluator.EvaluateWhere(tt.whereExpr)
			if err != nil {
				t.Errorf("EvaluateWhere() error = %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("EvaluateWhere() = %v, want %v\nExpression: %s\nDescription: %s",
					result, tt.expected, tt.whereExpr, tt.description)
			}
		})
	}
}

func TestWhereEvaluator_ComplexFunctionCombinations(t *testing.T) {
	envs := map[string]interface{}{
		"first_name":    "john",
		"last_name":     "doe",
		"email":         "john.doe@company.com",
		"salary":        75000.50,
		"bonus":         nil,
		"created_date":  "2022-03-15",
		"description":   "  senior engineer  ",
		"department":    "Engineering",
		"project_count": 5,
	}

	tests := []struct {
		name        string
		whereExpr   string
		expected    bool
		description string
	}{
		{
			"String function chain",
			"UPPER(CONCAT(first_name, ' ', last_name)) = 'JOHN DOE'",
			true,
			"Chain string functions together",
		},
		{
			"Math and string combination",
			"LENGTH(department) = 11",
			true,
			"Use string function result in comparison",
		},
		{
			"Date and math functions",
			"YEAR(created_date) = 2022",
			true,
			"Use date extraction in comparison",
		},
		{
			"NULL handling with string functions",
			"COALESCE(bonus, salary) > 70000",
			true,
			"Use COALESCE with numeric comparison",
		},
		{
			"Complex nested functions",
			"ROUND(salary, 2) = 75000.5 AND LENGTH(TRIM(description)) > 10",
			true,
			"Multiple nested function calls",
		},
		{
			"Function in LIKE pattern",
			"LOWER(email) LIKE '%company%'",
			true,
			"Use functions in LIKE patterns",
		},
		{
			"Function result comparison",
			"YEAR(created_date) = 2022",
			true,
			"Use function result in comparison",
		},
		{
			"Boolean function result",
			"CONTAINS(email, '@company.com') = TRUE AND LENGTH(first_name) >= 3",
			true,
			"Use boolean function result in comparison",
		},
	}

	evaluator := NewWhereEvaluator(envs)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := evaluator.EvaluateWhere(tt.whereExpr)
			if err != nil {
				t.Errorf("EvaluateWhere() error = %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("EvaluateWhere() = %v, want %v\nExpression: %s\nDescription: %s",
					result, tt.expected, tt.whereExpr, tt.description)
			}
		})
	}
}

func TestWhereEvaluator_FunctionErrorHandling(t *testing.T) {
	envs := map[string]interface{}{
		"text":   "hello",
		"number": 42,
	}

	tests := []struct {
		name        string
		whereExpr   string
		expected    bool
		description string
	}{
		{
			"Unknown function defaults to empty string",
			"UNKNOWN_FUNC(text) = ''",
			true,
			"Unknown functions return empty string",
		},
		{
			"Function with wrong arg count",
			"SUBSTR(text) = ''",
			true,
			"Function with insufficient arguments returns default",
		},
		{
			"Division by zero in MOD",
			"MOD(number, 0) = 0",
			true,
			"MOD with zero divisor returns 0",
		},
		{
			"Invalid date format",
			"YEAR('invalid-date') = YEAR(NOW())",
			true,
			"Invalid date falls back to current date",
		},
	}

	evaluator := NewWhereEvaluator(envs)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := evaluator.EvaluateWhere(tt.whereExpr)
			if err != nil {
				t.Errorf("EvaluateWhere() error = %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("EvaluateWhere() = %v, want %v\nExpression: %s\nDescription: %s",
					result, tt.expected, tt.whereExpr, tt.description)
			}
		})
	}
}

func BenchmarkWhereEvaluator_SimpleFunctions(b *testing.B) {
	envs := map[string]interface{}{
		"name":  "John Doe",
		"score": 85.5,
	}
	evaluator := NewWhereEvaluator(envs)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := evaluator.EvaluateWhere("UPPER(name) = 'JOHN DOE' AND ROUND(score) = 86")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWhereEvaluator_ComplexFunctions(b *testing.B) {
	envs := map[string]interface{}{
		"first_name": "john",
		"last_name":  "doe",
		"salary":     75000.50,
		"date":       "2022-03-15",
	}
	evaluator := NewWhereEvaluator(envs)
	complexExpr := "LENGTH(CONCAT(UPPER(first_name), ' ', UPPER(last_name))) > 5 AND ROUND(salary / 12, 2) > 6000 AND YEAR(date) >= 2020"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := evaluator.EvaluateWhere(complexExpr)
		if err != nil {
			b.Fatal(err)
		}
	}
}
