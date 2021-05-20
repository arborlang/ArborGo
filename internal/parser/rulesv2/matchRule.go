package rulesv2

import (
	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func matchRule(p *Parser) (ast.Node, error) {
	match := p.Next()
	if match.Token != tokens.MATCH {
		return nil, UnexpectedError(match, "match")
	}
	matchNode := &ast.MatchNode{}
	expressionNode, err := exprRule(p, true)
	if err != nil {
		return nil, err
	}
	curlyMaybe := p.Next()
	if curlyMaybe.Token != tokens.RCURLY {
		return nil, UnexpectedError(curlyMaybe, "{")
	}
	whenNodes := []*ast.WhenNode{}
	for {
		next := p.Next()
		if next.Token != tokens.PIPE && next.Token != tokens.LCURLY {
			return nil, UnexpectedError(next, "}", "|>")
		}
		if next.Token == tokens.LCURLY {
			break
		}
		whenCondition, err := exprRule(p, true)
		if err != nil {
			return nil, err
		}
		curly := p.Next()
		if curly.Token != tokens.RCURLY {
			return nil, UnexpectedError(curly, "{")
		}
		eval, err := ProgramRule(p, tokens.LCURLY)
		if err != nil {
			return nil, err
		}
		whenNode := &ast.WhenNode{
			Evaluate: eval,
			Case:     whenCondition,
		}
		whenNodes = append(whenNodes, whenNode)
	}
	matchNode.WhenCases = whenNodes
	matchNode.Match = expressionNode
	return matchNode, nil
}
