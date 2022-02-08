package rulesv2

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
	"github.com/arborlang/ArborGo/internal/tokens"
	// "github.com/arborlang/ArborGo/internal/tokens"
)

// ConstantsRule is a rule that corresponds with a number or a string
func ConstantsRule(p *Parser, canDoComparisons bool) (ast.Node, error) {
	lexeme := p.Next()
	constant := &ast.Constant{}
	constant.Value = lexeme.Value
	constant.Lexeme = lexeme
	tp := ""
	switch lexeme.Token {
	case tokens.FLOAT:
		tp = "Float"
	case tokens.NUMBER:
		tp = "Number"
	case tokens.STRINGVAL:
		tp = "String"
	case tokens.CHARVAL:
		tp = "Char"
	default:
		return nil, fmt.Errorf("unexpected type %s", lexeme)
	}
	constant.Type = &types.ConstantTypeNode{Name: tp}
	constant.Raw = lexeme.RuneVal
	switch p.Peek().Token {
	case tokens.ARTHOP:
		return MathOpRule(constant, p)
	case tokens.BOOLEAN:
		if !canDoComparisons {
			return nil, fmt.Errorf("unexpected character %q", p.Peek().Token)
		}
		return boolOperation(constant, p)
	case tokens.COMPARISON:
		return comparisonRule(constant, p)
	case tokens.PIPE:
		return pipeRule(constant, p)
	case tokens.LSQUARE:
		if tp != "String" {
			return nil, UnexpectedError(p.Peek())
		}
		return indexRule(constant, p)
	}
	return constant, nil
}

// func stringSliceRule(constant *ast.Constant, p *Parser) (ast.Node, error) {
// 	return

// }
