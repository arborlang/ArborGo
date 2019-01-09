package wast

import (
	"fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
	"strings"
)

// VisitDeclNode visits the decl Node
func (c *Compiler) VisitDeclNode(node *ast.DeclNode) (ast.VisitorMetaData, error) {
	tp := ""
	if node.Varname.Type != nil {
		tp = strings.Join(node.Varname.Type.Types, "")
	}
	loc := c.SymbolTable.GetSymbol(node.Varname.Name)
	if loc != nil {
		return ast.VisitorMetaData{}, fmt.Errorf("redefined symbol %s", node.Varname.Name)
	}
	c.SymbolTable.AddToScope(&Symbol{
		Name:       node.Varname.Name,
		Location:   "undefined",
		IsConstant: node.IsConstant,
		Type:       tp,
	})
	return ast.VisitorMetaData{
		Location: "",
		Types:    tp,
		SymbolData: &ast.SymbolData{
			Name:       node.Varname.Name,
			Type:       node.Varname.Type,
			IsConstant: node.IsConstant,
			IsNew:      true,
		},
	}, nil
}
