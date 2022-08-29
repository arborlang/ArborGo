package umlvisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

func (u *umlVisitor) VisitVarName(node *ast.VarName) (ast.Node, error) {
	label := u.getLabel("varname")
	u.writeLine("object \"VarName %s\" as %s {", node.Name, label)
	if node.Type != nil {
		u.writeLine("Type = %s", node.Type)
	} else {
		u.writeLine("Type = Unknown")
	}
	u.writeLine("}")
	return labeledNode(node, label), nil
}
