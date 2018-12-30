# Parser Rules

## Notation

This doc uses EBNF to describe the grammar rules. A lowercase string (eg `program`) denotes a rule. When wrapped in quotes (eg `";") it denotes a terminal token. When all uppercase, denotes a token.

```Text
program = { statement } | "EOF"
statement = expr ";"
expr =    decl
        | assignment
        | binop
        | func_call
        | pipe
        | strin
        | integer
        | character
        | if
        | ifelseif
        | pattern
decl = const | let
const = CONST varname
let = LET varname
assignment = (varname | decl) "=" expr
varname = VARNAME
binop = (expr "*" expr) | (expr "+" expr) | (expr "-" expr) | (expr "/" expr) (expr "**" expr)
pipe = func_call ["(" arglist ")"] "|>" { func_call }
func_call = varname "(" arglist ")"
```