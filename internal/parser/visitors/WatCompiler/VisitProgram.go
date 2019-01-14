package wast

import (
	// "bufio"
	// "bytes"
	// "fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
)

// VisitBlock visits a compiler block
func (c *Compiler) VisitBlock(block *ast.Program) (ast.VisitorMetaData, error) {
	metadata := ast.VisitorMetaData{}
	c.level++
	defer func() { c.level-- }()
	// defer c.Flush()
	c.SymbolTable.PushScope()
	defer c.SymbolTable.PopScope()
	// var b bytes.Buffer
	for _, stmt := range block.Nodes {
		value, err := stmt.Accept(c)
		if err != nil {
			return ast.VisitorMetaData{}, err
		}
		if _, ok := stmt.(*ast.ReturnNode); ok {
			metadata.Returns = append(metadata.Returns, value.Returns...)
		}
	}
	return metadata, nil
}

// VisitPipeNode visits the pipe node
func (c *Compiler) VisitPipeNode(node *ast.PipeNode) (ast.VisitorMetaData, error) {
	return ast.VisitorMetaData{}, nil
}

// VisitImportNode visits an import node
func (c *Compiler) VisitImportNode(node *ast.ImportNode) (ast.VisitorMetaData, error) {
	return ast.VisitorMetaData{}, nil
}

// VisitTypeNode visits a type node
func (c *Compiler) VisitTypeNode(node *ast.TypeNode) (ast.VisitorMetaData, error) {
	return ast.VisitorMetaData{}, nil
}

func (c *Compiler) VisitIndexNode(node *ast.IndexNode) (ast.VisitorMetaData, error) {
	return ast.VisitorMetaData{}, nil
}

func (c *Compiler) VisitSliceNode(node *ast.SliceNode) (ast.VisitorMetaData, error) {
	return ast.VisitorMetaData{}, nil
}
