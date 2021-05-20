package types

import "fmt"

// ArrayType is an array type
type ArrayType struct {
	SubType TypeNode
}

func (a *ArrayType) IsSatisfiedBy(t TypeNode) bool {
	tArr, ok := t.(*ArrayType)
	if !ok {
		return false
	}
	return a.SubType.IsSatisfiedBy(tArr.SubType)
}

func (a *ArrayType) String() string {
	return fmt.Sprintf("(%s)[]", a.SubType)
}
