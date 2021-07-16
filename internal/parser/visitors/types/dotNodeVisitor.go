package typevisitor

import (
	"fmt"
	"runtime/debug"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

var printDebug = false

func errorF(msg string, args ...interface{}) error {
	formattedMsg := fmt.Sprintf(msg, args...)
	if !printDebug {
		return fmt.Errorf(formattedMsg)
	}
	stack := debug.Stack()
	return fmt.Errorf(formattedMsg+"\n%s", string(stack))
}

func (t *typeVisitor) VisitDotNode(n *ast.DotNode) (ast.Node, error) {
	access, ok := n.VarName.(*ast.VarName)
	if !ok {
		return nil, errorF("did not get varname for value in dot node")
	}
	obj, _ := t.scope.LookupSymbol(access.Name)
	shape, ok := obj.Type.Type.(*types.ShapeType)
	if !ok {
		return nil, errorF("expected an object got %s instead", obj.Type.Type)
	}
	name, ok := n.Access.(*ast.VarName)
	if !ok {
		return nil, errorF("Not a var name")
	}
	_, ok = shape.Fields[name.Name]
	if !ok {
		return nil, fmt.Errorf("%s has no member %s", access.Name, name.Lexeme)
	}
	return n, nil
}
