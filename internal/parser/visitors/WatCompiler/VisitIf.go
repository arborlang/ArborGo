package wast

import (
	// "fmt"
	"github.com/arborlang/ArborGo/internal/parser/ast"
)

// VisitIfNode visits an if node
func (c *Compiler) VisitIfNode(node *ast.IfNode) (ast.VisitorMetaData, error) {
	// c.EmitFunc("drop")
	if len(node.ElseIfs) > 0 {
		return visitIfElseIf(node, c)
	}
	if node.Else != nil {
		return visitIfElse(node, c)
	}
	return visitIf(node, c, "")
}

func visitIf(node *ast.IfNode, c *Compiler, breakTo string) (ast.VisitorMetaData, error) {
	label1 := c.getLabelUniqueID("if_else_if")
	c.EmitFunc("block %s", label1)
	metadata, err := node.Condition.Accept(c)
	if err != nil {
		return metadata, err
	}
	tp := c.getType(metadata.Types.Types[0])
	c.EmitFunc("%s.eqz", tp)
	c.EmitFunc("br_if %s", label1)
	// c.Emit("(block")
	metadata, err = node.Body.Accept(c)
	if breakTo != "" {
		c.EmitFunc("br %s", breakTo)
	}
	// c.EmitFunc("drop")
	c.EmitFunc("end %s", label1)
	return ast.VisitorMetaData{}, nil
}

func visitIfElse(node *ast.IfNode, c *Compiler) (ast.VisitorMetaData, error) {
	label1 := c.getLabelUniqueID("if")
	label2 := c.getLabelUniqueID("if")
	c.EmitFunc("block %s", label1)
	c.EmitFunc("block %s", label2)
	metadata, err := node.Condition.Accept(c)
	if err != nil {
		return metadata, err
	}
	tp := c.getType(metadata.Types.Types[0])
	c.EmitFunc("%s.eqz", tp)
	c.EmitFunc("br_if %s", label2)
	// c.Emit("(block")
	metadata, err = node.Body.Accept(c)
	c.EmitFunc("br %s", label1)
	// c.EmitFunc("drop")
	c.EmitFunc("end %s", label2)
	node.Else.Accept(c)
	// c.EmitFunc("drop")
	c.EmitFunc("end %s", label1)
	return ast.VisitorMetaData{}, nil
}

func visitIfElseIf(node *ast.IfNode, c *Compiler) (ast.VisitorMetaData, error) {
	masterBlock := c.getLabelUniqueID("if")
	c.EmitFunc("block %s", masterBlock)
	metadata, err := visitIf(node, c, masterBlock)
	if err != nil {
		return metadata, err
	}
	for _, elseIf := range node.ElseIfs {
		metadata, err = visitIf(elseIf, c, masterBlock)
		if err != nil {
			return metadata, err
		}
	}
	if node.Else != nil {
		metadata, err = node.Else.Accept(c)
		if err != nil {
			return metadata, err
		}
	}
	// c.Emit("drop")
	c.EmitFunc("end %s", masterBlock)
	return metadata, err
}
