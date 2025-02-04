package consensus

import (
	"fmt"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/commontypes"
	"golang.org/x/crypto/sha3"
)

type observersCounter[T any] struct {
	data      T
	observers map[commontypes.OracleID]struct{}
}

// OracleMinObservation provides a way to ensure a minimum number of observations for
// some piece of data have occurred. It maintains an internal cache and provides a list
// of valid or invalid data points.
type OracleMinObservation[T any] interface {
	Add(data T, oracleID commontypes.OracleID)
	GetValid() []T
}

// minObservation is a helper object to filter data based on observation counts.
// It keeps track of all inputs, determines if they are consistent
// with one another, and whether they meet the required count threshold.
type oracleMinObservation[T any] struct {
	minObservation Threshold
	cache          map[cciptypes.Bytes32]*observersCounter[T]
	idFunc         func(T) [32]byte
}

// NewOracleMinObservation constructs a concrete MinObservation object. The
// supplied idFunc is used to generate a uniqueID for the type being observed.
func NewOracleMinObservation[T any](minThreshold Threshold, idFunc func(T) [32]byte) OracleMinObservation[T] {
	if idFunc == nil {
		idFunc = func(data T) [32]byte {
			return sha3.Sum256([]byte(fmt.Sprintf("%v", data)))
		}
	}
	return &oracleMinObservation[T]{
		minObservation: minThreshold,
		cache:          make(map[cciptypes.Bytes32]*observersCounter[T]),
		idFunc:         idFunc,
	}
}

func (cv *oracleMinObservation[T]) Add(data T, oracleID commontypes.OracleID) {
	id := cv.idFunc(data)
	if _, ok := cv.cache[id]; ok {
		cv.cache[id].observers[oracleID] = struct{}{}
	} else {
		cv.cache[id] = &observersCounter[T]{data: data, observers: make(map[commontypes.OracleID]struct{})}
	}
}

func (cv *oracleMinObservation[T]) GetValid() []T {
	var validated []T
	for _, rc := range cv.cache {
		if len(rc.observers) >= int(cv.minObservation) {
			rc := rc
			validated = append(validated, rc.data)
		}
	}
	return validated
}
