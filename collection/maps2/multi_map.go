package maps2

import "github.com/hsiafan/go-utils/collection/pair"

// MultiMap can hold multi value of one key
type MultiMap[K comparable, V any] map[K][]V

// Add adds new key-value
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
