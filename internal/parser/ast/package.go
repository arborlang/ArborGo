package ast

import "github.com/arborlang/ArborGo/internal/lexer"

type Package struct {
	Lexeme lexer.Lexeme
	Name   string
}

func (p *Package) Accept(v Visitor) (Node, error) {
	return v.VisitPackage(p)
}
