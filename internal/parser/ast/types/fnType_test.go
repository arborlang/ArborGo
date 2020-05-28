package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFnTypeIsATypeNode(t *testing.T) {
	assert := assert.New(t)

	assert.Implements((*TypeNode)(nil), new(FnType))
}

func TestFnIsSatisfiedByAnotherFn(t *testing.T) {
	assert := assert.New(t)

	fn1 := &FnType{
		Parameters: []TypeNode{
			&ConstantTypeNode{Name: "Bloop"},
			&TypeGuard{
				Types: []TypeNode{
					&ConstantTypeNode{Name: "Bloop"},
					&ConstantTypeNode{Name: "Blarp"},
				},
			},
		},
		ReturnVal: &ConstantTypeNode{Name: "Flop"},
	}
	fn2 := &FnType{
		Parameters: []TypeNode{
			&ConstantTypeNode{Name: "Bloop"},
			&TypeGuard{
				Types: []TypeNode{
					&ConstantTypeNode{Name: "Bloop"},
					&ConstantTypeNode{Name: "Blarp"},
				},
			},
		},
		ReturnVal: &ConstantTypeNode{Name: "Flop"},
	}

	assert.True(fn1.IsSatisfiedBy(fn2))
}

func TestFnIsSatisfiedByAnotherFnWithADifferentParam(t *testing.T) {
	assert := assert.New(t)

	fn1 := &FnType{
		Parameters: []TypeNode{
			&ConstantTypeNode{Name: "Bloop"},
			&TypeGuard{
				Types: []TypeNode{
					&ConstantTypeNode{Name: "Bloop"},
					&ConstantTypeNode{Name: "Blarp"},
				},
			},
		},
		ReturnVal: &ConstantTypeNode{Name: "Flop"},
	}
	fn2 := &FnType{
		Parameters: []TypeNode{
			&ConstantTypeNode{Name: "Bloop"},
			&ConstantTypeNode{Name: "Blarp"},
		},
		ReturnVal: &ConstantTypeNode{Name: "Flop"},
	}

	assert.True(fn1.IsSatisfiedBy(fn2))
}

func TestFnRejectsIfParamsDontMatch(t *testing.T) {
	assert := assert.New(t)

	fn1 := &FnType{
		Parameters: []TypeNode{
			&ConstantTypeNode{Name: "Bloop"},
			&TypeGuard{
				Types: []TypeNode{
					&ConstantTypeNode{Name: "Bloop"},
					&ConstantTypeNode{Name: "Blarp"},
				},
			},
		},
		ReturnVal: &ConstantTypeNode{Name: "Flop"},
	}
	fn2 := &FnType{
		Parameters: []TypeNode{
			&ConstantTypeNode{Name: "Bloop"},
		},
		ReturnVal: &ConstantTypeNode{Name: "Flop"},
	}

	assert.False(fn1.IsSatisfiedBy(fn2))

	fn2 = &FnType{
		Parameters: []TypeNode{
			&ConstantTypeNode{Name: "Bloop"},
			&ConstantTypeNode{Name: "Blarps"},
		},
		ReturnVal: &ConstantTypeNode{Name: "Flop"},
	}

	assert.False(fn1.IsSatisfiedBy(fn2))
}

func TestRejectsIfReturnDoesntMatch(t *testing.T) {
	assert := assert.New(t)

	fn1 := &FnType{
		Parameters: []TypeNode{
			&ConstantTypeNode{Name: "Bloop"},
			&TypeGuard{
				Types: []TypeNode{
					&ConstantTypeNode{Name: "Bloop"},
					&ConstantTypeNode{Name: "Blarp"},
				},
			},
		},
		ReturnVal: &ConstantTypeNode{Name: "Flop"},
	}

	fn2 := &FnType{
		Parameters: []TypeNode{
			&ConstantTypeNode{Name: "Bloop"},
			&ConstantTypeNode{Name: "Blarp"},
		},
		ReturnVal: &ConstantTypeNode{Name: "FlopNope"},
	}

	assert.False(fn1.IsSatisfiedBy(fn2))
}

func TestAnyotherTypeBreaks(t *testing.T) {

	assert := assert.New(t)

	fn1 := &FnType{
		Parameters: []TypeNode{
			&ConstantTypeNode{Name: "Bloop"},
			&TypeGuard{
				Types: []TypeNode{
					&ConstantTypeNode{Name: "Bloop"},
					&ConstantTypeNode{Name: "Blarp"},
				},
			},
		},
		ReturnVal: &ConstantTypeNode{Name: "Flop"},
	}
	fn2 := &ConstantTypeNode{Name: "Blarp"}

	assert.False(fn1.IsSatisfiedBy(fn2))
}
