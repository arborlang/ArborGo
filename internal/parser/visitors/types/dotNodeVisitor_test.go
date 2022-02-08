package typevisitor

import (
	"strings"
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/rulesv2"
	"github.com/stretchr/testify/assert"
)

func TestCanVisitDotNodeFine(t *testing.T) {
	assert := assert.New(t)

	testProg := `type Foobar {
		Name: String;
		Value: Number;	
	};
	
	fn SomeFunc(x: Foobar) -> String {
		return x.Name;	
	};`

	node, err := rulesv2.Parse(strings.NewReader(testProg))
	assert.NoError(err)
	assert.NotNil(node)
	tVisit := New(true)
	_, e := node.Accept(tVisit)
	assert.NoError(e)
}
