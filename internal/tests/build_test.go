package tests

import (
	"log"
	"os"
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/parser/rulesv2"
	typevisitor "github.com/arborlang/ArborGo/internal/parser/visitors/types"
	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	assert := assert.New(t)
	fName := "../../docs/example/example.ab"
	f, err := os.Open(fName)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	var node ast.Node

	t.Run("Assert that Example File can be parsed", func(t *testing.T) {
		node, err = rulesv2.Parse(f)

		assert.NoError(err)
		assert.NotNil(node)
	})

	t.Run("Assert that AST can be walked", func(t *testing.T) {
		var visitors []ast.Visitor = []ast.Visitor{
			typevisitor.New(false),
		}

		for _, visitor := range visitors {
			node, err = node.Accept(visitor)
			assert.NoError(err)
			assert.NotNil(node)
		}
	})
}
