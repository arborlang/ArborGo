package types

import "fmt"

// ExtendedType extends an existing Type
type ExtendedType struct {
	Extends TypeNode
	Shape   *ShapeType
}

// IsSatisfiedBy checks to see if this ExtendedType is good
func (e *ExtendedType) IsSatisfiedBy(t TypeNode) bool {
	return e.Extends.IsSatisfiedBy(t) && e.Shape.IsSatisfiedBy(t)
}

func (e *ExtendedType) String() string {
	return fmt.Sprintf("extends %s, adds %s", e.Extends, e.Shape)
}
