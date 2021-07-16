package ast

import "github.com/arborlang/ArborGo/internal/parser/ast/types"

// MethodDefinition Represents a method definition on a function
type MethodDefinition struct {
	FuncDef    *FunctionDefinitionNode
	TypeName   *VarName
	MethodName *VarName
}

func (m *MethodDefinition) Accept(v Visitor) (Node, error) {
	return v.VisitMethodDefinition(m)
}

func (m *MethodDefinition) GetType() types.TypeNode {
	return &types.FalseType{}
}
