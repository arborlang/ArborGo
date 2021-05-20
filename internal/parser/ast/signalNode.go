package ast

type SignalNode struct {
	Level        string
	ValueToRaise Node
}

func (s *SignalNode) Accept(v Visitor) (Node, error) {
	return v.VisitSignalNode(s)
}
