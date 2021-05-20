package ast

import "github.com/arborlang/ArborGo/internal/lexer"

// ImportNode represents an import statement
type ImportNode struct {
	Source     string
	ImportAs   string
	ExportName string
	Lexeme     lexer.Lexeme
	NextImport *ImportNode
}

// Accept a visitor
func (i *ImportNode) Accept(v Visitor) (Node, error) {
	return v.VisitImportNode(i)
}
