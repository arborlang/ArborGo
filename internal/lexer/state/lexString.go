package state

import (
	"github.com/radding/ArborGo/internal/lexer/internal"
	"github.com/radding/ArborGo/internal/tokens"
)

func lexString(lex *internal.Lexer) State {
	for {
		next := lex.Next()
		if next == '\\' {
			lex.Next()
		}
		if next == '"' {
			lex.Emit(tokens.STRINGVAL, nil)
			return lexText
		}
		if next == tokens.EOFChar {
			return lexError("String was never terminated! Reached EOF")
		}

	}
	return lexError("Holy Shit! You reached this error? Something fucked up!")
}
