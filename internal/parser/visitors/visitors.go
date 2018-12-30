package visitors

import (
	"github.com/radding/ArborGo/internal/parser/visitors/Compiler"
	"io"
)

// NewCompiler instantiates a brand new compiler
func NewCompiler(w io.Writer, filename string) *compiler.Compiler {
	w.Write([]byte(`(module`))
	return &compiler.Compiler{
		Writer:      w,
		SymbolTable: compiler.SymbolTable{},
	}
}
