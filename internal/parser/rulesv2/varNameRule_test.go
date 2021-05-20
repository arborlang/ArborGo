package rulesv2

import (
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
	"github.com/stretchr/testify/assert"
)

func TestCanParseAVarName(t *testing.T) {
	assert := assert.New(t)

	varName := `someName`
	p := parseTest(varName)
	node, err := varNameRule(p, false)
	realNode, ok := node.(*ast.VarName)
	assert.NoError(err)
	assert.True(ok)
	assert.Equal(realNode.Name, "someName")
}

func TestCanParseVarNameTypes(t *testing.T) {
	assert := assert.New(t)

	varName := `someName: Integer`
	p := parseTest(varName)
	node, err := varNameRule(p, true)
	realNode, ok := node.(*ast.VarName)
	assert.NoError(err)
	assert.True(ok)
	assert.Equal(realNode.Name, "someName")
	assert.True(realNode.Type.IsSatisfiedBy(&types.ConstantTypeNode{Name: "Integer"}))

	p = parseTest(varName)
	node, err = varNameRule(p, false)
	assert.Error(err)
}
