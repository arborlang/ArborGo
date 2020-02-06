package wast

import (
	// "fmt"
	"github.com/arborlang/ArborGo/internal/parser/ast"
)

// VisitReturnNode visits a return node
func (c *Compiler) VisitReturnNode(node *ast.ReturnNode) (ast.VisitorMetaData, error) {
	metadata, err := node.Expression.Accept(c)
	metadata.Returns = []string{metadata.Types.Types[0]}
	c.EmitFunc("call $__popstack__")
	c.EmitFunc("drop")
	c.EmitFunc("return")
	// fmt.Println(metadata.Types)
	if err != nil {
		return ast.VisitorMetaData{}, err
	}
	// if returnVal.Location != "STACK" {
	// 	c.Emit("(get_local %s)", returnVal.Location)
	// }
	return metadata, nil
}
