package bytes

import "unsafe"

// AsString converts a byte slice to a string without memory allocation.
func AsString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
