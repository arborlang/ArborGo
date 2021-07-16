package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

// Constant represents a constant definition
type Constant struct {
	Value  string
	Raw    []byte
	Lexeme lexer.Lexeme
	Type   types.TypeNode
}

// Accept visits the node
func (a *Constant) Accept(v Visitor) (Node, error) {
	return v.VisitConstant(a)
}

func (a *Constant) GetType() types.TypeNode {
	return a.Type
}
