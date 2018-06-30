package visitor

import (
	"io"

	"github.com/radding/ArborGo/internal/ast"
)

//CompilerVisitor visits the AST and outputs textual IR
type CompilerVisitor struct {
	writer io.Writer
}

//NewCompilerVisitor just returns a CompilerVisitor
func NewCompilerVisitor(w io.Writer) *CompilerVisitor {
	return &CompilerVisitor{
		writer: w,
	}
}

//PreVisit implements the Visitor
func (w *CompilerVisitor) PreVisit(node ast.Node, global, local *ast.CompilerContext) ast.Register {
	return ""
}

//PostVisit implements the Visitor Node
func (w *CompilerVisitor) PostVisit(node ast.Node, global, local *ast.CompilerContext) ast.Register {
	switch node.(type) {
	case *ast.Program:
		return w.VisitProgram(node, global, local)
	}
	return ""
}

//VisitProgram visits the Program
func (w *CompilerVisitor) VisitProgram(node ast.Node, global, local *ast.CompilerContext) ast.Register {
	return ""
}
