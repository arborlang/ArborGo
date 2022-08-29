package umlvisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

func (u *umlVisitor) VisitContinueNode(node *ast.ContinueNode) (ast.Node, error) {
	label := u.getLabel("continue")
	u.writeLine("object \"Continue\" as %s", label)
	if node.WithValue != nil {
		uml, _ := node.WithValue.Accept(u.v)
		u.connectNode(label, uml)
	}
	return labeledNode(node, label), nil
}
