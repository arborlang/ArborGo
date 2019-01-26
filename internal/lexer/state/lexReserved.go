package state

import "github.com/arborlang/ArborGo/internal/lexer/internal"

func lexReserved(lex *internal.Lexer) State {
	lex.Emit(isReserved(lex.CurrentGroup()), nil)
	return lexText
}
