package wast

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
)

// VisitVarName visits a varname node
func (c *Compiler) VisitVarName(node *ast.VarName) (ast.VisitorMetaData, error) {
	sym := c.SymbolTable.GetSymbol(node.Name)
	if sym == nil {
		return ast.VisitorMetaData{}, fmt.Errorf("symbol %s not defined", node.Name)
	}
	// tp := sym.Type
	// cmd := "i64.load"
	// switch tp {
	// case "char":
	// 	cmd = "i32.load"
	// case "float":
	// 	cmd = "f64.load"
	// }
	c.EmitFunc("get_local %s", sym.Location)
	return ast.VisitorMetaData{
		Location: sym.Location,
		Types:    sym.Type,
		SymbolData: &ast.SymbolData{
			Name:       node.Name,
			IsNew:      false,
			IsConstant: sym.IsConstant,
		},
	}, nil
}
