package deque

import (
	"container/list"
	"iter"
)

var _ Deque[string] = (*LinkedListDeque[string])(nil)

// LinkedListDeque represents a [Deque] implemented using a linked list.
type LinkedListDeque[T any] struct {
	l list.List
}

// NewLinkedListDeque creates a new LinkedListDeque
func NewLinkedListDeque[T any]() *LinkedListDeque[T] {
	q := &LinkedListDeque[T]{}
	q.l.Init()
	return q
}

// PopBack implements Deque.
func (q *LinkedListDeque[T]) PopBack() (T, bool) {
	e := q.l.Back()
	if e == nil {
		var zero T
		return zero, false
	}
	q.l.Remove(e)
	return e.Value.(T), true
}

// PopFront implements Deque.
func (q *LinkedListDeque[T]) PopFront() (T, bool) {
	e := q.l.Front()
	if e == nil {
		var zero T
		return zero, false
	}
	q.l.Remove(e)
	return e.Value.(T), true
}

// PushBack implements Deque.
func (q *LinkedListDeque[T]) PushBack(value T) {
	q.l.PushBack(value)
}

// PushFront implements Deque.
func (q *LinkedListDeque[T]) PushFront(value T) {
	q.l.PushFront(value)
}

// Size implements Deque.
func (q *LinkedListDeque[T]) Size() int {
	return q.l.Len()
}

// Values implements Deque.
func (q *LinkedListDeque[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := q.l.Front(); e != nil; e = e.Next() {
			if !yield(e.Value.(T)) {
				break
			}
		}
	}

}
