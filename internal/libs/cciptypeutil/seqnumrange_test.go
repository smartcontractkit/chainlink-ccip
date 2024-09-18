package cciptypeutil

import (
	"testing"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

func TestSeqNumRangeLimit(t *testing.T) {
	testCases := []struct {
		name string
		rng  ccipocr3.SeqNumRange
		n    uint64
		want ccipocr3.SeqNumRange
	}{
		{
			name: "no truncation",
			rng:  ccipocr3.NewSeqNumRange(0, 10),
			n:    11,
			want: ccipocr3.NewSeqNumRange(0, 10),
		},
		{
			name: "no truncation 2",
			rng:  ccipocr3.NewSeqNumRange(100, 110),
			n:    11,
			want: ccipocr3.NewSeqNumRange(100, 110),
		},
		{
			name: "truncation",
			rng:  ccipocr3.NewSeqNumRange(0, 10),
			n:    10,
			want: ccipocr3.NewSeqNumRange(0, 9),
		},
		{
			name: "truncation 2",
			rng:  ccipocr3.NewSeqNumRange(100, 110),
			n:    10,
			want: ccipocr3.NewSeqNumRange(100, 109),
		},
		{
			name: "empty",
			rng:  ccipocr3.NewSeqNumRange(0, 0),
			n:    0,
			want: ccipocr3.NewSeqNumRange(0, 0),
		},
		{
			name: "wrong range",
			rng:  ccipocr3.NewSeqNumRange(20, 15),
			n:    3,
			want: ccipocr3.NewSeqNumRange(20, 15),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := SeqNumRangeLimit(tc.rng, tc.n)
			if got != tc.want {
				t.Errorf("SeqNumRangeLimit(%v, %v) = %v; want %v", tc.rng, tc.n, got, tc.want)
			}
		})
	}
}
