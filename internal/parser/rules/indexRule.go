package rules

import (
	"fmt"
	"strconv"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func indexRule(name *ast.VarName, p *Parser) (ast.Node, error) {
	nxt := p.Next()
	if nxt.Token != tokens.NUMBER {
		return nil, fmt.Errorf("expected a number, got %s instead", nxt)
	}
	// index := strconv.ItoA
	index, err := strconv.Atoi(nxt.Value)
	if err != nil {
		return nil, err
	}
	ndxNode := &ast.IndexNode{
		Varname: name,
		Index:   index,
	}
	nxt = p.Next()
	if nxt.Token == tokens.COLON {
		return sliceRule(ndxNode, p)
	}
	if nxt.Token != tokens.RSQUARE {
		return nil, fmt.Errorf("expected ']', got %q instead", nxt)
	}
	return ndxNode, nil
}
