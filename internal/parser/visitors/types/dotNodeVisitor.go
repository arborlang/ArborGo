package typevisitor

import (
	"fmt"
	"runtime/debug"

	"github.com/arborlang/ArborGo/internal/parser/ast"
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
	nVarName, err := n.VarName.Accept(t.v)

	n.VarName = nVarName
	return n, err
	// access, ok := n.Access.(*ast.VarName)
	// if !ok {
	// 	return nil, errorF("did not get varname for value in dot node: %s", n.Lexeme)
	// }
	// obj, _ := t.scope.LookupSymbol(access.Name)
	// if obj == nil {
	// 	return nil, errorF("%s not defined", access)
	// }
	// obj = t.scope.ResolveType(obj)
	// shape, ok := obj.Type.Type.(*types.ShapeType)
	// if !ok {
	// 	return nil, errorF("expected an object got %s instead: %s", obj.Type.Type, obj.Lexeme)
	// }
	// name, ok := n.Access.(*ast.VarName)
	// if !ok {
	// 	return nil, errorF("Not a var name")
	// }
	// _, ok = shape.Fields[name.Name]
	// if !ok {
	// 	return nil, fmt.Errorf("%s has no member %s", access.Name, name.Lexeme)
	// }
	// return n, nil
}
