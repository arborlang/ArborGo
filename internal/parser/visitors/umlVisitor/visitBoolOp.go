package umlvisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

func (u *umlVisitor) VisitBoolOp(node *ast.BoolOp) (ast.Node, error) {
	label := u.getLabel("boolop")
	u.writeLine("object \"Boolean Operation\" as %s {", label)
	u.writeLine("Operation = %s", node.Condition)
	u.writeLine("}")

	maybeLeftSide, _ := node.LeftSide.Accept(u.v)
	maybeRightSide, _ := node.RightSide.Accept(u.v)
	u.connectNode(label, maybeLeftSide)
	u.connectNode(label, maybeRightSide)
	return labeledNode(node, label), nil
}
