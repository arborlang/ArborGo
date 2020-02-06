package wast

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
)

// VisitAssignment visits an assignment node
func (c *Compiler) VisitAssignment(assignment *ast.AssignmentNode) (ast.VisitorMetaData, error) {
	if _, ok := assignment.Value.(*ast.FunctionDefinitionNode); ok {
		return visitFunctionDefinitionNode(c, assignment)
	}
	location, err := assignment.AssignTo.Accept(c)
	if err != nil {
		return ast.VisitorMetaData{}, err
	}
	if location.SymbolData == nil {
		return ast.VisitorMetaData{}, fmt.Errorf("Didn't get symbol data from the assign to")
	}
	if err != nil {
		return ast.VisitorMetaData{}, err
	}
	sym := c.SymbolTable.GetSymbol(location.SymbolData.Name)
	c.currentAssignment = location.SymbolData.Name
	defer func() { c.currentAssignment = "" }()

	result, err := assignment.Value.Accept(c)
	if err != nil {
		return ast.VisitorMetaData{}, err
	}
	if location.SymbolData.Type != nil && !location.SymbolData.Type.IsValidType(result.Types.Types[0]) {
		return ast.VisitorMetaData{}, fmt.Errorf("can't assign %s to %s", result.Types, location.Types)
	}
	if location.SymbolData.IsNew {
		sym.Location = c.getUniqueID(location.Types.Types[0], location.SymbolData.Name)
		c.AddLocal(sym.Location, c.getType(result.Types.Types[0]))
		// c.locals = append(c.locals, locals{sym.Location, c.getType(result.Types)})
		location.Location = sym.Location
		if sym.Type == "" {
			sym.Type = result.Types.Types[0]
		}
		// c.SymbolTable.AddToScope(sym)
	}
	// _, isConstant := assignment.Value.(*ast.Constant)
	// if c.IsTopScope() && location.SymbolData.IsConstant && isConstant {
	// }
	if result.Location == "STACK" { // If the result is stored on the stack, emit the store command
		c.EmitFunc("set_local %s", location.Location)
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
	c.currentAssignment = location.SymbolData.Name
	if sym == nil && !isDeclNode {
		return ast.VisitorMetaData{}, fmt.Errorf("symbol %s not defined", location.Location)
	}
	result, err := assignment.Value.Accept(c)
	if err != nil {
		return ast.VisitorMetaData{}, err
	}
	if sym == nil {
		tp := location.Types
		if location.Types.Types[0] == "" {
			tp = result.Types
		}
		sym = &Symbol{
			Name:       location.Location,
			Location:   result.Location,
			Type:       tp.Types[0],
			IsConstant: declNode.IsConstant,
		}
		c.SymbolTable.AddToScope(sym)
	} else if sym.Type != result.Types.Types[0] {
		return ast.VisitorMetaData{}, fmt.Errorf("can not assign type %s to %s: %s", result.Types, sym.Name, sym.Type)
	} else if sym.IsConstant && !location.SymbolData.IsNew {
		return ast.VisitorMetaData{}, fmt.Errorf("reassigning constant symbol %s", sym.Name)
	}
	sym.Location = result.Location
	// _, isFunc := assignment.Value.(*ast.FunctionDefinitionNode)
	if c.IsTopScope() {
		c.currentFunc.export = true
		c.currentFunc.name = sym.Name
		c.currentFunc.mangle = sym.Location
	}
	return ast.VisitorMetaData{}, nil

}
