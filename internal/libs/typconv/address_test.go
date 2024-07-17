package typconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddressBytesToString(t *testing.T) {
	addr := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F}
	want := "0x000102030405060708090a0b0c0d0e0f"
	got := AddressBytesToString(addr)
	assert.Equal(t, want, got)
}
