package jsons

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalString(t *testing.T) {
	type myStruct struct {
		Name string `json:"name"`
	}
	var jsonStr = `{"name":"test"}`

	p, err := UnmarshalString[*myStruct](jsonStr)
	assert.NoError(t, err)
	assert.Equal(t, "test", p.Name)

	v, err := UnmarshalString[myStruct](jsonStr)
	assert.NoError(t, err)
	assert.Equal(t, "test", v.Name)

	i, err := UnmarshalString[any](jsonStr)
	assert.NoError(t, err)
	assert.Equal(t, map[string]any{"name": "test"}, i)
}
