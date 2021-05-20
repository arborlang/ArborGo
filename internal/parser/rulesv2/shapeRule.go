package rulesv2

import (
	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func constantShapeRule(p *Parser) (ast.Node, error) {
	shapeNode := &ast.ShapeNode{
		Fields: map[string]ast.Node{},
	}
	next := p.Next()
	if next.Token != tokens.RCURLY {
		return nil, UnexpectedError(next, "{")
	}
	for {
		next = p.Next()
		if next.Token != tokens.VARNAME && next.Token != tokens.LCURLY {
			return nil, UnexpectedError(next, "variable name", "}")
		}
		shouldBeColon := p.Next()
		if shouldBeColon.Token != tokens.COLON {
			return nil, UnexpectedError(shouldBeColon, ":")
		}
		exprNode, err := exprRule(p, true)
		if err != nil {
			return nil, err
		}
		comma := p.Next()
		if comma.Token != tokens.COMMA && comma.Token != tokens.LCURLY {
			return nil, UnexpectedError(comma, ",", "}")
		}
		if comma.Token == tokens.LCURLY {
			break
		}
		shapeNode.Fields[next.Value] = exprNode
	}

	return shapeNode, nil
}
