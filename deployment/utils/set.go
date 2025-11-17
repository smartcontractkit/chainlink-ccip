package utils

import (
	"iter"
	"maps"
)

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{m: map[T]struct{}{}}
}

func (s *Set[T]) Keys() iter.Seq[T] {
	return maps.Keys(s.m)
}

func (s *Set[T]) Has(value T) bool {
	_, exists := s.m[value]
	return exists
}

func (s *Set[T]) Add(value T) bool {
	exists := s.Has(value)
	s.m[value] = struct{}{}
	return exists
}
