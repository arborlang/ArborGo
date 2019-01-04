package compiler

import (
	"github.com/radding/ArborGo/internal/parser/ast"
)

// VisitBlock visits a compiler block
func (c *Compiler) VisitBlock(block *ast.Program) (ast.VisitorMetaData, error) {
	c.Level++
	defer func() { c.Level-- }()
	c.SymbolTable.PushScope()
	defer c.SymbolTable.PopScope()
	for _, stmt := range block.Nodes {
		if _, err := stmt.Accept(c); err != nil {
			return ast.VisitorMetaData{}, err
		}

	}
	return ast.VisitorMetaData{}, nil
}

// VisitBoolOp visits a boolean node
func (c *Compiler) VisitBoolOp(node *ast.BoolOp) (ast.VisitorMetaData, error) {
	return ast.VisitorMetaData{}, nil
}

// VisitFunctionCallNode visits a function call node
func (c *Compiler) VisitFunctionCallNode(node *ast.FunctionCallNode) (ast.VisitorMetaData, error) {
	return ast.VisitorMetaData{}, nil
}

// VisitMathOpNode Visits a math op node
func (c *Compiler) VisitMathOpNode(node *ast.MathOpNode) (ast.VisitorMetaData, error) {
	return ast.VisitorMetaData{}, nil
}

// VisitReturnNode visits a return node
func (c *Compiler) VisitReturnNode(node *ast.ReturnNode) (ast.VisitorMetaData, error) {
	return ast.VisitorMetaData{}, nil
}

// VisitPipeNode visits the pipe node
func (c *Compiler) VisitPipeNode(node *ast.PipeNode) (ast.VisitorMetaData, error) {
	return ast.VisitorMetaData{}, nil
}

// VisitIfNode visits an if node
func (c *Compiler) VisitIfNode(node *ast.IfNode) (ast.VisitorMetaData, error) {
	return ast.VisitorMetaData{}, nil
}

// VisitImportNode visits an import node
func (c *Compiler) VisitImportNode(node *ast.ImportNode) (ast.VisitorMetaData, error) {
	return ast.VisitorMetaData{}, nil
}

// VisitTypeNode visits a type node
func (c *Compiler) VisitTypeNode(node *ast.TypeNode) (ast.VisitorMetaData, error) {
	return ast.VisitorMetaData{}, nil
}
