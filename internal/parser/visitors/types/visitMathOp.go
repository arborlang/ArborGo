package typevisitor

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
)

func (t *typeVisitor) VisitMathOpNode(n *ast.MathOpNode) (ast.Node, error) {
	left, err := n.LeftSide.Accept(t.v)
	if err != nil {
		return nil, err
	}
	n.LeftSide = left
	right, err := n.RightSide.Accept(t.v)
	if err != nil {
		return nil, err
	}
	n.RightSide = right
	if !n.LeftSide.GetType().IsSatisfiedBy(n.RightSide.GetType()) {
		return nil, fmt.Errorf("can't %s %s and %s (at %s)", n.Operation, n.LeftSide.GetType(), n.RightSide.GetType(), n.Lexeme)
	}
	return n, nil
}
