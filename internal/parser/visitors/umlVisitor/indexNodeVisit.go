package umlvisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

func (u *umlVisitor) VisitIndexNode(node *ast.IndexNode) (ast.Node, error) {
	label := u.getLabel("index")
	u.writeLine("object \"Index\" as %s", label)
	maybeValue, _ := node.Varname.Accept(u.v)
	u.connectNode(label, maybeValue)
	maybeIndexVal, _ := node.Index.Accept(u.v)
	u.connectNode(label, maybeIndexVal)
	return labeledNode(node, label), nil
}
