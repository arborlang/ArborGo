package compiler

import (
	"bytes"
	"github.com/radding/ArborGo/internal/lexer"
	"github.com/radding/ArborGo/internal/parser/rules"
	"github.com/radding/ArborGo/internal/tokens"
	"os"
	"testing"
)

var testProgram = `
const test = (a: number, b: string): number -> {
	const numb = 1;
	const fa = 'a';

};

const test2 = (a: number, b: string): number ->  {
	
};

const something = (a: number, b:string, c: number): number -> {
};`

func TestCompilerVisitsFunctionCorrectly(t *testing.T) {
	parserStream := rules.New(lexer.Lex(bytes.NewReader([]byte(testProgram))))
	prog, err := rules.ProgramRule(parserStream, tokens.EOF)
	if err != nil {
		t.Fatal(err)
	}
	compiler := &Compiler{
		Writer: os.Stdout,
		Level:  0,
	}
	compiler.StartModule()
	if _, err := prog.Accept(compiler); err != nil {
		t.Fatal(err)
	}
	compiler.CloseModule()
}
