package typevisitor

import (
	"strings"
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
	"github.com/arborlang/ArborGo/internal/parser/rulesv2"
	"github.com/arborlang/ArborGo/internal/parser/visitors/base"
	"github.com/stretchr/testify/assert"
)

func TestVisitsTypeDefCorrectly(t *testing.T) {
	assert := assert.New(t)
	tpDefs := `type Thing String;
	type Thing2 {
		Method: fn () -> String;
		SomeProp: String;
	};
	`
	node, err := rulesv2.Parse(strings.NewReader(tpDefs))
	assert.NoError(err)
	assert.NotNil(node)

	tVisit := New()
	node, err = node.Accept(tVisit)
	assert.NoError(err)
	assert.NotNil(node)
	pkg, ok := node.(*ast.Program)
	assert.True(ok)
	first := pkg.Nodes[0]
	_, ok = first.(*ast.ExtendsNode)
	assert.True(ok)
	second := pkg.Nodes[1]
	_, ok = second.(*ast.TypeNode)
	assert.True(ok)

	realVisitor := tVisit.(*base.VisitorAdapter).Visitor.(*typeVisitor)
	scopedType, _ := realVisitor.scope.LookupType("Thing")
	assert.NotNil(scopedType)
	assert.IsType(&types.ExtendedType{}, scopedType.Type)
	scopedType, _ = realVisitor.scope.LookupType("Thing2")
	assert.NotNil(scopedType)
	assert.IsType(&types.ShapeType{}, scopedType.Type)
}
