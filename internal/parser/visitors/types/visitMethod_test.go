package typevisitor

import (
	"strings"
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/rulesv2"
	"github.com/stretchr/testify/assert"
)

func TestCanVisitMethodDefinitionProperly(t *testing.T) {
	assert := assert.New(t)
	tpDefs := `type Thing String;
	fn Thing::value() -> String {
		return "this is a test";
	};
	`
	node, err := rulesv2.Parse(strings.NewReader(tpDefs))
	assert.NoError(err)
	assert.NotNil(node)
	tVisit := New()
	n, e := node.Accept(tVisit)
	assert.NoError(e)
	assert.NotNil(n)
}