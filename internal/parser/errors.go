package parser

import (
	"fmt"

	"github.com/radding/ArborGo/internal/lexer"
)

//BaseError is the base Parser error that contains basic info about the error
type BaseError struct {
	LineNumber   int
	ColumnNumber int
}

//EOF is an error that is read at the end of the Input.
type EOF struct {
	BaseError
}

func (EOF) Error() string {
	return "normally encountered eof"
}

//PrematureEOF is an error that is returned when the parser encounters an EOF too early
type PrematureEOF struct {
	BaseError
}

func (PrematureEOF) Error() string {
	return "pre-maturally encountered an eof"
}

//UnrecognizedSymbol is raised when the symbol is not expected
type UnrecognizedSymbol struct {
	BaseError
	lexeme lexer.Lexeme
}

func (e UnrecognizedSymbol) Error() string {
	return fmt.Sprintf("unrecognized symbol %s encountered", e.lexeme.Value)
}

//NewUnrecognizedError creates an unrecognized error
func NewUnrecognizedError(lexeme lexer.Lexeme) UnrecognizedSymbol {
	return UnrecognizedSymbol{
		BaseError: BaseError{},
		lexeme:    lexeme,
	}
}
