package typevisitor

import (
	"strings"
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
	"github.com/arborlang/ArborGo/internal/parser/rulesv2"
	"github.com/arborlang/ArborGo/internal/parser/visitors/base"
	"github.com/stretchr/testify/require"
)

func TestVisitsTypeDefCorrectly(t *testing.T) {
	assert := require.New(t)

	tpDefs := `type Thing String;
	type Thing2 {
		Method: fn () -> String;
		SomeProp: String;
	};
	`
	node, err := rulesv2.Parse(strings.NewReader(tpDefs))
	assert.NoError(err)
	assert.NotNil(node)

	tVisit := New(true)
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
	scopedType, _ := realVisitor.scope.LookupSymbol("Thing")
	assert.NotNil(scopedType)
	assert.IsType(&types.ExtendedType{}, scopedType.Type.Type)
	scopedType, _ = realVisitor.scope.LookupSymbol("Thing2")
	assert.NotNil(scopedType)
	assert.IsType(&types.ShapeType{}, scopedType.Type.Type)
	shp := scopedType.Type.Type.(*types.ShapeType)
	assert.Contains(shp.Fields, "Method")
	// assert.IsType(&types.FnType{}, shp.Fields["Method"])
}

func TestExplicitExtends(t *testing.T) {
	assert := require.New(t)
	tpDefs := `
	type Thing extends String{
		Method: fn () -> String;
		SomeProp: String;
	};
	`

	node, err := rulesv2.Parse(strings.NewReader(tpDefs))
	assert.NoError(err)
	assert.NotNil(node)

	tVisit := New(true)
	node, err = node.Accept(tVisit)
	assert.NoError(err)

	realVisitor := tVisit.(*base.VisitorAdapter).Visitor.(*typeVisitor)
	scopedType, _ := realVisitor.scope.LookupSymbolInAllScopes("Thing")

	assert.NotNil(scopedType)
	assert.IsType(&types.ExtendedType{}, scopedType.Type.Type)
}
