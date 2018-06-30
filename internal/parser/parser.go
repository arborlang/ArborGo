//Package parser is the entrypoint to our arbor parser. This parser is where the magic happens
package parser

import (
	"github.com/radding/ArborGo/internal/ast"
	"github.com/radding/ArborGo/internal/lexer"
)

//Parser is a type that actually does the parsing of the arbor compiler. Parser takes a function
// that returns the next Lexeme on the stream (AKA a Reader) and parses the Applicable stream and
// returns an AST node
type Parser func(*lexer.BufferedReader) (ast.Node, error)

//Parse is the entry point to the Parser. The idea is that you pass the source code to the Lexer,
// then the Lexer gives the parser the Stream and the parser returns the AST
func Parse(reader *lexer.BufferedReader) (ast.Node, error) {
	return parseStatements(reader)
}
