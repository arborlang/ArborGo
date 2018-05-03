package state

import (
	"fmt"

	"github.com/radding/ArborGo/internal/tokens"

	"github.com/radding/ArborGo/internal/lexer/internal"
)

func lexChar(lex *internal.Lexer) State {
	next := lex.Next()
	if next == '\\' {
		lex.Next()
	}
	next = lex.Next()
	if next != '\'' {
		return lexError(fmt.Sprintf("Expected end of char (') got %q instead", next))
	}
	lex.Emit(tokens.CHARVAL)
	return lexText
}
