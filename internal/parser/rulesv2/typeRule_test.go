package rulesv2

import (
	"bytes"
	"testing"

	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
	"github.com/stretchr/testify/assert"
)

func parseTest(msg string) *Parser {
	tokStream := lexer.Lex(bytes.NewReader([]byte(msg)))
	return New(tokStream)
}

func TestCanParseTypes(t *testing.T) {
	assert := assert.New(t)

	msg := `Integer`
	parser := parseTest(msg)

	tp, err := typeRule(parser)
	assert.NoError(err)
	realTP, ok := tp.(*types.TypeGuard)
	assert.True(ok)
	assert.Len(realTP.Types, 1)

	assert.True(tp.IsSatisfiedBy(&types.ConstantTypeNode{Name: "Integer"}))
	assert.False(tp.IsSatisfiedBy(&types.ConstantTypeNode{Name: "String"}))

	msg = `String`
	parser = parseTest(msg)

	tp, err = typeRule(parser)
	assert.NoError(err)
	realTP, ok = tp.(*types.TypeGuard)
	assert.True(ok)
	assert.Len(realTP.Types, 1)

	assert.False(tp.IsSatisfiedBy(&types.ConstantTypeNode{Name: "Integer"}))
	assert.True(tp.IsSatisfiedBy(&types.ConstantTypeNode{Name: "String"}))
}

func TestCanParseGuardTypes(t *testing.T) {
	assert := assert.New(t)

	msg := `String | Integer`
	parser := parseTest(msg)

	tp, err := typeRule(parser)
	assert.NoError(err)
	realTP, ok := tp.(*types.TypeGuard)
	assert.True(ok)
	assert.Len(realTP.Types, 2)

	assert.True(tp.IsSatisfiedBy(&types.ConstantTypeNode{Name: "Integer"}))
	assert.True(tp.IsSatisfiedBy(&types.ConstantTypeNode{Name: "String"}))
	assert.False(tp.IsSatisfiedBy(&types.ConstantTypeNode{Name: "Foo"}))

	msg = `String | Integer | Boolean`
	parser = parseTest(msg)

	tp, err = typeRule(parser)
	assert.NoError(err)
	realTP, ok = tp.(*types.TypeGuard)
	assert.True(ok)
	assert.Len(realTP.Types, 3)

	assert.True(tp.IsSatisfiedBy(&types.ConstantTypeNode{Name: "Integer"}))
	assert.True(tp.IsSatisfiedBy(&types.ConstantTypeNode{Name: "String"}))
	assert.True(tp.IsSatisfiedBy(&types.ConstantTypeNode{Name: "Boolean"}))
}

func TestCanParseFunctionType(t *testing.T) {
	assert := assert.New(t)
	fn := `fn(obj: Integer, obj2: Integer | String) -> Integer`
	parser := parseTest(fn)
	tp, err := typeRule(parser)
	if !assert.NoError(err) {
		return
	}
	realTp, ok := tp.(*types.TypeGuard)
	assert.True(ok)
	assert.Len(realTp.Types, 1)
	params := []types.TypeNode{
		&types.TypeGuard{
			Types: []types.TypeNode{
				&types.ConstantTypeNode{Name: "Integer"},
			},
		},
		&types.TypeGuard{
			Types: []types.TypeNode{
				&types.ConstantTypeNode{Name: "Integer"},
				&types.ConstantTypeNode{Name: "String"},
			},
		},
	}
	fnType := &types.FnType{
		Parameters: params,
		ReturnVal:  &types.ConstantTypeNode{Name: "Integer"},
	}
	assert.True(realTp.IsSatisfiedBy(fnType))
}

func TestCanParseFuncWithNoParams(t *testing.T) {
	assert := assert.New(t)
	fn := `fn() -> Integer`
	parser := parseTest(fn)
	tp, err := typeRule(parser)
	if !assert.NoError(err) {
		return
	}
	_, ok := tp.(*types.TypeGuard)
	assert.True(ok)

}

func TestCanParseFunctionWithNoReturnType(t *testing.T) {
	assert := assert.New(t)
	fn := `fn(obj: Integer, obj2: Integer | String)`
	parser := parseTest(fn)
	tp, err := typeRule(parser)
	if !assert.NoError(err) {
		return
	}
	realTp, ok := tp.(*types.TypeGuard)
	assert.True(ok)
	assert.Len(realTp.Types, 1)
	params := []types.TypeNode{
		&types.TypeGuard{
			Types: []types.TypeNode{
				&types.ConstantTypeNode{Name: "Integer"},
			},
		},
		&types.TypeGuard{
			Types: []types.TypeNode{
				&types.ConstantTypeNode{Name: "Integer"},
				&types.ConstantTypeNode{Name: "String"},
			},
		},
	}
	fnType := &types.FnType{
		Parameters: params,
		ReturnVal:  &types.ConstantTypeNode{Name: "Integer"},
	}
	assert.True(realTp.IsSatisfiedBy(fnType))
}

func TestCanParseShapeTypes(t *testing.T) {
	assert := assert.New(t)
	shape := `{
		field: String | Integer;
		field2: Integer | fn () -> Integer;	
	}`
	p := parseTest(shape)
	tp, err := typeRule(p)
	if !assert.NoError(err) {
		return
	}

	realTp, ok := tp.(*types.TypeGuard)
	assert.True(ok)
	assert.Len(realTp.Types, 1)
	fields := make(map[string]types.TypeNode)

	fields["field"] = &types.ConstantTypeNode{Name: "String"}
	fields["field2"] = &types.ConstantTypeNode{Name: "Integer"}
	shapeTp := &types.ShapeType{
		Fields: fields,
	}
	assert.True(realTp.IsSatisfiedBy(shapeTp))
}

func TestCanParseArrayType(t *testing.T) {
	assert := assert.New(t)

	msg := `Integer[]`
	parser := parseTest(msg)

	tp, err := typeRule(parser)
	assert.NoError(err)
	realTP, ok := tp.(*types.TypeGuard)
	assert.True(ok)
	assert.Len(realTP.Types, 1)

	assert.True(tp.IsSatisfiedBy(&types.ArrayType{SubType: &types.ConstantTypeNode{Name: "Integer"}}))
	assert.False(tp.IsSatisfiedBy(&types.ArrayType{SubType: &types.ConstantTypeNode{Name: "String"}}))

	msg = `String[]`
	parser = parseTest(msg)

	tp, err = typeRule(parser)
	assert.NoError(err)
	realTP, ok = tp.(*types.TypeGuard)
	assert.True(ok)
	assert.Len(realTP.Types, 1)

	assert.False(tp.IsSatisfiedBy(&types.ArrayType{SubType: &types.ConstantTypeNode{Name: "Integer"}}))
	assert.True(tp.IsSatisfiedBy(&types.ArrayType{SubType: &types.ConstantTypeNode{Name: "String"}}))

}
