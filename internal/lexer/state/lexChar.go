package state

import (
	"fmt"

	"github.com/radding/ArborGo/internal/tokens"

	"github.com/radding/ArborGo/internal/lexer/internal"
)

func encodeChar(tok tokens.Token, value string) []byte {
	if value[1] == byte('\\') {
		switch value[1:3] {
		case `\n`:
			return []byte{39, 10, 39}
		case `\t`:
			return []byte{39, byte('\t'), 39}

		case `\b`:
			return []byte{39, byte('\b'), 39}
		case `\f`:
			return []byte{39, byte('\f'), 39}
		case `\r`:
			return []byte{39, 13, 39}
		case `\v`:
			return []byte{39, byte('\v'), 39}
		default:
			return []byte{39, value[2], 39}
		}
	}
	return []byte(value)
}

func lexChar(lex *internal.Lexer) State {
	next := lex.Next()
	if next == '\\' {
		lex.Next()
	}
	next = lex.Next()
	if next != '\'' {
		return lexError(fmt.Sprintf("Expected end of char (') got %q instead", next))
	}
	lex.Emit(tokens.CHARVAL, encodeChar)
	return lexText
}
