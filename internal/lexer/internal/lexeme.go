package internal

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/tokens"
)

//Lexeme is the value of the string
type Lexeme struct {
	Token   tokens.Token
	Value   string
	Column  int
	Line    int
	RuneVal []byte
}

func (lexeme Lexeme) String() string {
	return fmt.Sprintf("{ Token: %s, Value: %q, Line: %d, Column: %d }", lexeme.Token, lexeme.Value, lexeme.Line, lexeme.Column)
}
