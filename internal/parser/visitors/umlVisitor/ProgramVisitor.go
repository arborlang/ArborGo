package umlvisitor

import (
	"fmt"
	"io"
	"log"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
	"github.com/arborlang/ArborGo/internal/parser/scope"
	"github.com/arborlang/ArborGo/internal/parser/visitors/base"
)

func labeledNode(node ast.Node, label string) *umlNode {
	return &umlNode{
		realNode: node,
		label:    label,
	}
}

type umlNode struct {
	realNode ast.Node
	label    string
}

func (u *umlNode) Accept(v ast.Visitor) (ast.Node, error) {
	return u.realNode.Accept(v)
}
func (u *umlNode) GetType() types.TypeNode {
	return u.realNode.GetType()
}

type umlVisitor struct {
	writer       io.Writer
	v            *base.VisitorAdapter
	labelCounter map[string]int
}

func (u *umlVisitor) connectNode(label string, maybeUML ast.Node) {
	u.connectNodeWithLabel(label, maybeUML, "")
}

func (u *umlVisitor) createConnectWithLabel(label1, label2, connectionLabel string) {
	if len(connectionLabel) > 0 {
		u.writeLine("%s --> %s : %s", label1, label2, connectionLabel)
	} else {
		u.writeLine("%s --> %s", label1, label2)
	}
}

func (u *umlVisitor) connectNodeWithLabel(label string, maybeUML ast.Node, connectionLabel string) {
	if uml, ok := maybeUML.(*umlNode); ok {
		u.createConnectWithLabel(label, uml.label, connectionLabel)
	}
}

func (u *umlVisitor) getLabel(label string) string {
	number, ok := u.labelCounter[label]
	if !ok {
		number = 0
	}
	u.labelCounter[label] = number + 1
	return fmt.Sprintf("%s%d", label, number)
}

func (u *umlVisitor) connect(label1, label2 string) {
	u.createConnectWithLabel(label1, label2, "")
	// u.writeLine("%s --> %s", label1, label2)
}

func (u *umlVisitor) writeLine(str string, args ...interface{}) {
	u.writer.Write([]byte(fmt.Sprintf(str, args...) + "\n"))
}

func (u *umlVisitor) SetVisitor(v *base.VisitorAdapter) {
	u.v = v
}

func Visualize(node ast.Node, w io.Writer) (ast.Node, error) {
	v := New(w)
	w.Write([]byte("@startuml\n"))
	defer w.Write([]byte("@enduml\n"))
	return node.Accept(v)
}

func (u *umlVisitor) GetSymbolTable() *scope.SymbolTable {
	return nil
}

func New(writer io.Writer) ast.Visitor {
	return base.New(&umlVisitor{
		writer:       writer,
		labelCounter: map[string]int{},
	})
}

func (u *umlVisitor) getUMLObject(label string, node ast.Node) *umlObject {
	obj := newObject(u)
	obj.WithLabel(u.getLabel(label))
	obj.WithRawNode(node)
	return obj
}

func (u *umlVisitor) VisitProgram(prog *ast.Program) (ast.Node, error) {
	importUmlNodes := []*umlNode{}
	label := u.getLabel("program")
	for _, node := range prog.Imports {
		labeledNodeMaybe, _ := node.Accept(u.v)
		labeledNode, ok := labeledNodeMaybe.(*umlNode)
		if !ok {
			// log.Println("Unable to convert to uml node")
			panic("Unable to convert to uml node")
		} else {
			importUmlNodes = append(importUmlNodes, labeledNode)
		}
	}
	u.writeLine("object \"block\" as %s", label)
	for _, visNode := range importUmlNodes {
		u.writeLine("%s o-- %s", visNode.label, label)
	}
	for _, node := range prog.Nodes {
		maybeUML, _ := node.Accept(u.v)
		if uml, ok := maybeUML.(*umlNode); ok {
			u.writeLine("%s --> %s", label, uml.label)
		} else {
			log.Println("Not a UML node, I am confused")
		}
	}
	return labeledNode(prog, label), nil
}
