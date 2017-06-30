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
		{"foo", program.Statements[0].(*ReturnStatement).ReturnValue.TokenLiteral()},
	}

	for _, tt := range tests {
		if tt.expected != tt.actual {
			t.Errorf("got=%q, expected=%s", tt.actual, tt.expected)
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
						{
							Token: token.Token{Type: token.IDENT, Literal: "a"},
							Value: "a",
						},
						{
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
	}

	for _, tt := range tests {
		if tt.expected != tt.actual {
			t.Errorf("got=%q expected=%s", tt.actual, tt.expected)
		}
	}

	stmt, ok := program.Statements[0].(*ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ExpressionStatement. got=%T", program.Statements[0])
	}
	fl, ok := stmt.Expression.(*FunctionLiteral)
	if !ok {
		t.Fatalf("exp not *FunctionLiteral. got=%T", stmt.Expression)
	}

	if !testIdentifier(t, fl.Parameters[0], "a") {
		return
	}
	if !testIdentifier(t, fl.Parameters[1], "b") {
		return
	}

	rs, ok := fl.Body.Statements[0].(*ReturnStatement)
	if !ok {
		t.Fatalf("statement not *ReturnStatement. got=%T", fl.Body.Statements[0])
	}
	ie, ok := rs.ReturnValue.(*InfixExpression)
	if !ok {
		t.Fatalf("return value not *InfixExpression. got=%T", rs.ReturnValue)
	}
	if ie.String() != "(a + b)" {
		t.Errorf("InfixExpression wrong, expected=%s, got=%s", "(a + b)", ie.String())
	}
	if ie.Operator != "+" {
		t.Errorf("operator is not '+'. got=%s", ie.Operator)
	}
	if ie.TokenLiteral() != "+" {
		t.Errorf("InfixExpression.TokenLiteral expected=%s, got=%s", "+", ie.TokenLiteral())
	}
}

func TestReturnStatement(t *testing.T) {
	rt := ReturnStatement{
		Token: token.Token{Type: token.RETURN, Literal: "return"},
		ReturnValue: &IntegerLiteral{
			Token: token.Token{Type: token.INT, Literal: "5"},
			Value: 5,
		},
	}

	tests := []struct {
		expected string
		actual   string
	}{
		{"return 5;", rt.String()},
		{"return", rt.TokenLiteral()},
		{"5", rt.ReturnValue.(*IntegerLiteral).String()},
		{"5", rt.ReturnValue.(*IntegerLiteral).TokenLiteral()},
	}

	for _, tt := range tests {
		if tt.expected != tt.actual {
			t.Errorf("got=%v, expected=%v", tt.actual, tt.expected)
		}
	}
}

func TestIfExpression(t *testing.T) {
	ie := IfExpression{
		Token: token.Token{Type: token.IF, Literal: "if"},
		Condition: &Boolean{
			Token: token.Token{Type: token.TRUE, Literal: "true"},
			Value: true,
		},
		Consequence: &BlockStatement{
			Token: token.Token{Type: token.LBRACE, Literal: "{"},
			Statements: []Statement{
				&ReturnStatement{
					Token: token.Token{Type: token.RETURN, Literal: "return"},
					ReturnValue: &IntegerLiteral{
						Token: token.Token{Type: token.INT, Literal: "5"},
					},
				},
			},
		},
		Alternative: &BlockStatement{
			Token: token.Token{Type: token.LBRACE, Literal: "{"},
			Statements: []Statement{
				&ReturnStatement{
					Token: token.Token{Type: token.RETURN, Literal: "return"},
					ReturnValue: &IntegerLiteral{
						Token: token.Token{Type: token.INT, Literal: "10"},
					},
				},
			},
		},
	}

	tests := []struct {
		expected string
		actual   string
	}{
		{"if (true) { return 5; } else { return 10; }", ie.String()},
		{"if", ie.TokenLiteral()},
	}

	for _, tt := range tests {
		if tt.expected != tt.actual {
			t.Errorf("got=%v, expected=%v", tt.actual, tt.expected)
		}
	}
}

func TestCallExpression(t *testing.T) {
	ce := CallExpression{
		Token:     token.Token{Type: token.LPAREN, Literal: "("},
		Arguments: []Expression{},
		Function: &FunctionLiteral{
			Token: token.Token{Type: token.IDENT, Literal: "foo"},
			Body:  &BlockStatement{},
		},
	}

	tests := []struct {
		expected string
		actual   string
	}{
		{"foo() {  }()", ce.String()},
		{"(", ce.TokenLiteral()},
	}

	for _, tt := range tests {
		if tt.expected != tt.actual {
			t.Errorf("got=%v, expected=%v", tt.actual, tt.expected)
		}
	}
}

func testIdentifier(t *testing.T, exp Expression, value string) bool {
	ident, ok := exp.(*Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}

	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}

	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s. got=%s", value, ident.TokenLiteral())
		return false
	}

	return true
}
