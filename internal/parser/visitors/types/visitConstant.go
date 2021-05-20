package typevisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

// Visits a Constant Node
func (t *typeVisitor) VisitConstant(constant *ast.Constant) (ast.Node, error) {
	return &annotatedTypeNode{
		node: constant,
		tp:   constant.Type,
	}, nil
}
