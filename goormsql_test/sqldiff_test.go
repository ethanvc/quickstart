package main

import (
	"github.com/k0kubun/sqldef/database"
	"github.com/k0kubun/sqldef/parser"
	"github.com/k0kubun/sqldef/schema"
	"github.com/schemalex/schemalex/diff"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestSqlDiffSqldef(t *testing.T) {
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

func TestSqlDiffSchemalex(t *testing.T) {

	const sql1 = `
CREATE TABLE hoge (
    id INTEGER NOT NULL AUTO_INCREMENT,
    PRIMARY KEY (id)
);`
	const sql2 = `
CREATE TABLE hoge (
    id INTEGER NOT NULL AUTO_INCREMENT,
    c VARCHAR (20) NOT NULL DEFAULT "hoge",
    PRIMARY KEY (id)
);
`

	err := diff.Strings(os.Stdout, sql1, sql2)
	require.NoError(t, err)
}
