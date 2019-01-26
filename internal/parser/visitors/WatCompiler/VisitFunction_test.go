package wast

import (
	"bytes"
	"os"
	"testing"

	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/rules"
	"github.com/arborlang/ArborGo/internal/tokens"
)

var testProgram = `
const test = (a: number, b: string): number -> {
	const numb = 1;
	const fa = 'a';
	return numb;
};

const test2 = (a: number, b: string): number ->  {
	return 0;
};

const something = (a: number, b:string, c: number): number -> {
	return 0;
};`

func TestCompilerVisitsFunctionCorrectly(t *testing.T) {
	parserStream := rules.New(lexer.Lex(bytes.NewReader([]byte(testProgram))))
	prog, err := rules.ProgramRule(parserStream, tokens.EOF)
	if err != nil {
		t.Fatal(err)
	}
	compiler := &Compiler{
		Writer: os.Stdout,
		level:  0,
	}
	compiler.StartModule()
	if _, err := prog.Accept(compiler); err != nil {
		t.Fatal(err)
	}
	compiler.CloseModule()
}
