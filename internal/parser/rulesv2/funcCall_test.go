package rulesv2

import (
	"github.com/arborlang/ArborGo/internal/parser/ast"
	// "github.com/arborlang/ArborGo/internal/tokens"
	"testing"

	"github.com/stretchr/testify/assert"
)

var simpleCall = "(1 + 2, b, c * 12, x())"
var complicatedCall = "(fn() -> bool { return true; })"
var fullCall = "abc123(1 + 2, b, c * 12, x())"

func TestCanParseAFunctionCall(t *testing.T) {
	assert := assert.New(t)
	varName := &ast.VarName{}
	p := parseTest(simpleCall)
	varName.Name = "abc123"
	fun, err := functionCallRule(varName, p, true)
	if !assert.NoError(err) {
		t.Fatal()
	}
	funcCall := fun.(*ast.FunctionCallNode)
	if !assert.NotNil(funcCall) {
		t.Fatal()
	}
	assert.NotNil(funcCall.Arguments)
	assert.Len(funcCall.Arguments, 4)
	nm := funcCall.Definition.(*ast.VarName)
	assert.Equal("abc123", nm.Name)

	// Ensure that first arg is parsed as a MathNode
	firstArg := funcCall.Arguments[0].(*ast.MathOpNode)
	assert.NotNil(firstArg)

	//Ensure that second arg is parsed as a VarNameNode
	secondArg := funcCall.Arguments[1].(*ast.VarName)
	assert.NotNil(secondArg)

	thirdArg := funcCall.Arguments[2].(*ast.MathOpNode)
	assert.NotNil(thirdArg)

	fourthArg := funcCall.Arguments[3].(*ast.FunctionCallNode)
	assert.NotNil(fourthArg)

	p = parseTest(complicatedCall)

	_, err = functionCallRule(varName, p, false)
	if !assert.NoError(err) {
		t.Fatal()
	}
}

func TestParseFullCall(t *testing.T) {
	assert := assert.New(t)
	p := parseTest(fullCall)
	fun, err := varNameRule(p, false)
	if !assert.NoError(err) {
		t.Fatal()
	}
	funcCall := fun.(*ast.FunctionCallNode)
	if !assert.NotNil(funcCall, "can not convert Node to a FunctionCallNode") {
		t.Fatal()
	}
}
