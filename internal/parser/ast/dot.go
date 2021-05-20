package ast

type DotNode struct {
	VarName Node
	Access  Node
}

func (d *DotNode) Accept(v Visitor) (Node, error) {
	return v.VisitDotNode(d)
}
