package rulesv2

import (
	"bytes"
	"testing"

	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
	"github.com/stretchr/testify/assert"
)

var simpleDecl = "let abc: string;"
var constDecl = "const abc: float;"
var constDeclNoTypes = "const abc;"

var declWithEq = "let abc: number = 123;"
var declWithEqNoType = "let abc = 123;"

func TestParseDeclStandalone(t *testing.T) {
	// Test that let works
	assert := assert.New(t)
	tokStream := lexer.Lex(bytes.NewReader([]byte(simpleDecl)))
	parser := New(tokStream)
	statement, err := DeclRule(parser)
	assert.NoError(err)
	assert.NotNil(statement)

	decl, ok := statement.(*ast.DeclNode)
	assert.True(ok, "Could not convert node to DeclNode")
	assert.NotNil(decl)
	assert.Equal("abc", decl.Varname.Name)
	assert.False(decl.IsConstant)

	// Test that const works
	tokStream = lexer.Lex(bytes.NewReader([]byte(constDecl)))
	parser = New(tokStream)
	statement, err = DeclRule(parser)
	assert.NoError(err)
	assert.NotNil(statement)

	decl, ok = statement.(*ast.DeclNode)
	assert.True(ok, "Could not convert node to DeclNode")
	assert.NotNil(decl)
	assert.Equal("abc", decl.Varname.Name)
	assert.True(decl.IsConstant)
}

func TestParseDeclInExpression(t *testing.T) {
	assert := assert.New(t)
	tokStream := lexer.Lex(bytes.NewReader([]byte(simpleDecl)))
	parser := New(tokStream)
	statement, err := StatementRule(parser)
	assert.NoError(err)
	assert.NotNil(statement)

	decl, ok := statement.(*ast.DeclNode)
	assert.True(ok, "Could not convert node to DeclNode")
	assert.NotNil(decl)
	assert.Equal("abc", decl.Varname.Name)
	assert.False(decl.IsConstant)
}

func TestThisIsValidProgram(t *testing.T) {
	assert := assert.New(t)
	tokStream := lexer.Lex(bytes.NewReader([]byte(simpleDecl)))
	parser := New(tokStream)
	statement, err := ProgramRule(parser, tokens.EOF)
	assert.NoError(err)
	assert.NotNil(statement)
	prog, ok := statement.(*ast.Program)

	if !ok {
		assert.FailNow("Could not convert node to Prog")
	}

	assert.Len(prog.Nodes, 1)

	decl, ok := prog.Nodes[0].(*ast.DeclNode)
	if !ok {
		assert.FailNow("Could not convert node to Decl Node")
	}

	assert.NotNil(decl)
	assert.Equal("abc", decl.Varname.Name)
	assert.False(decl.IsConstant)
}

func TestAssignmentInDecl(t *testing.T) {
	assert := assert.New(t)
	tokStream := lexer.Lex(bytes.NewReader([]byte(declWithEq)))
	parser := New(tokStream)
	statement, err := DeclRule(parser)
	assert.NoError(err)
	assert.NotNil(statement)

	assign, ok := statement.(*ast.AssignmentNode)
	if !ok {
		assert.FailNow("failed to get assignment node")
	}
	assert.NotNil(assign)
	assert.NotNil(assign.AssignTo)

	decl, ok := assign.AssignTo.(*ast.DeclNode)
	if !ok {
		assert.FailNow("failed to convert the AssignTo to a decl node")
	}
	assert.NotNil(decl)
	assert.Equal(decl.Varname.Name, "abc")
}

func TestAssignmentInDeclNoTypes(t *testing.T) {
	assert := assert.New(t)
	tokStream := lexer.Lex(bytes.NewReader([]byte(declWithEqNoType)))
	parser := New(tokStream)
	statement, err := DeclRule(parser)
	assert.NoError(err)
	assert.NotNil(statement)

	assign, ok := statement.(*ast.AssignmentNode)
	if !ok {
		assert.FailNow("failed to get assignment node")
	}
	assert.NotNil(assign)
	assert.NotNil(assign.AssignTo)

	decl, ok := assign.AssignTo.(*ast.DeclNode)
	if !ok {
		assert.FailNow("failed to convert the AssignTo to a decl node")
	}
	assert.NotNil(decl)
	assert.Equal(decl.Varname.Name, "abc")
}

func TestFailsIfConstWithNoTypes(t *testing.T) {
	assert := assert.New(t)
	tokStream := lexer.Lex(bytes.NewReader([]byte(constDeclNoTypes)))
	parser := New(tokStream)
	_, err := DeclRule(parser)
	assert.Error(err)
}
