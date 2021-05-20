package rulesv2

import (
	"fmt"
	"strings"

	"github.com/arborlang/ArborGo/internal/lexer"
)

type unexpectedError struct {
	expected []string
	got      lexer.Lexeme
}

func (u unexpectedError) Error() string {
	if len(u.expected) == 0 {
		return fmt.Sprintf("unexpected token: %s", u.got)
	}
	return fmt.Sprintf("unexpected token: wanted %s, got %s", strings.Join(u.expected, " or"), u.got)
}

// UnexpectedError is an unexpected Lexeme
func UnexpectedError(got lexer.Lexeme, wanted ...string) unexpectedError {
	quotedGot := []string{}
	for _, i := range wanted {
		quotedGot = append(quotedGot, fmt.Sprintf("%q", i))
	}
	return unexpectedError{
		expected: quotedGot,
		got:      got,
	}
}
