package rules

import (
	"fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
	"github.com/radding/ArborGo/internal/tokens"
)

// gaurdParserRule gaurds a parser function from an unexpected keyworld after they run into the rules and pops the semicolon off if it encounters it.
func gaurdParserRule(p *Parser, rule func(*Parser) (ast.Node, error)) (ast.Node, error) {
	node, err := rule(p)
	if err != nil {
		return nil, err
	}
	if nxt := p.Next(); nxt.Token != tokens.SEMI {
		return nil, fmt.Errorf("unexpected token: got %s exected ';'", nxt)
	}
	return node, nil
}

// ExpressionRule parsers an expression. Expressions are any statement that has a return value you need to handle.
func ExpressionRule(p *Parser) (ast.Node, error) {
	next := p.Peek()
	switch next.Token {
	case tokens.NUMBER, tokens.STRINGVAL, tokens.CHARVAL, tokens.FLOAT:
		return ConstantsRule(p)
	case tokens.VARNAME:
		return varNameRule(false, p)
	case tokens.RPAREN:
		return functionDefinitionRule(p)
	default:
		return nil, fmt.Errorf("encountered unexpected token %s", next)
	}
}

// StatementRule parses a statement. Statement nodes defined as either a node that returns nothing or an expression followed by a semicolon.
func StatementRule(p *Parser) (ast.Node, error) {
	for next := p.Peek(); next.Token != tokens.SEMI; next = p.Peek() {
		switch next.Token {
		case tokens.LET, tokens.CONST:
			return gaurdParserRule(p, DeclRule)
		case tokens.RETURN:
			return gaurdParserRule(p, returnRule)
		case tokens.IF:
			return ifNodeRule(p)
		default:
			return gaurdParserRule(p, ExpressionRule)
		}

	}
	return nil, fmt.Errorf("encountered unknown token: %s", p.Peek())
}

// ProgramRule is basicaly the entry point for all Arbor files. Returns the root node
func ProgramRule(p *Parser, until tokens.Token) (ast.Node, error) {
	program := &ast.Program{}
	for next := p.Next(); next.Token != until; next = p.Next() {
		if next.Token == tokens.NEWLINE {
			continue
		}
		p.Backup()
		node, err := StatementRule(p)
		if err != nil {
			return nil, err
		}
		// program.AddChild(node)
		program.Nodes = append(program.Nodes, node)
	}
	return program, nil
}
