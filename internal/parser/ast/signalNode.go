package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

type SignalNode struct {
	Level        string
	ValueToRaise Node
	Lexeme       lexer.Lexeme
}

func (s *SignalNode) Accept(v Visitor) (Node, error) {
	return v.VisitSignalNode(s)
}

func (s *SignalNode) GetType() types.TypeNode {
	return &types.FalseType{}
}
