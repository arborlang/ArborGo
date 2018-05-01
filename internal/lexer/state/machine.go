package state

import (
	"github.com/radding/ArborGo/internal/lexer/internal"
)

//State is a representation of the Lexer state, the Lexer returns the next State function
type State func(*internal.Lexer) State

//Machine starts the state machine for the lexer
func Machine(lex *internal.Lexer) {
	//Run runs the lexer. This is basically the state machine of the Lexer
	for lexState := lexText; lexState != nil; {
		lexState = lexState(lex)
	}
	close(lex.Lexemes)
}
