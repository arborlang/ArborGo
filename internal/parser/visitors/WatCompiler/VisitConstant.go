package wast

import (
	"fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
	"strconv"
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
		Types:    tp,
	}, nil
}

func visitString(c *Compiler, node *ast.Constant) (ast.VisitorMetaData, error) {
	place := c.getUniqueID("string", "begin")
	val := node.Value[1 : len(node.Value)-1]
	data := []byte(val)
	data = append(data, 0x00)
	c.AddData(place, data)
	// c.locals = append(c.locals, locals{name: place, tp: "i64"})
	// val := node.Value[1 : len(node.Value)-1]
	// c.Emit("i64.const %d", len(val))
	// c.Emit("i32.const %d", 1)
	// c.Emit("call $__alloc__")
	// c.Emit("set_local %s", place)
	// for index, char := range node.Value {
	// 	c.Emit("get_local %s", place)
	// 	c.Emit("i64.const %d", index)
	// 	c.Emit("i64.add")
	// 	c.Emit("i32.const %d", char)
	// 	c.Emit("i32.store")
	// }
	return ast.VisitorMetaData{
		Location: strconv.Itoa(c.dataSize),
		Types:    "string",
	}, nil
}
