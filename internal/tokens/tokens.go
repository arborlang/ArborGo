//go:generate stringer -type=Token
package tokens

//Token is the token for the lexem
type Token int

const (
	//EOF is the eof char
	EOF Token = iota
	//RPAREN is the right parenthesis `(`
	RPAREN
	//LPAREN is the right parenthesis `)`
	LPAREN
	//VARNAME is the name
	VARNAME
	//ARROW is the -> symbol
	ARROW
	//COLON is :
	COLON
	//DCOLON is ::
	DCOLON
	//NUMBER is [0-9]+
	NUMBER
	//FLOAT is [0-9]+\.[0-9]+
	FLOAT
	//QUOTE is `
	QUOTE
	//DQOUTE is "
	DQOUTE
	//CHARVAL is '[.*]'
	CHARVAL
	//STRINGVAL is "[.*]"
	STRINGVAL
	//ERROR is an error
	ERROR
	//NEWLINE is the newline character
	NEWLINE
	//ARTHOP is an arithmetic operator
	ARTHOP
	//LOGICAL is a logical operator
	LOGICAL
	//NOT is the not operator
	NOT
	//BOOLEAN is a boolean operator
	BOOLEAN
	//EQUAL is an equal operator
	EQUAL
	//COMMA is the comma
	COMMA
	//PIPE is the pipe operator (|>)
	PIPE
	//SEMI is the semicolon
	SEMI
	//RCURLY is the '{' symbol
	RCURLY
	//LCURLY is the '}' symbol
	LCURLY
	// COMPARISON is a comparrisonOerator
	COMPARISON
	//LSQUARE is the '[' operator
	LSQUARE
	//RSQUARE is the ']' character
	RSQUARE
	//DOT is the '.' character
	DOT
	//AT is the '@' character
	AT

	//LET is the 'let' key word
	LET
	//FUNC is the 'function' keyword
	FUNC
	//STRING is the 'string' keyword
	STRING
	//CHAR is the 'char' keyword
	CHAR
	//DONE is the 'done' keyword
	DONE
	//RETURN is the 'return' key word
	RETURN
	// CONST is the 'const keyword
	CONST
	// IF is the 'if' keyword
	IF
	// ELSE is the 'else' keyword
	ELSE
	// FLOATWORD is the 'float' Keyword
	FLOATWORD
	// NUMBERWORD is the 'number' keyword
	NUMBERWORD
	// IMPORT is the 'import' keyword
	IMPORT
	// AS is the 'as' keyword
	AS
	// SHAPE is the 'shape' keyword
	SHAPE
	//TYPE is the 'type' keyword
	TYPE
	// FROM is the 'from' keyword
	FROM
	// INTERNAL is the 'internal' keyword
	INTERNAL
	//PACKAGE is the package keyword
	PACKAGE
	//NEW is the new keyword
	NEW
	//MATCH is the 'match' keyword
	MATCH
	// WHEN is the 'when' keyword
	WHEN
	// EXTENDS is the 'extends' keyword
	EXTENDS
	// IMPLEMENTS is the 'implements'
	IMPLEMENTS
	//CONTINUE is the 'continue' keyword
	CONTINUE
	//WITH is the 'with' keyword
	WITH
	// SIGNAL is the 'signal' keyword;
	SIGNAL
	// WARN is the 'warn' keyword;
	WARN
	// FATAL is the fatal keyword;
	FATAL
	// TRY is the try keyword
	TRY
	// HANDLE is the handle keyword
	HANDLE
	//SELF is the self keyword
	SELF
	//NOTFOUND if the symbol is not found
	NOTFOUND = -1
)

//EOFChar The actual Character for the EOF symbol
var EOFChar = rune(EOF)
