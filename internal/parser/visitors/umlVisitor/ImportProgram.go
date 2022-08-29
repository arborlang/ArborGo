package umlvisitor

import (
	"strings"

	"github.com/arborlang/ArborGo/internal/parser/ast"
)

func (u *umlVisitor) VisitImportNode(node *ast.ImportNode) (ast.Node, error) {
	label := u.getLabel("import")
	u.writeLine("object \"import %s\" as %s {", strings.Trim(node.Source, "\""), label)
	u.writeLine("\tImports = %s", node.ExportName)
	u.writeLine("\tAs = %s", node.ImportAs)
	u.writeLine("\tFrom = %s", node.Source)
	u.writeLine("}")
	return labeledNode(node, label), nil
}
