//Package lexer is the lexer for the arbor project. The lexer can either be run asyncrounously or syncrounously
package lexer

import (
	"fmt"
	"io"
	"log"
	"runtime/debug"

	participle "github.com/alecthomas/participle/lexer"
	"github.com/arborlang/ArborGo/internal/tokens"

	"github.com/arborlang/ArborGo/internal/lexer/internal"
	"github.com/arborlang/ArborGo/internal/lexer/state"
)

//Lexeme is the public lexeme for everyone
type Lexeme internal.Lexeme

func (lexeme Lexeme) String() string {
	return fmt.Sprintf("%q (Line: %d, Column: %d)", lexeme.Value, lexeme.Line, lexeme.Column)
}

//NewLexer creates and returns a new lexer instance
func NewLexer(in io.Reader) *internal.Lexer {
	return internal.NewLexer(in)
}

//RunMachine runs the lexer state machine
func RunMachine(lex *internal.Lexer) {
	state.Machine(lex)
}

//LexAsync lexs the input asyncronously. This returns a channel that will contain the identified lexemes
func LexAsync(in io.Reader) chan internal.Lexeme {
	lex := NewLexer(in)
	go RunMachine(lex)
	return lex.Lexemes
}

// Lex lexes the input syncrounously by returning a function that will listen for input on the channel else it will call the state function.
// 	This function returns a function that can be used to get the next lexeme
func Lex(in io.Reader) func() Lexeme {
	lex := NewLexer(in)
	stateFunc := state.LexText
	return func() (l Lexeme) {
		for {
			select {
			case lexeme := <-lex.Lexemes:
				return Lexeme(lexeme)
			default:
				defer func() {
					if r := recover(); r != nil {
						// log.Println("lex recovered ")
						log.Println("recovered in Lex. Something went wrong, outputting EOF and logging stack trace")
						log.Println(string(debug.Stack()))
						l = Lexeme{
							Token:  tokens.EOF,
							Value:  string(tokens.EOFChar),
							Column: -1,
							Line:   -1,
						}
					}
				}()
				stateFunc = stateFunc(lex)
			}
		}
	}
}

type PLexer struct {
	next     func() Lexeme
	fileName string
}

func (p *PLexer) Next() (participle.Token, error) {
	nxt := p.next()
	tok := participle.Token{}
	tok.Type = rune(nxt.Token)
	tok.Value = nxt.Value
	tok.Pos = participle.Position{
		Column:   nxt.Column,
		Line:     nxt.Line,
		Filename: p.fileName,
	}
	return tok, nil
}

type Definitions struct {
}

func (d *Definitions) Symbols() map[string]rune {
	return tokens.GetDefinitions()
}

func (d *Definitions) Lex(r io.Reader) (participle.Lexer, error) {
	lexer := &PLexer{}
	lexer.next = Lex(r)
	return lexer, nil
}
