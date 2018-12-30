package rules

import (
	"fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
	"github.com/radding/ArborGo/internal/tokens"
)

func pipeRule(left ast.Node, p *Parser) (ast.Node, error) {
	pipe := &ast.PipeNode{}
	pipe.LeftSide = left
	if next := p.Next(); next.Token != tokens.PIPE {
		return nil, fmt.Errorf("unexpected token %s, expected '|>'", next)
	}
	next, err := varNameRule(false, p)
	if err != nil {
		return nil, err
	}
	pipe.RightSide = next
	return pipe, nil
}
