package rulesv2

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func indexRule(name ast.Node, p *Parser) (ast.Node, error) {
	nxt := p.Next()
	if nxt.Token != tokens.LSQUARE {
		return nil, UnexpectedError(nxt, "[")
	}
	index, err := ExpressionRule(p)
	// index := strconv.ItoA
	// index, err := strconv.Atoi(nxt.Value)
	if err != nil {
		return nil, err
	}
	ndxNode := &ast.IndexNode{
		Varname: name,
		Index:   index,
	}
	nxt = p.Peek()
	if nxt.Token == tokens.COLON {
		return sliceRule(ndxNode, p)
	}
	if nxt.Token != tokens.RSQUARE {
		return nil, fmt.Errorf("expected ']', got %q instead", nxt)
	}
	p.Next()
	return ndxNode, nil
}
