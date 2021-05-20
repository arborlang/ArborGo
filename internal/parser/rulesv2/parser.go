/*
Package rulesv2 are the parser rules for arbor
*/
package rulesv2

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

// Parser Represents a parser Object
type Parser struct {
	Lex        func() lexer.Lexeme
	Lexemes    []lexer.Lexeme
	CurrentNdx int
	types      map[string]types.TypeNode
	isInSlice  bool
}

// New builds and returns a new lexer given a Lexer
func New(Lex func() lexer.Lexeme) *Parser {
	tps := make(map[string]types.TypeNode)
	tps["Integer"] = &types.ConstantTypeNode{Name: "Integer"}
	tps["String"] = &types.ConstantTypeNode{Name: "String"}
	tps["Float"] = &types.ConstantTypeNode{Name: "Float"}
	tps["Boolean"] = &types.ConstantTypeNode{Name: "Boolean"}
	return &Parser{
		Lex:        Lex,
		Lexemes:    []lexer.Lexeme{},
		CurrentNdx: -1,
		types:      tps,
	}
}

// PrintNextLexemes prints the next num lexemes
func (p *Parser) PrintNextLexemes(num int) {
	for i := 0; i < num; i++ {
		nxt := p.Next()
		fmt.Print(fmt.Sprintf("%q", nxt.Value), " ")
	}
	fmt.Println("")
	for i := 0; i < num; i++ {
		p.Backup()
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

// LookupType looks up a type and returns a type if it exists
func (p *Parser) LookupType(name string) types.TypeNode {
	if tp, ok := p.types[name]; ok {
		return tp
	}
	return nil
}

// AddType adds a type to the parser
func (p *Parser) AddType(name string, tp types.TypeNode) error {
	if _, ok := p.types[name]; ok {
		return fmt.Errorf("type with name %q already exists", name)
	}
	p.types[name] = tp
	return nil
}

func (p *Parser) setIsInSlice(newVal bool) {
	p.isInSlice = newVal
}
