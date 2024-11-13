package maps2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	var m = map[int]string{1: "1", 2: "2", 3: "3"}
	var nm = Convert(m, func(k int, v string) (int, string) {
		return k * 2, v + v
	})
	assert.Equal(t, map[int]string{2: "11", 4: "22", 6: "33"}, nm)

	var nm2 = Convert(m, func(k int, v string) (string, int) {
		return v, k
	})
	assert.Equal(t, map[string]int{"1": 1, "2": 2, "3": 3}, nm2)
}
