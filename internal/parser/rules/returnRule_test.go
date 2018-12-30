package rules

import (
	"github.com/radding/ArborGo/internal/parser/ast"
	"github.com/stretchr/testify/assert"
	"testing"
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
