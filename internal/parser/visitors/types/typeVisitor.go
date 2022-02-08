/**
package typevisitor transforms the tree into a tree with better type checking
*/
package typevisitor

import (
	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
	"github.com/arborlang/ArborGo/internal/parser/scope"
	"github.com/arborlang/ArborGo/internal/parser/visitors/base"
)

type annotatedTypeNode struct {
	node ast.Node
	tp   types.TypeNode
}

func (a *annotatedTypeNode) Accept(v ast.Visitor) (ast.Node, error) {
	return a.node.Accept(v)
}

type typeVisitor struct {
	v             *base.VisitorAdapter
	scope         *scope.SymbolTable
	dumpOnFailure bool
}

func addAllBase(scopetable *scope.SymbolTable, bases ...string) {
	for _, i := range bases {
		scopetable.AddToScope(i, &scope.SymbolData{
			Type: scope.TypeData{
				IsSealed: true,
				Type: &types.ConstantTypeNode{
					Name: i,
				},
			},
			IsConstant: true,
			Location:   "noop",
		})

	}
}

func New(dumpOnFailure bool) ast.Visitor {
	tVisitor := &typeVisitor{
		dumpOnFailure: dumpOnFailure,
	}
	tVisitor.scope = scope.NewTable()
	b := base.New(tVisitor)
	addAllBase(tVisitor.scope,
		"String",
		"Boolean",
		"Number",
		"Float",
		"Char",
	)
	return b
}

func (t *typeVisitor) SetVisitor(v *base.VisitorAdapter) {
	t.v = v
}
