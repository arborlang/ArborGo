package typevisitor

import (
	"strings"
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/rulesv2"
	"github.com/arborlang/ArborGo/internal/parser/visitors/base"
	"github.com/stretchr/testify/assert"
)

func TestCanDeriveTypes(t *testing.T) {
	assert := assert.New(t)

	tpDefs := `
	const elem = 1234;
	const elem2 = "1234";
	`
	node, err := rulesv2.Parse(strings.NewReader(tpDefs))

	assert.NoError(err)
	assert.NotNil(node)

	tVisit := New()
	_, e := node.Accept(tVisit)
	assert.NoError(e)

	real := tVisit.(*base.VisitorAdapter).Visitor.(*typeVisitor)
	data, _ := real.scope.LookupSymbolInAllScopes("elem")
	integer, _ := real.scope.LookupSymbolInAllScopes("Number")
	assert.NotNil(data)
	assert.True(data.Type.Type.IsSatisfiedBy(integer.Type.Type))

	data, _ = real.scope.LookupSymbolInAllScopes("elem2")
	str, _ := real.scope.LookupSymbolInAllScopes("String")
	assert.NotNil(data)
	assert.True(data.Type.Type.IsSatisfiedBy(str.Type.Type))
}

func TestFailsIfRedefined(t *testing.T) {
	assert := assert.New(t)

	tpDefs := `
		const elem = 1234;
		const elem = "1234";
	`
	node, err := rulesv2.Parse(strings.NewReader(tpDefs))

	assert.NoError(err)
	assert.NotNil(node)

	tVisit := New()
	_, e := node.Accept(tVisit)
	assert.Error(e)
	assert.Equal("elem is being redefined here: \"elem\" (Line: 3, Column: 24)", e.Error())
}

func TestFailsIfAssigningDifferentTypes(t *testing.T) {
	assert := assert.New(t)
	tpDefs := `
		let elem = 1234;
		elem = "1234";
	`
	node, err := rulesv2.Parse(strings.NewReader(tpDefs))

	assert.NoError(err)
	assert.NotNil(node)

	tVisit := New()
	_, e := node.Accept(tVisit)
	assert.Error(e)
	assert.Equal("can't assign String to Number at \"=\" (Line: 3, Column: 19)", e.Error())
}

func TestFailsIfAssigningToConst(t *testing.T) {
	assert := assert.New(t)

	tpDefs := `
		const elem = 1234;
		elem = 4;
	`
	node, err := rulesv2.Parse(strings.NewReader(tpDefs))

	assert.NoError(err)
	assert.NotNil(node)

	tVisit := New()
	_, e := node.Accept(tVisit)
	assert.Error(e)
	assert.Equal("elem is constant: defined here: \"elem\" (Line: 2, Column: 24)", e.Error())
}
