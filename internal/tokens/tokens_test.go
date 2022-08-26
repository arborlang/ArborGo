package tokens

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDefinition(t *testing.T) {
	assert := assert.New(t)
	defs := GetDefinitions()
	assert.NotEmpty(defs)
	assert.NotNil(defs["IMPORT"])
	assert.Len(defs, int(MAX)-1) // minus one because NOTFOUND is not in here
	assert.Equal(IMPORT, Token(defs["IMPORT"]))
}
