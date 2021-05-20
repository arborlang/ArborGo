package rulesv2

import (
	// "fmt"

	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func varNameRule(p *Parser, isDef bool) (ast.Node, error) {
	node := p.Next()
	if node.Token != tokens.VARNAME && node.Token != tokens.SELF {
		return nil, UnexpectedError(node, "Variable Name")
	}
	varname := &ast.VarName{
		Lexeme: node,
	}
	varname.Name = node.Value

	tok := p.Peek()
	if tok.Token == tokens.COLON && !p.isInSlice {
		if !isDef {
			return nil, fmt.Errorf("unexpected %s", tok)
		}
		p.Next()
		tp, err := typeRule(p)
		if err != nil {
			return nil, err
		}
		varname.Type = tp
	}
	switch p.Peek().Token {
	case tokens.EQUAL:
		return assignmentOperator(varname, p)
	case tokens.RPAREN:
		return functionCallRule(varname, p, true)
	case tokens.ARTHOP:
		return MathOpRule(varname, p)
	case tokens.BOOLEAN:
		return boolOperation(varname, p)
	case tokens.COMPARISON:
		return comparisonRule(varname, p)
	case tokens.LSQUARE:
		return indexRule(varname, p)
	case tokens.PIPE:
		return pipeRule(varname, p)
	case tokens.DOT:
		return dotRule(varname, p)
	}
	return varname, nil
}
