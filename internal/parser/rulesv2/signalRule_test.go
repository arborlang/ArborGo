package rulesv2

import (
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/stretchr/testify/assert"
)

var normalFatalRaiseParse = `fatal new FatalError();`
var normalWarnRaiseParse = `warn new Warning();`
var normalSignalRaiseParse = `signal 1 + 3;`

func TestCanParseFatalSignals(t *testing.T) {
	assert := assert.New(t)

	p := parseTest(normalFatalRaiseParse)

	node, err := parseSignal(p)
	assert.NoError(err)
	assert.NotNil(node)
	signalNode, ok := node.(*ast.SignalNode)
	assert.True(ok)
	assert.NotNil(signalNode)
	assert.Equal("fatal", signalNode.Level)
	assert.NotNil(signalNode.ValueToRaise)
	instantiateNode, ok := signalNode.ValueToRaise.(*ast.InstantiateNode)
	assert.True(ok)
	assert.NotNil(instantiateNode)
}

func TestCanParseWarnSignals(t *testing.T) {
	assert := assert.New(t)

	p := parseTest(normalWarnRaiseParse)

	node, err := parseSignal(p)
	assert.NoError(err)
	assert.NotNil(node)
	signalNode, ok := node.(*ast.SignalNode)
	assert.True(ok)
	assert.NotNil(signalNode)
	assert.Equal("warn", signalNode.Level)
	assert.NotNil(signalNode.ValueToRaise)
	instantiateNode, ok := signalNode.ValueToRaise.(*ast.InstantiateNode)
	assert.True(ok)
	assert.NotNil(instantiateNode)
}

func TestCanParseSignals(t *testing.T) {
	assert := assert.New(t)

	p := parseTest(normalSignalRaiseParse)

	node, err := parseSignal(p)
	assert.NoError(err)
	assert.NotNil(node)
	signalNode, ok := node.(*ast.SignalNode)
	assert.True(ok)
	assert.NotNil(signalNode)
	assert.Equal("signal", signalNode.Level)
	assert.NotNil(signalNode.ValueToRaise)
	instantiateNode, ok := signalNode.ValueToRaise.(*ast.MathOpNode)
	assert.True(ok)
	assert.NotNil(instantiateNode)
}

func TestCanParseFullExpression(t *testing.T) {
	assert := assert.New(t)

	p := parseTest(`foobar = warning new Warning();`)

	node, err := ExpressionRule(p)
	assert.NoError(err)
	assert.NotNil(node)

}
