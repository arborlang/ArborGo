package rules

import (
	"github.com/arborlang/ArborGo/internal/parser/ast"
)

func assignmentOperator(varname *ast.VarName, p *Parser) (ast.Node, error) {
	assignment := &ast.AssignmentNode{}
	assignment.AssignTo = varname
	assignment.Lexeme = p.Next()
	right, err := ExpressionRule(p)
	if err != nil {
		return nil, err
	}
	assignment.Value = right
	return assignment, nil
}
