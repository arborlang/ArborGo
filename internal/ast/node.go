package ast

//Register is just a string that represents a register, which is just some place holder that
// 	represents where some value will be stored. At this point it should just be the representation
// 	of where the value will end up. It doesn't refer to a physical or a VM's register.
type Register string

//Visitor is what walks down the AST and visits each node. The visitor should be responsible for
// 	compiling the the AST into its own compiler. This is abstracted away from the Compile method on
// 	the Node interface for greater flexibilty. For example, the Arbor compiler would want to walk
// 	down the AST and generate some IR, but an IDE would want to provide some autocorrection based on
// 	the current AST. If the idea of converting the Node to some text is tied to the Node interface,
// 	then it becomes much less flexible. This also allows us to keep the AST the same but provide
// 	different visitors to output different IRs (like if we want to go from LLVM to our own IR).
// 	all visitors visit the nodes in post order traversal and pre order traversal, that is we visit
//  all the children nodes, then we visit this node. This is why we have global and local context.
type Visitor interface {

	//PostVisit visits each node and does what it needs to do, it returns a register if it needs to
	// so that other nodes are aware of the result of the calculation (or rather where we put that).
	// Visit takes a the node we are currently visiting and the global compilation context and the
	// local compilation context. The difference between global and local context is that the global
	// contexts are things that effects the entire compilation, things like scope tables should be
	// here, while the local context is information about the current Node and information about its
	// children nodes, like what registers they put their values in. The local context only goes down
	// one level (from parent to immediate child). If something needs to be communicated either
	// further down the chain or up the chain, use the global context
	PostVisit(node Node, global, local *CompilerContext) Register

	//PreVisit basically does the same thing as PostVisit, but it is called before children are visited
	PreVisit(node Node, global, local *CompilerContext) Register
}

// Node is a node in the AST. This holds all of the contextual information about that node, like the
// 	value of the node.
type Node interface {
	//Walk Allows you to walk through the AST. Walk is responsible for instantiating the local
	// context and passing it to the children, calling the Visitor, and walking through its
	// children.
	Walk(global, local *CompilerContext, visitor Visitor) Register
}
