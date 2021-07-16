package typevisitor

import (
	"strings"
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/rulesv2"
	"github.com/stretchr/testify/assert"
)

func TestFunctionCall(t *testing.T) {
	assert := assert.New(t)

	program := `
	fn someTest(val: Number) -> Number {
		return val;
	};
	someTest(1234);
	`

	node, err := rulesv2.Parse(strings.NewReader(program))
	assert.NoError(err)
	assert.NotNil(node)
	tVisit := New()
	n, e := node.Accept(tVisit)
	assert.NoError(e)
	assert.NotNil(n)
}

func TestFunctionCallFailOnParamMismatch(t *testing.T) {
	assert := assert.New(t)
	program := `
	fn someTest(val: Number) -> Number {
		return val;
	};
	someTest("1234");
	`

	node, err := rulesv2.Parse(strings.NewReader(program))
	assert.NoError(err)
	assert.NotNil(node)
	tVisit := New()
	_, e := node.Accept(tVisit)
	assert.Error(e)
	assert.Equal("\"someTest\" (Line: 5, Column: 16) can not be called, signatures don't match. fn (String) -> Number vs fn (Number) -> Number", e.Error())

	program = `
	fn someTest(val: Number) -> Number {
		return val;
	};
	someTest(1234, "1234");
	`

	node, err = rulesv2.Parse(strings.NewReader(program))
	assert.NoError(err)
	assert.NotNil(node)
	tVisit = New()
	_, e = node.Accept(tVisit)
	assert.Error(e)
	assert.Equal("\"someTest\" (Line: 5, Column: 16) can not be called, signatures don't match. fn (Number, String) -> Number vs fn (Number) -> Number", e.Error())
}

func TestFailsIfCallsNonCallable(t *testing.T) {
	assert := assert.New(t)
	program := `
	const someTest = 1234;
	someTest("1234");
	`

	node, err := rulesv2.Parse(strings.NewReader(program))
	assert.NoError(err)
	assert.NotNil(node)
	tVisit := New()
	_, e := node.Accept(tVisit)
	assert.Error(e)
	assert.Equal("Type Number is not callable \"someTest\" (Line: 3, Column: 16)", e.Error())
}
