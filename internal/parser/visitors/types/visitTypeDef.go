package typevisitor

import (
	"fmt"

	"github.com/arborlang/ArborGo/internal/lexer"
	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/parser/ast/types"
	"github.com/arborlang/ArborGo/internal/parser/scope"
)

func (t *typeVisitor) verifyType(tp types.TypeNode, lexeme lexer.Lexeme) error {
	switch realType := tp.(type) {
	case *types.TypeGuard:
		for _, no := range realType.Types {
			err := t.verifyType(no, lexeme)
			if err != nil {
				return err
			}
		}
	case *types.ArrayType:
		return t.verifyType(realType.SubType, lexeme)
	case *types.ConstantTypeNode:
		otherType, _ := t.scope.LookupType(realType.Name)
		if otherType == nil {
			return fmt.Errorf("%s is not defined at %s", realType.Name, lexeme)
		}
		return nil
	case *types.ShapeType:
		for key, value := range realType.Fields {
			err := t.verifyType(value, lexeme)
			if err != nil {
				return fmt.Errorf("error with field %s at %s", key, lexeme)
			}
		}
	case *types.FnType:
		for _, params := range realType.Parameters {
			err := t.verifyType(params, lexeme)
			if err != nil {
				return err
			}
		}
		return t.verifyType(realType.ReturnVal, lexeme)
	case *types.ExtendedType:
		err := t.verifyType(realType.Extends, lexeme)
		if err != nil {
			return err
		}
		return t.verifyType(realType.Shape, lexeme)
	default:
		return nil
	}
	return nil
}

func (t *typeVisitor) deriveNewTypeNode(tn *ast.TypeNode) (types.TypeNode, error) {
	tp, ok := tn.Types.(*types.TypeGuard)
	if !ok {
		return nil, fmt.Errorf("expected a type guard, got %s", tn.Types)
	}
	if len(tp.Types) > 1 {
		return tp, nil
	}
	realType := tp.Types[0]
	switch rt := realType.(type) {
	case *types.ConstantTypeNode:
		typeToExtend, _ := t.scope.LookupType(rt.Name)
		if typeToExtend == nil {
			return nil, fmt.Errorf("%s is not defined", rt.Name)
		}
		return &types.ExtendedType{
			Extends: typeToExtend.Type,
		}, nil
	default:
		return rt, nil
	}
}

func (t *typeVisitor) VisitTypeNode(tn *ast.TypeNode) (ast.Node, error) {
	tp, _ := t.scope.LookupType(tn.VarName.Name)
	if tp != nil {
		return nil, fmt.Errorf("type %s is being redefined by %s", tp.Name, tn.Lexeme)
	}
	err := t.verifyType(tn.Types, tn.Lexeme)
	if err != nil {
		return nil, err
	}
	derivedType, err := t.deriveNewTypeNode(tn)
	if err != nil {
		return nil, err
	}
	t.scope.AddType(tn.VarName.Name, &scope.TypeData{
		Type:     derivedType,
		Name:     tn.VarName.Name,
		IsSealed: false,
	})
	if _, ok := derivedType.(*types.ExtendedType); ok {
		return &ast.ExtendsNode{
			Extend: tn.VarName,
		}, nil
	}
	return tn, nil
}
