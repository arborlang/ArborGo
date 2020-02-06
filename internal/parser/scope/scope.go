package scope

import (
	"fmt"
	"github.com/arborlang/ArborGo/internal/parser/ast"
)

// SymbolData is the data related to the symbol.
type SymbolData struct {
	// Type is the type of the node
	Type ast.TypeNode
	// Location is where the symbol is stored in the program
	Location string
}

// Scope Represents the current scope in the language
type Scope map[string]*SymbolData

// SymbolTable is a comprehensive list of symbols currently in scope
type SymbolTable struct {
	scopeStack []Scope
}

// PushNewScope Pushes a new scope on to the ScopeStack
func (s *SymbolTable) PushNewScope() {
	newScope := []Scope{Scope{}}
	s.scopeStack = append(newScope, s.scopeStack...)
}

// PopScope pops the global state
func (s *SymbolTable) PopScope() error {
	if len(s.scopeStack) <= 1 {
		return fmt.Errorf("Invalid operation: about to pop global stack")
	}
	s.scopeStack = s.scopeStack[1:]
	return nil
}

// AddToScope adds the variable to the scope
func (s *SymbolTable) AddToScope(name string, sym *SymbolData) {
	s.scopeStack[0][name] = sym
}

// LookupSymbol looks up the symbol in our table. returns the symboldata and the scope level (0 is the current scope).
// if symbol is not found, returns nil and -1
func (s *SymbolTable) LookupSymbol(name string) (*SymbolData, int) {
	for lvl, scope := range s.scopeStack {
		if sym, ok := scope[name]; ok {
			return sym, lvl
		}
	}
	return nil, -1
}

// NewTable generates and returns a new symbole table
func NewTable() *SymbolTable {
	scope := &SymbolTable{}
	scope.scopeStack = []Scope{}
	scope.PushNewScope()
	return scope
}
