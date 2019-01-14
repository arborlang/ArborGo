package ast

// IndexNode is the index node
type IndexNode struct {
	Varname *VarName
	index   int
}

// Accept a visitor
func (i *IndexNode) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitIndexNode(i)
}
