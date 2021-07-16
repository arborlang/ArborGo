package typevisitor

import (
	"strings"
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/rulesv2"
	"github.com/stretchr/testify/assert"
)

func TestAddsToScope(t *testing.T) {
	assert := assert.New(t)
	assert.True(true)
	tpDefs := `
	const elem: Number;
	`
	node, err := rulesv2.Parse(strings.NewReader(tpDefs))
	assert.NoError(err)
	assert.NotNil(node)
	tVisit := New()
	n, e := node.Accept(tVisit)
	assert.NoError(e)
	assert.NotNil(n)
}

func TestCanGetOtherTypes(t *testing.T) {
	assert := assert.New(t)
	tpDefs := `
	type Foobar String;
	const elem: Foobar;
	`
	node, err := rulesv2.Parse(strings.NewReader(tpDefs))
	assert.NoError(err)
	assert.NotNil(node)
	tVisit := New()
	n, e := node.Accept(tVisit)
	assert.NoError(e)
	assert.NotNil(n)
}
