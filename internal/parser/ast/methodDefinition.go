package ast

// MethodDefinition Represents a method definition on a function
type MethodDefinition struct {
	FuncDef    *FunctionDefinitionNode
	TypeName   *VarName
	MethodName *VarName
}

func (m *MethodDefinition) Accept(v Visitor) (Node, error) {
	return v.VisitMethodDefinition(m)
}
