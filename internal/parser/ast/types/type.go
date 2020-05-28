package types

// TypeNode is the basis for the type system. It represents a node that can take another node and returns true or
// false based on if it is satisfied by it.
type TypeNode interface {
	IsSatisfiedBy(t TypeNode) bool
}