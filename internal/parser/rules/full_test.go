package rules

import (
	"testing"

	"github.com/arborlang/ArborGo/internal/tokens"
	"github.com/stretchr/testify/assert"
)

var potentialProgram = `
const foo = (a: number, b: number) -> a + b;
const xyz = () -> {
	const a = 1;
	const b = 2;
	return foo(a, b);
};

let a = 1;
let b = 2;
const c = xyz() + foo(1, 2);
let d = "asbfasf";
if a < b {
	const x = 3;
	foo(x, b);
} else if a > b {
	const x = 3;
	foo(x, a);
} else {
	foo(a, b);
}
`

func TestCanParseProgram(t *testing.T) {
	assert := assert.New(t)
	p := createParser(potentialProgram)

	prog, err := ProgramRule(p, tokens.EOF)
	assert.NoError(err)
	assert.NotNil(prog)
}
