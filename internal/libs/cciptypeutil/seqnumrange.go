package cciptypeutil

import "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

// SeqNumRangeLimit limits the range to n elements by truncating the end if necessary.
func SeqNumRangeLimit(rng ccipocr3.SeqNumRange, n uint64) ccipocr3.SeqNumRange {
	numElems := rng.End() - rng.Start() + 1
	if numElems <= 0 {
		return rng
	}

	if uint64(numElems) > n {
		newEnd := rng.Start() + ccipocr3.SeqNum(n) - 1
		if newEnd > rng.End() { // overflow - do nothing
			return rng
		}
		rng.SetEnd(newEnd)
	}

	return rng
}
