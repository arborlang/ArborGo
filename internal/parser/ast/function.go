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
	Type       types.TypeNode
	Lexeme     lexer.Lexeme
}

// Accept visits the node
func (f *FunctionDefinitionNode) Accept(v Visitor) (Node, error) {
	return v.VisitFunctionDefinitionNode(f)
}

// GetType returns the function type
func (f *FunctionDefinitionNode) GetFnType() (*types.FnType, error) {
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

func (f *FunctionDefinitionNode) GetType() types.TypeNode {
	tp, _ := f.GetFnType()
	return tp
}

// Accept visits the node
func (f *FunctionCallNode) Accept(v Visitor) (Node, error) {
	return v.VisitFunctionCallNode(f)
}

func (f *FunctionCallNode) GetType() types.TypeNode {
	return f.Type
}
