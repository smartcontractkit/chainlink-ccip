package consensus

// Threshold is the minimum number of observations required for a given key.
type Threshold uint

// MultiThreshold is an interface that maps a key to a threshold.
type MultiThreshold[T any] interface {
	// Get the minimum number of observations required for a given key.
	Get(key T) (Threshold, bool)
}

// TwoFPlus1 is a common threshold mapping function that returns 2*f + 1.
// See ocr3types.QuorumTwoFPlusOne - guarantees an honest majority of observations
func TwoFPlus1(f int) Threshold {
	return Threshold(2*f + 1)
}

// FPlus1 is a common threshold mapping function that returns f + 1.
// See ocr3types.QuorumFPlusOne - Guarantees at least one honest observation
func FPlus1(f int) Threshold {
	return Threshold(f + 1)
}

func F(f int) Threshold { return Threshold(f) }

// MakeMultiThreshold constructs a threshold that maps each key to a threshold value.
func MakeMultiThreshold[K comparable](fChain map[K]int, mapping func(int) Threshold) MultiThreshold[K] {
	thresh := make(multiThreshold[K])
	for key, f := range fChain {
		thresh[key] = mapping(f)
	}
	return thresh
}

type multiThreshold[K comparable] map[K]Threshold

func (mt multiThreshold[K]) Get(key K) (Threshold, bool) {
	thresh, ok := mt[key]
	return thresh, ok
}

// MakeConstantThreshold constructs a threshold that maps each key to a constant threshold value.
func MakeConstantThreshold[T any](f Threshold) MultiThreshold[T] {
	return constantThreshold[T](f)
}

type constantThreshold[T any] uint

func (ct constantThreshold[T]) Get(_ T) (Threshold, bool) {
	return Threshold(ct), true
}

func GteFPlusOne(f, value int) bool { return value >= f+1 }

func LtFPlusOne(f, value int) bool { return value < f+1 }

func LtTwoFPlusOne(f, value int) bool { return value < 2*f+1 }
