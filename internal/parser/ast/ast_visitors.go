package ast

//Written by the generator, do not over write

type AnnotatedNodeVisitor interface {
	VisitAnnotatedNode(n *AnnotatedNode) (Node, error)
}

type AssignmentNodeVisitor interface {
	VisitAssignmentNode(n *AssignmentNode) (Node, error)
}

type BoolOpVisitor interface {
	VisitBoolOp(n *BoolOp) (Node, error)
}

type ComparisonVisitor interface {
	VisitComparison(n *Comparison) (Node, error)
}

type ConstantVisitor interface {
	VisitConstant(n *Constant) (Node, error)
}

type ContinueNodeVisitor interface {
	VisitContinueNode(n *ContinueNode) (Node, error)
}

type DeclNodeVisitor interface {
	VisitDeclNode(n *DeclNode) (Node, error)
}

type DecoratorNodeVisitor interface {
	VisitDecoratorNode(n *DecoratorNode) (Node, error)
}

type DotNodeVisitor interface {
	VisitDotNode(n *DotNode) (Node, error)
}

type ExtendsNodeVisitor interface {
	VisitExtendsNode(n *ExtendsNode) (Node, error)
}

type FunctionCallNodeVisitor interface {
	VisitFunctionCallNode(n *FunctionCallNode) (Node, error)
}

type FunctionDefinitionNodeVisitor interface {
	VisitFunctionDefinitionNode(n *FunctionDefinitionNode) (Node, error)
}

type HandleCaseNodeVisitor interface {
	VisitHandleCaseNode(n *HandleCaseNode) (Node, error)
}

type IfNodeVisitor interface {
	VisitIfNode(n *IfNode) (Node, error)
}

type ImplementsNodeVisitor interface {
	VisitImplementsNode(n *ImplementsNode) (Node, error)
}

type ImportNodeVisitor interface {
	VisitImportNode(n *ImportNode) (Node, error)
}

type IndexNodeVisitor interface {
	VisitIndexNode(n *IndexNode) (Node, error)
}

type InstantiateNodeVisitor interface {
	VisitInstantiateNode(n *InstantiateNode) (Node, error)
}

type InternalNodeVisitor interface {
	VisitInternalNode(n *InternalNode) (Node, error)
}

type MatchNodeVisitor interface {
	VisitMatchNode(n *MatchNode) (Node, error)
}

type MathOpNodeVisitor interface {
	VisitMathOpNode(n *MathOpNode) (Node, error)
}

type MethodDefinitionVisitor interface {
	VisitMethodDefinition(n *MethodDefinition) (Node, error)
}

type PackageVisitor interface {
	VisitPackage(n *Package) (Node, error)
}

type PipeNodeVisitor interface {
	VisitPipeNode(n *PipeNode) (Node, error)
}

type ProgramVisitor interface {
	VisitProgram(n *Program) (Node, error)
}

type ReturnNodeVisitor interface {
	VisitReturnNode(n *ReturnNode) (Node, error)
}

type ShapeNodeVisitor interface {
	VisitShapeNode(n *ShapeNode) (Node, error)
}

type SignalNodeVisitor interface {
	VisitSignalNode(n *SignalNode) (Node, error)
}

type SliceNodeVisitor interface {
	VisitSliceNode(n *SliceNode) (Node, error)
}

type TryNodeVisitor interface {
	VisitTryNode(n *TryNode) (Node, error)
}

type TypeNodeVisitor interface {
	VisitTypeNode(n *TypeNode) (Node, error)
}

type VarNameVisitor interface {
	VisitVarName(n *VarName) (Node, error)
}

type WhenNodeVisitor interface {
	VisitWhenNode(n *WhenNode) (Node, error)
}


type GenericVisitor interface {
	VisitAnyNode(n Node) (Node, error)
}

type Visitor interface {
	AnnotatedNodeVisitor
	AssignmentNodeVisitor
	BoolOpVisitor
	ComparisonVisitor
	ConstantVisitor
	ContinueNodeVisitor
	DeclNodeVisitor
	DecoratorNodeVisitor
	DotNodeVisitor
	ExtendsNodeVisitor
	FunctionCallNodeVisitor
	FunctionDefinitionNodeVisitor
	HandleCaseNodeVisitor
	IfNodeVisitor
	ImplementsNodeVisitor
	ImportNodeVisitor
	IndexNodeVisitor
	InstantiateNodeVisitor
	InternalNodeVisitor
	MatchNodeVisitor
	MathOpNodeVisitor
	MethodDefinitionVisitor
	PackageVisitor
	PipeNodeVisitor
	ProgramVisitor
	ReturnNodeVisitor
	ShapeNodeVisitor
	SignalNodeVisitor
	SliceNodeVisitor
	TryNodeVisitor
	TypeNodeVisitor
	VarNameVisitor
	WhenNodeVisitor

	GenericVisitor
}
