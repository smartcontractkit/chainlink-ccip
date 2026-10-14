package utils

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFractionalDigits(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  int
	}{
		{name: "two places", value: 0.29, want: 2},
		{name: "scientific notation", value: 1e-9, want: 9},
		{name: "integer", value: 100, want: 0},
		{name: "zero", value: 0, want: 0},
		{name: "nine places", value: 1.123456789, want: 9},
		{name: "NaN", value: math.NaN(), want: 0},
		{name: "+Inf", value: math.Inf(1), want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, FractionalDigits(tt.value))
		})
	}
}
