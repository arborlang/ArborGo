package typevisitor

import (
	"strings"
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/rulesv2"
	"github.com/arborlang/ArborGo/internal/parser/visitors/base"
	"github.com/stretchr/testify/assert"
)

func TestOperationsWorks(t *testing.T) {
	assert := assert.New(t)

	prog := `
		const elem = 1234 + 1234;
		1234 - 12;
		123 * 2;
		144 / 12;
	`

	node, err := rulesv2.Parse(strings.NewReader(prog))

	assert.NoError(err)
	assert.NotNil(node)

	tVisit := New(true)
	_, e := node.Accept(tVisit)
	assert.NoError(e)

	real := tVisit.(*base.VisitorAdapter).Visitor.(*typeVisitor)
	data, _ := real.scope.LookupSymbolInAllScopes("elem")

	integer, _ := real.scope.LookupSymbolInAllScopes("Number")
	assert.NotNil(data)
	assert.True(data.Type.Type.IsSatisfiedBy(integer.Type.Type))
}

func TestMathbetweenTwoTypesFails(t *testing.T) {
	assert := assert.New(t)

	prog := `
		const elem = 1234 + "1234";
	`

	node, err := rulesv2.Parse(strings.NewReader(prog))

	assert.NoError(err)
	assert.NotNil(node)

	tVisit := New(true)
	_, e := node.Accept(tVisit)
	assert.Error(e)
	assert.Equal("can't add Number and String (at \"+\" (Line: 2, Column: 52))", e.Error())
}
