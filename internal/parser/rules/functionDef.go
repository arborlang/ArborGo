package rules

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func functionDefinitionRule(p *Parser) (ast.Node, error) {
	funcDefNode := &ast.FunctionDefinitionNode{}
	topLexeme := p.Next()
	lexeme := topLexeme
	if lexeme.Token != tokens.RPAREN {
		return nil, fmt.Errorf("expected '(', got %s instead", lexeme)
	}
	for lexeme := p.Peek(); lexeme.Token != tokens.LPAREN; lexeme = p.Peek() {
		arg, err := varNameRule(true, p)
		if err != nil {
			return nil, err
		}
		varname, ok := arg.(*ast.VarName)
		if !ok {
			return nil, fmt.Errorf("expected Varname, got %s instead", arg)
		}
		funcDefNode.Arguments = append(funcDefNode.Arguments, varname)
		if peek := p.Peek(); peek.Token != tokens.COMMA && peek.Token != tokens.LPAREN {
			return nil, fmt.Errorf("Unexpected token %s", lexeme)
		} else if peek.Token == tokens.COMMA {
			p.Next()
		}
	}
	p.Next()
	if next := p.Next(); next.Token != tokens.COLON {
		return nil, fmt.Errorf("expected ':' got %s instead", next)
	}
	node, err := typeRules(p)
	if err != nil {
		return nil, err
	}
	funcDefNode.Returns = node.(*ast.TypeNode)
	if next := p.Next(); next.Token != tokens.ARROW {
		return nil, fmt.Errorf("expected '->', got %s instead", next)
	}
	var body ast.Node
	if next := p.Peek(); next.Token == tokens.RCURLY {
		p.Next()
		var err error
		body, err = ProgramRule(p, tokens.LCURLY)
		if err != nil {
			return nil, err
		}
		bodyNode := body.(*ast.Program)
		nodes := []ast.Node{}
		found := false
		for _, node := range bodyNode.Nodes {
			_, ok := node.(*ast.ReturnNode)
			nodes = append(nodes, node)
			if ok {
				found = true
				break
			}
		}
		bodyNode.Nodes = nodes
		if !found {
			return nil, fmt.Errorf("expected a return for all paths on line: %d", topLexeme.Line)
		}
		if err != nil {
			return nil, err
		}
	} else {
		var err error
		body, err = ExpressionRule(p)
		if err != nil {
			return nil, err
		}
	}

	funcDefNode.Body = body
	return funcDefNode, nil
}
