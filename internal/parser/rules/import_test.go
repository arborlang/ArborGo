package rules

import (
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/stretchr/testify/assert"
)

var import1 = `import "aa"`
var import2 = `import "aaa" as a`

func TestCanParseImport(t *testing.T) {
	assert := assert.New(t)
	p := createParser(import1)
	node, err := importRules(p)
	if !assert.NoError(err) {
		t.Fatal()
	}
	imp, _ := node.(*ast.ImportNode)
	if !assert.NotNil(imp, "can't convert node to an import Node") {
		t.Fatal()
	}
	assert.Equal(`"aa"`, imp.Source)

	p = createParser(import2)
	node, err = importRules(p)
	if !assert.NoError(err) {
		t.Fatal()
	}
	imp, _ = node.(*ast.ImportNode)
	if !assert.NotNil(imp, "can't convert node to an import Node") {
		t.Fatal()
	}
	assert.Equal(`"aaa"`, imp.Source)
	assert.Equal(`a`, imp.Name)
}
