package scope

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewScope(t *testing.T) {
	assert := assert.New(t)
	symTable := NewTable()
	assert.NotNil(symTable)
	assert.Len(symTable.scopeStack, 1)
}

func TestPushingAndPoppingScope(t *testing.T) {
	assert := assert.New(t)

	symTable := NewTable()
	err := symTable.PopScope()
	assert.Error(err)
	symTable.PushNewScope()
	symTable.PushNewScope()
	assert.Len(symTable.scopeStack, 3)
	err = symTable.PopScope()
	assert.NoError(err)
	assert.Len(symTable.scopeStack, 2)
	err = symTable.PopScope()
	assert.NoError(err)
	err = symTable.PopScope()
	assert.Error(err)
}

func TestSymbolLookUp(t *testing.T) {
	assert := assert.New(t)

	symTable := NewTable()
	sym := &SymbolData{Location: "xyz"}
	symTable.AddToScope("foo", sym)
	symTable.PushNewScope()
	sym3 := &SymbolData{Location: "abc"}
	symTable.AddToScope("bar", sym3)

	sym2, lvl := symTable.LookupSymbol("foo")
	assert.NotNil(sym2)
	assert.Equal(sym, sym2)
	assert.Equal(lvl, 1)

	sym2, lvl = symTable.LookupSymbol("bar")
	assert.NotNil(sym2)
	assert.Equal(sym3, sym2)
	assert.Equal(lvl, 0)

	sym2, lvl = symTable.LookupSymbol("barz")
	assert.Nil(sym2)
	assert.Equal(lvl, -1)
}
