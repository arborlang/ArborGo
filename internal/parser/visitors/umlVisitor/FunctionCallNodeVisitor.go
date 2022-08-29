package umlvisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

func (u *umlVisitor) VisitFunctionCallNode(node *ast.FunctionCallNode) (ast.Node, error) {
	label := u.getLabel("functionCall")
	u.writeLine("object \"Function Call\" as %s", label)
	maybeDef, _ := node.Definition.Accept(u.v)
	if def, ok := maybeDef.(*umlNode); ok {
		u.writeLine("%s --> %s : Definition", label, def.label)
	}
	for _, argNode := range node.Arguments {
		maybeUML, _ := argNode.Accept(u.v)
		u.connectNodeWithLabel(label, maybeUML, "Argument")
	}
	return labeledNode(node, label), nil
}
