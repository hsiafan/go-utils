package deque

import "iter"

var _ Deque[string] = (*ArrayDeque[string])(nil)

// ArrayDeque represents a [Deque] implemented using a dynamic array (circular buffer).
type ArrayDeque[T any] struct {
	data  []T
	front int
	rear  int
	size  int
}

// NewArrayDeque creates a new empty deque.
func NewArrayDeque[T any]() *ArrayDeque[T] {
	return &ArrayDeque[T]{
		data: make([]T, 16),
		rear: -1,
	}
}

// NewArrayDequeWithSize creates a new empty deque with a specified initial capacity.
func NewArrayDequeWithSize[T any](capacity int) *ArrayDeque[T] {
	return &ArrayDeque[T]{
		data: make([]T, capacity),
		rear: -1,
	}
}

// PushFront adds an element to the front of the deque.
func (d *ArrayDeque[T]) PushFront(value T) {
	if d.size == len(d.data) {
		d.resize(len(d.data) * 2)
	}
	d.front = (d.front - 1 + len(d.data)) % len(d.data)
	d.data[d.front] = value
	d.size++
}

// PushBack adds an element to the back of the deque.
func (d *ArrayDeque[T]) PushBack(value T) {
	if d.size == len(d.data) {
		d.resize(len(d.data) * 2)
	}
	d.rear = (d.rear + 1) % len(d.data)
	d.data[d.rear] = value
	d.size++
}

// PopFront removes and returns the element from the front of the deque.
func (d *ArrayDeque[T]) PopFront() (T, bool) {
	var zero T
	if d.size == 0 {
		return zero, false
	}
	value := d.data[d.front]
	d.front = (d.front + 1) % len(d.data)
	d.size--
	if d.size > 0 && d.size <= len(d.data)/4 {
		d.resize(len(d.data) / 2)
	}
	return value, true
}

// PopBack removes and returns the element from the back of the deque.
func (d *ArrayDeque[T]) PopBack() (T, bool) {
	var zero T
	if d.size == 0 {
		return zero, false
	}
	value := d.data[d.rear]
	d.rear = (d.rear - 1 + len(d.data)) % len(d.data)
	d.size--
	if d.size > 0 && d.size <= len(d.data)/4 {
		d.resize(len(d.data) / 2)
	}
	return value, true
}

// resize resizes the underlying array to a new capacity.
func (d *ArrayDeque[T]) resize(newCapacity int) {
	newData := make([]T, newCapacity)
	for i := 0; i < d.size; i++ {
		newData[i] = d.data[(d.front+i)%len(d.data)]
	}
	d.data = newData
	d.front = 0
	d.rear = d.size - 1
}

// Values returns all values in current Deque as a [iter.Seq].
func (d *ArrayDeque[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := 0; i < d.size; i++ {
			v := d.data[(d.front+i)%len(d.data)]
			if !yield(v) {
				break
			}
		}
	}
}

// Size returns the number of elements in the deque.
func (d *ArrayDeque[T]) Size() int {
	return d.size
}
