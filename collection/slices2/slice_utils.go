package slices2

import (
	"cmp"
	"slices"
)

// Map return a new slice with values applied func f on original slice.
func Map[S ~[]T, T any, R any](s S, f func(v T) R) []R {
	if s == nil {
		return nil
	}
	r := make([]R, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Filter return a new slice with values accepted by func f
func Filter[S ~[]T, T any](s S, f func(v T) bool) S {
	if s == nil {
		return nil
	}
	var r S
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

// SortBy sorts the slice in ascending order. The element is compared by value apply extract function on e.
func SortBy[S ~[]T, T any, O cmp.Ordered](s S, extract func(e T) O) {
	slices.SortFunc(s, func(e1, e2 T) int {
		v1 := extract(e1)
		v2 := extract(e2)
		if v1 < v2 {
			return -1
		} else if v1 == v2 {
			return 0
		} else {
			return 1
		}
	})
}

// SortStableBy sorts the slice in ascending order, while keeping the original order of equal elements.
// The element is compared by value apply extract function on e.
func SortStableBy[S ~[]T, T any, O cmp.Ordered](s S, extract func(e T) O) {
	slices.SortStableFunc(s, func(e1, e2 T) int {
		v1 := extract(e1)
		v2 := extract(e2)
		if v1 < v2 {
			return -1
		} else if v1 == v2 {
			return 0
		} else {
			return 1
		}
	})
}

// Copy copies and return the slice
func Copy[S ~[]T, T any](s S) S {
	if s == nil {
		return nil
	}
	ns := make(S, len(s))
	copy(ns, s)
	return ns
}
