package maps2

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedMap_PutAndRemove(t *testing.T) {
	m := NewLinkedMap[string, int]()
	m.Put("1", 1)
	value, ok := m.Get("1")
	assert.True(t, ok)
	assert.Equal(t, 1, value)

	m.Put("2", 2)
	m.Put("3", 3)
	m.Put("4", 4)
	m.Remove("2")

	assert.Equal(t, 3, m.Size())
	assert.False(t, m.Contains("2"))
	assert.Equal(t, []string{"1", "3", "4"}, slices.Collect(m.Keys()))
	assert.Equal(t, []int{1, 3, 4}, slices.Collect(m.Values()))
}

func TestLinkedMap_Copy(t *testing.T) {
	m := NewLinkedMap[string, int]()
	m.Put("1", 1)
	m.Put("2", 2)
	m.Put("3", 3)

	nm := m.Copy()
	assert.Equal(t, 3, nm.Size())
	assert.Equal(t, []string{"1", "2", "3"}, slices.Collect(nm.Keys()))

}

func TestLinkedMap_All(t *testing.T) {
	m := NewLinkedMap[string, int]()
	m.Put("1", 1)
	m.Put("2", 2)
	m.Put("3", 3)

	assert.Equal(t, []string{"1", "2", "3"}, slices.Collect(m.Keys()))
	assert.Equal(t, []int{1, 2, 3}, slices.Collect(m.Values()))
}

func TestLinkedMap_Clear(t *testing.T) {
	m := NewLinkedMap[string, int]()
	m.Put("1", 1)
	m.Put("2", 2)

	m.Clear()
	assert.Equal(t, 0, m.Size())
	assert.False(t, m.Contains("1"))

	m.Put("1", 1)
	assert.Equal(t, 1, m.Size())
}
