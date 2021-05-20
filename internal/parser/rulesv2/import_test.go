package rulesv2

import (
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/stretchr/testify/assert"
)

func TestCanParseImportSimple(t *testing.T) {
	assert := assert.New(t)
	importStmt := `import foobar from "github.com/arborlang/stdlib"`
	p := parseTest(importStmt)
	importNode, err := importRule(p)
	assert.NoError(err)
	assert.NotNil(importNode)
	if !assert.IsType(&ast.ImportNode{}, importNode) {
		return
	}
	node := importNode.(*ast.ImportNode)
	assert.Equal("foobar", node.ImportAs)
	assert.Equal("foobar", node.ExportName)
	assert.Equal(`"github.com/arborlang/stdlib"`, node.Source)
	assert.Nil(node.NextImport)
}

func TestCanParseImportAs(t *testing.T) {
	assert := assert.New(t)
	importStmt := `import foobar as something from "github.com/arborlang/stdlib"`
	p := parseTest(importStmt)
	importNode, err := importRule(p)
	assert.NoError(err)
	assert.NotNil(importNode)
	if !assert.IsType(&ast.ImportNode{}, importNode) {
		return
	}
	node := importNode.(*ast.ImportNode)
	assert.Equal("something", node.ImportAs)
	assert.Equal("foobar", node.ExportName)
	assert.Equal(`"github.com/arborlang/stdlib"`, node.Source)
	assert.Nil(node.NextImport)
}

func TestCanParseImportPieces(t *testing.T) {
	assert := assert.New(t)
	importStmt := `import { foobar } from "github.com/arborlang/stdlib"`
	p := parseTest(importStmt)
	importNode, err := importRule(p)
	assert.NoError(err)
	assert.NotNil(importNode)
	if !assert.IsType(&ast.ImportNode{}, importNode) {
		return
	}
	node := importNode.(*ast.ImportNode)
	assert.Equal("", node.ImportAs)
	assert.Equal("", node.ExportName)
	assert.Equal(`"github.com/arborlang/stdlib"`, node.Source)
	assert.NotNil(node.NextImport)
	if !assert.IsType(&ast.ImportNode{}, node.NextImport) {
		return
	}
	assert.Equal("foobar", node.NextImport.ImportAs)
	assert.Equal("foobar", node.NextImport.ExportName)
	assert.Nil(node.NextImport.NextImport)
}

func TestCanParseImportPiecesAliased(t *testing.T) {
	assert := assert.New(t)
	importStmt := `import { foobar as barbaz } from "github.com/arborlang/stdlib"`
	p := parseTest(importStmt)
	importNode, err := importRule(p)
	assert.NoError(err)
	assert.NotNil(importNode)
	if !assert.IsType(&ast.ImportNode{}, importNode) {
		return
	}
	node := importNode.(*ast.ImportNode)
	assert.Equal("", node.ImportAs)
	assert.Equal("", node.ExportName)
	assert.Equal(`"github.com/arborlang/stdlib"`, node.Source)
	assert.NotNil(node.NextImport)
	if !assert.IsType(&ast.ImportNode{}, node.NextImport) {
		return
	}
	assert.Equal("barbaz", node.NextImport.ImportAs)
	assert.Equal("foobar", node.NextImport.ExportName)
	assert.Nil(node.NextImport.NextImport)
}

func TestCanParseImportMultiples(t *testing.T) {
	assert := assert.New(t)
	importStmt := `import { foobar, bar } from "github.com/arborlang/stdlib"`
	p := parseTest(importStmt)
	importNode, err := importRule(p)
	assert.NoError(err)
	assert.NotNil(importNode)
	if !assert.IsType(&ast.ImportNode{}, importNode) {
		return
	}
	node := importNode.(*ast.ImportNode)
	assert.Equal("", node.ImportAs)
	assert.Equal("", node.ExportName)
	assert.Equal(`"github.com/arborlang/stdlib"`, node.Source)
	assert.NotNil(node.NextImport)
	if !assert.IsType(&ast.ImportNode{}, node.NextImport) {
		return
	}
	assert.Equal("foobar", node.NextImport.ImportAs)
	assert.Equal("foobar", node.NextImport.ExportName)
	assert.NotNil(node.NextImport.NextImport)
	assert.Equal("bar", node.NextImport.NextImport.ImportAs)
	assert.Equal("bar", node.NextImport.NextImport.ExportName)
	assert.Nil(node.NextImport.NextImport.NextImport)
}
