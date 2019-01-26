package rules

import (
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
	"github.com/stretchr/testify/assert"
)

var one = "type"
var two = "type | number;"
var three = "type | number | string,"
var four = "type ->"

var functionType = "(number, string | char, name) -> bool;"
var functionType2 = "string | (number, string | char, name) -> bool;"

func TestCanParseTypes(t *testing.T) {
	assert := assert.New(t)
	p := createParser(one)
	node, err := typeRules(p)
	if !assert.NoError(err) {
		t.Fatal()
	}
	tp := node.(*ast.TypeNode)
	if !assert.NotNil(tp, "Could not convert to TypeNode") {
		t.Fatal()
	}
	assert.Len(tp.Types, 1)
	assert.Equal("type", tp.Types[0])

	p = createParser(two)
	node, err = typeRules(p)
	if !assert.NoError(err) {
		t.Fatal()
	}
	tp = node.(*ast.TypeNode)
	if !assert.NotNil(tp, "Could not convert to TypeNode") {
		t.Fatal()
	}
	assert.Len(tp.Types, 2)
	assert.Equal("type", tp.Types[0])
	assert.Equal("number", tp.Types[1])

	p = createParser(three)
	node, err = typeRules(p)
	if !assert.NoError(err) {
		t.Fatal()
	}
	tp = node.(*ast.TypeNode)
	if !assert.NotNil(tp, "Could not convert to TypeNode") {
		t.Fatal()
	}
	assert.Len(tp.Types, 3)
	assert.Equal("type", tp.Types[0])
	assert.Equal("number", tp.Types[1])
	assert.Equal("string", tp.Types[2])

	p = createParser(four)
	node, err = typeRules(p)
	if !assert.NoError(err) {
		t.Fatal()
	}
	tp = node.(*ast.TypeNode)
	if !assert.NotNil(tp, "Could not convert to TypeNode") {
		t.Fatal()
	}
	assert.Len(tp.Types, 1)
	assert.Equal("type", tp.Types[0])

	p = createParser(functionType)
	node, err = typeRules(p)
	if !assert.NoError(err) {
		t.Fatal()
	}
	tp = node.(*ast.TypeNode)
	if !assert.NotNil(tp, "Could not convert to TypeNode") {
		t.Fatal()
	}
	assert.Len(tp.Types, 1)
	assert.Equal("func(number,string|char,name)bool", tp.Types[0])

	p = createParser(functionType2)
	node, err = typeRules(p)
	if !assert.NoError(err) {
		t.Fatal()
	}
	tp = node.(*ast.TypeNode)
	if !assert.NotNil(tp, "Could not convert to TypeNode") {
		t.Fatal()
	}
	assert.Len(tp.Types, 2)
	assert.Equal("func(number,string|char,name)bool", tp.Types[1])

}

func TestCanParseFunctionReturnType(t *testing.T) {
	assert := assert.New(t)
	p := createParser(functionType)
	assert.Equal(tokens.RPAREN, p.Next().Token)
	funcDef, err := parseFunctionDefinitionType(p)
	assert.NoError(err)
	assert.Equal("func(number,string|char,name)bool", funcDef)
}
