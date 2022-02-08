package rulesv2

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func parseExport(p *Parser) (ast.Node, error) {
	nxt := p.Next()
	if nxt.Token != tokens.INTERNAL {
		return nil, fmt.Errorf("expected %q, got %s", "export", nxt)
	}
	var node ast.Node
	var err error
	nxt = p.Peek()
	switch nxt.Token {
	case tokens.CONST, tokens.LET:
		node, err = DeclRule(p)
	case tokens.VARNAME:
		node, err = varNameRule(p, true)
	case tokens.FUNC:
		node, err = functionDefinitionRule(p)
	case tokens.TYPE:
		node, err = typeDefRule(p)
	default:
		err = fmt.Errorf("expected a function, declaration, or variable, got %s", nxt)
	}
	if err != nil {
		return nil, err
	}
	_, isVarname := node.(*ast.VarName)
	_, isAssignment := node.(*ast.AssignmentNode)
	if !isVarname && !isAssignment {
		return nil, fmt.Errorf("expected a variable name, declaration or function definition on line %d", nxt.Line)
	}
	exp := &ast.InternalNode{
		Lexeme:     nxt,
		Expression: node,
	}
	return exp, nil
}
