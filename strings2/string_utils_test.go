package strings2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPadLeft(t *testing.T) {
	assert.Equal(t, "123", PadLeft("123", 2, '0'))
	assert.Equal(t, "123", PadLeft("123", 3, '0'))
	assert.Equal(t, "00123", PadLeft("123", 5, '0'))
}

func TestPadRight(t *testing.T) {
	assert.Equal(t, "123", PadRight("123", 2, '0'))
	assert.Equal(t, "123", PadRight("123", 3, '0'))
	assert.Equal(t, "12300", PadRight("123", 5, '0'))
}

func TestCompareLower(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected int
	}{
		{"", "", 0},
		{"Hello", "hello", 0},
		{"Hello", "Helloo", -1},
		{"Hello", "Hell", 1},
		{"apple", "Banana", -1},
		{"Orange", "orange", 0},
		{"grape", "Grape", 0},
		{"kiwi", "Kiwi", 0},
		{"melon", "Melon", 0},
		{"melon", "mElon", 0},
		{"melon", "MelonS", -1},
		{"melonS", "melon", 1},
		{"你好Hello", "你好hello", 0},
		{"你好Hello", "你好helloo", -1},
	}

	for _, test := range tests {
		result := CompareLower(test.s1, test.s2)
		if result < 0 && test.expected >= 0 ||
			result > 0 && test.expected <= 0 ||
			result == 0 && test.expected != 0 {
			t.Errorf("CaseInsensitiveCompare(%q, %q) = %d; expected %d", test.s1, test.s2, result, test.expected)
		}
	}
}
