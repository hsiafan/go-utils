package optional

import "errors"

// Optional is a generic type that represents an optional value.
// It can be used to indicate that a value may or may not be present.
// Param T is the type of the value that may be present.
type Optional[T any] struct {
	value   T
	present bool
}

// Of creates a new Optional instance with the given value and presence flag.
func Of[T any](v T, present bool) Optional[T] {
	return Optional[T]{value: v, present: present}
}

// OfValue creates a new Optional instance with the given value and sets the presence flag to true.
func OfValue[T any](v T) Optional[T] {
	return Optional[T]{value: v, present: true}
}

// Empty creates a new Optional instance with no value and sets the presence flag to false.
func Empty[T any]() Optional[T] {
	return Optional[T]{present: false}
}

// IsPresent checks if the Optional instance contains a value.
func (v Optional[T]) IsPresent() bool {
	return v.present
}

// IsEmpty checks if the Optional instance does not contain a value.
func (v Optional[T]) IsEmpty() bool {
	return !v.present
}

var ErrorValueNotPresent = errors.New("value not present")

// Unwrap returns the value and a boolean indicating whether the value is present.
// If the value is not present, it returns the zero value of T and false.
func (v Optional[T]) Unwrap() (T, bool) {
	if !v.present {
		return v.value, false
	}
	return v.value, true
}

// Get retrieves the value from the Optional instance if it is present.
func (v Optional[T]) Get() T {
	if !v.present {
		panic(ErrorValueNotPresent)
	}
	return v.value
}

// GetOrElse retrieves the value from the Optional instance if it is present; otherwise, it returns the provided default value.
func (v Optional[T]) GetOrElse(defaultValue T) T {
	if v.present {
		return v.value
	}
	return defaultValue
}

// GetOrZero retrieves the value from the Optional instance if it is present; otherwise, it returns the zero value of T.
func (v Optional[T]) GetOrZero() T {
	if v.present {
		return v.value
	}
	var zero T
	return zero
}
