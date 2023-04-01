package typevisitor

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
)

func (t *typeVisitor) VisitAssignmentNode(n *ast.AssignmentNode) (ast.Node, error) {
	var vname *ast.VarName
	value, err := n.Value.Accept(t.v)
	if err != nil {
		return nil, err
	}
	n.Value = value
	switch tp := n.AssignTo.(type) {
	case *ast.DeclNode:
		vname = tp.Varname
		if vname.Type == nil {
			tp.Varname.Type = n.Value.GetType()
		}
		node, err := tp.Accept(t.v)
		if err != nil {
			return nil, err
		}
		n.AssignTo = node
	case *ast.VarName:
		vname = tp
		elem, _ := t.scope.LookupSymbol(vname.Name)
		if elem != nil && elem.IsConstant {
			return nil, fmt.Errorf("%s is constant: defined here: %s", vname.Name, elem.Lexeme)
		}
	default:
		return nil, fmt.Errorf("unexpected node %s", tp)
	}
	vnameNode, err := vname.Accept(t.v)
	if err != nil {
		return nil, err
	}
	vname = vnameNode.(*ast.VarName)

	// if !vname.GetType().IsSatisfiedBy(n.Value.GetType()) {
	// 	return nil, fmt.Errorf("can't assign %s to %s at %s", n.Value.GetType(), vname.Type, n.Lexeme)
	// }
	return n, nil
}
