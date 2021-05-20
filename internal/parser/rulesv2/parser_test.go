package rulesv2

import (
	"bytes"
	"testing"

	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/tokens"
	"github.com/stretchr/testify/assert"
)

var test = `
let name = fn () -> {
	return butt
}

x = a + b
value = 'a'
str = "abc dea"

fn testB (a: int, b: string, c: float) -> {
	return a
}
`

var test2 = `
import stdlib as foobar from "github.com/arborlang/stdlib";

type Test {
	Example: fn () -> bool;
	Name: string;
};

fn Test::__Construct() {
};


fn Bazzer (val: foobar) -> number {
	match val {
		|> {foobar: 1, bazBar: 2} {
			return;
		}
		|> val.foobar == 10 || val.bazBar == 2 {
			return 0;
		}
	}
	return 1;
};

fn Test::Decorator<TCall> (val: TCall) -> TCall {
	if val == "something" {
		return help;
	} else if val.value > 1 {
		return narp;
	} else {
		hello;
	}
	const val = fatal new SomeRandomError();
	return val[33];
};

@Decorator
fn DecoratedVal (value: number) -> number{
	value.toString().someOtherAccess;
	value.toString() == "buttt"[0:2];
	self.tryThing();
	try {
		signal new SomeRandomSignal();
		const warnResponse = warn new SomeRandomWarning();
		fatal new SomeRandomFatal();
	} handle (e: SomeRandomSignal) {
		continue;
	} handle (w: SomeRandomWarning) {
		continue with "RetryWarning";
	};
	return 2 + 1;
};

let test = new Test();
test |> DecoratedVal;

type Stringer extends iface1 {
	Thing: fn (xyz: String) -> String;
};

fn Stringer::Thing(xyz: String) -> String {
	return xyz;
};

type Find implements ABC, XYZ, ZZZ Boop;

type Thing String;
`

func normalizeLexemes(l lexer.Lexeme) lexer.Lexeme {
	return lexer.Lexeme{
		Token:  l.Token,
		Value:  l.Value,
		Column: 0,
		Line:   0,
	}
}

func TestCanParseTestProgram(t *testing.T) {
	assert := assert.New(t)

	p := parseTest(test2)
	_, err := ProgramRule(p, tokens.EOF)
	assert.NoError(err)
}
func TestParserStream(t *testing.T) {
	assert := assert.New(t)
	tokStream := lexer.Lex(bytes.NewReader([]byte(test)))
	parser := New(tokStream)
	next := parser.Next()

	assert.NotNil(next)
	assert.Equal(normalizeLexemes(next), lexer.Lexeme{Token: tokens.LET, Value: "let"})
	peek := parser.Peek()
	assert.NotNil(peek)
	assert.Equal(normalizeLexemes(peek), lexer.Lexeme{Token: tokens.VARNAME, Value: "name"})

	peek = parser.Next()
	assert.NotNil(peek)
	assert.Equal(normalizeLexemes(peek), lexer.Lexeme{Token: tokens.VARNAME, Value: "name"})

	peek = parser.Peek()
	assert.NotNil(peek)
	assert.Equal(normalizeLexemes(peek), lexer.Lexeme{Token: tokens.EQUAL, Value: "="})

	peek = parser.Previous()
	assert.NotNil(peek)
	assert.Equal(normalizeLexemes(peek), lexer.Lexeme{Token: tokens.LET, Value: "let"})

	parser.Backup()
	peek = parser.Peek()
	assert.NotNil(peek)
	assert.Equal(normalizeLexemes(peek), lexer.Lexeme{Token: tokens.VARNAME, Value: "name"})
}
