package rules

import (
	"testing"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/stretchr/testify/assert"
)

func TestReturnWorks(t *testing.T) {
	retStmt := "return a + b + c + d()"
	assert := assert.New(t)
	p := createParser(retStmt)
	ret, err := returnRule(p)
	if !assert.NoError(err) {
		t.Fatal()
	}
	assert.NotNil(ret)
	retNode := ret.(*ast.ReturnNode)
	if !assert.NotNil(retNode, "Could not convert to ReturnNode") {
		t.Fatal()
	}
}
