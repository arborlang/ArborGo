package ast

//Program is node that represents the Root of an AST. All valid Arbor ASTs should have this at the
// root
type Program struct {
	//Statements are the Statements in the program
	Statemetents []Node
}

// Walk iterates over every Statement in the the program
func (p *Program) Walk(global, local *CompilerContext, v Visitor) Register {
	v.PreVisit(p, global, nil)
	for _, statement := range p.Statemetents {
		statement.Walk(global, nil, v)
	}
	return v.PostVisit(p, global, nil)
}
