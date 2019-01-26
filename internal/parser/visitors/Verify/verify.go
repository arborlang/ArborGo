package verify

import (
	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/parser/visitors/WatCompiler"
	"github.com/arborlang/ArborGo/internal/parser/visitors/base"
)

// Visitor verifies the AST
type Visitor struct {
	level   int
	visitor *base.Visitor
	symbols compiler.SymbolTable
}

// SetVisitor sets the visitor
func (v *Visitor) SetVisitor(b *base.Visitor) {
	v.visitor = b
}

// New Returns a new Visitor
func New() *base.Visitor {
	return base.New(&Visitor{
		level: 0,
	})
}

// NoTypeDeclError means there is no type on the decl node
type NoTypeDeclError struct{}

func (n NoTypeDeclError) Error() string {
	return "no type defined on declatation"
}

// VisitDeclNode verify the decl node
func (v *Visitor) VisitDeclNode(node *ast.DeclNode) (ast.VisitorMetaData, error) {
	if node.Varname.Type == nil {
		return ast.VisitorMetaData{}, NoTypeDeclError{}
	}
	return ast.VisitorMetaData{}, nil
}

// VisitAssignment visits an assignment node
// func (v *Visitor) VisitAssignment(node *ast.AssignmentNode) (ast.VisitorMetaData, error) {

// }
