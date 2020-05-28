package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

//VarName Represents a name for a variable
type VarName struct {
	Name   string
	Type   *types.TypeNode
	Lexeme lexer.Lexeme
}

// Accept does nothing
func (vn *VarName) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitVarName(vn)
}

//DeclNode is a node that
type DeclNode struct {
	Varname    VarName
	IsConstant bool
}

// Accept does nothing
func (d *DeclNode) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitDeclNode(d)
}

// AddChild just sets the
func (d *DeclNode) AddChild(c Node) error {
	d.Varname = *c.(*VarName)
	return nil
}
