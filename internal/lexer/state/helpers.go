package state

import "github.com/arborlang/ArborGo/internal/tokens"

//isReserved checks if the value is a reserved key word
func isReserved(test string) tokens.Token {
	return tokens.FindKeyword(test)
}

//isWhitespace tells you if this character is a whitespace, character
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == tokens.EOFChar
}

//isAlphaNumeric tells you if the rune is an alpha numeric rune (a-zA-Z0-9) and if it is an underscore
func isAlphaNumeric(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || ch == '_'
}
