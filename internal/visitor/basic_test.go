package visitor

import (
	"testing"

	"github.com/radding/ArborGo/internal/ast"
	"github.com/stretchr/testify/assert"
)

func TestBasicImplementsVisitor(t *testing.T) {
	assert := assert.New(t)

	assert.Implements((*ast.Visitor)(nil), new(CompilerVisitor))
}
