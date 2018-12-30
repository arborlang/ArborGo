package rules

import (
	"bytes"
	"github.com/radding/ArborGo/internal/lexer"
	"github.com/radding/ArborGo/internal/tokens"
	"github.com/stretchr/testify/assert"
	"testing"
)

var test = `
let name = () ->
	return butt
done

x = a + b
value = 'a'
str = "abc dea"

let test = (a, b, c) ->
	return a
done
`

func normalizeLexemes(l lexer.Lexeme) lexer.Lexeme {
	return lexer.Lexeme{
		Token:  l.Token,
		Value:  l.Value,
		Column: 0,
		Line:   0,
	}
}
func TestParserStream(t *testing.T) {
	assert := assert.New(t)
	tokStream := lexer.Lex(bytes.NewReader([]byte(test)))
	parser := New(tokStream)
	next := parser.Next()

	assert.NotNil(next)
	assert.Equal(normalizeLexemes(next), lexer.Lexeme{Token: tokens.LET, Value: "let"})
	peek := parser.Peek()
	assert.NotNil(peek)
	assert.Equal(normalizeLexemes(peek), lexer.Lexeme{Token: tokens.VARNAME, Value: "name"})

	peek = parser.Next()
	assert.NotNil(peek)
	assert.Equal(normalizeLexemes(peek), lexer.Lexeme{Token: tokens.VARNAME, Value: "name"})

	peek = parser.Peek()
	assert.NotNil(peek)
	assert.Equal(normalizeLexemes(peek), lexer.Lexeme{Token: tokens.EQUAL, Value: "="})

	peek = parser.Previous()
	assert.NotNil(peek)
	assert.Equal(normalizeLexemes(peek), lexer.Lexeme{Token: tokens.LET, Value: "let"})

	parser.Backup()
	peek = parser.Peek()
	assert.NotNil(peek)
	assert.Equal(normalizeLexemes(peek), lexer.Lexeme{Token: tokens.VARNAME, Value: "name"})
}
