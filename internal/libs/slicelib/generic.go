package slicelib

import (
	"sort"

	mapset "github.com/deckarep/golang-set/v2"
)

// CountUnique counts the unique items of the provided slice.
func CountUnique[T comparable](items []T) int {
	m := make(map[T]struct{})
	for _, item := range items {
		m[item] = struct{}{}
	}
	return len(m)
}

// Flatten flattens a slice of slices into a single slice.
func Flatten[T any](slices [][]T) []T {
	res := make([]T, 0)
	for _, s := range slices {
		res = append(res, s...)
	}
	return res
}

func Filter[T any](slice []T, valid func(T) bool) []T {
	res := make([]T, 0, len(slice))
	for _, item := range slice {
		if valid(item) {
			res = append(res, item)
		}
	}
	return res
}

func Map[T any, T2 any](slice []T, mapper func(T) T2) []T2 {
	res := make([]T2, len(slice))
	for i, item := range slice {
		res[i] = mapper(item)
	}
	return res
}

func ToSortedSlice[T ~int | ~int64 | ~uint64](set mapset.Set[T]) []T {
	elements := set.ToSlice()
	sort.Slice(elements, func(i, j int) bool {
		return elements[i] < elements[j]
	})
	return elements
}
