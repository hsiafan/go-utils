package floats

import (
	"strconv"
	"unsafe"
)

// Types is constraint for all float types.
type Types interface {
	~float32 | ~float64
}

// ToString format floats to string, no using exponent.
// The precision prec controls the number of digits after the decimal point.
func ToString[T Types](v T, prec int) string {
	return strconv.FormatFloat(float64(v), 'f', prec, int(unsafe.Sizeof(v))*8)
}

// Parse parses a float value string. Return an error if is not a valid float string.
func Parse[T Types](s string) (T, error) {
	var zero T
	v, err := strconv.ParseFloat(s, int(unsafe.Sizeof(zero))*8)
	return T(v), err
}

// SafeParse parses a float value string. Return the default value if is not a valid float string.
func SafeParse[T Types](s string, defaultValue T) T {
	v, err := Parse[T](s)
	if err == nil {
		return v
	}
	return defaultValue
}
