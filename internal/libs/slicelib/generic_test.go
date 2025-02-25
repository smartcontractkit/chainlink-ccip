package slicelib

import (
	"testing"

	"github.com/stretchr/testify/assert"

	mapset "github.com/deckarep/golang-set/v2"
)

func TestCountUnique(t *testing.T) {
	testCases := []struct {
		name     string
		items    []string
		expCount int
	}{
		{
			name:     "empty slice",
			items:    []string{},
			expCount: 0,
		},
		{
			name:     "no duplicate",
			items:    []string{"a", "b", "c"},
			expCount: 3,
		},
		{
			name:     "with duplicate",
			items:    []string{"a", "a", "b", "c", "b"},
			expCount: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expCount, CountUnique(tc.items))
		})
	}
}

func TestFlatten(t *testing.T) {
	testCases := []struct {
		name       string
		slices     [][]int
		expFlatten []int
	}{
		{
			name:       "empty slice",
			slices:     [][]int{},
			expFlatten: []int{},
		},
		{
			name:       "no duplicate",
			slices:     [][]int{{1, 2}, {3, 4}, {5, 6}},
			expFlatten: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:       "with duplicate",
			slices:     [][]int{{1, 2}, {1, 2}, {3, 4}, {5, 6}},
			expFlatten: []int{1, 2, 1, 2, 3, 4, 5, 6},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expFlatten, Flatten(tc.slices))
		})
	}
}

func TestFilter(t *testing.T) {
	type person struct {
		id   string
		name string
		age  int
	}

	testCases := []struct {
		name       string
		items      []person
		valid      func(person) bool
		expResults []person
	}{
		{
			name:       "empty slice",
			items:      []person{},
			valid:      func(p person) bool { return p.age > 20 },
			expResults: []person{},
		},
		{
			name: "no valid item",
			items: []person{
				{id: "1", name: "Alice", age: 18},
				{id: "2", name: "Bob", age: 20},
				{id: "3", name: "Charlie", age: 19},
			},
			valid:      func(p person) bool { return p.age > 20 },
			expResults: []person{},
		},
		{
			name: "with valid item",
			items: []person{
				{id: "1", name: "Alice", age: 18},
				{id: "2", name: "Bob", age: 25},
				{id: "3", name: "Charlie", age: 19},
			},
			valid: func(p person) bool { return p.age > 20 },
			expResults: []person{
				{id: "2", name: "Bob", age: 25},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expResults, Filter(tc.items, tc.valid))
		})
	}
}

func TestMap(t *testing.T) {
	type person struct {
		name string
	}

	testCases := []struct {
		name       string
		items      []person
		valid      func(person) string
		expResults []string
	}{
		{
			name:       "empty slice",
			items:      []person{},
			valid:      func(p person) string { return p.name },
			expResults: []string{},
		},
		{
			name: "no valid item",
			items: []person{
				{name: "Alice"},
				{name: "Bob"},
				{name: "Charlie"},
			},
			valid:      func(p person) string { return p.name },
			expResults: []string{"Alice", "Bob", "Charlie"},
		},
		{
			name: "with nonsense",
			items: []person{
				{name: "Alice"},
				{name: "Bob"},
				{name: "Charlie"},
			},
			valid:      func(p person) string { return "nonsense" },
			expResults: []string{"nonsense", "nonsense", "nonsense"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expResults, Map(tc.items, tc.valid))
		})
	}
}

func TestToSortedSlice(t *testing.T) {
	testCases := []struct {
		name        string
		input       []int64
		expElements []int64
	}{
		{
			name:        "empty set",
			input:       []int64{},
			expElements: []int64{},
		},
		{
			name:        "single element",
			input:       []int64{1},
			expElements: []int64{1},
		},
		{
			name:        "multiple elements unsorted",
			input:       []int64{3, 1, 4, 2},
			expElements: []int64{1, 2, 3, 4},
		},
		{
			name:        "duplicate elements",
			input:       []int64{3, 1, 3, 2, 1},
			expElements: []int64{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			set := mapset.NewSet[int64]()
			for _, v := range tc.input {
				set.Add(v)
			}
			assert.Equal(t, tc.expElements, ToSortedSlice(set))
		})
	}
}
