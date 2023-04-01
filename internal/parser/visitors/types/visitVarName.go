package typevisitor

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
)

type NotDefinedError struct {
	varName *ast.VarName
}

func (n *NotDefinedError) Error() string {
	return fmt.Sprintf("%q is not defined: %s", n.varName.Name, n.varName.Lexeme)
}

func (t *typeVisitor) VisitVarName(n *ast.VarName) (ast.Node, error) {
	fmt.Println("in varname", n.Lexeme)
	if n.Type == nil {
		info, _ := t.scope.LookupSymbol(n.Name)
		if info == nil {
			return nil, &NotDefinedError{
				varName: n,
			}
		}
		n.Type = info.Type.Type
	} else {
		tp := t.expandTypeObject(n.Type)
		if tp == nil {
			return nil, errorF("%s not defined", n.Type)
		}
		// err := t.verifyType(n.Type, n.Lexeme)
		n.Type = tp
	}
	return n, nil
}
