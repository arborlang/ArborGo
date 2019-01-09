package wast

import (
	"fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
)

// VisitAssignment visits an assignment node
func (c *Compiler) VisitAssignment(assignment *ast.AssignmentNode) (ast.VisitorMetaData, error) {
	if _, ok := assignment.Value.(*ast.FunctionDefinitionNode); ok {
		return visitFunctionDefinitionNode(c, assignment)
	}
	location, err := assignment.AssignTo.Accept(c)
	if location.SymbolData == nil {
		return ast.VisitorMetaData{}, fmt.Errorf("Didn't get symbol data from the assign to")
	}
	if err != nil {
		return ast.VisitorMetaData{}, err
	}
	sym := c.SymbolTable.GetSymbol(location.SymbolData.Name)
	result, err := assignment.Value.Accept(c)
	if err != nil {
		return ast.VisitorMetaData{}, err
	}
	if location.SymbolData.Type != nil && !location.SymbolData.Type.IsValidType(result.Types) {
		return ast.VisitorMetaData{}, fmt.Errorf("can't assign %s to %s", result.Types, location.Types)
	}
	if location.SymbolData.IsNew {
		sym.Location = c.getUniqueID(location.Types, location.SymbolData.Name)
		c.locals = append(c.locals, locals{sym.Location, c.getType(result.Types)})
		location.Location = sym.Location
		if sym.Type == "" {
			sym.Type = result.Types
		}
		// c.SymbolTable.AddToScope(sym)
		// c.Emit("ß(local %s i64)", sym.Location)
	}
	// _, isConstant := assignment.Value.(*ast.Constant)
	// if c.IsTopScope() && location.SymbolData.IsConstant && isConstant {
	// 	c.Emit("(export ")
	// }
	if result.Location == "STACK" { // If the result is stored on the stack, emit the store command
		// c.Emit("(%s %s)", cmd, location.Location)
		c.Emit("set_local %s", location.Location)
		return ast.VisitorMetaData{}, nil
	}
	sym.Location = result.Location // If the result is not on the stack, why load it and then store it? just change the location
	return ast.VisitorMetaData{}, nil
}

func visitFunctionDefinitionNode(c *Compiler, assignment *ast.AssignmentNode) (ast.VisitorMetaData, error) {
	location, err := assignment.AssignTo.Accept(c)
	declNode, isDeclNode := assignment.AssignTo.(*ast.DeclNode)
	if err != nil {
		return ast.VisitorMetaData{}, err
	}
	sym := c.SymbolTable.GetSymbol(location.SymbolData.Name)
	if sym == nil && !isDeclNode {
		return ast.VisitorMetaData{}, fmt.Errorf("symbol %s not defined", location.Location)
	}
	result, err := assignment.Value.Accept(c)
	if err != nil {
		return ast.VisitorMetaData{}, err
	}
	if sym == nil {
		tp := location.Types
		if location.Types == "" {
			tp = result.Types
		}
		sym = &Symbol{
			Name:       location.Location,
			Location:   result.Location,
			Type:       tp,
			IsConstant: declNode.IsConstant,
		}
		c.SymbolTable.AddToScope(sym)
	} else if sym.Type != result.Types {
		return ast.VisitorMetaData{}, fmt.Errorf("can not assign type %s to %s: %s", result.Types, sym.Name, sym.Type)
	} else if sym.IsConstant && !location.SymbolData.IsNew {
		return ast.VisitorMetaData{}, fmt.Errorf("reassigning constant symbol %s", sym.Name)
	}
	sym.Location = result.Location
	// _, isFunc := assignment.Value.(*ast.FunctionDefinitionNode)
	if c.IsTopScope() {
		c.Emit("(export \"%s\" %s)", sym.Name, result.Exportable)
	}
	return ast.VisitorMetaData{}, nil

}
