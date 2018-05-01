package internal

import (
	"github.com/radding/ArborGo/internal/tokens"
)

//Lexeme is the value of the string
type Lexeme struct {
	Token tokens.Token
	Value string
}
