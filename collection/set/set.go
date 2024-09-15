package set

import (
	"iter"
)

type empty struct{}

// Set is set implemented by map
type Set[T comparable] map[T]empty

// New creates new set
func New[T comparable](values ...T) Set[T] {
	s := Set[T]{}
	for _, v := range values {
		s[v] = empty{}
	}
	return s
}

// Collect collects items into a Set.
func Collect[T comparable](seq iter.Seq[T]) Set[T] {
	s := Set[T]{}
	for v := range seq {
		s[v] = empty{}
	}
	return s
}

// CollectWithError collects items into a Set. If error occurred, return a nil set, and an error.
func CollectWithError[T comparable](seq iter.Seq2[T, error]) (Set[T], error) {
	s := Set[T]{}
	for v, err := range seq {
		if err != nil {
			return nil, err
		}
		s[v] = empty{}
	}
	return s, nil
}

// NewWithSize creates new set with expected size.
// A empty set is allocated with enough space to hold the specified number of elements.
func NewWithSize[T comparable](size int) Set[T] {
	s := make(Set[T], size)
	return s
}

// Contains reports is given value exists in this set
func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

// Add adds new element to set
func (s Set[T]) Add(v T) {
	s[v] = empty{}
}

// AddAll adds all values to set
func (s Set[T]) AddAll(values ...T) {
	for _, v := range values {
		s[v] = empty{}
	}
}

// Extend adds all values to set
func (s Set[T]) Extend(s2 Set[T]) {
	for v := range s2 {
		s[v] = empty{}
	}
}

// Remove removes element from set if it exists.
func (s Set[T]) Remove(v T) {
	delete(s, v)
}

// Remove removes elements from set if it exists.
func (s Set[T]) RemoveAll(values ...T) {
	for _, v := range values {
		delete(s, v)
	}
}

// Copy return a new set with same elements as the original one.
func (s Set[T]) Copy() Set[T] {
	ns := make(Set[T], len(s))
	for v, _ := range s {
		ns[v] = empty{}
	}
	return ns
}

// Values returns all values in Set as a [iter.Seq].
func (s Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for k := range s {
			if !yield(k) {
				break
			}
		}
	}
}

// Values returns all values in Set as a [iter.Seq].
//
// Deprecated: use [Set.All] instead.
func (s Set[T]) Values() iter.Seq[T] {
	return s.All()
}

// ToSlice return a slice contains the elements in the set
func (s Set[T]) ToSlice() []T {
	slice := make([]T, len(s))
	for v, _ := range s {
		slice = append(slice, v)
	}
	return slice
}

// Size returns the element count of set.
func (s Set[T]) Size() int {
	return len(s)
}

// Union returns a new set contains elements in one of sets.
func (s Set[T]) Union(s2 Set[T]) Set[T] {
	r := make(Set[T])
	r.Extend(s)
	r.Extend(s2)
	return s
}

// Intersection returns a new set contains element in both sets.
func (s Set[T]) Intersection(s2 Set[T]) Set[T] {
	r := make(Set[T])
	for v := range s {
		if s2.Contains(v) {
			r.Add(v)
		}
	}
	return r
}

// Difference returns a set contains elements in s but not in s2
func (s Set[T]) Difference(s2 Set[T]) Set[T] {
	r := make(Set[T])
	for v := range s {
		if !s2.Contains(v) {
			r.Add(v)
		}
	}
	return r
}
