package ast

import (
	"github.com/radding/ArborGo/internal/lexer"
)

//Declaration is a decleration node
type Declaration struct {
	NameToken lexer.Lexeme
	TypeToken lexer.Lexeme
	IsConst   bool
}

//NewDeclaration returns a new declaration node
func NewDeclaration(name, typeTok lexer.Lexeme, isConst bool) *Declaration {
	return &Declaration{
		NameToken: name,
		TypeToken: typeTok,
		IsConst:   isConst,
	}
}

//Walk compiles the expressions tha
func (s *Declaration) Walk(global, local *CompilerContext, v Visitor) Register {
	v.PreVisit(s, global, nil)
	return v.PostVisit(s, global, nil)
}
