package wast

import (
	"fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
)

//VisitBoolOp visits a boolean op
func (c *Compiler) VisitBoolOp(node *ast.BoolOp) (ast.VisitorMetaData, error) {
	// c.Emit("(block)")
	leftSide, err := node.LeftSide.Accept(c)
	if err != nil {
		return ast.VisitorMetaData{}, err
	}
	// rightSide, err := node.RightSide.Accept(c)
	// if err != nil {
	// 	return ast.VisitorMetaData{}, err
	// }
	// if leftSide.Types != rightSide.Types {
	// 	return ast.VisitorMetaData{}, fmt.Errorf("can't compare two different types %s %s %s", leftSide.Types, node.Condition, rightSide.Types)
	// }
	tp := ""
	switch leftSide.Types {
	case "char", "bool":
		tp = "i32"
	case "float":
		tp = "f64"
	default:
		tp = "i64"
	}
	c.EmitFunc("%s.const 0", tp)
	c.EmitFunc("%s.ne", tp)
	rightSide, err := node.RightSide.Accept(c)
	if err != nil {
		return ast.VisitorMetaData{}, err
	}
	if rightSide.Types != leftSide.Types {
		return ast.VisitorMetaData{}, fmt.Errorf("can't compare two different types: %s %s %s", leftSide.Types, node.Condition, rightSide.Types)
	}
	c.EmitFunc("%s.const 0", tp)
	c.EmitFunc("%s.ne", tp)

	switch node.Condition {
	case "and":
		c.EmitFunc("i32.and")
	case "or":
		c.EmitFunc("i32.or")
	}
	return ast.VisitorMetaData{
		Types: "bool",
	}, nil
}
