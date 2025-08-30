package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveIthElement(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		index    int
		expected []int
	}{
		{"Remove middle element", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 4, 5}},
		{"Remove first element", []int{1, 2, 3, 4, 5}, 0, []int{2, 3, 4, 5}},
		{"Remove last element", []int{1, 2, 3, 4, 5}, 4, []int{1, 2, 3, 4}},
		{"Index out of bounds (negative)", []int{1, 2, 3, 4, 5}, -1, []int{1, 2, 3, 4, 5}},
		{"Index out of bounds (too large)", []int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{"Single element slice", []int{1}, 0, []int{}},
		{"Empty slice", []int{}, 0, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveIthElement(tt.slice, tt.index)
			assert.Equal(t, tt.expected, result)
		})
	}
}
