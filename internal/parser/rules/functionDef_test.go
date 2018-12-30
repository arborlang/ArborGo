package rules

import (
	"github.com/radding/ArborGo/internal/parser/ast"
	"github.com/radding/ArborGo/internal/tokens"
	"github.com/stretchr/testify/assert"
	"testing"
)

var funcStr1 = `
(a: number, b: number) -> {
	a + b;	
};`

var funcStrFull = `
let x = (a: number, b: number) -> {
	a + b;
};`

var funcStr2 = `(a: number, b: number) -> a + b;`

func TestCanParseFunctionDef(t *testing.T) {
	assert := assert.New(t)
	p := createParser(funcStr1)

	fun, err := functionDefinitionRule(p)
	if !assert.NoError(err) {
		t.Fatal()
	}
	assert.NotNil(fun)
	funcNode := fun.(*ast.FunctionDefinitionNode)
	if !assert.NotNil(funcNode, "Failed to convert to a function def node") {
		t.Fatal()
	}
	assert.Len(funcNode.Arguments, 2)
	block := funcNode.Body.(*ast.Program)
	if !assert.NotNil(block, "could not convert block to Program node") {
		t.Fatal()
	}
	assert.Len(block.Nodes, 1)

	// Now test a simple function
	p = createParser(funcStr2)
	fun, err = functionDefinitionRule(p)
	if !assert.NoError(err) {
		t.Fatal()
	}

	assert.NotNil(fun)
	funcNode = fun.(*ast.FunctionDefinitionNode)
	if !assert.NotNil(funcNode, "Failed to convert to a function def node") {
		t.Fatal()
	}
	assert.Len(funcNode.Arguments, 2)
	block, ok := funcNode.Body.(*ast.Program)
	if !assert.False(ok, "converted block to a program node?") {
		t.Fatal()
	}
	expr := funcNode.Body.(*ast.MathOpNode)
	if !assert.NotNil(expr, "Failed to convert to a MathOpNode") {
		t.Fatal()
	}

}

func TestCanBeAValidProgram(t *testing.T) {
	assert := assert.New(t)
	p := createParser(funcStrFull)

	prog, err := ProgramRule(p, tokens.EOF)
	if !assert.NoError(err) {
		t.Fatal()
	}
	assert.NotNil(prog)

	p = createParser(funcStr2)
	prog, err = ProgramRule(p, tokens.EOF)
	if !assert.NoError(err) {
		t.Fatal()
	}
	assert.NotNil(prog)

	p = createParser(funcStr1)
	prog, err = ProgramRule(p, tokens.EOF)
	if !assert.NoError(err) {
		t.Fatal()
	}
	assert.NotNil(prog)
}
