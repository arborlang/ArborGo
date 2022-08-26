package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

//Program is the Root node in a file
type Program struct {
	Imports []*ImportNode `@@*`
	Nodes   []Node
	Lexeme  lexer.Lexeme
}

// Accept Accepts a vistor
func (s *Program) Accept(v Visitor) (Node, error) {
	return v.VisitProgram(s)
}

func (s *Program) GetType() types.TypeNode {
	return &types.FalseType{}
}
