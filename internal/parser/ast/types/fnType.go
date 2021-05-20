package types

import (
	"fmt"
	"strings"
)

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
	if f.ReturnVal == nil {
		return true
	}
	if f2.ReturnVal == nil {
		return false
	}
	return f.ReturnVal.IsSatisfiedBy(f2.ReturnVal)
}

func (f *FnType) String() string {
	parms := []string{}
	for _, tp := range f.Parameters {
		parms = append(parms, tp.String())
	}
	params := strings.Join(parms, ", ")
	return fmt.Sprintf("fn (%s) -> %s", params, f.ReturnVal)
}

func mangleShape(shape *ShapeType) string {
	name := &strings.Builder{}
	name.Write([]byte("sh"))
	for field, tp := range shape.Fields {
		name.Write([]byte(field))
		name.Write([]byte(MangleName(tp)))
	}
	return name.String()
}

// MangleName mangles a type name
func MangleName(t TypeNode) string {
	switch rt := t.(type) {
	case *FnType:
		name := &strings.Builder{}
		name.Write([]byte(`_fn`))
		for _, tp := range rt.Parameters {
			name.Write([]byte(MangleName(tp)))
		}
		name.Write([]byte("_"))
		name.Write([]byte(MangleName(rt.ReturnVal)))
		return name.String()
	case *ConstantTypeNode:
		return rt.Name
	case *ArrayType:
		return fmt.Sprintf("_ar%s", MangleName(rt.SubType))
	case *ShapeType:
		return mangleShape(rt)
	case *ExtendedType:
		extMangle := MangleName(rt.Extends)
		addsMangle := MangleName(rt.Shape)
		return fmt.Sprintf("_ex%s%s", extMangle, addsMangle)
	case *TypeGuard:
		name := &strings.Builder{}
		name.Write([]byte(`_gd`))
		for _, i := range rt.Types {
			name.Write([]byte(MangleName(i)))
		}
		return name.String()
	case *TrueType:
		return "_al"
	case *FalseType:
		return "_ne"
	default:
		return ""
	}

}
