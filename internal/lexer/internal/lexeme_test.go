package internal

import (
	"testing"

	"github.com/radding/ArborGo/internal/tokens"
)

func TestLexerReturns(t *testing.T) {
	lexer := RegisterLexeme(".*", tokens.VARNAME)

	if lexer == nil {
		t.Error("lexer is nil!")
	}

	lexeme, err := lexer("Test this!")

	if err != nil {
		t.Fatal(err)
	}
	if lexeme == nil {
		t.Error("lexeme is nil")
	}

	if lexeme.Token != tokens.VARNAME {
		t.Errorf("Expected tokens.VARNAME, got %d instead", lexeme.Token)
	}

	if lexeme.Value != "Test this!" {
		t.Errorf("Expected tokens.VARNAME, got %s instead", lexeme.Value)
	}

	lexer = RegisterLexeme("[0-9]+[a-zA-Z_]", tokens.ARROW)

	lexeme, err = lexer("1234A")

	if err != nil {
		t.Fatal(err)
	}

	if lexeme.Token != tokens.ARROW {
		t.Error("Exepected tokens.ARROW got", lexeme.Token, "instead")
	}

	if lexeme.Value != "1234A" {
		t.Error("Expected 1234A, got", lexeme.Value, "instead")
	}
}
