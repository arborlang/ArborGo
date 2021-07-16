package typevisitor

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

func (t *typeVisitor) VisitFunctionCallNode(n *ast.FunctionCallNode) (ast.Node, error) {
	varname := n.Definition.(*ast.VarName)

	info, _ := t.scope.LookupSymbol(varname.Name)
	if info == nil {
		return nil, fmt.Errorf("%s is not defined at %s", varname.Name, varname.Lexeme)
	}
	fnType, ok := info.Type.Type.(*types.FnType)
	if !ok {
		return nil, fmt.Errorf("Type %s is not callable %s", info.Type.Type, varname.Lexeme)
	}
	n.Type = fnType.ReturnVal
	params := []ast.Node{}
	tps := []types.TypeNode{}
	for _, arg := range n.Arguments {
		param, err := arg.Accept(t.v)
		if err != nil {
			return nil, err
		}
		params = append(params, param)
		tps = append(tps, param.GetType())
	}
	derivedFnTp := &types.FnType{
		ReturnVal:  n.Type,
		Parameters: tps,
	}
	if !fnType.IsSatisfiedBy(derivedFnTp) {
		return nil, fmt.Errorf("%s can not be called, signatures don't match. %s vs %s", varname.Lexeme, derivedFnTp, fnType)
	}
	return n, nil
}
