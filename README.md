# SQL Parser/Evaluate

A Go project that uses ANTLR4 to parse SQL statements, focusing on WHERE clause evaluation with environment variable matching.

## Features

- Parse SQL WHERE clauses using ANTLR4
- Match WHERE conditions against environment variables
- Evaluate boolean expressions in WHERE clauses
- Support for common SQL operators (=, !=, <, >, <=, >=, LIKE, IN, etc.)

## Installation

```bash
go get github.com/KeArcus/sqlevaluate
```

For development:
```bash
# Clone the repository
git clone https://github.com/KeArcus/sqlevaluate.git
cd sqlevaluate
go mod tidy
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/KeArcus/sqlevaluate/parser"
)

func main() {
    envs := map[string]interface{}{
        "user_id": 123,
        "status": "active",
        "region": "us-west-2",
    }
    
    // Evaluate complete SQL statement
    sql := "SELECT * FROM users WHERE user_id = 123 AND status = 'active'"
    result, err := parser.EvaluateSQL(sql, envs)
    if err != nil {
        panic(err)
    }
    fmt.Printf("SQL result: %v\n", result)
    
    // Evaluate WHERE expression only
    whereExpr := "user_id = 123 AND status = 'active'"
    result2, err := parser.EvaluateWhere(whereExpr, envs)
    if err != nil {
        panic(err)
    }
    fmt.Printf("WHERE result: %v\n", result2)
}
```

## Development

1. Install ANTLR4 tool
2. Generate parser code from grammar
3. Run tests

```bash
# Generate parser from grammar
antlr -Dlanguage=Go -o parser grammar/SQL.g4

# Run examples
go run examples/basic_usage.go
go run examples/functions_example.go

# Run tests
go test ./parser
```
