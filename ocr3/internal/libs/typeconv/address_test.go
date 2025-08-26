package typconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeepNRightBytes(t *testing.T) {
	b := []byte{0x01, 0x02, 0x03, 0x04}
	assert.Equal(t, []byte{0x03, 0x04}, KeepNRightBytes(b, 2))
	assert.Equal(t, []byte{}, KeepNRightBytes(b, 0))
	assert.Equal(t, []byte{0x01, 0x02, 0x03, 0x04}, KeepNRightBytes(b, 21))
}
