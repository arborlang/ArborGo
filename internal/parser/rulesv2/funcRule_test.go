package rulesv2

import (
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/stretchr/testify/assert"
)

func TestCanFunc(t *testing.T) {
	assert := assert.New(t)

	msg := `fn fooBar(arg: number) -> number {
		return 1;
	}`
	parser := parseTest(msg)

	tree, err := functionDefinitionRule(parser)
	assert.NoError(err)
	assert.NotNil(tree)
	if !assert.IsType(&ast.AssignmentNode{}, tree) {
		return
	}
	assignmentNode := tree.(*ast.AssignmentNode)
	assert.NotNil(assignmentNode.AssignTo)
	assert.NotNil(assignmentNode.Value)
	if !assert.IsType(&ast.FunctionDefinitionNode{}, assignmentNode.Value) {
		return
	}
	if !assert.IsType(&ast.VarName{}, assignmentNode.AssignTo) {
		return
	}
	funcDef := assignmentNode.Value.(*ast.FunctionDefinitionNode)
	assert.NotNil(funcDef.Body)
	assert.Len(funcDef.Arguments, 1)
	assert.NotNil(funcDef.Returns)
	varName := assignmentNode.AssignTo.(*ast.VarName)
	assert.Equal("fooBar", varName.Name)
}

func TestCanParseMultipleArguments(t *testing.T) {
	assert := assert.New(t)

	msg := `fn fooBar(arg: number, arg1: number) -> number {
		return 1;
	}`
	parser := parseTest(msg)

	tree, err := functionDefinitionRule(parser)
	assert.NoError(err)
	assert.NotNil(tree)
	if !assert.IsType(&ast.AssignmentNode{}, tree) {
		return
	}
}

func TestCanParseNoReturnValue(t *testing.T) {
	assert := assert.New(t)

	msg := `fn fooBar(arg: number, arg1: number) {
		return 1;
	}`
	parser := parseTest(msg)

	tree, err := functionDefinitionRule(parser)
	assert.NoError(err)
	assert.NotNil(tree)
	if !assert.IsType(&ast.AssignmentNode{}, tree) {
		return
	}
}

func TestCanParseAnonFunction(t *testing.T) {

	assert := assert.New(t)

	msg := `fn (arg: number, arg1: number) {
		return 1;
	}`
	parser := parseTest(msg)

	tree, err := functionDefinitionRule(parser)
	assert.NoError(err)
	assert.NotNil(tree)
	if !assert.IsType(&ast.FunctionDefinitionNode{}, tree) {
		return
	}
}

func TestCanParseGenericFunction(t *testing.T) {
	assert := assert.New(t)
	msg := `fn <TGeneric>(arg: number, arg1: number) {
		return 1;
	}`
	parser := parseTest(msg)

	tree, err := functionDefinitionRule(parser)
	assert.NoError(err)
	assert.NotNil(tree)
	funcTree, ok := tree.(*ast.FunctionDefinitionNode)
	assert.True(ok)
	assert.Len(funcTree.GenericTypeNames, 1)

	msg = `fn <TGeneric, TGeneric>(arg: number, arg1: number) {
		return 1;
	}`
	parser = parseTest(msg)

	tree, err = functionDefinitionRule(parser)
	assert.NoError(err)
	assert.NotNil(tree)
	funcTree, ok = tree.(*ast.FunctionDefinitionNode)
	assert.True(ok)
	assert.Len(funcTree.GenericTypeNames, 2)
}
