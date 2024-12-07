package slices2

import (
	"testing"

	"github.com/hsiafan/go-utils/collection/pair"
	"github.com/stretchr/testify/assert"
)

func TestSortBy(t *testing.T) {
	pairs := []pair.Pair[string, int]{
		pair.Of("1", 1),
		pair.Of("3", 3),
		pair.Of("2", 2),
		pair.Of("2", 1),
	}
	SortBy(pairs, pair.Pair[string, int].Value)
	assert.Equal(t, pair.Of("2", 2), pairs[2])
	assert.Equal(t, pair.Of("3", 3), pairs[3])
}

func TestSortStableBy(t *testing.T) {
	pairs := []pair.Pair[string, int]{
		pair.Of("1", 1),
		pair.Of("3", 3),
		pair.Of("2", 2),
		pair.Of("2", 1),
	}
	SortStableBy(pairs, pair.Pair[string, int].Value)
	assert.Equal(t, []pair.Pair[string, int]{
		pair.Of("1", 1),
		pair.Of("2", 1),
		pair.Of("2", 2),
		pair.Of("3", 3),
	}, pairs)
}

func TestSlice(t *testing.T) {
	s := []int{1, 2, 3}
	assert.Equal(t, []int{1}, Slice(s, 0, 1))
	assert.Equal(t, []int{1, 2}, Slice(s, 0, -1))
	assert.Equal(t, []int{2}, Slice(s, -2, -1))
	assert.Equal(t, []int{1}, Slice(s, -5, 1))
	assert.Equal(t, []int{1}, Slice(s, -5, 1))
	assert.Equal(t, []int{2, 3}, Slice(s, 1, 100))
	assert.Equal(t, []int{}, Slice(s, 5, 2))
}

func TestLastN(t *testing.T) {
	s := []int{1, 2, 3}
	assert.Equal(t, []int{3}, LastN(s, 1))
	assert.Equal(t, []int{1, 2, 3}, LastN(s, 4))
	assert.Equal(t, []int{}, LastN(s, -1))
}

func TestFirstN(t *testing.T) {
	s := []int{1, 2, 3}
	assert.Equal(t, []int{1}, FirstN(s, 1))
	assert.Equal(t, []int{1, 2, 3}, FirstN(s, 4))
	assert.Equal(t, []int{}, FirstN(s, -1))
}
