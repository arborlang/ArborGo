package typevisitor

import (
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/stretchr/testify/assert"
)

func TestCanVisitConstantNode(t *testing.T) {
	assert := assert.New(t)

	constantNode := &ast.Constant{}

	v := New()
	ret, err := constantNode.Accept(v)
	assert.NoError(err)

	annotated, ok := ret.(*annotatedTypeNode)
	assert.True(ok)

	assert.Equal(constantNode, annotated.node)

}
