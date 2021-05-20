package rulesv2

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func importMultiple(importAs *ast.ImportNode, p *Parser) error {
	newImportModule := &ast.ImportNode{}
	newImportModule.Source = importAs.Source
	importAs.NextImport = newImportModule
	nxt := p.Next()
	if nxt.Token != tokens.VARNAME {
		return fmt.Errorf("expected a variable name, not %s", nxt)
	}
	newImportModule.ExportName = nxt.Value
	newImportModule.ImportAs = nxt.Value
	nxt = p.Peek()
	if nxt.Token == tokens.AS {
		p.Next()
		nxt = p.Next()
		if nxt.Token != tokens.VARNAME {
			return fmt.Errorf("expected a variable name, not %s", nxt)
		}
		newImportModule.ImportAs = nxt.Value
	}
	nxt = p.Peek()
	if nxt.Token == tokens.COMMA {
		p.Next()
		return importMultiple(newImportModule, p)
	}
	if nxt.Token != tokens.LCURLY {
		return fmt.Errorf("expected %q or %q, got %s", "}", ",", nxt)
	}
	p.Next()
	return nil
}

// import stdlib from "github.com/arborlang/stdlib"
func importRule(p *Parser) (ast.Node, error) {
	next := p.Next()
	importNode := &ast.ImportNode{}
	if next.Token != tokens.IMPORT {
		return nil, fmt.Errorf("expected %q, not %s", "import", next)
	}
	next = p.Next()
	importAs, export := "", ""
	if next.Token == tokens.VARNAME {
		importAs = next.Value
		export = next.Value
		importNode.ImportAs = importAs
		importNode.ExportName = export
		nxt := p.Peek()
		if nxt.Token == tokens.AS {
			p.Next()
			nxt = p.Next()
			if nxt.Token != tokens.VARNAME {
				return nil, fmt.Errorf("expected a variable name, got %s", nxt)
			}
			importNode.ImportAs = nxt.Value
		}
	} else if next.Token == tokens.RCURLY {
		if err := importMultiple(importNode, p); err != nil {
			return nil, err
		}
	}
	nxt := p.Next()
	if nxt.Token != tokens.FROM {
		return nil, fmt.Errorf("expected %q, not %s", "from", nxt)
	}
	source := p.Next()
	if source.Token != tokens.STRINGVAL {
		return nil, fmt.Errorf("expected an import string, got %s", source)
	}
	importNode.Source = source.Value
	return importNode, nil
}
