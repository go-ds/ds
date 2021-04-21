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
	// Returns (nil, false) if the Container is empty.
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

// Slice is a slice wrapper which provides various handy methods.
type Slice interface {
	Container
	IndexRanger
	// Raw returns the raw slice of Slice.
	Raw() []interface{}
	// Pop removes the last element from a Slice and returns it.
	// Returns (nil, false) if there is no more element.
	Pop() (interface{}, bool)
	// PopFront removes the first element from a Slice and returns it.
	// Returns (nil, false) if there is no more element.
	PopFront()(interface{}, bool)
	// Append appends new elements to the end of a Slice.
	Append(... interface{}) Slice
	// Prepend inserts new elements at the start of a Slice.
	Prepend(... interface{}) Slice
	// Concat combines two Slices.
	Concat(slice Slice) Slice
	// Reverse reverses the elements in a Slice in place.
	Reverse() Slice
	// Sort sorts a Slice in place.
	Sort(compare func(raw []interface{}, i, j int) bool) Slice
	// Slice returns a copy of a section of a Slice.
	Slice(...int) Slice
	// Splice removes elements from a Slice and, if necessary, inserts
	// new elements in their place, returning the deleted elements.
	// Remove all elements after the start position(including start one)
	// if deleteCount is -1.
	Splice(start int, deleteCount int, elements ...interface{}) Slice
	// Map projects every element in Slice with the projection function
	// and returns a Slice that contains all the results.
	Map(project func(interface{}) interface{}) Slice
	// Filter returns the elements of a Slice that meet the condition
	// specified in a predicate function.
	Filter(predicate func(interface{}) bool) Slice
	// Reject returns the elements of a Slice that does not meet the
	// condition specified in a predicate function.
	Reject(predicate func(interface{}) bool) Slice
	// Every determines whether all the elements of a Slice satisfy the
	// specified predicate function.
	Every(predicate func(interface{}) bool) bool
	// Some determines whether the specified predicate function returns
	// true for any element of a Slice.
	Some(predicate func(interface{}) bool) bool
	// Reduce calls the specified callback function for all the elements in a Slice.
	// The return value of the callback function is the accumulated result, and is
	// provided as an argument in the next call to the callback function.
	Reduce(fn func(previousValue, currentValue interface{}, currentIndex int) interface{}, initialValue interface{}) interface{}
	// ReduceRight calls the specified callback function for all the elements in
	// a Slice, in descending order. The return value of the callback function is the
	// accumulated result, and is provided as an argument in the next call to the
	// callback function.
	ReduceRight(fn func(previousValue, currentValue interface{}, currentIndex int) interface{}, initialValue interface{}) interface{}
}

// Slicer can convert all elements in a Container to a Slice.
type Slicer interface {
	// ToSlice returns a Slice with all elements in the Container.
	ToSlice() Slice
}
