package rules

import (
	"fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
	"github.com/radding/ArborGo/internal/tokens"
)

func isAType(tok tokens.Token) bool {
	return tok == tokens.VARNAME || tok == tokens.STRING || tok == tokens.CHAR || tok == tokens.NUMBERWORD || tok == tokens.FLOATWORD || tok == tokens.FUNC
}

func isAllowable(tok tokens.Token) bool {
	return tok == tokens.COMMA || tok == tokens.SEMI || tok == tokens.ARROW
}

func typeRules(p *Parser) (ast.Node, error) {
	tp := &ast.TypeNode{}
	nxt := p.Next()
	if !isAType(nxt.Token) {
		return nil, fmt.Errorf("expected a type, got %s instead", nxt)
	}
	tp.Types = append(tp.Types, nxt.Value)
	nxt = p.Peek()
	if nxt.Token == tokens.PIPE {
		for nxt := p.Peek(); nxt.Token == tokens.PIPE; nxt = p.Peek() {
			p.Next()        // Skip over the pipe
			nxt := p.Next() // get the type name
			if !isAType(nxt.Token) {
				return nil, fmt.Errorf("expected a type, got %s instead", nxt)
			}
			tp.Types = append(tp.Types, nxt.Value)
		}
		if next := p.Peek(); !isAllowable(next.Token) {
			return nil, fmt.Errorf("unexpected token while parsing types: %s", next)
		}
	}
	return tp, nil
}
