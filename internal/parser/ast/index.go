package ast

import "github.com/arborlang/ArborGo/internal/lexer"

// IndexNode is the index node
type IndexNode struct {
	Varname *VarName
	Index   int
	Lexeme  lexer.Lexeme
}

// Accept a visitor
func (i *IndexNode) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitIndexNode(i)
}
