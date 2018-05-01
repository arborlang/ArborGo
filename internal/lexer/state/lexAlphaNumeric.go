package state

import (
	"github.com/radding/ArborGo/internal/lexer/internal"
	"github.com/radding/ArborGo/internal/tokens"
)

func lexAlphaNumeric(lex *internal.Lexer) State {
	for isAlphaNumeric(lex.Next()) {
	}
	lex.Backup()
	if isReserved(lex.CurrentGroup()) >= 0 {
		return lexReserved
	}
	lex.Emit(tokens.VARNAME)
	return lexText

}
