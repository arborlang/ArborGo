package wast

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type locals struct {
	name string
	tp   string
}

// Compiler traverses the AST and converts it to WASM
type Compiler struct {
	Writer       io.Writer
	SymbolTable  SymbolTable
	DeclLocation int
	level        int
	buffer       bytes.Buffer
	locals       []locals
}

// IsTopScope indicates if we are at the top level scope
func (c *Compiler) IsTopScope() bool {
	return c.level == 1
}

// getUniqueID gets a unique id for a thing
func (c *Compiler) getUniqueID(tp, name string) string {
	c.DeclLocation++
	return strings.Join([]string{"$" + tp, name, strconv.Itoa(c.DeclLocation)}, "_")
}

func (c *Compiler) getLabelUniqueID(tp string) string {
	loc := c.getUniqueID("label", tp)
	return fmt.Sprintf("$label%s", loc)
}

//StartModule starts the wat module
func (c *Compiler) StartModule() {
	c.level = 0
	c.Emit("(module")
	c.Emit(`(import "env" "__putch__" (func $__putch__ (param i32)))`)
}

//CloseModule ends the wat module
// The idea is that you would use this like so:
//		compiler.StartModule()
//		ast.Accept(compiler)
//		compiler.EndModule()
func (c *Compiler) CloseModule() {
	c.Emit(")")
	c.Flush()
}

// Emit emits the compiled instructions
func (c *Compiler) Emit(msg string, data ...interface{}) {
	instr := fmt.Sprintf(msg, data...)
	c.buffer.Write([]byte(fmt.Sprintf("%s\n", instr)))
	// c.Writer.Write([]byte(fmt.Sprintf("%s\n", instr)))
}

// Flush flushes the buffer
func (c *Compiler) Flush() {
	c.buffer.WriteTo(c.Writer)
	c.buffer.Reset()
}

// PrependAndFlush prepends sectiopn to the writer
func (c *Compiler) PrependAndFlush(section []byte) {
	c.Writer.Write(section)
	c.buffer.WriteTo(c.Writer)
	c.buffer.Reset()
}

// Clone clones the compiler with a different writer
func (c *Compiler) Clone(w io.Writer) *Compiler {
	return &Compiler{
		Writer:       w,
		SymbolTable:  c.SymbolTable,
		DeclLocation: c.DeclLocation,
		level:        c.level,
		buffer:       c.buffer,
	}
}

func (c *Compiler) getType(tp string) string {
	switch tp {
	case "char", "bool":
		return "i32"
	case "float":
		return "f64"
	default:
		return "i64"
	}
}
