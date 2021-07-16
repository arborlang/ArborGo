package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

// PipeNode reperesent a node that is a pipe operator
type PipeNode struct {
	LeftSide  Node
	RightSide Node
	Lexeme    lexer.Lexeme
}

// Accept accepts the visitor
func (p *PipeNode) Accept(v Visitor) (Node, error) {
	return v.VisitPipeNode(p)
}

func (p *PipeNode) GetType() types.TypeNode {
	return p.RightSide.GetType()
}
