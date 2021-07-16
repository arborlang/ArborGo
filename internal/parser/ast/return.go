package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

// ReturnNode Represents a return statement
type ReturnNode struct {
	Expression Node
	Lexeme     lexer.Lexeme
}

// Accept allows the vistor to visit
func (r *ReturnNode) Accept(v Visitor) (Node, error) {
	return v.VisitReturnNode(r)
}

func (r *ReturnNode) GetType() types.TypeNode {
	return r.Expression.GetType()
}
