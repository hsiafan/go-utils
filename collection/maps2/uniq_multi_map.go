package maps2

import (
	"github.com/hsiafan/go-utils/collection/pair"
	"github.com/hsiafan/go-utils/collection/set"
)

// UniqMultiMap can hold multi value of one key, the same value only store once
type UniqMultiMap[K comparable, V comparable] map[K]set.Set[V]

// Add adds new key-value
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
