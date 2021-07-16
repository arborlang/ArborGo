package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

// IfNode represens an if statemet
type IfNode struct {
	Condition Node
	Body      Node
	ElseIfs   []*IfNode
	Else      Node
	ReturnTo  string
	Lexeme    lexer.Lexeme
}

// Accept implements a node
func (i *IfNode) Accept(v Visitor) (Node, error) {
	return v.VisitIfNode(i)
}

func (i *IfNode) GetType() types.TypeNode {
	return &types.FalseType{}
}
