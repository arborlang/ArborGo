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
	v     *base.VisitorAdapter
	scope *scope.SymbolTable
}

func New() ast.Visitor {
	tVisitor := &typeVisitor{}
	tVisitor.scope = scope.NewTable()
	b := base.New(tVisitor)
	tVisitor.scope.AddType("String", &scope.TypeData{
		IsSealed: true,
		Type: &types.ConstantTypeNode{
			Name: "String",
		},
	})
	tVisitor.scope.AddType("Boolean", &scope.TypeData{
		IsSealed: true,
		Type: &types.ConstantTypeNode{
			Name: "Boolean",
		},
	})
	tVisitor.scope.AddType("Integer", &scope.TypeData{
		IsSealed: true,
		Type: &types.ConstantTypeNode{
			Name: "Integer",
		},
	})
	tVisitor.scope.AddType("Number", &scope.TypeData{
		IsSealed: true,
		Type: &types.ConstantTypeNode{
			Name: "Number",
		},
	})
	tVisitor.scope.AddType("Float", &scope.TypeData{
		IsSealed: true,
		Type: &types.ConstantTypeNode{
			Name: "Float",
		},
	})
	tVisitor.scope.AddType("Char", &scope.TypeData{
		IsSealed: true,
		Type: &types.ConstantTypeNode{
			Name: "Char",
		},
	})
	return b
}

func (t *typeVisitor) SetVisitor(v *base.VisitorAdapter) {
	t.v = v
}
