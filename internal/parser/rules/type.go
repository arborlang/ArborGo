package rules

import (
	"fmt"
	"github.com/radding/ArborGo/internal/lexer"
	"github.com/radding/ArborGo/internal/parser/ast"
	"github.com/radding/ArborGo/internal/tokens"
	"strings"
)

func isAType(tok tokens.Token) bool {
	return tok == tokens.VARNAME || tok == tokens.STRING || tok == tokens.CHAR || tok == tokens.NUMBERWORD || tok == tokens.FLOATWORD || tok == tokens.RPAREN
}

func isAllowable(tok tokens.Token) bool {
	return tok == tokens.COMMA || tok == tokens.SEMI || tok == tokens.ARROW
}

func isPipe(lexeme lexer.Lexeme) bool {
	return lexeme.Token == tokens.LOGICAL && lexeme.Value == "|"
}

// parseFunctionDefinitionType is a function that takes what looks like a function type definition and distiles it dowwn to a simpler format.
// For example (number, string, someType) -> something | nil gets turned into func(number,string,someType)something|nil
func parseFunctionDefinitionType(p *Parser) (string, error) {
	// Get all parameter types
	tp := []string{"func("}
	internalTps := []string{}
	for next := p.Peek(); next.Token != tokens.LPAREN; next = p.Peek() {
		node, err := typeRules(p)
		if err != nil {
			return "", err
		}
		typeNode := node.(*ast.TypeNode)
		internalTps = append(internalTps, strings.Join(typeNode.Types, "|"))
		if nxt := p.Peek(); nxt.Token != tokens.COMMA && nxt.Token != tokens.LPAREN {
			return "", fmt.Errorf("expected ',' or ')' got %s instead", nxt)
		} else if nxt.Token == tokens.COMMA {
			p.Next()
		}
	}
	tp = append(tp, strings.Join(internalTps, ","))
	tp = append(tp, ")")
	p.Next() // This should be the closing paren

	// Get the return type
	nxt := p.Next()
	// There should always be a return type
	if nxt.Token != tokens.ARROW {
		return "", fmt.Errorf("expected symbol '->', got %s instead", nxt)
	}
	node, err := typeRules(p)
	if err != nil {
		return "", err
	}
	returnNode := node.(*ast.TypeNode)
	tp = append(tp, strings.Join(returnNode.Types, "|"))
	return strings.Join(tp, ""), nil
}

func typeRules(p *Parser) (ast.Node, error) {
	tp := &ast.TypeNode{}
	nxt := p.Next()
	if !isAType(nxt.Token) {
		return nil, fmt.Errorf("expected a type, got %s instead", nxt)
	}
	tpStr := nxt.Value
	if nxt.Token == tokens.RPAREN {
		var err error
		tpStr, err = parseFunctionDefinitionType(p)
		if err != nil {
			return nil, err
		}
	}
	tp.Types = append(tp.Types, tpStr)

	nxt = p.Peek()
	if isPipe(nxt) {
		for nxt := p.Peek(); isPipe(nxt); nxt = p.Peek() {
			p.Next()        // Skip over the pipe
			nxt := p.Next() // get the type name
			if !isAType(nxt.Token) {
				return nil, fmt.Errorf("expected a type, got %s instead", nxt)
			}
			tpStr := nxt.Value
			if nxt.Token == tokens.RPAREN {
				var err error
				tpStr, err = parseFunctionDefinitionType(p)
				if err != nil {
					return nil, err
				}
			}
			tp.Types = append(tp.Types, tpStr)
		}
		if next := p.Peek(); !isAllowable(next.Token) {
			return nil, fmt.Errorf("unexpected token while parsing types: %s", next)
		}
	}
	return tp, nil
}
