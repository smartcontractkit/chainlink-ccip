package utils

// Optional represents a value that may or may not be explicitly set. This
// is useful in scenarios where a changeset only needs to partially update
// a subset of fields in a large configuration struct.
type Optional[T any] struct {
	// Valid indicates whether the provided value should be used. If this is
	// false (the default), then the provided value will be ignored.
	Valid bool `json:"valid" yaml:"valid"`

	// This only has an effect when `Valid` is set to true. If this is the
	// case, then the provided value should be used.
	Value T `json:"value" yaml:"value"`
}

// Infer returns the contained value if `Valid` is true; otherwise, it returns the provided fallback.
func (o Optional[T]) Infer(fallback T) T {
	if o.Valid {
		return o.Value
	}

	return fallback
}
