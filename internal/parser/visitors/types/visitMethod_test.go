package typevisitor

import (
	"strings"
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/ast/types"
	"github.com/arborlang/ArborGo/internal/parser/rulesv2"
	"github.com/arborlang/ArborGo/internal/parser/visitors/base"
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
	tVisit := New(true)
	n, e := node.Accept(tVisit)
	assert.NoError(e)
	assert.NotNil(n)

	real := tVisit.(*base.VisitorAdapter).Visitor.(*typeVisitor)
	data, _ := real.scope.LookupSymbolInAllScopes("Thing")
	assert.True(data.IsType)
	assert.IsType(&types.ExtendedType{}, data.Type.Type)
	dataTP := data.Type.Type.(*types.ExtendedType)
	assert.NotNil(dataTP.Shape.Fields["value"])
}
