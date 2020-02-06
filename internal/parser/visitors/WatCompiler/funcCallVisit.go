package wast

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
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
		switch varName.Name {
		case "__putch__":
			c.EmitFunc("call $__putch__")
		case "__alloc__":
			c.EmitFunc("call $__alloc__")
		case "len":
			c.EmitFunc("call $__len__")
			return ast.VisitorMetaData{
				Location: "$__len__",
				Types:    &ast.TypeNode{Types: []string{"number"}},
			}, nil
		case "pause":
			c.EmitFunc("call $__break__")
		default:
			sym := c.SymbolTable.GetSymbol(varName.Name)
			if sym == nil {
				return ast.VisitorMetaData{}, fmt.Errorf("no function named %s", varName.Name)
			}
			c.EmitFunc(";;calling %s", varName.Name)
			c.EmitFunc("call %s", sym.Location)
		}
	}
	return ast.VisitorMetaData{}, nil
}
