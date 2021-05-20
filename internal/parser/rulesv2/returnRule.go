package rulesv2

import (
	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func returnRule(p *Parser) (ast.Node, error) {
	nxt := p.Next()
	if nxt.Token != tokens.RETURN {
		return nil, UnexpectedError(nxt, "return")
	}
	returnNode := &ast.ReturnNode{
		Lexeme: nxt,
	}
	peek := p.Peek()
	if peek.Token == tokens.SEMI {
		return returnNode, nil
	}
	expr, err := ExpressionRule(p)
	if err != nil {
		return nil, err
	}
	returnNode.Expression = expr
	return returnNode, nil
}
