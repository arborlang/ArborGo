package parser

import (
	"fmt"

	"github.com/radding/ArborGo/internal/ast"
	"github.com/radding/ArborGo/internal/lexer"
	"github.com/radding/ArborGo/internal/tokens"
)

func parseStatements(reader *lexer.BufferedReader) (ast.Node, error) {
	program := &ast.Program{}
	for {
		node, err := parseStatement(reader)
		fmt.Println("Fuck", node)
		if err != nil {
			if _, ok := err.(EOF); !ok {
				return nil, err
			}
			return program, nil
		}
		program.Statemetents = append(program.Statemetents, node)
	}
}

func parseStatement(reader *lexer.BufferedReader) (ast.Node, error) {
	switch lexeme := reader.Peek(); lexeme.Token {
	case tokens.EOF:
		return nil, EOF{}
	case tokens.CONST, tokens.LET:
		return parseDecl(reader)
	}
	return nil, nil
}

func parseDecl(reader *lexer.BufferedReader) (ast.Node, error) {
	isConst := reader.Next().Token == tokens.CONST
	if reader.Peek().Token != tokens.VARNAME {
		return nil, NewUnrecognizedError(reader.Next())
	}
	nameLexeme := reader.Next()
	if reader.Peek().Token != tokens.COLON {
		return nil, NewUnrecognizedError(reader.Next())
	}
	reader.Next()
	if lexeme := reader.Peek(); !tokens.IsType(lexeme.Token) {
		return nil, NewUnrecognizedError(reader.Next())
	}
	typeLexeme := reader.Next()
	return ast.NewDeclaration(nameLexeme, typeLexeme, isConst), nil
}
