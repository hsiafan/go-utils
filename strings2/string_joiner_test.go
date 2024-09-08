package strings2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoiner_Join(t *testing.T) {
	j := Joiner{Prefix: "[", Suffix: "]", Separator: ","}
	assert.Equal(t, "[]", j.Join([]string{}))
	assert.Equal(t, "[1]", j.Join([]string{1}))
	assert.Equal(t, "[1,2]", j.Join([]string{"1", "2"}))

}

func TestJoiner_Split(t *testing.T) {
	j := Joiner{Prefix: "[", Suffix: "]", Separator: ","}
	assert.Equal(t, []string{"1", "2"}, j.Split("[1,2]"))
	assert.Equal(t, []string{"1"}, j.Split("[1]"))
	assert.Equal(t, []string{""}, j.Split("[]"))
	assert.Equal(t, []string{"1", "2"}, j.Split("1,2"))

}
