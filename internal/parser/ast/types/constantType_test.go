package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstantTypeIsTypeNode(t *testing.T) {
	assert := assert.New(t)

	assert.Implements((*TypeNode)(nil), new(ConstantTypeNode))
}

func TestCanBeSatisfiedByAnotherConstantType(t *testing.T) {
	assert := assert.New(t)

	c1 := &ConstantTypeNode{Name: "FooBar"}
	c2 := &ConstantTypeNode{Name: "FooBar"}
	assert.True(c1.IsSatisfiedBy(c2))
	assert.True(c2.IsSatisfiedBy(c1))

	c2.Name = "bloop"
	assert.False(c1.IsSatisfiedBy(c2))
	assert.False(c2.IsSatisfiedBy(c1))
}

type NonConstant struct {
}

func (nc *NonConstant) IsSatisfiedBy(n TypeNode) bool {
	return false
}
func (nc *NonConstant) String() string {
	return "narp"
}

func TestNonConstantNodeFailsImmediately(t *testing.T) {
	assert := assert.New(t)

	c1 := &ConstantTypeNode{Name: "FooBar"}
	n := &NonConstant{}
	assert.False(c1.IsSatisfiedBy(n))
}
