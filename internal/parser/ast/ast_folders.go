package ast

//Written by the generator, do not over write

type AnnotatedNodeFolder interface {
	FoldAnnotatedNode(n *AnnotatedNode) (Node, error)
}

type AssignmentNodeFolder interface {
	FoldAssignmentNode(n *AssignmentNode) (Node, error)
}

type BoolOpFolder interface {
	FoldBoolOp(n *BoolOp) (Node, error)
}

type ComparisonFolder interface {
	FoldComparison(n *Comparison) (Node, error)
}

type ConstantFolder interface {
	FoldConstant(n *Constant) (Node, error)
}

type DeclNodeFolder interface {
	FoldDeclNode(n *DeclNode) (Node, error)
}

type DecoratorNodeFolder interface {
	FoldDecoratorNode(n *DecoratorNode) (Node, error)
}

type DotNodeFolder interface {
	FoldDotNode(n *DotNode) (Node, error)
}

type FunctionCallNodeFolder interface {
	FoldFunctionCallNode(n *FunctionCallNode) (Node, error)
}

type FunctionDefinitionNodeFolder interface {
	FoldFunctionDefinitionNode(n *FunctionDefinitionNode) (Node, error)
}

type IfNodeFolder interface {
	FoldIfNode(n *IfNode) (Node, error)
}

type ImplementsNodeFolder interface {
	FoldImplementsNode(n *ImplementsNode) (Node, error)
}

type ImportNodeFolder interface {
	FoldImportNode(n *ImportNode) (Node, error)
}

type IndexNodeFolder interface {
	FoldIndexNode(n *IndexNode) (Node, error)
}

type InstantiateNodeFolder interface {
	FoldInstantiateNode(n *InstantiateNode) (Node, error)
}

type InternalNodeFolder interface {
	FoldInternalNode(n *InternalNode) (Node, error)
}

type MatchNodeFolder interface {
	FoldMatchNode(n *MatchNode) (Node, error)
}

type MathOpNodeFolder interface {
	FoldMathOpNode(n *MathOpNode) (Node, error)
}

type MethodDefinitionFolder interface {
	FoldMethodDefinition(n *MethodDefinition) (Node, error)
}

type PackageFolder interface {
	FoldPackage(n *Package) (Node, error)
}

type PipeNodeFolder interface {
	FoldPipeNode(n *PipeNode) (Node, error)
}

type ProgramFolder interface {
	FoldProgram(n *Program) (Node, error)
}

type ReturnNodeFolder interface {
	FoldReturnNode(n *ReturnNode) (Node, error)
}

type ShapeNodeFolder interface {
	FoldShapeNode(n *ShapeNode) (Node, error)
}

type SliceNodeFolder interface {
	FoldSliceNode(n *SliceNode) (Node, error)
}

type TypeNodeFolder interface {
	FoldTypeNode(n *TypeNode) (Node, error)
}

type VarNameFolder interface {
	FoldVarName(n *VarName) (Node, error)
}

type WhenNodeFolder interface {
	FoldWhenNode(n *WhenNode) (Node, error)
}

type Folder interface {
	AnnotatedNodeFolder
	AssignmentNodeFolder
	BoolOpFolder
	ComparisonFolder
	ConstantFolder
	DeclNodeFolder
	DecoratorNodeFolder
	DotNodeFolder
	FunctionCallNodeFolder
	FunctionDefinitionNodeFolder
	IfNodeFolder
	ImplementsNodeFolder
	ImportNodeFolder
	IndexNodeFolder
	InstantiateNodeFolder
	InternalNodeFolder
	MatchNodeFolder
	MathOpNodeFolder
	MethodDefinitionFolder
	PackageFolder
	PipeNodeFolder
	ProgramFolder
	ReturnNodeFolder
	ShapeNodeFolder
	SliceNodeFolder
	TypeNodeFolder
	VarNameFolder
	WhenNodeFolder
}
