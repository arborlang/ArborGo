package ast

// MathOpNode is a struct representing a type of Mathematical operation ('+', '-', '/', '*')
//! <Leftside> [math operation] <RightSide>
type MathOpNode struct {
	LeftSide  Node
	RightSide Node
	Operation string
}

// Accept accepts the Visitor
func (m *MathOpNode) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitMathOpNode(m)
}
