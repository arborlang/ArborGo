package rules

import (
	"bytes"
	"github.com/radding/ArborGo/internal/lexer"
	"github.com/radding/ArborGo/internal/parser/ast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createParser(msg string) *Parser {
	tokStream := lexer.Lex(bytes.NewReader([]byte(msg)))
	parser := New(tokStream)
	return parser
}

func aaa(t *testing.T) {
	// Test that
	sampleDecl := "let abc"
	assert := assert.New(t)
	tokStream := lexer.Lex(bytes.NewReader([]byte(sampleDecl)))
	parser := New(tokStream)
	statement, err := DeclRule(parser)
	assert.NoError(err)
	assert.NotNil(statement)

	decl, ok := statement.(*ast.DeclNode)
	assert.True(ok, "Could not convert node to DeclNode")
	assert.NotNil(decl)
	assert.Equal("abc", decl.Varname.Name)
	assert.False(decl.IsConstant)
}
