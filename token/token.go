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
	// STRING A sequence of characters
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

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
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
