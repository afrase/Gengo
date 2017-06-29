package token

import (
	"testing"
)

func TestLookupIdent(t *testing.T) {
	tests := []struct {
		input    string
		expected Type
	}{
		{input: "let", expected: LET},
		{input: "fn", expected: FUNCTION},
		{input: "true", expected: TRUE},
		{input: "if", expected: IF},
		{input: "false", expected: FALSE},
		{input: "else", expected: ELSE},
		{input: "return", expected: RETURN},
		{input: "fooBar", expected: IDENT},
	}

	for i, tt := range tests {
		tok := LookupIdent(tt.input)

		if tok != tt.expected {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", i, tt.expected, tok)
		}
	}
}
