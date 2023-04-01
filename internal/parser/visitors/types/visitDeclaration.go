package typevisitor

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/parser/scope"
)

func (t *typeVisitor) VisitDeclNode(n *ast.DeclNode) (ast.Node, error) {
	other, _ := t.scope.LookupSymbol(n.Varname.Name)
	if other != nil {
		return nil, fmt.Errorf("%s is being redefined here: %s", n.Varname.Name, n.Varname.Lexeme)
	}
	varName, err := n.Varname.Accept(t.v)
	if _, ok := err.(*NotDefinedError); !ok && err != nil {
		return nil, err
	}
	if varName != nil {
		vName, _ := varName.(*ast.VarName)
		n.Varname = vName
	}
	t.scope.AddToScope(n.Varname.Name, &scope.SymbolData{
		Type: scope.TypeData{
			Type:     n.Varname.GetType(),
			IsSealed: false,
		},
		Lexeme:     n.Varname.Lexeme,
		Location:   n.Varname.Name,
		IsConstant: n.IsConstant,
	})
	return n, nil
}
