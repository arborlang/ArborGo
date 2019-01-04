package compiler

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Compiler traverses the AST and converts it to WASM
type Compiler struct {
	Writer       io.Writer
	SymbolTable  SymbolTable
	DeclLocation int
	Level        int
}

// getUniqueID gets a unique id for a thing
func (c *Compiler) getUniqueID(tp, name string) string {
	c.DeclLocation++
	return strings.Join([]string{"$" + tp, name, strconv.Itoa(c.DeclLocation)}, "_")
}

//StartModule starts the wat module
func (c *Compiler) StartModule() {
	c.Level = 0
	c.Emit("(module")
}

//CloseModule ends the wat module
// The idea is that you would use this like so:
//		compiler.StartModule()
//		ast.Accept(compiler)
//		compiler.EndModule()
func (c *Compiler) CloseModule() {
	c.Emit(")")
}

// Emit emits thecompiled instructions
func (c *Compiler) Emit(msg string, data ...interface{}) {
	instr := fmt.Sprintf(msg, data...)
	c.Writer.Write([]byte(fmt.Sprintf("%s\n", instr)))
}
