package ast

// IfNode represens an if statemet
type IfNode struct {
	Condition Node
	Body      Node
	ElseIfs   []*IfNode
	Else      Node
	ReturnTo  string
}

// Accept implements a node
func (i *IfNode) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitIfNode(i)
}
