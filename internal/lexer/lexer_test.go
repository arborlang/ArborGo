package lexer

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/radding/ArborGo/internal/tokens"
)

func TestRunMachine(t *testing.T) {
	test := `this is a test function let there 1234
	be light! 12345.1234 0x123abcdef 0b011 1234 1234-1222`
	lexer := NewLexer(bytes.NewReader([]byte(test)))

	go RunMachine(lexer)

	for lexeme := range lexer.Lexemes {
		fmt.Printf("Token: %s, Value: %q\n", lexeme.Token, lexeme.Value)
	}
}

func TestLexAsync(t *testing.T) {
	test := `this is a test function let there 1234
	be light! 12345.1234 0x123abcdef 0b011 1234 1234-1222`

	lexemes := LexAsync(bytes.NewReader([]byte(test)))

	for lexeme := range lexemes {
		fmt.Printf("Token: %s, Value: %q\n", lexeme.Token, lexeme.Value)
	}
}

func TestLexSync(t *testing.T) {
	test := `this is a test function let there 1234
	be light! 12345.1234 0x123abcdef 0b011 1234 1234-1222`

	getNext := Lex(bytes.NewReader([]byte(test)))

	for lexeme := getNext(); lexeme.Token != tokens.EOF && lexeme.Token != tokens.ERROR; {
		fmt.Printf("Token: %s, Value: %q\n", lexeme.Token, lexeme.Value)
		lexeme = getNext()
	}
}
