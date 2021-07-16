package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

// IndexNode is the index node
type IndexNode struct {
	Varname Node
	Index   Node
	Lexeme  lexer.Lexeme
}

// Accept a visitor
func (i *IndexNode) Accept(v Visitor) (Node, error) {
	return v.VisitIndexNode(i)
}

func (i *IndexNode) GetType() types.TypeNode {
	if sliceList, ok := i.Varname.GetType().(*types.ArrayType); ok {
		return sliceList.SubType
	}
	return &types.FalseType{}
}
