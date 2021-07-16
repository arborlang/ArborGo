package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

// BoolOp represents a boolean operation ('||' and '&&')
type BoolOp struct {
	LeftSide  Node
	RightSide Node
	Condition string
	Lexeme    lexer.Lexeme
}

// Accept visits the node
func (a *BoolOp) Accept(v Visitor) (Node, error) {
	return v.VisitBoolOp(a)
}

func (a *BoolOp) GetType() types.TypeNode {
	return &types.ConstantTypeNode{
		Name: "Boolean",
	}
}
