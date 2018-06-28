package ast

import (
	"context"
	"io"
)

//Register is just a string that represents a register, which is just some place holder that
// 	represents where some value will be stored. At this point it should just be the representation
// 	of where the value will end up. It doesn't refer to a physical or a VM's register.
type Register string

// Node is a node in the AST.
type Node interface {
	// Compile writes out the result of the node to the the io.Writer and returns the "Register"
	// 	that the operation wrote to. The context is the context for the currently compiled lexical stream
	//	Compile also takes its children nodes to compile (if applicable)
	Compile(context.Context, io.Writer, ...Node) Register
}

// Compiler is a function that implements the Node interface.
type Compiler func(context.Context, io.Writer, ...Node) Register

//Compile satisfies the Node interface
func (n Compiler) Compile(c context.Context, w io.Writer, ns ...Node) Register {
	return n(c, w, ns...)
}
