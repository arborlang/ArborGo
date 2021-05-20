package rulesv2

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func decoratorRule(p *Parser) (ast.Node, error) {
	at := p.Next()
	if at.Token != tokens.AT {
		return nil, fmt.Errorf("expected %q, got %s instead", "@", at)
	}
	decoratorName := p.Next()
	if decoratorName.Token != tokens.VARNAME {
		return nil, fmt.Errorf("expected a variable name, got %s", decoratorName)
	}
	decoratedNode := &ast.DecoratorNode{
		Name: &ast.VarName{
			Name:   decoratorName.Value,
			Lexeme: decoratorName,
		},
	}
	nxt := p.Peek()
	if nxt.Token != tokens.FUNC && nxt.Token != tokens.TYPE {
		return nil, fmt.Errorf("expected %q or %q, got %s instead", "fn", "type", nxt)
	}
	node, err := ExpressionRule(p)
	if err != nil {
		return nil, err
	}
	decoratedNode.Decorates = node
	return decoratedNode, nil
}
