package compiler

import (
	"fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
	"strings"
)

// VisitFunctionDefinitionNode visits a function definition ndde
func (c *Compiler) VisitFunctionDefinitionNode(node *ast.FunctionDefinitionNode) (ast.VisitorMetaData, error) {
	tempName := []string{}
	for _, varName := range node.Arguments {
		tempName = append(
			tempName,
			fmt.Sprintf("%s%s", varName.Name, strings.Join(varName.Type.Types, "|")),
		)
	}
	tempName = append(tempName, strings.Join(node.Returns.Types, "|"))
	name := strings.Join(tempName, "_")
	name = c.getUniqueID("func", name)

	c.Emit("(func %s", name)
	args := []string{}
	for _, arg := range node.Arguments {
		name := c.getUniqueID(strings.Join(arg.Type.Types, ""), arg.Name)
		args = append(args, name)
		c.Emit("(param %s i64)", name)
	}
	c.Emit("(result i64)")
	c.Flush()

	// for _, arg := range args {
	// 	c.Emit("(local %s i64)", arg)
	// }
	node.Body.Accept(c)
	// locals := []byte{}
	// for _, sym := range c.SymbolTable.currentScope {
	// 	fmt.Println("{", sym.Location, sym.Type, sym.IsConstant, sym.Name, "}")
	// }
	c.Emit(")")
	// c.PrependAndFlush(locals)

	return ast.VisitorMetaData{
		Location:   name,
		Exportable: fmt.Sprintf("(func %s)", name),
	}, nil
}
