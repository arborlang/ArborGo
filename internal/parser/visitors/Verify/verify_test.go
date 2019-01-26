package verify

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/rules"
	"github.com/arborlang/ArborGo/internal/tokens"
	"github.com/stretchr/testify/assert"
)

var TestCases = []struct {
	// Name     string
	TestCase string
	IsValid  bool
}{
	{`const bap = 1;`, true},
	{`const bap:number;`, false},
	{`let bap = 1;`, true},
	{`let bap: number;`, true},
}

func TestIsValid(t *testing.T) {
	assert := assert.New(t)
	for _, testCase := range TestCases {
		t.Run(testCase.TestCase, func(t *testing.T) {
			parserStream := rules.New(lexer.Lex(bytes.NewReader([]byte(testCase.TestCase))))
			prog, err := rules.ProgramRule(parserStream, tokens.EOF)
			if err != nil {
				t.Fatal(err)
			}
			verify := New()
			_, err = prog.Accept(verify)
			assert.Equal(testCase.IsValid, err == nil)
			fmt.Println(err)
		})
	}
}
