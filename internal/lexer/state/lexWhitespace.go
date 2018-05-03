package state

import (
	"github.com/radding/ArborGo/internal/lexer/internal"
	"github.com/radding/ArborGo/internal/tokens"
)

func lexWhiteSpace(lex *internal.Lexer) State {
	lex.AcceptWhile(" \t")
	lex.Ignore()
	next := lex.Next()
	if next == '\n' {
		lex.Emit(tokens.NEWLINE)
		lex.AcceptWhile("\n")
		lex.Ignore()
	} else {
		lex.Backup()
	}
	return lexText
}
