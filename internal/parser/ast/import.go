package ast

import (
	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
)

// ImportNode represents an import statement
type ImportNode struct {
	ExportName string `IMPORT @VARNAME`
	ImportAs   string `('as' @VARNAME)?`
	Source     string `'from' @STRINGVAL`
	Lexeme     lexer.Lexeme
	NextImport *ImportNode
}

// Accept a visitor
func (i *ImportNode) Accept(v Visitor) (Node, error) {
	return v.VisitImportNode(i)
}

func (i *ImportNode) GetType() types.TypeNode {
	return &types.FalseType{}
}
