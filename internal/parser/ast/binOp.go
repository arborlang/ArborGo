package ast

// BoolOp represents a boolean operation ('||' and '&&')
type BoolOp struct {
	LeftSide  Node
	RightSide Node
	Condition string
}

// Accept visits the node
func (a *BoolOp) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitBoolOp(a)
}
