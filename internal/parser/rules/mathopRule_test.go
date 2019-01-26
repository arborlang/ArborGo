package rules

import (
	"bytes"
	"testing"

	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/stretchr/testify/assert"
)

var simpleMath = `+ 1;` // don't put the numbers in the fron because that should be caught by the expressionRule func
var simpleMath2 = `+ 2 - 3 * 4 / 5;`

var mathIsExpression = `1 + 2`

func TestMathRuleWorks(t *testing.T) {
	assert := assert.New(t)
	tokStream := lexer.Lex(bytes.NewReader([]byte(simpleMath)))
	parser := New(tokStream)

	constant := &ast.Constant{}
	constant.Value = "2"
	constant.Type = "NUMBER"

	math, err := MathOpRule(constant, parser)
	assert.NoError(err)
	assert.NotNil(math)
	mathNode, ok := math.(*ast.MathOpNode)
	if !ok {
		assert.FailNow("can not convert node to MathNode")
	}
	assert.NotNil(mathNode)
	assert.Equal("add", mathNode.Operation)

	// Assert that both sides are constants
	leftside, ok := mathNode.LeftSide.(*ast.Constant)
	if !ok {
		assert.FailNow("cannot convert left side to a constant")
	}
	assert.Equal("2", leftside.Value)
	assert.Equal("NUMBER", leftside.Type)

	rightSide, ok := mathNode.RightSide.(*ast.Constant)
	if !ok {
		assert.FailNow("cannot convert right side to a constant")
	}
	assert.Equal("1", rightSide.Value)
	assert.Equal("NUMBER", rightSide.Type)
}

func TestMathWorksWithMoreComplexAdditions(t *testing.T) {
	assert := assert.New(t)
	tokStream := lexer.Lex(bytes.NewReader([]byte(simpleMath2)))
	parser := New(tokStream)

	constant := &ast.Constant{}
	constant.Value = "2"
	constant.Type = "NUMBER"
	math, err := MathOpRule(constant, parser)
	assert.NoError(err)
	assert.NotNil(math)
	mathNode, ok := math.(*ast.MathOpNode)
	if !ok {
		assert.FailNow("can not convert node to MathNode")
	}

	//Manually traverse the ast now.
	// should be 2 + <AST 2 - <AST 3 * <AST 4 / 5>>>
	leftSide, ok := mathNode.LeftSide.(*ast.Constant)
	if !ok {
		assert.FailNow("could not convert left side to a constant")
	}
	assert.Equal("2", leftSide.Value)
	rightSide, ok := mathNode.RightSide.(*ast.MathOpNode)
	if !ok {
		assert.FailNow("could not convert right side to a constant")
	}
	assert.NotNil(rightSide)
	assert.Equal("add", mathNode.Operation)

	//should be 2 - <AST 3 * <AST 4 / 5>>
	assert.Equal("sub", rightSide.Operation)
	leftSide, ok = rightSide.LeftSide.(*ast.Constant)
	if !ok {
		assert.FailNow("could not convert left side to a constant")
	}
	assert.Equal("2", leftSide.Value)
	rightSide, ok = rightSide.RightSide.(*ast.MathOpNode)
	if !ok {
		assert.FailNow("could not convert right side to a constant")
	}
	assert.NotNil(rightSide)

	//should be 3 * <AST 4 / 5>
	assert.Equal("mul_s", rightSide.Operation)
	leftSide, ok = rightSide.LeftSide.(*ast.Constant)
	if !ok {
		assert.FailNow("could not convert left side to a constant")
	}
	assert.Equal("3", leftSide.Value)
	rightSide, ok = rightSide.RightSide.(*ast.MathOpNode)
	if !ok {
		assert.FailNow("could not convert right side to a constant")
	}
	assert.NotNil(rightSide)

	//should 4 / 5
	assert.Equal("div_s", rightSide.Operation)
	leftSide, ok = rightSide.LeftSide.(*ast.Constant)
	if !ok {
		assert.FailNow("could not convert left side to a constant")
	}
	assert.Equal("4", leftSide.Value)
	rightSide2, ok := rightSide.RightSide.(*ast.MathOpNode)
	if ok {
		assert.FailNow("converted to a math node?")
	}
	assert.Nil(rightSide2)

	rightSideConst, ok := rightSide.RightSide.(*ast.Constant)
	if !ok {
		assert.FailNow("Failed to convert to a constant")
	}
	assert.Equal("5", rightSideConst.Value)
	assert.NotNil(rightSideConst)
}

func TestIsExpression(t *testing.T) {
	assert := assert.New(t)
	tokStream := lexer.Lex(bytes.NewReader([]byte(mathIsExpression)))
	parser := New(tokStream)

	math, err := ExpressionRule(parser)
	if !assert.NoError(err) {
		t.Fatal()
	}
	assert.NotNil(math)
	mathNode := math.(*ast.MathOpNode)
	assert.NotNil(mathNode, "Could not get expressionRule back from parser")
	leftSide := mathNode.LeftSide.(*ast.Constant)
	rightSide := mathNode.RightSide.(*ast.Constant)
	assert.NotNil(leftSide, "LeftSide could not be converted to a constant")
	assert.NotNil(rightSide, "RightSide could not be converted")
	assert.Equal("add", mathNode.Operation)
	assert.Equal("1", leftSide.Value)
	assert.Equal("2", rightSide.Value)
}
