package rulesv2

import (
	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func dotVarNameRule(p *Parser) (ast.Node, error) {
	node := p.Next()
	if node.Token != tokens.VARNAME {
		return nil, UnexpectedError(node, "Variable Name")
	}
	varname := &ast.VarName{}
	varname.Name = node.Value
	switch p.Peek().Token {
	case tokens.RPAREN:
		return functionCallRule(varname, p, false)
	case tokens.LSQUARE:
		return indexRule(varname, p)
	case tokens.DOT:
		return dotRule(varname, p)
	default:
		return varname, nil
	}

}

func dotRule(varName ast.Node, p *Parser) (ast.Node, error) {
	dot := p.Next()
	if dot.Token != tokens.DOT {
		return nil, UnexpectedError(dot, ".")
	}
	varNameTok := p.Peek()
	if varNameTok.Token != tokens.VARNAME {
		return nil, UnexpectedError(varNameTok, "VARNAME")
	}
	nextNode, err := dotVarNameRule(p)
	if err != nil {
		return nil, err
	}
	node := &ast.DotNode{
		VarName: varName,
		Access:  nextNode,
	}
	switch p.Peek().Token {
	case tokens.EQUAL:
		return assignmentOperator(node, p)
	case tokens.RPAREN:
		return functionCallRule(node, p, false)
	case tokens.ARTHOP:
		return MathOpRule(node, p)
	case tokens.BOOLEAN:
		return boolOperation(node, p)
	case tokens.COMPARISON:
		return comparisonRule(node, p)
	case tokens.LSQUARE:
		return indexRule(node, p)
	case tokens.PIPE:
		return pipeRule(node, p)
	case tokens.DOT:
		return dotRule(node, p)
	}
	return node, nil
}
