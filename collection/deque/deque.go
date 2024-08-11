package deque

import (
	"iter"
)

// Deque interface.
// Implements:
//  1. [ArrayDeque]
//  2. TBA
type Deque[T any] interface {
	// PushFront adds an element to the front of the deque.
	PushFront(value T)
	// PushBack adds an element to the back of the deque.
	PushBack(value T)
	// PopFront removes and returns the element from the front of the deque.
	PopFront() (T, bool)
	// PopBack removes and returns the element from the back of the deque.
	PopBack() (T, bool)
	// Values returns all values in current Deque as a [iter.Seq].
	Values() iter.Seq[T]
	// Size returns the number of elements in the deque.
	Size() int
}
