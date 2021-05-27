package scope

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

// SymbolData is the data related to the symbol.
type SymbolData struct {
	// Type is the type of the node
	Type TypeData
	// Location is where the symbol is stored in the program
	Location string
	// Lexeme is the lexeme where this type was defined
	Lexeme lexer.Lexeme
	// IsConstant denotes if this symbol can be reassigned
	IsConstant bool
	// Methods are the methods on the type
	Methods map[string][]*types.FnType
}

// TypeData represents some information about a type
type TypeData struct {
	IsSealed bool
	Adds     []types.TypeNode
	Name     string
	Type     types.TypeNode
}

// IsSatisfiedBy checks if the symbol is satisfied by this type
func (t *TypeData) IsSatisfiedBy(tp types.TypeNode) bool {
	return t.Type.IsSatisfiedBy(tp)
}

// IsSatisfiedBy checks if the symbol is satisfied by this type
func (s *SymbolData) IsSatisfiedBy(t types.TypeNode) bool {
	return s.Type.IsSatisfiedBy(t)
}

// Scope Represents the current scope in the language
type Scope map[string]*SymbolData
type TypeScope map[string]*TypeData

type levelStack struct {
	values [][]int
}

func (l *levelStack) push(value int) {
	vals := []int{value}
	if len(l.values)-1 >= 0 {
		vals = append(l.values[len(l.values)-1], value)
	}
	l.values = append(l.values, vals)
}

func (l *levelStack) pop() {
	l.values = l.values[:len(l.values)-1]
}

func (l *levelStack) top() []int {
	return l.values[len(l.values)-1]
}

// SymbolTable is a comprehensive list of symbols currently in scope
type SymbolTable struct {
	scopeStack    []Scope
	typeStack     []TypeScope
	lvlStack      levelStack
	pushOperation int
	scopesCanGrow bool
}

// PushNewScope Pushes a new scope on to the ScopeStack
func (s *SymbolTable) PushNewScope() {
	if s.scopesCanGrow == true {
		s.scopeStack = append(s.scopeStack, Scope{})
		s.typeStack = append(s.typeStack, TypeScope{})
	}
	s.lvlStack.push(s.pushOperation)
	s.pushOperation++
}

// PopScope pops the global state
func (s *SymbolTable) PopScope() error {
	if len(s.lvlStack.values) == 1 {
		return fmt.Errorf("Invalid operation: about to pop global stack")
	}
	s.lvlStack.pop()
	return nil
}

// AddToScope adds the variable to the scope
func (s *SymbolTable) AddToScope(name string, sym *SymbolData) {
	currentLevel := s.lvlStack.top()[len(s.lvlStack.top())-1]
	s.scopeStack[currentLevel][name] = sym
}

func (s *SymbolTable) AddType(name string, sym *TypeData) {
	currentLevel := s.lvlStack.top()[len(s.lvlStack.top())-1]
	s.typeStack[currentLevel][name] = sym
}

// LookupSymbol looks up the symbol in our table. returns the symboldata and the scope level (0 is the current scope).
// if symbol is not found, returns nil and -1
func (s *SymbolTable) LookupSymbol(name string) (*SymbolData, int) {
	for i := len(s.lvlStack.top()) - 1; i >= 0; i-- {
		lvl := s.lvlStack.top()[i]
		fmt.Println(s.scopeStack[lvl])
		if sym, ok := s.scopeStack[lvl][name]; ok {
			return sym, lvl
		}
	}
	return nil, -1
}

// LookupType looks up the symbol in our table. returns the symboldata and the scope level (0 is the current scope).
// if symbol is not found, returns nil and -1
func (s *SymbolTable) LookupType(name string) (*TypeData, int) {
	for i := len(s.lvlStack.top()) - 1; i >= 0; i-- {
		lvl := s.lvlStack.top()[i]
		if sym, ok := s.typeStack[lvl][name]; ok {
			return sym, lvl
		}
	}
	return nil, -1
}

func (s *SymbolTable) getCurrentLevel() int {
	return s.lvlStack.top()[len(s.lvlStack.top())-1]
}

func (s *SymbolTable) ResetStackAndLockScope() {
	s.lvlStack.values = [][]int{{0}}
	s.pushOperation = 0
	s.scopesCanGrow = false
}

// NewTable generates and returns a new symbole table
func NewTable() *SymbolTable {
	scope := &SymbolTable{}
	scope.scopeStack = []Scope{}
	scope.typeStack = []TypeScope{}
	scope.pushOperation = 0
	scope.scopesCanGrow = true
	scope.lvlStack = levelStack{
		values: [][]int{},
	}
	scope.PushNewScope()
	return scope
}
