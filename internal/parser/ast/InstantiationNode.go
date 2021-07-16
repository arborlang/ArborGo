package ast

import "github.com/arborlang/ArborGo/internal/parser/ast/types"

type InstantiateNode struct {
	FunctionCallNode *FunctionCallNode
}

func (i *InstantiateNode) Accept(v Visitor) (Node, error) {
	return v.VisitInstantiateNode(i)
}

func (i *InstantiateNode) GetType() types.TypeNode {
	return i.FunctionCallNode.Definition.GetType()
}
