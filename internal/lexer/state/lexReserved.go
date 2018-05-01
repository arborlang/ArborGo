package state

import "github.com/radding/ArborGo/internal/lexer/internal"

func lexReserved(lex *internal.Lexer) State {
	lex.Emit(isReserved(lex.CurrentGroup()))
	return lexText
}
