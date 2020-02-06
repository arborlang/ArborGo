package verify

import (
	"github.com/arborlang/ArborGo/internal/parser/ast"
)

// VisitBlock visits a compiler block
func (v *Visitor) VisitBlock(block *ast.Program) (md ast.VisitorMetaData, e error) {
	v.level++
	v.symbols.PushNewScope()
	defer func() {
		v.level--
		e = v.symbols.PopScope()
	}()
	// v.symbols.PushScope()
	// defer v.symbols.PopScope()
	v.visitor.ShouldCallVisitor = false
	v.visitor.VisitBlock(block)
	md = ast.VisitorMetaData{}
	e = nil
	return
}
