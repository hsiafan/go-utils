package set

import (
	"iter"

	"github.com/hsiafan/go-utils/collection/maps2"
)

// LinkedSet is a set that keeps the order of elements.
type LinkedSet[T comparable] maps2.LinkedMap[T, empty]

// New creates a new LinkedSet.
func NewLinkedSet[T comparable](values ...T) *LinkedSet[T] {
	m := maps2.NewLinkedMap[T, empty]()
	for _, v := range values {
		m.Put(v, empty{})
	}
	return (*LinkedSet[T])(m)
}

// Contains reports is given value exists in this set
func (s *LinkedSet[T]) Contains(v T) bool {
	return s.m().Contains(v)
}

// Add adds new element to set
func (s *LinkedSet[T]) Add(v T) {
	s.m().Put(v, empty{})
}

// AddAll adds all values to set
func (s *LinkedSet[T]) AddAll(values ...T) {
	for _, v := range values {
		s.m().Put(v, empty{})
	}
}

// AddSet adds all values to set
func (s *LinkedSet[T]) AddSet(s2 *LinkedSet[T]) {
	for v := range s2.All() {
		s.m().Put(v, empty{})
	}
}

// Remove removes element from set if it exists.
func (s *LinkedSet[T]) Remove(v T) {
	s.m().Remove(v)
}

// Remove removes all elements from set.
func (s *LinkedSet[T]) RemoveAll(values ...T) {
	for _, v := range values {
		s.m().Remove(v)
	}
}

// Copy return a new set with same elements as the original one.
func (s *LinkedSet[T]) Copy() *LinkedSet[T] {
	return (*LinkedSet[T])(s.m().Copy())
}

// Values returns all values in Set as a [iter.Seq].
func (s *LinkedSet[T]) All() iter.Seq[T] {
	return s.m().Keys()
}

// ToSlice return a slice contains the elements in the set.
func (s *LinkedSet[T]) ToSlice() []T {
	slice := make([]T, s.Size())
	for v := range s.All() {
		slice = append(slice, v)
	}
	return slice
}

// Size returns the element count of set.
func (s *LinkedSet[T]) Size() int {
	return s.m().Size()
}

// Clear removes all elements from set.
func (s *LinkedSet[T]) Clear() {
	s.m().Clear()
}

func (s *LinkedSet[T]) m() *maps2.LinkedMap[T, empty] {
	return (*maps2.LinkedMap[T, empty])(s)
}
