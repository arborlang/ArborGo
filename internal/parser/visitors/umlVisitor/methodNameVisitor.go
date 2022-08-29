package umlvisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

func (u *umlVisitor) VisitMethodDefinition(node *ast.MethodDefinition) (ast.Node, error) {
	label := u.getLabel("methodDef")
	u.writeLine("object \"Define %s on %s\" as %s", node.MethodName.Name, node.TypeName.Name, label)
	maybeUML, _ := node.FuncDef.Accept(u.v)
	if uml, ok := maybeUML.(*umlNode); ok {
		u.connect(label, uml.label)
	}
	return labeledNode(node, label), nil
}
