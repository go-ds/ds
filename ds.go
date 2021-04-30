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
type KeyRangerFunc func(key interface{}) bool

// KeyRanger enables to traverse a Container with a KeyRangerFunc.
type KeyRanger interface {
	// RangeWithKey iterates a Container with a KeyRangerFunc.
	// Stop iterating if the KeyRangerFunc returns false.
	RangeWithKey(KeyRangerFunc)
}

// KVRangerFunc is an iteration function for ranging a KVRanger.
type KVRangerFunc func(key, value interface{}) bool

// KVRanger enables to traverse a Container with a KVRangerFunc.
type KVRanger interface {
	// RangeKV iterates a Container with a KVRangerFunc.
	// Stop iterating if the KVRangerFunc returns false.
	RangeKV(KVRangerFunc)
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
	// Push adds an element to the top of Stack.
	Push(interface{})
	// Pop ejects the most recently added element that was not yet removed
	// and removes it.
	Pop() interface{}
}

// Queue represents a first-in-first-out (FIFO) data structure.
type Queue interface {
	Container
	Peeker
	// Push appends an element to the end of Queue.
	Push(interface{})
	// Pop ejects the start element of Queue and removes it.
	Pop() interface{}
}

// PriorityQueue is an abstract data structure similar to a regular
// Queue in which each element additionally has a "priority"
// associated with it.
//
// In a priority queue, an element with high priority is served before
// an element with low priority.
type PriorityQueue interface {
	Container
	Peeker
	// Push appends an element to the PriorityQueue.
	Push(interface{})
	// Pop ejects the highest "priority" element of PriorityQueue, and
	// removes it.
	Pop() interface{}
}

// MonotoneQueue is a variant of the priority queue abstract structure
// in which the priorities of extracted elements are required to form a
// monotonic sequence. That is, for a priority queue in which each
// successively extracted item is the one with the minimum priority (a min-heap),
// the minimum priority should be monotonically increasing. Conversely for
// a max-heap the maximum priority should be monotonically decreasing.
//
// A necessary and sufficient condition on a MonotoneQueue is that one never
// attempts to add an element with lower priority than the most recently
// extracted one.
type MonotoneQueue interface {
	Container
	Peeker
	// Push appends an element to the end of MonotoneQueue if it makes the
	// queue still monotony, and removes lower priority elements if necessary.
	Push(interface{})
	// Pop ejects the start element of MonotoneQueue and removes it.
	Pop() interface{}
}

// Set is an abstract data structure that can store unique values,
// without any particular order.
type Set interface {
	Container
	// Add adds the element to Set, if it is not present already.
	Add(interface{}) Set
	// Has checks whether the element is in the Set.
	Has(interface{}) bool
	// Delete removes the element from Set, if it is present.
	Delete(interface{})
}

// Map is an abstract data structure composed of a Container of
// (key, value) pairs, such that each possible key appears at most
// once in the Container.
type Map interface {
	Container
	// Add adds a new (key,value) pair to the Map, mapping
	// the new key to its new value.
	Add(interface{}, interface{}) Map
	// Get finds the value (if any) that is bound to a given key.
	Get(interface{}) (interface{}, bool)
	// Has checks whether the key is in the Map.
	Has(interface{}) bool
	// Delete removes a (key,value) pair from the Map, unmapping
	// a given key from its value.
	Delete(interface{})
}

// Tree is an abstract data structure that simulates a hierarchical
// tree structure, with a root value and subtrees of children with
// a parent node, represented as a set of linked nodes.
type Tree interface {
	Container
	// Size is the numbers of nodes in the Tree, aka degree.
	Size() int
	// Height is the length of the longest downward path to a leaf
	// from the root.
	Height() int
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
	PopFront() (interface{}, bool)
	// Append appends new elements to the end of a Slice.
	Append(...interface{}) Slice
	// Prepend inserts new elements at the start of a Slice.
	Prepend(...interface{}) Slice
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

// Comparer is an interface has a Compare method for a ordered Container. It can carry
// value and is comparable to other Comparer items.
type Comparer interface {
	// Compare compares itself to other item and returns:
	//	negative	self  < other
	//	zero		self == other
	//	positive	self  > other
	Compare(other Comparer) int
}
