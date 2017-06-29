package ast

import (
	"testing"

	"github.com/afrase/Gengo/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

func TestEmptyProgram(t *testing.T) {
	program := &Program{}

	if program.String() != "" {
		t.Errorf("got=%q expected=%s", program.String(), "")
	}
	if program.TokenLiteral() != "" {
		t.Errorf("got=%q expected=%s", program.TokenLiteral(), "")
	}
}

func TestSingleStatementProgram(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&ReturnStatement{
				Token: token.Token{Type: token.RETURN, Literal: "return"},
				ReturnValue: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "foo"},
					Value: "foo",
				},
			},
		},
	}

	tests := []struct {
		expected string
		actual   string
	}{
		{"return foo;", program.String()},
		{"return", program.TokenLiteral()},
	}

	for _, tt := range tests {
		if tt.expected != tt.actual {
			t.Errorf("got=%q expected=%s", tt.actual, tt.expected)
		}
	}
}

func TestExpressionStatement(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&ExpressionStatement{
				Token: token.Token{Type: token.LPAREN, Literal: "("},
				Expression: &PrefixExpression{
					Operator: "!",
					Token:    token.Token{Type: token.BANG, Literal: "!"},
					Right: &Boolean{
						Value: true,
						Token: token.Token{Type: token.TRUE, Literal: "true"},
					},
				},
			},
		},
	}

	tests := []struct {
		expected string
		actual   string
	}{
		{"(!true)", program.String()},
		{"(", program.Statements[0].TokenLiteral()},
	}

	for _, tt := range tests {
		if tt.expected != tt.actual {
			t.Errorf("got=%q expected=%s", tt.actual, tt.expected)
		}
	}
}

func TestFunctionLiteral(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&ExpressionStatement{
				Token: token.Token{Type: token.IDENT, Literal: "fn"},
				Expression: &FunctionLiteral{
					Token: token.Token{Type: token.IDENT, Literal: "add"},
					Parameters: []*Identifier{
						&Identifier{
							Token: token.Token{Type: token.IDENT, Literal: "a"},
							Value: "a",
						},
						&Identifier{
							Token: token.Token{Type: token.IDENT, Literal: "b"},
							Value: "b",
						},
					},
					Body: &BlockStatement{
						Token: token.Token{Type: token.LBRACE, Literal: "{"},
						Statements: []Statement{
							&ReturnStatement{
								Token: token.Token{Type: token.RETURN, Literal: "return"},
								ReturnValue: &InfixExpression{
									Token: token.Token{Type: token.PLUS, Literal: "+"},
									Left: &Identifier{
										Token: token.Token{Type: token.IDENT, Literal: "a"},
										Value: "a",
									},
									Right: &Identifier{
										Token: token.Token{Type: token.IDENT, Literal: "b"},
										Value: "b",
									},
									Operator: "+",
								},
							},
						},
					},
				},
			},
		},
	}

	tests := []struct {
		expected string
		actual   string
	}{
		{"add(a, b) { return (a + b); }", program.String()},
		{"fn", program.Statements[0].TokenLiteral()},
		{"add", program.Statements[0].(*ExpressionStatement).Expression.(*FunctionLiteral).Token.Literal},
	}

	for _, tt := range tests {
		if tt.expected != tt.actual {
			t.Errorf("got=%q expected=%s", tt.actual, tt.expected)
		}
	}
}
