package rulesv2

import (
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/stretchr/testify/assert"
)

var sliceRuleStr = ":]"
var sliceRuleWithEnd = ":12]"
var sliceRuleWithEndAndStep = ":23:123]"
var sliceRulesWithAbc = ":abc]"
var sliceRulesWithAbcStep = ":abc:xyz]"

var fullIndex = "[0:]"
var fullIndexWithEnd = "[0:-1]"
var fullIndexWithEndAbc = "[0:abc]"
var fullIndexEndWithStep = "[0:abc:-1]"

func TestCanParseSlice(t *testing.T) {
	assert := assert.New(t)
	index := &ast.IndexNode{}
	p := parseTest(sliceRuleStr)
	slice, err := sliceRule(index, p)
	assert.NoError(err)
	assert.NotNil(slice)
	assert.IsType(&ast.SliceNode{}, slice)

	index = &ast.IndexNode{}
	p = parseTest(sliceRuleWithEnd)
	slice, err = sliceRule(index, p)
	assert.NoError(err)
	assert.NotNil(slice)
	assert.IsType(&ast.SliceNode{}, slice)

	index = &ast.IndexNode{}
	p = parseTest(sliceRuleWithEndAndStep)
	slice, err = sliceRule(index, p)
	assert.NoError(err)
	assert.NotNil(slice)
	assert.IsType(&ast.SliceNode{}, slice)

	index = &ast.IndexNode{}
	p = parseTest(sliceRulesWithAbc)
	slice, err = sliceRule(index, p)
	assert.NoError(err)
	assert.NotNil(slice)
	assert.IsType(&ast.SliceNode{}, slice)

	index = &ast.IndexNode{}
	p = parseTest(sliceRulesWithAbcStep)
	slice, err = sliceRule(index, p)
	assert.NoError(err)
	assert.NotNil(slice)
	assert.IsType(&ast.SliceNode{}, slice)
}

func TestCanParseFromIndex(t *testing.T) {
	assert := assert.New(t)
	varName := &ast.VarName{}

	p := parseTest(fullIndex)
	slice, err := indexRule(varName, p)
	assert.NoError(err)
	assert.NotNil(slice)
	if !assert.IsType(&ast.SliceNode{}, slice) {
		t.Fatal()
	}
	realSlice := slice.(*ast.SliceNode)
	assert.IsType(&ast.Constant{}, realSlice.Start)
	assert.Equal(varName, realSlice.Varname)
	assert.Nil(realSlice.End)
	assert.Nil(realSlice.Step)

	p = parseTest(fullIndexWithEnd)
	slice, err = indexRule(varName, p)
	assert.NoError(err)
	assert.NotNil(slice)
	if !assert.IsType(&ast.SliceNode{}, slice) {
		t.Fatal()
	}
	realSlice = slice.(*ast.SliceNode)
	assert.NotNil(realSlice.End)
	assert.NotNil(realSlice.Start)
	assert.Nil(realSlice.Step)

	p = parseTest(fullIndexEndWithStep)
	slice, err = indexRule(varName, p)
	assert.NoError(err)
	assert.NotNil(slice)
	if !assert.IsType(&ast.SliceNode{}, slice) {
		t.Fatal()
	}
	realSlice = slice.(*ast.SliceNode)
	assert.NotNil(realSlice.End)
	assert.NotNil(realSlice.Start)
	assert.NotNil(realSlice.Step)
}
