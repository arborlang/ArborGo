package wast

import (
	"fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
)

// VisitComparison Visits a comparison node
func (c *Compiler) VisitComparison(node *ast.Comparison) (ast.VisitorMetaData, error) {
	leftSide, err := node.LeftSide.Accept(c)
	if err != nil {
		return ast.VisitorMetaData{}, err
	}
	rightSide, err := node.RightSide.Accept(c)
	if err != nil {
		return ast.VisitorMetaData{}, err
	}

	fmt.Println(leftSide.Types, "!=", rightSide.Types)
	if leftSide.Types != rightSide.Types {
		return ast.VisitorMetaData{}, fmt.Errorf("can't compare object: types don't match")
	}
	tp := ""
	switch leftSide.Types {
	case "char":
		tp = "i32"
	case "float":
		tp = "f64"
	default:
		tp = "i64"
	}
	op := ""
	switch node.Operation {
	case "lt":
		op = "lt_u"
	case "lte":
		op = "le_u"
	case "gt":
		op = "gt_u"
	case "gte":
		op = "ge_u"
	case "eq":
		op = "eq"
	case "neq":
		op = "ne"
	}
	c.EmitFunc("%s.%s", tp, op)
	return ast.VisitorMetaData{
		Types: "bool",
	}, nil
}
