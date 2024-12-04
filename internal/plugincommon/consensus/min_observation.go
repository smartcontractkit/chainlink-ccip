package consensus

import (
	"fmt"

	"golang.org/x/crypto/sha3"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

type counter[T any] struct {
	data  T
	count uint
}

// MinObservation provides a way to ensure a minimum number of observations for
// some piece of data have occurred. It maintains an internal cache and provides a list
// of valid or invalid data points.
type MinObservation[T any] interface {
	Add(data T)
	GetValid() []T
}

// minObservation is a helper object to filter data based on observation counts.
// It keeps track of all inputs, determines if they are consistent
// with one another, and whether they meet the required count threshold.
type minObservation[T any] struct {
	minObservation Threshold
	cache          map[cciptypes.Bytes32]*counter[T]
	idFunc         func(T) [32]byte
}

// NewMinObservation constructs a concrete MinObservation object. The
// supplied idFunc is used to generate a uniqueID for the type being observed.
func NewMinObservation[T any](minThreshold Threshold, idFunc func(T) [32]byte) MinObservation[T] {
	if idFunc == nil {
		idFunc = func(data T) [32]byte {
			return sha3.Sum256([]byte(fmt.Sprintf("%v", data)))
		}
	}
	return &minObservation[T]{
		minObservation: minThreshold,
		cache:          make(map[cciptypes.Bytes32]*counter[T]),
		idFunc:         idFunc,
	}
}

func (cv *minObservation[T]) Add(data T) {
	id := cv.idFunc(data)
	if _, ok := cv.cache[id]; ok {
		cv.cache[id].count++
	} else {
		cv.cache[id] = &counter[T]{data: data, count: 1}
	}
}

func (cv *minObservation[T]) GetValid() []T {
	var validated []T
	for _, rc := range cv.cache {
		if rc.count >= uint(cv.minObservation) {
			rc := rc
			validated = append(validated, rc.data)
		}
	}
	return validated
}
