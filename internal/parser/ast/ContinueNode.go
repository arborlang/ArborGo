package ast

type ContinueNode struct {
	WithValue Node
}

func (c *ContinueNode) Accept(v Visitor) (Node, error) {
	return v.VisitContinueNode(c)
}
