package umlvisitor

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
)

func (u *umlVisitor) VisitFunctionDefinitionNode(node *ast.FunctionDefinitionNode) (ast.Node, error) {
	label := u.getLabel("function")

	u.writeLine("object \"function definition\" as %s {", label)
	if node.Returns == nil {
		u.writeLine("Returns = void")
	} else {
		u.writeLine("Returns = %s", node.Returns)
	}
	u.writeLine("Arguments")
	u.writeLine("}")

	for _, varName := range node.Arguments {
		maybeUML, _ := varName.Accept(u.v)
		u.connectNode(fmt.Sprintf("%s::Arguments", label), maybeUML)
	}

	for _, genericName := range node.GenericTypeNames {
		maybeUML, _ := genericName.Accept(u.v)
		u.connectNodeWithLabel(label, maybeUML, "Generic Params")
	}

	maybeUML, _ := node.Body.Accept(u.v)
	if uml, ok := maybeUML.(*umlNode); ok {
		u.connect(label, uml.label)
	}
	return labeledNode(node, label), nil
}
