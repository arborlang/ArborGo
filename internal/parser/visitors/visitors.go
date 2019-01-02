package visitors

import (
	"github.com/radding/ArborGo/internal/parser/visitors/Compiler"
	"io"
)

// NewCompiler instantiates a brand new compiler
func NewCompiler(w io.Writer, filename string) *compiler.Compiler {
	comp := &compiler.Compiler{
		Writer:      w,
		SymbolTable: compiler.SymbolTable{},
	}

	comp.StartModule()
	return comp
}
