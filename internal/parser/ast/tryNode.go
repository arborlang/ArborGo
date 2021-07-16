package ast

import "github.com/arborlang/ArborGo/internal/parser/ast/types"

type TryNode struct {
	Tries       Node
	HandleCases []*HandleCaseNode
}

func (t *TryNode) Accept(v Visitor) (Node, error) {
	return v.VisitTryNode(t)
}

func (t *TryNode) GetType() types.TypeNode {
	return &types.FalseType{}
}

type HandleCaseNode struct {
	VariableName string
	Type         types.TypeNode
	Case         Node
}

func (t *HandleCaseNode) Accept(v Visitor) (Node, error) {
	return v.VisitHandleCaseNode(t)
}

func (t *HandleCaseNode) GetType() types.TypeNode {
	return &types.FalseType{}
}
