package umlvisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

func (u *umlVisitor) VisitDotNode(node *ast.DotNode) (ast.Node, error) {
	label := u.getLabel("dot")
	u.writeLine("object \"Dot accessor\" as %s", label)
	maybeVarName, _ := node.VarName.Accept(u.v)
	u.connectNodeWithLabel(label, maybeVarName, "access")
	maybeAccessor, _ := node.Access.Accept(u.v)
	u.connectNodeWithLabel(label, maybeAccessor, "get Value")

	return labeledNode(node, label), nil
}
