package ast

import "github.com/arborlang/ArborGo/internal/lexer"

type ImplementsNode struct {
	Implements []*VarName
	Lexeme     lexer.Lexeme
	Type       *TypeNode
}

func (i *ImplementsNode) Accept(v Visitor) (Node, error) {
	return v.VisitImplementsNode(i)
}
