package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

// ExportNode represents an export statement
type InternalNode struct {
	Lexeme     lexer.Lexeme
	Expression Node
}

// Accept visits the node
func (e *InternalNode) Accept(v Visitor) (Node, error) {
	return v.VisitInternalNode(e)
}

func (e *InternalNode) GetType() types.TypeNode {
	return &types.FnType{}
}
