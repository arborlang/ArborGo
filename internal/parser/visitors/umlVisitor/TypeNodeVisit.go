package umlvisitor

import (
	"github.com/arborlang/ArborGo/internal/parser/ast"
)

func (u *umlVisitor) VisitTypeNode(node *ast.TypeNode) (ast.Node, error) {
	label := u.getLabel("type")
	u.writeLine("object \"type %s\" as %s {", node.VarName.Name, label)
	defer u.writeLine("}")
	u.writeLine("extends = %t", node.Extends)
	u.writeLine("type = %s", node.Types)
	return labeledNode(node, label), nil
}
