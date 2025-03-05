package typconv

// KeepNRightBytes returns the last n bytes of the given byte slice.
// Example: KeepNRightBytes([]byte{0x01, 0x02, 0x03, 0x04}, 2) -> []byte{0x03, 0x04}
func KeepNRightBytes(b []byte, n uint) []byte {
	if n >= uint(len(b)) {
		return b
	}
	return b[uint(len(b))-n:]
}
