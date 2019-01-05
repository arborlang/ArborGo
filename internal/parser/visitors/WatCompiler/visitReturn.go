package compiler

import (
	"github.com/radding/ArborGo/internal/parser/ast"
)

// VisitReturnNode visits a return node
func (c *Compiler) VisitReturnNode(node *ast.ReturnNode) (ast.VisitorMetaData, error) {
	_, err := node.Expression.Accept(c)
	if err != nil {
		return ast.VisitorMetaData{}, err
	}
	// if returnVal.Location != "STACK" {
	// 	c.Emit("(get_local %s)", returnVal.Location)
	// }
	return ast.VisitorMetaData{}, nil
}
