package internal

import (
	"fmt"

	"github.com/radding/ArborGo/internal/tokens"
)

//Lexeme is the value of the string
type Lexeme struct {
	Token tokens.Token
	Value string
}

func (lexeme Lexeme) String() string {
	return fmt.Sprintf("{ Token: %s, Value: %q }", lexeme.Token, lexeme.Value)
}
