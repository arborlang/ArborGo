package participle

import (
	"fmt"
	"io"

	"github.com/alecthomas/participle"
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast"
)

type Entry struct {
	Properties []*Property `@@*`
	Sections   []*Section  `@@*`
}

type Section struct {
	Identifier string      `"[" @Ident "]"`
	Properties []*Property `@@*`
}

type Property struct {
	Key   string `@Ident "="`
	Value *Value `@@`
}

type Value struct {
	String *string  `  @String`
	Number *float64 `| @Float`
}

func Parse(source io.Reader) error {
	parser, err := participle.Build(&ast.Program{}, participle.Lexer(&lexer.Definitions{}))
	if err != nil {
		fmt.Println(err)
		return err
	}
	astNode := &ast.Program{}
	err = parser.ParseString("import blah from \"boop\"", astNode)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
