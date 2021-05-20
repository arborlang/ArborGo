package types

// TrueType is a type that is satisfied by everything
type TrueType struct{}

// IsSatisfiedBy tests if this type is satisfied by t. (It is)
func (tr *TrueType) IsSatisfiedBy(t TypeNode) bool {
	return true
}

func (tr *TrueType) String() string {
	return "always"
}

// FalseType is a type that is satisfied by noone
type FalseType struct{}

// IsSatisfiedBy tests if this type is satisfied by t (It is Not)
func (f *FalseType) IsSatisfiedBy(t TypeNode) bool {
	return false
}

func (f *FalseType) String() string {
	return "never"
}
