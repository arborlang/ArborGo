package ast

//Program is the Root node in a file
type Program struct {
	Nodes []Node
}

// Accept Accepts a vistor
func (s *Program) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitBlock(s)
}
