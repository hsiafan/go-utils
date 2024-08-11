package strings2

import (
	"strings"
	"unicode"
	"unicode/utf8"
	"unsafe"

	"github.com/hsiafan/go-utils/bytes"
)

// PrependIfMissing adds a prefix to string, if string do not has this prefix.
func PrependIfMissing(s string, prefix string) string {
	if !strings.HasPrefix(s, prefix) {
		return prefix + s
	}
	return s
}

// AppendIfMissing adds a suffix to string, if string do not has this prefix.
func AppendIfMissing(s string, suffix string) string {
	if !strings.HasSuffix(s, suffix) {
		return s + suffix
	}
	return s
}

// PadLeft adds pad to string at left, if string len is little than width.
func PadLeft(s string, width int, c byte) string {
	if len(s) >= width {
		return s
	}
	bs := make([]byte, width)
	for i := 0; i < width-len(s); i++ {
		bs[i] = c
	}
	copy(bs[width-len(s):], s)
	return bytes.AsString(bs)
}

// PadRight adds pad to string at right, if string len is little than width.
func PadRight(s string, width int, c byte) string {
	if len(s) >= width {
		return s
	}
	bs := make([]byte, width)
	copy(bs[:len(s)], s)
	for i := len(s); i < width; i++ {
		bs[i] = c
	}
	return bytes.AsString(bs)
}

// CompareLower compares two strings without considering their case, all upper-case char are compared as small-case.
// It returns:
// - a negative number if s1 < s2,
// - zero if s1 == s2,
// - a positive number if s1 > s2
func CompareLower(s, t string) int {
	// ASCII fast path
	i := 0
	for ; i < len(s) && i < len(t); i++ {
		sr := s[i]
		tr := t[i]
		if sr|tr >= utf8.RuneSelf {
			goto hasUnicode
		}

		if sr == tr {
			continue
		}

		if 'A' <= sr && sr <= 'Z' {
			sr += 'a' - 'A'
		}

		if 'A' <= tr && tr <= 'Z' {
			tr += 'a' - 'A'
		}
		if sr == tr {
			continue
		}
		return int(sr) - int(tr)
	}
	// Check if we've exhausted both strings.
	return len(s) - len(t)

hasUnicode:
	s = s[i:]
	t = t[i:]
	for _, sr := range s {
		// If t is exhausted the strings are not equal.
		if len(t) == 0 {
			return 1
		}

		// Extract first rune from second string.
		var tr rune
		if t[0] < utf8.RuneSelf {
			tr, t = rune(t[0]), t[1:]
		} else {
			r, size := utf8.DecodeRuneInString(t)
			tr, t = r, t[size:]
		}

		if sr == tr {
			continue
		}

		sr = unicode.ToLower(sr)
		tr = unicode.ToLower(tr)
		if sr == tr {
			continue
		}
		return int(sr) - int(tr)
	}

	// First string is empty, so check if the second one.
	return -len(t)
}

// AsBytes converts a string to a byte slice without memory allocation.
func AsBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
