package state

import (
	"github.com/radding/ArborGo/internal/lexer/internal"
	"github.com/radding/ArborGo/internal/tokens"
)

//lex start is the first state
func lexText(lex *internal.Lexer) State {
	for {
		next := lex.Next()
		if next == tokens.EOFChar {
			break
		}
		switch {
		case next == '+' || next == '-':
			if lex.Peek() >= '0' && lex.Peek() <= '9' {
				lex.Backup()
				return lexNumeric
			}
			lex.Backup()
			return nil
		case next >= '0' && next <= '9':
			lex.Backup()
			return lexNumeric
		case isAlphaNumeric(next):
			lex.Backup()
			return lexAlphaNumeric
		case isWhitespace(next):
			lex.Backup()
			return lexWhiteSpace
		}
	}
	lex.EmitEOF()
	return nil
}

//LexText is the public entry point for all lexer states
func LexText(lex *internal.Lexer) State {
	return lexText
}
