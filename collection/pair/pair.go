package pair

// Pair is a key-value pair.
type Pair[K, V any] struct {
	key   K
	value V
}

// Of creates a new Pair
func Of[K, V any](key K, value V) Pair[K, V] {
	return Pair[K, V]{key, value}
}

// Unpack unpacks pair to key, value
func (p Pair[K, V]) Unpack() (K, V) {
	return p.key, p.value
}

// Key return key of pair
func (p Pair[K, V]) Key() K {
	return p.key
}

// Get value return value of pair
func (p Pair[K, V]) Value() V {
	return p.value
}

// First returns the first value of pair
func (p Pair[K, V]) First() K {
	return p.key
}

// Second returns the second value of pair
func (p Pair[K, V]) Second() V {
	return p.value
}
