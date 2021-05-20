package rulesv2

import (
	"reflect"
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
	"github.com/stretchr/testify/assert"
)

var pipeSimple = `|> bash() |> repo`

var pipeWithVarName = `xyz |> bash |> repo;`
var pipeWithString = `"xyz" |> bash |> repo;`
var pipeWithNumber = `123 |> bash |> repo;`

func TestSimplePipe(t *testing.T) {
	assert := assert.New(t)
	varname := &ast.VarName{}
	p := parseTest(pipeSimple)
	pipe, err := pipeRule(varname, p)
	if !assert.NoError(err) {
		t.Fatal()
	}
	assert.NotNil(pipe)
	pipeNode, _ := pipe.(*ast.PipeNode)
	if !assert.NotNil(pipeNode, "failed to convert node to pipe to PipeNode, got %s", reflect.TypeOf(pipe)) {
		t.Fatal()
	}
	rightSide, _ := pipeNode.RightSide.(*ast.PipeNode)
	if !assert.NotNil(rightSide, "failed to convert node to pipe to PipeNode, got %s", reflect.TypeOf(pipeNode.RightSide)) {
		t.Fail()
	}
	leftSide, _ := rightSide.LeftSide.(*ast.FunctionCallNode)
	if !assert.NotNil(leftSide, "failed to convert leftside to functionCallNode, got %s", reflect.TypeOf(pipeNode.LeftSide)) {
		t.Fail()
	}
	rightSideVarName, _ := rightSide.RightSide.(*ast.VarName)
	if !assert.NotNil(rightSideVarName, "failed to convert right side to a varname, got %s", reflect.TypeOf(rightSide.RightSide)) {
		t.Fatal()
	}
}

func TestCanParseAsExpression(t *testing.T) {
	assert := assert.New(t)
	p := parseTest(pipeWithNumber)
	_, err := ProgramRule(p, tokens.EOF)
	assert.NoError(err)

	p = parseTest(pipeWithString)
	_, err = ProgramRule(p, tokens.EOF)
	assert.NoError(err)

	p = parseTest(pipeWithVarName)
	_, err = ProgramRule(p, tokens.EOF)
	assert.NoError(err)
}
