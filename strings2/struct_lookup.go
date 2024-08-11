package strings2

import (
	"reflect"
	"unsafe"
)

type structLookup struct {
	cache map[reflect.Type]structInfo
}

type structInfo struct {
	fields map[string]int
	embeds []int
}

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

func newStructLookup() *structLookup {
	return &structLookup{
		cache: map[reflect.Type]structInfo{},
	}
}

func (s *structLookup) lookup(v reflect.Value, name string) (any, bool) {
	t := v.Type()
	si := s.loadStructFields(t)
	if idx, ok := si.fields[name]; ok {
		fv := v.Field(idx)
		if !fv.CanInterface() {
			p := (*reflectValue)(unsafe.Pointer(&fv))
			p.flag = p.flag & ^flagRO
		}
		return fv.Interface(), true
	}

	for _, index := range si.embeds {
		if r, ok := s.lookup(v.Field(index), name); ok {
			return r, ok
		}
	}
	return nil, false
}

func (s *structLookup) loadStructFields(t reflect.Type) structInfo {
	v, ok := s.cache[t]
	if ok {
		return v
	}
	var embeds []int
	var fields = map[string]int{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.Anonymous {
			fields[f.Name] = i
		} else if f.Type.Kind() == reflect.Struct {
			embeds = append(embeds, i)
		}
	}
	si := structInfo{
		fields: fields,
		embeds: embeds,
	}
	s.cache[t] = si
	return si
}
