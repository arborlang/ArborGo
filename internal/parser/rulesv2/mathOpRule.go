package rulesv2

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

// MathOpRule parses a mathematical operation
func MathOpRule(left ast.Node, p *Parser) (ast.Node, error) {
	opCodeLexeme := p.Next()
	mathNode := &ast.MathOpNode{
		Lexeme: opCodeLexeme,
	}
	mathNode.LeftSide = left
	if opCodeLexeme.Token != tokens.ARTHOP {
		return nil, fmt.Errorf("Unexpected token, expected math symbol, got %s", opCodeLexeme)
	}
	switch opCodeLexeme.Value {
	case "+":
		mathNode.Operation = "add"
	case "-":
		mathNode.Operation = "sub"
	case "*":
		mathNode.Operation = "mul_s"
	case "/":
		mathNode.Operation = "div_s"
	}
	node, err := ExpressionRule(p)
	if err != nil {
		return nil, err
	}
	mathNode.RightSide = node
	return mathNode, nil
}
