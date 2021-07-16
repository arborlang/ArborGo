package typevisitor

import (
	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/parser/scope"
)

func (t *typeVisitor) VisitFunctionDefinitionNode(def *ast.FunctionDefinitionNode) (ast.Node, error) {
	t.scope.PushNewScope()
	defer t.scope.PopScope()
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
