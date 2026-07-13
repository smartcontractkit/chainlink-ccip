package adapters

import (
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
	assert.Equal(t, common.LeftPadBytes(common.HexToAddress(evmAddr).Bytes(), 32), out[0])
	assert.Equal(t, common.FromHex(cantonAddr), out[1])
}

func TestParseOffRampSourceOnRampAddresses_Dedupes(t *testing.T) {
	t.Parallel()

	addr := "0xa5f4d6b956c610c147282e1c180fcd04cfbed6cf8a0244289a1be44c7e784330"
	out, err := parseOffRampSourceOnRampAddresses([]string{addr, addr})
	require.NoError(t, err)
	require.Len(t, out, 1)
}

func TestParseOffRampSourceOnRampAddresses_InvalidLength(t *testing.T) {
	t.Parallel()

	_, err := parseOffRampSourceOnRampAddresses([]string{"0x0102"})
	require.Error(t, err)
}
