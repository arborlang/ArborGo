package rulesv2

import (
	"fmt"
	"io"

	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/tokens"
)

// guardParserRule gaurds a parser function from an unexpected keyworld after they run into the rules and pops the semicolon off if it encounters it.
func guardParserRule(p *Parser, rule func(*Parser) (ast.Node, error)) (ast.Node, error) {
	node, err := rule(p)
	if err != nil {
		return nil, err
	}
	if nxt := p.Next(); nxt.Token != tokens.SEMI {
		return nil, fmt.Errorf("unexpected token: got %s exected %q", nxt, ";")
	}
	return node, nil
}

func instantiateRule(p *Parser) (ast.Node, error) {
	nxt := p.Next()
	if nxt.Token != tokens.NEW {
		return nil, fmt.Errorf("expected %q, got %s instead", "new", nxt)
	}
	node := &ast.InstantiateNode{}
	varName := p.Next()
	if varName.Token != tokens.VARNAME {
		return nil, fmt.Errorf("expected a constructor, got %s instead", nxt)
	}
	varNameNode := &ast.VarName{
		Name:   varName.Value,
		Lexeme: varName,
	}
	someNode, err := functionCallRule(varNameNode, p, false)
	if funcCallNode, ok := someNode.(*ast.FunctionCallNode); ok {
		node.FunctionCallNode = funcCallNode
	} else if err != nil {
		err = fmt.Errorf("didn't get a return value from function call that makes sense")
	}
	return node, err
}

func exprRule(p *Parser, allowSquare bool) (ast.Node, error) {
	next := p.Peek()
	switch next.Token {
	case tokens.NUMBER, tokens.STRINGVAL, tokens.CHARVAL, tokens.FLOAT:
		return ConstantsRule(p, true)
	case tokens.VARNAME:
		return varNameRule(p, false)
	case tokens.FUNC:
		return functionDefinitionRule(p)
	case tokens.TYPE:
		return typeDefRule(p)
	case tokens.NEW:
		return instantiateRule(p)
	case tokens.RCURLY:
		return constantShapeRule(p)
	case tokens.FATAL, tokens.WARN, tokens.SIGNAL:
		return parseSignal(p)
	case tokens.CONTINUE:
		return continueRule(p)
	case tokens.SELF:
		return varNameRule(p, false)
	default:
		if allowSquare && next.Token == tokens.RSQUARE {
			return nil, nil
		}
		return nil, fmt.Errorf("encountered unexpected token %s", next)
	}
}

// ExpressionRule parsers an expression. Expressions are any statement that has a return value you need to handle.
func ExpressionRule(p *Parser) (ast.Node, error) {
	return exprRule(p, false)
}

func packageRule(p *Parser) (ast.Node, error) {
	nxt := p.Next()
	if nxt.Token != tokens.PACKAGE {
		return nil, fmt.Errorf("expected \"package\" got: %q (line: %d column: %d)", nxt.Value, nxt.Line, nxt.Column)
	}
	nxt = p.Next()
	if nxt.Token != tokens.VARNAME {
		return nil, fmt.Errorf("expected a variable name, got %s", nxt)
	}
	return &ast.Package{
		Name:   nxt.Value,
		Lexeme: nxt,
	}, nil
}

// StatementRule parses a statement. Statement nodes defined as either a node that returns nothing or an expression followed by a semicolon.
func StatementRule(p *Parser) (ast.Node, error) {
	for next := p.Peek(); next.Token != tokens.SEMI; next = p.Peek() {
		switch next.Token {
		case tokens.LET, tokens.CONST:
			return guardParserRule(p, DeclRule)
		case tokens.RETURN:
			return guardParserRule(p, returnRule)
		case tokens.IF:
			return ifNodeRule(p)
		case tokens.IMPORT:
			return guardParserRule(p, importRule)
		case tokens.INTERNAL:
			return guardParserRule(p, parseExport)
		case tokens.PACKAGE:
			return guardParserRule(p, packageRule)
		case tokens.AT:
			return guardParserRule(p, decoratorRule)
		case tokens.MATCH:
			return matchRule(p)
		case tokens.TRY:
			return parseTryBlock(p)
		// case tokens.SHAPE:
		// 	return
		default:
			return guardParserRule(p, ExpressionRule)
		}

	}
	tok := p.Peek()
	return nil, fmt.Errorf("encountered unknown token: %s", tok)
}

func sliceContains(tok tokens.Token, slice []tokens.Token) bool {
	for _, allowedTok := range slice {
		if allowedTok == tok {
			return true
		}
	}
	return false
}

func typeDefRule(p *Parser) (ast.Node, error) {
	nxt := p.Next()
	if nxt.Token != tokens.TYPE {
		return nil, fmt.Errorf("unexpected token, expected 'type' got %s", nxt)
	}
	nxt = p.Next()
	if nxt.Token != tokens.VARNAME {
		return nil, fmt.Errorf("unexpected token, expected a variable name, got %s", nxt)
	}
	// if
	peek := p.Peek()
	// if peek.Token != tokens.VARNAME && peek.Token != tokens.RCURLY && peek.Token != tokens.EXTENDS && peek.Token !=  {
	doesExtend := false
	var extends *ast.VarName = nil
	doesImplement := false
	varName := &ast.VarName{
		Name:   nxt.Value,
		Lexeme: nxt,
	}
	if peek.Token == tokens.EXTENDS {
		p.Next()
		doesExtend = true
		nxt := p.Next()
		if nxt.Token != tokens.VARNAME {
			return nil, UnexpectedError(nxt, "a type name")
		}
		extends = &ast.VarName{

			Name:   nxt.Value,
			Lexeme: nxt,
		}
	}
	generics, err := parseGenericDef(p)
	if err != nil {
		return nil, err
	}
	peek = p.Peek()
	if !sliceContains(peek.Token, []tokens.Token{tokens.VARNAME, tokens.RCURLY, tokens.IMPLEMENTS}) {
		return nil, UnexpectedError(peek, "a type", "{", "extends", "implements")
	}
	varNames := []*ast.VarName{}
	if peek.Token == tokens.IMPLEMENTS {
		p.Next()
		doesImplement = true
		for {
			nxt := p.Next()
			if nxt.Token != tokens.VARNAME {
				return nil, UnexpectedError(nxt, "a type name")
			}
			varNames = append(varNames, &ast.VarName{
				Name:   nxt.Value,
				Lexeme: nxt,
			})
			peek := p.Peek()
			if peek.Token == tokens.COMMA {
				p.Next()
			} else {
				break
			}
		}
		doesImplement = true
	}
	tp, err := typeRule(p)
	if err != nil {
		return nil, err
	}
	typeNode := &ast.TypeNode{
		Types:        tp,
		VarName:      varName,
		Extends:      doesExtend,
		Lexeme:       nxt,
		GenericTypes: generics,
	}
	if doesImplement {
		return &ast.ImplementsNode{
			Implements: varNames,
			Type:       typeNode,
			Lexeme:     nxt,
		}, nil
	}
	if extends != nil {
		return &ast.ExtendsNode{
			Name:   varName,
			Extend: extends,
			Adds:   tp,
		}, nil
	}
	return typeNode, nil
}

// ProgramRule is basicaly the entry point for all Arbor files. Returns the root node
func ProgramRule(p *Parser, until tokens.Token) (ast.Node, error) {
	program := &ast.Program{}
	for next := p.Next(); next.Token != until; next = p.Next() {
		if next.Token == tokens.NEWLINE {
			continue
		}
		p.Backup()
		node, err := StatementRule(p)
		if err != nil {
			return nil, err
		}
		// program.AddChild(node)
		if importNode, ok := node.(*ast.ImportNode); ok {
			program.Imports = append(program.Imports, importNode)
			continue
		}
		program.Nodes = append(program.Nodes, node)
	}
	return program, nil
}

// Parse creates the full AST
func Parse(reader io.Reader) (ast.Node, error) {
	tokStream := lexer.Lex(reader)
	p := New(tokStream)
	return ProgramRule(p, tokens.EOF)
}
