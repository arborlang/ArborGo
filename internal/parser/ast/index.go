package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
)

// IndexNode is the index node
type IndexNode struct {
	Varname Node
	Index   Node
	Lexeme  lexer.Lexeme
}

// Accept a visitor
func (i *IndexNode) Accept(v Visitor) (Node, error) {
	return v.VisitIndexNode(i)
}
