package typevisitor

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
)

func (t *typeVisitor) VisitVarName(n *ast.VarName) (ast.Node, error) {
	if n.Type == nil {
		info, _ := t.scope.LookupSymbol(n.Name)
		if info == nil {
			return nil, fmt.Errorf("%q is not defined: %s", n.Name, n.Lexeme)
		}
		n.Type = info.Type.Type
	} else {
		err := t.verifyType(n.Type, n.Lexeme)
		if err != nil {
			return nil, err
		}
	}
	return n, nil
}
