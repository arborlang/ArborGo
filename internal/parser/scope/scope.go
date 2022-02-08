package scope

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
	"github.com/olekukonko/tablewriter"
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
	// The Constructors of the Type
	Constructors []*types.FnType
	// IsType is weather this symbol is a type
	IsType bool
}

func (s *SymbolData) String() string {
	data, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
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
	lvlStack      levelStack
	pushOperation int
	scopesCanGrow bool
}

func (s *SymbolTable) renderSpecificTable(getLevel func(chan int)) []string {
	parts := []string{}
	lvlToRender := make(chan int)
	topOLvlStack := s.lvlStack.top()
	go getLevel(lvlToRender)
	for i := range lvlToRender {
		table := [][]string{}
		lvl := topOLvlStack[i]
		scope := s.scopeStack[lvl]
		parts = append(parts, fmt.Sprintf("Scope Level: %d", i))
		for key, val := range scope {
			table = append(table, []string{key, val.String()})
		}
		buf := new(bytes.Buffer)
		tableWriter := tablewriter.NewWriter(buf)
		tableWriter.SetHeader([]string{"Name", "Symbol"})
		tableWriter.AppendBulk(table)
		tableWriter.Render()
		parts = append(parts, buf.String())
	}

	return parts
}

func (s *SymbolTable) String() string {
	parts := []string{}
	parts = append(parts, fmt.Sprintf(
		"scopeStackTop: %d. Scope Length: %d",
		s.GetCurrentLevel(),
		len(s.scopeStack)),
	)
	parts = append(parts, "VISIBLE TABLE:")

	visibleParts := s.renderSpecificTable(func(channel chan int) {
		defer close(channel)
		topOLvlStack := s.lvlStack.top()
		fmt.Println("LEVELS", topOLvlStack)
		for i := topOLvlStack[len(topOLvlStack)-1]; i >= 0; i-- {
			channel <- topOLvlStack[i]
		}
	})

	parts = append(parts, visibleParts...)
	parts = append(parts, "FULL TABLE:")

	visibleParts = s.renderSpecificTable(func(channel chan int) {
		defer close(channel)
		for i := len(s.scopeStack) - 1; i >= 0; i-- {
			channel <- i
		}
	})
	parts = append(parts, visibleParts...)
	return strings.Join(parts, "\n")
}

// PushNewScope Pushes a new scope on to the ScopeStack
func (s *SymbolTable) PushNewScope() {
	if s.scopesCanGrow == true {
		s.scopeStack = append(s.scopeStack, Scope{})
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
	currentLevel := s.GetCurrentLevel()
	s.scopeStack[currentLevel][name] = sym
}

// LookupSymbol looks up the symbol in our table. returns the symboldata and the scope level (0 is the current scope).
// if symbol is not found, returns nil and -1
func (s *SymbolTable) LookupSymbol(name string) (*SymbolData, int) {
	for i := len(s.lvlStack.top()) - 1; i >= 0; i-- {
		lvl := s.lvlStack.top()[i]
		if sym, ok := s.scopeStack[lvl][name]; ok {
			return sym, lvl
		}
	}
	return nil, -1
}

// GetCurrentLevel gets the current level of scope
func (s *SymbolTable) GetCurrentLevel() int {
	return s.lvlStack.top()[len(s.lvlStack.top())-1]
}

// ResetStackAndLockScope resets the stack so we are looking at the global scope again, and then prevents the scope from growing again
func (s *SymbolTable) ResetStackAndLockScope() {
	s.lvlStack.values = [][]int{{0}}
	s.pushOperation = 1
	s.scopesCanGrow = false
}

func (s *SymbolTable) LookupSymbolInAllScopes(name string) (*SymbolData, int) {
	for i, data := range s.scopeStack {
		info, ok := data[name]
		if ok {
			return info, i
		}
	}
	return nil, -1
}

// NewTable generates and returns a new symbole table
func NewTable() *SymbolTable {
	scope := &SymbolTable{}
	scope.scopeStack = []Scope{}
	scope.pushOperation = 0
	scope.scopesCanGrow = true
	scope.lvlStack = levelStack{
		values: [][]int{},
	}
	scope.PushNewScope()
	return scope
}

func (s *SymbolTable) ResolveType(sym *SymbolData) *SymbolData {
	if sym == nil {
		return sym
	}
	tp := sym.Type.Type
	gd, ok := sym.Type.Type.(*types.TypeGuard)
	if ok {
		if len(gd.Types) > 1 || len(gd.Types) == 0 {
			return sym
		}
		tp = gd.Types[0]
	}
	constTp, ok := tp.(*types.ConstantTypeNode)
	if ok {
		if sym.IsType {
			return sym
		}
		otherSym, _ := s.LookupSymbol(constTp.Name)
		return s.ResolveType(otherSym)
	}
	return sym
}
