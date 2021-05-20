package rulesv2

import (
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/stretchr/testify/assert"
)

var indexConst = "[123]"
var indexVal1 = "[abc]"
var indexCanAdd = "[123 + abc]"
var index2 = "[foobar()]"

var exprIndex = "foobar[1234]"

func TestCanParseIndex(t *testing.T) {
	assert := assert.New(t)
	p := parseTest(indexConst)
	varName := &ast.VarName{
		Name: "Foobar",
	}
	ndxNode, err := indexRule(varName, p)
	if !assert.NoError(err, "Error is defined") {
		t.Fatal()
	}
	ndx, _ := ndxNode.(*ast.IndexNode)
	if !assert.IsType(&ast.IndexNode{}, ndx) {
		t.Fatal()
	}
	assert.NotNil(ndx)
	assert.Equal(varName, ndx.Varname)
	assert.IsType(&ast.Constant{}, ndx.Index)
}

func TestCanParseVarNameIndex(t *testing.T) {
	assert := assert.New(t)
	p := parseTest(indexVal1)
	varName := &ast.VarName{
		Name: "Foobar",
	}
	ndxNode, err := indexRule(varName, p)
	if !assert.NoError(err, "Error is defined") {
		t.Fatal()
	}
	ndx, _ := ndxNode.(*ast.IndexNode)
	if !assert.IsType(&ast.IndexNode{}, ndx) {
		t.Fatal()
	}
	assert.NotNil(ndx)
	assert.Equal(varName, ndx.Varname)
	assert.IsType(&ast.VarName{}, ndx.Index)
}

func TestCanParseAdditionIndex(t *testing.T) {
	assert := assert.New(t)
	p := parseTest(indexCanAdd)
	varName := &ast.VarName{
		Name: "Foobar",
	}
	ndxNode, err := indexRule(varName, p)
	if !assert.NoError(err, "Error is defined") {
		t.Fatal()
	}
	ndx, _ := ndxNode.(*ast.IndexNode)
	if !assert.IsType(&ast.IndexNode{}, ndx) {
		t.Fatal()
	}
	assert.NotNil(ndx)
	assert.Equal(varName, ndx.Varname)
	assert.IsType(&ast.MathOpNode{}, ndx.Index)
}

func TestCanParseFunctionCallIndex(t *testing.T) {
	assert := assert.New(t)
	p := parseTest(index2)
	varName := &ast.VarName{
		Name: "Foobar",
	}
	ndxNode, err := indexRule(varName, p)
	if !assert.NoError(err, "Error is defined") {
		t.Fatal()
	}
	ndx, _ := ndxNode.(*ast.IndexNode)
	if !assert.IsType(&ast.IndexNode{}, ndx) {
		t.Fatal()
	}
	assert.NotNil(ndx)
	assert.Equal(varName, ndx.Varname)
	assert.IsType(&ast.FunctionCallNode{}, ndx.Index)
}

func TestCanParseIndexFromExpression(t *testing.T) {
	assert := assert.New(t)
	p := parseTest(exprIndex)
	va, err := ExpressionRule(p)
	assert.NoError(err)
	assert.IsType(&ast.IndexNode{}, va)

}
