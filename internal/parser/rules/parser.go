/*
Package rules are the parser rules for arbor
*/
package rules

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast"
)

// Parser Represents a parser Object
type Parser struct {
	Lex        func() lexer.Lexeme
	Lexemes    []lexer.Lexeme
	CurrentNdx int
}

// New builds and returns a new lexer given a Lexer
func New(Lex func() lexer.Lexeme) *Parser {
	return &Parser{
		Lex:        Lex,
		Lexemes:    []lexer.Lexeme{},
		CurrentNdx: -1,
	}
}

// Next gets the next Lexeme in the stream
func (p *Parser) Next() lexer.Lexeme {
	if len(p.Lexemes) > p.CurrentNdx+1 {
		p.CurrentNdx++
		nxt := p.Lexemes[p.CurrentNdx]
		return nxt
	}
	nxt := p.Lex()
	p.Lexemes = append(p.Lexemes, nxt)
	p.CurrentNdx++
	return nxt
}

// Peek looks at the next lexeme in the stream
func (p *Parser) Peek() lexer.Lexeme {
	if len(p.Lexemes) >= p.CurrentNdx+2 {
		return p.Lexemes[p.CurrentNdx+1]
	}
	nxt := p.Lex()
	p.Lexemes = append(p.Lexemes, nxt)
	return nxt
}

// Previous looks at the last token returned
func (p *Parser) Previous() lexer.Lexeme {
	if p.CurrentNdx <= 0 {
		return lexer.Lexeme{}
	}
	return p.Lexemes[p.CurrentNdx-1]
}

// Backup backs the token stream up one
func (p *Parser) Backup() {
	p.CurrentNdx--
}

// ParseRule is a function that takes a parser and returns the next parse rule, the current Node and an error
type ParseRule func(p *Parser) (ParseRule, ast.Node, error)
