package ast

import "github.com/arborlang/ArborGo/internal/lexer"

// Comparison represents a comparison operator
type Comparison struct {
	LeftSide  Node
	RightSide Node
	Operation string
	Lexeme    lexer.Lexeme
}

// Accept visits the node
func (a *Comparison) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitComparison(a)
}
