package types

// ConstantTypeNode is a node that is a built in or constant value (eg: "some_string")
type ConstantTypeNode struct {
	Name    string
	IsArray bool
}

// NewConstant returns a new Constant
func NewConstant(name string) *ConstantTypeNode {
	return &ConstantTypeNode{
		Name: name,
	}
}

// IsSatisfiedBy checks the name directly and checks if they match and returns true. If the node is not a constant
// node, or the name doesn't match, return false
func (c *ConstantTypeNode) IsSatisfiedBy(t TypeNode) bool {
	if c2, ok := t.(*ConstantTypeNode); ok {
		return c.Name == c2.Name && c2.IsArray == c.IsArray
	}
	return false
}

func (c *ConstantTypeNode) String() string {
	name := c.Name
	if c.IsArray {
		name = name + "[]"
	}
	return name
}
