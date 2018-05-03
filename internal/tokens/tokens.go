package tokens

//Token is the token for the lexem
type Token int

const (
	//EOF is the eof char
	EOF Token = iota
	//RPAREN is the right parenthesis
	RPAREN
	//LPAREN is the right parenthesis
	LPAREN
	//VARNAME is the name
	VARNAME
	//ARROW is the -> symbol
	ARROW
	//SEMICOLON is ;
	SEMICOLON
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
	//NOTFOUND if the symbol is not found
	NOTFOUND = -1
)

func (tok Token) String() string {
	switch tok {
	case EOF:
		return "EOF"
	case RPAREN:
		return "RPAREN"
	case LPAREN:
		return "LPAREN"
	case VARNAME:
		return "VARNAME"
	case ARROW:
		return "ARROW"
	case SEMICOLON:
		return "SEMICOLON"
	case NUMBER:
		return "NUMBER"
	case FLOAT:
		return "FLOAT"
	case QUOTE:
		return "QUOTE"
	case DQOUTE:
		return "DQOUTE"
	case CHARVAL:
		return "CHARVAL"
	case STRINGVAL:
		return "STRINGVAL"
	case ERROR:
		return "ERROR"
	case NEWLINE:
		return "NEWLINE"
	case LET:
		return "LET"
	case FUNC:
		return "FUNC"
	case STRING:
		return "STRING"
	case CHAR:
		return "CHAR"
	case DONE:
		return "DONE"
	case RETURN:
		return "RETURN"
	case ARTHOP:
		return "ARTHOP"
	case LOGICAL:
		return "LOGICAL"
	case NOT:
		return "NOT"
	case BOOLEAN:
		return "BOOLEAN"
	case EQUAL:
		return "EQUAL"
	case COMMA:
		return "COMMA"
	default:
		return "NOTFOUND"
	}
}

//EOFChar The actual Character for the EOF symbol
var EOFChar = rune(EOF)
