package rules

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func returnRule(p *Parser) (ast.Node, error) {
	if nxt := p.Next(); nxt.Token != tokens.RETURN {
		return nil, fmt.Errorf("unexpected token, expected 'return' got %s", nxt)
	}
	returnNode := &ast.ReturnNode{}
	expr, err := ExpressionRule(p)
	if err != nil {
		return nil, err
	}
	returnNode.Expression = expr
	return returnNode, nil
}
