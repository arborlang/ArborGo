package state

import (
	"fmt"

	"github.com/radding/ArborGo/internal/lexer/internal"
	"github.com/radding/ArborGo/internal/tokens"
)

func lexNumeric(lex *internal.Lexer) State {
	isArithmetic := func(ch rune) bool {
		return ch == '+' || ch == '-' || ch == '*' || ch == '/'
	}
	acceptNumbers := func() {
		lex.Accept("+-")
		acceptable := "0123456789"
		if lex.Accept("0") {
			if lex.Accept("bB") {
				//match binary 0bxxxxxx or 0Bxxxxxxx
				acceptable = "01"
			} else if lex.Accept("xX") {
				//match hexidecimal 0X... or 0x....
				acceptable = "0123456789abcdefABCDEF"
			}
		}
		lex.AcceptWhile(acceptable)
	}
	acceptNumbers()
	if lex.Accept(".") {
		acceptNumbers()
		if !isWhitespace(lex.Peek()) && !isArithmetic(lex.Peek()) {
			return lexError(fmt.Sprintf("float mismatch: %s", lex.CurrentGroup()+string(lex.Peek())))
		}
		lex.Emit(tokens.FLOAT)
		if lex.Accept("+-/*") {
			lex.Emit(tokens.ARTHOP)
		}
		return lexText
	}
	if !isWhitespace(lex.Peek()) && !isArithmetic(lex.Peek()) {
		return lexError(fmt.Sprintf("number mismatch: %s", lex.CurrentGroup()+string(lex.Peek())))
	}
	lex.Emit(tokens.NUMBER)
	if lex.Accept("+-/*") {
		lex.Emit(tokens.ARTHOP)
	}
	return lexText
}
