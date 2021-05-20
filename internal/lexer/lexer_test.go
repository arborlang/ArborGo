package lexer

import (
	"bytes"
	"testing"

	"github.com/arborlang/ArborGo/internal/lexer/internal"
	"github.com/arborlang/ArborGo/internal/tokens"
)

var test = `
fn name = () ->
	return butt
done

x = a + b
value = 'a'
str = "abc dea"

fn test = (a, b, c) ->
	return a
done
`

var expectedTokenStream = []internal.Lexeme{
	internal.Lexeme{Token: tokens.FUNC, Value: "fn"},
	internal.Lexeme{Token: tokens.VARNAME, Value: "name"},
	internal.Lexeme{Token: tokens.EQUAL, Value: "="},
	internal.Lexeme{Token: tokens.RPAREN, Value: "("},
	internal.Lexeme{Token: tokens.LPAREN, Value: ")"},
	internal.Lexeme{Token: tokens.ARROW, Value: "->"},
	internal.Lexeme{Token: tokens.RETURN, Value: "return"},
	internal.Lexeme{Token: tokens.VARNAME, Value: "butt"},
	internal.Lexeme{Token: tokens.DONE, Value: "done"},
	internal.Lexeme{Token: tokens.VARNAME, Value: "x"},
	internal.Lexeme{Token: tokens.EQUAL, Value: "="},
	internal.Lexeme{Token: tokens.VARNAME, Value: "a"},
	internal.Lexeme{Token: tokens.ARTHOP, Value: "+"},
	internal.Lexeme{Token: tokens.VARNAME, Value: "b"},
	internal.Lexeme{Token: tokens.VARNAME, Value: "value"},
	internal.Lexeme{Token: tokens.EQUAL, Value: "="},
	internal.Lexeme{Token: tokens.CHARVAL, Value: "'a'"},
	internal.Lexeme{Token: tokens.VARNAME, Value: "str"},
	internal.Lexeme{Token: tokens.EQUAL, Value: "="},
	internal.Lexeme{Token: tokens.STRINGVAL, Value: `"abc dea"`},
	internal.Lexeme{Token: tokens.FUNC, Value: "fn"},
	internal.Lexeme{Token: tokens.VARNAME, Value: "test"},
	internal.Lexeme{Token: tokens.EQUAL, Value: "="},
	internal.Lexeme{Token: tokens.RPAREN, Value: "("},
	internal.Lexeme{Token: tokens.VARNAME, Value: "a"},
	internal.Lexeme{Token: tokens.COMMA, Value: ","},
	internal.Lexeme{Token: tokens.VARNAME, Value: "b"},
	internal.Lexeme{Token: tokens.COMMA, Value: ","},
	internal.Lexeme{Token: tokens.VARNAME, Value: "c"},
	internal.Lexeme{Token: tokens.LPAREN, Value: ")"},
	internal.Lexeme{Token: tokens.ARROW, Value: "->"},
	internal.Lexeme{Token: tokens.RETURN, Value: "return"},
	internal.Lexeme{Token: tokens.VARNAME, Value: "a"},
	internal.Lexeme{Token: tokens.DONE, Value: "done"},
	internal.Lexeme{Token: tokens.EOF, Value: string(tokens.EOFChar)},
}

func TestRunMachine(t *testing.T) {

	lexer := NewLexer(bytes.NewReader([]byte(test)))

	go RunMachine(lexer)
	index := 0
	for lexeme := range lexer.Lexemes {
		if index >= len(expectedTokenStream) {
			t.Fatal("Received token stream is longer than expected")
		}
		correctLexeme := expectedTokenStream[index]
		if correctLexeme.Token != lexeme.Token || correctLexeme.Value != lexeme.Value {
			t.Errorf("Lexemes don't match at position %v: expected %s, got %s", index, correctLexeme, lexeme)
		}
		index++
	}
}

func TestLexAsync(t *testing.T) {

	lexemes := LexAsync(bytes.NewReader([]byte(test)))

	index := 0
	for lexeme := range lexemes {
		if index >= len(expectedTokenStream) {
			t.Fatal("Received token stream is longer than expected")
		}
		correctLexeme := expectedTokenStream[index]
		if correctLexeme.Token != lexeme.Token || correctLexeme.Value != lexeme.Value {
			t.Errorf("Lexemes don't match at position %v: expected %s, got %s", index, correctLexeme, lexeme)
		}
		index++
	}
}

func TestLexSync(t *testing.T) {
	getNext := Lex(bytes.NewReader([]byte(test)))
	index := 0
	for lexeme := getNext(); lexeme.Token != tokens.EOF && lexeme.Token != tokens.ERROR; {
		if index >= len(expectedTokenStream) {
			t.Fatal("Received token stream is longer than expected")
		}
		correctLexeme := expectedTokenStream[index]
		if correctLexeme.Token != lexeme.Token || correctLexeme.Value != lexeme.Value {
			t.Errorf("Lexemes don't match at position %v: expected %s, got %s", index, correctLexeme, lexeme)
		}
		lexeme = getNext()
		index++
	}
}

func TestDoubleQuote(t *testing.T) {
	getNext := Lex(bytes.NewReader([]byte("Test::Foo")))
	expectedToks := []internal.Lexeme{{
		Value: "Test",
		Token: tokens.VARNAME,
	}, {
		Value: "::",
		Token: tokens.DCOLON,
	}, {
		Value: "Foo",
		Token: tokens.VARNAME,
	}}
	index := 0
	for lexeme := getNext(); lexeme.Token != tokens.EOF && lexeme.Token != tokens.ERROR; {
		if index >= len(expectedToks) {
			t.Fatal("Received token stream is longer than expected")
		}
		correctLexeme := expectedToks[index]
		if correctLexeme.Token != lexeme.Token || correctLexeme.Value != lexeme.Value {
			t.Errorf("Lexemes don't match at position %v: expected %s, got %s", index, correctLexeme, lexeme)
		}
		lexeme = getNext()
		index++
	}

}
