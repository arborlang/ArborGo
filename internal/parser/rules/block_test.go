package rules

import (
	"github.com/radding/ArborGo/internal/parser/ast"
	"github.com/stretchr/testify/assert"
	"testing"
)

var simpleBlock = `{ a + b;}`
var complexBlock = `{
	const x: number = 1;
	let b: number = 2 + a;
	x();
	z + 4;
}`

func TestCanParseBlock(t *testing.T) {
	assert := assert.New(t)
	p := createParser(simpleBlock)
	block, err := parseBlockRule(p)
	if !assert.NoError(err) {
		t.Fatal()
	}
	blockNode := block.(*ast.Program)
	if !assert.NotNil(blockNode, "could not convert node to a BlockNode") {
		t.Fatal()
	}
	assert.Len(blockNode.Nodes, 1)

	add := blockNode.Nodes[0].(*ast.MathOpNode)
	if !assert.NotNil(add, "couldn't convert the node to MathOp") {
		t.Fatal()
	}

	p = createParser(complexBlock)
	block, err = parseBlockRule(p)
	if !assert.NoError(err) {
		t.Fatal()
	}
	blockNode = block.(*ast.Program)
	if !assert.NotNil(blockNode, "could not convert node to a BlockNode") {
		t.Fatal()
	}
	assert.Len(blockNode.Nodes, 4)
}
