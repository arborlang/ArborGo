package ast

import "github.com/arborlang/ArborGo/internal/lexer"

// TypeNode represents a type node
type TypeNode struct {
	Types  []string
	Lexeme lexer.Lexeme
}

// Accept a type visitor
func (t *TypeNode) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitTypeNode(t)
}

// IsValidType Makes sure that a given type stisifies the gaurd
func (t *TypeNode) IsValidType(tp string) bool {
	for _, typ := range t.Types {
		if typ == tp {
			return true
		}
	}
	return false
}

// IsPointer denotes whether a type is a pointer or not
func (t *TypeNode) IsPointer() bool {
	return t.IsValidType("string") || t.IsValidType("array")
}
