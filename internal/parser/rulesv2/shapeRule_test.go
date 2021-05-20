package rulesv2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstantShape(t *testing.T) {
	assert := assert.New(t)
	obj := `{ object: 1234, value: 1234}`
	p := parseTest(obj)

	shape, err := constantShapeRule(p)
	assert.NoError(err)
	assert.NotNil(shape)
}
