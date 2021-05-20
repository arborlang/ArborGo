package typevisitor

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
)

func (t *typeVisitor) VisitMethodDefinition(node *ast.MethodDefinition) (ast.Node, error) {
	tp, _ := t.scope.LookupType(node.TypeName.Name)
	if tp == nil {
		return nil, fmt.Errorf("type %s is not defined here: %s", node.TypeName.Name, node.FuncDef.Lexeme)
	}
	if tp.IsSealed {
		return nil, fmt.Errorf("type %s is sealed here: %s", node.TypeName.Name, node.FuncDef.Lexeme)
	}
	return node, nil
}
