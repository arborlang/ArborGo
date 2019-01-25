package wast

import (
	// "fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
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
	c.stackPointer += 4
	c.EmitFunc("i32.const %d", c.stackPointer)
	loc, err := node.Varname.Accept(c)
	if err != nil {
		return ast.VisitorMetaData{}, err
	}
	c.EmitFunc("i32.const 1")
	c.EmitFunc("i32.sub")
	c.EmitFunc("i32.store")
	c.EmitFunc("get_local %s", loc.Location)
	// c.EmitFunc("call $__break__")
	// fmt.Println(node.Varname.Name, ":", node.Start, "->", node.End)
	return ast.VisitorMetaData{}, nil
}
