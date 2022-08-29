package umlvisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

func (u *umlVisitor) VisitPipeNode(node *ast.PipeNode) (ast.Node, error) {
	label := u.getLabel("pipeline")
	u.writeLine("object \"Pipeline Node\" as %s", label)
	uml, _ := node.LeftSide.Accept(u.v)
	u.connectNode(label, uml)
	uml, _ = node.RightSide.Accept(u.v)
	u.connectNode(label, uml)
	return labeledNode(node, label), nil
}
