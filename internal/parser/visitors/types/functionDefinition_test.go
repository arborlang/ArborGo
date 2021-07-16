package typevisitor

import (
	"strings"
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/rulesv2"
	"github.com/stretchr/testify/assert"
)

func TestFunctionDef(t *testing.T) {
	assert := assert.New(t)

	program := `
	fn someTest(val: Number) -> Number {
		return val;
	};
	`

	node, err := rulesv2.Parse(strings.NewReader(program))
	assert.NoError(err)
	assert.NotNil(node)
	tVisit := New()
	n, e := node.Accept(tVisit)
	assert.NoError(e)
	assert.NotNil(n)
}

func TestFunctionDefFailsIfTypeNotDefined(t *testing.T) {
	assert := assert.New(t)

	program := `
	fn someTest(val: Foobar) -> Number {
		return val;
	};
	`

	node, err := rulesv2.Parse(strings.NewReader(program))
	assert.NoError(err)
	assert.NotNil(node)
	tVisit := New()
	_, e := node.Accept(tVisit)
	assert.Error(e)
}
