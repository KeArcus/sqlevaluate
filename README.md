# ANTLR4 SQL Parser

A Go project that uses ANTLR4 to parse SQL statements, focusing on WHERE clause evaluation with environment variable matching.

## Features

- Parse SQL WHERE clauses using ANTLR4
- Match WHERE conditions against environment variables
- Evaluate boolean expressions in WHERE clauses
- Support for common SQL operators (=, !=, <, >, <=, >=, LIKE, IN, etc.)

## Installation

```bash
go mod tidy
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/antlr4-sqlparser/parser"
)

func main() {
    envs := map[string]interface{}{
        "user_id": 123,
        "status": "active",
        "region": "us-west-2",
    }
    
    sql := "SELECT * FROM users WHERE user_id = 123 AND status = 'active'"
    result, err := parser.EvaluateWhere(sql, envs)
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("WHERE clause matches: %v\n", result)
}
```

## Development

1. Install ANTLR4 tool
2. Generate parser code from grammar
3. Run tests

```bash
# Generate parser from grammar
antlr4 -Dlanguage=Go -o parser grammar/SQL.g4

# Run tests
go test ./...
```
