package state

import (
	"github.com/arborlang/ArborGo/internal/lexer/internal"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func lexAlphaNumeric(lex *internal.Lexer) State {
	for isAlphaNumeric(lex.Next()) {
	}
	lex.Backup()
	if isReserved(lex.CurrentGroup()) >= 0 {
		return lexReserved
	}
	lex.Emit(tokens.VARNAME, nil)
	return lexText

}
