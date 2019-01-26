package rules

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func importRules(p *Parser) (ast.Node, error) {
	importNode := &ast.ImportNode{}
	if tok := p.Next(); tok.Token != tokens.IMPORT {
		return nil, fmt.Errorf("expected 'import', got %s", tok)
	}
	next := p.Next()
	if next.Token != tokens.STRINGVAL {
		return nil, fmt.Errorf("Expected a path, got %s instead", next)
	}
	importNode.Source = next.Value
	next = p.Peek()
	if next.Token == tokens.AS {
		p.Next()
		next = p.Next()
		if next.Token != tokens.VARNAME {
			return nil, fmt.Errorf("expected a name, got %s instead", next)
		}
		importNode.Name = next.Value
	}
	return importNode, nil
}
