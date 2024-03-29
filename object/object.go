package object

import (
	"bytes"
	"fmt"
	"strings"

	"Gengo/ast"
)

const (
	// STRING Represents a string object.
	STRING = "STRING"
	// INTEGER Represents an integer object.
	INTEGER = "INTEGER"
	// FLOAT Represents a float object.
	FLOAT = "FLOAT"
	// BOOLEAN Represents a boolean object.
	BOOLEAN = "BOOLEAN"
	// ARRAY Represents an array object.
	ARRAY = "ARRAY"
	// NULL Represents a null object.
	NULL = "NULL"
	// FUNCTION Represents a function.
	FUNCTION = "FUNCTION"
	// BUILTIN Represents a built-in function.
	BUILTIN = "BUILTIN"
	// RETURN_VALUE Represents the return value.
	RETURN_VALUE = "RETURN_VALUE"
	// ERROR Represents an error object.
	ERROR = "ERROR"
)

// ObjectType The base object type.
type ObjectType string

// Object The interface of the object.
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer The integer type.
type Integer struct {
	Value int64
}

// Inspect A String of the value.
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Type The object's type.
func (i *Integer) Type() ObjectType {
	return INTEGER
}

// Float A float type.
type Float struct {
	Value float64
}

// Inspect A String of the value.
func (f *Float) Inspect() string {
	return fmt.Sprintf("%f", f.Value)
}

// Type The object's type.
func (f *Float) Type() ObjectType {
	return FLOAT
}

// Boolean type.
type Boolean struct {
	Value bool
}

// Inspect A string of the value.
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// Type The object's type.
func (b *Boolean) Type() ObjectType {
	return BOOLEAN
}

// Null type.
type Null struct{}

// Inspect A string of the type.
func (n *Null) Inspect() string {
	return "null"
}

// Type The object's type.
func (n *Null) Type() ObjectType {
	return NULL
}

// ReturnValue A value returned by a call.
type ReturnValue struct {
	Value Object
}

// Type The object's type.
func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE
}

// Inspect A string of the type.
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

// Error type.
type Error struct {
	Message string
}

// Type The object's type.
func (e *Error) Type() ObjectType {
	return ERROR
}

// Inspect A string of the type.
func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

// String type.
type String struct {
	Value string
}

// Type The object's type.
func (s *String) Type() ObjectType {
	return STRING
}

// Inspect A string of the type.
func (s *String) Inspect() string {
	return s.Value
}

// Array type.
type Array struct {
	Elements []Object
}

// Type The object's type.
func (ao *Array) Type() ObjectType {
	return ARRAY
}

// Inspect A string of the type.
func (ao *Array) Inspect() string {
	var elements []string

	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	return "[" + strings.Join(elements, ", ") + "]"
}

// Function type.
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type The object's type.
func (f *Function) Type() ObjectType {
	return FUNCTION
}

// Inspect A string of the type.
func (f *Function) Inspect() string {
	var out bytes.Buffer
	var params []string

	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

// BuiltinFunction type.
type BuiltinFunction func(args ...Object) Object

// Builtin type.
type Builtin struct {
	Fn BuiltinFunction
}

// Type The object's type.
func (b *Builtin) Type() ObjectType {
	return BUILTIN
}

// Inspect A string of the type.
func (b *Builtin) Inspect() string {
	return "builtin function"
}
