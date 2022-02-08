package rulesv2

import (
	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func continueRule(p *Parser) (ast.Node, error) {
	nxt := p.Next()
	if nxt.Token != tokens.CONTINUE {
		return nil, UnexpectedError(nxt, "continue")
	}
	contNode := &ast.ContinueNode{}
	contNode.Lexeme = nxt
	nxt = p.Peek()
	if nxt.Token == tokens.WITH {
		p.Next()
		exprNode, err := ExpressionRule(p)
		if err != nil {
			return nil, err
		}
		contNode.WithValue = exprNode
	}
	return contNode, nil
}
