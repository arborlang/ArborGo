package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

// Comparison represents a comparison operator
type Comparison struct {
	LeftSide  Node
	RightSide Node
	Operation string
	Lexeme    lexer.Lexeme
}

// Accept visits the node
func (a *Comparison) Accept(v Visitor) (Node, error) {
	return v.VisitComparison(a)
}

func (a *Comparison) GetType() types.TypeNode {
	return a.LeftSide.GetType()
}
