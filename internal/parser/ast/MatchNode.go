package ast

import "github.com/arborlang/ArborGo/internal/parser/ast/types"

type WhenNode struct {
	Case     Node
	Evaluate Node
}

type MatchNode struct {
	Match     Node
	WhenCases []*WhenNode
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
