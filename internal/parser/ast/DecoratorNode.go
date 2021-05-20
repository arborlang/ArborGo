package ast

type DecoratorNode struct {
	Name      *VarName
	Decorates Node
}

func (d *DecoratorNode) Accept(v Visitor) (Node, error) {
	return v.VisitDecoratorNode(d)
}
