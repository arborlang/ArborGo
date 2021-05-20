package ast

import "github.com/arborlang/ArborGo/internal/lexer"

// MathOpNode is a struct representing a type of Mathematical operation ('+', '-', '/', '*')
//! <Leftside> [math operation] <RightSide>
type MathOpNode struct {
	LeftSide  Node
	RightSide Node
	Operation string
	Lexeme    lexer.Lexeme
}

// Accept accepts the Visitor
func (m *MathOpNode) Accept(v Visitor) (Node, error) {
	return v.VisitMathOpNode(m)
}
