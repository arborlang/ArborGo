package state

import (
	"github.com/radding/ArborGo/internal/lexer/internal"
	"github.com/radding/ArborGo/internal/tokens"
)

func isComparison(lex *internal.Lexer, next rune) bool {
	isLT := next == '<'
	isLTE := next == '<' && lex.Peek() == '='
	isGT := next == '>'
	isGTE := next == '>' && lex.Peek() == '='
	isEq := next == '=' && lex.Peek() == '='
	isNeq := next == '!' && lex.Peek() == '='
	if isLTE || isGTE || isEq || isNeq {
		lex.Next()
	}
	return isLT || isLTE || isGT || isGTE || isEq || isNeq
}

//lexTextis the first state
func lexText(lex *internal.Lexer) State {
	for {
		next := lex.Next()
		if next == tokens.EOFChar {
			break
		}
		switch {
		case next == '+' || next == '-':
			if next == '-' && lex.Peek() == '>' {
				lex.Next()
				lex.Emit(tokens.ARROW, nil)
				return lexText
			}
			if lex.Peek() >= '0' && lex.Peek() <= '9' {
				lex.Backup()
				return lexNumeric
			}
			lex.Emit(tokens.ARTHOP, nil)
		case next == '/' && lex.Peek() == '/':
			return func(lex *internal.Lexer) State {
				for {
					next := lex.Next()
					if next == '\n' {
						lex.Ignore()
						return lexText
					}
				}
			}
		case next == '/' && lex.Peek() == '*':
			return func(lex *internal.Lexer) State {
				for {
					next := lex.Next()
					if next == '*' && lex.Peek() == '/' {
						lex.Next()
						lex.Ignore()
						return lexText
					}
				}
			}
		case isComparison(lex, next):
			lex.Emit(tokens.COMPARISON, nil)
		case next == '|' && lex.Peek() == '>':
			lex.Next()
			lex.Emit(tokens.PIPE, nil)
		case next == '*' || next == '/':
			lex.Emit(tokens.ARTHOP, nil)
		case next >= '0' && next <= '9':
			lex.Backup()
			return lexNumeric
		case next == '!':
			lex.Emit(tokens.NOT, nil)
		case (next == '&' && lex.Peek() == '&') || (next == '|' && lex.Peek() == '|'):
			lex.Next()
			lex.Emit(tokens.BOOLEAN, nil)
		case next == '&' || next == '|' || next == '^':
			lex.Emit(tokens.LOGICAL, nil)
		case isAlphaNumeric(next):
			lex.Backup()
			return lexAlphaNumeric
		case isWhitespace(next):
			lex.Backup()
			return lexWhiteSpace
		case next == '=':
			lex.Emit(tokens.EQUAL, nil)
		case next == '(':
			lex.Emit(tokens.RPAREN, nil)
		case next == ')':
			lex.Emit(tokens.LPAREN, nil)
		case next == '\'':
			return lexChar
		case next == '"':
			return lexString
		case next == ':':
			lex.Emit(tokens.COLON, nil)
		case next == ',':
			lex.Emit(tokens.COMMA, nil)
		case next == ';':
			lex.Emit(tokens.SEMI, nil)
		case next == '{':
			lex.Emit(tokens.RCURLY, nil)
		case next == '}':
			lex.Emit(tokens.LCURLY, nil)

		}
	}
	lex.EmitEOF()
	return nil
}

//LexText is the public entry point for all lexer states
func LexText(lex *internal.Lexer) State {
	return lexText
}
