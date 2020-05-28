package types

// FnType represents a function type
type FnType struct {
	Parameters []TypeNode
	ReturnVal  TypeNode
}

// IsSatisfiedBy checks to see if this value is satisfied by another type.
//  the only type that can satisfy this is another function.
func (f *FnType) IsSatisfiedBy(t TypeNode) bool {
	f2, ok := t.(*FnType)
	if !ok {
		return false
	}
	if len(f.Parameters) != len(f2.Parameters) {
		return false
	}
	for i, param := range f.Parameters {
		if !param.IsSatisfiedBy(f2.Parameters[i]) {
			return false
		}
	}
	return f.ReturnVal.IsSatisfiedBy(f2.ReturnVal)
}
