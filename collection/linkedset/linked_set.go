package linkedset

import (
	"iter"

	"github.com/hsiafan/go-utils/collection/linkedmap"
)

type empty struct{}

// Set is a set that keeps the order of elements.
type Set[T comparable] linkedmap.Map[T, empty]

// New creates a new LinkedSet.
func New[T comparable](values ...T) *Set[T] {
	m := linkedmap.New[T, empty]()
	for _, v := range values {
		m.Put(v, empty{})
	}
	return (*Set[T])(m)
}

// Contains reports is given value exists in this set
func (s *Set[T]) Contains(v T) bool {
	return s.m().Contains(v)
}

// Add adds new element to set
func (s *Set[T]) Add(v T) {
	s.m().Put(v, empty{})
}

// AddAll adds all values to set
func (s *Set[T]) AddAll(values ...T) {
	for _, v := range values {
		s.m().Put(v, empty{})
	}
}

// AddSet adds all values to set
func (s *Set[T]) AddSet(s2 *Set[T]) {
	for v := range s2.All() {
		s.m().Put(v, empty{})
	}
}

// Remove removes element from set if it exists.
func (s *Set[T]) Remove(v T) {
	s.m().Remove(v)
}

// Remove removes all elements from set.
func (s *Set[T]) RemoveAll(values ...T) {
	for _, v := range values {
		s.m().Remove(v)
	}
}

// Copy return a new set with same elements as the original one.
func (s *Set[T]) Copy() *Set[T] {
	return (*Set[T])(s.m().Copy())
}

// Values returns all values in Set as a [iter.Seq].
func (s *Set[T]) All() iter.Seq[T] {
	return s.m().Keys()
}

// ToSlice return a slice contains the elements in the set.
func (s *Set[T]) ToSlice() []T {
	slice := make([]T, s.Size())
	for v := range s.All() {
		slice = append(slice, v)
	}
	return slice
}

// Size returns the element count of set.
func (s *Set[T]) Size() int {
	return s.m().Size()
}

// Clear removes all elements from set.
func (s *Set[T]) Clear() {
	s.m().Clear()
}

func (s *Set[T]) m() *linkedmap.Map[T, empty] {
	return (*linkedmap.Map[T, empty])(s)
}
