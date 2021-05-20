package ast

import "github.com/arborlang/ArborGo/internal/lexer"

// SliceNode Represnets the  slice node
type SliceNode struct {
	Varname Node
	Start   Node
	End     Node
	Step    Node
	Lexeme  lexer.Lexeme
}

// Accept allows the vistor to visit
func (r *SliceNode) Accept(v Visitor) (Node, error) {
	return v.VisitSliceNode(r)
}
