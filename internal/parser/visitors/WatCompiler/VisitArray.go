package wast

import (
	// "fmt"
	"github.com/arborlang/ArborGo/internal/parser/ast"
)

// VisitIndexNode visists a node that gets the index of an element in an array
func (c *Compiler) VisitIndexNode(node *ast.IndexNode) (ast.VisitorMetaData, error) {
	_, err := node.Varname.Accept(c)
	c.EmitFunc(";;Start to load an index")
	if err != nil {
		return ast.VisitorMetaData{}, err
	}
	c.EmitFunc("i32.const %d", (node.Index+1)*4)
	c.EmitFunc("i32.add")
	c.EmitFunc("i32.load")
	c.EmitFunc(";;end to load an index")
	return ast.VisitorMetaData{}, nil
}

//VisitSliceNode visits a node as a slice
func (c *Compiler) VisitSliceNode(node *ast.SliceNode) (ast.VisitorMetaData, error) {
	c.EmitFunc(";; Slicing an array")
	// c.stackPointer += 4
	// location := c.stackPointer
	localName := c.getUniqueID("stack", "pointer")
	c.AddLocal(localName, "i32")
	c.EmitFunc("i64.const 4")
	c.EmitFunc("call $__allocstack__")
	c.EmitFunc("set_local %s", localName)
	c.EmitFunc("get_local %s", localName)
	// c.EmitFunc("i32.const %d", location)
	_, err := node.Varname.Accept(c)
	if err != nil {
		return ast.VisitorMetaData{}, err
	}
	c.EmitFunc("i32.load")
	c.EmitFunc("i32.const 1")
	c.EmitFunc("i32.sub")
	c.EmitFunc("i32.store")
	// c.stackPointer += 4
	// c.EmitFunc("i32.const %d", location)
	// c.EmitFunc("call $__")
	c.EmitFunc("get_local %s", localName)
	c.EmitFunc(";; Done slicing an array")
	return ast.VisitorMetaData{}, nil
}
