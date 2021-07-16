package ast

import "github.com/arborlang/ArborGo/internal/parser/ast/types"

type DecoratorNode struct {
	Name      *VarName
	Decorates Node
}

func (d *DecoratorNode) Accept(v Visitor) (Node, error) {
	return v.VisitDecoratorNode(d)
}

func (d *DecoratorNode) GetType() types.TypeNode {
	return d.Name.GetType()
}
