package rules

import (
	"fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
	"github.com/radding/ArborGo/internal/tokens"
)

//DeclRule defines how to parse a rule that begins with decleration (`const` or `let`)
func DeclRule(p *Parser) (ast.Node, error) {
	d := &ast.DeclNode{}
	tp := p.Next()
	name := p.Peek()
	if name.Token != tokens.VARNAME {
		return nil, fmt.Errorf("expected token VARNAME, got %s: %s", name.Token.String(), name.Value)
	}
	nameNode, err := varNameRule(true, p)
	if err != nil {
		return nil, err
	}

	switch node := nameNode.(type) {
	case *ast.AssignmentNode:
		if tp.Token == tokens.CONST {
			d.IsConstant = true
		}
		d.AddChild(node.AssignTo)
		node.AssignTo = d

		return node, nil
	case *ast.VarName:
		if tp.Token == tokens.CONST {
			d.IsConstant = true
		}
		d.AddChild(nameNode)
		return d, nil
	}

	return nil, fmt.Errorf("got bad node back from parser")
}
