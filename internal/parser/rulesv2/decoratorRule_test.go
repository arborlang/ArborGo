package rulesv2

import (
	"testing"

	"github.com/arborlang/ArborGo/internal/tokens"
	"github.com/stretchr/testify/assert"
)

func TestDecorators(t *testing.T) {
	assert := assert.New(t)
	simple := `@Simple
	fn Decorated() {
		
	};`

	p := parseTest(simple)
	_, err := ProgramRule(p, tokens.EOF)
	assert.NoError(err)
}
