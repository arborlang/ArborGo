package types

import "fmt"

// ShapeType represents a Shape. Which is just a data record
type ShapeType struct {
	Fields map[string]TypeNode
}

// IsSatisfiedBy Checks to see if this shape is satisfied by n
func (s *ShapeType) IsSatisfiedBy(n TypeNode) bool {
	s2, ok := n.(*ShapeType)
	if !ok {
		maybeExtended, ok := n.(*ExtendedType)
		if !ok {
			return false
		}
		s2 = maybeExtended.Shape
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

func (s *ShapeType) String() string {
	fields := []string{}
	for fieldName, fieldType := range s.Fields {
		fields = append(fields, fmt.Sprintf("%s: %s", fieldName, fieldType))
	}
	return fmt.Sprintf("{%s}", fields)
}
