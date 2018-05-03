package internal

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/radding/ArborGo/internal/tokens"
)

//Lexer is the lexer data implementation
type Lexer struct {
	name     string
	position int
	start    int
	input    string
	width    int
	col      int
	line     int
	Lexemes  chan Lexeme
}

//NewLexer takes an input reader and
func NewLexer(in io.Reader) *Lexer {
	buf := new(bytes.Buffer)
	buf.ReadFrom(in)
	return &Lexer{
		Lexemes: make(chan Lexeme, 4),
		input:   buf.String(),
		line:    1,
		col:     1,
	}
}

//SetName sets the name for the lexer (it should be the file name)
func (lexer *Lexer) SetName(name string) {
	lexer.name = name
}

//Errorf emits an error token on the stream
func (lexer *Lexer) Errorf(msg string) {
	lexer.Lexemes <- Lexeme{
		Token: tokens.ERROR,
		Value: fmt.Sprintf("%s:%d:%d: %s", lexer.name, lexer.line, lexer.col, msg),
	}
}

//Emit puts a lexeme on the lexemes channel
func (lexer *Lexer) Emit(tok tokens.Token) {
	lexer.Lexemes <- Lexeme{
		Token: tok,
		Value: lexer.CurrentGroup(),
	}
	lexer.start = lexer.position
	if tok == tokens.NEWLINE {
		lexer.col = 1
		lexer.line++
	}
}

//EmitEOF emits the EOF lexeme
func (lexer *Lexer) EmitEOF() {
	lexer.Lexemes <- Lexeme{
		Token: tokens.EOF,
		Value: string(tokens.EOFChar),
	}
	lexer.start = lexer.position
}

//Next returns the next rune in the in the input
func (lexer *Lexer) Next() (ch rune) {
	if lexer.position >= len(lexer.input) {
		lexer.width = 0
		return tokens.EOFChar
	}
	ch, lexer.width = utf8.DecodeRuneInString(lexer.input[lexer.position:])
	lexer.position += lexer.width
	lexer.col++
	return ch
}

//Ignore skips over the next rune
func (lexer *Lexer) Ignore() {
	lexer.start = lexer.position
}

//Backup backs up the scanner by one rune, resets width to what it was before next, theoretically, we can call this as many times as necessary
func (lexer *Lexer) Backup() {
	lexer.position -= lexer.width
	_, lexer.width = utf8.DecodeRuneInString(lexer.input[lexer.position:])
}

//Peek looks at the next rune without advancing the position
func (lexer *Lexer) Peek() rune {
	next := lexer.Next()
	lexer.Backup()
	return next
}

//Accept returns if next rune is in the valid string values. It backsup if the lexer doesn't accept
func (lexer *Lexer) Accept(valid string) bool {
	if strings.IndexRune(valid, lexer.Next()) >= 0 {
		return true
	}
	lexer.Backup()
	return false
}

//AcceptWhile keeps running while the input is valid
func (lexer *Lexer) AcceptWhile(valid string) {
	for strings.IndexRune(valid, lexer.Next()) >= 0 {

	}
	lexer.Backup()
}

//CurrentGroup returns the current group
func (lexer *Lexer) CurrentGroup() string {
	return lexer.input[lexer.start:lexer.position]
}
