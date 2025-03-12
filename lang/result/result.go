package result

// Result is a generic type that represents a value or an error.
// It is used to handle operations that can succeed or fail without using exceptions.
type Result[T any] struct {
	value T
	err   error
}

// New creates a new Result instance with the given value and error.
// If the error is nil, the value is considered valid; otherwise, it indicates an error.
func Of[T any](value T, err error) Result[T] {
	return Result[T]{value: value, err: err}
}

// OfValue creates a new Result instance with the given value and no error.
func OfValue[T any](value T) Result[T] {
	return Result[T]{value: value, err: nil}
}

// OfError creates a new Result instance with an error and a zero value of type T.
func OfError[T any](err error) Result[T] {
	return Result[T]{value: *new(T), err: err}
}

// IsError checks if the Result contains an error.
func (r Result[T]) IsError() bool {
	return r.err != nil
}

// IsSuccess checks if the Result contains a valid value (no error).
func (r Result[T]) IsSuccess() bool {
	return r.err == nil
}

// Unwrap retrieves the value and error from the Result instance.
// If the Result contains an error, it returns the value and the error; otherwise, it returns the value and nil.
func (r Result[T]) Unwrap() (T, error) {
	if r.IsError() {
		return r.value, r.err
	}
	return r.value, nil
}

// Get retrieves the value from the Result instance if it is valid (no error).
// If the Result contains an error, it panics.
func (r Result[T]) Get() T {
	if r.IsError() {
		panic(r.err)
	}
	return r.value
}

// GetOrElse retrieves the value from the Result instance if it is valid (no error); otherwise, it returns the provided default value.
func (r Result[T]) GetOrElse(defaultValue T) T {
	if r.IsSuccess() {
		return r.value
	}
	return defaultValue
}

// Error retrieves the error from the Result instance if it contains an error.
// If the Result is valid (no error), it returns nil.
func (r Result[T]) Error() error {
	if r.IsError() {
		return r.err
	}
	return nil
}
