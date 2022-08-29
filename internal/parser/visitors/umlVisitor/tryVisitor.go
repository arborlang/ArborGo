package umlvisitor

import (
	"log"

	"github.com/arborlang/ArborGo/internal/parser/ast"
)

func (u *umlVisitor) VisitTryNode(node *ast.TryNode) (ast.Node, error) {
	label := u.getLabel("try")
	u.writeLine("object \"Try Node\" as %s", label)
	maybe, _ := node.Tries.Accept(u.v)
	u.connectNodeWithLabel(label, maybe, "Tries")
	for _, handle := range node.HandleCases {
		maybeHandle, _ := handle.Accept(u.v)
		u.connectNodeWithLabel(label, maybeHandle, "handle")
	}
	return labeledNode(node, label), nil
}

func (u *umlVisitor) VisitHandleCaseNode(node *ast.HandleCaseNode) (ast.Node, error) {
	if node == nil {
		log.Println("handleCase is null")
		return node, nil
	}
	label := u.getLabel("handle")

	u.writeLine("object \"Handle Node\" as %s {", label)
	u.writeLine("%s = %s", node.VariableName, node.Type)
	u.writeLine("}")

	maybeHandle, _ := node.Case.Accept(u.v)
	u.connectNode(label, maybeHandle)
	return labeledNode(node, label), nil
}
