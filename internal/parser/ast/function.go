package ast

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

// FunctionDefinitionNode represents a function definition
type FunctionDefinitionNode struct {
	Arguments        []*VarName
	Body             Node
	Returns          types.TypeNode
	Lexeme           lexer.Lexeme
	GenericTypeNames []*VarName
}

// FunctionCallNode represents a function call
type FunctionCallNode struct {
	Arguments  []Node
	Definition Node
}

// Accept visits the node
func (f *FunctionDefinitionNode) Accept(v Visitor) (Node, error) {
	return v.VisitFunctionDefinitionNode(f)
}

// GetType returns the function type
func (f *FunctionDefinitionNode) GetType() (*types.FnType, error) {
	paramTypes := []types.TypeNode{}
	for _, arg := range f.Arguments {
		if arg.Type == nil {
			return nil, fmt.Errorf("No type for argument")
		}
		paramTypes = append(paramTypes, arg.Type)
	}
	return &types.FnType{
		Parameters: paramTypes,
		ReturnVal:  f.Returns,
	}, nil
}

// Accept visits the node
func (f *FunctionCallNode) Accept(v Visitor) (Node, error) {
	return v.VisitFunctionCallNode(f)
}
