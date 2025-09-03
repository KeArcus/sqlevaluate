package parser

import (
	"testing"
)

func TestWhereEvaluator_BasicComparisons(t *testing.T) {
	envs := map[string]interface{}{
		"user_id": 123,
		"status":  "active",
		"age":     25,
		"score":   85.5,
		"active":  true,
	}

	tests := []struct {
		name      string
		whereExpr string
		expected  bool
	}{
		{"Equal integer", "user_id = 123", true},
		{"Equal string", "status = 'active'", true},
		{"Not equal", "user_id != 456", true},
		{"Greater than", "age > 20", true},
		{"Less than", "age < 30", true},
		{"Greater equal", "age >= 25", true},
		{"Less equal", "age <= 25", true},
		{"Boolean true", "active = TRUE", true},
		{"Float comparison", "score > 80.0", true},
		{"Failed comparison", "user_id = 456", false},
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
				t.Errorf("EvaluateWhere() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestWhereEvaluator_LogicalOperators(t *testing.T) {
	envs := map[string]interface{}{
		"user_id": 123,
		"status":  "active",
		"age":     25,
	}

	tests := []struct {
		name      string
		whereExpr string
		expected  bool
	}{
		{"AND true", "user_id = 123 AND status = 'active'", true},
		{"AND false", "user_id = 123 AND status = 'inactive'", false},
		{"OR true", "user_id = 123 OR status = 'inactive'", true},
		{"OR false", "user_id = 456 OR status = 'inactive'", false},
		{"Complex expression", "(user_id = 123 AND status = 'active') OR age > 30", true},
		{"Parentheses", "(user_id = 456 OR status = 'active') AND age < 30", true},
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
				t.Errorf("EvaluateWhere() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestWhereEvaluator_IN_Operations(t *testing.T) {
	envs := map[string]interface{}{
		"region":  "us-west-2",
		"status":  "active",
		"user_id": 123,
	}

	tests := []struct {
		name      string
		whereExpr string
		expected  bool
	}{
		{"IN string match", "region IN ('us-west-2', 'us-east-1')", true},
		{"IN string no match", "region IN ('eu-west-1', 'ap-south-1')", false},
		{"IN number match", "user_id IN (123, 456, 789)", true},
		{"IN number no match", "user_id IN (456, 789)", false},
		{"NOT IN match", "region NOT IN ('eu-west-1', 'ap-south-1')", true},
		{"NOT IN no match", "region NOT IN ('us-west-2', 'us-east-1')", false},
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
				t.Errorf("EvaluateWhere() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestWhereEvaluator_LIKE_Operations(t *testing.T) {
	envs := map[string]interface{}{
		"name":  "John Doe",
		"email": "john@example.com",
	}

	tests := []struct {
		name      string
		whereExpr string
		expected  bool
	}{
		{"LIKE prefix", "name LIKE 'John%'", true},
		{"LIKE suffix", "email LIKE '%example.com'", true},
		{"LIKE contains", "name LIKE '%Doe%'", true},
		{"LIKE no match", "name LIKE 'Jane%'", false},
		{"NOT LIKE match", "name NOT LIKE 'Jane%'", true},
		{"NOT LIKE no match", "name NOT LIKE 'John%'", false},
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
				t.Errorf("EvaluateWhere() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestWhereEvaluator_NULL_Operations(t *testing.T) {
	envs := map[string]interface{}{
		"name":        "John",
		"description": nil,
	}

	tests := []struct {
		name      string
		whereExpr string
		expected  bool
	}{
		{"IS NULL true", "description IS NULL", true},
		{"IS NULL false", "name IS NULL", false},
		{"IS NOT NULL true", "name IS NOT NULL", true},
		{"IS NOT NULL false", "description IS NOT NULL", false},
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
				t.Errorf("EvaluateWhere() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestEvaluateSQL_CompleteSQL(t *testing.T) {
	envs := map[string]interface{}{
		"user_id": 123,
		"status":  "active",
	}

	tests := []struct {
		name     string
		sql      string
		expected bool
	}{
		{
			"Complete SQL with WHERE",
			"SELECT * FROM users WHERE user_id = 123 AND status = 'active'",
			true,
		},
		{
			"Complete SQL no match",
			"SELECT id, name FROM users WHERE user_id = 456",
			false,
		},
		{
			"SQL without WHERE",
			"SELECT * FROM users",
			true, // No WHERE clause means always match
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := EvaluateSQL(tt.sql, envs)
			if err != nil {
				t.Errorf("EvaluateSQL() error = %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("EvaluateSQL() = %v, want %v", result, tt.expected)
			}
		})
	}
}
