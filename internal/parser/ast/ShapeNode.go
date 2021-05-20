package ast

type ShapeNode struct {
	Fields map[string]Node
}

func (s *ShapeNode) Accept(v Visitor) (Node, error) {
	return v.VisitShapeNode(s)
}
