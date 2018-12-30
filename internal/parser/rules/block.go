package rules

import (
	"fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
	"github.com/radding/ArborGo/internal/tokens"
)

func parseBlockRule(p *Parser) (ast.Node, error) {
	next := p.Next()
	if next.Token != tokens.RCURLY {
		return nil, fmt.Errorf("expected '{', got %s", next)
	}
	return ProgramRule(p, tokens.LCURLY)
}
