package maps2

import (
	"github.com/hsiafan/go-utils/collection/pair"
)

// Contains reports if given key exists in the map.
func Contains[M ~map[K]V, K comparable, V any](m M, k K) bool {
	_, ok := m[k]
	return ok
}

// GetOrElse returns the value for given key; return default value if key is not exists.
func GetOrElse[M ~map[K]V, K comparable, V any](m M, k K, defaultV V) V {
	if v, ok := m[k]; ok {
		return v
	}
	return defaultV
}

// AddMap adds all key-values in map2 into map. If the key-value in m2 already exists in map, it will be override.
func AddMap[M ~map[K]V, K comparable, V any](m M, another M) {
	for k, v := range another {
		m[k] = v
	}
}

// AddIfAbsent add key-value to map if key not exists.
// It returns the value for this key.
func AddIfAbsent[M ~map[K]V, K comparable, V any](m M, k K, defaultV V) V {
	if v, ok := m[k]; ok {
		return v
	}
	m[k] = defaultV
	return defaultV
}

// ComputeIfAbsent add key-value to map if key not exists, the value is computed by compute func.
// It returns the value for this key.
func ComputeIfAbsent[M ~map[K]V, K comparable, V any](m M, k K, compute func(K) V) V {
	if v, ok := m[k]; ok {
		return v
	}
	v := compute(k)
	m[k] = v
	return v
}

// Map creates a new map, with new values compute by compute func.
func Map[M ~map[K]V, K comparable, V any, U any](m M, k K, compute func(K, V) U) map[K]U {
	if m == nil {
		return nil
	}
	rm := make(map[K]U, len(m))
	for k, v := range m {
		rm[k] = compute(k, v)
	}
	return rm
}

// Filter creates a new map with new values accept by predicate func.
func Filter[M ~map[K]V, K comparable, V any](m M, k K, predicate func(K, V) bool) map[K]V {
	if m == nil {
		return nil
	}
	rm := make(map[K]V)
	for k, v := range m {
		if predicate(k, v) {
			rm[k] = v
		}
	}
	return rm
}

// Merge merges all key-value pairs in given maps into one new Map.
// If key-value exists in multi map, the last key-value will be used.
func Merge[M ~map[K]V, K comparable, V any](maps ...M) M {
	var size = 0
	for _, m := range maps {
		size += len(m)
	}
	nm := make(M, size)
	for _, m := range maps {
		for k, v := range m {
			nm[k] = v
		}
	}
	return nm
}

// Entries returns a Pair slice contains the key-value items in map.
func Entries[M ~map[K]V, K comparable, V any](m M) []pair.Pair[K, V] {
	ps := make([]pair.Pair[K, V], len(m))
	for k, v := range m {
		ps = append(ps, pair.Of(k, v))
	}
	return ps
}
