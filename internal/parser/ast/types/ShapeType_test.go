package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShapeTypeIsTypeNode(t *testing.T) {
	assert := assert.New(t)

	assert.Implements((*TypeNode)(nil), new(ShapeType))
}

func TestShapeNodeIsSatisfiedByExactShape(t *testing.T) {
	assert := assert.New(t)

	fields := make(map[string]TypeNode)
	fields["foo"] = &ConstantTypeNode{Name: "Fooer"}
	shape := &ShapeType{
		Fields: fields,
	}
	shape2 := &ShapeType{
		Fields: fields,
	}
	assert.True(shape.IsSatisfiedBy(shape2))
	assert.True(shape2.IsSatisfiedBy(shape))
}

func TestShapeNodeRejectsShapeWithSameFieldNameButDifferentTypes(t *testing.T) {
	assert := assert.New(t)

	fields := make(map[string]TypeNode)
	fields["foo"] = &ConstantTypeNode{Name: "Fooer"}
	shape := &ShapeType{
		Fields: fields,
	}
	fields2 := make(map[string]TypeNode)
	fields2["foo"] = &ConstantTypeNode{Name: "Fooer2"}
	shape2 := &ShapeType{
		Fields: fields2,
	}
	assert.False(shape.IsSatisfiedBy(shape2))
	assert.False(shape2.IsSatisfiedBy(shape))
}

func TestShapeCanBeSatisifiedByDifferentShapes(t *testing.T) {
	assert := assert.New(t)

	fields := make(map[string]TypeNode)
	fields["foo"] = &ConstantTypeNode{Name: "Fooer"}
	shape := &ShapeType{
		Fields: fields,
	}
	fields2 := make(map[string]TypeNode)
	fields2["foo"] = &ConstantTypeNode{Name: "Fooer"}
	fields2["barer"] = &ConstantTypeNode{Name: "Barer"}
	shape2 := &ShapeType{
		Fields: fields2,
	}
	assert.True(shape.IsSatisfiedBy(shape2))
	assert.False(shape2.IsSatisfiedBy(shape))
}

func TestWillRejectAnythingOtherThanAnotherShape(t *testing.T) {
	assert := assert.New(t)

	fields := make(map[string]TypeNode)
	fields["foo"] = &ConstantTypeNode{Name: "Fooer"}
	shape := &ShapeType{
		Fields: fields,
	}
	shape2 := &ConstantTypeNode{Name: "Fooer"}
	assert.False(shape.IsSatisfiedBy(shape2))
}
