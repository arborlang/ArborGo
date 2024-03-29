package rulesv2

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func comparisonRule(leftSide ast.Node, p *Parser) (ast.Node, error) {
	comp := &ast.Comparison{}
	comp.LeftSide = leftSide
	next := p.Next()
	comp.Lexeme = next
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
	case "!=":
		comp.Operation = "neq"
	}
	rightSide, err := ExpressionRule(p)
	if err != nil {
		return nil, err
	}
	comp.RightSide = rightSide
	return comp, nil
}
