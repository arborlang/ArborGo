package arborbuild

import (
	"io"

	"github.com/arborlang/ArborGo/internal/parser"
	"github.com/arborlang/ArborGo/internal/parser/visitors/WatCompiler"
)

// BuildToWast builds an Arbor file to web assembly text format
func BuildToWast(data io.Reader, w io.Writer) error {
	comp := wast.New(w)
	defer comp.CloseModule()
	if err := parser.Compile(data, comp); err != nil {
		return err
	}
	return nil
}
