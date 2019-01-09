package ast

// Constant represents a constant definition
type Constant struct {
	Value string
	Type  string
	Raw   []byte
}

// Accept visits the node
func (a *Constant) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitConstant(a)
}
