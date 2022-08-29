package umlvisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

func (u *umlVisitor) VisitComparison(node *ast.Comparison) (ast.Node, error) {
	label := u.getLabel("comparison")
	u.writeLine("object \"Comparison Node\" as %s {", label)
	u.writeLine("Operation = %s", node.Operation)
	u.writeLine("}")
	maybeLeftSide, _ := node.LeftSide.Accept(u.v)
	u.connectNode(label, maybeLeftSide)
	maybeRightSide, _ := node.RightSide.Accept(u.v)
	u.connectNode(label, maybeRightSide)
	return labeledNode(node, label), nil
}
