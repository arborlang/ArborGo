package umlvisitor

import (
	"github.com/arborlang/ArborGo/internal/parser/ast"
)

func (u *umlVisitor) VisitIfNode(node *ast.IfNode) (ast.Node, error) {
	label := u.getLabel("if")
	labelCond := u.getLabel("condition")
	u.writeLine("object \"if node\" as %s", label)
	// node.Condition.Accept(u.v)
	u.writeLine("package \"condition\" as %s {", labelCond)
	node.Condition.Accept(u.v)
	u.writeLine("}")
	u.connect(label, labelCond)
	maybeBody, _ := node.Body.Accept(u.v)
	if bodyUML, ok := maybeBody.(*umlNode); ok {
		u.writeLine("%s --> %s : when true", labelCond, bodyUML.label)
	}
	for _, elseIfs := range node.ElseIfs {
		maybeIf, _ := elseIfs.Accept(u.v)
		ifUml := maybeIf.(*umlNode)
		u.writeLine("%s --> %s : otherwise maybe", labelCond, ifUml.label)
	}
	if node.Else != nil {
		maybeElse, _ := node.Else.Accept(u.v)
		if elseUML, ok := maybeElse.(*umlNode); ok {
			u.writeLine("%s --> %s : otherwise", labelCond, elseUML.label)
		}

	}
	return labeledNode(node, label), nil
}
