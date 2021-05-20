package rulesv2

import (
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/stretchr/testify/assert"
)

func TestCanExportThing(t *testing.T) {
	assert := assert.New(t)
	testExportAssignment := `internal const Foobar = 1;`
	p := parseTest(testExportAssignment)
	exportNode, err := parseExport(p)
	assert.NoError(err)
	assert.NotNil(exportNode)
	if !assert.IsType(&ast.InternalNode{}, exportNode) {
		return
	}
	exp := exportNode.(*ast.InternalNode)
	assert.IsType(&ast.AssignmentNode{}, exp.Expression)
}

func TestCanExportVarname(t *testing.T) {
	assert := assert.New(t)
	testExportAssignment := `internal Foobar;`
	p := parseTest(testExportAssignment)
	exportNode, err := parseExport(p)
	assert.NoError(err)
	assert.NotNil(exportNode)
	if !assert.IsType(&ast.InternalNode{}, exportNode) {
		return
	}
	exp := exportNode.(*ast.InternalNode)
	assert.IsType(&ast.VarName{}, exp.Expression)
}

func TestCanExportFunction(t *testing.T) {
	assert := assert.New(t)
	testExportAssignment := `internal fn Foobar (value: string) {
		return 0;	
	}`
	p := parseTest(testExportAssignment)
	exportNode, err := parseExport(p)
	assert.NoError(err)
	assert.NotNil(exportNode)
	if !assert.IsType(&ast.InternalNode{}, exportNode) {
		return
	}
	exp := exportNode.(*ast.InternalNode)
	assert.IsType(&ast.AssignmentNode{}, exp.Expression)
}
