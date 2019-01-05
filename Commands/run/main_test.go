package main

import (
	"github.com/radding/ArborGo/lib/plugins"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImplementsPluginInterface(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*plugins.Command)(nil), new(Run))
}
