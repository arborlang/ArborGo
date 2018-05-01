package lexer

import (
	"io"

	"github.com/radding/ArborGo/internal/lexer/state"

	"github.com/radding/ArborGo/internal/lexer/internal"
)

//NewLexer returns a new lexer instance
func NewLexer(in io.Reader) *internal.Lexer {
	return internal.NewLexer(in)
}

//RunMachine runs the lexer
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
func Lex(in io.Reader) func() internal.Lexeme {
	lex := NewLexer(in)
	stateFunc := state.LexText
	return func() internal.Lexeme {
		for {
			select {
			case lexeme := <-lex.Lexemes:
				return lexeme
			default:
				stateFunc = stateFunc(lex)
			}
		}
	}
}
