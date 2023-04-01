/*
*
package typevisitor transforms the tree into a tree with better type checking
*/
package typevisitor

import (
	"fmt"

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

func GetScope(vistor ast.Visitor) (*scope.SymbolTable, error) {
	v, ok := vistor.(*base.VisitorAdapter)
	if !ok {
		return nil, fmt.Errorf("Can't get Symbol Table")
	}
	return v.Visitor.GetSymbolTable(), nil
}

func addAllBase(scopetable *scope.SymbolTable, bases ...string) {
	for _, i := range bases {
		scopetable.AddToScope(i, &scope.SymbolData{
			Type: scope.TypeData{
				IsSealed: true,
				Type: &types.ConstantTypeNode{
					Name:   i,
					IsBase: true,
				},
			},
			IsConstant: true,
			IsType:     true,
			Location:   "noop",
		})

	}
}

func (t *typeVisitor) GetSymbolTable() *scope.SymbolTable {
	return t.scope
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
