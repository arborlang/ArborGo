package ast

// ImportNode represents an import statement
type ImportNode struct {
	Source string
	Name   string
}

// Accept a visitor
func (i *ImportNode) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitImportNode(i)
}
