package umlvisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

func (u *umlVisitor) VisitInstantiateNode(node *ast.InstantiateNode) (ast.Node, error) {
	label := u.getLabel("instantiate")

	u.writeLine("object \"instantiate node\" as %s", label)
	maybeFn, _ := node.FunctionCallNode.Accept(u.v)
	u.connectNode(label, maybeFn)

	return labeledNode(node, label), nil
}
