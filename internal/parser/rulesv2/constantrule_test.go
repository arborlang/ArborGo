package rulesv2

import (
	"bytes"
	"testing"

	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
	"github.com/stretchr/testify/assert"
)

var numberConst = `12;`
var floatConst = `123.01;`
var stringConst = `"hello";`
var charConst = `'f';`

func TestConstantWorks(t *testing.T) {
	assert := assert.New(t)
	tokStream := lexer.Lex(bytes.NewReader([]byte(numberConst)))
	parser := New(tokStream)
	constNode, err := ConstantsRule(parser, true)
	assert.NoError(err)
	assert.NotNil(constNode)
	actNode, ok := constNode.(*ast.Constant)
	if !ok {
		assert.FailNow("can't convert Node to ast.Constant")
	}
	assert.Equal(actNode.Value, "12")
	assert.True(actNode.Type.IsSatisfiedBy(&types.ConstantTypeNode{Name: "Number"}))

	tokStream = lexer.Lex(bytes.NewReader([]byte(floatConst)))
	parser = New(tokStream)
	constNode, err = ConstantsRule(parser, true)
	assert.NoError(err)
	assert.NotNil(constNode)
	actNode, ok = constNode.(*ast.Constant)
	if !ok {
		assert.FailNow("can't convert Node to ast.Constant")
	}
	assert.Equal(actNode.Value, "123.01")
	assert.True(actNode.Type.IsSatisfiedBy(&types.ConstantTypeNode{Name: "Float"}))

	tokStream = lexer.Lex(bytes.NewReader([]byte(stringConst)))
	parser = New(tokStream)
	constNode, err = ConstantsRule(parser, true)
	assert.NoError(err)
	assert.NotNil(constNode)
	actNode, ok = constNode.(*ast.Constant)
	if !ok {
		assert.FailNow("can't convert Node to ast.Constant")
	}
	assert.Equal(actNode.Value, "\"hello\"")
	assert.True(actNode.Type.IsSatisfiedBy(&types.ConstantTypeNode{Name: "String"}))

	tokStream = lexer.Lex(bytes.NewReader([]byte(charConst)))
	parser = New(tokStream)
	constNode, err = ConstantsRule(parser, true)
	assert.NoError(err)
	assert.NotNil(constNode)
	actNode, ok = constNode.(*ast.Constant)
	if !ok {
		assert.FailNow("can't convert Node to ast.Constant")
	}
	assert.Equal(actNode.Value, "'f'")
	assert.True(actNode.Type.IsSatisfiedBy(&types.ConstantTypeNode{Name: "Char"}))
}

func TestCanParseStringSlice(t *testing.T) {
	assert := assert.New(t)
	tokStream := lexer.Lex(bytes.NewReader([]byte(`"string"[0]`)))
	parser := New(tokStream)
	constNode, err := ConstantsRule(parser, true)
	assert.NoError(err)
	assert.NotNil(constNode)
	actNode, ok := constNode.(*ast.IndexNode)
	if !ok {
		assert.FailNow("can't convert Node to ast.IndexNode")
	}
	value, ok := actNode.Varname.(*ast.Constant)
	if !ok {
		assert.FailNow("can't convert Node to ast.Constant")
	}
	assert.Equal(value.Value, "\"string\"")
	assert.True(value.Type.IsSatisfiedBy(&types.ConstantTypeNode{Name: "String"}))

}
