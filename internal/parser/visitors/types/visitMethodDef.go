package typevisitor

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

func (t *typeVisitor) VisitMethodDefinition(node *ast.MethodDefinition) (ast.Node, error) {
	tp, _ := t.scope.LookupSymbol(node.TypeName.Name)
	if tp == nil {
		return nil, fmt.Errorf("type %s is not defined here: %s", node.TypeName.Name, node.FuncDef.Lexeme)
	}
	if tp.Type.IsSealed {
		return nil, fmt.Errorf("type %s is sealed here: %s", node.TypeName.Name, node.FuncDef.Lexeme)
	}
	funcDef, err := node.FuncDef.Accept(t.v)
	if err != nil {
		return nil, err
	}
	node.FuncDef = funcDef.(*ast.FunctionDefinitionNode)
	toAdd, err := node.FuncDef.GetFnType()
	if err != nil {
		return nil, err
	}
	switch value := tp.Type.Type.(type) {
	case *types.ShapeType:
		value.Fields[node.MethodName.Name] = toAdd
	case *types.ExtendedType:
		value.Shape.Fields[node.MethodName.Name] = toAdd
	default:
		return nil, fmt.Errorf("Can't add %s to %s because it is not an object", node.MethodName.Lexeme, node.TypeName.Name)
	}
	return node, nil
}
