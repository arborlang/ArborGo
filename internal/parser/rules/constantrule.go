package rules

import (
	"github.com/radding/ArborGo/internal/parser/ast"
	"github.com/radding/ArborGo/internal/tokens"
)

// ConstantsRule is a rule that corresponds with a number or a string
func ConstantsRule(p *Parser) (ast.Node, error) {
	lexeme := p.Next()
	constant := &ast.Constant{}
	constant.Value = lexeme.Value
	constant.Type = lexeme.Token.String()
	constant.Raw = lexeme.RuneVal
	switch p.Peek().Token {
	case tokens.ARTHOP:
		return MathOpRule(constant, p)
	case tokens.BOOLEAN:
		return boolOperation(constant, p)
	case tokens.COMPARISON:
		return comparisonRule(constant, p)
	case tokens.PIPE:
		return pipeRule(constant, p)
	}
	return constant, nil
}
