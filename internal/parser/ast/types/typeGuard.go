package types

import "strings"

// TypeGuard is a type guard that allows you to pick one type. Looks like: Number | String
type TypeGuard struct {
	Types []TypeNode
}

// IsSatisfiedBy Takes a node and checks each of the types and find the first node to be satisfied
func (t *TypeGuard) IsSatisfiedBy(n TypeNode) bool {
	g, ok := n.(*TypeGuard)
	if ok {
		for _, i := range t.Types {
			for _, j := range g.Types {
				if i.IsSatisfiedBy(j) {
					return true
				}
			}
		}
	}
	for _, i := range t.Types {
		if i.IsSatisfiedBy(n) {
			return true
		}
	}
	return false
}

func (t *TypeGuard) String() string {
	strs := []string{}
	for _, tp := range t.Types {
		strs = append(strs, tp.String())
	}
	return strings.Join(strs, " | ")
}
