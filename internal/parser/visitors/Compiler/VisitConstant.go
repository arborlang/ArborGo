package compiler

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
		c.Emit("(i64.const %d)", number)
	case "CHARVAL":
		tp = "char"
		c.Emit("(i32.const %d)", rune(node.Value[0]))
	case "FLOAT":
		tp = "float"
		number, err := strconv.ParseFloat(node.Value, 64)
		if err != nil {
			return ast.VisitorMetaData{}, err
		}
		c.Emit("(f64.const %d)", number)
	default:
		return ast.VisitorMetaData{}, fmt.Errorf("encountered unknown constant")
	}
	return ast.VisitorMetaData{
		Location: "STACK",
		Types:    tp,
	}, nil
}

func visitString(c *Compiler, node *ast.Constant) (ast.VisitorMetaData, error) {

	return ast.VisitorMetaData{}, nil
}
