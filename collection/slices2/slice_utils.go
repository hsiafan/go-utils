package slices2

import (
	"cmp"
	"iter"
	"slices"
)

// CollectWithError collects values in seq to slices. If err occurred, return nil slice and the err.
func CollectWithError[T any](seq iter.Seq2[T, error]) ([]T, error) {
	var s []T
	for v, err := range seq {
		if err != nil {
			return nil, err
		}
		s = append(s, v)
	}
	return s, nil
}

// Convert return a new slice with values applied func f on original slice.
func Convert[S ~[]T, T any, R any](s S, f func(v T) R) []R {
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

// FirstN returns the first n elements of the slice.
// If n is greater than the length of the slice, the original slice is returned.
func FirstN[T any](s []T, n int) []T {
	if n <= 0 {
		return nil
	}
	if n >= len(s) {
		return s
	}
	return s[:n]
}

// LastN returns the last n elements of the slice.
// If n is greater than the length of the slice, the original slice is returned.
func LastN[T any](s []T, n int) []T {
	if n <= 0 {
		return nil
	}
	if n >= len(s) {
		return s
	}
	return s[len(s)-n:]
}

// First returns the first element of the slice.
// If the slice is empty, the zero value of the element type is returned.
func First[T any](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	return s[0]
}

// FirstOrElse returns the first element of the slice.
// If the slice is empty, the defaultValue is returned.
func FirstOrElse[T any](s []T, defaultValue T) T {
	if len(s) == 0 {
		return defaultValue
	}
	return s[0]
}

// Last returns the last element of the slice.
// If the slice is empty, the zero value of the element type is returned.
func Last[T any](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	return s[len(s)-1]
}

// LastOrElse returns the last element of the slice.
// If the slice is empty, the defaultValue is returned.
func LastOrElse[T any](s []T, defaultValue T) T {
	if len(s) == 0 {
		return defaultValue
	}
	return s[len(s)-1]
}
