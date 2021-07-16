// Code generated by "stringer -type=Token"; DO NOT EDIT.

package tokens

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[EOF-0]
	_ = x[RPAREN-1]
	_ = x[LPAREN-2]
	_ = x[VARNAME-3]
	_ = x[ARROW-4]
	_ = x[COLON-5]
	_ = x[DCOLON-6]
	_ = x[NUMBER-7]
	_ = x[FLOAT-8]
	_ = x[QUOTE-9]
	_ = x[DQOUTE-10]
	_ = x[CHARVAL-11]
	_ = x[STRINGVAL-12]
	_ = x[ERROR-13]
	_ = x[NEWLINE-14]
	_ = x[ARTHOP-15]
	_ = x[LOGICAL-16]
	_ = x[NOT-17]
	_ = x[BOOLEAN-18]
	_ = x[EQUAL-19]
	_ = x[COMMA-20]
	_ = x[PIPE-21]
	_ = x[SEMI-22]
	_ = x[RCURLY-23]
	_ = x[LCURLY-24]
	_ = x[COMPARISON-25]
	_ = x[LSQUARE-26]
	_ = x[RSQUARE-27]
	_ = x[DOT-28]
	_ = x[AT-29]
	_ = x[LET-30]
	_ = x[FUNC-31]
	_ = x[STRING-32]
	_ = x[CHAR-33]
	_ = x[DONE-34]
	_ = x[RETURN-35]
	_ = x[CONST-36]
	_ = x[IF-37]
	_ = x[ELSE-38]
	_ = x[FLOATWORD-39]
	_ = x[NUMBERWORD-40]
	_ = x[IMPORT-41]
	_ = x[AS-42]
	_ = x[SHAPE-43]
	_ = x[TYPE-44]
	_ = x[FROM-45]
	_ = x[INTERNAL-46]
	_ = x[PACKAGE-47]
	_ = x[NEW-48]
	_ = x[MATCH-49]
	_ = x[WHEN-50]
	_ = x[EXTENDS-51]
	_ = x[IMPLEMENTS-52]
	_ = x[CONTINUE-53]
	_ = x[WITH-54]
	_ = x[SIGNAL-55]
	_ = x[WARN-56]
	_ = x[FATAL-57]
	_ = x[TRY-58]
	_ = x[HANDLE-59]
	_ = x[SELF-60]
}

const _Token_name = "EOFRPARENLPARENVARNAMEARROWCOLONDCOLONNUMBERFLOATQUOTEDQOUTECHARVALSTRINGVALERRORNEWLINEARTHOPLOGICALNOTBOOLEANEQUALCOMMAPIPESEMIRCURLYLCURLYCOMPARISONLSQUARERSQUAREDOTATLETFUNCSTRINGCHARDONERETURNCONSTIFELSEFLOATWORDNUMBERWORDIMPORTASSHAPETYPEFROMINTERNALPACKAGENEWMATCHWHENEXTENDSIMPLEMENTSCONTINUEWITHSIGNALWARNFATALTRYHANDLESELF"

var _Token_index = [...]uint16{0, 3, 9, 15, 22, 27, 32, 38, 44, 49, 54, 60, 67, 76, 81, 88, 94, 101, 104, 111, 116, 121, 125, 129, 135, 141, 151, 158, 165, 168, 170, 173, 177, 183, 187, 191, 197, 202, 204, 208, 217, 227, 233, 235, 240, 244, 248, 256, 263, 266, 271, 275, 282, 292, 300, 304, 310, 314, 319, 322, 328, 332}

func (i Token) String() string {
	if i < 0 || i >= Token(len(_Token_index)-1) {
		return "Token(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Token_name[_Token_index[i]:_Token_index[i+1]]
}
