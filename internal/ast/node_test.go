package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompilerImplementsNode(t *testing.T) {
	assert := assert.New(t)

	assert.Implements((*Node)(nil), new(Compiler))
}
