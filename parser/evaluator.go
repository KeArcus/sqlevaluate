package parser

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/antlr4-go/antlr/v4"
)

// WhereEvaluator evaluates WHERE clause conditions against environment variables
type WhereEvaluator struct {
	envs map[string]interface{}
}

// NewWhereEvaluator creates a new WHERE clause evaluator
func NewWhereEvaluator(envs map[string]interface{}) *WhereEvaluator {
	return &WhereEvaluator{envs: envs}
}

// EvaluateSQL parses a complete SQL statement and evaluates the WHERE clause
func (e *WhereEvaluator) EvaluateSQL(sql string) (bool, error) {
	input := antlr.NewInputStream(sql)
	lexer := NewSQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewSQLParser(stream)

	// Parse the SQL statement
	tree := parser.Sql_statement()

	// Find the WHERE clause
	selectStmt := tree.Select_statement()
	if selectStmt == nil {
		return true, nil // No WHERE clause means always match
	}

	whereClause := selectStmt.Where_clause()
	if whereClause == nil {
		return true, nil // No WHERE clause means always match
	}

	// Evaluate the WHERE condition
	return e.evaluateCondition(whereClause.Condition()), nil
}

// EvaluateWhere directly evaluates a WHERE condition expression
func (e *WhereEvaluator) EvaluateWhere(whereExpr string) (bool, error) {
	// Wrap the WHERE expression in a complete SQL statement for parsing
	sql := fmt.Sprintf("SELECT * FROM dummy WHERE %s", whereExpr)
	return e.EvaluateSQL(sql)
}

// evaluateCondition evaluates a condition based on its type
func (e *WhereEvaluator) evaluateCondition(ctx IConditionContext) bool {
	switch c := ctx.(type) {
	case *AndConditionContext:
		return e.evaluateCondition(c.GetChild(0).(IConditionContext)) &&
			e.evaluateCondition(c.GetChild(2).(IConditionContext))

	case *OrConditionContext:
		return e.evaluateCondition(c.GetChild(0).(IConditionContext)) ||
			e.evaluateCondition(c.GetChild(2).(IConditionContext))

	case *ParenthesisConditionContext:
		return e.evaluateCondition(c.Condition())

	case *ComparisonConditionContext:
		return e.evaluateComparison(c.Comparison_expression())

	case *InConditionContext:
		return e.evaluateInExpr(c.GetChild(0).(IExpressionContext), c.Value_list(), false)

	case *NotInConditionContext:
		return e.evaluateInExpr(c.GetChild(0).(IExpressionContext), c.Value_list(), true)

	case *LikeConditionContext:
		// Find the expression contexts (skip LIKE terminal)
		leftExpr := c.GetChild(0).(IExpressionContext)
		rightExpr := c.GetChild(2).(IExpressionContext)
		return e.evaluateLikeExpr(leftExpr, rightExpr, false)

	case *NotLikeConditionContext:
		// Find the expression contexts (skip NOT LIKE terminals)
		leftExpr := c.GetChild(0).(IExpressionContext)
		rightExpr := c.GetChild(3).(IExpressionContext)
		return e.evaluateLikeExpr(leftExpr, rightExpr, true)

	case *IsNullConditionContext:
		return e.evaluateIsNullExpr(c.GetChild(0).(IExpressionContext), false)

	case *IsNotNullConditionContext:
		return e.evaluateIsNullExpr(c.GetChild(0).(IExpressionContext), true)

	default:
		return false
	}
}

// evaluateComparison evaluates comparison expressions (=, !=, <, >, etc.)
func (e *WhereEvaluator) evaluateComparison(ctx IComparison_expressionContext) bool {
	leftValue := e.evaluateExpression(ctx.GetChild(0).(IExpressionContext))
	operator := ctx.Comparison_operator().GetText()
	rightValue := e.evaluateExpression(ctx.GetChild(2).(IExpressionContext))

	return e.compareValues(leftValue, operator, rightValue)
}

// evaluateExpression evaluates expressions (functions, columns, values)
func (e *WhereEvaluator) evaluateExpression(ctx IExpressionContext) interface{} {
	switch {
	case ctx.Function_call() != nil:
		return e.evaluateFunction(ctx.Function_call())
	case ctx.Column_name() != nil:
		columnName := ctx.Column_name().GetText()
		if value, exists := e.envs[columnName]; exists {
			return value
		}
		return nil
	case ctx.Value() != nil:
		return e.parseValue(ctx.Value())
	default:
		return nil
	}
}

// evaluateIn evaluates IN and NOT IN conditions
func (e *WhereEvaluator) evaluateIn(columnName string, valueList IValue_listContext, isNot bool) bool {
	envValue, exists := e.envs[columnName]
	if !exists {
		return isNot // If column doesn't exist, NOT IN is true, IN is false
	}

	values := e.parseValueList(valueList)
	found := false
	for _, v := range values {
		if e.compareValues(envValue, "=", v) {
			found = true
			break
		}
	}

	return found != isNot
}

// evaluateInExpr evaluates IN and NOT IN conditions with expressions
func (e *WhereEvaluator) evaluateInExpr(exprCtx IExpressionContext, valueList IValue_listContext, isNot bool) bool {
	leftValue := e.evaluateExpression(exprCtx)
	values := e.parseValueList(valueList)

	found := false
	for _, v := range values {
		if e.compareValues(leftValue, "=", v) {
			found = true
			break
		}
	}

	return found != isNot
}

// evaluateLikeExpr evaluates LIKE and NOT LIKE conditions with expressions
func (e *WhereEvaluator) evaluateLikeExpr(leftExpr, rightExpr IExpressionContext, isNot bool) bool {
	leftValue := e.evaluateExpression(leftExpr)
	rightValue := e.evaluateExpression(rightExpr)

	leftStr := fmt.Sprintf("%v", leftValue)
	rightStr := fmt.Sprintf("%v", rightValue)

	// Convert SQL LIKE pattern to Go regex-like matching
	pattern := strings.ReplaceAll(rightStr, "%", ".*")
	pattern = strings.ReplaceAll(pattern, "_", ".")

	matched := e.simplePatternMatch(leftStr, pattern)
	return matched != isNot
}

// evaluateIsNullExpr evaluates IS NULL and IS NOT NULL conditions with expressions
func (e *WhereEvaluator) evaluateIsNullExpr(exprCtx IExpressionContext, isNot bool) bool {
	value := e.evaluateExpression(exprCtx)
	isNull := value == nil
	return isNull != isNot
}

// evaluateLike evaluates LIKE and NOT LIKE conditions
func (e *WhereEvaluator) evaluateLike(columnName, pattern string, isNot bool) bool {
	envValue, exists := e.envs[columnName]
	if !exists {
		return isNot
	}

	envStr := fmt.Sprintf("%v", envValue)
	// Remove quotes from pattern
	pattern = strings.Trim(pattern, "'\"")

	// Convert SQL LIKE pattern to Go regex-like matching
	// % matches any sequence of characters
	// _ matches any single character
	pattern = strings.ReplaceAll(pattern, "%", ".*")
	pattern = strings.ReplaceAll(pattern, "_", ".")

	// Simple pattern matching - not full regex for simplicity
	matched := e.simplePatternMatch(envStr, pattern)
	return matched != isNot
}

// evaluateIsNull evaluates IS NULL and IS NOT NULL conditions
func (e *WhereEvaluator) evaluateIsNull(columnName string, isNot bool) bool {
	envValue, exists := e.envs[columnName]
	if !exists {
		return !isNot // If column doesn't exist, IS NULL is true, IS NOT NULL is false
	}

	isNull := envValue == nil
	return isNull != isNot
}

// parseValue converts ANTLR value context to Go interface{}
func (e *WhereEvaluator) parseValue(ctx IValueContext) interface{} {
	switch {
	case ctx.String_literal() != nil:
		str := ctx.String_literal().GetText()
		return strings.Trim(str, "'\"") // Remove quotes

	case ctx.Number_literal() != nil:
		numStr := ctx.Number_literal().GetText()
		if strings.Contains(numStr, ".") {
			if f, err := strconv.ParseFloat(numStr, 64); err == nil {
				return f
			}
		} else {
			if i, err := strconv.ParseInt(numStr, 10, 64); err == nil {
				return i
			}
		}
		return numStr

	case ctx.Boolean_literal() != nil:
		return ctx.Boolean_literal().GetText() == "TRUE"

	case ctx.NULL() != nil:
		return nil

	default:
		return ctx.GetText()
	}
}

// parseValueList converts a list of expressions to Go slice
func (e *WhereEvaluator) parseValueList(ctx IValue_listContext) []interface{} {
	var values []interface{}
	for _, exprCtx := range ctx.AllExpression() {
		values = append(values, e.evaluateExpression(exprCtx))
	}
	return values
}

// compareValues compares two values using the given operator
func (e *WhereEvaluator) compareValues(left interface{}, operator string, right interface{}) bool {
	// Convert both values to comparable types
	leftVal := e.normalizeValue(left)
	rightVal := e.normalizeValue(right)

	switch operator {
	case "=":
		return e.valuesEqual(leftVal, rightVal)
	case "!=", "<>":
		return !e.valuesEqual(leftVal, rightVal)
	case "<":
		return e.compareNumeric(leftVal, rightVal) < 0
	case "<=":
		cmp := e.compareNumeric(leftVal, rightVal)
		return cmp < 0 || e.valuesEqual(leftVal, rightVal)
	case ">":
		return e.compareNumeric(leftVal, rightVal) > 0
	case ">=":
		cmp := e.compareNumeric(leftVal, rightVal)
		return cmp > 0 || e.valuesEqual(leftVal, rightVal)
	default:
		return false
	}
}

// normalizeValue converts values to comparable types
func (e *WhereEvaluator) normalizeValue(val interface{}) interface{} {
	if val == nil {
		return nil
	}

	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String:
		return val.(string)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(v.Uint())
	case reflect.Float32, reflect.Float64:
		return v.Float()
	case reflect.Bool:
		return v.Bool()
	default:
		return fmt.Sprintf("%v", val)
	}
}

// valuesEqual checks if two normalized values are equal
func (e *WhereEvaluator) valuesEqual(left, right interface{}) bool {
	if left == nil || right == nil {
		return left == right
	}

	// Try numeric comparison first
	if leftNum, leftOk := e.toFloat64(left); leftOk {
		if rightNum, rightOk := e.toFloat64(right); rightOk {
			return leftNum == rightNum
		}
	}

	// Fall back to string comparison
	return fmt.Sprintf("%v", left) == fmt.Sprintf("%v", right)
}

// compareNumeric compares two values numerically
func (e *WhereEvaluator) compareNumeric(left, right interface{}) int {
	leftNum, leftOk := e.toFloat64(left)
	rightNum, rightOk := e.toFloat64(right)

	if !leftOk || !rightOk {
		// Fall back to string comparison
		leftStr := fmt.Sprintf("%v", left)
		rightStr := fmt.Sprintf("%v", right)
		if leftStr < rightStr {
			return -1
		} else if leftStr > rightStr {
			return 1
		}
		return 0
	}

	if leftNum < rightNum {
		return -1
	} else if leftNum > rightNum {
		return 1
	}
	return 0
}

// toFloat64 converts a value to float64 if possible
func (e *WhereEvaluator) toFloat64(val interface{}) (float64, bool) {
	switch v := val.(type) {
	case int64:
		return float64(v), true
	case float64:
		return v, true
	case int:
		return float64(v), true
	case float32:
		return float64(v), true
	case string:
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f, true
		}
	}
	return 0, false
}

// simplePatternMatch performs simple pattern matching for LIKE operations
func (e *WhereEvaluator) simplePatternMatch(text, pattern string) bool {
	// Simple implementation - in a real scenario you might want to use regex
	if pattern == ".*" {
		return true
	}

	// Handle patterns without wildcards
	if !strings.Contains(pattern, ".*") && !strings.Contains(pattern, ".") {
		return text == pattern
	}

	// Handle contains patterns like ".*abc.*" (check this first)
	if strings.HasPrefix(pattern, ".*") && strings.HasSuffix(pattern, ".*") {
		middle := strings.TrimPrefix(strings.TrimSuffix(pattern, ".*"), ".*")
		if middle == "" {
			return true // ".*.*" matches everything
		}
		return strings.Contains(text, middle)
	}

	// Handle prefix patterns like "abc.*"
	if strings.HasSuffix(pattern, ".*") && !strings.HasPrefix(pattern, ".*") {
		prefix := strings.TrimSuffix(pattern, ".*")
		return strings.HasPrefix(text, prefix)
	}

	// Handle suffix patterns like ".*abc"
	if strings.HasPrefix(pattern, ".*") && !strings.HasSuffix(pattern, ".*") {
		suffix := strings.TrimPrefix(pattern, ".*")
		return strings.HasSuffix(text, suffix)
	}

	return false
}

// evaluateFunction evaluates SQL functions
func (e *WhereEvaluator) evaluateFunction(ctx IFunction_callContext) interface{} {
	funcName := strings.ToUpper(ctx.Function_name().GetText())

	var args []interface{}
	if ctx.Function_args() != nil {
		for _, argCtx := range ctx.Function_args().AllExpression() {
			args = append(args, e.evaluateExpression(argCtx))
		}
	}

	return e.executeFunction(funcName, args)
}

// executeFunction executes the actual function logic
func (e *WhereEvaluator) executeFunction(funcName string, args []interface{}) interface{} {
	switch funcName {
	// String functions
	case "UPPER":
		if len(args) == 1 {
			if str, ok := args[0].(string); ok {
				return strings.ToUpper(str)
			}
		}
		return ""

	case "LOWER":
		if len(args) == 1 {
			if str, ok := args[0].(string); ok {
				return strings.ToLower(str)
			}
		}
		return ""

	case "LENGTH":
		if len(args) == 1 {
			str := fmt.Sprintf("%v", args[0])
			return int64(len(str))
		}
		return int64(0)

	case "SUBSTR":
		if len(args) >= 2 {
			str := fmt.Sprintf("%v", args[0])
			if start, ok := e.toInt(args[1]); ok {
				startIdx := int(start) - 1 // SQL uses 1-based indexing
				if startIdx < 0 || startIdx >= len(str) {
					return ""
				}
				if len(args) >= 3 {
					if length, ok := e.toInt(args[2]); ok {
						endIdx := startIdx + int(length)
						if endIdx > len(str) {
							endIdx = len(str)
						}
						return str[startIdx:endIdx]
					}
				}
				return str[startIdx:]
			}
		}
		return ""

	case "TRIM":
		if len(args) == 1 {
			str := fmt.Sprintf("%v", args[0])
			return strings.TrimSpace(str)
		}
		return ""

	case "CONCAT":
		var result strings.Builder
		for _, arg := range args {
			result.WriteString(fmt.Sprintf("%v", arg))
		}
		return result.String()

	case "REPLACE":
		if len(args) == 3 {
			str := fmt.Sprintf("%v", args[0])
			old := fmt.Sprintf("%v", args[1])
			new := fmt.Sprintf("%v", args[2])
			return strings.ReplaceAll(str, old, new)
		}
		return ""

	case "CONTAINS":
		if len(args) == 2 {
			str := fmt.Sprintf("%v", args[0])
			substr := fmt.Sprintf("%v", args[1])
			return strings.Contains(str, substr)
		}
		return false

	// Math functions
	case "ABS":
		if len(args) == 1 {
			if num, ok := e.toFloat64(args[0]); ok {
				return math.Abs(num)
			}
		}
		return 0.0

	case "ROUND":
		if len(args) >= 1 {
			if num, ok := e.toFloat64(args[0]); ok {
				if len(args) >= 2 {
					if precision, ok := e.toInt(args[1]); ok {
						multiplier := math.Pow(10, float64(precision))
						return math.Round(num*multiplier) / multiplier
					}
				}
				return math.Round(num)
			}
		}
		return 0.0

	case "CEIL":
		if len(args) == 1 {
			if num, ok := e.toFloat64(args[0]); ok {
				return math.Ceil(num)
			}
		}
		return 0.0

	case "FLOOR":
		if len(args) == 1 {
			if num, ok := e.toFloat64(args[0]); ok {
				return math.Floor(num)
			}
		}
		return 0.0

	case "MOD":
		if len(args) == 2 {
			if num1, ok1 := e.toFloat64(args[0]); ok1 {
				if num2, ok2 := e.toFloat64(args[1]); ok2 && num2 != 0 {
					return math.Mod(num1, num2)
				}
			}
		}
		return 0.0

	// Date/Time functions
	case "NOW":
		return time.Now().Format("2006-01-02 15:04:05")

	case "DATE":
		if len(args) == 1 {
			str := fmt.Sprintf("%v", args[0])
			if t, err := time.Parse("2006-01-02 15:04:05", str); err == nil {
				return t.Format("2006-01-02")
			}
			if t, err := time.Parse("2006-01-02", str); err == nil {
				return t.Format("2006-01-02")
			}
		}
		return time.Now().Format("2006-01-02")

	case "YEAR":
		if len(args) == 1 {
			str := fmt.Sprintf("%v", args[0])
			if t, err := time.Parse("2006-01-02", str); err == nil {
				return int64(t.Year())
			}
			if t, err := time.Parse("2006-01-02 15:04:05", str); err == nil {
				return int64(t.Year())
			}
		}
		return int64(time.Now().Year())

	case "MONTH":
		if len(args) == 1 {
			str := fmt.Sprintf("%v", args[0])
			if t, err := time.Parse("2006-01-02", str); err == nil {
				return int64(t.Month())
			}
			if t, err := time.Parse("2006-01-02 15:04:05", str); err == nil {
				return int64(t.Month())
			}
		}
		return int64(time.Now().Month())

	case "DAY":
		if len(args) == 1 {
			str := fmt.Sprintf("%v", args[0])
			if t, err := time.Parse("2006-01-02", str); err == nil {
				return int64(t.Day())
			}
			if t, err := time.Parse("2006-01-02 15:04:05", str); err == nil {
				return int64(t.Day())
			}
		}
		return int64(time.Now().Day())

	// NULL handling functions
	case "COALESCE":
		for _, arg := range args {
			if arg != nil {
				return arg
			}
		}
		return nil

	case "NULLIF":
		if len(args) == 2 {
			if e.valuesEqual(args[0], args[1]) {
				return nil
			}
			return args[0]
		}
		return nil

	case "ISNULL":
		if len(args) >= 1 {
			if args[0] == nil {
				if len(args) >= 2 {
					return args[1] // Return replacement value
				}
				return true
			}
			return args[0]
		}
		return nil

	default:
		// Unknown function, return empty string
		return ""
	}
}

// toInt converts value to int64
func (e *WhereEvaluator) toInt(val interface{}) (int64, bool) {
	switch v := val.(type) {
	case int64:
		return v, true
	case int:
		return int64(v), true
	case float64:
		return int64(v), true
	case float32:
		return int64(v), true
	case string:
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return i, true
		}
	}
	return 0, false
}
