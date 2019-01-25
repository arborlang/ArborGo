package rules

import (
	"fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
	"github.com/radding/ArborGo/internal/tokens"
	"strconv"
)

func sliceRule(ndxNode *ast.IndexNode, p *Parser) (ast.Node, error) {
	slice := &ast.SliceNode{}
	slice.Varname = ndxNode.Varname
	slice.Start = ndxNode.Index
	number := p.Next()
	if number.Token == tokens.RSQUARE {
		slice.End = -1
		return slice, nil
	}
	if number.Token != tokens.NUMBER {
		return nil, fmt.Errorf("expected a number, got %s instead", number)
	}
	end, err := strconv.Atoi(number.Value)
	if err != nil {
		return nil, err
	}
	slice.End = end
	return slice, nil
}
