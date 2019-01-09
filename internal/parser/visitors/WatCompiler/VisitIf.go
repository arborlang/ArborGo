package wast

import (
	// "fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
)

// VisitIfNode visits an if node
func (c *Compiler) VisitIfNode(node *ast.IfNode) (ast.VisitorMetaData, error) {
	c.Emit("drop")
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
	c.Emit("block %s", label1)
	metadata, err := node.Condition.Accept(c)
	if err != nil {
		return metadata, err
	}
	tp := c.getType(metadata.Types)
	c.Emit("%s.eqz", tp)
	c.Emit("br_if %s", label1)
	// c.Emit("(block")
	metadata, err = node.Body.Accept(c)
	if breakTo != "" {
		c.Emit("br %s", breakTo)
	}
	c.Emit("drop")
	c.Emit("end %s", label1)
	return ast.VisitorMetaData{}, nil
}

func visitIfElse(node *ast.IfNode, c *Compiler) (ast.VisitorMetaData, error) {
	label1 := c.getLabelUniqueID("if")
	label2 := c.getLabelUniqueID("if")
	c.Emit("block %s", label1)
	c.Emit("block %s", label2)
	metadata, err := node.Condition.Accept(c)
	if err != nil {
		return metadata, err
	}
	tp := c.getType(metadata.Types)
	c.Emit("%s.eqz", tp)
	c.Emit("br_if %s", label2)
	// c.Emit("(block")
	metadata, err = node.Body.Accept(c)
	c.Emit("br %s", label1)
	c.Emit("drop")
	c.Emit("end %s", label2)
	node.Else.Accept(c)
	c.Emit("drop")
	c.Emit("end %s", label1)
	return ast.VisitorMetaData{}, nil
}

func visitIfElseIf(node *ast.IfNode, c *Compiler) (ast.VisitorMetaData, error) {
	masterBlock := c.getLabelUniqueID("if")
	c.Emit("block %s", masterBlock)
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
	c.Emit("end %s", masterBlock)
	return metadata, err
}
