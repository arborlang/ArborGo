package rules

import (
	"fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
	"github.com/radding/ArborGo/internal/tokens"
)

// MathOpRule parses a mathematical operation
func MathOpRule(left ast.Node, p *Parser) (ast.Node, error) {
	mathNode := &ast.MathOpNode{}
	mathNode.LeftSide = left
	opCodeLexeme := p.Next()
	if opCodeLexeme.Token != tokens.ARTHOP {
		return nil, fmt.Errorf("Unexpected token, expected math symbol, got %s", opCodeLexeme)
	}
	switch opCodeLexeme.Value {
	case "+":
		mathNode.Operation = "add"
	case "-":
		mathNode.Operation = "sub"
	case "*":
		mathNode.Operation = "mult"
	case "/":
		mathNode.Operation = "div"
	}
	node, err := ExpressionRule(p)
	if err != nil {
		return nil, err
	}
	mathNode.RightSide = node
	return mathNode, nil
}
