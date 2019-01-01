package compiler

import (
	"fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
)

// VisitBlock visits a compiler block
func (c *Compiler) VisitBlock(block *ast.Program) (ast.VisitorMetaData, error) {
	c.SymbolTable.PushScope()
	for _, stmt := range block.Nodes {
		stmt.Accept(c)
	}
	return ast.VisitorMetaData{}, nil
}

// VisitAssignment visits an assignment node
func (c *Compiler) VisitAssignment(assignment *ast.AssignmentNode) (ast.VisitorMetaData, error) {
	location, err := assignment.AssignTo.Accept(c)
	if err != nil {
		return ast.VisitorMetaData{}, err
	}
	sym := c.SymbolTable.GetSymbol(location.Location)
	if sym == nil {
		return ast.VisitorMetaData{}, fmt.Errorf("symbol %s not defined", location.Location)
	}
	result, err := assignment.Value.Accept(c)
	if err != nil {
		return ast.VisitorMetaData{}, err
	}
	fmt.Println(result)
	return ast.VisitorMetaData{}, nil
}

// VisitBoolOp visits a boolean node
func (c *Compiler) VisitBoolOp(node *ast.BoolOp) (ast.VisitorMetaData, error) {
	return ast.VisitorMetaData{}, nil
}

// VisitComparison Visits a comparison node
func (c *Compiler) VisitComparison(node *ast.Comparison) (ast.VisitorMetaData, error) {
	return ast.VisitorMetaData{}, nil
}

// VisitConstant visits the constant object
func (c *Compiler) VisitConstant(node *ast.Constant) (ast.VisitorMetaData, error) {
	return ast.VisitorMetaData{}, nil
}

// VisitFunctionDefinitionNode visits a function definition ndde
func (c *Compiler) VisitFunctionDefinitionNode(node *ast.FunctionDefinitionNode) (ast.VisitorMetaData, error) {
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

// VisitVarName visits a varname node
func (c *Compiler) VisitVarName(node *ast.VarName) (ast.VisitorMetaData, error) {
	return ast.VisitorMetaData{}, nil
}

// VisitDeclNode visits the decl Node
func (c *Compiler) VisitDeclNode(node *ast.DeclNode) (ast.VisitorMetaData, error) {
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
