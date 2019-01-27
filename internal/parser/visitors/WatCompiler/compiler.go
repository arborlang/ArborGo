package wast

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type data struct {
	name   string
	data   []byte
	len    int
	offSet int
}

func (d data) writeTo(w io.Writer) {
	str := []string{}
	for _, byt := range d.data {
		str = append(str, fmt.Sprintf("\\%X", byt))
	}
	out := fmt.Sprintf("(data (i32.const %d) \"%s\")\n", d.offSet, strings.Join(str, ""))

	w.Write([]byte(out))
}

// Compiler traverses the AST and converts it to WASM
type Compiler struct {
	Writer         io.Writer
	SymbolTable    SymbolTable
	DeclLocation   int
	level          int
	buffer         bytes.Buffer
	funcCount      int
	functions      []function
	nameToFunction map[string]function
	data           []data
	currentFunc    *function
	dataSize       int
	// stackPointer      int
	currentAssignment string
}

// AddData to the module
func (c *Compiler) AddData(name string, byteArr []byte) {
	c.data = append(c.data, data{name: name, data: byteArr, len: len(byteArr), offSet: c.dataSize})
	c.dataSize += len(byteArr)
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
	c.EmitFirst("(module")
	c.EmitFirst(`(import "env" "__putch__" (func $__putch__ (param i32)))`)
	c.EmitFirst(`(import "env" "__alloc__" (func $__alloc__ (param i64) (param i32) (result i64)))`)
	c.EmitFirst(`(import "env" "__break__" (func $__break__ (result i64)))`)
	c.EmitFirst(`(import "env" "__pushstack__" (func $__pushstack__ (result i32)))`)
	c.EmitFirst(`(import "env" "__popstack__" (func $__popstack__ (result i32)))`)
	c.EmitFirst(`(import "env" "__incrementstack__" (func $__allocstack__ (param i64) (result i32)))`)
	c.EmitFirst(`(import "env" "__stacktop__" (func $__stacktop__ (result i32)))`)
	c.EmitFirst(`(import "env" "STACKTOP_ASM" (global $__STACKTOP_IMPORT__ i32))`)
	c.EmitFirst(`(global  $__STACKTOP__ (mut i32) (get_global $__STACKTOP_IMPORT__))`)
	c.EmitFirst("(memory 1)")
	c.EmitFirst(`(func $__len__ (param $pointer i32) (result i64)
		get_local $pointer
		i64.load
	)`)
	// c.EmitFirst("(func $__stackPush__")
	c.SymbolTable.PushScope()
	c.SymbolTable.AddToScope(&Symbol{
		Name:     "len",
		Location: "$__len__",
		Type:     "number",
	})
}

//CloseModule ends the wat module
// The idea is that you would use this like so:
//		compiler.StartModule()
//		ast.Accept(compiler)
//		compiler.EndModule()
func (c *Compiler) CloseModule() {

	c.EmitFirst("(table %d anyfunc)", c.funcCount)
	for _, dat := range c.data {
		dat.writeTo(c.Writer)
	}
	for _, fun := range c.functions {
		fun.writeTo(c.Writer)
	}
	if c.currentFunc != nil {
		c.currentFunc.writeTo(c.Writer)
	}
	c.EmitFirst(")")
	c.Flush()
}

// EmitFirst emits the compiled instructions
func (c *Compiler) EmitFirst(msg string, data ...interface{}) {
	instr := fmt.Sprintf(msg, data...)
	c.Writer.Write([]byte(fmt.Sprintf("%s\n", instr)))
	// c.Writer.Write([]byte(fmt.Sprintf("%s\n", instr)))
}

// EmitFunc emits the compiled instructions
func (c *Compiler) EmitFunc(msg string, data ...interface{}) {
	instr := fmt.Sprintf(msg, data...)
	c.currentFunc.code.Write([]byte(fmt.Sprintf("%s\n", instr)))
}

//New constructs a new compiler
func New(w io.Writer) *Compiler {
	comp := &Compiler{
		Writer:         w,
		SymbolTable:    SymbolTable{},
		nameToFunction: map[string]function{},
	}

	comp.StartModule()
	return comp
}

// StartFunc is a function
func (c *Compiler) StartFunc() {
	if c.currentFunc == nil {
		c.currentFunc = &function{
			code:     bytes.NewBuffer([]byte{}),
			locals:   []locals{},
			tblIndex: c.funcCount,
		}
		c.funcCount++
		return
	}
	c.funcCount++
	c.functions = append(c.functions, *c.currentFunc)
	c.nameToFunction[c.currentFunc.mangle] = *c.currentFunc
	c.currentFunc = &function{
		code:     bytes.NewBuffer([]byte{}),
		locals:   []locals{},
		tblIndex: c.funcCount,
	}
}

// AddLocal adds a local variable
func (c *Compiler) AddLocal(name string, tp string) {
	c.currentFunc.locals = append(c.currentFunc.locals, locals{name: name, tp: tp})
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
