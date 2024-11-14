package maps2

import "iter"

// node is one linked list node
type node[K comparable, V any] struct {
	k    K
	v    V
	prev *node[K, V]
	next *node[K, V]
}

// LinkedMap is a map with nodes maintained by a linked list, it can keep the order of keys.
type LinkedMap[K comparable, V any] struct {
	m    map[K]*node[K, V]
	head *node[K, V]
	tail *node[K, V]
}

// NewLinkedMap creates a new LinkedMap.
func NewLinkedMap[K comparable, V any]() *LinkedMap[K, V] {
	return &LinkedMap[K, V]{m: make(map[K]*node[K, V])}
}

// Contains returns true if key exists.
func (m *LinkedMap[K, V]) Contains(k K) bool {
	_, ok := m.m[k]
	return ok
}

// Get returns value for key.
func (m *LinkedMap[K, V]) Get(k K) (v V, ok bool) {
	if n, ok := m.m[k]; ok {
		return n.v, true
	}
	return
}

// GetOrZero returns value for key; return zero value if key is not exists.
func (m *LinkedMap[K, V]) GetOrZero(k K) V {
	if n, ok := m.m[k]; ok {
		return n.v
	}
	var zero V
	return zero
}

// Put adds or sets value for key.
func (m *LinkedMap[K, V]) Put(k K, v V) {
	if n, ok := m.m[k]; ok {
		n.v = v
		return
	}
	n := &node[K, V]{k: k, v: v}
	m.m[k] = n
	m.insertNode(n)
}

// PutMap adds/sets all key-values in another map.
func (m *LinkedMap[K, V]) PutMap(another LinkedMap[K, V]) {
	for n := another.head; n != nil; n = n.next {
		m.Put(n.k, n.v)
	}
}

func (m *LinkedMap[K, V]) insertNode(n *node[K, V]) {
	if m.head == nil {
		m.head = n
		m.tail = n
	} else {
		m.tail.next = n
		n.prev = m.tail
		m.tail = n
	}
}

// Remove removes key.
func (m *LinkedMap[K, V]) Remove(k K) {
	if n, ok := m.m[k]; ok {
		m.removeNode(n)
		delete(m.m, k)
	}
}

// RemoveAll removes all keys.
func (m *LinkedMap[K, V]) RemoveAll(keys ...K) {
	for _, k := range keys {
		m.Remove(k)
	}
}

func (m *LinkedMap[K, V]) removeNode(n *node[K, V]) {
	if n.prev == nil {
		m.head = n.next
	} else {
		n.prev.next = n.next
	}
	if n.next == nil {
		m.tail = n.prev
	} else {
		n.next.prev = n.prev
	}
}

// Copy returns a new LinkedMap with same key-values.
func (m *LinkedMap[K, V]) Copy() *LinkedMap[K, V] {
	nm := NewLinkedMap[K, V]()
	for n := m.head; n != nil; n = n.next {
		nm.Put(n.k, n.v)
	}
	return nm
}

// All returns all key-value pairs as a sequence.
// The order of key-value pairs is the same as they were added to map.
func (m *LinkedMap[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for n := m.head; n != nil; n = n.next {
			if !yield(n.k, n.v) {
				break
			}
		}
	}
}

// Keys returns all keys as a sequence.
// The order of keys is the same as they were added to map.
func (m *LinkedMap[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		for n := m.head; n != nil; n = n.next {
			if !yield(n.k) {
				break
			}
		}
	}
}

// Values returns all values as a sequence.
// The order of values is the same as key-values pairs were added to map.
func (m *LinkedMap[K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for n := m.head; n != nil; n = n.next {
			if !yield(n.v) {
				break
			}
		}
	}
}

// Size returns the size of the map.
func (m *LinkedMap[K, V]) Size() int {
	return len(m.m)
}

// Clear clears the map.
func (m *LinkedMap[K, V]) Clear() {
	clear(m.m)
	m.head = nil
	m.tail = nil
}
