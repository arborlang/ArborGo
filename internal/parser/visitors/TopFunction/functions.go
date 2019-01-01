package functions

import (
	"fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
	"github.com/radding/ArborGo/internal/parser/visitors/base"
)

// Argument defines the arguments for a function
type Argument struct {
	Name  string   // Name is the name of the argument
	Types []string // Type is the type of argument
}

// Symbol is the function definition its self
type Symbol struct {
	Name       string     // Is the name of the function
	Arguments  []Argument // the list of arguments the function takes
	ReturnType string     // what the function returns
	IsConstant bool
}

// FunctionAnalyzer gets all of the top level function definitions for a given module.
type FunctionAnalyzer struct {
	visitor   *base.Visitor
	ShouldGet bool
	functions map[string]Symbol
	Called    int
}

// New returns a base visitor
func New() *base.Visitor {
	return base.New(&FunctionAnalyzer{
		ShouldGet: false,
		functions: make(map[string]Symbol),
		Called:    0,
	})
}

// SetVisitor sets the visitor
func (f *FunctionAnalyzer) SetVisitor(v *base.Visitor) {
	f.visitor = v
}

// VisitBlock visits a block node
func (f *FunctionAnalyzer) VisitBlock(block *ast.Program) (ast.VisitorMetaData, error) {
	f.Called++
	if f.Called > 1 {
		return ast.VisitorMetaData{}, nil
	}
	f.visitor.ShouldCallVisitor = false
	return f.visitor.VisitBlock(block)
}

// VisitAssignment visits an Assignment Node
func (f *FunctionAnalyzer) VisitAssignment(block *ast.AssignmentNode) (ast.VisitorMetaData, error) {
	var ok bool
	var funcNode *ast.FunctionDefinitionNode
	if funcNode, ok = block.Value.(*ast.FunctionDefinitionNode); !ok {
		// ignore anything that is not a function definition.
		return ast.VisitorMetaData{}, nil
	}
	args := []Argument{}
	for _, arg := range funcNode.Arguments {
		args = append(args, Argument{
			Name:  arg.Name,
			Types: arg.Type.Types,
		})
	}
	switch assignTo := block.AssignTo.(type) {
	case *ast.DeclNode:
		_, ok := f.functions[assignTo.Varname.Name]
		if ok {
			return ast.VisitorMetaData{}, fmt.Errorf("%s is being redefined", assignTo.Varname.Name)
		}

		f.functions[assignTo.Varname.Name] = Symbol{
			Name:       assignTo.Varname.Name,
			IsConstant: assignTo.IsConstant,
			Arguments:  args,
		}
		return ast.VisitorMetaData{}, nil
	case *ast.VarName:
		varname, ok := f.functions[assignTo.Name]
		if !ok {
			return ast.VisitorMetaData{}, fmt.Errorf("%s not defined", assignTo.Name)
		}
		if varname.IsConstant {
			return ast.VisitorMetaData{}, fmt.Errorf("%s is being redefined", assignTo.Name)
		}
		f.functions[assignTo.Name] = Symbol{
			Name:       assignTo.Name,
			IsConstant: false,
			Arguments:  args,
		}
		return ast.VisitorMetaData{}, nil
	default:
		return ast.VisitorMetaData{}, fmt.Errorf("What the fuck are you assigning to?")
	}
}
