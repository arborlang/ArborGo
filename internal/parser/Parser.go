package parser

import (
	"github.com/radding/ArborGo/internal/lexer"
	"github.com/radding/ArborGo/internal/parser/ast"
	"github.com/radding/ArborGo/internal/parser/rules"
	"github.com/radding/ArborGo/internal/tokens"
)

// Parse takes a lexer and parsers it to its AST
func Parse(lex func() lexer.Lexeme) (*ast.Program, error) {
	parserStream := rules.New(lex)
	prog, err := rules.ProgramRule(parserStream, tokens.EOF)
	if err != nil {
		return nil, err
	}
	return prog.(*ast.Program), nil
}
