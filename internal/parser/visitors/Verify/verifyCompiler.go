package verify

import (
	"github.com/radding/ArborGo/internal/parser/ast"
)

// VisitBlock visits a compiler block
func (v *Visitor) VisitBlock(block *ast.Program) (ast.VisitorMetaData, error) {
	v.level++
	defer func() { v.level-- }()
	v.symbols.PushScope()
	defer v.symbols.PopScope()
	v.visitor.ShouldCallVisitor = false
	v.visitor.VisitBlock(block)
	return ast.VisitorMetaData{}, nil
}
