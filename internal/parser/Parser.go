package parser

import (
	"github.com/radding/ArborGo/internal/lexer"
	"github.com/radding/ArborGo/internal/parser/ast"
	"github.com/radding/ArborGo/internal/parser/rules"
	"github.com/radding/ArborGo/internal/tokens"
	"io"
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

// Compile compiles an arbor file
func Compile(input io.Reader, visitors ...ast.Visitor) error {
	parserStream := rules.New(lexer.Lex(input))
	prog, err := rules.ProgramRule(parserStream, tokens.EOF)
	if err != nil {
		return err
	}
	for _, vistor := range visitors {
		_, err := prog.Accept(vistor)
		if err != nil {
			return err
		}
	}
	return nil
}
