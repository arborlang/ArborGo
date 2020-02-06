package base

import (
	"github.com/arborlang/ArborGo/internal/parser/ast"
)

// VisitorHider is a simple way to set and hide the visitor
type VisitorHider interface {
	SetVisitor(v *Visitor)
}

// Visitor represents a top level Visitor that walks the tree but does nothing. Useful for doing analysis on the AST by other visitors
// For example, if I want to collect all of the function definitions of an arbor file, I would define a struct that is composed of this
// visitor and implements the VisitFunctionDefinitionNode function.
type Visitor struct {
	visitor           VisitorHider // visitor implments a visitor interface
	ShouldCallVisitor bool
}

// New returns a new Visitor
func New(visitor VisitorHider) *Visitor {
	v := &Visitor{
		visitor:           visitor,
		ShouldCallVisitor: true,
	}
	visitor.SetVisitor(v)
	return v
}

// GetVisitor gets the underlying visitor
func (v *Visitor) GetVisitor() interface{} {
	return v.visitor
}

// VisitBlock visits a compiler block
func (v *Visitor) VisitBlock(block *ast.Program) (ast.VisitorMetaData, error) {
	if visitor, ok := v.visitor.(ast.BlockVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitBlock(block)
	}
	v.ShouldCallVisitor = true
	for _, stmt := range block.Nodes {
		stmt.Accept(v)
	}
	return ast.VisitorMetaData{}, nil
}

// VisitAssignment visits an assignment node
func (v *Visitor) VisitAssignment(assignment *ast.AssignmentNode) (ast.VisitorMetaData, error) {
	if visitor, ok := v.visitor.(ast.AssignmentVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitAssignment(assignment)
	}
	v.ShouldCallVisitor = true
	assignment.AssignTo.Accept(v)
	assignment.Value.Accept(v)
	return ast.VisitorMetaData{}, nil
}

// VisitBoolOp visits a boolean node
func (v *Visitor) VisitBoolOp(node *ast.BoolOp) (ast.VisitorMetaData, error) {
	if visitor, ok := v.visitor.(ast.BoolOpVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitBoolOp(node)
	}
	v.ShouldCallVisitor = true
	node.LeftSide.Accept(v)
	node.RightSide.Accept(v)
	return ast.VisitorMetaData{}, nil
}

// VisitComparison Visits a comparison node
func (v *Visitor) VisitComparison(node *ast.Comparison) (ast.VisitorMetaData, error) {
	if visitor, ok := v.visitor.(ast.ComparisonVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitComparison(node)
	}
	v.ShouldCallVisitor = true
	node.LeftSide.Accept(v)
	node.RightSide.Accept(v)
	return ast.VisitorMetaData{}, nil
}

// VisitConstant visits the constant object
func (v *Visitor) VisitConstant(node *ast.Constant) (ast.VisitorMetaData, error) {
	if visitor, ok := v.visitor.(ast.ConstantVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitConstant(node)
	}
	v.ShouldCallVisitor = true
	return ast.VisitorMetaData{}, nil
}

// VisitFunctionDefinitionNode visits a function definition ndde
func (v *Visitor) VisitFunctionDefinitionNode(node *ast.FunctionDefinitionNode) (ast.VisitorMetaData, error) {
	if visitor, ok := v.visitor.(ast.FunctionDefinitionNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitFunctionDefinitionNode(node)
	}
	v.ShouldCallVisitor = true
	for _, arg := range node.Arguments {
		arg.Accept(v)
	}
	node.Body.Accept(v)
	return ast.VisitorMetaData{}, nil
}

// VisitFunctionCallNode visits a function call node
func (v *Visitor) VisitFunctionCallNode(node *ast.FunctionCallNode) (ast.VisitorMetaData, error) {
	if visitor, ok := v.visitor.(ast.FunctionCallNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitFunctionCallNode(node)
	}
	v.ShouldCallVisitor = true
	node.Definition.Accept(v)
	for _, arg := range node.Arguments {
		arg.Accept(v)
	}
	return ast.VisitorMetaData{}, nil
}

// VisitMathOpNode Visits a math op node
func (v *Visitor) VisitMathOpNode(node *ast.MathOpNode) (ast.VisitorMetaData, error) {
	if visitor, ok := v.visitor.(ast.MathOpNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitMathOpNode(node)
	}
	v.ShouldCallVisitor = true
	node.LeftSide.Accept(v)
	node.RightSide.Accept(v)
	return ast.VisitorMetaData{}, nil
}

// VisitReturnNode visits a return node
func (v *Visitor) VisitReturnNode(node *ast.ReturnNode) (ast.VisitorMetaData, error) {
	if visitor, ok := v.visitor.(ast.ReturnNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitReturnNode(node)
	}
	v.ShouldCallVisitor = true
	node.Expression.Accept(v)
	return ast.VisitorMetaData{}, nil
}

// VisitVarName visits a varname node
func (v *Visitor) VisitVarName(node *ast.VarName) (ast.VisitorMetaData, error) {
	if visitor, ok := v.visitor.(ast.VarNameVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitVarName(node)
	}
	v.ShouldCallVisitor = true
	node.Type.Accept(v)
	return ast.VisitorMetaData{}, nil
}

// VisitDeclNode visits the decl Node
func (v *Visitor) VisitDeclNode(node *ast.DeclNode) (ast.VisitorMetaData, error) {
	if visitor, ok := v.visitor.(ast.DeclNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitDeclNode(node)
	}
	v.ShouldCallVisitor = true
	node.Varname.Accept(v)
	return ast.VisitorMetaData{}, nil
}

// VisitPipeNode visits the pipe node
func (v *Visitor) VisitPipeNode(node *ast.PipeNode) (ast.VisitorMetaData, error) {
	if visitor, ok := v.visitor.(ast.PipeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitPipeNode(node)
	}
	v.ShouldCallVisitor = true
	node.LeftSide.Accept(v)
	node.RightSide.Accept(v)
	return ast.VisitorMetaData{}, nil
}

// VisitIfNode visits an if node
func (v *Visitor) VisitIfNode(node *ast.IfNode) (ast.VisitorMetaData, error) {
	if visitor, ok := v.visitor.(ast.IfVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitIfNode(node)
	}
	v.ShouldCallVisitor = true
	node.Condition.Accept(v)
	node.Body.Accept(v)
	for _, elseIf := range node.ElseIfs {
		elseIf.Accept(v)
	}
	if node.Else != nil {
		els := node.Else
		els.Accept(v)
	}
	return ast.VisitorMetaData{}, nil
}

// VisitImportNode visits an import node
func (v *Visitor) VisitImportNode(node *ast.ImportNode) (ast.VisitorMetaData, error) {
	if visitor, ok := v.visitor.(ast.ImportVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitImportNode(node)
	}
	v.ShouldCallVisitor = true
	return ast.VisitorMetaData{}, nil
}

// VisitTypeNode visits a type node
func (v *Visitor) VisitTypeNode(node *ast.TypeNode) (ast.VisitorMetaData, error) {
	if visitor, ok := v.visitor.(ast.TypeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitTypeNode(node)
	}
	v.ShouldCallVisitor = true
	return ast.VisitorMetaData{}, nil
}

// VisitIndexNode visits an index node
func (v *Visitor) VisitIndexNode(node *ast.IndexNode) (ast.VisitorMetaData, error) {
	if visitor, ok := v.visitor.(ast.IndexVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitIndexNode(node)
	}
	v.ShouldCallVisitor = true
	return ast.VisitorMetaData{}, nil
}

// VisitSliceNode visits a slice node
func (v *Visitor) VisitSliceNode(node *ast.SliceNode) (ast.VisitorMetaData, error) {
	if visitor, ok := v.visitor.(ast.SliceVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitSliceNode(node)
	}
	v.ShouldCallVisitor = true
	return ast.VisitorMetaData{}, nil
}
