package umlvisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

func (u *umlVisitor) VisitSliceNode(slice *ast.SliceNode) (ast.Node, error) {
	label := u.getLabel("slice")

	u.writeLine("object \"slice\" as %s", label)
	maybe, _ := slice.Varname.Accept(u.v)
	u.connectNodeWithLabel(label, maybe, "Var Name")

	if slice.Start != nil {
		maybe, _ = slice.Start.Accept(u.v)
		u.connectNodeWithLabel(label, maybe, "Start")

	}

	if slice.End != nil {
		maybe, _ = slice.End.Accept(u.v)
		u.connectNodeWithLabel(label, maybe, "End")
	}

	if slice.Step != nil {
		maybe, _ = slice.Step.Accept(u.v)
		u.connectNodeWithLabel(label, maybe, "Step")
	}

	return labeledNode(slice, label), nil
}
