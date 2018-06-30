//Package arbor is the public api for the arbor programming langauge.
package arbor

import (
	"io"

	"github.com/radding/ArborGo/internal/ast"
	"github.com/radding/ArborGo/internal/lexer"
	"github.com/radding/ArborGo/internal/parser"
)

// Parse takes some input string and returns the AST for the string or an error if it failed.
// This method does both the Lexing and the parsing and returns the ast for you to be able to
// 	walk through it.
func Parse(r io.Reader) (ast.Node, error) {
	lexReader := lexer.Lex(r)
	return parser.Parse(lexer.NewBufferedReader(lexReader))
}
