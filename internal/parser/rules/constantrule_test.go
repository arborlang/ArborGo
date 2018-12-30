package rules

import (
	"bytes"
	"github.com/radding/ArborGo/internal/lexer"
	"github.com/radding/ArborGo/internal/parser/ast"
	"github.com/stretchr/testify/assert"
	"testing"
)

var numberConst = `12;`
var floatConst = `123.01;`
var stringConst = `"hello";`
var charConst = `'f';`

func TestConstantWorks(t *testing.T) {
	assert := assert.New(t)
	tokStream := lexer.Lex(bytes.NewReader([]byte(numberConst)))
	parser := New(tokStream)
	constNode, err := ConstantsRule(parser)
	assert.NoError(err)
	assert.NotNil(constNode)
	actNode, ok := constNode.(*ast.Constant)
	if !ok {
		assert.FailNow("can't convert Node to ast.Constant")
	}
	assert.Equal(actNode.Value, "12")
	assert.Equal(actNode.Type, "NUMBER")

	tokStream = lexer.Lex(bytes.NewReader([]byte(floatConst)))
	parser = New(tokStream)
	constNode, err = ConstantsRule(parser)
	assert.NoError(err)
	assert.NotNil(constNode)
	actNode, ok = constNode.(*ast.Constant)
	if !ok {
		assert.FailNow("can't convert Node to ast.Constant")
	}
	assert.Equal(actNode.Value, "123.01")
	assert.Equal(actNode.Type, "FLOAT")

	tokStream = lexer.Lex(bytes.NewReader([]byte(stringConst)))
	parser = New(tokStream)
	constNode, err = ConstantsRule(parser)
	assert.NoError(err)
	assert.NotNil(constNode)
	actNode, ok = constNode.(*ast.Constant)
	if !ok {
		assert.FailNow("can't convert Node to ast.Constant")
	}
	assert.Equal(actNode.Value, "\"hello\"")
	assert.Equal(actNode.Type, "STRINGVAL")

	tokStream = lexer.Lex(bytes.NewReader([]byte(charConst)))
	parser = New(tokStream)
	constNode, err = ConstantsRule(parser)
	assert.NoError(err)
	assert.NotNil(constNode)
	actNode, ok = constNode.(*ast.Constant)
	if !ok {
		assert.FailNow("can't convert Node to ast.Constant")
	}
	assert.Equal(actNode.Value, "'f'")
	assert.Equal(actNode.Type, "CHARVAL")
}
