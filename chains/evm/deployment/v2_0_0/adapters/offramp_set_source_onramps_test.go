package adapters

import (
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseOffRampSourceOnRampAddresses(t *testing.T) {
	t.Parallel()

	evmAddr := "0x0000000000000000000000005e47fdcf6d4a4529424cbdfdddd33b08a3da5faa"
	cantonAddr := "0xa5f4d6b956c610c147282e1c180fcd04cfbed6cf8a0244289a1be44c7e784330"

	out, err := parseOffRampSourceOnRampAddresses([]string{evmAddr, cantonAddr})
	require.NoError(t, err)
	require.Len(t, out, 2)
	assert.Equal(t, common.FromHex(evmAddr), out[0])
	assert.Equal(t, common.FromHex(cantonAddr), out[1])
}

// TestParseOffRampSourceOnRampAddresses_NoPadding pins the encoding contract: whatever the
// operator supplies is stored as-is. The source chain owns its encoding, so this adapter
// must not reshape it to a length that happens to suit EVM.
func TestParseOffRampSourceOnRampAddresses_NoPadding(t *testing.T) {
	t.Parallel()

	short := "0x5e47fdcf6d4a4529424cbdfdddd33b08a3da5faa"
	out, err := parseOffRampSourceOnRampAddresses([]string{short})
	require.NoError(t, err)
	require.Len(t, out, 1)
	assert.Len(t, out[0], 20, "a 20-byte input must not be widened to 32")
	assert.Equal(t, common.FromHex(short), out[0])
}

func TestParseOffRampSourceOnRampAddresses_Dedupes(t *testing.T) {
	t.Parallel()

	addr := "0xa5f4d6b956c610c147282e1c180fcd04cfbed6cf8a0244289a1be44c7e784330"
	out, err := parseOffRampSourceOnRampAddresses([]string{addr, addr})
	require.NoError(t, err)
	require.Len(t, out, 1)
}

func TestParseOffRampSourceOnRampAddresses_Invalid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		addr string
	}{
		{name: "empty", addr: "0x"},
		{name: "not hex", addr: "nope"},
		{name: "missing 0x prefix", addr: "a5f4d6b9"},
		// The wire format length-prefixes the onramp address with a uint8.
		{name: "over 255 bytes", addr: "0x" + strings.Repeat("ab", 256)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := parseOffRampSourceOnRampAddresses([]string{test.addr})
			require.Error(t, err)
		})
	}
}
