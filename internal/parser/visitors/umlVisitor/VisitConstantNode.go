package umlvisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

func (u *umlVisitor) VisitConstant(node *ast.Constant) (ast.Node, error) {
	label := u.getLabel("constant")
	u.writeLine("object \"constant\" as %s {", label)
	u.writeLine("Type = %s", node.Type)
	u.writeLine("Value = %s", node.Value)
	u.writeLine("}")
	return labeledNode(node, label), nil
}
