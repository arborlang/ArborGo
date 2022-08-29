package umlvisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

func (u *umlVisitor) VisitInternalNode(node *ast.InternalNode) (ast.Node, error) {
	label := u.getLabel("internal")
	u.writeLine("object \"Internal Node\" as %s", label)
	uml, _ := node.Expression.Accept(u.v)
	u.connectNode(label, uml)
	return labeledNode(node, label), nil
}
