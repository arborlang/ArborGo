package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsATypeNode(t *testing.T) {
	assert := assert.New(t)

	assert.Implements((*TypeNode)(nil), new(TypeGuard))
}

func TestIsSatisfiedByOne(t *testing.T) {
	assert := assert.New(t)

	tg := &TypeGuard{
		Types: []TypeNode{
			&ConstantTypeNode{Name: "Floop"},
			&ConstantTypeNode{Name: "Bloop"},
			&ConstantTypeNode{Name: "bar"},
		},
	}

	assert.True(tg.IsSatisfiedBy(&ConstantTypeNode{Name: "bar"}))
	assert.True(tg.IsSatisfiedBy(&ConstantTypeNode{Name: "Bloop"}))
	assert.False(tg.IsSatisfiedBy(&ConstantTypeNode{Name: "Bloops"}))
}

func TestIsSatisfiedByAnotherTypeGuard(t *testing.T) {
	assert := assert.New(t)

	tg := &TypeGuard{
		Types: []TypeNode{
			&ConstantTypeNode{Name: "Floop"},
			&ConstantTypeNode{Name: "Bloop"},
			&ConstantTypeNode{Name: "bar"},
		},
	}
	tg2 := &TypeGuard{
		Types: []TypeNode{
			&ConstantTypeNode{Name: "Bloop"},
			&ConstantTypeNode{Name: "bar"},
			&ConstantTypeNode{Name: "Floop"},
		},
	}

	assert.True(tg.IsSatisfiedBy(tg2))
	assert.False(tg.IsSatisfiedBy(&ConstantTypeNode{Name: "Bloops"}))
	assert.False(tg2.IsSatisfiedBy(&ConstantTypeNode{Name: "Bloops"}))
}
