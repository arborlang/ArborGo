package base

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
)

// VisitorHider is a simple way to set and hide the visitor
type VisitorHider interface {
	SetVisitor(v *VisitorAdapter)
}

// VisitorAdapter represents a top level VisitorAdapter that walks the tree but does nothing. Useful for doing analysis on the AST by other visitors
// For example, if I want to collect all of the function definitions of an arbor file, I would define a struct that is composed of this
// visitor and implements the VisitFunctionDefinitionNode function.
type VisitorAdapter struct {
	Visitor           VisitorHider // visitor implements a visitor interface
	ShouldCallVisitor bool
}

// New returns a new Visitor
func New(visitor VisitorHider) *VisitorAdapter {
	v := &VisitorAdapter{
		Visitor:           visitor,
		ShouldCallVisitor: true,
	}
	visitor.SetVisitor(v)
	return v
}

// GetVisitor gets the underlying visitor
func (v *VisitorAdapter) GetVisitor() interface{} {
	return v.Visitor
}

func (v *VisitorAdapter) VisitAnyNode(node ast.Node) (ast.Node, error) {
	if visitor, ok := v.Visitor.(ast.GenericVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitAnyNode(node)
	}
	return nil, nil
}

func (v *VisitorAdapter) VisitExtendsNode(extends *ast.ExtendsNode) (ast.Node, error) {
	v.VisitAnyNode(extends)
	if visitor, ok := v.Visitor.(ast.ExtendsNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitExtendsNode(extends)
	}
	return extends, nil
}

func (v *VisitorAdapter) VisitAnnotatedNode(node *ast.AnnotatedNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if vi, ok := v.Visitor.(ast.AnnotatedNodeVisitor); ok {

		return vi.VisitAnnotatedNode(node)
	}
	return node, nil
}

// VisitProgram visits a compiler block
func (v *VisitorAdapter) VisitProgram(block *ast.Program) (ast.Node, error) {
	v.VisitAnyNode(block)
	node, err := v.VisitAnyNode(block)
	if node != nil || err != nil {
		return node, err
	}
	if visitor, ok := v.Visitor.(ast.ProgramVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitProgram(block)
	}
	v.ShouldCallVisitor = true
	statements := []ast.Node{}
	for _, stmt := range block.Nodes {
		stmt, err := stmt.Accept(v)
		if err != nil {
			return nil, err
		}
		statements = append(statements, stmt)
	}
	block.Nodes = statements
	return block, nil
}

func (v *VisitorAdapter) VisitImplementsNode(implementsNode *ast.ImplementsNode) (ast.Node, error) {
	v.VisitAnyNode(implementsNode)
	if visitor, ok := v.Visitor.(ast.ImplementsNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitImplementsNode(implementsNode)
	}
	return implementsNode, nil
}

// VisitAssignment visits an assignment node
func (v *VisitorAdapter) VisitAssignmentNode(assignment *ast.AssignmentNode) (ast.Node, error) {
	v.VisitAnyNode(assignment)
	if visitor, ok := v.Visitor.(ast.AssignmentNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitAssignmentNode(assignment)
	}
	v.ShouldCallVisitor = true
	assignTo, err := assignment.AssignTo.Accept(v)
	if err != nil {
		return assignment, err
	}
	value, err := assignment.Value.Accept(v)
	if err != nil {
		return assignment, err
	}
	assignment.AssignTo = assignTo
	assignment.Value = value
	return assignment, nil
}

// VisitBoolOp visits a boolean node
func (v *VisitorAdapter) VisitBoolOp(node *ast.BoolOp) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.BoolOpVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitBoolOp(node)
	}
	v.ShouldCallVisitor = true
	leftOp, err := node.LeftSide.Accept(v)
	if err != nil {
		return nil, err
	}
	rightOP, err := node.RightSide.Accept(v)
	if err != nil {
		return nil, err
	}
	node.LeftSide = leftOp
	node.RightSide = rightOP
	return node, nil
}

// VisitComparison Visits a comparison node
func (v *VisitorAdapter) VisitComparison(node *ast.Comparison) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.ComparisonVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitComparison(node)
	}
	v.ShouldCallVisitor = true
	comparison, err := node.LeftSide.Accept(v)
	if err != nil {
		return nil, err
	}
	right, err := node.RightSide.Accept(v)
	if err != nil {
		return nil, err
	}
	node.LeftSide = comparison
	node.RightSide = right
	return node, nil
}

// VisitConstant visits the constant object
func (v *VisitorAdapter) VisitConstant(node *ast.Constant) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.ConstantVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitConstant(node)
	}
	v.ShouldCallVisitor = true
	return node, nil
}

// VisitFunctionDefinitionNode visits a function definition ndde
func (v *VisitorAdapter) VisitFunctionDefinitionNode(node *ast.FunctionDefinitionNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.FunctionDefinitionNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitFunctionDefinitionNode(node)
	}
	v.ShouldCallVisitor = true
	args := []*ast.VarName{}
	for _, arg := range node.Arguments {
		argRes, err := arg.Accept(v)
		if err != nil {
			return nil, err
		}
		args = append(args, argRes.(*ast.VarName))
	}
	body, err := node.Body.Accept(v)
	node.Arguments = args
	node.Body = body
	return node, err
}

// VisitFunctionCallNode visits a function call node
func (v *VisitorAdapter) VisitFunctionCallNode(node *ast.FunctionCallNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.FunctionCallNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitFunctionCallNode(node)
	}
	v.ShouldCallVisitor = true
	def, err := node.Definition.Accept(v)
	if err != nil {
		return nil, err
	}
	node.Definition = def
	args := []ast.Node{}
	for _, arg := range node.Arguments {
		arg, err := arg.Accept(v)
		if err != nil {
			return nil, err
		}
		args = append(args, arg)
	}
	node.Arguments = args
	return node, nil
}

// VisitMathOpNode Visits a math op node
func (v *VisitorAdapter) VisitMathOpNode(node *ast.MathOpNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.MathOpNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitMathOpNode(node)
	}
	v.ShouldCallVisitor = true
	left, err := node.LeftSide.Accept(v)
	if err != nil {
		return nil, err
	}
	right, err := node.RightSide.Accept(v)
	if err != nil {
		return nil, err
	}
	node.LeftSide = left
	node.RightSide = right
	return node, nil
}

// VisitReturnNode visits a return node
func (v *VisitorAdapter) VisitReturnNode(node *ast.ReturnNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.ReturnNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitReturnNode(node)
	}
	v.ShouldCallVisitor = true
	if node.Expression != nil {
		expr, err := node.Expression.Accept(v)
		if err != nil {
			return nil, err
		}
		node.Expression = expr
	}

	return node, nil
}

// VisitVarName visits a varname node
func (v *VisitorAdapter) VisitVarName(node *ast.VarName) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.VarNameVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitVarName(node)
	}
	v.ShouldCallVisitor = true
	return node, nil
}

// VisitDeclNode visits the decl Node
func (v *VisitorAdapter) VisitDeclNode(node *ast.DeclNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.DeclNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitDeclNode(node)
	}
	v.ShouldCallVisitor = true
	vName, err := node.Varname.Accept(v)
	node.Varname = vName.(*ast.VarName)
	return node, err
}

// VisitPipeNode visits the pipe node
func (v *VisitorAdapter) VisitPipeNode(node *ast.PipeNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.PipeNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitPipeNode(node)
	}
	v.ShouldCallVisitor = true
	left, err := node.LeftSide.Accept(v)
	if err != nil {
		return nil, err
	}
	right, err := node.RightSide.Accept(v)
	if err != nil {
		return nil, err
	}
	node.LeftSide = left
	node.RightSide = right
	return node, nil
}

// VisitIfNode visits an if node
func (v *VisitorAdapter) VisitIfNode(node *ast.IfNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.IfNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitIfNode(node)
	}
	v.ShouldCallVisitor = true
	condition, err := node.Condition.Accept(v)
	if err != nil {
		return nil, err
	}
	body, err := node.Body.Accept(v)
	if err != nil {
		return nil, err
	}
	elIfs := []*ast.IfNode{}
	for _, elseIf := range node.ElseIfs {
		elIf, err := elseIf.Accept(v)
		if err != nil {
			return nil, err
		}
		elIfs = append(elIfs, elIf.(*ast.IfNode))
	}
	els := node.Else
	if els != nil {
		els, err = els.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	node.Condition = condition
	node.Body = body
	node.ElseIfs = elIfs
	node.Else = els
	return node, nil
}

// VisitImportNode visits an import node
func (v *VisitorAdapter) VisitImportNode(node *ast.ImportNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.ImportNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitImportNode(node)
	}
	v.ShouldCallVisitor = true
	return node, nil
}

// VisitTypeNode visits a type node
func (v *VisitorAdapter) VisitTypeNode(node *ast.TypeNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.TypeNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitTypeNode(node)
	}
	v.ShouldCallVisitor = true
	return node, nil
}

// VisitIndexNode visits an index node
func (v *VisitorAdapter) VisitIndexNode(node *ast.IndexNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.IndexNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitIndexNode(node)
	}
	v.ShouldCallVisitor = true
	return node, nil
}

// VisitSliceNode visits a slice node
func (v *VisitorAdapter) VisitSliceNode(node *ast.SliceNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.SliceNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitSliceNode(node)
	}
	v.ShouldCallVisitor = true
	return node, nil
}

// VisitSliceNode visits a slice node
func (v *VisitorAdapter) VisitDecoratorNode(node *ast.DecoratorNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.DecoratorNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitDecoratorNode(node)
	}
	v.ShouldCallVisitor = true
	name, err := node.Name.Accept(v)
	if err != nil {
		return nil, err
	}
	node.Name = name.(*ast.VarName)
	decorates, err := node.Decorates.Accept(v)
	if err != nil {
		return nil, err
	}
	node.Decorates = decorates
	return node, nil
}

func (v *VisitorAdapter) VisitDotNode(node *ast.DotNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.DotNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitDotNode(node)
	}
	v.ShouldCallVisitor = true
	varName, err := node.VarName.Accept(v)
	if err != nil {
		return nil, err
	}
	node.VarName = varName
	access, err := node.Access.Accept(v)
	if err != nil {
		return nil, err
	}
	node.Access = access
	return node, nil

}

func (v *VisitorAdapter) VisitInstantiateNode(node *ast.InstantiateNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.InstantiateNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitInstantiateNode(node)
	}
	v.ShouldCallVisitor = true
	callNode, err := node.FunctionCallNode.Accept(v)
	node.FunctionCallNode = callNode.(*ast.FunctionCallNode)
	return node, err
}

func (v *VisitorAdapter) VisitInternalNode(node *ast.InternalNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.InternalNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitInternalNode(node)
	}
	v.ShouldCallVisitor = true
	expr, err := node.Expression.Accept(v)
	node.Expression = expr
	return node, err
}

func (v *VisitorAdapter) VisitMatchNode(node *ast.MatchNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.MatchNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitMatchNode(node)
	}
	v.ShouldCallVisitor = true
	match, err := node.Match.Accept(v)
	if err != nil {
		return nil, err
	}
	node.Match = match
	whens := []*ast.WhenNode{}
	for _, clause := range node.WhenCases {
		when, err := clause.Accept(v)
		if err != nil {
			return nil, err
		}
		whens = append(whens, when.(*ast.WhenNode))
	}
	node.WhenCases = whens
	return node, nil
}

func (v *VisitorAdapter) VisitMethodDefinition(node *ast.MethodDefinition) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.MethodDefinitionVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitMethodDefinition(node)
	}
	v.ShouldCallVisitor = true
	def, err := node.FuncDef.Accept(v)
	if err != nil {
		return nil, err
	}
	node.FuncDef = def.(*ast.FunctionDefinitionNode)
	name, err := node.MethodName.Accept(v)
	if err != nil {
		return nil, err
	}
	node.MethodName = name.(*ast.VarName)
	tpName, err := node.TypeName.Accept(v)
	if err != nil {
		return nil, err
	}
	node.TypeName = tpName.(*ast.VarName)
	return node, nil
}

func (v *VisitorAdapter) VisitPackage(node *ast.Package) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.PackageVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitPackage(node)
	}
	v.ShouldCallVisitor = true
	return node, nil
}

func (v *VisitorAdapter) VisitShapeNode(node *ast.ShapeNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.ShapeNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitShapeNode(node)
	}
	v.ShouldCallVisitor = true
	fields := map[string]ast.Node{}
	for name, field := range node.Fields {
		field, err := field.Accept(v)
		if err != nil {
			return nil, err
		}
		fields[name] = field
	}
	node.Fields = fields
	return node, nil
}

func (v *VisitorAdapter) VisitWhenNode(node *ast.WhenNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.WhenNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitWhenNode(node)
	}
	cas, err := node.Case.Accept(v)
	if err != nil {
		return nil, err
	}
	node.Case = cas
	eval, err := node.Evaluate.Accept(v)
	if err != nil {
		return nil, err
	}
	node.Evaluate = eval
	return node, nil
}

func (v *VisitorAdapter) VisitContinueNode(node *ast.ContinueNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.ContinueNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitContinueNode(node)
	}
	nd := node.WithValue
	if nd != nil {
		nd, err := nd.Accept(v)
		if err != nil {
			return nil, err
		}
		node.WithValue = nd
	}
	return node, nil
}

func (v *VisitorAdapter) VisitSignalNode(node *ast.SignalNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.SignalNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitSignalNode(node)
	}
	val, err := node.ValueToRaise.Accept(v)
	node.ValueToRaise = val
	return node, err
}

func (v *VisitorAdapter) VisitTryNode(node *ast.TryNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.TryNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitTryNode(node)
	}
	tries, err := node.Tries.Accept(v)
	if err != nil {
		return nil, err
	}
	node.Tries = tries
	handleCases := []*ast.HandleCaseNode{}
	for _, hNode := range node.HandleCases {
		cs, err := hNode.Accept(v)
		if err != nil {
			return nil, err
		}
		hCs, ok := cs.(*ast.HandleCaseNode)
		if !ok {
			return nil, fmt.Errorf("got back a missunderstood node")
		}
		handleCases = append(handleCases, hCs)
	}
	node.HandleCases = handleCases
	return node, nil
}

func (v *VisitorAdapter) VisitHandleCaseNode(node *ast.HandleCaseNode) (ast.Node, error) {
	v.VisitAnyNode(node)
	if visitor, ok := v.Visitor.(ast.HandleCaseNodeVisitor); ok && v.ShouldCallVisitor {
		return visitor.VisitHandleCaseNode(node)
	}
	cs, err := node.Case.Accept(v)
	if err != nil {
		return nil, err
	}
	node.Case = cs
	return node, nil
}
