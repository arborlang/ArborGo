package wast

import (
	// "fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
)

// VisitReturnNode visits a return node
func (c *Compiler) VisitReturnNode(node *ast.ReturnNode) (ast.VisitorMetaData, error) {
	metadata, err := node.Expression.Accept(c)
	metadata.Returns = []string{metadata.Types}
	c.Emit("return")
	// fmt.Println(metadata.Types)
	if err != nil {
		return ast.VisitorMetaData{}, err
	}
	// if returnVal.Location != "STACK" {
	// 	c.Emit("(get_local %s)", returnVal.Location)
	// }
	return metadata, nil
}
