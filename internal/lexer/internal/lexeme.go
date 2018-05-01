package internal

import (
	"regexp"

	"github.com/radding/ArborGo/internal/tokens"
)

//Lexeme is the value of the string
type Lexeme struct {
	Token tokens.Token
	Value string
}

//Tokenizer is what takes the string and turns it into a lexeme
type Tokenizer func(value string) (*Lexeme, error)

//RegisterLexeme returns a function that can identify matches and returns the lexeme
func RegisterLexeme(regexVal string, token tokens.Token) Tokenizer {
	regex, err := regexp.Compile(regexVal)

	return func(value string) (*Lexeme, error) {
		if err != nil {
			return nil, err
		}
		matched := regex.Match([]byte(value))
		if matched {
			return &Lexeme{
				Token: token,
				Value: value,
			}, nil
		}
		return nil, nil
	}
}
