//Package lexer is the lexer for the arbor project. The lexer can either be run asyncrounously or syncrounously
package lexer

import (
	"io"

	"github.com/radding/ArborGo/internal/lexer/internal"
	"github.com/radding/ArborGo/internal/lexer/state"
)

//Lexeme is the public lexeme for everyone
type Lexeme internal.Lexeme

// Reader is a type of function that consumes a Lexer and returns one lexeme every time it is called.
// 	When the reader is done, it will return an EOF symbol.
type Reader func() internal.Lexeme

//NewLexer creates and returns a new lexer instance
func NewLexer(in io.Reader) *internal.Lexer {
	return internal.NewLexer(in)
}

//RunMachine runs the lexer state machine
func RunMachine(lex *internal.Lexer) {
	state.Machine(lex)
}

//LexAsync lexs the input asyncronously. This returns a channel that will contain the identified lexemes
func LexAsync(in io.Reader) chan internal.Lexeme {
	lex := NewLexer(in)
	go RunMachine(lex)
	return lex.Lexemes
}

// Lex lexes the input syncrounously by returning a function that will listen for input on the channel else it will call the state function.
// 	This function returns a function that can be used to get the next lexeme
func Lex(in io.Reader) Reader {
	lex := NewLexer(in)
	stateFunc := state.LexText
	return func() internal.Lexeme {
		for {
			select {
			case lexeme := <-lex.Lexemes:
				return lexeme
			default:
				stateFunc = stateFunc(lex)
			}
		}
	}
}

//BufferedReader Allows us to peak and view the lex strea
type BufferedReader struct {
	reader Reader
	buffer []Lexeme
}

//NewBufferedReader gives us a new buffered reader
func NewBufferedReader(reader Reader) *BufferedReader {
	return &BufferedReader{reader: reader, buffer: make([]Lexeme, 0)}
}

//Next gets the next Lexeme, either from the buffer or the reader its self
func (b *BufferedReader) Next() Lexeme {
	var lexeme Lexeme
	if len(b.buffer) > 0 {
		lexeme, b.buffer = b.buffer[0], b.buffer[1:]
	} else {
		lexeme = Lexeme(b.reader())
	}
	return lexeme
}

//Look looks at the token N places in the future without consuming the tokens.
func (b *BufferedReader) Look(n int) Lexeme {
	leftOver := n - len(b.buffer)
	if leftOver > 0 {
		for i := 0; i < leftOver+1; i++ {
			b.buffer = append(b.buffer, Lexeme(b.reader()))
		}
	}
	return b.buffer[n-1]
}

//Peek Gets the next Lexeme in the stream, but doesn't consume it.
func (b *BufferedReader) Peek() Lexeme {
	return b.Look(1)
}
