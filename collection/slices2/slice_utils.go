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

// Slice return a part of a slice. It treat start and end index likes python list slice and js array slice,
// the negative index counts back from the end of the slice. It is safe, never panics.
//
// For start:
//   - If start is negative, -len(s) < start < 0, start + len(s) is used.
//   - if start < -len(s), 0 is used.
//   - If start > len(s), len(s) is used.
//
// For end:
//   - If end is negative, -len(s) < end < 0, end + len(s) is used.
//   - if end < -len(s), 0 is used.
//   - If end > len(s), len(s) is used.
//
// If end implies a position before or at the position that start implies, an empty slice is returned.
func Slice[S ~[]T, T any](s S, start, end int) S {
	if start < 0 {
		if start > -len(s) {
			start += len(s)
		} else {
			start = 0
		}
	} else if start > len(s) {
		start = len(s)
	}

	if end < 0 {
		if end > -len(s) {
			end += len(s)
		} else {
			end = 0
		}
	} else if end > len(s) {
		end = len(s)
	}

	if start > end {
		start = end
	}

	return s[start:end]
}

// SliceToEnd likes Slice, but without the end index. The len(s) is used as the end index.
func SliceToEnd[S ~[]T, T any](s S, start int) S {
	return Slice(s, start, len(s))
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
