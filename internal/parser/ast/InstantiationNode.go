package ast

type InstantiateNode struct {
	FunctionCallNode *FunctionCallNode
}

func (i *InstantiateNode) Accept(v Visitor) (Node, error) {
	return v.VisitInstantiateNode(i)
}
