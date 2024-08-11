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
