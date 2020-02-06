package ast

import "github.com/arborlang/ArborGo/internal/lexer"

// ImportNode represents an import statement
type ImportNode struct {
	Source string
	Name   string
	Lexeme lexer.Lexeme
}

// Accept a visitor
func (i *ImportNode) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitImportNode(i)
}
