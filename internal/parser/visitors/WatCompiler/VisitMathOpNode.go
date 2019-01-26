package wast

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
)

// VisitMathOpNode Visits a math op node
func (c *Compiler) VisitMathOpNode(node *ast.MathOpNode) (ast.VisitorMetaData, error) {
	lsMetadata, err := node.LeftSide.Accept(c)
	if err != nil {
		return lsMetadata, err
	}
	rsMetadata, err := node.RightSide.Accept(c)
	if err != nil {
		return rsMetadata, err
	}
	if lsMetadata.Types != rsMetadata.Types {
		return lsMetadata, fmt.Errorf("can't %s two different types", node.Operation)
	}
	tp := c.getType(lsMetadata.Types)
	c.EmitFunc("%s.%s", tp, node.Operation)
	return lsMetadata, nil

}
