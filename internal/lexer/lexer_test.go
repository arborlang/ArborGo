package lexer

import (
	"bytes"
	"testing"

	"github.com/radding/ArborGo/internal/lexer/internal"
	"github.com/radding/ArborGo/internal/tokens"
	"github.com/stretchr/testify/assert"
)

// TODO: Rewrite these tests using testify

var test = `
func name = () ->
	return butt
done

x = a + b
value = 'a'
str = "abc dea"

func test = (a, b, c) ->
	return a
done
`

var expectedTokenStream = []internal.Lexeme{
	internal.Lexeme{Token: tokens.NEWLINE, Value: "\n"},
	internal.Lexeme{Token: tokens.FUNC, Value: "func"},
	internal.Lexeme{Token: tokens.VARNAME, Value: "name"},
	internal.Lexeme{Token: tokens.EQUAL, Value: "="},
	internal.Lexeme{Token: tokens.RPAREN, Value: "("},
	internal.Lexeme{Token: tokens.LPAREN, Value: ")"},
	internal.Lexeme{Token: tokens.ARROW, Value: "->"},
	internal.Lexeme{Token: tokens.NEWLINE, Value: "\n"},
	internal.Lexeme{Token: tokens.RETURN, Value: "return"},
	internal.Lexeme{Token: tokens.VARNAME, Value: "butt"},
	internal.Lexeme{Token: tokens.NEWLINE, Value: "\n"},
	internal.Lexeme{Token: tokens.DONE, Value: "done"},
	internal.Lexeme{Token: tokens.NEWLINE, Value: "\n"},
	internal.Lexeme{Token: tokens.VARNAME, Value: "x"},
	internal.Lexeme{Token: tokens.EQUAL, Value: "="},
	internal.Lexeme{Token: tokens.VARNAME, Value: "a"},
	internal.Lexeme{Token: tokens.ARTHOP, Value: "+"},
	internal.Lexeme{Token: tokens.VARNAME, Value: "b"},
	internal.Lexeme{Token: tokens.NEWLINE, Value: "\n"},
	internal.Lexeme{Token: tokens.VARNAME, Value: "value"},
	internal.Lexeme{Token: tokens.EQUAL, Value: "="},
	internal.Lexeme{Token: tokens.CHARVAL, Value: "'a'"},
	internal.Lexeme{Token: tokens.NEWLINE, Value: "\n"},
	internal.Lexeme{Token: tokens.VARNAME, Value: "str"},
	internal.Lexeme{Token: tokens.EQUAL, Value: "="},
	internal.Lexeme{Token: tokens.STRINGVAL, Value: `"abc dea"`},
	internal.Lexeme{Token: tokens.NEWLINE, Value: "\n"},
	internal.Lexeme{Token: tokens.FUNC, Value: "func"},
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
	internal.Lexeme{Token: tokens.NEWLINE, Value: "\n"},
	internal.Lexeme{Token: tokens.RETURN, Value: "return"},
	internal.Lexeme{Token: tokens.VARNAME, Value: "a"},
	internal.Lexeme{Token: tokens.NEWLINE, Value: "\n"},
	internal.Lexeme{Token: tokens.DONE, Value: "done"},
	internal.Lexeme{Token: tokens.NEWLINE, Value: "\n"},
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

func TestBufferedReaderPerformsTheSame(t *testing.T) {
	getNext := Lex(bytes.NewReader([]byte(test)))
	index := 0
	bufferedReader := NewBufferedReader(getNext)
	for lexeme := bufferedReader.Next(); lexeme.Token != tokens.EOF && lexeme.Token != tokens.ERROR; {
		if index >= len(expectedTokenStream) {
			t.Fatal("Received token stream is longer than expected")
		}
		correctLexeme := expectedTokenStream[index]
		if correctLexeme.Token != lexeme.Token || correctLexeme.Value != lexeme.Value {
			t.Errorf("Lexemes don't match at position %v: expected %s, got %s", index, correctLexeme, lexeme)
		}
		lexeme = bufferedReader.Next()
		index++
	}
}

func TestBufferedReadersPeekDoesntBreak(t *testing.T) {
	assert := assert.New(t)
	getNext := Lex(bytes.NewReader([]byte(test)))
	bufferedReader := NewBufferedReader(getNext)

	letsTakeAPeak := bufferedReader.Peek()
	whatNow := bufferedReader.Next()
	realToken := expectedTokenStream[0]

	assert.Equal(realToken.Token, letsTakeAPeak.Token)
	assert.Equal(realToken.Value, letsTakeAPeak.Value)
	assert.Equal(whatNow.Token, letsTakeAPeak.Token)
	assert.Equal(whatNow.Value, letsTakeAPeak.Value)

	letsTakeAPeak = bufferedReader.Peek()
	whatNow = bufferedReader.Next()
	realToken = expectedTokenStream[1]

	assert.Equal(realToken.Token, letsTakeAPeak.Token)
	assert.Equal(realToken.Value, letsTakeAPeak.Value)
	assert.Equal(whatNow.Token, letsTakeAPeak.Token)
	assert.Equal(whatNow.Value, letsTakeAPeak.Value)

	whatNow = bufferedReader.Next()
	letsTakeAPeak = bufferedReader.Peek()

	assert.NotEqual(whatNow.Token, letsTakeAPeak.Token)
	assert.NotEqual(whatNow.Value, letsTakeAPeak.Value)
}

func TestBufferedReaderLookWorks(t *testing.T) {
	assert := assert.New(t)
	getNext := Lex(bytes.NewReader([]byte(test)))
	bufferedReader := NewBufferedReader(getNext)

	letsTakeAPeak := bufferedReader.Look(1)
	realToken := expectedTokenStream[0]

	assert.Equal(realToken.Token, letsTakeAPeak.Token)
	assert.Equal(realToken.Value, letsTakeAPeak.Value)

	letsTakeAPeak = bufferedReader.Look(4)
	realToken = expectedTokenStream[3]
	assert.Equal(realToken.Token, letsTakeAPeak.Token)
	assert.Equal(realToken.Value, letsTakeAPeak.Value)

	letsTakeAPeak = bufferedReader.Look(1)
	realToken = expectedTokenStream[0]
	assert.Equal(realToken.Token, letsTakeAPeak.Token)
	assert.Equal(realToken.Value, letsTakeAPeak.Value)

	letsTakeAPeak = bufferedReader.Next()
	assert.Equal(realToken.Token, letsTakeAPeak.Token)
	assert.Equal(realToken.Value, letsTakeAPeak.Value)
}
