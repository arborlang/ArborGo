package types

type VoidType struct{}

func (v *VoidType) IsSatisfiedBy(t TypeNode) bool {
	_, ok := t.(*VoidType)
	return ok
}

func (v *VoidType) String() string {
	return "void"
}
