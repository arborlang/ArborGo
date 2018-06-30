package parser

import (
	"bytes"
	"testing"

	"github.com/radding/ArborGo/internal/ast"
	"github.com/radding/ArborGo/internal/lexer"

	"github.com/stretchr/testify/assert"
)

func Lexer(in string) *lexer.BufferedReader {
	lx := lexer.Lex(bytes.NewReader([]byte(in)))
	return lexer.NewBufferedReader(lx)
}

func TestStatements(t *testing.T) {
	assert := assert.New(t)
	msg := ``

	tree, err := parseStatements(Lexer(msg))

	assert.NoError(err)
	assert.Implements((*ast.Node)(nil), tree)

	//Testing tbat the tree forms with a single node

}

func TestDeclNode(t *testing.T) {
	assert := assert.New(t)
	msg := `let foo: char`

	tree, err := parseDecl(Lexer(msg))

	assert.NoError(err)
	assert.NotNil(tree)

	decl, ok := tree.(*ast.Declaration)
	assert.True(ok)
	assert.Equal("char", decl.TypeToken.Value)
	assert.Equal("foo", decl.NameToken.Value)
	assert.False(decl.IsConst)

	msg = `let foo:: char`

	tree, err = parseDecl(Lexer(msg))
	assert.Error(err)
}

func TestStatementsParsesDecl(t *testing.T) {
	assert := assert.New(t)
	msg := "let foo: char"

	tree, err := parseStatements(Lexer(msg))
	assert.NoError(err)

	program, ok := tree.(*ast.Program)
	assert.True(ok)
	assert.Len(program.Statemetents, 1)

	node, ok := program.Statemetents[0].(*ast.Declaration)
	assert.True(ok)
	assert.Equal("foo", node.NameToken.Value)
}
