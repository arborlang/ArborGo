package ast

// Node represents a Node in the AST
type Node interface {
	// Accept allows a visitor to travers the tree
	Accept(Visitor) (VisitorMetaData, error)
}

//SymbolData is some data for the symbol in our symbol table
type SymbolData struct {
	Name       string
	Type       *TypeNode
	IsConstant bool
	IsNew      bool
}

// VisitorMetaData is what the visitor will return to represent the conclusion of it work on a node
type VisitorMetaData struct {
	Body       string
	Location   string
	Exportable string
	Types      string
	SymbolData *SymbolData
	Returns    []string
}
