package ast

import "github.com/arborlang/ArborGo/internal/parser/ast/types"

type ContinueNode struct {
	WithValue Node
}

func (c *ContinueNode) Accept(v Visitor) (Node, error) {
	return v.VisitContinueNode(c)
}

func (c *ContinueNode) GetType() types.TypeNode {
	return &types.FalseType{}
}
