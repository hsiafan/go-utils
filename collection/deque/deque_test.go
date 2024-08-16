package deque

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeque(t *testing.T) {
	deque := New[int]()

	// Test initial state
	assert.Equal(t, 0, deque.Size(), "Initial size should be 0")

	// Test Append and Prepend
	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushFront(0)
	assert.Equal(t, 3, deque.Size(), "Size should be 3 after adding three elements")

	assert.Equal(t, []int{0, 1, 2}, slices.Collect(deque.Values()))

	// Test PopFront
	value, ok := deque.PopFront()
	assert.True(t, ok, "PopFront should return true when deque is not empty")
	assert.Equal(t, 0, value, "PopFront should return the first element, which is 0")

	// Test PopBack
	value, ok = deque.PopBack()
	assert.True(t, ok, "PopBack should return true when deque is not empty")
	assert.Equal(t, 2, value, "PopBack should return the last element, which is 2")

	// Test Size after pops
	assert.Equal(t, 1, deque.Size(), "Size should be 1 after two pops")

	// Test Pop on empty deque
	_, ok = deque.PopFront()
	assert.True(t, ok, "PopFront should succeed")
	_, ok = deque.PopFront()
	assert.False(t, ok, "PopFront should return false when deque is empty")
	_, ok = deque.PopBack()
	assert.False(t, ok, "PopBack should return false when deque is empty")

	// Test resizing
	for i := 0; i < 8; i++ {
		deque.PushBack(i)
	}
	assert.Equal(t, 8, deque.Size(), "Size should be 8 after adding elements")
	for i := 0; i < 8; i++ {
		_, ok = deque.PopFront()
		assert.True(t, ok, "PopFront should return true for non-empty deque")
	}
	assert.Equal(t, 0, deque.Size(), "Size should be 0 after popping all elements")
}

func TestDequeWithString(t *testing.T) {
	deque := New[string]()

	// Test Append and Prepend with strings
	deque.PushBack("world")
	deque.PushFront("hello")
	assert.Equal(t, 2, deque.Size(), "Size should be 2 after adding two elements")

	// Test PopFront with strings
	value, ok := deque.PopFront()
	assert.True(t, ok, "PopFront should return true when deque is not empty")
	assert.Equal(t, "hello", value, "PopFront should return the first element, which is 'hello'")

	// Test PopBack with strings
	value, ok = deque.PopBack()
	assert.True(t, ok, "PopBack should return true when deque is not empty")
	assert.Equal(t, "world", value, "PopBack should return the last element, which is 'world'")
}

func TestDequeEmpty(t *testing.T) {
	deque := New[int]()

	// Test PopFront on empty deque
	value, ok := deque.PopFront()
	assert.False(t, ok, "PopFront should return false when deque is empty")
	assert.Equal(t, 0, value, "PopFront should return zero value when deque is empty")

	// Test PopBack on empty deque
	value, ok = deque.PopBack()
	assert.False(t, ok, "PopBack should return false when deque is empty")
	assert.Equal(t, 0, value, "PopBack should return zero value when deque is empty")
}

func TestResizing(t *testing.T) {

	deque := New[int]()

	// Test resizing up
	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushFront(0)
	assert.Equal(t, 3, deque.Size(), "Size should be 3 after resizing up")

	// Test contents after resizing
	value, ok := deque.PopFront()
	assert.True(t, ok, "PopFront should return true when deque is not empty")
	assert.Equal(t, 0, value, "PopFront should return the first element, which is 0")
	value, ok = deque.PopFront()
	assert.True(t, ok, "PopFront should return true when deque is not empty")
	assert.Equal(t, 1, value, "PopFront should return the first element, which is 1")
	value, ok = deque.PopFront()
	assert.True(t, ok, "PopFront should return true when deque is not empty")
	assert.Equal(t, 2, value, "PopFront should return the first element, which is 2")

	// Test resizing down
	for i := 0; i < 16; i++ {
		deque.PushBack(i)
	}
	for i := 0; i < 14; i++ {
		_, _ = deque.PopFront()
	}
	assert.Equal(t, 2, deque.Size(), "Size should be 2 after resizing down")
}
