package base

import (
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/stretchr/testify/assert"
)

func TestImplementsVisitor(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*ast.Visitor)(nil), new(VisitorAdapter))
}
