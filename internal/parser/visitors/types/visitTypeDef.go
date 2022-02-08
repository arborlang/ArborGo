package typevisitor

import (
	"fmt"
	"runtime/debug"

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
		return nil
	case *types.ArrayType:
		return t.verifyType(realType.SubType, lexeme)
	case *types.ConstantTypeNode:
		otherType, _ := t.scope.LookupSymbol(realType.Name)
		if otherType == nil {
			return fmt.Errorf("no such type %s at %s", realType.Name, lexeme)
		}
		return nil
	case *types.ShapeType:
		for _, value := range realType.Fields {
			err := t.verifyType(value, lexeme)
			if err != nil {
				return fmt.Errorf("error %s", err)
			}
		}
		return nil
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
	case *types.VoidType:
		return nil
	default:
		if tp == nil {
			if t.dumpOnFailure {
				debug.PrintStack()
				fmt.Println(t.scope)
			}
			return fmt.Errorf("Type is nil here: %s. Type: %s", lexeme, realType)
		}
		return nil
	}
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
		typeToExtend, _ := t.scope.LookupSymbol(rt.Name)
		if typeToExtend == nil {
			return nil, fmt.Errorf("%s is not defined", rt.Name)
		}
		return &types.ExtendedType{
			Extends: typeToExtend.Type.Type,
			Shape: &types.ShapeType{
				Fields: map[string]types.TypeNode{},
			},
		}, nil
	default:
		return rt, nil
	}
}

func (t *typeVisitor) VisitTypeNode(tn *ast.TypeNode) (ast.Node, error) {
	tp, _ := t.scope.LookupSymbol(tn.VarName.Name)
	if tp != nil {
		return nil, fmt.Errorf("type %s is being redefined by %s", tp.Type.Name, tn.Lexeme)
	}
	err := t.verifyType(tn.Types, tn.Lexeme)
	if err != nil {
		return nil, err
	}
	derivedType, err := t.deriveNewTypeNode(tn)
	if err != nil {
		return nil, err
	}
	t.scope.AddToScope(tn.VarName.Name, &scope.SymbolData{
		Type: scope.TypeData{
			Type:     derivedType,
			Name:     tn.VarName.Name,
			IsSealed: false,
		},
		Location:   "noop",
		IsConstant: false,
		IsType:     true,
	})
	if _, ok := derivedType.(*types.ExtendedType); ok {
		return &ast.ExtendsNode{
			Extend: tn.VarName,
		}, nil
	}
	return tn, nil
}

func (t *typeVisitor) VisitExtendsNode(n *ast.ExtendsNode) (ast.Node, error) {
	parent, _ := t.scope.LookupSymbol(n.Extend.Name)
	if parent == nil {
		return nil, fmt.Errorf("%s is not defined, can't extend here: %s", n.Extend.Name, n.Extend.Lexeme)
	}
	extended := &types.ExtendedType{}
	extended.Extends = parent.Type.Type
	shapeInGaurd, ok := n.Adds.(*types.TypeGuard)
	if !ok {
		return nil, fmt.Errorf("unexpected type")
	}
	if len(shapeInGaurd.Types) != 1 {
		return nil, fmt.Errorf("extends on type guard %s", n.Name.Lexeme)
	}
	shape, ok := shapeInGaurd.Types[0].(*types.ShapeType)
	if !ok {
		return nil, fmt.Errorf("expected a shape, got a %s", shapeInGaurd.Types[0].String())
	}
	extended.Shape = shape
	t.scope.AddToScope(n.Name.Name, &scope.SymbolData{
		Type: scope.TypeData{
			Type: extended,
		},
		IsType:     true,
		IsConstant: true,
	})
	return n, nil
}
