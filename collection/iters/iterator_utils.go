package iters

import (
	"iter"

	"github.com/hsiafan/go-utils/collection/pair"
)

// Map maps to a new Seq with values applied func convert
func Map[T any, R any](seq iter.Seq[T], convert func(v T) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range seq {
			if !yield(convert(v)) {
				break
			}
		}
	}
}

// MapToSeq2 maps Seq to a Seq2 with values applied func convert
func MapToSeq2[T any, K, V any](seq iter.Seq[T], convert func(v T) (K, V)) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for v := range seq {
			if !yield(convert(v)) {
				break
			}
		}
	}
}

// MapToSeq  maps Seq2 to a Seq with values applied func convert
func MapToSeq[K, V any, T any](seq iter.Seq2[K, V], convert func(K, V) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for k, v := range seq {
			if !yield(convert(k, v)) {
				break
			}
		}
	}
}

// MapToPairSeq convert iter.Seq2 to iter.Seq of pair.Pair
func MapToPairSeq[K, V any](seq iter.Seq2[K, V]) iter.Seq[pair.Pair[K, V]] {
	return func(yield func(pair.Pair[K, V]) bool) {
		for k, v := range seq {
			if !yield(pair.Of(k, v)) {
				break
			}
		}
	}
}

// Filter returns a new Seq contains the values accepted by predicate
func Filter[T any](seq iter.Seq[T], predicate func(v T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if predicate(v) {
				if !yield(v) {
					break
				}
			}
		}
	}
}

// Filter2 returns a new Seq2 contains the values accepted by predicate
func Filter2[K, V any](seq iter.Seq2[K, V], predicate func(k K, v V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if predicate(k, v) {
				if !yield(k, v) {
					break
				}
			}
		}
	}
}

// Drop returns a sequence containing all elements except first n elements.
func Drop[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		var i int
		for e := range seq {
			if i >= n {
				if !yield(e) {
					break
				}
			}
			i++
		}
	}
}

// Drop2 returns a sequence containing all elements except first n elements.
func Drop2[K, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		var i int
		for k, v := range seq {
			if i >= n {
				if !yield(k, v) {
					break
				}
			}
			i++
		}
	}
}

// Take returns a sequence containing first n elements.
func Take[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		var i int
		for e := range seq {
			if i < n {
				if !yield(e) {
					break
				}
			} else {
				break
			}
			i++
		}
	}
}

// Take2 returns a sequence containing first n entries.
func Take2[K, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		var i int
		for k, v := range seq {
			if i < n {
				if !yield(k, v) {
					break
				}
			} else {
				break
			}
			i++
		}
	}
}

// Indexed returns a new Seq2 with index
func Indexed[T any](seq iter.Seq[T]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		var i = 0
		for v := range seq {
			if !yield(i, v) {
				break
			}
			i++
		}
	}
}
