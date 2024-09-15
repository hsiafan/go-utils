package reflects

import (
	"errors"
	"reflect"
	"sync"
	"unsafe"
)

// StructLookup used to deal reflect operations for struct fields and methods.
// The lookup instance should be reused to improve the performance.
type StructLookup struct {
	cache sync.Map //map[reflect.Type]structInfo
}

// structInfo contains struct meta cache to accelerate the reflect operations.
type structInfo struct {
	fields map[string]int // The field name to index map
	// The embed fields names to index map.
	// The names is retrieve recursively from all embed type fields;
	// The values is the indexes for each level' struct field, from the inner to outer.
	embedFields map[string][]int
}

var defaultStructLookup = sync.OnceValue(func() *StructLookup {
	return NewStructLookup()
})

// DefaultStructLookup returns the default StructLookup instance.
func DefaultStructLookup() *StructLookup {
	return defaultStructLookup()
}

// NewStructLookup returns a new created StructLookup instance.
func NewStructLookup() *StructLookup {
	return &StructLookup{
		cache: sync.Map{},
	}
}

// Field get the field value for the giving struct value and field name.
// It will look into embed struct fields if no match Field is found.
//
// param v: a reflect contains a struct value, or contains a pointer to struct value.
func (s *StructLookup) Field(v reflect.Value, name string) (reflect.Value, bool) {
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		panic(errors.New("not a struct value"))
	}

	t := v.Type()
	si := s.loadStructFields(t)
	if idx, ok := si.fields[name]; ok {
		fv := v.Field(idx)
		if !fv.CanInterface() {
			// try to read unexported fields
			p := (*reflectValue)(unsafe.Pointer(&fv))
			p.flag = p.flag & ^flagRO
		}
		return fv, true
	}

	if indexes, ok := si.embedFields[name]; ok {
		fv := v
		for i := len(indexes) - 1; i >= 0; i-- {
			fv = fv.Field(indexes[i])
			if fv.Kind() == reflect.Pointer {
				fv = fv.Elem()
			}
		}
		if !fv.CanInterface() {
			// try to read unexported fields
			p := (*reflectValue)(unsafe.Pointer(&fv))
			p.flag = p.flag & ^flagRO
		}
		return fv, true
	}

	var zero reflect.Value
	return zero, false
}

func (s *StructLookup) loadStructFields(t reflect.Type) structInfo {
	v, ok := s.cache.Load(t)
	if ok {
		return v.(structInfo)
	}

	var embeds = map[string][]int{}
	var fields = map[string]int{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.Anonymous {
			fields[f.Name] = i
		}
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		ft := f.Type
		if f.Anonymous {
			if ft.Kind() == reflect.Pointer {
				ft = ft.Elem()
			}
			if ft.Kind() == reflect.Struct {
				si := s.loadStructFields(ft)
				for name, idx := range si.fields {
					if _, ok := fields[name]; !ok {
						embeds[name] = []int{idx, i}
					}
				}
				for name, indexes := range si.embedFields {
					if _, ok := fields[name]; !ok {
						embeds[name] = append(indexes, i)
					}
				}
			}
		}
	}
	si := structInfo{
		fields:      fields,
		embedFields: embeds,
	}

	s.cache.Store(t, si)
	return si
}
