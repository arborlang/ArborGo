package ast

import "github.com/arborlang/ArborGo/internal/lexer"

// ReturnNode Represents a return statement
type ReturnNode struct {
	Expression Node
	Lexeme     lexer.Lexeme
}

// Accept allows the vistor to visit
func (r *ReturnNode) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitReturnNode(r)
}
