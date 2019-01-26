package rules

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func boolOperation(leftSide ast.Node, p *Parser) (ast.Node, error) {
	boolNode := &ast.BoolOp{}
	boolNode.LeftSide = leftSide
	booleanOperator := p.Next()
	switch booleanOperator.Token {
	case tokens.BOOLEAN:
		boolNode.Condition = "and"
		if booleanOperator.Value == "||" {
			boolNode.Condition = "or"
		}
	default:
		return nil, fmt.Errorf("expected '&&', '||', '<', '>', '==', '<=', or '=>', go %s instead", booleanOperator)
	}
	next, err := ExpressionRule(p)
	if err != nil {
		return nil, err
	}
	boolNode.RightSide = next
	return boolNode, nil
}
