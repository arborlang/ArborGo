package visitors

import (
	"github.com/radding/ArborGo/internal/parser/visitors/WatCompiler"
	"io"
)

// NewWastCompiler instantiates a brand new compiler
func NewWastCompiler(w io.Writer, filename string) *wast.Compiler {
	return wast.New(w)
}
