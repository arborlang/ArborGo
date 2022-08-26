package typevisitor

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/parser/scope"
)

func (t *typeVisitor) VisitFunctionDefinitionNode(def *ast.FunctionDefinitionNode) (ast.Node, error) {
	t.scope.PushNewScope()
	defer t.scope.PopScope()
	for _, arg := range def.Arguments {
		err := t.verifyType(arg.Type, arg.Lexeme)
		if err != nil {
			return def, err
		}
	}
	generics := []ast.VarName{}
	for _, generic := range def.GenericTypeNames {
		generics = append(generics, *generic)
	}
	fmt.Println("Generics", generics)
	if len(def.GenericTypeNames) > 0 {
		fmt.Println("GenericS!")
	}
	if def.Returns != nil {
		fmt.Println("Lets find all return nodes and get the types")
	}
	err := t.verifyType(def.Returns, def.Lexeme)
	if err != nil {
		return def, err
	}
	for _, arg := range def.Arguments {
		t.scope.AddToScope(arg.Name, &scope.SymbolData{
			Type: scope.TypeData{
				Type:     arg.Type,
				IsSealed: true,
			},
			IsConstant: false,
			Lexeme:     arg.Lexeme,
		})
	}
	body, err := def.Body.Accept(t.v)
	if err != nil {
		return nil, err
	}
	def.Body = body
	return def, nil
}
