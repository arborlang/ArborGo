package rules

import (
	"github.com/radding/ArborGo/internal/tokens"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ifStmt = `
if a {
	let c = a + b;	
}`

var ifElseStmt = `
if a {
	let c = a + b;
} else {
	let c = a - b;
}`

var ifElseIfElseStmt = `
if a {
	let c = a + b;
} else if c {
	const d = 5;
} else {
	let c = a - b;
}`

var testMatrix = []struct {
	Name     string
	TestCase string
}{
	{"If statement", ifStmt},
	{"If/Else statement", ifElseStmt},
	{"If/Else If/Else Statement", ifElseIfElseStmt},
}

func TestCanParseIfStatements(t *testing.T) {
	assert := assert.New(t)
	for _, caseStmt := range testMatrix {
		t.Run("if statment: "+caseStmt.Name, func(t *testing.T) {
			p := createParser(caseStmt.TestCase)
			_, err := ifNodeRule(p)
			if !assert.NoError(err) {
				t.Fatal()
			}
		})
	}

	for _, caseStmt := range testMatrix {
		t.Run("Program statment: "+caseStmt.Name, func(t *testing.T) {
			p := createParser(caseStmt.TestCase)
			_, err := ProgramRule(p, tokens.EOF)
			if !assert.NoError(err) {
				t.Fatal()
			}
		})
	}
}
