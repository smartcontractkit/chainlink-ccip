package utils

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// Optional represents a value that may or may not be explicitly set. This
// is useful in scenarios where a changeset only needs to partially update
// a subset of fields in a large configuration struct. In YAML files, it's
// not strictly necessary to define both `value` and `valid` keys; instead
// see optional_test.go for the supported YAML shapes and their semantics.
type Optional[T any] struct {
	// Valid indicates whether the provided value should be used. If this is
	// false (the default), then the provided value will be ignored.
	Valid bool `json:"valid" yaml:"valid"`

	// This only has an effect when `Valid` is set to true. If this is the
	// case, then the provided value should be used.
	Value T `json:"value" yaml:"value"`
}

// NewOptional creates a new Optional with the provided value and sets Valid to true.
func NewOptional[T any](value T) Optional[T] {
	return Optional[T]{
		Value: value,
		Valid: true,
	}
}

// GetOrDefault returns the contained value if `Valid` is true; otherwise, it returns the provided fallback.
func (o Optional[T]) GetOrDefault(fallback T) T {
	if o.Valid {
		return o.Value
	}

	return fallback
}

// Get returns the contained value and a boolean indicating whether it is valid.
func (o Optional[T]) Get() (T, bool) {
	return o.Value, o.Valid
}

// UnmarshalYAML supports three YAML representations:
//   - Very-Verbose: `field: {value: 32, valid: false}` → backwards-compatible explicit control
//   - Semi-verbose: `field: {value: 32}` → {Value: 32, Valid: true} (inferred)
//   - Shorthand: `field: 32` → {Value: 32, Valid: true}
func (o *Optional[T]) UnmarshalYAML(node *yaml.Node) error {
	// Explicit null (`field: ~` or `field:`) — reset to zero value with Valid=false
	if node.Kind == yaml.ScalarNode && node.Tag == "!!null" {
		var zero T
		o.Value = zero
		o.Valid = false
		return nil
	}

	// Shorthand: `field: <scalar>` — infer value and valid=true
	if node.Kind == yaml.ScalarNode {
		if err := node.Decode(&o.Value); err != nil {
			return err
		}
		o.Valid = true
		return nil
	}

	// If we're NOT using the shorthand, then we should expect a mapping node
	if node.Kind != yaml.MappingNode {
		return fmt.Errorf("optional: expected scalar or mapping, got node kind %v", node.Kind)
	}

	// Scan keys present in the mapping
	hasValid, hasValue := false, false
	for i := 0; i < len(node.Content)-1; i += 2 {
		switch node.Content[i].Value {
		case "valid":
			hasValid = true
		case "value":
			hasValue = true
		}
	}

	// Decode via a local anonymous struct to avoid infinite recursion
	// (a named type alias doesn't work cleanly with generics in Go)
	var raw struct {
		Valid bool `yaml:"valid"`
		Value T    `yaml:"value"`
	}
	if err := node.Decode(&raw); err != nil {
		return err
	}

	// Infer whether `Valid` should be true or false based on which keys were present in the mapping
	o.Value = raw.Value
	switch {
	case hasValid:
		o.Valid = raw.Valid // explicit — backwards-compatible
	case hasValue:
		o.Valid = true // value present, no valid key → infer true
	default:
		o.Valid = false // empty mapping → zero value, Valid stays false
	}

	return nil
}
