package maps2

import (
	"iter"

	"github.com/hsiafan/go-utils/collection/pair"
)

// MultiMap can hold multi value of one key
type MultiMap[K comparable, V any] map[K][]V

func New[K comparable, V any]() MultiMap[K, V] {
	return MultiMap[K, V]{}
}

func Collect[K comparable, V any](seq iter.Seq2[K, V]) MultiMap[K, V] {
	m := MultiMap[K, V]{}
	for k, v := range seq {
		m[k] = append(m[k], v)
	}
	return m
}

// Get gets values for key
func (m MultiMap[K, V]) Get(k K) []V {
	return m[k]
}

// Add adds new key-value
func (m MultiMap[K, V]) Add(k K, v V) {
	m[k] = append(m[k], v)
}

// Remove removes k and all values
func (m MultiMap[K, V]) Remove(k K) {
	delete(m, k)
}

func (m MultiMap[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, vs := range m {
			for _, v := range vs {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

func (m MultiMap[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		for k, _ := range m {
			if !yield(k) {
				break
			}
		}
	}
}

func (m MultiMap[K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, vs := range m {
			for _, v := range vs {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Entries returns a Pair slice contains the key-value items in map.
func (m MultiMap[K, V]) Entries() []pair.Pair[K, V] {
	var entries []pair.Pair[K, V]
	for k, vs := range m {
		for _, v := range vs {
			entries = append(entries, pair.Of(k, v))
		}
	}
	return entries
}
