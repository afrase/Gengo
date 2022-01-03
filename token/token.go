package token

// Type A string representing a token
type Type string

// Token A token in the AST.
type Token struct {
	Type    Type
	Literal string
}

const (
	// ILLEGAL An illegal token.
	ILLEGAL = "ILLEGAL"
	// EOF The end of the file.
	EOF = "EOF"
	// IDENT An identifier. (e.g. a, b, foobar, x, y, ...)
	IDENT = "IDENT"
	// INT An integer. (e.g. 1, 2, 9539, ...)
	INT = "INT"
	// FLOAT A float. (e.g. 0.1, 3.1459, ...)
	FLOAT = "FLOAT"
	// STRING A sequence of characters. (e.g. foo, bar, foo1bar, ...)
	STRING = "STRING"

	// ASSIGN The token for assignment.
	ASSIGN = "="
	// PLUS The token for addition.
	PLUS = "+"
	// MINUS The token for subtraction.
	MINUS = "-"
	// BANG The token for the "bang" operator.
	BANG = "!"
	// ASTERISK The token for multiplication.
	ASTERISK = "*"
	// SLASH The token for division.
	SLASH = "/"
	// POW The token for Power of.
	POW = "**"
	// EQ The token used to check for equality.
	EQ = "=="
	// NOTEQ The token used to check for the opposite of equality.
	NOTEQ = "!="

	// LT The token for less-than.
	LT = "<"
	// GT The token for greater-than.
	GT = ">"

	// COMMA A token used as a delimiter.
	COMMA = ","
	// SEMICOLON A token used to denote the end of a statement.
	SEMICOLON = ";"
	// COLON A token used for hash literals.
	COLON = ":"

	// LPAREN The token for an opening parenthesis.
	LPAREN = "("
	// RPAREN The token for a closing parenthesis.
	RPAREN = ")"
	// LBRACE The token for an opening curly brace.
	LBRACE = "{"
	// RBRACE The token for a closing curly brace.
	RBRACE = "}"
	// LBRACKET The token for an opening bracket.
	LBRACKET = "["
	// RBRACKET The token for a closing bracket.
	RBRACKET = "]"

	// Keywords

	// FUNCTION The token for a function.
	FUNCTION = "FUNCTION"
	// LET The token for identifier assignment.
	LET = "LET"
	// TRUE The token for a "true" value.
	TRUE = "TRUE"
	// FALSE The token for a "false" value.
	FALSE = "FALSE"
	// IF The token for an "if" statement.
	IF = "IF"
	// ELSE The token for an "else" statement.
	ELSE = "ELSE"
	// RETURN The token for the return of a function.
	RETURN = "RETURN"
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent Convert a string to a TokenType
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
