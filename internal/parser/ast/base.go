package ast

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

//go:generate ../../../utils/gen_visitors -type=Node

// Node represents a Node in the AST
type Node interface {
	// Accept allows a visitor to traverse the tree
	Accept(Visitor) (Node, error)
}

//SymbolData is some data for the symbol in our symbol table
type SymbolData struct {
	Name       string
	Type       *TypeNode
	IsConstant bool
	IsNew      bool
}

// VisitorMetaData is what the visitor will return to represent the conclusion of it work on a node
type AnnotatedNode struct {
	VarName  string
	Location string
	Types    types.TypeNode
	IsDefine bool
	Lexeme   lexer.Lexeme
	Node     Node
}

func (a AnnotatedNode) Accept(v Visitor) (Node, error) {
	if a.Node == nil {
		return AnnotatedNode{}, fmt.Errorf("annotated node has no node")
	}
	return a.Node.Accept(v)
}
