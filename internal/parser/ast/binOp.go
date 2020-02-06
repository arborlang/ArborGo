package ast

import "github.com/arborlang/ArborGo/internal/lexer"

// BoolOp represents a boolean operation ('||' and '&&')
type BoolOp struct {
	LeftSide  Node
	RightSide Node
	Condition string
	Lexeme    lexer.Lexeme
}

// Accept visits the node
func (a *BoolOp) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitBoolOp(a)
}
