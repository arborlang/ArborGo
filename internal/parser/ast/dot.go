package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

type DotNode struct {
	Lexeme  lexer.Lexeme
	VarName Node
	Access  Node
}

func (d *DotNode) Accept(v Visitor) (Node, error) {
	return v.VisitDotNode(d)
}

func (d *DotNode) GetType() types.TypeNode {
	return d.Access.GetType()
}
