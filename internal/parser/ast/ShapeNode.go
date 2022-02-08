package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

type ShapeNode struct {
	Fields       map[string]Node
	Lexeme       lexer.Lexeme
	GenericNames []*VarName
}

func (s *ShapeNode) Accept(v Visitor) (Node, error) {
	return v.VisitShapeNode(s)
}

func (s *ShapeNode) GetType() types.TypeNode {
	return &types.FalseType{}
}
