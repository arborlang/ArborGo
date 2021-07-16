package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

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

func (m *MathOpNode) GetType() types.TypeNode {
	return m.LeftSide.GetType()
}
