package 

//Written by the generator, do not over write


type GenericVisitor interface {
	VisitAnyNode(n Node) (Node, error)
}

type Visitor interface {

	GenericVisitor
}
