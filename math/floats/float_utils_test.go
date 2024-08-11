package floats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	assert.Equal(t, "3.0", ToString(3.0, 1))
	assert.Equal(t, "3.14", ToString(3.1415926, 2))
}

func TestParse(t *testing.T) {

}

func TestSafeParse(t *testing.T) {
}
