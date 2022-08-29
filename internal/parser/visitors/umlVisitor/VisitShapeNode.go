package umlvisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

func (u *umlVisitor) VisitShapeNode(node *ast.ShapeNode) (ast.Node, error) {
	label := u.getLabel("shape")
	u.writeLine("object \"Shape Node\" as %s", label)
	return labeledNode(node, label), nil
}
