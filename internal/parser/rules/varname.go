package rules

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

//varNameRule parses a varname
func varNameRule(shouldEnforceTypes bool, p *Parser) (ast.Node, error) {
	node := p.Next()
	if node.Token != tokens.VARNAME {
		return nil, fmt.Errorf("expected a name, got %s", node)
	}
	varname := &ast.VarName{}
	varname.Name = node.Value
	tok := p.Peek()
	didSeeColon := false
	if tok.Token == tokens.COLON {
		didSeeColon = true
		p.Next()
		tp, err := typeRules(p)
		if err != nil {
			return nil, err
		}
		varname.Type = tp.(*ast.TypeNode)
	}
	switch p.Peek().Token {
	case tokens.EQUAL:
		return assignmentOperator(varname, p)
	case tokens.RPAREN:
		return functionCallRule(varname, p)
	case tokens.ARTHOP:
		return MathOpRule(varname, p)
	case tokens.BOOLEAN:
		return boolOperation(varname, p)
	case tokens.COMPARISON:
		return comparisonRule(varname, p)
	case tokens.LSQUARE:
		p.Next() // Fuck consitency ammiright?
		return indexRule(varname, p)
	case tokens.PIPE:
		return pipeRule(varname, p)
	}
	if didSeeColon == false && shouldEnforceTypes {
		return nil, fmt.Errorf("Expected ':', got %s instead", tok)
	}
	return varname, nil
}
