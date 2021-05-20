package ast

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

//VarName Represents a name for a variable
type VarName struct {
	Name   string
	Type   types.TypeNode
	Lexeme lexer.Lexeme
}

// Accept does nothing
func (vn *VarName) Accept(v Visitor) (Node, error) {
	return v.VisitVarName(vn)
}

//DeclNode is a node that
type DeclNode struct {
	Varname    *VarName
	IsConstant bool
}

// Accept does nothing
func (d *DeclNode) Accept(v Visitor) (Node, error) {
	return v.VisitDeclNode(d)
}

// AddChild just sets the
func (d *DeclNode) AddChild(c Node) error {
	varname, ok := c.(*VarName)
	if !ok {
		return fmt.Errorf("got unexpected node. Wanted a varname got %s", c)
	}
	d.Varname = varname
	return nil
}
