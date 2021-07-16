package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

type Package struct {
	Lexeme lexer.Lexeme
	Name   string
}

func (p *Package) Accept(v Visitor) (Node, error) {
	return v.VisitPackage(p)
}

func (p *Package) GetType() types.TypeNode {
	return &types.FalseType{}
}
