package umlvisitor

import "github.com/arborlang/ArborGo/internal/parser/ast"

func (u *umlVisitor) VisitMatchNode(node *ast.MatchNode) (ast.Node, error) {
	label := u.getLabel("match")
	u.writeLine("object \"Match Node\" as %s", label)
	uml, _ := node.Match.Accept(u.v)
	u.connectNodeWithLabel(label, uml, "Match")
	for _, whenNode := range node.WhenCases {
		uml, _ := whenNode.Accept(u.v)
		u.connectNodeWithLabel(label, uml, "When")
	}
	return labeledNode(node, label), nil
}

func (u *umlVisitor) VisitWhenNode(node *ast.WhenNode) (ast.Node, error) {
	label := u.getLabel("when")
	u.writeLine("object \"When Node \" as %s", label)
	uml, _ := node.Case.Accept(u.v)
	u.connectNodeWithLabel(label, uml, "When")
	uml, _ = node.Evaluate.Accept(u.v)
	u.connectNodeWithLabel(label, uml, "Then")
	return labeledNode(node, label), nil
}
