package ioutil

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLines(t *testing.T) {
	var str = "1\n2\n3\n"

	{
		var r = strings.NewReader(str)
		var lines []string
		for line, err := range Lines(r) {
			assert.NoError(t, err)
			lines = append(lines, line)
			break
		}
		assert.Equal(t, []string{"1"}, lines)
	}

	{
		var r = strings.NewReader(str)
		var lines []string
		for line, err := range Lines(r) {
			assert.NoError(t, err)
			lines = append(lines, line)
		}
		assert.Equal(t, []string{"1", "2", "3"}, lines)
	}

}
