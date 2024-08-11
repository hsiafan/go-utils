package iters

import (
	"iter"

	"github.com/hsiafan/go-utils/collection/pair"
)

// Map
func Map[T any, R any](seq iter.Seq[T], convert func(v T) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range seq {
			if !yield(convert(v)) {
				break
			}
		}
	}
}

// MapToSeq2
func MapToSeq2[T any, K, V any](seq iter.Seq[T], convert func(v T) (K, V)) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for v := range seq {
			if !yield(convert(v)) {
				break
			}
		}
	}
}

// MapToSeq1
func MapToSeq1[K, V any, T any](seq iter.Seq2[K, V], convert func(K, V) T) iter.Seq[T] {
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

// Filter
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

// Indexed
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
