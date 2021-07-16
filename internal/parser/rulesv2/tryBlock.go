package rulesv2

import (
	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func parseHandleBlock(p *Parser) (*ast.HandleCaseNode, bool, error) {
	// p.PrintNextLexemes(10)
	nxt := p.Next()
	if nxt.Token != tokens.HANDLE {
		return nil, false, nil
	}
	handleNode := &ast.HandleCaseNode{}

	nxt = p.Next()
	if nxt.Token != tokens.RPAREN {
		return nil, false, UnexpectedError(nxt, "(")
	}
	nxt = p.Next()
	if nxt.Token != tokens.VARNAME {
		return nil, false, UnexpectedError(nxt, "variable name")
	}

	varNameVal := nxt.Value
	nxt = p.Next()
	if nxt.Token != tokens.COLON {
		return nil, false, UnexpectedError(nxt, ":")
	}
	tp, err := typeRule(p)
	if err != nil {
		return nil, false, err
	}
	handleNode.Type = tp
	handleNode.VariableName = varNameVal
	if err != nil {
		return nil, false, err
	}

	nxt = p.Next()
	if nxt.Token != tokens.LPAREN {
		return nil, false, UnexpectedError(nxt, ")")
	}
	nxt = p.Next()
	if nxt.Token != tokens.RCURLY {
		return nil, false, UnexpectedError(nxt, "{")
	}
	block, err := ProgramRule(p, tokens.LCURLY)
	if err != nil {
		return nil, false, err
	}
	handleNode.Case = block
	return handleNode, true, nil
	// handles = append(handles, handleNode)

}

func parseTryBlock(p *Parser) (ast.Node, error) {
	nxt := p.Next()
	if nxt.Token != tokens.TRY {
		return nil, UnexpectedError(nxt, "try")
	}
	tryNode := &ast.TryNode{}
	curly := p.Next()
	if curly.Token != tokens.RCURLY {
		return nil, UnexpectedError(curly, "{")
	}

	block, err := ProgramRule(p, tokens.LCURLY)
	if err != nil {
		return nil, err
	}
	tryNode.Tries = block
	handles := []*ast.HandleCaseNode{}

	handleNext := p.Peek()
	if handleNext.Token != tokens.HANDLE {
		return nil, UnexpectedError(handleNext, "handle")
	}
	handleNode, cont, err := parseHandleBlock(p)
	if err != nil {
		return nil, err
	}
	handles = append(handles, handleNode)
	for cont {
		handleNode, cont, err = parseHandleBlock(p)
		if err != nil {
			return nil, err
		}

		handles = append(handles, handleNode)
	}

	return tryNode, nil
}
