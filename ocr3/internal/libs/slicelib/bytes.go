package slicelib

// LeftPadBytes zero-pads slice to the left up to length l.
// Cribbed from go-ethereum/common/bytes.go
func LeftPadBytes(slice []byte, l int) []byte {
	if l <= len(slice) {
		return slice
	}

	padded := make([]byte, l)
	copy(padded[l-len(slice):], slice)

	return padded
}
