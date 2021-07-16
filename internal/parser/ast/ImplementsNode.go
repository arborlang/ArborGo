package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

type ImplementsNode struct {
	Implements []*VarName
	Lexeme     lexer.Lexeme
	Type       *TypeNode
}

func (i *ImplementsNode) Accept(v Visitor) (Node, error) {
	return v.VisitImplementsNode(i)
}

func (i *ImplementsNode) GetType() types.TypeNode {
	return &types.FalseType{}
}
