package wast

import (
	// "encoding/binary"
	"fmt"
	"strconv"

	"github.com/arborlang/ArborGo/internal/parser/ast"
)

// VisitConstant visits the constant object
func (c *Compiler) VisitConstant(node *ast.Constant) (ast.VisitorMetaData, error) {
	// var value []byte
	tp := ""
	switch node.Type {
	case "STRINGVAL":
		return visitString(c, node)
	case "NUMBER":
		tp = "number"
		number, err := strconv.Atoi(node.Value)
		if err != nil {
			return ast.VisitorMetaData{}, err
		}
		c.EmitFunc("i64.const %d", number)
	case "CHARVAL":
		tp = "char"
		c.EmitFunc("i32.const %d", rune(node.Raw[1]))
	case "FLOAT":
		tp = "float"
		number, err := strconv.ParseFloat(node.Value, 64)
		if err != nil {
			return ast.VisitorMetaData{}, err
		}
		c.EmitFunc("f64.const %f", number)
	default:
		return ast.VisitorMetaData{}, fmt.Errorf("encountered unknown constant")
	}
	return ast.VisitorMetaData{
		Location: "STACK",
		Types:    &ast.TypeNode{Types: []string{tp}},
	}, nil
}

// Converts a number to a little endian
func littleEndian(number int32) uint32 {
	// LE := binary.LittleEndian
	// BE := binary.BigEndian
	// b := make([]byte, 4)
	// LE.PutUint32(b, uint32(number))
	// fmt.Printf("Value: %X", BE.Uint32(b))
	return uint32(number)
	// return BE.Uint32(b)
}

func visitString(c *Compiler, node *ast.Constant) (ast.VisitorMetaData, error) {
	c.EmitFunc(";; Compiling a string baby!")
	place := c.getUniqueID("string", "begin")
	val := node.Value[1 : len(node.Value)-1]
	c.AddLocal(place, "i32")
	c.EmitFunc("call $__stacktop__")
	c.EmitFunc("set_local %s", place)
	c.EmitFunc("call $__stacktop__")
	c.EmitFunc("i32.const %d", littleEndian(int32(len(val))))
	c.EmitFunc("i32.store")
	for _, char := range val {
		c.EmitFunc("i64.const 4")
		c.EmitFunc("call $__allocstack__")
		c.EmitFunc("i32.const %d", littleEndian(int32(char)))

		c.EmitFunc("i32.store")
	}
	c.EmitFunc(";; Done with that string")
	return ast.VisitorMetaData{
		Location: strconv.Itoa(c.dataSize),
		Types:    &ast.TypeNode{Types: []string{"string"}},
	}, nil
}
