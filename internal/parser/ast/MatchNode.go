package ast

type WhenNode struct {
	Case     Node
	Evaluate Node
}

type MatchNode struct {
	Match     Node
	WhenCases []*WhenNode
}

func (m *MatchNode) Accept(v Visitor) (Node, error) {
	return v.VisitMatchNode(m)
}

func (w *WhenNode) Accept(v Visitor) (Node, error) {
	return v.VisitWhenNode(w)
}
