package rulesv2

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
	"github.com/arborlang/ArborGo/internal/tokens"
)

func parseGenericDef(p *Parser) ([]*ast.VarName, error) {
	genericTypeNames := []*ast.VarName{}
	peek := p.Peek()
	if p.Peek().Token == tokens.COMPARISON && p.Peek().Value == "<" {
		p.Next()
		peek = p.Peek()
		for peek.Token != tokens.COMPARISON && peek.Value != ">" {
			typeName := p.Next()
			if typeName.Token != tokens.VARNAME {
				return nil, fmt.Errorf("Expected a Type Name, got %s instead", typeName)
			}
			genericTypeNames = append(genericTypeNames, &ast.VarName{
				Name:   typeName.Value,
				Lexeme: typeName,
			})
			peek = p.Next()
			if !(peek.Token == tokens.COMMA || (peek.Token == tokens.COMPARISON && peek.Value == ">")) {
				return genericTypeNames, fmt.Errorf("expected %q or %q, got %s instead", ",", ">", peek)
			}
		}
	}
	return genericTypeNames, nil
}

// func paramTypeParam(p *)

func paramsParser(p *Parser) ([]*ast.VarName, error) {
	paren := p.Next()
	if paren.Token != tokens.RPAREN {
		return nil, fmt.Errorf("expecting %q, instead got %s", "(", paren)
	}
	nodes := []*ast.VarName{}
	for nxt := p.Peek(); nxt.Token != tokens.LPAREN; nxt = p.Peek() {
		tok := p.Peek()
		varNameNode, err := varNameRule(p, true)
		varName := varNameNode.(*ast.VarName)
		if err != nil {
			return nil, err
		}
		if varName.Type == nil {
			return nil, fmt.Errorf("no type defined for %q", tok)
		}
		nodes = append(nodes, varName)
		nxt = p.Peek()
		if nxt.Token == tokens.COMMA {
			p.Next()
		} else if nxt.Token != tokens.LPAREN {
			return nil, fmt.Errorf("expected %q or %q, instead got %s", ")", ",", nxt)
		}
	}
	p.Next() // This should be the ")"
	return nodes, nil
}

func functionDefinitionRule(p *Parser) (ast.Node, error) {
	// funcDef := ast.FunctionDefinitionNode{}
	var asignNode *ast.AssignmentNode = nil
	var methodDef *ast.MethodDefinition = nil
	next := p.Next()
	if next.Token != tokens.FUNC {
		return nil, fmt.Errorf("expecting %q, got %s", "fn", next)
	}
	varName := p.Peek()
	if varName.Token == tokens.VARNAME {
		p.Next()
		retVal := &ast.VarName{
			Name:   varName.Value,
			Lexeme: varName,
		}
		asignNode = &ast.AssignmentNode{
			Lexeme: varName,
		}
		asignNode.AssignTo = &ast.DeclNode{
			Lexeme:     varName,
			Varname:    retVal,
			IsConstant: true,
		}
	}

	funcNode := &ast.FunctionDefinitionNode{
		Lexeme: varName,
	}
	peek := p.Peek()
	if peek.Token == tokens.DCOLON {
		p.Next()
		methodDef = &ast.MethodDefinition{
			Lexeme: peek,
		}
		methodDef.TypeName = asignNode.AssignTo.(*ast.DeclNode).Varname
		methodName := p.Next()
		if methodName.Token != tokens.VARNAME {
			return nil, fmt.Errorf("expected var name, got %s instead", methodName)
		}
		methodDef.MethodName = &ast.VarName{
			Name:   methodName.Value,
			Lexeme: methodName,
		}
		asignNode = nil
	}
	peek = p.Peek()
	genericTypeNames, err := parseGenericDef(p)
	funcNode.GenericTypeNames = genericTypeNames
	if err != nil {
		return funcNode, err
	}
	//Generic parsing
	// if peek.Token == tokens.COMPARISON && peek.Value == "<" {
	// 	p.Next()
	// 	peek = p.Peek()
	// 	for peek.Token != tokens.COMPARISON && peek.Value != ">" {
	// 		typeName := p.Next()
	// 		if typeName.Token != tokens.VARNAME {
	// 			return nil, fmt.Errorf("Expected a Type Name, got %s instead", typeName)
	// 		}
	// 		funcNode.GenericTypeNames = append(funcNode.GenericTypeNames, &ast.VarName{
	// 			Name:   typeName.Value,
	// 			Lexeme: typeName,
	// 		})
	// 		peek = p.Next()
	// 		if !(peek.Token == tokens.COMMA || (peek.Token == tokens.COMPARISON && peek.Value == ">")) {
	// 			return nil, fmt.Errorf("expected %q or %q, got %s instead", ",", ">", peek)
	// 		}
	// 	}

	// }
	params, err := paramsParser(p)
	if err != nil {
		return nil, err
	}
	tps := []types.TypeNode{}
	for _, varName := range params {
		tps = append(tps, varName.Type)
	}
	funcNode.Arguments = params
	next = p.Next()
	if next.Token != tokens.ARROW && next.Token != tokens.RCURLY {
		return nil, fmt.Errorf("expected %q or %q, instead got %s", "->", "{", next)
	}
	if next.Token == tokens.ARROW {
		retType, err := typeRule(p)
		if err != nil {
			return nil, err
		}
		funcNode.Returns = retType
		next = p.Next()
	}

	if next.Token != tokens.RCURLY {
		return nil, fmt.Errorf("expected %q, instead got %s", "{", next)
	}
	funcNode.Body, err = ProgramRule(p, tokens.LCURLY)
	if err != nil {
		return nil, err
	}
	if asignNode != nil {
		asignNode.Value = funcNode
		return asignNode, nil
	}
	if methodDef != nil {
		methodDef.FuncDef = funcNode
		return methodDef, nil
	}
	return funcNode, nil
}
