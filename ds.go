// Package gods provides core interfaces and functions for data structures.
package gods

// Container is a basic interface that all data structures implement.
type Container interface {
	// Empty indicates if the Container is empty.
	Empty() bool
	// Size retrieves Container size.
	Size() int
	// Clear resets Container, it will be empty with size 0.
	Clear()
}

// IndexRangerFunc is an iteration function for ranging an IndexRanger.
type IndexRangerFunc func(index int, value interface{}) bool

// IndexRanger enables to traverse a Container with an IndexRangerFunc.
type IndexRanger interface {
	// RangeWithIndex iterates a Container with an IndexRangerFunc.
	// Stop iterating if the IndexRangerFunc returns false.
	RangeWithIndex(IndexRangerFunc)
}

// KeyRangerFunc is an iteration function for ranging a KeyRanger.
type KeyRangerFunc func(key, value interface{}) bool

// KeyRanger enables to traverse a Container with a KeyRangerFunc.
type KeyRanger interface {
	// RangeWithKey iterates a Container with a KeyRangerFunc.
	// Stop iterating if the KeyRangerFunc returns false.
	RangeWithKey(KeyRangerFunc)
}

// Peeker gives access to the top without modifying the Container.
type Peeker interface {
	// Peek inspects topmost element of Container without modifying the Container.
	// (nil, false) are returned if the Container is empty.
	Peek() (interface{}, bool)
}

// Stack represents a last-in-first-out (LIFO) data structure provides
// the principal push and pop operations, as well as a method to peek at the
// top element on the stack.
type Stack interface {
	Container
	Peeker
	// Push adds an element to the Stack.
	Push(interface{})
	// Pop removes the most recently added element that was not yet removed.
	Pop() interface{}
}

// Queue represents a first-in-first-out (FIFO) data structure.
type Queue interface {
	Container
	// Push appends an element to the end of Queue.
	Push(interface{})
	// Pop removes the start element of Queue.
	Pop() interface{}
}
