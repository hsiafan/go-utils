package reflects

import "unsafe"

type flag uintptr

const (
	flagStickyRO flag = 1 << 5 // unexported not embedded field
	flagEmbedRO  flag = 1 << 6 // unexported embedded field
	flagRO       flag = flagStickyRO | flagEmbedRO
)

type reflectValue struct {
	typ_ unsafe.Pointer
	ptr  unsafe.Pointer
	flag
}
