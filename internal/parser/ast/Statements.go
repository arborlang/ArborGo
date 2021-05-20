package ast

import "github.com/arborlang/ArborGo/internal/lexer"

//Program is the Root node in a file
type Program struct {
	Nodes  []Node
	Lexeme lexer.Lexeme
}

// Accept Accepts a vistor
func (s *Program) Accept(v Visitor) (Node, error) {
	return v.VisitProgram(s)
}
