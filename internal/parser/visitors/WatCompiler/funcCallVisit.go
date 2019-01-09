package wast

import (
	"github.com/radding/ArborGo/internal/parser/ast"
)

// VisitFunctionCallNode visits a function call node
func (c *Compiler) VisitFunctionCallNode(node *ast.FunctionCallNode) (ast.VisitorMetaData, error) {
	for _, arg := range node.Arguments {
		metadata, err := arg.Accept(c)
		if err != nil {
			return metadata, err
		}
	}
	varName, ok := node.Definition.(*ast.VarName)
	if ok {
		if varName.Name == "__putch__" {
			c.Emit("call $__putch__")
		}
	}
	return ast.VisitorMetaData{}, nil
}
