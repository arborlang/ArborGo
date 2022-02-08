package rulesv2

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func sliceRule(ndxNode *ast.IndexNode, p *Parser) (ast.Node, error) {
	p.setIsInSlice(true)
	defer p.setIsInSlice(false)
	col := p.Next()
	if col.Token != tokens.COLON {
		return nil, fmt.Errorf("expected %q, not %s", ":", col)
	}
	slice := &ast.SliceNode{
		Lexeme: col,
	}
	slice.Varname = ndxNode.Varname
	slice.Start = ndxNode.Index
	number := p.Peek()
	if number.Token == tokens.RSQUARE {
		p.Next()
		slice.End = nil
		return slice, nil
	}
	end, err := exprRule(p, true)
	if err != nil {
		return nil, err
	}
	slice.End = end
	nxt := p.Next()
	if nxt.Token == tokens.RSQUARE {
		return slice, nil
	}
	if nxt.Token != tokens.COLON {
		return nil, fmt.Errorf("expected %q or %q, not %s", ":", "]", nxt)
	}
	step, err := exprRule(p, true)
	if err != nil {
		return nil, err
	}
	slice.Step = step
	next := p.Next()
	if next.Token != tokens.RSQUARE {
		return nil, fmt.Errorf("expected %q, not %s", "]", next)
	}
	return slice, nil
}
