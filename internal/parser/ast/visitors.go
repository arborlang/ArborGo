package ast

/**
* This file defines the interfaces needed to walk the AST. Each  Interface defines what the Node can be spoken to
 */

// AssignmentVisitor visists an Assignment Node
type AssignmentVisitor interface {
	VisitAssignment(*AssignmentNode) (VisitorMetaData, error)
}

// BoolOpVisitor visits a BoolOP node
type BoolOpVisitor interface {
	VisitBoolOp(*BoolOp) (VisitorMetaData, error)
}

// ComparisonVisitor visits a Comparison Node
type ComparisonVisitor interface {
	VisitComparison(*Comparison) (VisitorMetaData, error)
}

// ConstantVisitor visits a ConstantNode
type ConstantVisitor interface {
	VisitConstant(*Constant) (VisitorMetaData, error)
}

// FunctionDefinitionNodeVisitor visits a FunctionDefinitionNode
type FunctionDefinitionNodeVisitor interface {
	VisitFunctionDefinitionNode(*FunctionDefinitionNode) (VisitorMetaData, error)
}

// FunctionCallNodeVisitor visits a FunctionCallNode
type FunctionCallNodeVisitor interface {
	VisitFunctionCallNode(*FunctionCallNode) (VisitorMetaData, error)
}

// MathOpNodeVisitor visits a MathOpNode
type MathOpNodeVisitor interface {
	VisitMathOpNode(*MathOpNode) (VisitorMetaData, error)
}

// ReturnNodeVisitor visits a return node
type ReturnNodeVisitor interface {
	VisitReturnNode(*ReturnNode) (VisitorMetaData, error)
}

// VarNameVisitor Visits a vistorNode
type VarNameVisitor interface {
	VisitVarName(*VarName) (VisitorMetaData, error)
}

// DeclNodeVisitor visits a DeclNode
type DeclNodeVisitor interface {
	VisitDeclNode(*DeclNode) (VisitorMetaData, error)
}

// BlockVisitor visits a Program Node (aka A block)
type BlockVisitor interface {
	VisitBlock(*Program) (VisitorMetaData, error)
}

// PipeVisitor visits a Pipe node
type PipeVisitor interface {
	VisitPipeNode(*PipeNode) (VisitorMetaData, error)
}

// IfVisitor Visits an if statement
type IfVisitor interface {
	VisitIfNode(*IfNode) (VisitorMetaData, error)
}

//ImportVisitor visits the import node
type ImportVisitor interface {
	VisitImportNode(*ImportNode) (VisitorMetaData, error)
}

// TypeVisitor vistis a type node
type TypeVisitor interface {
	VisitTypeNode(*TypeNode) (VisitorMetaData, error)
}

// IndexVisitor is the index visitor of the
type IndexVisitor interface {
	VisitIndexNode(*IndexNode) (VisitorMetaData, error)
}

//SliceVisitor visits a slice node
type SliceVisitor interface {
	VisitSliceNode(*SliceNode) (VisitorMetaData, error)
}

// Visitor can visit every type of node. The Visitor will be responsible for walking into the children
type Visitor interface {
	AssignmentVisitor
	BoolOpVisitor
	VarNameVisitor
	ComparisonVisitor
	ConstantVisitor
	FunctionDefinitionNodeVisitor
	FunctionCallNodeVisitor
	MathOpNodeVisitor
	ReturnNodeVisitor
	DeclNodeVisitor
	BlockVisitor
	PipeVisitor
	IfVisitor
	ImportVisitor
	TypeVisitor
	IndexVisitor
	SliceVisitor
}
