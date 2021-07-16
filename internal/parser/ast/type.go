package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

// TypeNode represents a type node
type TypeNode struct {
	Types   types.TypeNode
	Lexeme  lexer.Lexeme
	VarName *VarName
	Extends bool
}

// Accept a type visitor
func (t *TypeNode) Accept(v Visitor) (Node, error) {
	return v.VisitTypeNode(t)
}

func (t *TypeNode) GetType() types.TypeNode {
	return t.Types
}

// IsValidType Makes sure that a given type stisifies the gaurd
func (t *TypeNode) IsValidType(tp types.TypeNode) bool {
	return t.Types.IsSatisfiedBy(tp)
}

// IsPointer denotes whether a type is a pointer or not
func (t *TypeNode) IsPointer() bool {
	return t.IsValidType(&types.ConstantTypeNode{Name: "String"}) || t.IsValidType(&types.ConstantTypeNode{Name: "Array"})
}
