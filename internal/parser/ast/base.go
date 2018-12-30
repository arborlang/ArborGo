package ast

// Node represents a Node in the AST
type Node interface {
	// Accept allows a visitor to travers the tree
	Accept(Visitor) (VisitorMetaData, error)
}

// VisitorMetaData is what the visitor will return to represent the conclusion of it work on a node
type VisitorMetaData struct {
	Body     string
	Location string
}
