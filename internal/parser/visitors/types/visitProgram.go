package typevisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

func (t *typeVisitor) VisitProgam(prog *ast.Program) (ast.Node, error) {
	t.scope.PushNewScope()
	defer t.scope.PopScope()
	blocks := []ast.Node{}
	for _, block := range prog.Nodes {
		newBlock, err := block.Accept(t.v)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, newBlock)
	}
	prog.Nodes = blocks
	return prog, nil
}
