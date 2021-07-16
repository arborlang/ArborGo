package rulesv2

import (
	// "fmt"

	// "github.com/arborlang/ArborGo/internal/parser/ast"
	"fmt"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
	"github.com/arborlang/ArborGo/internal/tokens"
)

// fn (name: Type) -> Type
func parseFnType(p *Parser) (types.TypeNode, error) {
	nxt := p.Next()
	if nxt.Token != tokens.RPAREN {
		return nil, fmt.Errorf("unexpected token %s, expected \"(\"", nxt.Value)
	}
	// varName := p.Next()
	params := make([]types.TypeNode, 0)
	peek := p.Peek()
	if peek.Token == tokens.LPAREN {
		p.Next()
	} else {
		for nxt.Token != tokens.LPAREN {
			varName := p.Next()
			if varName.Token != tokens.VARNAME {
				return nil, fmt.Errorf("unexpected token %s, expected \"VARNAME\"", varName)
			}
			colon := p.Next()
			if colon.Token != tokens.COLON {
				return nil, fmt.Errorf("unexpected token %s, expected \":\"", colon)
			}
			tp, err := typeRule(p)
			if err != nil {
				return nil, err
			}
			params = append(params, tp)
			nxt = p.Next()
			if !(nxt.Token == tokens.COMMA || nxt.Token == tokens.LPAREN) {
				return nil, fmt.Errorf("unexpected token %s, expected %q or %q", nxt, ",", ")")
			}
		}

	}
	arrow := p.Peek()
	var retType types.TypeNode = nil
	var err error
	if arrow.Token == tokens.ARROW {
		p.Next()
		retType, err = typeRule(p)
		if err != nil {
			return nil, err
		}
	}
	return &types.FnType{
		Parameters: params,
		ReturnVal:  retType,
	}, nil
}

func parseShapeType(p *Parser) (types.TypeNode, error) {
	fields := make(map[string]types.TypeNode)
	for nxt := p.Next(); nxt.Token != tokens.LCURLY; nxt = p.Next() {
		if nxt.Token != tokens.VARNAME {
			return nil, fmt.Errorf("unexpected token %s. Expected a %q", nxt, "Variable Name")
		}
		colon := p.Next()
		if colon.Token != tokens.COLON {
			return nil, fmt.Errorf("unexpected token %s, expected a %q", colon, ":")
		}
		var tp types.TypeNode
		_, err := guardParserRule(p, func(p *Parser) (ast.Node, error) {
			var err error = nil
			tp, err = typeRule(p)
			return nil, err
		})
		if err != nil {
			return nil, err
		}
		if _, ok := fields[nxt.Value]; ok {
			return nil, fmt.Errorf("redefining field %s on shape", nxt)
		}

		fields[nxt.Value] = tp
	}
	shapeTP := &types.ShapeType{
		Fields: fields,
	}
	return shapeTP, nil
}

func subTypeRule(p *Parser) (types.TypeNode, error) {
	nxt := p.Next()
	var tp types.TypeNode = nil
	var err error = nil
	if nxt.Token == tokens.VARNAME {
		tp = &types.ConstantTypeNode{Name: nxt.Value}
		// peek := p.Peek()
		// if
	}
	if nxt.Token == tokens.STRINGVAL {
		tp = &types.ConstantTypeNode{Name: nxt.Value}
	}
	if nxt.Token == tokens.CHARVAL {
		tp = &types.ConstantTypeNode{Name: nxt.Value}
	}
	if nxt.Token == tokens.NUMBER {
		tp = &types.ConstantTypeNode{Name: nxt.Value}
	}
	if nxt.Token == tokens.FLOAT {
		tp = &types.ConstantTypeNode{Name: nxt.Value}
	}
	if nxt.Token == tokens.FUNC {
		tp, err = parseFnType(p)
	}
	if nxt.Token == tokens.RCURLY {
		tp, err = parseShapeType(p)
	}
	if err == nil {
		peek := p.Peek()
		if peek.Token == tokens.LSQUARE {
			p.Next()
			tp = &types.ArrayType{
				SubType: tp,
			}
			nxt := p.Next()
			if nxt.Token != tokens.RSQUARE {
				return nil, UnexpectedError(nxt, "]")
			}
		}
	}
	return tp, err
}

/**
*  A type rule is something like Int | String | SomeType;
 */
func typeRule(p *Parser) (types.TypeNode, error) {
	tpGuard := &types.TypeGuard{Types: make([]types.TypeNode, 0)}
	for {
		tp, err := subTypeRule(p)
		if err != nil {
			return nil, err
		}
		tpGuard.Types = append(tpGuard.Types, tp)
		if p.Peek().Token != tokens.LOGICAL || p.Peek().Value != "|" {
			break
		}
		p.Next()
	}
	return tpGuard, nil
}
