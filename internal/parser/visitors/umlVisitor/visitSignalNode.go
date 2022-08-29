package umlvisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

func (u *umlVisitor) VisitSignalNode(node *ast.SignalNode) (ast.Node, error) {
	label := u.getLabel("signal")
	u.writeLine("object \"Signal Node\" as %s {", label)
	u.writeLine("Level = %s", node.Level)
	u.writeLine("}")
	maybeUML, _ := node.ValueToRaise.Accept(u.v)
	u.connectNode(label, maybeUML)
	return labeledNode(node, label), nil
}
