package functions

import (
	"bytes"
	"github.com/radding/ArborGo/internal/lexer"
	"github.com/radding/ArborGo/internal/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

var funcs = `
let x = (a: number, b: number) -> a + b;
const y = () -> {
	const xyz = a + b;
	let z = () -> xyz * 3;
	return y() + xyz;
};
const foo = (a: number, b: number,  c: number) -> {
	return a() + b(c);
};
`

func TestCanGetAllFucntionsInClass(t *testing.T) {
	assert := assert.New(t)
	tokStream := lexer.Lex(bytes.NewReader([]byte(funcs)))
	prog, err := parser.Parse(tokStream)
	if err != nil {
		t.Fatal(err)
	}
	funcVisitor := New()
	_, err = prog.Accept(funcVisitor)
	if !assert.NoError(err) {
		t.Fatal()
	}
	visitor := funcVisitor.GetVisitor().(*FunctionAnalyzer)
	assert.Contains(visitor.functions, "x")
	assert.Contains(visitor.functions, "foo")
	assert.Contains(visitor.functions, "y")
	assert.NotContains(visitor.functions, "z")
	assert.Equal(visitor.functions["x"].IsConstant, false)
	assert.Equal(visitor.functions["x"].Arguments[0].Name, "a")
	assert.Equal(visitor.functions["x"].Arguments[0].Types[0], "number")
	assert.Equal(visitor.functions["x"].Arguments[1].Name, "b")
	assert.Equal(visitor.functions["x"].Arguments[1].Types[0], "number")
}
