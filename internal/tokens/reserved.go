package tokens

//ReservedWords is the reserved words in arbor
var ReservedWords = []string{
	"let",
	"func",
	"string",
	"char",
	"done",
	"return",
	"const",
	"if",
	"else",
	"float",
	"number",
	"import",
	"as",
}

//FindKeyword finds and returns a token associated with that key word, if that isn't a key word, it returns the NOTFOUND token
func FindKeyword(str string) Token {
	for ndx, reserved := range ReservedWords {
		if reserved == str {
			return Token(ndx + int(LET))
		}
	}
	return NOTFOUND
}
