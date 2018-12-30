package ast

// FunctionDefinitionNode represents a function definition
type FunctionDefinitionNode struct {
	Arguments []*VarName
	Body      Node
}

// FunctionCallNode represents a function call
type FunctionCallNode struct {
	Arguments  []Node
	Definition Node
}

// Accept visits the node
func (f *FunctionDefinitionNode) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitFunctionDefinitionNode(f)
}

// Accept visits the node
func (f *FunctionCallNode) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitFunctionCallNode(f)
}
