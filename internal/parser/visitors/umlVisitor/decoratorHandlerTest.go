package umlvisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

func (u *umlVisitor) VisitDecoratorNode(node *ast.DecoratorNode) (ast.Node, error) {
	label := u.getLabel("decorator")

	u.writeLine("object \"DecoratorNode\" as %s", label)
	uml, _ := node.Decorates.Accept(u.v)
	u.connectNode(label, uml)

	return labeledNode(node, label), nil
}
