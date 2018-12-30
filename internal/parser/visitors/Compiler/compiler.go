package compiler

import (
	"io"
	"strconv"
	"strings"
)

// Compiler traverses the AST and converts it to WASM
type Compiler struct {
	Writer       io.Writer
	SymbolTable  SymbolTable
	DeclLocation int
}

// getUniqueID gets a unique id for a thing
func (c *Compiler) getUniqueID(tp, name string) string {
	c.DeclLocation++
	return strings.Join([]string{tp, name, strconv.Itoa(c.DeclLocation)}, "")
}

//StartModule starts the wat module
func (c *Compiler) StartModule() {
	c.Writer.Write([]byte("(module"))
}

//CloseModule ends the wat module
// The idea is that you would use this like so:
//		compiler.StartModule()
//		ast.Accept(compiler)
//		compiler.EndModule()
func (c *Compiler) CloseModule() {
	c.Writer.Write([]byte(")"))
}
