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
	assert.Equal(symTable.lvlStack.values, [][]int{{0}})
}

func TestPushingAndPoppingScope(t *testing.T) {
	assert := assert.New(t)

	symTable := NewTable()
	err := symTable.PopScope()
	assert.Error(err)
	symTable.PushNewScope()
	symTable.PushNewScope()
	assert.Len(symTable.scopeStack, 3)
	assert.Equal(2, symTable.lvlStack.top()[len(symTable.lvlStack.top())-1])
	err = symTable.PopScope()
	assert.Equal(1, symTable.lvlStack.top()[len(symTable.lvlStack.top())-1])
	assert.NoError(err)
	assert.Len(symTable.scopeStack, 3)
	err = symTable.PopScope()
	assert.Equal(0, symTable.lvlStack.top()[len(symTable.lvlStack.top())-1])
	assert.NoError(err)
	err = symTable.PopScope()
	assert.Equal(0, symTable.lvlStack.top()[len(symTable.lvlStack.top())-1])
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
	assert.Equal(0, lvl)

	sym2, lvl = symTable.LookupSymbol("bar")
	assert.NotNil(sym2)
	assert.Equal(sym3, sym2)
	assert.Equal(1, lvl)

	sym2, lvl = symTable.LookupSymbol("barz")
	assert.Nil(sym2)
	assert.Equal(-1, lvl)
}

func TestPopAndPushScopeDoesntFallover(t *testing.T) {
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
	assert.Equal(0, lvl)

	sym2, lvl = symTable.LookupSymbol("bar")
	assert.NotNil(sym2)
	assert.Equal(sym3, sym2)

	symTable.PopScope()
	symTable.PushNewScope()

	sym4 := &SymbolData{Location: "abc1"}
	symTable.AddToScope("bar2", sym4)
	sym2, lvl = symTable.LookupSymbol("bar")
	assert.Nil(sym2)
	sym2, lvl = symTable.LookupSymbol("bar2")
	assert.NotNil(sym3)
	assert.NotEqual(sym3, sym2)
	assert.Equal(sym4, sym2)
	assert.Len(symTable.scopeStack, 3)
}

func TestResetAndLock(t *testing.T) {
	assert := assert.New(t)

	symTable := NewTable()
	sym := &SymbolData{Location: "xyz"}
	symTable.AddToScope("foo", sym)
	symTable.PushNewScope()
	sym3 := &SymbolData{Location: "abc"}
	symTable.AddToScope("bar", sym3)
	symTable.PopScope()
	symTable.PushNewScope()
	sym4 := &SymbolData{Location: "abc1"}
	symTable.AddToScope("bar2", sym4)

	assert.Len(symTable.scopeStack, 3)
	symTable.ResetStackAndLockScope()
	assert.False(symTable.scopesCanGrow)
	assert.Equal(1, symTable.pushOperation)

	sym2, lvl := symTable.LookupSymbol("foo")
	assert.NotNil(sym2)
	assert.Equal(sym, sym2)
	assert.Equal(0, lvl)
	sym2, lvl = symTable.LookupSymbol("bar")
	assert.Nil(sym2)
	sym2, lvl = symTable.LookupSymbol("bar2")
	assert.Nil(sym2)

	symTable.PushNewScope()
	assert.Len(symTable.scopeStack, 3)
	sym2, lvl = symTable.LookupSymbol("foo")
	assert.NotNil(sym2)
	assert.Equal(sym, sym2)
	sym2, lvl = symTable.LookupSymbol("bar")
	assert.NotNil(sym2)
	sym2, lvl = symTable.LookupSymbol("bar2")
	assert.Nil(sym2)

	symTable.PopScope()
	symTable.PushNewScope()

	sym2, lvl = symTable.LookupSymbol("bar")
	assert.Nil(sym2)
	sym2, lvl = symTable.LookupSymbol("bar2")
	assert.NotNil(sym2)
}

func TestStack(t *testing.T) {
	assert := assert.New(t)

	stack := levelStack{
		values: [][]int{{0}},
	}

	stack.push(1)
	assert.Equal(stack.values, [][]int{{0}, {0, 1}})
	stack.push(2)
	assert.Equal(stack.values, [][]int{{0}, {0, 1}, {0, 1, 2}})
	stack.pop()
	assert.Equal(stack.values, [][]int{{0}, {0, 1}})
	stack.push(3)
	assert.Equal(stack.values, [][]int{{0}, {0, 1}, {0, 1, 3}})

	assert.Equal(stack.top(), []int{0, 1, 3})
}
