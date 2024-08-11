package ints

import (
	"strconv"
	"unsafe"
)

// Types is constraint for all int types.
type Types interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// ToString convert int to Decimal string
func ToString[T Types](v T) string {
	var zero T
	if zero-1 < 0 {
		return strconv.FormatInt(int64(v), 10)
	} else {
		return strconv.FormatUint(uint64(v), 10)
	}

}

// Parse parses decimal string to int value, return an error if it is not a valid decimal number string.
func Parse[T Types](s string) (T, error) {
	var zero T
	if zero-1 < 0 {
		v, err := strconv.ParseInt(s, 10, int(unsafe.Sizeof(zero))*8)
		return T(v), err
	} else {
		v, err := strconv.ParseUint(s, 10, int(unsafe.Sizeof(zero))*8)
		return T(v), err
	}
}

// Parse parses decimal string to int value, return default value if it is not a valid decimal number string.
func SafeParse[T Types](s string, defaultValue T) T {
	v, err := Parse[T](s)
	if err == nil {
		return v
	}
	return defaultValue
}
