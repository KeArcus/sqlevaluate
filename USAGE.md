# Usage Guide

## Overview

This project provides a SQL WHERE clause parser and evaluator using ANTLR4. It allows you to:

1. Parse SQL statements and extract WHERE conditions
2. Evaluate WHERE conditions against environment variables 
3. Support for complex expressions with AND, OR, parentheses
4. Handle various SQL operators: =, !=, <, >, <=, >=, LIKE, IN, IS NULL, etc.

## Quick Start

```go
package main

import (
    "fmt"
    "KeArcus/sqlevaluate/parser"
)

func main() {
    // Define your environment variables
    envs := map[string]interface{}{
        "user_id": 123,
        "status":  "active", 
        "age":     25,
    }

    // Evaluate a complete SQL statement
    sql := "SELECT * FROM users WHERE user_id = 123 AND status = 'active'"
    result, err := parser.EvaluateSQL(sql, envs)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Query matches: %v\n", result) // Output: true

    // Or evaluate just the WHERE expression
    whereExpr := "age > 20 AND status = 'active'"
    result2, err := parser.EvaluateWhere(whereExpr, envs)
    if err != nil {
        panic(err)
    }
    fmt.Printf("WHERE matches: %v\n", result2) // Output: true
}
```

## Supported Operations

### Comparison Operators
- `=` - Equal
- `!=`, `<>` - Not equal  
- `<` - Less than
- `<=` - Less than or equal
- `>` - Greater than
- `>=` - Greater than or equal

### Logical Operators
- `AND` - Logical AND
- `OR` - Logical OR
- `()` - Parentheses for grouping

### Special Operators
- `IN (value1, value2, ...)` - Value in list
- `NOT IN (value1, value2, ...)` - Value not in list
- `LIKE 'pattern'` - Pattern matching with % and _
- `NOT LIKE 'pattern'` - Negated pattern matching
- `IS NULL` - Null check
- `IS NOT NULL` - Not null check

### SQL Functions

#### String Functions
- `UPPER(string)` - Convert to uppercase
- `LOWER(string)` - Convert to lowercase
- `LENGTH(string)` - Get string length
- `SUBSTR(string, start [, length])` - Extract substring
- `TRIM(string)` - Remove leading/trailing whitespace
- `CONCAT(string1, string2, ...)` - Concatenate strings
- `REPLACE(string, old, new)` - Replace substrings
- `CONTAINS(string, substring)` - Check if string contains substring

#### Math Functions
- `ABS(number)` - Absolute value
- `ROUND(number [, precision])` - Round number
- `CEIL(number)` - Round up to nearest integer
- `FLOOR(number)` - Round down to nearest integer
- `MOD(number, divisor)` - Modulo operation

#### Date/Time Functions
- `NOW()` - Current timestamp
- `DATE(datetime)` - Extract date part
- `YEAR(date)` - Extract year
- `MONTH(date)` - Extract month
- `DAY(date)` - Extract day

#### NULL Handling Functions
- `COALESCE(value1, value2, ...)` - Return first non-null value
- `NULLIF(value1, value2)` - Return NULL if values are equal
- `ISNULL(value [, replacement])` - Check for NULL or replace NULL

### Data Types
- **Strings**: `'hello'`, `"world"`
- **Numbers**: `123`, `45.67`
- **Booleans**: `TRUE`, `FALSE`
- **NULL**: `NULL`

## Examples

### Basic Comparisons
```go
envs := map[string]interface{}{
    "user_id": 123,
    "name":    "John",
    "score":   85.5,
}

// String comparison
parser.EvaluateWhere("name = 'John'", envs) // true

// Numeric comparison  
parser.EvaluateWhere("user_id > 100", envs) // true
parser.EvaluateWhere("score >= 85.0", envs) // true
```

### Logical Operations
```go
// AND operation
parser.EvaluateWhere("user_id = 123 AND name = 'John'", envs) // true

// OR operation
parser.EvaluateWhere("user_id = 456 OR name = 'John'", envs) // true

// Parentheses grouping
parser.EvaluateWhere("(user_id = 123 OR user_id = 456) AND name = 'John'", envs) // true
```

### Pattern Matching
```go
envs := map[string]interface{}{
    "email": "john@example.com",
    "name":  "John Doe",
}

// Prefix matching
parser.EvaluateWhere("name LIKE 'John%'", envs) // true

// Suffix matching  
parser.EvaluateWhere("email LIKE '%example.com'", envs) // true

// Contains matching
parser.EvaluateWhere("name LIKE '%Doe%'", envs) // true
```

### IN Operations
```go
envs := map[string]interface{}{
    "region": "us-west-2",
    "status": "active",
}

// IN operation
parser.EvaluateWhere("region IN ('us-west-2', 'us-east-1')", envs) // true

// NOT IN operation
parser.EvaluateWhere("status NOT IN ('pending', 'disabled')", envs) // true
```

### NULL Checks
```go
envs := map[string]interface{}{
    "name":        "John",
    "description": nil,
}

// IS NULL
parser.EvaluateWhere("description IS NULL", envs) // true

// IS NOT NULL
parser.EvaluateWhere("name IS NOT NULL", envs) // true
```

### SQL Functions
```go
envs := map[string]interface{}{
    "name":         "john doe",
    "email":        "john@example.com",
    "salary":       85750.50,
    "created_date": "2022-03-15",
    "bonus":        nil,
}

// String functions
parser.EvaluateWhere("UPPER(name) = 'JOHN DOE'", envs) // true
parser.EvaluateWhere("LENGTH(name) = 8", envs) // true  
parser.EvaluateWhere("SUBSTR(name, 1, 4) = 'john'", envs) // true
parser.EvaluateWhere("CONTAINS(email, '@example')", envs) // true

// Math functions
parser.EvaluateWhere("ROUND(salary) = 85751", envs) // true
parser.EvaluateWhere("ABS(salary) > 80000", envs) // true

// Date functions
parser.EvaluateWhere("YEAR(created_date) = 2022", envs) // true
parser.EvaluateWhere("MONTH(created_date) >= 3", envs) // true

// NULL handling
parser.EvaluateWhere("COALESCE(bonus, 0) = 0", envs) // true
parser.EvaluateWhere("ISNULL(bonus) = TRUE", envs) // true

// Complex nested functions
parser.EvaluateWhere("LENGTH(UPPER(TRIM(name))) = 8", envs) // true
parser.EvaluateWhere("YEAR(created_date) = 2022 AND ROUND(salary) > 85000", envs) // true
```

## Error Handling

The parser will return an error if:
- SQL syntax is invalid
- Grammar parsing fails
- Internal evaluation errors occur

Always check for errors:

```go
result, err := parser.EvaluateSQL(sql, envs)
if err != nil {
    log.Printf("SQL evaluation failed: %v", err)
    return
}
```

## Testing

Run the test suite:

```bash
go test ./parser -v
```

## Performance Notes

- The parser creates a new ANTLR parse tree for each evaluation
- For high-frequency operations, consider caching parsed expressions
- Environment variable lookups are O(1) using Go maps
- Complex expressions with many AND/OR operations may impact performance
