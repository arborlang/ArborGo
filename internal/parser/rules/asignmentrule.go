package rules

import (
	"github.com/radding/ArborGo/internal/parser/ast"
)

func assignmentOperator(varname *ast.VarName, p *Parser) (ast.Node, error) {
	assignment := &ast.AssignmentNode{}
	assignment.AssignTo = varname
	p.Next() // We are here, there for we know that we saw an equal and can safely skip it
	right, err := ExpressionRule(p)
	if err != nil {
		return nil, err
	}
	assignment.Value = right
	return assignment, nil
}
