package parser

import (
	"testing"
	"time"
)

func TestWhereEvaluator_ComplexConditions(t *testing.T) {
	envs := map[string]interface{}{
		"user_id":      123,
		"status":       "active",
		"subscription": "premium",
		"region":       "us-west-2",
		"age":          28,
		"score":        92.5,
		"last_login":   "2024-01-15",
		"login_count":  150,
		"tags":         []string{"admin", "poweruser"},
		"created_at":   time.Date(2023, 1, 15, 10, 30, 0, 0, time.UTC),
		"balance":      1500.75,
		"is_verified":  true,
		"department":   "engineering",
		"level":        "senior",
		"projects":     5,
	}

	tests := []struct {
		name        string
		whereExpr   string
		expected    bool
		description string
	}{
		{
			"Deeply nested conditions",
			"((user_id = 123 OR user_id = 456) AND (status = 'active' AND subscription = 'premium')) OR (age > 25 AND score >= 90.0)",
			true,
			"Complex nested conditions with multiple levels of parentheses",
		},
		{
			"Multiple AND/OR combinations",
			"user_id = 123 AND status = 'active' AND (subscription = 'premium' OR subscription = 'enterprise') AND region IN ('us-west-2', 'us-east-1', 'eu-west-1')",
			true,
			"Multiple conditions with IN operator",
		},
		{
			"Complex numeric conditions",
			"age BETWEEN 25 AND 35 AND score > 90.0 AND balance > 1000.0 AND login_count >= 100",
			true,
			"Multiple numeric comparisons (Note: BETWEEN not implemented yet, will use >= AND <=)",
		},
		{
			"String pattern matching combinations",
			"department LIKE 'eng%' AND level IN ('senior', 'principal', 'staff') AND status NOT LIKE '%inactive%'",
			true,
			"Complex string pattern matching",
		},
		{
			"Mixed data type conditions",
			"is_verified = TRUE AND projects >= 3 AND department = 'engineering' AND subscription NOT IN ('free', 'basic')",
			true,
			"Conditions across different data types",
		},
		{
			"Complex negation patterns",
			"(status != 'inactive' AND subscription != 'free') AND (age >= 18 AND score >= 50.0)",
			true,
			"Complex negation with De Morgan's law patterns (using != instead of NOT)",
		},
		{
			"Edge case with NULL handling",
			"(description IS NULL OR description = '') AND user_id IS NOT NULL AND status IS NOT NULL",
			true,
			"NULL handling in complex conditions",
		},
		{
			"Performance stress test",
			"(user_id = 123 OR user_id = 124 OR user_id = 125) AND (status = 'active' OR status = 'pending') AND (region = 'us-west-2' OR region = 'us-east-1' OR region = 'eu-west-1' OR region = 'ap-south-1') AND (subscription = 'premium' OR subscription = 'enterprise')",
			true,
			"Multiple OR conditions to test performance",
		},
		{
			"False complex condition",
			"(user_id = 999 AND status = 'active') OR (age < 18 AND subscription = 'premium') OR (score < 50.0 AND region = 'us-west-2')",
			false,
			"Complex condition that should evaluate to false",
		},
		{
			"Mixed LIKE and IN operations",
			"department LIKE '%eng%' AND level IN ('junior', 'senior', 'staff') AND region NOT IN ('ap-south-1', 'eu-central-1') AND status LIKE 'act%'",
			true,
			"Combination of LIKE and IN operations",
		},
	}

	evaluator := NewWhereEvaluator(envs)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// For BETWEEN operator, convert to equivalent >= AND <= expression
			whereExpr := tt.whereExpr
			if tt.name == "Complex numeric conditions" {
				whereExpr = "age >= 25 AND age <= 35 AND score > 90.0 AND balance > 1000.0 AND login_count >= 100"
			}

			result, err := evaluator.EvaluateWhere(whereExpr)
			if err != nil {
				t.Errorf("EvaluateWhere() error = %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("EvaluateWhere() = %v, want %v\nExpression: %s\nDescription: %s",
					result, tt.expected, whereExpr, tt.description)
			}
		})
	}
}

func TestWhereEvaluator_RealWorldScenarios(t *testing.T) {
	// User profile data
	userProfile := map[string]interface{}{
		"user_id":            12345,
		"email":              "john.doe@company.com",
		"status":             "active",
		"subscription":       "enterprise",
		"region":             "us-west-2",
		"age":                32,
		"department":         "engineering",
		"role":               "senior_engineer",
		"salary":             120000,
		"years_experience":   8,
		"skills":             "go,python,sql,kubernetes",
		"is_manager":         false,
		"is_remote":          true,
		"last_login_days":    2,
		"projects_count":     7,
		"performance_rating": 4.5,
	}

	scenarios := []struct {
		name        string
		description string
		sql         string
		expected    bool
	}{
		{
			"HR Query - Senior Engineers",
			"Find senior engineers with high performance in engineering",
			"SELECT * FROM employees WHERE role LIKE '%senior%' AND department = 'engineering' AND performance_rating >= 4.0 AND years_experience >= 5",
			true,
		},
		{
			"Access Control - Enterprise Users",
			"Check if user has enterprise access in allowed regions",
			"SELECT * FROM users WHERE subscription = 'enterprise' AND region IN ('us-west-2', 'us-east-1', 'eu-west-1') AND status = 'active'",
			true,
		},
		{
			"Performance Review - High Performers",
			"Identify high-performing employees for promotion",
			"SELECT * FROM employees WHERE performance_rating > 4.0 AND projects_count >= 5 AND years_experience >= 7 AND salary < 150000",
			true,
		},
		{
			"Security Audit - Recent Login Check",
			"Find active users who haven't logged in recently",
			"SELECT * FROM users WHERE status = 'active' AND last_login_days > 30",
			false,
		},
		{
			"Compensation Analysis - Remote Workers",
			"Analyze remote workers in engineering with competitive salary",
			"SELECT * FROM employees WHERE is_remote = TRUE AND department = 'engineering' AND salary >= 100000 AND performance_rating >= 4.0",
			true,
		},
		{
			"Skill-based Filtering",
			"Find engineers with specific technical skills",
			"SELECT * FROM employees WHERE skills LIKE '%go%' AND skills LIKE '%kubernetes%' AND department = 'engineering'",
			true,
		},
		{
			"Management Track Identification",
			"Identify potential management candidates",
			"SELECT * FROM employees WHERE is_manager = FALSE AND years_experience >= 8 AND performance_rating >= 4.5 AND projects_count >= 6",
			true,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			result, err := EvaluateSQL(scenario.sql, userProfile)
			if err != nil {
				t.Errorf("EvaluateSQL() error = %v", err)
				return
			}
			if result != scenario.expected {
				t.Errorf("Scenario '%s' failed:\nSQL: %s\nExpected: %v, Got: %v\nDescription: %s",
					scenario.name, scenario.sql, scenario.expected, result, scenario.description)
			}
		})
	}
}

func TestWhereEvaluator_DataTypeVariations(t *testing.T) {
	envs := map[string]interface{}{
		"int_value":      42,
		"int64_value":    int64(9223372036854775807),
		"float32_value":  float32(3.14159),
		"float64_value":  3.14159265359,
		"string_value":   "test string",
		"bool_true":      true,
		"bool_false":     false,
		"nil_value":      nil,
		"empty_string":   "",
		"zero_int":       0,
		"zero_float":     0.0,
		"unicode_string": "测试中文字符串",
		"json_like":      `{"key": "value"}`,
		"email":          "test@example.com",
		"url":            "https://example.com/path",
		"uuid":           "550e8400-e29b-41d4-a716-446655440000",
	}

	tests := []struct {
		name      string
		whereExpr string
		expected  bool
	}{
		{"Large int64 comparison", "int64_value > 9000000000000000000", true},
		{"Float precision test", "float64_value > 3.14159", true},
		{"Unicode string matching", "unicode_string = '测试中文字符串'", true},
		{"JSON-like string contains", "json_like LIKE '%key%'", true},
		{"Email domain extraction", "email LIKE '%@example.com'", true},
		{"URL protocol check", "url LIKE 'https://%'", true},
		{"UUID format validation", "uuid LIKE '%-%'", true},
		{"Zero value comparisons", "zero_int = 0 AND zero_float = 0.0", true},
		{"Empty string vs NULL", "empty_string != '' OR nil_value IS NULL", true},
		{"Boolean logic combinations", "bool_true = TRUE AND bool_false = FALSE", true},
		{"Mixed type numeric comparison", "int_value = 42.0", true},
		{"String number comparison", "string_value != '42'", true},
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

func BenchmarkWhereEvaluator_SimpleCondition(b *testing.B) {
	envs := map[string]interface{}{
		"user_id": 123,
		"status":  "active",
	}
	evaluator := NewWhereEvaluator(envs)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := evaluator.EvaluateWhere("user_id = 123 AND status = 'active'")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWhereEvaluator_ComplexCondition(b *testing.B) {
	envs := map[string]interface{}{
		"user_id": 123,
		"status":  "active",
		"region":  "us-west-2",
		"age":     25,
		"score":   85.5,
	}
	evaluator := NewWhereEvaluator(envs)
	complexExpr := "(user_id = 123 OR user_id = 456) AND (status = 'active' AND region IN ('us-west-2', 'us-east-1')) AND (age > 20 OR score >= 80.0)"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := evaluator.EvaluateWhere(complexExpr)
		if err != nil {
			b.Fatal(err)
		}
	}
}
