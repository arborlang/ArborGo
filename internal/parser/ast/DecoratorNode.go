package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

type DecoratorNode struct {
	Name      *VarName
	Lexeme    lexer.Lexeme
	Decorates Node
}

func (d *DecoratorNode) Accept(v Visitor) (Node, error) {
	return v.VisitDecoratorNode(d)
}

func (d *DecoratorNode) GetType() types.TypeNode {
	return d.Name.GetType()
}
