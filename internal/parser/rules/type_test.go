package rules

import (
	"github.com/radding/ArborGo/internal/parser/ast"
	"github.com/stretchr/testify/assert"
	"testing"
)

var one = "type"
var two = "type | number"
var three = "type | number | string"

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
}
