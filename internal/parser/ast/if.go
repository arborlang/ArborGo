package ast

import "github.com/arborlang/ArborGo/internal/lexer"

// IfNode represens an if statemet
type IfNode struct {
	Condition Node
	Body      Node
	ElseIfs   []*IfNode
	Else      Node
	ReturnTo  string
	Lexeme    lexer.Lexeme
}

// Accept implements a node
func (i *IfNode) Accept(v Visitor) (VisitorMetaData, error) {
	return v.VisitIfNode(i)
}
