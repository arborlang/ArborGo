package rulesv2

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

//DeclRule defines how to parse a rule that begins with decleration (`const` or `let`)
func DeclRule(p *Parser) (ast.Node, error) {
	d := &ast.DeclNode{}
	tp := p.Next()
	name := p.Peek()
	if name.Token != tokens.VARNAME {
		return nil, fmt.Errorf("expected token VARNAME, got %s: %s on line %d:%d", name.Token.String(), name.Value, name.Line, name.Column)
	}
	nameNode, err := varNameRule(p, true)
	if err != nil {
		return nil, err
	}

	switch node := nameNode.(type) {
	case *ast.AssignmentNode:
		if tp.Token == tokens.CONST {
			d.IsConstant = true
		}
		// vName, ok := node.AssignTo.(*ast.VarName)
		// if ok == false {
		// 	return nil, fmt.Errorf("got bad node back in assign to: expected a varname go", node.AssignTo)
		// }
		// d.Varname = vName
		err = d.AddChild(node.AssignTo)
		if err != nil {
			return nil, err
		}
		// if d.Varname.Type == nil {
		// 	d.Var
		// }
		node.AssignTo = d

		return node, nil
	case *ast.VarName:
		if tp.Token == tokens.CONST {
			d.IsConstant = true
			if node.Type == nil {
				return nil, fmt.Errorf("can not define a const with out a type")
			}
		}
		d.AddChild(nameNode)
		// if len(node.Type.Types) == 0 {
		// 	return nil, fmt.Errorf("ambiguous type declaration for variable %q Line %d:%d", node.Name, node.Lexeme.Line, node.Lexeme.Column)
		// }
		return d, nil
	}

	return nil, fmt.Errorf("got bad node back from parser")
}
