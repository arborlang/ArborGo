package ast

// ReturnNode Represents a return statement
type ReturnNode struct {
	Expression Node
}

// Accept allows the vistor to visit
func (r *ReturnNode) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitReturnNode(r)
}
