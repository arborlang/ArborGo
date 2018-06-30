package ast

import "errors"

//CompilerContext refers to the currrent compilation context in the
type CompilerContext struct {
	context map[string]interface{}
}

//Add adds to the context
func (c *CompilerContext) Add(name string, val interface{}) {
	c.context[name] = val
}

//Get retrieves from the context
func (c *CompilerContext) Get(name string) (interface{}, error) {
	val, ok := c.context[name]
	if !ok {
		return nil, errors.New("not in the context")
	}
	return val, nil
}
