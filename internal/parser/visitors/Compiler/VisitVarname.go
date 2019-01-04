package compiler

import (
	"fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
)

// VisitVarName visits a varname node
func (c *Compiler) VisitVarName(node *ast.VarName) (ast.VisitorMetaData, error) {
	sym := c.SymbolTable.GetSymbol(node.Name)
	if sym == nil {
		return ast.VisitorMetaData{}, fmt.Errorf("symbol %s not defined", node.Name)
	}
	tp := sym.Type
	cmd := "i64.load"
	switch tp {
	case "char":
		cmd = "i32.load"
	case "float":
		cmd = "f64.load"
	}
	c.Emit("(%s %s)", cmd, sym.Location)
	return ast.VisitorMetaData{
		Location: sym.Location,
		SymbolData: &ast.SymbolData{
			Name:       node.Name,
			IsNew:      false,
			IsConstant: sym.IsConstant,
		},
	}, nil
}