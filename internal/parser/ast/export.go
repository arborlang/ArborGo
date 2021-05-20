package ast

// ExportNode represents an export statement
type InternalNode struct {
	Expression Node
}

// Accept visits the node
func (e *InternalNode) Accept(v Visitor) (Node, error) {
	return v.VisitInternalNode(e)
}
