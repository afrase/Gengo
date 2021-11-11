package ast

import (
	"bytes"
	"strings"

	"github.com/afrase/Gengo/token"
)

// Node A single node in the AST.
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement represents a statement node.
type Statement interface {
	Node
	// Used to help the compiler tell the difference between a Statement and
	// an Expression.
	statementNode()
}

// Expression represents an expression node.
type Expression interface {
	Node
	expressionNode()
}

// Program consists of an array of statement nodes.
type Program struct {
	Statements []Statement
}

// TokenLiteral The literal value of the token. Used only for debugging and testing.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// LetStatement Are used to assign a value to an identifier.
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// TokenLiteral The literal value of the token.
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " " + ls.Name.String() + " = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

func (ls *LetStatement) statementNode() {}

// Identifier The name a value is assigned to.
type Identifier struct {
	Token token.Token
	Value string
}

// TokenLiteral The literal value of the token.
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) expressionNode() {}

// StringLiteral The value of a string.
type StringLiteral struct {
	Token token.Token
	Value string
}

// TokenLiteral The literal value of the token.
func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token.Literal
}

func (sl *StringLiteral) String() string {
	return sl.Token.Literal
}

func (sl *StringLiteral) expressionNode() {}

// ReturnStatement Used to return a value from a function.
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

// TokenLiteral The literal value of the token.
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) statementNode() {}

// ExpressionStatement An expression is something that returns a value.
type ExpressionStatement struct {
	Token      token.Token // first token of the expression
	Expression Expression
}

// TokenLiteral The literal value of the token.
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (es *ExpressionStatement) statementNode() {}

// IntegerLiteral A literal integer
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

// TokenLiteral The literal value of the token.
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) expressionNode() {
}

// FloatLiteral A literal float
type FloatLiteral struct {
	Token token.Token
	Value float64
}

// TokenLiteral The literal value of the token.
func (fl *FloatLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FloatLiteral) String() string {
	return fl.Token.Literal
}

func (fl *FloatLiteral) expressionNode() {
}

// PrefixExpression Prefixes an expression. Like `-5`, `!true`, `-add(1,2)`
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

// TokenLiteral The literal value of the token.
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

func (pe *PrefixExpression) expressionNode() {}

// InfixExpression An infix expression
type InfixExpression struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

// TokenLiteral The literal value of the token.
func (oe *InfixExpression) TokenLiteral() string {
	return oe.Token.Literal
}

func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}

func (oe *InfixExpression) expressionNode() {}

// Boolean A boolean
type Boolean struct {
	Token token.Token
	Value bool
}

// TokenLiteral The literal value of the token.
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Token.Literal
}

func (b *Boolean) expressionNode() {}

// IfExpression An if expression
type IfExpression struct {
	Token       token.Token // the 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

// TokenLiteral The literal value of the token.
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if (")
	out.WriteString(ie.Condition.String())
	out.WriteString(") ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString(" else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

func (ie *IfExpression) expressionNode() {}

// BlockStatement A block of code
type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement
}

// TokenLiteral The literal value of the token.
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	out.WriteString("{ ")
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	out.WriteString(" }")

	return out.String()
}

func (bs *BlockStatement) statementNode() {}

// FunctionLiteral A function literal
type FunctionLiteral struct {
	Token      token.Token // The 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

// TokenLiteral The literal value of the token.
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	var params []string
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

func (fl *FunctionLiteral) expressionNode() {}

// CallExpression A call expression
type CallExpression struct {
	Token     token.Token // the '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

// TokenLiteral The literal value of the token.
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *CallExpression) String() string {
	var out bytes.Buffer

	var args []string
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

func (ce *CallExpression) expressionNode() {}
