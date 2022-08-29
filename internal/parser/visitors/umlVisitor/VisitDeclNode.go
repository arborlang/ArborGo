package umlvisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

func (u *umlVisitor) VisitDeclNode(node *ast.DeclNode) (ast.Node, error) {
	label := u.getLabel("decl")
	u.writeLine("object \"declaration node\" as %s {", label)
	u.writeLine("IsConstant = %t", node.IsConstant)
	u.writeLine("}")
	maybeUML, _ := node.Varname.Accept(u.v)
	if uml, ok := maybeUML.(*umlNode); ok {
		u.connect(label, uml.label)
	}
	return labeledNode(node, label), nil
}
