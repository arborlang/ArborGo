package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

type WhenNode struct {
	Case     Node
	Evaluate Node
	Lexeme   lexer.Lexeme
}

type MatchNode struct {
	Match     Node
	WhenCases []*WhenNode
	Lexeme    lexer.Lexeme
}

func (m *MatchNode) Accept(v Visitor) (Node, error) {
	return v.VisitMatchNode(m)
}

func (w *WhenNode) Accept(v Visitor) (Node, error) {
	return v.VisitWhenNode(w)
}

func (m *MatchNode) GetType() types.TypeNode {
	return &types.FalseType{}
}
func (m *WhenNode) GetType() types.TypeNode {
	return &types.FalseType{}
}
