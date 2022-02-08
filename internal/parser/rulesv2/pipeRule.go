package rulesv2

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func pipeRule(left ast.Node, p *Parser) (ast.Node, error) {
	next := p.Next()
	pipe := &ast.PipeNode{
		Lexeme: next,
	}
	pipe.LeftSide = left
	if next.Token != tokens.PIPE {
		return nil, fmt.Errorf("unexpected token %s, expected '|>'", next)
	}
	nextNode, err := varNameRule(p, false)
	if err != nil {
		return nil, err
	}
	pipe.RightSide = nextNode
	return pipe, nil
}
