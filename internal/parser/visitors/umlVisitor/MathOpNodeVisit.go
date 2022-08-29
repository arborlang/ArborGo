package umlvisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

func (u *umlVisitor) VisitMathOpNode(node *ast.MathOpNode) (ast.Node, error) {
	label := u.getLabel("mathop")
	u.writeLine("object \"Math operation\" as %s {", label)
	u.writeLine("Operation = %s", node.Operation)
	u.writeLine("}")
	maybeUml, _ := node.LeftSide.Accept(u.v)
	u.connectNode(label, maybeUml)
	maybeUml, _ = node.RightSide.Accept(u.v)
	u.connectNode(label, maybeUml)
	return labeledNode(node, label), nil
}
