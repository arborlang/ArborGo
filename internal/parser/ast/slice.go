package ast

// SliceNode Represnets the  slice node
type SliceNode struct {
	Varname *VarName
	Start   int
	End     int
	Step    int
}

// Accept allows the vistor to visit
func (r *SliceNode) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitSliceNode(r)
}
