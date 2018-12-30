package ast

// AssignmentNode is a node that represents an asignment operator
type AssignmentNode struct {
	AssignTo Node
	Value    Node
}

// Accept visits the node
func (a *AssignmentNode) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitAssignment(a)
}
