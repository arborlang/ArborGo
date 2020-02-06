package verify

import (
	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/parser/scope"
	// "github.com/arborlang/ArborGo/internal/parser/visitors/WatCompiler"
	"github.com/arborlang/ArborGo/internal/parser/visitors/base"
)

// Visitor verifies the correctness of the AST
type Visitor struct {
	level   int
	visitor *base.Visitor
	symbols *scope.SymbolTable
	// symbols compiler.SymbolTable
}

// SetVisitor sets the visitor
func (v *Visitor) SetVisitor(b *base.Visitor) {
	v.visitor = b
}

// New Returns a new Visitor
func New() *base.Visitor {
	return base.New(&Visitor{
		symbols: scope.NewTable(),
		level:   0,
	})
}

// NoTypeDeclError means there is no type on the decl node
type NoTypeDeclError struct{}

func (n NoTypeDeclError) Error() string {
	return "no type defined on declaration"
}

// VisitDeclNode verify the decl node
func (v *Visitor) VisitDeclNode(node *ast.DeclNode) (ast.VisitorMetaData, error) {
	if node.Varname.Type == nil {
		return ast.VisitorMetaData{}, NoTypeDeclError{}
	}
	return ast.VisitorMetaData{}, nil
}
