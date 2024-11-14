package maps2

import (
	"iter"

	"github.com/hsiafan/go-utils/collection/pair"
	"github.com/hsiafan/go-utils/collection/set"
)

// UniqMultiMap can hold multi value of one key, the same value only store once
type UniqMultiMap[K comparable, V comparable] map[K]set.Set[V]

func NewUniq[K comparable, V comparable]() UniqMultiMap[K, V] {
	return UniqMultiMap[K, V]{}
}

func CollectUniq[K comparable, V comparable](seq iter.Seq2[K, V]) UniqMultiMap[K, V] {
	m := UniqMultiMap[K, V]{}
	for k, v := range seq {
		m.Add(k, v)
	}
	return m
}

// Get returns values for key
func (m UniqMultiMap[K, V]) Get(k K) set.Set[V] {
	return m[k]
}

// Add adds new key-value
func (m UniqMultiMap[K, V]) Add(k K, v V) {
	if s, ok := m[k]; ok {
		s.Add(v)
	} else {
		m[k] = set.New(v)
	}
}

// Remove removes k and all values
func (m UniqMultiMap[K, V]) Remove(k K) {
	delete(m, k)
}

func (m UniqMultiMap[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, vs := range m {
			for v := range vs {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

func (m UniqMultiMap[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		for k, _ := range m {
			if !yield(k) {
				break
			}
		}
	}
}

func (m UniqMultiMap[K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, vs := range m {
			for v := range vs {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Entries returns a Pair slice contains the key-value items in map.
func (m UniqMultiMap[K, V]) Entries() []pair.Pair[K, V] {
	var entries []pair.Pair[K, V]
	for k, vs := range m {
		for v := range vs {
			entries = append(entries, pair.Of(k, v))
		}
	}
	return entries
}
