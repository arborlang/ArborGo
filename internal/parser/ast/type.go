package ast

// TypeNode represents a type node
type TypeNode struct {
	Types []string
}

// Accept a type visitor
func (t *TypeNode) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitTypeNode(t)
}

// IsValidType Makes sure that a given type stisifies the gaurd
func (t *TypeNode) IsValidType(tp string) bool {
	for _, typ := range t.Types {
		if typ == tp {
			return true
		}
	}
	return false
}
