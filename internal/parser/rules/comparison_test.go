package rules

import (
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/stretchr/testify/assert"
)

// var lte := "<= 12"
// var lt := "< 12"
// var gt := "> 12"
// var gte := ">= 12"
// var eq := "== 12"

var simpleTests = []struct {
	name      string
	testCase  string
	operation string
}{
	{"Less than or Equal", "<= 12", "lte"},
	{"Less Than", "< 12", "lt"},
	{"Greater than", "> 12", "gt"},
	{"Greater than or equal", ">= 12", "gte"},
	{"Equal", "== 12", "eq"},
}

var complexTests = []struct {
	name      string
	testCase  string
	operation string
}{
	{"Complex Less than or Equal with a varname", "abc <= 12", "lte"},
	{"Complex Less Than with a varname", "abc < 12", "lt"},
	{"Complex Greater than with a varname", "abc > 12", "gt"},
	{"Complex Greater than or equal with a varname", "abc >= 12", "gte"},
	{"complex Equal with a varname", "abc == 12", "eq"},

	{"Complex Less than or Equal with a function", "abc() <= 12", "lte"},
	{"Complex Less Than with a function", "abc() < 12", "lt"},
	{"Complex Greater than with a function", "abc() > 12", "gt"},
	{"Complex Greater than or equal with a function", "abc() >= 12", "gte"},
	{"complex Equal with a function", "abc() == 12", "eq"},

	{"Complex Less than or Equal with a constant", "112 <= 12", "lte"},
	{"Complex Less Than with a constant", "112 < 12", "lt"},
	{"Complex Greater than with a constant", "112 > 12", "gt"},
	{"Complex Greater than or equal with a constant", "112 >= 12", "gte"},
	{"complex Equal with a constant", "112 == 12", "eq"},
}

func TestCanParseAComparssion(t *testing.T) {
	assert := assert.New(t)
	for _, testCase := range simpleTests {
		t.Run(testCase.name, func(t *testing.T) {

			p := createParser(testCase.testCase)
			varName := &ast.VarName{}
			comp, err := comparisonRule(varName, p)
			if !assert.NoError(err) {
				t.Fatal()
			}
			assert.NotNil(comp)
			compNode, _ := comp.(*ast.Comparison)
			if !assert.NotNil(compNode) {
				t.Fatal()
			}
			assert.Equal(testCase.operation, compNode.Operation)
			rightSide, _ := compNode.RightSide.(*ast.Constant)
			if !assert.NotNil(rightSide) {
				t.Fatal()
			}
			assert.Equal("12", rightSide.Value)
		})
	}
}

func TestCanBeAValidExpression(t *testing.T) {
	assert := assert.New(t)
	for _, testCase := range complexTests {
		t.Run(testCase.name, func(t *testing.T) {

			p := createParser(testCase.testCase)
			comp, err := ExpressionRule(p)
			if !assert.NoError(err) {
				t.Fatal()
			}
			assert.NotNil(comp)
			compNode, _ := comp.(*ast.Comparison)
			if !assert.NotNil(compNode) {
				t.Fatal()
			}
			assert.Equal(testCase.operation, compNode.Operation)
			rightSide, _ := compNode.RightSide.(*ast.Constant)
			if !assert.NotNil(rightSide) {
				t.Fatal()
			}
			assert.Equal("12", rightSide.Value)

		})
	}
}
