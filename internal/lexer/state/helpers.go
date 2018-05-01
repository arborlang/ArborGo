package state

import "github.com/radding/ArborGo/internal/tokens"

//isReserved checks if the value is a reserved key word
func isReserved(test string) tokens.Token {
	return tokens.FindKeyword(test)
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == tokens.EOFChar
}

func isAlphaNumeric(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || ch == '_'
}
