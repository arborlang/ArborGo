package umlvisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

func (u *umlVisitor) VisitAssignmentNode(node *ast.AssignmentNode) (ast.Node, error) {
	label := u.getLabel("assignment")
	u.writeLine("object \"Assignment\" as %s", label)
	maybeUMLValue, _ := node.Value.Accept(u.v)
	u.connectNode(label, maybeUMLValue)
	maybeUMLAssignTo, _ := node.AssignTo.Accept(u.v)
	u.connectNodeWithLabel(label, maybeUMLAssignTo, "Assign To")
	return labeledNode(node, label), nil
}
