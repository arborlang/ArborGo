package rulesv2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanParseMatch(t *testing.T) {
	assert := assert.New(t)
	match := `
	match val {
		|> {foobar: 1, bazBar: 2 } {
			return;
		}
		|> val.foobar == 10 || val.bazBar == 2 {
			return 0;
		}
	}
	`
	p := parseTest(match)

	_, err := matchRule(p)
	assert.NoError(err)
}
