package compiler

import (
	"fmt"
)

// Symbol defines a reference
type Symbol struct {
	Name       string
	Type       string
	Location   string
	IsConstant bool
}

// Scope defines a scope
type Scope map[string]Symbol

// SymbolTable defines a way to look up symbols in a given file
type SymbolTable struct {
	fileName     string
	currentScope Scope
	scopeStack   []Scope
	useGlobal    bool
	// GlobalScope  Scope
}

// PushScope adds a new scope to the symbol table
func (s *SymbolTable) PushScope() error {
	// s.useGlobal = false
	oldScope := s.currentScope
	newScope := make(Scope)
	s.scopeStack = append(s.scopeStack, oldScope)
	s.currentScope = newScope
	return nil
}

// PopScope removes the last scope from the list
func (s *SymbolTable) PopScope() error {
	if len(s.scopeStack) == 1 {
		s.useGlobal = true
		s.currentScope = s.scopeStack[0]
		s.scopeStack = []Scope{}
		return nil
	} else if len(s.scopeStack) == 0 {
		return fmt.Errorf("can not pop non exsistent scope")
	}
	lastScope, stack := s.scopeStack[len(s.scopeStack)-1], s.scopeStack[:len(s.scopeStack)-1]
	s.currentScope = lastScope
	s.scopeStack = stack
	return nil
}

// AddToScope adds a symbol to the current scope
func (s *SymbolTable) AddToScope(sy Symbol) error {
	// if s.useGlobal {
	// 	s.GlobalScope[sy.Name] = sy
	// } else {
	s.currentScope[sy.Name] = sy
	// }
	return nil
}

// GetSymbol looks up a symbol in the symbol table and returns a pointer to it if it finds it, else return none
func (s *SymbolTable) GetSymbol(name string) *Symbol {
	// if s.useGlobal {
	// 	sym, ok := s.GlobalScope[name]
	// 	if !ok {
	// 		return nil
	// 	}
	// 	return &sym
	// }
	sym, ok := s.currentScope[name]
	if !ok {
		for i := len(s.scopeStack) - 1; i >= 0; i-- {
			sym, ok := s.scopeStack[i][name]
			if ok {
				return &sym
			}
		}
		return nil
	}
	return &sym
}
