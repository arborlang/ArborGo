package wast

import (
	"bytes"
	"fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
	"strings"
)

// VisitFunctionDefinitionNode visits a function definition ndde
func (c *Compiler) VisitFunctionDefinitionNode(node *ast.FunctionDefinitionNode) (ast.VisitorMetaData, error) {
	c.StartFunc()
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
	signature := &bytes.Buffer{}
	signature.Write([]byte(fmt.Sprintf("func %s", name)))
	args := []string{}
	for _, arg := range node.Arguments {
		name := c.getUniqueID(strings.Join(arg.Type.Types, ""), arg.Name)
		args = append(args, name)
		signature.Write([]byte(fmt.Sprintf("(param %s i64)", name)))
	}
	signature.Write([]byte("(result i64)"))
	c.currentFunc.signature = signature.String()

	metadata, err := node.Body.Accept(c)
	if err != nil {
		return ast.VisitorMetaData{}, err
	}
	// for _, local := range c.locals {
	// 	lc := fmt.Sprintf("(local %s %s)\n", local.name, local.tp)
	// 	locals = append(locals, []byte(lc)...)
	// }
	for _, retType := range metadata.Returns {
		if !node.Returns.IsValidType(retType) {
			fmt.Println("return type:", retType, "want:", node.Returns)
			return ast.VisitorMetaData{}, fmt.Errorf("function does not return a valid type")
		}
	}

	return ast.VisitorMetaData{
		Location:   name,
		Exportable: fmt.Sprintf("(func %s)", name),
	}, nil
}
