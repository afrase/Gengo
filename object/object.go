package object

import "fmt"

const (
	// STRING_OBJ Represents a string object
	STRING_OBJ = "STRING"
	// INTEGER_OBJ Represents an integer object
	INTEGER_OBJ = "INTEGER"
	// FLOAT_OBJ Represents a float object
	FLOAT_OBJ = "FLOAT"
	// BOOLEAN_OBJ Represents a boolean object
	BOOLEAN_OBJ = "BOOLEAN"
	// NULL_OBJ Represents a null object
	NULL_OBJ = "NULL"
	// RETURN_VALUE_OBJ Represents the return value
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	// ERROR_OBJ Represents an error object
	ERROR_OBJ = "ERROR"
)

// ObjectType The base object type
type ObjectType string

// Object The interface of the object
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer The integer type
type Integer struct {
	Value int64
}

// Inspect
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Type
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

// Float
type Float struct {
	Value float64
}

// Inspect
func (f *Float) Inspect() string {
	return fmt.Sprintf("%f", f.Value)
}

// Type
func (f *Float) Type() ObjectType {
	return FLOAT_OBJ
}

// Boolean
type Boolean struct {
	Value bool
}

// Inspect
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// Type
func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

// Null
type Null struct{}

// Inspect
func (n *Null) Inspect() string {
	return "null"
}

// Type
func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

// ReturnValue
type ReturnValue struct {
	Value Object
}

// Type
func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

// Inspect
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

// Error
type Error struct {
	Message string
}

// Type
func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}

// Inspect
func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

// String
type String struct {
	Value string
}

// Type
func (s *String) Type() ObjectType {
	return STRING_OBJ
}

// Inspect
func (s *String) Inspect() string {
	return s.Value
}
