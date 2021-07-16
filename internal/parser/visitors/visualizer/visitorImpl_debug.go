package visualizer

import (
	"fmt"
	"strings"

	"github.com/arborlang/ArborGo/internal/parser/ast"
)

func (v *Visualizer) VisitAnyNode(n ast.Node) (ast.Node, error) {
	return n, nil
}

func (v *Visualizer) VisitExtendsNode(extends *ast.ExtendsNode) (ast.Node, error) {
	t := v.appendName(fmt.Sprintf("%s extends %s", extends.Name.Name, extends.Extend.Name))
	v.makeNewLevel(t)
	v.appendName(extends.Adds.String())
	v.popLevel()
	return extends, nil
}

func (v *Visualizer) VisitAnnotatedNode(node *ast.AnnotatedNode) (ast.Node, error) {
	v.appendName(fmt.Sprintf("@%s", node.VarName))
	return node, nil
}

// VisitProgram visits a compiler block
func (v *Visualizer) VisitProgram(block *ast.Program) (ast.Node, error) {
	t := v.appendName("{...}")
	// TODO: Add support for children here
	v.makeNewLevel(t)
	defer v.popLevel()
	for _, stmt := range block.Nodes {
		_, err := stmt.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return block, nil
}

func (v *Visualizer) VisitImplementsNode(implementsNode *ast.ImplementsNode) (ast.Node, error) {
	v.appendName(fmt.Sprintf("%s implements %s", implementsNode.Lexeme, implementsNode.GetType()))
	return implementsNode, nil
}

// VisitAssignment visits an assignment node
func (v *Visualizer) VisitAssignmentNode(assignment *ast.AssignmentNode) (ast.Node, error) {
	t := v.appendName("=")
	assignment.AssignTo.Accept(v)
	// TODO: Add support for children here
	v.makeNewLevel(t)
	defer v.popLevel()
	assignment.Value.Accept(v)
	return assignment, nil
}

// VisitBoolOp visits a boolean node
func (v *Visualizer) VisitBoolOp(node *ast.BoolOp) (ast.Node, error) {
	t := v.appendName(node.Condition)
	v.makeNewLevel(t)
	defer v.popLevel()
	// TODO: Add support for children here
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
func (v *Visualizer) VisitComparison(node *ast.Comparison) (ast.Node, error) {
	// TODO: Add support for children here
	t := v.appendName(node.Operation)
	v.makeNewLevel(t)
	defer v.popLevel()
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
func (v *Visualizer) VisitConstant(node *ast.Constant) (ast.Node, error) {
	v.appendName(node.Value)
	return node, nil
}

// VisitFunctionDefinitionNode visits a function definition ndde
func (v *Visualizer) VisitFunctionDefinitionNode(node *ast.FunctionDefinitionNode) (ast.Node, error) {
	args := []string{}
	for _, arg := range node.Arguments {
		args = append(args, arg.Name, fmt.Sprintf("%s | %s", arg.Name, arg.Type))
	}
	t := v.appendName(fmt.Sprintf("fn %s(%s) -> %s", node.Lexeme, strings.Join(args, ", "), node.Returns))
	v.makeNewLevel(t)
	defer v.popLevel()
	body, err := node.Body.Accept(v)
	node.Body = body
	return node, err
}

// VisitFunctionCallNode visits a function call node
func (v *Visualizer) VisitFunctionCallNode(node *ast.FunctionCallNode) (ast.Node, error) {
	t := v.appendName(fmt.Sprintf("call %s", node.Definition))
	//TODO: Support children
	v.makeNewLevel(t)
	defer v.popLevel()
	for _, arg := range node.Arguments {
		_, err := arg.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return node, nil
}

// VisitMathOpNode Visits a math op node
func (v *Visualizer) VisitMathOpNode(node *ast.MathOpNode) (ast.Node, error) {
	t := v.appendName(node.Operation)
	v.makeNewLevel(t)
	defer v.popLevel()
	// TODO: Support children
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
func (v *Visualizer) VisitReturnNode(node *ast.ReturnNode) (ast.Node, error) {
	t := v.appendName("return")
	// TODO: support children
	v.makeNewLevel(t)
	defer v.popLevel()
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
func (v *Visualizer) VisitVarName(node *ast.VarName) (ast.Node, error) {
	v.appendName(node.Name)
	return node, nil
}

// VisitDeclNode visits the decl Node
func (v *Visualizer) VisitDeclNode(node *ast.DeclNode) (ast.Node, error) {
	str := "let"
	if node.IsConstant {
		str = "const"
	}
	v.appendName(fmt.Sprintf("%s %s", str, node.Varname.Name))
	return node, nil
}

// VisitPipeNode visits the pipe node
func (v *Visualizer) VisitPipeNode(node *ast.PipeNode) (ast.Node, error) {
	t := v.appendName("|>")
	// TODO: Support children
	v.makeNewLevel(t)
	defer v.popLevel()
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
func (v *Visualizer) VisitIfNode(node *ast.IfNode) (ast.Node, error) {
	t := v.appendName("if")
	v.makeNewLevel(t)
	defer v.popLevel()
	condition, err := node.Condition.Accept(v)
	if err != nil {
		return nil, err
	}
	t = v.appendName("")
	v.makeNewLevel(t)
	body, err := node.Body.Accept(v)
	if err != nil {
		return nil, err
	}
	v.popLevel()
	elIfs := []*ast.IfNode{}
	for _, elseIf := range node.ElseIfs {
		t = v.appendName("else")
		v.makeNewLevel(t)
		elIf, err := elseIf.Accept(v)
		if err != nil {
			return nil, err
		}
		elIfs = append(elIfs, elIf.(*ast.IfNode))
		v.popLevel()
	}
	els := node.Else
	if els != nil {
		t = v.appendName("else")
		v.makeNewLevel(t)
		els, err = els.Accept(v)
		if err != nil {
			return nil, err
		}
		v.popLevel()
	}
	node.Condition = condition
	node.Body = body
	node.ElseIfs = elIfs
	node.Else = els
	return node, nil
}

// VisitImportNode visits an import node
func (v *Visualizer) VisitImportNode(node *ast.ImportNode) (ast.Node, error) {
	v.appendName(fmt.Sprintf("import %s from %s", node.ImportAs, node.Source))
	return node, nil
}

// VisitTypeNode visits a type node
func (v *Visualizer) VisitTypeNode(node *ast.TypeNode) (ast.Node, error) {
	v.appendName(fmt.Sprintf("type %s", node.Types))
	return node, nil
}

// VisitIndexNode visits an index node
func (v *Visualizer) VisitIndexNode(node *ast.IndexNode) (ast.Node, error) {
	t := v.appendName("index")
	v.makeNewLevel(t)
	defer v.popLevel()
	t = v.appendName("value to index")
	v.makeNewLevel(t)
	node.Varname.Accept(v)
	v.popLevel()
	t = v.appendName("index value")
	v.makeNewLevel(t)
	node.Index.Accept(v)
	v.popLevel()
	return node, nil
}

// VisitSliceNode visits a slice node
func (v *Visualizer) VisitSliceNode(node *ast.SliceNode) (ast.Node, error) {
	t := v.appendName("index")
	v.makeNewLevel(t)
	defer v.popLevel()
	t = v.appendName("value to index")
	v.makeNewLevel(t)
	node.Varname.Accept(v)
	v.popLevel()
	t = v.appendName("start value")
	v.makeNewLevel(t)
	node.Start.Accept(v)
	v.popLevel()
	t = v.appendName("end value")
	v.makeNewLevel(t)
	node.End.Accept(v)
	v.popLevel()
	if node.Step != nil {
		t = v.appendName("step value")
		v.makeNewLevel(t)
		node.Step.Accept(v)
		v.popLevel()
	}
	return node, nil
}

// VisitSliceNode visits a slice node
func (v *Visualizer) VisitDecoratorNode(node *ast.DecoratorNode) (ast.Node, error) {
	t := v.appendName("@%s", node.Name.Name)
	v.makeNewLevel(t)
	defer v.popLevel()
	node.Decorates.Accept(v)
	return node, nil
}

func (v *Visualizer) VisitDotNode(node *ast.DotNode) (ast.Node, error) {
	t := v.appendName(".")
	v.makeNewLevel(t)
	v.popLevel()
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

func (v *Visualizer) VisitInstantiateNode(node *ast.InstantiateNode) (ast.Node, error) {
	t := v.appendName("new")
	v.makeNewLevel(t)
	defer v.popLevel()
	callNode, err := node.FunctionCallNode.Accept(v)
	node.FunctionCallNode = callNode.(*ast.FunctionCallNode)
	return node, err
}

func (v *Visualizer) VisitInternalNode(node *ast.InternalNode) (ast.Node, error) {
	t := v.appendName("internal")
	v.makeNewLevel(t)
	defer v.popLevel()
	expr, err := node.Expression.Accept(v)
	node.Expression = expr
	return node, err
}

func (v *Visualizer) VisitMatchNode(node *ast.MatchNode) (ast.Node, error) {
	t := v.appendName("match")
	v.makeNewLevel(t)
	defer v.popLevel()
	node.Match.Accept(v)
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

func (v *Visualizer) VisitMethodDefinition(node *ast.MethodDefinition) (ast.Node, error) {
	t := v.appendName("%s.%s", node.TypeName.Name, node.MethodName.Name)
	v.makeNewLevel(t)
	defer v.popLevel()
	_, err := node.FuncDef.Accept(v)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (v *Visualizer) VisitPackage(node *ast.Package) (ast.Node, error) {

	return node, nil
}

func (v *Visualizer) VisitShapeNode(node *ast.ShapeNode) (ast.Node, error) {
	t := v.appendName("shape")
	v.makeNewLevel(t)
	defer v.popLevel()
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

func (v *Visualizer) VisitWhenNode(node *ast.WhenNode) (ast.Node, error) {
	t := v.appendName("when")
	v.makeNewLevel(t)
	defer v.popLevel()
	// TODO: children
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

func (v *Visualizer) VisitContinueNode(node *ast.ContinueNode) (ast.Node, error) {
	t := v.appendName("continue")
	v.makeNewLevel(t)
	defer v.popLevel()
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

func (v *Visualizer) VisitSignalNode(node *ast.SignalNode) (ast.Node, error) {
	t := v.appendName("signal %s", node.Level)
	v.makeNewLevel(t)
	defer v.popLevel()
	val, err := node.ValueToRaise.Accept(v)
	node.ValueToRaise = val
	return node, err
}

func (v *Visualizer) VisitTryNode(node *ast.TryNode) (ast.Node, error) {
	t := v.appendName("try")
	v.makeNewLevel(t)
	defer v.popLevel()
	t = v.appendName("main body")
	v.makeNewLevel(t)
	node.Tries.Accept(v)
	v.popLevel()
	for _, hNode := range node.HandleCases {
		hNode.Accept(v)
	}
	return node, nil
}

func (v *Visualizer) VisitHandleCaseNode(node *ast.HandleCaseNode) (ast.Node, error) {

	t := v.appendName("handle(%s %s)", node.VariableName, node.Type)
	v.makeNewLevel(t)
	defer v.popLevel()
	node.Case.Accept(v)
	return node, nil
}
