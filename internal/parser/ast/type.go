package ast

// TypeNode represents a type node
type TypeNode struct {
	Types []string
}

// Accept a type visitor
func (t *TypeNode) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitTypeNode(t)
}
