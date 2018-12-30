package rules

import (
	"fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
	"github.com/radding/ArborGo/internal/tokens"
)

func comparisonRule(leftSide ast.Node, p *Parser) (ast.Node, error) {
	comp := &ast.Comparison{}
	comp.LeftSide = leftSide
	next := p.Next()
	if next.Token != tokens.COMPARISON {
		return nil, fmt.Errorf("expected a comparison operator, got %s instead", next)
	}
	switch next.Value {
	case "<=":
		comp.Operation = "lte"
	case "<":
		comp.Operation = "lt"
	case ">":
		comp.Operation = "gt"
	case ">=":
		comp.Operation = "gte"
	case "==":
		comp.Operation = "eq"
	}
	rightSide, err := ExpressionRule(p)
	if err != nil {
		return nil, err
	}
	comp.RightSide = rightSide
	return comp, nil
}
