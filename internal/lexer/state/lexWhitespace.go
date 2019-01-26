package state

import (
	"github.com/arborlang/ArborGo/internal/lexer/internal"
	// "github.com/arborlang/ArborGo/internal/tokens"
)

func lexWhiteSpace(lex *internal.Lexer) State {
	lex.AcceptWhile(" \t")
	for next := lex.Peek(); next == '\n'; next = lex.Peek() {
		lex.Next()
		lex.NewLine()
	}
	lex.Ignore()
	// next := lex.Next()
	// if next == '\n' {
	// 	lex.Emit(tokens.NEWLINE)
	// 	lex.AcceptWhile("\n")
	// 	lex.Ignore()
	// } else {
	// 	lex.Backup()
	// }
	return lexText
}
