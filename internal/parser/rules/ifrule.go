package rules

import (
	"fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
	"github.com/radding/ArborGo/internal/tokens"
)

func ifNodeRule(p *Parser) (ast.Node, error) {
	ifNode := &ast.IfNode{}
	if next := p.Next(); next.Token != tokens.IF {
		return nil, fmt.Errorf("unexpected token, expected 'if' got %s", next)
	}
	condition, err := ExpressionRule(p)
	if err != nil {
		return nil, err
	}
	ifNode.Condition = condition
	if next := p.Next(); next.Token != tokens.RCURLY {
		return nil, fmt.Errorf("expected '{' got %s", next)
	}
	body, err := ProgramRule(p, tokens.LCURLY)
	if err != nil {
		return nil, err
	}
	ifNode.Body = body
	for next := p.Peek(); next.Token == tokens.ELSE; next = p.Peek() {
		p.Next() // This is an ELSE token, just skip and ignore
		if p.Peek().Token == tokens.IF {
			elif, err := ifNodeRule(p)
			if err != nil {
				return nil, err
			}
			ifNode.ElseIfs = append(ifNode.ElseIfs, elif.(*ast.IfNode))
		} else {
			if next := p.Next(); next.Token != tokens.RCURLY {
				return nil, fmt.Errorf("expected '{' got %s", next)
			}
			elseBlock, err := ProgramRule(p, tokens.LCURLY)
			if err != nil {
				return nil, err
			}
			ifNode.Body = elseBlock
			break
		}
	}

	return ifNode, nil
}
