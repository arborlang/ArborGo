package umlvisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

func (u *umlVisitor) VisitReturnNode(returnNode *ast.ReturnNode) (ast.Node, error) {
	label := u.getLabel("return")
	u.writeLine("object \"Return Node\" as %s", label)
	maybeNode, _ := returnNode.Expression.Accept(u.v)
	u.connectNode(label, maybeNode)
	return labeledNode(returnNode, label), nil
}
