package ast

import "github.com/arborlang/ArborGo/internal/lexer"

// Constant represents a constant definition
type Constant struct {
	Value  string
	Type   string
	Raw    []byte
	Lexeme lexer.Lexeme
}

// Accept visits the node
func (a *Constant) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitConstant(a)
}
