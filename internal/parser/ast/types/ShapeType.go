package types

// ShapeType represents a Shape. Which is just a data record
type ShapeType struct {
	Fields map[string]TypeNode
}

// IsSatisfiedBy Checks to see if this shape is satisfied by n
func (s *ShapeType) IsSatisfiedBy(n TypeNode) bool {
	s2, ok := n.(*ShapeType)
	if !ok {
		return false
	}
	for fieldName, tp := range s.Fields {
		fld, ok := s2.Fields[fieldName]
		if !ok {
			return false
		}
		if !tp.IsSatisfiedBy(fld) {
			return false
		}
	}
	return true
}