package tokens

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var reserved = []struct {
	Name     string
	Test     string
	Expected Token
}{
	{"let keyword", "let", LET},
	{"const keyword", "const", CONST},
	{"number keyword", "number", NUMBERWORD},
	{"float", "float", FLOATWORD},
}

func TestCanFindReservedKeyWord(t *testing.T) {
	assert := assert.New(t)
	for _, test := range reserved {
		t.Run(test.Name, func(t *testing.T) {
			str := FindKeyword(test.Test).String()
			fmt.Println(str)
			assert.Equal(test.Expected, FindKeyword(test.Test), "%s: %d != %s:%d", test.Expected, test.Expected, FindKeyword(test.Test), FindKeyword(test.Test))
		})
	}
}
