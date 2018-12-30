package ast

// Comparison represents a comparison operator
type Comparison struct {
	LeftSide  Node
	RightSide Node
	Operation string
}

// Accept visits the node
func (a *Comparison) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitComparison(a)
}
