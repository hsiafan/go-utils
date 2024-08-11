package ints

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	assert.Equal(t, "1", ToString(1))
	assert.Equal(t, "-1", ToString(-1))

	assert.Equal(t, "1", ToString(uint(1)))
	assert.Equal(t, "255", ToString(uint8(0xFF)))
}

func TestParse(t *testing.T) {
	v, err := Parse[int]("1")
	assert.NoError(t, err)
	assert.Equal(t, 1, v)

	u, err := Parse[uint]("1")
	assert.NoError(t, err)
	assert.Equal(t, uint(1), u)

	_, err = Parse[uint]("-1")
	assert.Error(t, err)
}

func TestSafeParse(t *testing.T) {
	assert.Equal(t, 1, SafeParse[int]("1", 0))
	assert.Equal(t, 1, SafeParse[int]("+1", 0))
	assert.Equal(t, 0, SafeParse[int]("a1", 0))
	assert.Equal(t, 0, SafeParse[int]("", 0))
}
