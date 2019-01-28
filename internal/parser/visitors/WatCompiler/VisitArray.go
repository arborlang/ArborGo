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
	arrayPosition := c.getUniqueID("current_array", "position")
	arraySize := c.getUniqueID("current_array", "size")
	c.AddLocal(arrayPosition, "i32")
	c.AddLocal(arraySize, "i32")
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
	c.EmitFunc("tee_local %s", arrayPosition)
	c.EmitFunc("i32.load")
	c.EmitFunc("tee_local %s", arraySize)
	c.EmitFunc("i32.const 1")
	c.EmitFunc("i32.sub")
	c.EmitFunc("i32.store")

	loop := c.getLabelUniqueID("loop")
	loopEnd := c.getLabelUniqueID("loopEnd")
	// Start the block for the loop
	c.EmitFunc("block %s", loop)
	c.EmitFunc("loop %s", loopEnd)

	// Get the array size. If it is equal to zero, break out of the loop
	c.EmitFunc("get_local %s", arraySize)
	c.EmitFunc("i32.eqz")
	c.EmitFunc("br_if %s", loop)

	// First add four to the arrayPostition
	c.EmitFunc("i32.const 4")
	c.EmitFunc("get_local %s", arrayPosition)
	c.EmitFunc("i32.add")

	// Store the new pointer, but also we need to load it
	c.EmitFunc("tee_local %s", arrayPosition)
	c.EmitFunc("i32.load")

	//Increment the stack by 4
	c.EmitFunc("i64.const 4")
	c.EmitFunc("call $__allocstack__")
	// Store the value at the new position
	// c.EmitFunc("")
	c.EmitFunc("i32.store")
	//Decrement the arraysize by 1
	c.EmitFunc("get_local %s", arraySize)
	c.EmitFunc("i32.const 1")
	c.EmitFunc("i32.sub")
	c.EmitFunc("set_local %s", arraySize)
	// Jump back to the start of the loop
	c.EmitFunc("br %s", loopEnd)
	c.EmitFunc("end %s", loopEnd)
	c.EmitFunc("end %s", loop)

	// c.stackPointer += 4
	// c.EmitFunc("i32.const %d", location)
	// c.EmitFunc("call $__")

	//Must be last command to emit
	c.EmitFunc("get_local %s", localName)
	c.EmitFunc(";; Done slicing an array")
	return ast.VisitorMetaData{}, nil
}
