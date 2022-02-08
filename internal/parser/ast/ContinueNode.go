package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

type ContinueNode struct {
	WithValue Node
	Lexeme    lexer.Lexeme
}

func (c *ContinueNode) Accept(v Visitor) (Node, error) {
	return v.VisitContinueNode(c)
}

func (c *ContinueNode) GetType() types.TypeNode {
	return &types.FalseType{}
}
