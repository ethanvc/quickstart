package main

import (
	"github.com/k0kubun/sqldef/database"
	"github.com/k0kubun/sqldef/parser"
	"github.com/k0kubun/sqldef/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSqlDiffMain(t *testing.T) {
	currentSql := `
CREATE TABLE new_table (
    id INT PRIMARY KEY,
    name VARCHAR(100)
)
`
	desireSql := `
CREATE TABLE new_table (
    id INT PRIMARY KEY,
    name VARCHAR(100),
    age int
)
`
	sqlParser := database.NewParser(parser.ParserModeMysql)
	result, err := schema.GenerateIdempotentDDLs(
		schema.GeneratorModeMysql, sqlParser, desireSql, currentSql, database.GeneratorConfig{}, "")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))
}
