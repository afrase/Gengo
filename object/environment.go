package object

// Environment associates strings with objects
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get the value of name
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set the value of name to val
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

// NewEnvironment returns a new Environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// NewEnclosedEnvironment returns a new Environment with outer set.
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}
