package rulesv2

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func functionCallRule(varname ast.Node, p *Parser, recurse bool) (ast.Node, error) {
	tok := p.Next()
	if tok.Token != tokens.RPAREN {
		return nil, fmt.Errorf("expected '(', got %s", tok)
	} // Discard the first paren
	funcCall := &ast.FunctionCallNode{
		Lexeme: tok,
	}
	for next := p.Peek(); next.Token != tokens.LPAREN; next = p.Peek() {
		argument, err := ExpressionRule(p)
		if err != nil {
			return nil, err
		}
		funcCall.Arguments = append(funcCall.Arguments, argument)
		if nxt := p.Peek(); nxt.Token == tokens.COMMA {
			p.Next()
		}
	}
	// if p.Peek().Token == tokens.LPAREN {
	p.Next()
	// }
	funcCall.Definition = varname
	if recurse {
		switch p.Peek().Token {
		case tokens.BOOLEAN:
			return boolOperation(funcCall, p)
		case tokens.COMPARISON:
			return comparisonRule(funcCall, p)
		case tokens.ARTHOP:
			return MathOpRule(funcCall, p)
		case tokens.PIPE:
			return pipeRule(funcCall, p)
		}
	}
	return funcCall, nil
}
