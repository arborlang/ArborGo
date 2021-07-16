package ast

import "github.com/arborlang/ArborGo/internal/parser/ast/types"

// ExtendsNode represents an extension to a node
type ExtendsNode struct {
	Name   *VarName
	Extend *VarName
	Adds   types.TypeNode
}

func (e *ExtendsNode) Accept(v Visitor) (Node, error) {
	return v.VisitExtendsNode(e)
}

func (e *ExtendsNode) GetType() types.TypeNode {
	return &types.FalseType{}
}
