package linkedset

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedSet(t *testing.T) {
	set := New(1, 2, 3)
	assert.Equal(t, 3, set.Size())
	assert.Equal(t, []int{1, 2, 3}, slices.Collect(set.All()))
	set.Remove(2)
	assert.Equal(t, []int{1, 3}, slices.Collect(set.All()))
}

func TestLinkedSet_Add(t *testing.T) {
	set := New[int]()
	set.Add(1)
	assert.True(t, set.Contains(1))
}

func TestLinkedSet_AddAll(t *testing.T) {
	set := New[int]()
	set.AddAll(1, 2, 3)
	assert.True(t, set.Contains(1))
	assert.True(t, set.Contains(2))
	assert.True(t, set.Contains(3))
	assert.Equal(t, 3, set.Size())
}

func TestLinkedSet_AddSet(t *testing.T) {
	set1 := New(1, 2)
	set2 := New(3, 4)
	set1.AddSet(set2)
	assert.True(t, set1.Contains(3))
	assert.True(t, set1.Contains(4))
	assert.Equal(t, 4, set1.Size())
}

func TestLinkedSet_Remove(t *testing.T) {
	set := New(1, 2, 3)
	set.Remove(2)
	assert.False(t, set.Contains(2))
	assert.Equal(t, 2, set.Size())
}

func TestLinkedSet_RemoveAll(t *testing.T) {
	set := New(1, 2, 3, 4)
	set.RemoveAll(2, 3)
	assert.False(t, set.Contains(2))
	assert.False(t, set.Contains(3))
	assert.Equal(t, 2, set.Size())
}

func TestLinkedSet_Clear(t *testing.T) {
	set := New(1, 2, 3)
	set.Clear()
	assert.Equal(t, 0, set.Size())
	assert.False(t, set.Contains(1))

	set.Add(1)
	assert.True(t, set.Contains(1))
}

func TestLinkedSet_Copy(t *testing.T) {
	set := New(1, 2, 3)
	copied := set.Copy()
	assert.Equal(t, 3, copied.Size())
	assert.True(t, copied.Contains(1))
	assert.True(t, copied.Contains(2))
	assert.True(t, copied.Contains(3))
}
