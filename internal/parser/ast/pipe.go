package ast

import "github.com/arborlang/ArborGo/internal/lexer"

// PipeNode reperesent a node that is a pipe operator
type PipeNode struct {
	LeftSide  Node
	RightSide Node
	Lexeme    lexer.Lexeme
}

// Accept accepts the visitor
func (p *PipeNode) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitPipeNode(p)
}
