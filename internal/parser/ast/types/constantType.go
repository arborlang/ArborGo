package types

// ConstantTypeNode is a node that is a built in or constant value (eg: "some_string")
type ConstantTypeNode struct {
	Name string
}

// IsSatisfiedBy checks the name directly and checks if they match and returns true. If the node is not a constant 
// node, or the name doesn't match, return false
func (c *ConstantTypeNode) IsSatisfiedBy(t TypeNode) bool {
	if c2, ok := t.(*ConstantTypeNode); ok {
		return c.Name == c2.Name
	}
	return false
}