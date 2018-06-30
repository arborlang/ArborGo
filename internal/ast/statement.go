package ast

//Expression is the expression in the statement node.
type Expression struct {
	expressions []Node
}

//Walk compiles the expressions tha
func (s *Expression) Walk(global, local *CompilerContext, v Visitor) Register {
	v.PreVisit(s, global, nil)
	for _, expr := range s.expressions {
		expr.Walk(global, nil, v)
	}
	return v.PostVisit(s, global, nil)
}
